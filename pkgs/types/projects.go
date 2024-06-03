package types

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Project struct {
	Name         string       `yaml:"name"`
	Industry     string       `yaml:"industry"`
	StartDate    string       `yaml:"startDate"`
	EndDate      string       `yaml:"endDate"`
	Role         string       `yaml:"role"`
	TeamSize     int          `yaml:"teamSize"`
	Description  string       `yaml:"description"`
	Technologies []Technology `yaml:"technologies"`
}

type Technology struct {
	DisplayName string `yaml:"displayName"`
	URL         string `yaml:"url"`
	IconifyID   string `yaml:"iconifyID"`
	Colour      string `yaml:"colour"`
}

type ProjectData struct {
	Projects []Project `yaml:"projects"`
}

func (p *ProjectData) ParseData(reader io.Reader) error {
	return yaml.NewDecoder(reader).Decode(p)
}
