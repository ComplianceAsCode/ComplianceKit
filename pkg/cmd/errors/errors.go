/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package errors

import (
	"context"
	"fmt"
	"os"
)

// CheckError is for a Single check failure
func CheckError(err error) {
	if err != nil {
		if err != context.Canceled {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("An error occurred: %v\n", err))
		}
		os.Exit(1)
	}
}
