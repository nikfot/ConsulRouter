# Build time info
PACKAGE = nikfot/ConsulRouter
VERSION=$(shell git describe --tags)
BUILD=$(shell git rev-parse HEAD)
DATE=$(shell git show -s --format=%ci ${BUILD})
NAME?=consulrouter
DNSSERVER?=10.0.0.0

REGISTRY = 'nikfot'

# Binary output file
BINARY  = consulrouter

# Setup ldflags
LDFLAGS = -ldflags "-X '${PACKAGE}/health.Version=${VERSION}' -X '${PACKAGE}/health.Build=${BUILD}' -X '${PACKAGE}/health.Date=${DATE}'"

${BINARY}: main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -v -a -o ${BINARY}

tools:
	go get -insecure gopkg.in/yaml.v2
	go get -insecure github.com/satori/go.uuid
	
clean:
	if [ -f ${BINARY} ]; then rm -rf ${BINARY}; fi



docker:
	docker build --no-cache --label VERSION=$(VERSION) --tag nikfot/${BINARY}:latest .
	docker tag  nikfot/${BINARY}  ${REGISTRY}/${BINARY}:$(VERSION)
	docker push nikfot/${BINARY}:$(VERSION)
	docker push nikfot/${BINARY}:latest

run:
	docker run -d --name ${NAME} --dns ${DNSSERVER} -p 8080:8080 nikfot/consulrouter:latest
