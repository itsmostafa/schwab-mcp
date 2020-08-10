# Go Fundamentals

### The go build Command

- To produce an executable binary file, run this in the same directory with `main.go`:
    - `go build`

- To reduce the file size, you can add these additional flags to strip debugging information and the symbol table:
    - `go build -ldflags "-w -s"`

### Cross-Compiling

- Cross-Compiling allows you to create a binary that can run on different architecture. 

- To cross-compile, you need to set the constraints `GOOS` (operating system) and `GOARCH` (architecture).
    - For example:
        - `GOOS="linux" GOARCH="amd64" go build hello.go`