BINARY=learn-unittest
test:
	go test -v -cover ./...

dev:
	@go mod tidy
	@go mod vendor
	@go build

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -mod=vendor -a -installsuffix cgo -ldflags '-w'

dependencies:
	@echo "> Installing the server dependencies ..."
	@go mod download
	@go mod vendor

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	./learn-unittest api

run-dev: dev run

mock:
	mockgen --source=usecase/posts/repository/post.go -destination=shared/mock/repository/post.go --package mock
	mockgen --source=usecase/posts/presenter/post.go -destination=shared/mock/presenter/post.go --package mock
	mockgen --source=usecase/posts/interactor/post.go -destination=shared/mock/interactor/post.go --package mock
	mockgen --source=usecase/author/repository/author.go -destination=shared/mock/repository/author.go --package mock
	

.PHONY: clean build
