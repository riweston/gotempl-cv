package types

import (
	"io"

	"gopkg.in/yaml.v3"
)

type CertificateData struct {
	Certificates []Certificate `yaml:"certificates"`
}

type Certificate struct {
	DisplayName string `yaml:"displayName"`
	Issuer      string `yaml:"issuer"`
	BadgePath   string `yaml:"badgePath"`
	IssueDate   string `yaml:"issueDate"`
	ExpireDate  string `yaml:"expireDate"`
}

func (c *CertificateData) ParseData(reader io.Reader) error {
	err := yaml.NewDecoder(reader).Decode(c)
	if err != nil {
		return err
	}
	return nil
}
