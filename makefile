
GO ?= go

tidy:
	$(GO) mod tidy

download: tidy
	$(GO) mod download

run: download
	$(GO) run ./cmd/app/main.go


clear:
	RUN=clear $(GO) run ./cmd/app/main.go

init: clear
	RUN=init $(GO) run ./cmd/app/main.go
