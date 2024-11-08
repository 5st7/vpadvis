package repository

import (
	"context"
	"fmt"

	"github.com/5st7/vpadvis/domain"

	vpav1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type VPARepository interface {
	GetAllRecommendedResources(namespace string) ([]domain.WorkloadResource, error)
}

type vpaRepository struct {
	client          client.Client
	deploymentRepo  DeploymentRepository
	statefulSetRepo StatefulSetRepository
	daemonSetRepo   DaemonSetRepository
}

func NewVPARepository(client client.Client, deploymentRepo DeploymentRepository, statefulSetRepo StatefulSetRepository, daemonSetRepo DaemonSetRepository) VPARepository {
	return &vpaRepository{
		client:          client,
		deploymentRepo:  deploymentRepo,
		statefulSetRepo: statefulSetRepo,
		daemonSetRepo:   daemonSetRepo,
	}
}

func (r *vpaRepository) GetAllRecommendedResources(namespace string) ([]domain.WorkloadResource, error) {
	// ネームスペース内のすべてのVPAリソースを取得
	vpaList := &vpav1.VerticalPodAutoscalerList{}
	if err := r.client.List(context.Background(), vpaList, client.InNamespace(namespace)); err != nil {
		return nil, err
	}

	var allWorkloadResources []domain.WorkloadResource

	for _, vpa := range vpaList.Items {
		targetRef := vpa.Spec.TargetRef
		if targetRef == nil {
			continue
		}

		var workloadResource *domain.WorkloadResource
		var err error
		switch targetRef.Kind {
		case "Deployment":
			workloadResource, err = r.deploymentRepo.GetDeploymentResources(targetRef.Name, namespace)
		case "StatefulSet":
			workloadResource, err = r.statefulSetRepo.GetStatefulSetResources(targetRef.Name, namespace)
		case "DaemonSet":
			workloadResource, err = r.daemonSetRepo.GetDaemonSetResources(targetRef.Name, namespace)
		default:
			err = fmt.Errorf("unsupported workload type: %s", targetRef.Kind)
		}

		if err != nil {
			return nil, err
		}

		// VPAの推奨値をworkloadResourceに追加
		for _, container := range vpa.Status.Recommendation.ContainerRecommendations {
			for i, resource := range workloadResource.ContainerResources {
				if container.ContainerName == resource.ContainerName {
					workloadResource.ContainerResources[i].Recommended = domain.Resource{
						Request: domain.ResourceRequest{
							CPU:    container.Target.Cpu().String(),
							Memory: container.Target.Memory().String(),
						},
						Limit: domain.ResourceRequest{
							CPU:    container.UpperBound.Cpu().String(),
							Memory: container.UpperBound.Memory().String(),
						},
					}
				}
			}
		}

		allWorkloadResources = append(allWorkloadResources, *workloadResource)
	}

	return allWorkloadResources, nil
}
