build:
	go build -o gcraft

generate:
	./gcraft create my-app

all: build generate