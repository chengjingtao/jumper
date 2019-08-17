COMMIT=$(shell git rev-parse --short HEAD)
BUILDDATE=$(shell date +%Y%m%d%H%M)


build: build-darwin build-linux build-win

run:
	go run ./cmd/jump.go


build-darwin:
	go build -o ./bin/darwin/jump \
		-ldflags '-w -s -X main.version=${COMMIT} -X main.buildDate=${BUILDDATE}' \
		./cmd/

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/linux-amd64/jump \
		-ldflags '-w -s -X main.version=${COMMIT} -X main.buildDate=${BUILDDATE}' \
		./cmd/

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/windows-amd64/jump.exe \
		-ldflags '-w -s -X main.version=${COMMIT} -X main.buildDate=${BUILDDATE}' \
		./cmd/