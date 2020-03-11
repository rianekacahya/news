-include config
REST_MAIN := "$(CURDIR)/cmd/rest"
BIN_REST := "$(CURDIR)/bin/rest"
EVENT_MAIN := "$(CURDIR)/cmd/event"
BIN_EVENT := "$(CURDIR)/bin/event"

.PHONY: prepare

prepare: clean init fetch

init:
	@go mod init github.com/rianekacahya/news

fetch:
	@go mod tidy

build-rest:
	@go build -i -v -o $(BIN_REST) $(REST_MAIN)

build-event:
	@go build -i -v -o $(BIN_EVENT) $(EVENT_MAIN)

build: build-rest build-event

run-rest:
	@go run $(CURDIR)/cmd/rest/main.go

run-event:
	@go run $(CURDIR)/cmd/event/main.go

deploy: prepare build

run:
	@docker-compose up -d

stop:
	@docker-compose stop

clean:
	@rm -f $(CURDIR)/go.mod $(CURDIR)/go.sum \
	@rm -rf $(CURDIR)/bin