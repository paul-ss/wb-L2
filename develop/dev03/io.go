package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type Flags struct {
	ColumnSort    int
	NumberSort    bool
	ReverseSort   bool
	UniqueStrings bool
	FileNames     []string
}

func ParceFlags() *Flags {
	var flags Flags
	flag.IntVar(&flags.ColumnSort, "k", 0, "specify column number")
	flag.BoolVar(&flags.NumberSort, "n", false, "sort by number")
	flag.BoolVar(&flags.ReverseSort, "r", false, "reverse sort order")
	flag.BoolVar(&flags.UniqueStrings, "u", false, "delete repetitions")

	flag.Parse()
	flags.FileNames = flag.Args()

	return &flags
}

func ReadLines(f *Flags) ([]string, error) {
	var files []io.Closer
	var readers []io.Reader
	defer func() {
		for _, file := range files {
			if err := file.Close(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}()

	for _, name := range f.FileNames {
		file, err := os.Open(name)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
		readers = append(readers, file)
	}

	scanner := bufio.NewScanner(io.MultiReader(readers...))

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

func WriteLines(lines []string) {
	for _, l := range lines {
		fmt.Println(l)
	}
}
