package domain

type ContainerRecommendation struct {
	ContainerName  string          `json:"containerName"`
	LowerBound     ResourceRequest `json:"lowerBound"`
	Target         ResourceRequest `json:"target"`
	UncappedTarget ResourceRequest `json:"uncappedTarget"`
	UpperBound     ResourceRequest `json:"upperBound"`
}

type VPARecommendation struct {
	DeploymentName           string                    `json:"deploymentName"`
	ContainerRecommendations []ContainerRecommendation `json:"containerRecommendations"`
}
