export GO111MODULE=on
GOIMPORTS=goimports
BINARY_NAME=places
DOCKER_TAG=`date -u +%Y%m%d`

HASRICHGO := $(shell which richgo)
ifdef HASRICHGO
    GOTEST=richgo test
else
    GOTEST=go test
endif

all: dep build
dep: dep-go
build: build-go
install: install-bin install-assets
test: test-go
test-all: test
fmt: fmt-go
upgrade: upgrade-go
start:
	go run cmd/places/places.go start -d
stop:
	go run cmd/places/places.go stop
terminal:
	docker-compose exec places bash
migrate:
	go run cmd/places/places.go migrate
install-bin:
	scripts/build.sh prod ~/.local/bin/$(BINARY_NAME)
install-assets:
	$(info Installing assets)
	mkdir -p ~/.config/places
	mkdir -p ~/.cache/places
	mkdir -p ~/.local/share/places
	cp -r assets/static ~/.local/share/places
	rsync -a -v --ignore-existing assets/config/*.yml ~/.config/places
	find ~/.local/share/places -name '.*' -type f -delete
dep-go:
	go build -v ./...
build-go:
	rm -f $(BINARY_NAME)
	scripts/build.sh debug $(BINARY_NAME)
build-static:
	rm -f $(BINARY_NAME)
	scripts/build.sh static $(BINARY_NAME)
test-go:
	$(info Running all Go unit tests...)
	$(GOTEST) -count=1 -tags=slow -timeout 20m ./internal/...
test-verbose:
	$(info Running all Go unit tests in verbose mode...)
	$(GOTEST) -tags=slow -timeout 20m -v ./internal/...
test-short:
	$(info Running short Go unit tests in verbose mode...)
	$(GOTEST) -short -timeout 5m -v ./internal/...
test-race:
	$(info Running all Go unit tests with race detection in verbose mode...)
	$(GOTEST) -tags=slow -race -timeout 60m -v ./internal/...
test-codecov:
	$(info Running all Go unit tests with code coverage report for codecov...)
	go test -count=1 -tags=slow -timeout 30m -coverprofile=coverage.txt -covermode=atomic -v ./internal/...
	scripts/codecov.sh
test-coverage:
	$(info Running all Go unit tests with code coverage report...)
	go test -count=1 -tags=slow -timeout 30m -coverprofile=coverage.txt -covermode=atomic -v ./internal/...
	go tool cover -html=coverage.txt -o coverage.html
clean:
	rm -f $(BINARY_NAME)
	rm -f *.log
docker-development:
	scripts/docker-build.sh places-dev $(DOCKER_TAG)
	scripts/docker-push.sh places-dev $(DOCKER_TAG)
docker-places:
	scripts/docker-build.sh places $(DOCKER_TAG)
	scripts/docker-push.sh places $(DOCKER_TAG)
fmt-go:
	goimports -w internal cmd
	go fmt ./internal/... ./cmd/...
tidy:
	go mod tidy
upgrade-go:
	go mod tidy
	go get -u
