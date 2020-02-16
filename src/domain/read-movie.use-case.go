package domain

type readMovieUseCase struct {
	reader   DataReader
	notifier DataNotifier
}

// NewReadMovieUseCase is a factory
func NewReadMovieUseCase(reader DataReader, notifier DataNotifier) *readMovieUseCase {
	return &readMovieUseCase{
		reader:   reader,
		notifier: notifier,
	}
}

func (usecase *readMovieUseCase) ReadMovieTitle() {
	titles := usecase.reader.Read()
	for _, title := range titles {
		usecase.notifier.Notify(title)
	}
}
