.PHONY: vendor
build: clean vendor
	go build -mod=vendor -o dist/dsnote main.go

clean:
	rm -rf dist

run: vendor
	go run -mod=vendor main.go

test:
	go test --coverprofile=coverage.out ./...

cover: test
	go tool cover -func=coverage.out

vendor:
	go get
	go mod vendor
