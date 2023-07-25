package main

import (
	"assigment2/sorter"
	"flag"
	"fmt"
)

func main() {
	ints := flag.Bool("int", false, "sort integers")
	floats := flag.Bool("float", false, "sort floats")
	strings := flag.Bool("string", false, "sort strings")
	mix := flag.Bool("mix", false, "sort mixed types")
	flag.Parse()

	if *ints {
		fmt.Println(sorter.SortInts(flag.Args()))
	} else if *floats {
		fmt.Println(sorter.SortFloats(flag.Args()))
	} else if *strings {
		fmt.Println(sorter.SortStrings(flag.Args()))
	} else if *mix {
		numbers, words := sorter.SortMix(flag.Args())
		for _, num := range numbers {
			fmt.Print(num, " ")
		}
		for _, word := range words {
			fmt.Print(word, " ")
		}
		fmt.Println()
	}
}
