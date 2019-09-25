/*
 Copyright (C) 2019 ComplianceKit Contributors. See LICENSE.md for license.
*/

package main

import (
	"os"

	"github.com/complianceascode/compliancekit/pkg/cmd/ckit"
)

func main() {
	cmd := ckit.ComplianceKitRootCommand(os.Stdin, os.Stdout, os.Stderr)
	cmd.Execute()
}
