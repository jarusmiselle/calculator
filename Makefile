.PHONY: build clean

build: clean
	go build -o ./build/calc ./cmd

clean:
	rm -rf ./build