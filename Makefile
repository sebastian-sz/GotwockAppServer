test:
	go test ./...

format_code:
	go fmt ./...

build:
	go build -o bin/GotwockAppServer

clean:
	rm -r bin/*

lint:
	@unformatted_files=$(shell gofmt -l .);\
	if [ "$$unformatted_files" = "" ]; then\
		echo "Content properly formatted";\
	else\
		echo "The following files would be reformatted: ";\
		echo $$unformatted_files;\
		echo "";\
		exit 1;\
	fi

sync_dependencies:
	go mod tidy

e2e_tests:
	python3 end2end_tests/run_e2e_tests.py
