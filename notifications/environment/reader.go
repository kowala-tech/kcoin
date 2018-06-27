package environment

//go:generate moq -out reader_mock.go . Reader

// Reader describes an environment reader
type Reader interface {
	Read(string) string
}
