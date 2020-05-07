package update

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/complianceascode/compliancekit/internal/clifactory"
	v1 "github.com/complianceascode/compliancekit/pkg/api/v1"
	"github.com/complianceascode/compliancekit/pkg/cmd/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	name string
)

// NewCmdUpdateRule creates the compliance documentation in Gitbook format.
func NewCmdUpdateRule(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rule",
		Short: "create Extensible Configuration Checklist Description Format (XCCDF) rule",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunCreateRule(out, cmd, args)
			errors.CheckError(err)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "rule name")

	return cmd
}

// RunCreateRule runs export when specified in cli
func RunCreateRule(out io.Writer, cmd *cobra.Command, args []string) error {
	dir := "/home/ralford/development/testing/scap-security-guide/linux_os/guide/system/auditing"
	name := clifactory.CliReader("Rule name", name)
	/*if len(args) > 0 {
		name = args[0]
	}

	rule := new(v1.Rule)
	rule.ID = clifactory.CliReader("Rule name", name, clifactory.ReaderConfig{Regex: "^[a-zA-Z_\\-]+$"})
	rule.Title = clifactory.CliReader("Rule title", title)
	rule.Description = clifactory.CliReader("Rule description", desc, clifactory.ReaderConfig{Multiline: true})
	rule.Rationale = clifactory.CliReader("Rule rationale", rationale)
	rule.Severity = clifactory.CliReader("Rule severity", strings.ToLower(severity))

	d, err := yaml.Marshal(&rule)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
	*/

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if info.Name() == name {
			path += "/rule.yml"
			fmt.Println(path)
		}
		if !info.IsDir() && info.Name() == "rule.yml" {

			//fmt.Println(path)
			f, err := os.Open(path)
			if err != nil {
				err = err
			}

			var cfg v1.Rule
			decoder := yaml.NewDecoder(f)
			err = decoder.Decode(&cfg)
			if err != nil {
				err = err
			}
			//fmt.Println(cfg)

		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
		return err
	}

	return nil
}
