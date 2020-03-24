package vendorverify

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadModFile(t *testing.T) {
	type args struct {
		path string
	}

	currentPath, err := os.Getwd()
	if err != nil {
		t.Errorf("cannot get current path: %v", err)
	}

	tests := []struct {
		name    string
		args    args
		want    *ModFile
		wantErr bool
	}{
		{
			name: "load_single_require_mod",
			args: args{
				path: filepath.Join(currentPath, "/testdata/single_require.mod"),
			},
			want: &ModFile{
				Packages: []*Package{
					{
						Name:    "github.com/gin-gonic/gin",
						Version: "v1.5.0",
						Tag:     "",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "load_single_with_tag_mod",
			args: args{
				path: filepath.Join(currentPath, "/testdata/single_require_with_tag.mod"),
			},
			want: &ModFile{
				Packages: []*Package{
					{
						Name:    "github.com/gin-gonic/gin",
						Version: "v1.5.0",
						Tag:     "// indirect",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "load_multiple_require_mod",
			args: args{
				path: filepath.Join(currentPath, "/testdata/multiple_require.mod"),
			},
			want: &ModFile{
				Packages: []*Package{
					{
						Name:    "github.com/spf13/cobra",
						Version: "v0.0.5",
						Tag:     "",
					},
					{
						Name:    "github.com/stretchr/testify",
						Version: "v1.2.2",
						Tag:     "",
					},
					{
						Name:    "golang.org/x/crypto",
						Version: "v0.0.0-20190820162420-60c769a6c586",
						Tag:     "// indirect",
					},
					{
						Name:    "golang.org/x/sys",
						Version: "v0.0.0-20190826190057-c7b8b68b1456",
						Tag:     "// indirect",
					},
					{
						Name:    "golang.org/x/text",
						Version: "v0.3.2",
						Tag:     "// indirect",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "file_not_exist",
			args: args{
				path: "abc.mod",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadModFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadModFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadModFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadModuleFile(t *testing.T) {
	type args struct {
		path string
	}

	currentPath, err := os.Getwd()
	if err != nil {
		t.Errorf("cannot get current path: %v", err)
	}

	tests := []struct {
		name    string
		args    args
		want    map[string]*Package
		wantErr bool
	}{
		{
			name: "common_modules_file",
			args: args{
				path: filepath.Join(currentPath, "/testdata/common_modules.txt"),
			},
			want: map[string]*Package{
				"github.com/davecgh/go-spew": &Package{
					Name:    "github.com/davecgh/go-spew",
					Version: "v1.1.1",
					Tag:     "",
				},
				"github.com/fsnotify/fsnotify": &Package{
					Name:    "github.com/fsnotify/fsnotify",
					Version: "v1.4.7",
					Tag:     "",
				},
				"github.com/hashicorp/hcl": &Package{
					Name:    "github.com/hashicorp/hcl",
					Version: "v1.0.0",
					Tag:     "",
				},
				"github.com/inconshreveable/mousetrap": &Package{
					Name:    "github.com/inconshreveable/mousetrap",
					Version: "v1.0.0",
					Tag:     "",
				},
			},
			wantErr: false,
		},
		{
			name: "no_packages_in_modules_file",
			args: args{
				path: filepath.Join(currentPath, "/testdata/no_packages_in_modules.txt"),
			},
			want:    map[string]*Package{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadModulesFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadModulesFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadModulesFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
