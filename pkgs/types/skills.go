package types

import (
	"gopkg.in/yaml.v3"
	"io"
)

type SkillData struct {
	Name        string `yaml:"name"`
	IconifyID   string `yaml:"iconifyID"`
	DisplayName string `yaml:"displayName"`
	Level       int    `yaml:"level"`
}

type SkillDataList []SkillData

func (s *SkillDataList) ParseData(reader io.Reader) error {
	return yaml.NewDecoder(reader).Decode(s)
}

type SkillLearningData struct {
	Name        string `yaml:"name"`
	DisplayName string `yaml:"displayName"`
}

type SkillLearningDataList []SkillLearningData

func (s *SkillLearningDataList) ParseData(reader io.Reader) error {
	return yaml.NewDecoder(reader).Decode(s)
}
