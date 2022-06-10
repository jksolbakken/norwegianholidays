holidays:
	go build -o bin/holidays cmd/holidays/*.go

test: fmt vet
	go test ./... -coverprofile cover.out -short
fmt:
	go fmt ./...
vet:
	go vet ./...

