package main

import (
	"bufio"
	"fmt"
	"mac/lib"
	"os"
)

func main() {
	args := os.Args
	usage := "usage: mac Filename"

	if len(args) == 1 {
		fmt.Printf("No file specified\n\n%v\n", usage)
		return
	} else if len(args) > 2 {
		fmt.Printf("Too many arguments\n\n%v\n", usage)
		return
	} else if args[1] == "-h" || args[1] == "--help" {
		fmt.Printf("%v\n", usage)
		return
	}

	file, err := os.Open(args[1])
	lib.CheckErr("File not Found", err)

	out, err := os.Create("out.txt")
	lib.CheckErr("Couldn't create output file", err)

	defer file.Close()
	defer out.Close()

	var macs []lib.Mac

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := lib.Mac{
			Address: scanner.Text(),
		}
		macs = append(macs, x)
	}

	lib.CheckErr("Couldn't Create Scanner object", scanner.Err())

	w := bufio.NewWriter(out)

	for _, m := range macs {
		m.Company = lib.GetResult(m)

		_, err2 := fmt.Fprintf(w, "%v: %v\n", m.Address, m.Company)
		lib.CheckErr("Couldn't write buffer to output", err2)
	}
	w.Flush()
}
