package repository

import (
	"context"

	"github.com/5st7/vpadvis/domain"

	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DaemonSetRepository interface for retrieving resource configurations from DaemonSets.
type DaemonSetRepository interface {
	GetDaemonSetResources(daemonSetName, namespace string) (*domain.WorkloadResource, error)
}

type daemonSetRepository struct {
	client client.Client
}

func NewDaemonSetRepository(client client.Client) DaemonSetRepository {
	return &daemonSetRepository{client: client}
}

func (r *daemonSetRepository) GetDaemonSetResources(daemonSetName, namespace string) (*domain.WorkloadResource, error) {
	daemonSet := &appsv1.DaemonSet{}
	err := r.client.Get(context.Background(), client.ObjectKey{Name: daemonSetName, Namespace: namespace}, daemonSet)
	if err != nil {
		return nil, err
	}

	containerResources := []domain.ContainerResource{}
	for _, container := range daemonSet.Spec.Template.Spec.Containers {
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
		WorkloadName:       daemonSetName,
		ContainerResources: containerResources,
	}, nil
}
