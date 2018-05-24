VERSION?=0.0.0-dev
FLAGS?="-X main.version=${VERSION}"
GOOS?=darwin

build:
	@go build -ldflags ${FLAGS} -o dist/ckies_${GOOS}

run:
	@go run -ldflags ${FLAGS} main.go ${CMD}

test:
	@ginkgo ckies