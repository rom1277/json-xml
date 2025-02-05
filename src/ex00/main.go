package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("f", "", "path to the file")
	flag.Parse()
	if *fileName == "" {
		fmt.Println("Please provide a filename using -f flag.")
		return
	}
	var reader DBReader
	reader, err := GetTypeDB(*fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// можно не использовать
	//функция reader.Reader(*fileName) уже самостоятельно открывает и закрывает файл
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("не удалось открыть файл:", err)
		return
	}
	defer file.Close()
	recipes, err := reader.Reader(*fileName)
	if err != nil {
		fmt.Println("")
		return
	}
	writer, err := reader.Writer(*recipes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(writer)
}
