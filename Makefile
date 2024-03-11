
FILEPATH = ${HOME}/.local/share/todo-cli/todos.json
path 	 = ./cmd/main.go
binary   = bin/todo-cli
version  = 1.0.0
build	   = $(shell git rev-parse HEAD)
ldflags  = -ldflags "-X 'github.com/Tecu23/todo-cli/cmd/command.version=$(version)'
ldflags += -X 'github.com/Tecu23/todo-cli/cmd/command.build=$(build)'
ldflags += -X 'github.com/Tecu23/todo-cli/cmd/command.todoFile=$(FILEPATH)'"

build:
	@go build -o $(binary) $(ldflags) ${path}

install: export GOPATH=${HOME}/go
install:
	@mkdir -p ${HOME}/.local/share/todo-cli
	@touch ${HOME}/.local/share/todo-cli/todos.json
	@mv $(binary) $(GOPATH)/${binary}

uninstall: export GOPATH=${HOME}/go
uninstall:
	@rm -rf ${HOME}/.local/share/todo-cli
	@rm ${GOPATH}/${binary}

test:
	@go test ./... -cover -coverprofile c.out
	@go tool cover -html=c.out -o coverage.html

clean:
	@rm -rf $(binary) c.out coverage.html


