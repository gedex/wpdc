// Copyright 2016 The wpdc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"go/scanner"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gedex/wpdc"
)

var exitCode = 0

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runWpdc()
	os.Exit(exitCode)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: wpdc [path ...]\n")
}

func report(err error) {
	scanner.PrintError(os.Stderr, err)
	exitCode = 2
}

func processFile(filename string, in io.Reader, out io.Writer, stdin bool) error {
	if in == nil {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	}

	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	res := wpdc.Check(src)
	if len(res) > 0 {
		fmt.Fprintf(out, "Found %d deprecateds being used in %s:\n", len(res), filename)
		for _, r := range res {
			fmt.Fprintf(out, "* line %d uses deprecated `%s` as listed in %s\n", r.Line, r.DeprecatedName, r.DeprecatedSource)
		}
		fmt.Fprintln(out, "------------\n")
	}

	return nil
}

func walkDir(path string) {
	filepath.Walk(path, visitFile)
}

func visitFile(path string, f os.FileInfo, err error) error {
	if err == nil && isPHPFile(f) {
		err = processFile(path, nil, os.Stdout, false)
	}
	if err != nil {
		report(err)
	}
	return nil
}

func isPHPFile(f os.FileInfo) bool {
	// Ignore non-PHP files.
	name := f.Name()

	return !f.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".php")
}

var parseFlags = func() []string {
	flag.Parse()
	return flag.Args()
}

func runWpdc() {
	flag.Usage = usage
	paths := parseFlags()

	if len(paths) == 0 {
		if err := processFile("<standard input>", os.Stdin, os.Stdout, true); err != nil {
			report(err)
		}
		return
	}

	for _, path := range paths {
		switch dir, err := os.Stat(path); {
		case err != nil:
			report(err)
		case dir.IsDir():
			walkDir(path)
		default:
			if err := processFile(path, nil, os.Stdout, false); err != nil {
				report(err)
			}
		}
	}
}
