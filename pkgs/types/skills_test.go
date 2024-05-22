package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
		want    SkillDataList
	}{
		{
			name: "Test SkillData_ParseData",
			args: args{
				filePath: "./examples/skills.yml",
				data:     &SkillDataList{},
			},
			wantErr: assert.NoError,
			want: []SkillData{
				{
					Name:        "foo",
					DisplayName: "bar",
					Level:       1,
				},
				{
					Name:        "baz",
					DisplayName: "qux",
					Level:       2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, ReadYamlFile(tt.args.filePath, tt.args.data), fmt.Sprintf("ReadYamlFile(%v, %v)", tt.args.filePath, tt.args.data))
			assert.Equal(t, tt.want, *tt.args.data.(*SkillDataList), "ReadYamlFile(%v, %v)", tt.args.filePath, tt.args.data)
		})
	}
}
