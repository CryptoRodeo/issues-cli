package formatter

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/CryptoRodeo/issues-cli/pkg/models"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// Colors for severity levels
var (
	criticalColor = color.New(color.FgRed, color.Bold).SprintFunc()
	majorColor    = color.New(color.FgYellow, color.Bold).SprintFunc()
	minorColor    = color.New(color.FgBlue).SprintFunc()
	infoColor     = color.New(color.FgGreen).SprintFunc()
	boldColor     = color.New(color.Bold).SprintFunc()
)

// GetSeverityColor returns the colored string for a severity level
func GetSeverityColor(severity string) string {
	switch strings.ToLower(severity) {
	case "critical":
		return criticalColor(severity)
	case "major":
		return majorColor(severity)
	case "minor":
		return minorColor(severity)
	case "info":
		return infoColor(severity)
	default:
		return severity
	}
}

// PrintIssuesTable prints a table of issues
func PrintIssuesTable(issues []models.Issue) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Type", "Severity", "State", "Detected"})

	// Set minimum width for each column
	table.SetColMinWidth(0, 36) // ID width (UUID)
	table.SetColMinWidth(1, 40) // Title width
	table.SetColMinWidth(2, 12) // Type width
	table.SetColMinWidth(3, 12) // Severity width
	table.SetColMinWidth(4, 12) // State width
	table.SetColMinWidth(5, 30) // Detected width

	table.SetAutoWrapText(true)
	table.SetRowLine(true)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("-")
	table.SetHeaderLine(true)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	for _, issue := range issues {
		detectedAt := issue.DetectedAt.Format("2006-01-02 15:04:05")
		id := issue.ID

		table.Append([]string{
			id,
			issue.Title,
			issue.IssueType,
			GetSeverityColor(issue.Severity),
			issue.State,
			detectedAt,
		})
	}

	table.Render()
	fmt.Printf("\nFound %d issue(s)\n", len(issues))
}

// PrintIssueDetails prints detailed information about an issue
func PrintIssueDetails(issue *models.Issue) {
	fmt.Println()
	fmt.Println(boldColor("Issue Details:"))
	fmt.Printf("%s: %s\n", boldColor("ID"), issue.ID)
	fmt.Printf("%s: %s\n", boldColor("Title"), issue.Title)
	fmt.Printf("%s:\n%s\n", boldColor("Description"), issue.Description)
	fmt.Printf("%s: %s\n", boldColor("Type"), issue.IssueType)
	fmt.Printf("%s: %s\n", boldColor("Severity"), GetSeverityColor(issue.Severity))
	fmt.Printf("%s: %s\n", boldColor("State"), issue.State)
	fmt.Printf("%s: %s\n", boldColor("Detected At"), formatTime(issue.DetectedAt))

	if issue.ResolvedAt != nil {
		fmt.Printf("%s: %s\n", boldColor("Resolved At"), formatTime(*issue.ResolvedAt))
	}

	fmt.Println()
	fmt.Println(boldColor("Scope:"))
	fmt.Printf("%s: %s\n", boldColor("Type"), issue.Scope.ResourceType)
	fmt.Printf("%s: %s\n", boldColor("Name"), issue.Scope.ResourceName)
	fmt.Printf("%s: %s\n", boldColor("Namespace"), issue.Scope.ResourceNamespace)

	if len(issue.Links) > 0 {
		fmt.Println()
		fmt.Println(boldColor("Links:"))
		for _, link := range issue.Links {
			blue := color.New(color.FgBlue).SprintFunc()
			fmt.Printf("• %s: %s\n", blue(link.Title), link.URL)
		}
	}

	if len(issue.RelatedFrom) > 0 {
		fmt.Println()
		fmt.Println(boldColor("Related Issues:"))
		for _, related := range issue.RelatedFrom {
			if related.Target != nil {
				fmt.Printf("• %s: %s\n", related.Target.ID, related.Target.Title)
			}
		}
	}
}

// Helper function to format time
func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
