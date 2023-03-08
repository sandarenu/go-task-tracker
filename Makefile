LINUX := linux
AMD64 := amd64

FILES := domain.go db_access.go main.go

run:
	go run $(FILES)

build:
	GOARCH=$(AMD64) GOOS=$(LINUX) go build -o bin/tt-$(LINUX) $(FILES)

clean:
	go clean
	rm -rf bin/*