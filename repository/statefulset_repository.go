package repository

import (
	"context"

	"github.com/5st7/vpadvis/domain"

	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// StatefulSetRepository interface for retrieving resource configurations from StatefulSets.
type StatefulSetRepository interface {
	GetStatefulSetResources(statefulSetName, namespace string) (*domain.WorkloadResource, error)
}

type statefulSetRepository struct {
	client client.Client
}

func NewStatefulSetRepository(client client.Client) StatefulSetRepository {
	return &statefulSetRepository{client: client}
}

func (r *statefulSetRepository) GetStatefulSetResources(statefulSetName, namespace string) (*domain.WorkloadResource, error) {
	statefulSet := &appsv1.StatefulSet{}
	err := r.client.Get(context.Background(), client.ObjectKey{Name: statefulSetName, Namespace: namespace}, statefulSet)
	if err != nil {
		return nil, err
	}

	containerResources := []domain.ContainerResource{}
	for _, container := range statefulSet.Spec.Template.Spec.Containers {
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
		WorkloadName:       statefulSetName,
		ContainerResources: containerResources,
	}, nil
}
