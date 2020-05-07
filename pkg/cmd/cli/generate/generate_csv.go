package generate

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	v1 "github.com/complianceascode/compliancekit/pkg/api/v1"
	"github.com/complianceascode/compliancekit/pkg/cmd/errors"
	"github.com/spf13/cobra"
)

var (
	name, desc, rationale, severity, complete, title string
)

// NewCmdDocsGitBook creates the compliance documentation in Gitbook format.
func NewCmdGenerateCSV(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "csv",
		Short: "Create CSV file from XCCDF file",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunGenerateCSV(out, cmd, args)
			errors.CheckError(err)
		},
	}

	cmd.Flags().StringVarP(&name, "file-name", "n", "", "Generated file name")
	cmd.Flags().StringVarP(&severity, "stig-template", "s", "", "confidentiality, integrity, availability impact of rule")

	return cmd
}

// RunCreateRule runs export when specified in cli
func RunGenerateCSV(out io.Writer, cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		fmt.Println("Missing XCCDF file")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("File reading error", err)
		return err
	}
	xccdf := new(v1.Benchmark)
	xml.Unmarshal([]byte(data), &xccdf)

	/*output, err := xml.MarshalIndent(xccdf, "", "  ")

	if err != nil {

		fmt.Printf("error: %v\n", err)

	}*/
	xccdf.GetRules()
	//os.Stdout.Write(output)
	//fmt.Println(xccdf)
	/*
		for _, rule := range xccdf.GetProfileRules("stig") {
			fmt.Println(rule.Idref)
		}
	*/
	/*
		for i := range xccdf.Profiles {
			if xccdf.Profiles[i].Id == "stig" {
				for j, k := range xccdf.Profiles[i].Selects {
					for g := range xccdf.Groups {
						fmt.Println(g)
						if k.IdRef == g. {
							fmt.Println(k.Idref)
						}
					}
				}
			}
		}
	*/

	//fmt.Println(string(xccdf))
	//fmt.Println("Contents of file:", string(data))

	return nil
}
