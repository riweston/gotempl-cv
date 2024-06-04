package types

import (
	"io"

	"gopkg.in/yaml.v3"
)

// SkillData is a struct used to unmarshal the skills.yml file
type SkillData struct {
	// Skills is a list of skills that the user has
	Skills []Skill `yaml:"skills"`
	// SkillsLearning is a list of skills that the user is learning
	SkillsLearning []SkillLearning `yaml:"skillsLearning"`
}

// Skill is a struct that represents a skill that the user has
type Skill struct {
	DisplayName string `yaml:"displayName"`
	// IconifyID uses https://icon-sets.iconify.design/?query=google+cloud to get the icon
	// For example, the icon for Google Cloud is "logos:google-cloud"
	IconifyID string `yaml:"iconifyID"`
	// Level is the level of the skill, this can be a number from 1 to 5
	Level int `yaml:"level"`
}

// ParseData parses the data from the reader into the SkillData and implements the DataParser interface
func (s *SkillData) ParseData(reader io.Reader) error {
	return yaml.NewDecoder(reader).Decode(s)
}

// SkillLearning is a struct that represents a skill that the user is learning
type SkillLearning struct {
	DisplayName string `yaml:"displayName"`
	// IconifyID uses https://icon-sets.iconify.design/?query=google+cloud to get the icon
	// For example, the icon for Google Cloud is "logos:google-cloud"
	IconifyID string `yaml:"iconifyID"`
	Colour    string `yaml:"colour"`
}

// ParseData parses the data from the reader into the SkillLearning and implements the DataParser interface
func (s *SkillLearning) ParseData(reader io.Reader) error {
	return yaml.NewDecoder(reader).Decode(s)
}
