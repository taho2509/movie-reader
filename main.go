package main

import (
	"fmt"
	"os"
	"taho2509/movie-reader/src/domain"
	"taho2509/movie-reader/src/interfaces"
)

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		panic("Two arguments required")
	}

	reader := interfaces.NewFileReader(argsWithProg[1])
	notifier := interfaces.NewNatsNotifier()
	readMovieUseCase := domain.NewReadMovieUseCase(reader, notifier)
	readMovieUseCase.ReadMovieTitle()
	fmt.Println("Finish")
}
