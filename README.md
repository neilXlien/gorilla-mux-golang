# gorilla-mux-golang

### Download and Install

#### Binary Distributions

Official binary distributions are available at https://golang.org/dl/.

After downloading a binary release, visit https://golang.org/doc/install
or load [doc/install.html](./doc/install.html) in your web browser for installation
instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://golang.org/doc/install/source or load [doc/install-source.html](./doc/install-source.html)
in your web browser for source installation instructions.

After installing Go, you need to first be inside your GOPATH(The directory where all Go projects will be kept). Your GOPATH will be set automatically once you have Go installed. In your GOPATH, you will have three folders bin, pkg and src. In your src, paste a folder called Desktop, in it you run this command: 

```sh
cd Desktop/go-workspace/src/gorilla-mux
```

## Install
[correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/gorilla/mux
go run main.go
```
listening on http://localhost:3000
