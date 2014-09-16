NAME = guberno

all:
	make format
	make build

format:
	find . -name "*.go" -type f -exec gofmt -s=true -w=true {} \;
	find . -name "*.go" -type f -exec goimports -w=true {} \;

test:
	go test ./...

build:
	go build -ldflags "-X main.buildtime `date +%s` -X main.version `git describe --always`"

clean:
	rm "$(NAME)"
	rm *.pprof
	rm *.pdf

install:
	cp "$(NAME)" /usr/local/bin

callgraph:
	go tool pprof --pdf "$(NAME)" cpu.pprof > callgraph.pdf

memograph:
	go tool pprof --pdf "$(NAME)" mem.pprof > memograph.pdf

dependencies_save:
	godep save ./...

dependencies_restore:
	godep restore ./...
