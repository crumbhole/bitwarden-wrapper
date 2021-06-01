.PHONY:	all clean code-vet code-fmt test

DEPS := $(shell find . -type f -name "*.go" -printf "%p ")

all: code-vet code-fmt test

clean:
	$(RM) -rf build

test:
	go test ./...

code-vet: $(DEPS)
## Run go vet for this project. More info: https://golang.org/cmd/vet/
	@echo go vet
	go vet $$(go list ./... )

code-fmt: $(DEPS)
## Run go fmt for this project
	@echo go fmt
	go fmt $$(go list ./... )
