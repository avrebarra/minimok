# simple makefile @avrebarra
NAME=minimok
VERSION=v1
COVERAGE_MIN=0.0

## coverage: Show coverage report in browser
coverage: test
	go tool cover -html=cp.out

## --------: 

## test: Run test and enforce go coverage
test:
	go test ./... -coverprofile cp.out

## test_coverage: Enforce test coverage percentage
test_coverage: test
	$(eval COVERAGE_CURRENT = $(shell go tool cover -func=cp.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' ))
	$(eval COVERAGE_PASSED = $(shell echo "$(COVERAGE_CURRENT) >= $(COVERAGE_MIN)" | bc -l ))

	@if [ $(COVERAGE_PASSED) == 0 ] ; then \
		echo "coverage below threshold"; \
		exit 2; \
    fi

## benchmark: Run benchmark test
benchmark:
	go test -bench=.

## watch: development with air
watch:
	air -c .air/development.air.toml

## build: Build binary applications
build:
	@go generate ./...
	@echo building binary to ./dist/${NAME}
	@go build -o ./dist/${NAME} .

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run with parameter options: "
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
