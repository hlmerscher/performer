default: test

test:
	go test ./...

test-verbose:
	go test -v ./...
