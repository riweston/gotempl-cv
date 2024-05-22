package types

import (
	"gopkg.in/yaml.v3"
	"io"
)

type BioData struct {
	Name               string      `yaml:"name"`
	ProfilePicturePath string      `yaml:"profilePicturePath"`
	Headline           string      `yaml:"headline"`
	Location           string      `yaml:"location"`
	Email              string      `yaml:"email"`
	About              string      `yaml:"about"`
	Tags               BioDataTags `yaml:"tags"`
}

type BioDataTags struct {
	OpenToOpportunities bool `yaml:"openToOpportunities"`
	Remote              bool `yaml:"remote"`
	FullTime            bool `yaml:"fullTime"`
}

func (b *BioData) ParseData(reader io.Reader) error {
	err := yaml.NewDecoder(reader).Decode(b)
	if err != nil {
		return err
	}
	return nil
}
