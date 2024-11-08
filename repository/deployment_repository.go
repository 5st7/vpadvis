package repository

import (
	"context"
	"github.com/5st7/vpadvis/domain"

	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DeploymentRepository interface for retrieving resource configurations from Deployments.
type DeploymentRepository interface {
	GetDeploymentResources(deploymentName, namespace string) (*domain.WorkloadResource, error)
}

type deploymentRepository struct {
	client client.Client
}

func NewDeploymentRepository(client client.Client) DeploymentRepository {
	return &deploymentRepository{client: client}
}

func (r *deploymentRepository) GetDeploymentResources(deploymentName, namespace string) (*domain.WorkloadResource, error) {
	deployment := &appsv1.Deployment{}
	err := r.client.Get(context.Background(), client.ObjectKey{Name: deploymentName, Namespace: namespace}, deployment)
	if err != nil {
		return nil, err
	}

	containerResources := []domain.ContainerResource{}
	for _, container := range deployment.Spec.Template.Spec.Containers {
		containerResources = append(containerResources, domain.ContainerResource{
			ContainerName: container.Name,
			Current: domain.Resource{
				Request: domain.ResourceRequest{
					CPU:    container.Resources.Requests.Cpu().String(),
					Memory: container.Resources.Requests.Memory().String(),
				},
				Limit: domain.ResourceRequest{
					CPU:    container.Resources.Limits.Cpu().String(),
					Memory: container.Resources.Limits.Memory().String(),
				},
			},
		})
	}

	return &domain.WorkloadResource{
		WorkloadName:       deploymentName,
		ContainerResources: containerResources,
	}, nil
}
