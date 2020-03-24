package vendorverify

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// StartVerify start verify
func StartVerify(modFilePath, modulesFilePath string) error {
	// load go.mod file
	modFile, err := LoadModFile(modFilePath)
	if err != nil {
		return err
	}

	// load modules.txt file
	modulesFile, err := LoadModulesFile(modulesFilePath)
	if err != nil {
		return err
	}

	err = verify(modFile, modulesFile)
	if err != nil {
		return err
	}

	return nil
}

// it just verify package name and version, the tag will be ignore
func verify(modFile *ModFile, modulesFile map[string]*Package) error {
	var (
		errBuf bytes.Buffer
	)
	for _, pkg := range modFile.Packages {
		if v, ok := modulesFile[pkg.Name]; ok {
			if pkg.Name != v.Name || pkg.Version != v.Version {
				errMsg := fmt.Sprintf("%s %s is not sync to vendor\r\n", pkg.Name, pkg.Version)
				errBuf.WriteString(errMsg)
			}
			continue
		}

		errMsg := fmt.Sprintf("%s %s is not exist in vendor\r\n", pkg.Name, pkg.Version)
		errBuf.WriteString(errMsg)
	}

	errStr := strings.TrimSpace(errBuf.String())
	if len(errStr) != 0 {
		return errors.New(errStr)
	}

	return nil
}
