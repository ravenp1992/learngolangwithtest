# Go Commands

> Use go run when you want to treat a Go program like a script and run the source code immediately.

```terminal
go run hello.go

```

> Use go build to create a binary that is distributed for other people to use. Most of the time,\
> this is what you want to do. Use the -o flag to give the binary a different name or location.

```terminal
go build hello.go
```

> Or if you want to store it in a different location, use the -o flag. For example, if we wanted\
> to compile our code to a binary called "hello_world", we would use:

```terminal
go build -o hello_word hello.go
```

> This downloads hey and all of its dependencies, builds the program, and installs the binary in\
> your $GOPATH/bin directory.

```terminal
go install github.com/rakyll/hey@latest
```

Even if you haven't seen a Makefile before, it's not too difficult to figure out what is going on.\
Each possible operation is called a _target_. The **.DEFAULT_GOAL** defines which target is run\
When no target is specified. In our ase, we are going to run the build target. Next we have definitions.\
The word before the colon ( : ) is the name of the target. Any words after the target (like vet in the line build: vet)\
are the other targets that must be run before the specified target runs. The tasks that are performed by the target are on\
the indented lines after the target. (The .PHONY line keeps make from getting confused if you ever createt)\
a directory in your project with the same name as a target.)

```Makefile
.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build hello.go
.PHONY:build

```

We can install a secondary Go environment. For example, if you are currently running version 1.22 and\
wanter to try out version 1.22.1, you would use the following commands:

```terminal
go get golang.org/dl/go.1.22.1
go1.22.1 download

<!-- to use -->
go1.22.1 build
```

Once you have validated that your code works, you can deletet the secondary environment by finding its\
GOROOT, deleting it, and then deleting its binary from your $GOPATTH/bin directory.

```terminal
go1.22.1 env GOROOT
rm -rf $(go1.22.1 env GOROOT)
rm $(go env GOPATH)/bin/go1.22.1
```
