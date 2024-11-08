package infrastructure

import (
	"log"

	vpav1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func NewK8sClient() (client.Client, error) {
	_ = vpav1.AddToScheme(scheme.Scheme)
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to load Kubernetes config: %v", err)
		return nil, err
	}
	return client.New(cfg, client.Options{Scheme: scheme.Scheme})
}
