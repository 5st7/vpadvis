package interfaces

import (
	"fmt"
	"github.com/5st7/vpadvis/domain"
	"os"
	"text/tabwriter"
)

type PlainTextFormatter struct{}

func (f *PlainTextFormatter) PrintAllRecommendations(workloadResources []domain.WorkloadResource) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Workload\tContainer\tCurrent Request CPU\tCurrent Request Memory\tCurrent Limit CPU\tCurrent Limit Memory\tRecommended CPU\tRecommended Memory")
	fmt.Fprintln(w, "--------\t---------\t--------------------\t-------------------\t---------------\t----------------\t---------------\t----------------")
	for _, workloadResource := range workloadResources {
		for _, container := range workloadResource.ContainerResources {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
				workloadResource.WorkloadName, container.ContainerName,
				displayOrNull(container.Current.Request.CPU), displayOrNull(container.Current.Request.Memory),
				displayOrNull(container.Current.Limit.CPU), displayOrNull(container.Current.Limit.Memory),
				displayOrNull(container.Recommended.Request.CPU), displayOrNull(convertMemory((container.Recommended.Request.Memory))),
			)
		}
	}
	w.Flush()
}
