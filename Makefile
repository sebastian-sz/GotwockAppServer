test:
	go test ./...

format_code:
	go fmt ./...

build:
	go build -o bin/GotwockAppServer

clean:
	rm -r bin/*