NAME = lablog

all:
	make format
	make test
	make build

generate:
	cd src/web/; go-bindata -pkg="web" html/

format:
	find . -name "*.go" -not -path './vendor/*' -type f -exec goimports -w=true {} \;

test:
	rm -rf vendor/
	go test -v ./...
	git checkout -- vendor/

build:
	go build -ldflags "-X github.com/AlexanderThaller/lablog/src/commands.buildTime=`date +%s` -X github.com/AlexanderThaller/lablog/src/commands.buildVersion=`git describe --always`" -o "$(NAME)"

install:
	cp "$(NAME)" /usr/local/bin

uninstall:
	rm "/usr/local/bin/$(NAME)"

clean:
	rm "$(NAME)"
	rm *.pprof
	rm *.pdf

callgraph:
	go tool pprof --pdf "$(NAME)" cpu.pprof > callgraph.pdf

memograph:
	go tool pprof --pdf "$(NAME)" mem.pprof > memograph.pdf

bench:
	rm -rf vendor/
	mkdir -p benchmarks/`git describe --always`/
	go test -test.benchmem=true -test.bench . 2> /dev/null | tee benchmarks/`git describe --always`/`date +%s`
	git checkout -- vendor/

coverage:
	rm -rf vendor/
	rm -f coverage.out
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o=/tmp/coverage.html
	git checkout -- vendor/

lint:
	rm -rf vendor/
	gometalinter ./...; git checkout -- vendor/
