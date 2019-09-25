package create

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	v1 "github.com/complianceascode/compliancekit/pkg/api/v1"
	"github.com/complianceascode/compliancekit/pkg/cmd/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	name, desc, rationale, severity, complete, title string
)

// NewCmdDocsGitBook creates the compliance documentation in Gitbook format.
func NewCmdCreateRule(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rule",
		Short: "create Extensible Configuration Checklist Description Format (XCCDF) rule",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunCreateRule(out, cmd, args)
			errors.CheckError(err)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "rule name")
	cmd.Flags().StringVarP(&desc, "description", "d", "", "rule description")
	cmd.Flags().StringVarP(&rationale, "rationale", "r", "", "why should this rule be implemented")
	cmd.Flags().StringVarP(&severity, "severity", "s", "unknown", "confidentiality, integrity, availability impact of rule")
	cmd.Flags().StringVarP(&title, "title", "t", "", "rule title")

	return cmd
}

// RunCreateRule runs export when specified in cli
func RunCreateRule(out io.Writer, cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		cmd.MarkFlagRequired("name")
	}

	rule := new(v1.Rule)
	rule.ID = CliReader("Rule name", name)
	rule.Description = desc
	rule.Rationale = rationale
	rule.Severity = strings.ToLower(severity)
	rule.Title = title

	d, err := yaml.Marshal(&rule)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))

	return nil
}

// CliReader handles getting input from command line interactively
func CliReader(msg string, optVal string) string {
	if optVal == "" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf(" %s: ", msg)
		for scanner.Scan() {
			return scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
	return optVal
}
