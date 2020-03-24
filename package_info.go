package vendorverify

import (
	"errors"
	"strings"
)

// Package store the package information
type Package struct {
	Name    string
	Version string
	Tag     string
}

// ParsePackage get Package object from string
func ParsePackage(p string) (*Package, error) {
	var tag string
	packageInfoArray := strings.SplitN(p, " ", 3)
	if len(packageInfoArray) < 2 {
		return nil, errors.New("format invalid")
	} else if len(packageInfoArray) == 3 {
		tag = packageInfoArray[2]
	}

	return &Package{
		Name:    packageInfoArray[0],
		Version: packageInfoArray[1],
		Tag:     tag,
	}, nil
}
