.PHONY: gobuild godeps gofmt golint goinstall uibuild uideps uilint build \
	deps fmt lint test all

all: build
build: gobuild uibuild
deps: godeps uideps
fmt: gofmt
lint: golint uilint
test: gotest

gobuild:
	go install -a ./...
godeps:
	go get -v -t github.com/amiforus/tchaik/cmd/...
gofmt:
	go fmt ./...
golint:
	./scripts/verify-gofmt.sh ./**/*.go
	go vet ./...
gotest:
	go test ./...

uibuild:
	cd cmd/tchaik/ui; gulp
uideps:
	cd cmd/tchaik/ui; npm install
uilint:
	cd cmd/tchaik/ui; gulp lint
