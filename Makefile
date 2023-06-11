GO_BUILD_ENV := CGO_ENABLED=0

all: build

build:	
	$(GO_BUILD_ENV) go build -v -o ./bin/fulbot .
