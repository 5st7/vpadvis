package interfaces

import (
	"github.com/5st7/vpadvis/domain"
	"github.com/dustin/go-humanize"
	"strconv"
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
	case "plaintext":
		return &PlainTextFormatter{}
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
	mem, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return "valid"
	}
	return humanize.IBytes(mem)
}
