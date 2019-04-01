setup:
	go get -u github.com/Masterminds/glide
	go install github.com/Masterminds/glide
	go get -u -f gitlab.com/beacon-software/go-embed
	go install gitlab.com/beacon-software/go-embed
	go get -u -f golang.org/x/tools/cmd/goimports
	go install golang.org/x/tools/cmd/goimports

update: setup
	glide cc
	glide update --strip-vendor

get: setup
	glide install --strip-vendor

test: get
	go test -cover -p 1 ./...

tools: setup
	go generate ./codegen
	go install ./codegen

gen: setup tools _gen fmt

_gen:
	go generate ./...

fmt: setup
	goimports -local gitlab.com/beacon-software/ -w .

example: gen
	go run example/main.gen.go
