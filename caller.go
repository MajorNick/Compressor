package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var comp, decomp bool
	var path, name string
	flag.BoolVar(&comp, "compress", false, "compress")
	flag.BoolVar(&decomp, "decompress", false, "decompress")
	flag.StringVar(&path, "path", "", "Get Path Of File")
	flag.StringVar(&name, "name", "", "compressed file's name")
	flag.Parse()
	if (comp && decomp) || (!comp && !decomp) {
		GetHelp()
		return
	}
	if comp {
		checkIfFileExists(path)
		if name == "" {
			GetHelp()
			return
		}
		dir, _ := filepath.Split(path)
		_, err := os.Stat(dir + name)
		if err == nil {
			fmt.Println("File Named", name, "Already Exists")
			return
		}

		Compress(path, name)
	}
	if decomp {
		checkIfFileExists(path)
		if name == "" {
			GetHelp()
			return
		}
		dir, _ := filepath.Split(path)
		_, err := os.Stat(dir + name)
		if err == nil {
			fmt.Println("file named ", name, "alreade Exists")
			return
		}

		Decompress(path, name)

	}
}
func GetHelp() {
	fmt.Fprintf(os.Stdout, "Usage: (Option) -path=(\"Actual Path\")\n Options: \n")
	fmt.Fprintf(os.Stdout, " -compress To Compress File  \n -decompress To Decompress File\n")
	return
}
func checkIfFileExists(path string) {
	_, err := os.Stat(path)
	if err != nil {
		GetHelp()
		return
	}
}
