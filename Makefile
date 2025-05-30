VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")

build:
	go build -ldflags="-X 'github.com/Alwin18/gcraft/cmd.Version=$(VERSION)'" -o gcraft

generate:
	rm -rf my-app
	./gcraft create my-app

all: build generate