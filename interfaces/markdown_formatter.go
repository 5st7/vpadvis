package interfaces

import (
	"fmt"
	"github.com/5st7/vpadvis/domain"
	"os"
	"text/tabwriter"
)

type MarkdownFormatter struct{}

func (f *MarkdownFormatter) PrintAllRecommendations(workloadResources []domain.WorkloadResource) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "| Workload | Container | Current Request CPU | Current Request Memory | Current Limit CPU | Current Limit Memory | Recommended CPU | Recommended Memory |")
	fmt.Fprintln(w, "|----------|-----------|---------------------|------------------------|-------------------|----------------------|-----------------|--------------------|")
	for _, workloadResource := range workloadResources {
		for _, container := range workloadResource.ContainerResources {
			fmt.Fprintf(w, "| %s | %s | %s | %s | %s | %s | %s | %s |\n",
				workloadResource.WorkloadName, container.ContainerName,
				displayOrNull(container.Current.Request.CPU), displayOrNull(container.Current.Request.Memory),
				displayOrNull(container.Current.Limit.CPU), displayOrNull(container.Current.Limit.Memory),
				displayOrNull(container.Recommended.Request.CPU), displayOrNull(convertMemory((container.Recommended.Request.Memory))))
		}
	}
	w.Flush()
}
