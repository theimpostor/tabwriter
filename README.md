# TabWriter
Exposes golang's text/tabwriter as a command line filter

## Installation

#### go 1.16 or later:

```
go install -v github.com/theimpostor/tabwriter@latest
```

#### go 1.15 or earlier:
```
GO111MODULE=on go get github.com/theimpostor/tabwriter@latest
```

This will install the latest version of osc52 to `$GOPATH/bin`. To find out where `$GOPATH` is, run `go env GOPATH`

## Usage
```
Usage: ./tabwriter [file1 [...fileN]]
Copies file contents to stdout passing through text.tabwriter
With no arguments, will read from stdin.
```
