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

This will install the latest version of tabwriter to `$GOPATH/bin`. To find out where `$GOPATH` is, run `go env GOPATH`

## Usage
```
Usage: tabwriter [file1 [...fileN]]
Copies file contents to stdout passing through text/tabwriter.
With no arguments, will read from stdin.
  -alignRight
        see https://pkg.go.dev/text/tabwriter
  -debug
        see https://pkg.go.dev/text/tabwriter
  -discardEmptyColumns
        see https://pkg.go.dev/text/tabwriter
  -filterHTML
        see https://pkg.go.dev/text/tabwriter
  -minwidth int
        minimal cell width including any padding
  -padchar string
        ASCII char used for padding
  -padding int
        padding added to a cell before computing its width
  -stripEscape
        see https://pkg.go.dev/text/tabwriter
  -tabIndent
        see https://pkg.go.dev/text/tabwriter
  -tabwidth int
        width of tab characters (equivalent number of spaces)
```
