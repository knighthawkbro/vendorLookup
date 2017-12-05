package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"vendorLookup/lib"

	"github.com/atotto/clipboard"
)

func main() {
	var usage = `
VendorLookup takes in a file of MAC addresses in any format separated by newlines and spits out all the vendor specified for those mac addresses.

Usage:
	vendorLookup [--copy] <FileName>

Examples:
	# Look up a bulk list of MAC Addresses
	vendorLookup mac.txt

`
	macReg := "(([0-9A-Fa-f]{2}[-:]){5}[0-9A-Fa-f]{2})|(([0-9A-Fa-f]{4}.){2}[0-9A-Fa-f]{4})"
	copy := flag.Bool("copy", false, "Copies a string of MAC addresses from the Clipboard.")

	flag.Usage = func() {
		fmt.Printf(usage)
	}
	flag.Parse()

	var macs []lib.Mac

	if *copy {
		text, err := clipboard.ReadAll()
		lib.CheckErr("Nothing in clipboard", err)
		scanner := bufio.NewScanner(strings.NewReader(text))
		for scanner.Scan() {
			match, _ := regexp.MatchString(macReg, scanner.Text())
			var x lib.Mac
			if match {
				x = lib.Mac{
					Address: scanner.Text(),
				}
			} else {
				fmt.Println("Not A MAC")
				os.Exit(1)
			}

			macs = append(macs, x)
		}
	} else {
		if flag.NArg() == 0 {
			flag.Usage()
			os.Exit(1)
		}

		file := flag.Arg(0)
		f, err := os.Open(file)
		lib.CheckErr("File not Found", err)
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			x := lib.Mac{
				Address: scanner.Text(),
			}
			macs = append(macs, x)
		}
		lib.CheckErr("Couldn't Create Scanner object", scanner.Err())
	}

	out, err := os.Create("out.txt")
	lib.CheckErr("Couldn't create output file", err)
	defer out.Close()

	w := bufio.NewWriter(out)

	for _, m := range macs {
		m.Company = lib.GetResult(m)

		_, err2 := fmt.Fprintf(w, "%v: %v\n", m.Address, m.Company)
		lib.CheckErr("Couldn't write buffer to output", err2)
	}
	w.Flush()
}
