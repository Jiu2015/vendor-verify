package vendorverify

import (
	"reflect"
	"testing"
)

func TestParsePackage(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		want    *Package
		wantErr bool
	}{
		{
			name: "one content",
			args: args{
				p: "github.com/gin-gonic/gin",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "two contents",
			args: args{
				p: "github.com/gin-gonic/gin v1.0.0",
			},
			want: &Package{
				Name:    "github.com/gin-gonic/gin",
				Version: "v1.0.0",
				Tag:     "",
			},
			wantErr: false,
		},
		{
			name: "three contents",
			args: args{
				p: "github.com/gin-gonic/gin v1.0.0 //",
			},
			want: &Package{
				Name:    "github.com/gin-gonic/gin",
				Version: "v1.0.0",
				Tag:     "//",
			},
			wantErr: false,
		},
		{
			name: "four contents",
			args: args{
				p: "github.com/gin-gonic/gin v1.0.0 // indirect",
			},
			want: &Package{
				Name:    "github.com/gin-gonic/gin",
				Version: "v1.0.0",
				Tag:     "// indirect",
			},
			wantErr: false,
		},
		{
			name: "five contents",
			args: args{
				p: "github.com/gin-gonic/gin v1.0.0 // indirect aaaa",
			},
			want: &Package{
				Name:    "github.com/gin-gonic/gin",
				Version: "v1.0.0",
				Tag:     "// indirect aaaa",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePackage(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePackage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
