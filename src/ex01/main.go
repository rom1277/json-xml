package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	oldfile := flag.String("old", "", "path to the file")
	newfile := flag.String("new", "", "path to the file")
	flag.Parse()

	if *oldfile == "" || *newfile == "" {
		fmt.Println("Please specify two file names using -old and -new flags.")
	}
	var file1, file2 DBReader
	file1, err := GetTypeDB(*oldfile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	file2, err = GetTypeDB(*newfile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	openfile1, err := os.Open(*oldfile)
	if err != nil {
		fmt.Println("не удалось открыть файл:", err)
		return
	}
	defer openfile1.Close()
	menu1, err := file1.Reader(*oldfile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	openfile2, err := os.Open(*newfile)
	if err != nil {
		fmt.Println("не удалось открыть файл:", err)
		return
	}
	defer openfile2.Close()
	menu2, err := file2.Reader(*newfile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	Сompare(menu1, menu2)
}
