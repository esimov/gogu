build:
	go build -v ./...

test:
	go test -v ./... -race -cover -run=. -count=1

coverage:
	go test -v ./... -run=. -cover -covermode=count -coverprofile=coverage.out    
	go tool cover -func=coverage.out -o coverage.out

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...

lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...