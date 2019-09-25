/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package create

import (
	"io"

	"github.com/spf13/cobra"
)

// NewCmdDocs creates the compliance documentation.
func NewCmdCreate(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create compliance components",
	}
	cmd.AddCommand(NewCmdCreateRule(out))
	return cmd
}
