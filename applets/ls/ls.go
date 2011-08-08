package ls

import (
	"flag"
	"io/ioutil"
	"os"
	"os/user"
	"tabwriter"
	"fmt"
	"math"
)

var (
	flagSet       = flag.NewFlagSet("ls", flag.PanicOnError)
	longFlag      = flagSet.Bool("l", false, "Long, detailed listing")
	recursiveFlag = flagSet.Bool("r", false, "Recurse into directories")
	humanFlag     = flagSet.Bool("h", false, "Output sizes in a human readable format")
	out           = tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
)

func Ls(call []string) os.Error {
	flagSet.Parse(call[1:])

	if flagSet.NArg() == 0 {
		cwd, e := os.Getwd()
		if e != nil {
			return e
		}
		e = list(cwd, "")
		if e != nil {
			return e
		}
	} else {
		for _, dir := range flagSet.Args() {
			e := list(dir, "")
			if e != nil {
				return e
			}
		}
	}
	out.Flush()
	return nil
}

func list(dir, prefix string) os.Error {
	entries, e := ioutil.ReadDir(dir)
	if e != nil {
		return e
	}

	for _, entry := range entries {
		printEntry(entry)
		if entry.IsDirectory() && *recursiveFlag {
			folder := prefix + "/" + entry.Name
			fmt.Fprintf(out, "%s:\n", folder)
			e := list(dir+"/"+entry.Name, folder)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

func printEntry(e *os.FileInfo) {
	fmt.Fprintf(out, "%s%s\t", e.Name, getEntryTypeString(e))
	if *longFlag {
		fmt.Fprintf(out, "%s\t", getModeString(e.Mode))
		fmt.Fprintf(out, "%s\t", getSizeString(e.Size))
		fmt.Fprintf(out, "%s\t", getUserString(e.Uid))
	}
	fmt.Fprintln(out, "")
}

var accessSymbols = "xwr"

func getModeString(mode uint32) (s string) {
	for i := 8; i >= 0; i-- {
		if mode&(1<<uint(i)) == 0 {
			s += "-"
		} else {
			char := i % 3
			s += accessSymbols[char : char+1]
		}
	}
	return
}

var sizeSymbols = "BkMGT"

func getSizeString(size int64) (s string) {
	if !*humanFlag {
		return fmt.Sprintf("%9dB", size)
	}
	var power int
	if size == 0 {
		power = 0
	} else {
		power = int(math.Log(float64(size)) / math.Log(1024.0))
	}
	if power > len(sizeSymbols)-1 {
		power = len(sizeSymbols) - 1
	}
	rSize := float64(size) / math.Pow(1024, float64(power))
	return fmt.Sprintf("%7.3f%s", rSize, sizeSymbols[power:power+1])
}

func getEntryTypeString(e *os.FileInfo) string {
	if e.IsDirectory() {
		return "/"
	} else if e.IsBlock() {
		return "<>"
	} else if e.IsFifo() {
		return ">>"
	} else if e.IsSymlink() {
		return "@"
	} else if e.IsSocket() {
		return "&"
	} else if e.IsRegular() && (e.Mode&0001 == 0001) {
		return "*"
	}
	return ""
}

func getUserString(id int) string {
	username, e := user.LookupId(id)
	if e != nil {
		return "???"
	}
	return username.Username
}