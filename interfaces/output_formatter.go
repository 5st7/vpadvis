package interfaces

import (
	"github.com/5st7/vpadvis/domain"
	"github.com/docker/go-units"
)

// OutputFormatter defines an interface for outputting recommendations in different formats.
type OutputFormatter interface {
	PrintAllRecommendations([]domain.WorkloadResource)
}

// NewFormatter creates an OutputFormatter based on the format type.
func NewFormatter(formatType string) OutputFormatter {
	switch formatType {
	case "markdown":
		return &MarkdownFormatter{}
	default:
		return &MarkdownFormatter{}
	}
}
func displayOrNull(value string) string {
	if value == "0" || value == "" {
		return "null"
	}
	return value
}
func convertMemory(value string) string {
	if value == "0" || value == "" {
		return "null"
	}
	sizeInBytes, err := units.RAMInBytes(value)
	if err != nil {
		return "invalid"
	}
	return units.HumanSize(float64(sizeInBytes))
}
