package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"
)

func filterFn(fname string, filter io.Writer) {
	var f *os.File
	var err error

	if fname == "-" {
		f = os.Stdin
	} else {
		if f, err = os.Open(fname); err != nil {
			log.Fatalf("Failed to open file %v: %v", fname, err)
		} else {
			defer f.Close()
		}
	}

	if _, err = io.Copy(filter, f); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var fnames []string

	flag.Usage = func() {
		template := `Usage:
%s [file1 [...fileN]]
Copies file contents to stdout passing through text/tabwriter.
With no arguments, will read from stdin.
`
		fmt.Fprintf(flag.CommandLine.Output(), template, os.Args[0])
		flag.PrintDefaults()
	}

	var minwidth, tabwidth, padding int
	var padchar string
	flag.IntVar(&minwidth, "minwidth", 0, "minimal cell width including any padding")
	flag.IntVar(&tabwidth, "tabwidth", 0, "width of tab characters (equivalent number of spaces)")
	flag.IntVar(&padding, "padding", 0, "padding added to a cell before computing its width")
	flag.StringVar(&padchar, "padchar", "", "ASCII char used for padding")
	filterHTML := flag.Bool("filterHTML", false, "see https://pkg.go.dev/text/tabwriter")
	stripEscape := flag.Bool("stripEscape", false, "see https://pkg.go.dev/text/tabwriter")
	alignRight := flag.Bool("alignRight", false, "see https://pkg.go.dev/text/tabwriter")
	discardEmptyColumns := flag.Bool("discardEmptyColumns", false, "see https://pkg.go.dev/text/tabwriter")
	tabIndent := flag.Bool("tabIndent", false, "see https://pkg.go.dev/text/tabwriter")
	debug := flag.Bool("debug", false, "see https://pkg.go.dev/text/tabwriter")

	flag.Parse()

	var flags uint
	if *filterHTML {
		flags |= tabwriter.FilterHTML
	}
	if *stripEscape {
		flags |= tabwriter.StripEscape
	}
	if *alignRight {
		flags |= tabwriter.AlignRight
	}
	if *discardEmptyColumns {
		flags |= tabwriter.DiscardEmptyColumns
	}
	if *tabIndent {
		flags |= tabwriter.TabIndent
	}
	if *debug {
		flags |= tabwriter.Debug
	}

	if len(flag.Args()) > 0 {
		fnames = flag.Args()
	} else {
		fnames = []string{"-"}
	}

	out := tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar[0], flags)

	for _, fname := range fnames {
		filterFn(fname, out)
	}

	out.Flush()
}
