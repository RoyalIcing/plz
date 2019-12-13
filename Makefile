test: *.go
	go test .

run: *.go
	go run .

run_help:
	go run . help

run_help_http:
	go run . help http

run_http_no_args:
	go run . http

run_http_408:
	go run . http 408
