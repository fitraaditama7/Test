BINARY=app

test:
	go test ./...

cover:
	go test -cover ./...  -coverprofile=coverage.out && go tool cover -html=coverage.out

run:
	@go run main.go

download:
	@go mod download
	@go mod tidy

build:
	@go build -o ${BINARY} main.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean download cover test run build