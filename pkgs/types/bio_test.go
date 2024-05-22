package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadYamlFile(t *testing.T) {
	type args struct {
		filePath string
		data     DataParser
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		want    BioData
	}{
		{
			name: "Test ReadYamlFile",
			args: args{
				filePath: "./examples/bio.yml",
				data:     &BioData{},
			},
			wantErr: assert.NoError,
			want: BioData{
				Name:     "foo",
				Headline: "bar",
				Location: "baz",
				Email:    "qux",
				About:    "quux",
				Tags: BioDataTags{
					OpenToOpportunities: true,
					Remote:              true,
					FullTime:            true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, ReadYamlFile(tt.args.filePath, tt.args.data), fmt.Sprintf("ReadYamlFile(%v, %v)", tt.args.filePath, tt.args.data))
			assert.Equal(t, tt.want, *tt.args.data.(*BioData), "ReadYamlFile(%v, %v)", tt.args.filePath, tt.args.data)
		})
	}
}
