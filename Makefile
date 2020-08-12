MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
PACKAGES:=$(shell go list ./... | sed -n '1!p' | grep -v /vendor/)
LDFLAGS:=-ldflags "-X main.go=${VERSION}"

default: run

depends:
	../../../../bin/glide up

test:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

cover: test
	go tool cover -html=coverage-all.out

run:
	cd locales && wti pull -f kyloren.* && cd ..
	go run ${LDFLAGS} main.go

build: clean
	go build ${LDFLAGS} -a -o main main.go

clean:
	rm -rf main coverage.out coverage-all.out

coverage: 
	go tool cover -html=coverage.out

wti: 
	cd locales && wti pull kyloren*