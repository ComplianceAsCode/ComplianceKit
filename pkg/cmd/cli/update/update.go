/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package update

import (
	"io"

	"github.com/spf13/cobra"
)

// NewCmdUpdate creates the compliance documentation.
func NewCmdUpdate(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update compliance components",
	}
	cmd.AddCommand(NewCmdUpdateRule(out))
	return cmd
}
