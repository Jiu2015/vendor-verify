package cmd

import (
	"errors"
	"os"
	"path/filepath"

	vendorverify "vendor-verify"

	"github.com/spf13/cobra"
)

// the work Path command, it used to get the work Path for this tools
type workPathCommand struct {
	cmd *cobra.Command

	Parameters struct {
		// Path current work path
		Path string

		// ModFilePath go.mod path
		ModFilePath string

		// VendorModFilePath modules.txt path which will be in the root path of vendor
		VendorModFilePath string
	}
}

func (w *workPathCommand) Command() *cobra.Command {
	if w.cmd != nil {
		return w.cmd
	}

	w.cmd = &cobra.Command{
		Use:   "verify",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return w.Execute(args)
		},
	}

	w.cmd.Flags().StringVarP(&w.Parameters.Path,
		"path",
		"p",
		"",
		"the source code root Path")

	return w.cmd
}

func (w *workPathCommand) Execute(args []string) error {
	err := w.ValidateParameters()
	if err != nil {
		return err
	}
	return vendorverify.StartVerify(w.Parameters.ModFilePath, w.Parameters.VendorModFilePath)
}

func (w *workPathCommand) ValidateParameters() error {
	pwd, err := os.Getwd()
	if err != nil {
		return errors.New("cannot get current Path")
	}

	// if not provide the Path, it will use current Path as source code Path
	if len(w.Parameters.Path) <= 0 {
		w.Parameters.Path = pwd
	}

	// Check if or not go.mod file exist
	modFilePath := filepath.Join(w.Parameters.Path, "go.mod")
	_, err = os.Stat(modFilePath)
	if err != nil {
		return errors.New("cannot found go mod files")
	}
	// Set mod file path to parameter struct
	w.Parameters.ModFilePath = modFilePath

	// Check if or not vendor folder exist
	vendorPath := filepath.Join(w.Parameters.Path, "vendor")
	vendorFolder, err := os.Stat(vendorPath)
	if err != nil {
		return errors.New("it looks like you have not use vendor")
	}

	// Check the vendor is or not a folder
	if !vendorFolder.IsDir() {
		return errors.New("it looks like you have not use vendor")
	}

	// Set vendor modules path to parameter struct
	vendorModFilePath := filepath.Join(vendorPath, "modules.txt")
	w.Parameters.VendorModFilePath = vendorModFilePath

	return nil
}
