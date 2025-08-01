APP_NAME=caas-eks-api
GO_FILES=$(shell find . -type f -name '*.go')
GO_FMT=$(shell gofmt -l $(GO_FILES))

.PHONY: all build run fmt vet clean swagger

all: build

build:
	go build -o $(APP_NAME) ./cmd

run:
	go run ./cmd/main.go

swagger:
	swag init --parseDependency --parseInternal

clean:
	rm -f $(APP_NAME)
	go run main.go

fmt:
	@echo "Checking formatting..."
	@if [ "$(GO_FMT)" ]; then \
		echo "Files not formatted:"; \
		echo "$(GO_FMT)"; \
		exit 1; \
	else \
		echo "All files formatted"; \
	fi

vet:
	go vet ./...
