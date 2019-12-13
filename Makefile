test: *.go
	go test .

run: *.go
	go run .

run_http_408:
	go run . http 408
