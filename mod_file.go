package vendorverify

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// ModFile go mod/modules package infos
type ModFile struct {
	Packages []*Package
}

// AppendPackage append package to ModFile
func (m *ModFile) AppendPackage(pk *Package) {
	m.Packages = append(m.Packages, pk)
}

// NewModFile create new ModFile object
func NewModFile() *ModFile {
	return &ModFile{
		Packages: []*Package{},
	}
}

// LoadModFile it will load go.mod file
func LoadModFile(path string) (*ModFile, error) {
	var requireBody bool
	modFile := NewModFile()

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("the go mod not exist")
	}
	defer f.Close()

	r := bufio.NewScanner(f)

	for r.Scan() {
		tmp := r.Text()
		if strings.Contains(tmp, "require") {
			// The require has multi-line
			if strings.Contains(tmp, "(") {
				requireBody = true
				continue
			}

			// Single line
			content := strings.TrimSpace(tmp[7:])
			tmpPackage, err := ParsePackage(content)
			if err != nil {
				return nil, err
			}
			modFile.AppendPackage(tmpPackage)
		}

		if requireBody {
			// requires stop
			if strings.Contains(tmp, ")") {
				requireBody = false
				break
			}

			tmpPackage, err := ParsePackage(strings.TrimSpace(tmp))
			if err != nil {
				return nil, err
			}
			modFile.AppendPackage(tmpPackage)
		}
	}

	return modFile, nil
}

// LoadModulesFile load modules.txt file
// it will return map[string]*package, the key is package name,
// the value is package pointer
func LoadModulesFile(path string) (map[string]*Package, error) {
	modMap := make(map[string]*Package)
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		tmpContent := scanner.Text()
		if strings.HasPrefix(tmpContent, "#") {
			tmpPackage, err := ParsePackage(tmpContent[2:])
			if err != nil {
				return nil, err
			}
			modMap[tmpPackage.Name] = tmpPackage
		}
	}

	return modMap, nil
}
