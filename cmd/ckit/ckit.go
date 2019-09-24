/*
 Copyright (C) 2019 ComplianceKit Contributors. See LICENSE.md for license.
*/

package main

import (
	"os"
	"fmt"
	"path"
)

// Run the ComplianceKit Command structure
func Run() error {
	cmd := ckitRootCommand(os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}

func main() {
//	fmt.Printf("%s", path.Base(os.Args[0]))
	if err := Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
