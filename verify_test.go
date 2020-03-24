package vendorverify

import "testing"

func Test_verify(t *testing.T) {
	type args struct {
		modFile     *ModFile
		modulesFile map[string]*Package
	}

	// 1. not exist
	// 3. version is different
	// 4. verify successfully
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "modules_not_exist",
			args: args{
				modFile: &ModFile{
					Packages: []*Package{
						&Package{
							Name:    "golang.org/x/crypto",
							Version: "v0.0.0-20190820162420-60c769a6c586",
							Tag:     "// indirect",
						},
						&Package{
							Name:    "golang.org/x/text",
							Version: "v0.3.2",
							Tag:     "// indirect",
						},
					},
				},
				modulesFile: map[string]*Package{
					"github.com/hashicorp/hcl": &Package{
						Name:    "github.com/hashicorp/hcl",
						Version: "v1.0.0",
						Tag:     "",
					},
					"github.com/davecgh/go-spew": &Package{
						Name:    "github.com/davecgh/go-spew",
						Version: "v1.1.1",
						Tag:     "",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "version_is_different",
			args: args{
				modFile: &ModFile{
					Packages: []*Package{
						&Package{
							Name:    "golang.org/x/crypto",
							Version: "v0.0.0-20190820162420-60c769a6c586",
							Tag:     "// indirect",
						},
						&Package{
							Name:    "golang.org/x/text",
							Version: "v0.3.2",
							Tag:     "// indirect",
						},
					},
				},
				modulesFile: map[string]*Package{
					"golang.org/x/crypto": &Package{
						Name:    "golang.org/x/crypto",
						Version: "v1.0.0",
						Tag:     "",
					},
					"golang.org/x/text": &Package{
						Name:    "golang.org/x/text",
						Version: "v0.3.2",
						Tag:     "",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "verify_successfully",
			args: args{
				modFile: &ModFile{
					Packages: []*Package{
						&Package{
							Name:    "golang.org/x/crypto",
							Version: "v0.0.0-20190820162420-60c769a6c586",
							Tag:     "// indirect",
						},
						&Package{
							Name:    "golang.org/x/text",
							Version: "v0.3.2",
							Tag:     "// indirect",
						},
					},
				},
				modulesFile: map[string]*Package{
					"golang.org/x/crypto": &Package{
						Name:    "golang.org/x/crypto",
						Version: "v0.0.0-20190820162420-60c769a6c586",
						Tag:     "// indirect",
					},
					"golang.org/x/text": &Package{
						Name:    "golang.org/x/text",
						Version: "v0.3.2",
						Tag:     "// indirect",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "modules_are_more_than_mod",
			args: args{
				modFile: &ModFile{
					Packages: []*Package{
						&Package{
							Name:    "golang.org/x/crypto",
							Version: "v0.0.0-20190820162420-60c769a6c586",
							Tag:     "// indirect",
						},
						&Package{
							Name:    "golang.org/x/text",
							Version: "v0.3.2",
							Tag:     "// indirect",
						},
					},
				},
				modulesFile: map[string]*Package{
					"golang.org/x/crypto": &Package{
						Name:    "golang.org/x/crypto",
						Version: "v0.0.0-20190820162420-60c769a6c586",
						Tag:     "// indirect",
					},
					"golang.org/x/text": &Package{
						Name:    "golang.org/x/text",
						Version: "v0.3.2",
						Tag:     "// indirect",
					},
					"github.com/davecgh/go-spew": &Package{
						Name:    "github.com/davecgh/go-spew",
						Version: "v1.1.1",
						Tag:     "",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verify(tt.args.modFile, tt.args.modulesFile); (err != nil) != tt.wantErr {
				t.Errorf("verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
