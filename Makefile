
path 	 = ./cmd/main.go
binary   = bin/todo-cli
version  = 0.0.1
build	   = $(shell git rev-parse HEAD)
ldflags  = -ldflags "-X 'github.com/Tecu23/todo-cli/cmd/command.version=$(version)'
ldflags += -X 'github.com/Tecu23/todo-cli/cmd/command.build=$(build)'"

build:
	go build -o $(binary) $(ldflags) ${path}

install: export GOPATH=$${HOME}/go
install:
	mv $(binary) $(GOPATH)/${binary}

test:
	go test ./... -cover -coverprofile c.out
	go tool cover -html=c.out -o coverage.html

clean:
	rm -rf $(binary) c.out coverage.html
