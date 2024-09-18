package types

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Job struct {
	Company       string   `yaml:"company"`
	LogoPath      string   `yaml:"logoPath"`
	LinkedIn      string   `yaml:"linkedIn"`
	Website       string   `yaml:"website"`
	CurrentRole   string   `yaml:"currentRole"`
	StartDate     string   `yaml:"startDate"`
	EndDate       string   `yaml:"endDate"`
	Description   string   `yaml:"description"`
	ShortDesc     string   `yaml:"shortDesc"`
	PreviousRoles []string `yaml:"previousRoles"`
}

type WorkExperienceData struct {
	Jobs []Job `yaml:"jobs"`
}

func (w *WorkExperienceData) ParseData(reader io.Reader) error {
	return yaml.NewDecoder(reader).Decode(w)
}
