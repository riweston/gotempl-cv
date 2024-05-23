package types

import (
	"io"

	"gopkg.in/yaml.v3"
)

type BioData struct {
	Name               string   `yaml:"name"`
	ProfilePicturePath string   `yaml:"profilePicturePath"`
	Headline           string   `yaml:"headline"`
	Location           string   `yaml:"location"`
	Email              string   `yaml:"email"`
	About              string   `yaml:"about"`
	Tags               []string `yaml:"tags"`
}

func (b *BioData) ParseData(reader io.Reader) error {
	err := yaml.NewDecoder(reader).Decode(b)
	if err != nil {
		return err
	}
	return nil
}
