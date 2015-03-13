default:
	make format
	make test
	make style

format:
	gofmt -s=true -w=true *.go
	goimports -w=true *.go

style:
	golint *.go

test:
	go test -test.v=true

coverage:
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

bench:
	go test -test.benchmem=true -test.bench . 2> /dev/null
