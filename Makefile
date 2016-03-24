NAME = lablog

all:
	make dependencies
	make generate
	make format
	make test
	make build

dependencies:
	go get -u github.com/jteeuwen/go-bindata/...

generate:
	cd src/web/; go-bindata -pkg="web" templates/

format:
	find . -name "*.go" -not -path './vendor/*' -type f -exec goimports -w=true {} \;

crossbuild:
	make dependencies
	make generate
	make format
	make test

	rm -rf "bin"

	# linux - amd64
	env GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/AlexanderThaller/lablog/cmd.buildTime=`date +%s` -X github.com/AlexanderThaller/lablog/cmd.buildVersion=`git describe --tag --always`" -o "bin/$(NAME)_linux_adm64"
	xz --best --extreme "bin/$(NAME)_linux_adm64"

	# linux - arm
	env GOOS=linux GOARCH=arm go build -ldflags "-X github.com/AlexanderThaller/lablog/cmd.buildTime=`date +%s` -X github.com/AlexanderThaller/lablog/cmd.buildVersion=`git describe --tag --always`" -o "bin/$(NAME)_linux_arm"
	xz --best --extreme "bin/$(NAME)_linux_arm"

	# freebsd - amd64
	env GOOS=freebsd GOARCH=amd64 go build -ldflags "-X github.com/AlexanderThaller/lablog/cmd.buildTime=`date +%s` -X github.com/AlexanderThaller/lablog/cmd.buildVersion=`git describe --tag --always`" -o "bin/$(NAME)_freebsd_amd64"
	xz --best --extreme "bin/$(NAME)_freebsd_amd64"

	# darwin - amd64
	env GOOS=darwin GOARCH=amd64 go build -ldflags "-X github.com/AlexanderThaller/lablog/cmd.buildTime=`date +%s` -X github.com/AlexanderThaller/lablog/cmd.buildVersion=`git describe --tag --always`" -o "bin/$(NAME)_darwin_amd64"
	xz --best --extreme "bin/$(NAME)_darwin_amd64"

test:
	GO15VENDOREXPERIMENT=1 go test `GO15VENDOREXPERIMENT=1 go list ./... | grep -v '/vendor/'`

build:
	go build -ldflags "-X github.com/AlexanderThaller/lablog/cmd.buildTime=`date +%s` -X github.com/AlexanderThaller/lablog/cmd.buildVersion=`git describe --tag --always`" -o "$(NAME)"

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
	mkdir -p benchmarks/`git describe --always`/
	go test -test.benchmem=true -test.bench . 2> /dev/null | tee benchmarks/`git describe --always`/`date +%s`

coverage:
	rm -f coverage.out
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o=/tmp/coverage.html

lint:
	GO15VENDOREXPERIMENT=1 gometalinter `GO15VENDOREXPERIMENT=1 go list ./... | grep -v '/vendor/'`
