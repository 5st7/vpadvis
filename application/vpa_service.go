package application

import (
	"github.com/5st7/vpadvis/domain"
	"github.com/5st7/vpadvis/repository"
)

type VPAService interface {
	GetAllVPARecommendations(namespace string) ([]domain.WorkloadResource, error)
}

type vpaService struct {
	vpaRepo repository.VPARepository
}

func NewVPAService(vpaRepo repository.VPARepository) VPAService {
	return &vpaService{vpaRepo: vpaRepo}
}

func (s *vpaService) GetAllVPARecommendations(namespace string) ([]domain.WorkloadResource, error) {
	return s.vpaRepo.GetAllRecommendedResources(namespace)
}
