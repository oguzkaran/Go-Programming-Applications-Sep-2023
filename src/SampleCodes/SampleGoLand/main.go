/*
------------------------------------------------------------------------------------------------------------------------

	Aşağıdaki demo örneği inceleyiniz

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"os"
)

func printFileInformation(de os.DirEntry) {
	fi, err := de.Info()
	if err == nil {
		if fi.IsDir() {
			fmt.Printf("%s <DIR>\n", fi.Name())
		} else {
			fmt.Printf("%s %d\n", fi.Name(), fi.Size())
		}

	} else {
		fmt.Printf("Error:%s\n", err.Error())
	}

}

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Wrong number of arguments!...\n")
		os.Exit(1)
	}

	if modfile.IsDirectoryPath(os.Args[1]) {
		entries, err := os.ReadDir(os.Args[1])
		if err == nil {
			for _, de := range entries {
				printFileInformation(de)
			}
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "ReadDir:%s\n", err.Error())
		}
	} else {
		fmt.Println("Is not a directory")
	}
}
