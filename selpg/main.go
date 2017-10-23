package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"
	"sync"
)


var (
	// lg = log.Println
	// lf = log.Printf
	lg = func(args ...interface{}) {}
	lf = func(args ...interface{}) {}

	fin  = os.Stdin
	fout = os.Stdout

	wg sync.WaitGroup
)

// Default/convention.
const (
	illegalPageIndex        = math.MaxInt32
	illegalLineCountPerPage = -1
	DefaultLineCountPerPage = 72
)

// CommandLine args.
var (
	pageIndexStart int
	pageIndexEnd   int

	lineCountPerPage              = DefaultLineCountPerPage
	separatingPagesByFormfeedChar bool

	printerDestination string
	inputFilePath      string
)

// Helpers depend on CommandLine args. Assume flag.Parsed().
var (
	pageIndexStartAndEndValidator = func() {
		args := os.Args
		lg("os.Args: ", len(args), ", ", args)

		if pageIndexStart == illegalPageIndex || pageIndexEnd == illegalPageIndex {
			goto valueInvalid
		}
		if pageIndexStart <= 0 || pageIndexStart > pageIndexEnd {
			goto valueInvalid
		}

		if !strings.HasPrefix(args[1], "-s") {
			log.Fatalln("[Error] ... Required flag `-s` should be the 1st .")
		} else if strings.HasPrefix(args[1], "-s=") {
			if !strings.HasPrefix(args[2], "-e") {
				goto positionInvalid
			}
		} else {
			if !strings.HasPrefix(args[3], "-e") {
				goto positionInvalid
			}
		}
		return
	positionInvalid:
		log.Fatalln("[Error] ... Required flag `-e` should be the 2nd.")
	valueInvalid:
		log.Fatalln("[Error] ... Values of `-s` and `-e` are invalid.")
	}
	pageSeparationValidator = func() {
		if lineCountPerPage <= 0 {
			log.Fatalln("[Error] ... Number-of-lines-in-per-page could not be <= 0.")
		}
		if lineCountPerPage != DefaultLineCountPerPage && separatingPagesByFormfeedChar {
			log.Fatalln("[Error] ... `-l` is not compatible with `-f`.")
		}
	}
	pageSeparater = map[bool]func(line string) (catchingNewPage bool, offsetEndPrevPage int){
		// use Formfeed-Char
		true: func(line string) (bool, int) {
			for i, u := range line {
				if u == '\f' {
					return true, i
				}
			}
			return false, len(line)
		},
		// use line-count
		false: func(line string) (bool, int) {
			lineCountHandled++
			if lineCountHandled == lineCountPerPage {
				lineCountHandled = 0
				return true, len(line)
			}
			return false, len(line)
		},
	}
	inputRouter = func() {
		// NOTE: may redundant
		switch flag.NArg() {
		case 0:
			fInfo, err := os.Stdin.Stat()
			mode := fInfo.Mode() & os.ModeNamedPipe
			lg("os.Stdin.Stat() --> fInfo = ", fInfo, ", err = ", err, ", fInfo.Size() = ", fInfo.Size())
			lg("fInfo.Mode() = ", fInfo.Mode(), ", os.ModeNamedPipe = ", os.ModeNamedPipe, ", &--> ", mode)
			if err != nil || fInfo.Size() == 0 {
				// log.Fatalln("[Error] ... No input-file, nor pipe-in.")
			}
			// keep fin = os.Stdin, default
		case 1:
			lg("... With exactly 1 Args, set as inputFilePath.")
			inputFilePath = flag.Arg(0)
			if _, err := os.Stat(inputFilePath); os.IsNotExist(err) {
				log.Fatalf("[Error] ... File %v does not exist.\n", inputFilePath)
			}
		default:
			log.Fatalln("[Error] ... Only allow at most 1 input-path.")
		}

		if inputFilePath != "" {
			f, err := os.Open(inputFilePath)
			if err != nil {
				log.Fatal(err)
			}
			fin = f
			// TODO: wrap
		}
	}

	outputRouter = func() {
		if printerDestination != "" {
			// f, err := os.Open("/dev/usb/" + printerDestination) // "/dev/usb/lp1"
			// FIXME: not sure whether works ...
			cmd := exec.Command("lp", "-d", printerDestination)
			cmd.Stdin = fout
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Start(); err != nil {
				log.Fatal(err)
			}
		}
	}
)

// Static, just record working status.
var (
	lineCountHandled int
	pageIndexCurrent = 1
)

func init() {
	// required, 1st 2nd.
	flag.IntVar(&pageIndexStart, "s", illegalPageIndex, "The index of the first page to be printed.")
	flag.IntVar(&pageIndexEnd, "e", illegalPageIndex, "The index of the end page to be printed.")

	// optional.
	flag.IntVar(&lineCountPerPage, "l", DefaultLineCountPerPage, "The custom value for number-of-lines-in-per-page. (not compatible with `-f`)")
	flag.BoolVar(&separatingPagesByFormfeedChar, "f", false, "Choose if the input pages should be separated by Formfeed-Character('\\f'). (not compatible with `-l`)")
	flag.StringVar(&printerDestination, "d", "", "The named-printer to print the selected content.")
}

func main() {
	defer fin.Close()
	defer fout.Close()

	flag.Parse()

	lg(pageIndexStart, pageIndexEnd, lineCountPerPage, separatingPagesByFormfeedChar)
	lg("flag.NArg(), flag.Args() = ", flag.NArg(), flag.Args())

	argsValidator()
	inputRouter()
	outputRouter()

	print()

	wg.Wait()
}

func argsValidator() {
	pageIndexStartAndEndValidator()
	pageSeparationValidator()
}

// TODO: batch buffer reading, prevent read-all-once ...
func readInput() (<-chan string, error) {
	channel := make(chan string)
	scanner := bufio.NewScanner(fin)
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	wg.Add(1)
	go func() {
		for scanner.Scan() {
			channel <- scanner.Text() + "\n"
		}
		defer wg.Done()
		close(channel)
	}()

	return channel, nil
}

func print() {
	writer := bufio.NewWriter(fout)

	inputLines, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	for line := range inputLines {
		catchingNewPage, offsetEndPrevPage := pageSeparater[separatingPagesByFormfeedChar](line)

		if availableToPrint() {
			if _, err := writer.WriteString(line[:offsetEndPrevPage]); err != nil {
				panic(err)
			}
			if catchingNewPage {
				if _, err := writer.WriteString("\f" + line[offsetEndPrevPage:]); err != nil {
					panic(err)
				}
			}
		}
		if catchingNewPage {
			pageIndexCurrent++
		}
	}
	if err := writer.Flush(); err != nil {
		panic(err)
	}
}

func availableToPrint() bool {
	return pageIndexStart <= pageIndexCurrent && pageIndexCurrent <= pageIndexEnd
}
