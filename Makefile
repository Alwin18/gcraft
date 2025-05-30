build:
	go build -o gcraft

generate:
	rm -rf my-app
	./gcraft create my-app

all: build generate