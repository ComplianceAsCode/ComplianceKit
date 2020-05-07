/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package ckit

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/complianceascode/compliancekit/pkg/cmd/cli/create"
	"github.com/complianceascode/compliancekit/pkg/cmd/cli/generate"
	"github.com/complianceascode/compliancekit/pkg/cmd/cli/repo"
	"github.com/complianceascode/compliancekit/pkg/cmd/cli/update"
	cliversion "github.com/complianceascode/compliancekit/pkg/cmd/cli/version"
	"github.com/complianceascode/compliancekit/pkg/cmd/errors"
	"github.com/complianceascode/compliancekit/version"
	"github.com/spf13/cobra"
)

// Verbose boolean for turning on/off verbosity
var (
	Verbose       bool
	Version       bool
	usageTemplate = `{{if gt (len .Aliases) 0}}
Aliases:
{{.NameAndAliases}}{{end}}Usage:
  {{.CommandPath}} [global-options]{{if .HasAvailableSubCommands}} COMMAND{{end}}{{if .HasAvailableFlags}} [command-options]{{end}}{{if .HasExample}}
Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Options:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableLocalFlags}}

Options:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
See "{{.CommandPath}} help <TOPIC>" for more information on a specific topic.
See "{{.CommandPath}} <COMMAND> --help" for more information about a command.
{{end}}
`
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func name() string {
	return filepath.Base(os.Args[0])
}

// NewRootCommand Main ckit command cli
// Add new commands/subcommands for new verbs in this function
func ComplianceKitRootCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   name(),
		Short: "ComplianceKit CLI Tool",
		Long: `ComplianceKit is a command-line interface (CLI) that
allows users to easily construct compliance content`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := RunGlobalFlags(out, cmd)
			errors.CheckError(err)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmds.SetUsageTemplate(usageTemplate)
	cmds.SuggestionsMinimumDistance = 1
	cmds.ResetFlags()
	// Global Options
	cmds.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "Run in verbose mode")
	cmds.PersistentFlags().BoolVarP(&Version, "version", "v", false, "Print the version")

	// Add new main commands here
	cmds.AddCommand(create.NewCmdCreate(out))
	cmds.AddCommand(generate.NewCmdGenerate(out))
	cmds.AddCommand(repo.NewCmdGet(out))
	cmds.AddCommand(update.NewCmdUpdate(out))
	cmds.AddCommand(cliversion.NewCmdVersion(out))
	disableFlagsInUseLine(cmds)

	return cmds
}

// RunGlobalFlags runs global options when specified in cli
func RunGlobalFlags(out io.Writer, cmd *cobra.Command) error {
	flagVersion := Version
	flagVerbose := Verbose

	if flagVerbose {
		log.SetOutput(os.Stderr)
		log.Println("Running in verbose mode")
	}

	if flagVersion {
		version.PrintVersion()
	}

	return nil

}

// disableFlagsInUseLine do not add a `[flags]` to the end of the usage line.
func disableFlagsInUseLine(cmd *cobra.Command) {
	visitAll(cmd, func(cmds *cobra.Command) {
		cmds.DisableFlagsInUseLine = true
	})
}

// visitAll will traverse all commands from the root.
// This is different from the VisitAll of cobra.Command where only parents
// are checked.
func visitAll(cmds *cobra.Command, fn func(*cobra.Command)) {
	for _, cmd := range cmds.Commands() {
		visitAll(cmd, fn)
	}
	fn(cmds)
}
