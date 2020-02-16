package domain

// DataNotifier is used to read an input streams
type DataNotifier interface {
	Notify(string)
}
