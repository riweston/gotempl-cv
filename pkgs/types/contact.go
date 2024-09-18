package types

import (
	"io"

	"gopkg.in/yaml.v3"
)

type ContactData struct {
	Email    string `yaml:"email"`
	Location string `yaml:"location"`
	Phone    string `yaml:"phone"`
	Website  string `yaml:"website"`
	LinkedIn string `yaml:"linkedin"`
	GitHub   string `yaml:"github"`
}

func (b *ContactData) ParseData(reader io.Reader) error {
	err := yaml.NewDecoder(reader).Decode(b)
	if err != nil {
		return err
	}
	return nil
}
