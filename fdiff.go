package main

import "fmt"

func helptext() {
	fmt.Println("fdiff       version dev 171020\n")
	fmt.Println("Usage: fdiff [OPTION]... FILE1/DIR1 FILE2/DIR2\n\n")
	fmt.Println("\t-v, --verbose         be more verbose")
	fmt.Println("\t-q, --quite           be more quite")
	fmt.Println("\t-d, --directory       diff directories")
	fmt.Println("\t-p, --patch           create patch after diff")
	fmt.Println("\t-o, --output <dir>    set patch output directory")
	fmt.Println("\t-u, --update <dir>    update FILE1/DIR1 with patch in <dir>")
}

func main() {
	helptext();
}
