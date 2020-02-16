# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool

# file names
BINARY_NAME=movie-reader
COVER_OUT_NAME=cover.out
COVER_HTML_NAME=cover.html
FILE_IMPORT=movies.csv

build: 
	$(GOBUILD) -v ./...

start:
	./$(BINARY_NAME)

run:
	$(GORUN) main.go $(FILE_IMPORT)

test:
	$(GOTEST) ./... -coverprofile $(COVER_OUT_NAME)
	$(GOTOOL) cover -html=$(COVER_OUT_NAME) -o $(COVER_HTML_NAME)

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(COVER_OUT_NAME)
	rm -f $(COVER_HTML_NAME)
