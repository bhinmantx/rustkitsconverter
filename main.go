package main

import (
	"flag"
	"fmt"
)

func main() {
	var oldkits_filename, output_filename string
	fmt.Println("About to parse vars. Guess it's slow when it's empty?")
	flag.StringVar(&oldkits_filename, "old", "Kits.json", "name of the old kits file")
	flag.StringVar(&output_filename, "new", "kits_data.json", "name of the new kits file")
	flag.Parse()

	fmt.Println(oldkits_filename)

	extracted := extractOldKits(oldkits_filename)
	fmt.Println(output_filename)
	convertOldtoNew(extracted, output_filename)
}
