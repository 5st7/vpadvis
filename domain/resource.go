package domain

// ResourceRequest contains CPU and Memory specifications.
type ResourceRequest struct {
	CPU    string
	Memory string
}

// Resource defines the CPU and Memory settings for requests and limits.
type Resource struct {
	Request ResourceRequest
	Limit   ResourceRequest
}

// ContainerResource holds the current and recommended resource configurations for each container.
type ContainerResource struct {
	ContainerName string
	Current       Resource
	Recommended   Resource
}

// WorkloadResource represents the resource settings for an entire workload (e.g., Deployment).
type WorkloadResource struct {
	WorkloadName       string
	ContainerResources []ContainerResource
}
