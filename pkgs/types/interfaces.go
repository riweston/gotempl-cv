package types

import "io"

type DataParser interface {
	ParseData(reader io.Reader) error
}
