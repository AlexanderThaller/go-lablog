NAME = lablog

all:
	make format
	make test
	make build

format:
	find . -name "*.go" -type f -exec gofmt -s=true -w=true {} \;
	find . -name "*.go" -type f -exec goimports -w=true {} \;

test:
	go test

build:
	go build -ldflags "-X main.buildTime `date +%s` -X main.buildVersion `git describe --always`" -o "$(NAME)"

clean:
	rm "$(NAME)"
	rm *.pprof
	rm *.pdf

install:
	cp "$(NAME)" /usr/local/bin

uninstall:
	rm "/usr/local/bin/$(NAME)"

callgraph:
	go tool pprof --pdf "$(NAME)" cpu.pprof > callgraph.pdf

memograph:
	go tool pprof --pdf "$(NAME)" mem.pprof > memograph.pdf

dependencies_save:
	godep save ./...

dependencies_restore:
	godep restore ./...

bench:
	go test -test.benchmem=true -test.bench . 2> /dev/null

coverage:
	rm -f coverage.out
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
