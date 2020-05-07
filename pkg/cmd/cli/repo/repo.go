/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package repo

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/Masterminds/vcs"
	"github.com/complianceascode/compliancekit/pkg/cmd/errors"
	"github.com/spf13/cobra"
)

var (
	repository string
	directory  string
)

// NewCmdDocs creates the compliance documentation.
func NewCmdGet(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Download compliance repository from GitHub",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunGet(out, cmd, args)
			errors.CheckError(err)
		},
	}

	cmd.Flags().StringVarP(&repository, "repository", "r", "https://github.com/complianceascode/content", "Version control repository")
	cmd.Flags().StringVarP(&directory, "directory", "d", "", "Local directory to clone repostory")

	return cmd
}

func RunGet(out io.Writer, cmd *cobra.Command, args []string) error {
	var local string
	u, err := url.Parse(repository)

	if len(directory) > 0 {
		local = directory
	} else {
		local = filepath.Base(u.Path)
	}

	if _, err := os.Stat(local); !os.IsNotExist(err) {
		err := fmt.Errorf("Local clone '%s' already exists! Skipping...", local)
		return err
	}

	fmt.Printf("Cloning '%s' into '%s'...\n", u, local)
	repo, err := vcs.NewRepo(repository, local)
	if err != nil {
		fmt.Printf("Error, repo")

	}

	err = repo.Get()
	if err != nil {
		fmt.Printf("Error: get")
	}

	return nil
}
