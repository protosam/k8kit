GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
build-all:
	GOOS=windows GOARCH=amd64 go build -o bin/k8kit-amd64-windows.exe ./main.go
	GOOS=linux GOARCH=amd64 go build -o bin/k8kit-amd64-linux ./main.go
	GOOS=linux GOARCH=arm64 go build -o bin/k8kit-arm64-linux ./main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/k8kit-amd64-darwin ./main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/k8kit-arm64-darwin ./main.go

build:
ifeq ($(GOOS), windows)
	go build -o bin/k8kit-$(GOARCH)-$(GOOS).exe ./main.go
else
	go build -o bin/k8kit-$(GOARCH)-$(GOOS) ./main.go
endif

clean:
	rm -rf ./bin
