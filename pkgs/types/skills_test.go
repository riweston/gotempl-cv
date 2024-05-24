package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkillData_ParseData(t *testing.T) {
	type args struct {
		filePath string
		data     DataParser
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		want    SkillData
	}{
		{
			name: "Test SkillData_ParseData",
			args: args{
				filePath: "./examples/skills.yml",
				data:     &SkillData{},
			},
			wantErr: assert.NoError,
			want: SkillData{
				Skills: []Skill{

					{
						DisplayName: "bar",
						IconifyID:   "logos:baz",
						Level:       1,
					},
					{
						DisplayName: "qux",
						IconifyID:   "logos:quux",
						Level:       2,
					},
				},
				SkillsLearning: []SkillLearning{
					{
						DisplayName: "quux",
						IconifyID:   "devicon:corge",
					},
					{
						DisplayName: "grault",
						IconifyID:   "devicon:garply",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, ReadYamlFile(tt.args.filePath, tt.args.data), fmt.Sprintf("ReadYamlFile(%v, %v)", tt.args.filePath, tt.args.data))
			assert.Equal(t, tt.want, *tt.args.data.(*SkillData), "ReadYamlFile(%v, %v)", tt.args.filePath, tt.args.data)
		})
	}
}
