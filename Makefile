include .env
export

docker-build:
	docker build -t capital-gains .

docker-run:
	docker run -it --rm --name capital-gains capital-gains

deps:
	go mod tidy && go get -u ./cmd

build:
	go build -o bin/capital-gains ./cmd

run:
	go run -race ./cmd/main.go

tests:
	go clean -testcache && go test ./... -cover -v

coverage:
	go test ./... -coverprofile coverage.out && go tool cover -func coverage.out && go tool cover -html=coverage.out -o coverage.html && open coverage.html

mock-generate:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@go generate ./...
