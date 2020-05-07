/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package generate

import (
	"io"

	"github.com/spf13/cobra"
)

// NewCmdDocs creates the compliance documentation.
func NewCmdGenerate(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate compliance content",
	}
	cmd.AddCommand(NewCmdGenerateCSV(out))
	return cmd
}
