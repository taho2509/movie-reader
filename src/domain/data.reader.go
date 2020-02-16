package domain

// DataReader is used to read an input streams
type DataReader interface {
	Read() []string
}
