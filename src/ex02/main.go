package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	oldfile := flag.String("old", "", "path to the file")
	newfile := flag.String("new", "", "path to the file")
	flag.Parse()

	if *oldfile == "" || *newfile == "" {
		fmt.Println("Please specify two file names using -old and -new flags.")
	}
	openOld, err := os.Open(*oldfile)
	if err != nil {
		fmt.Println("не удалось открыть файл:", err)
		return
	}
	defer openOld.Close()

	openNew, err := os.Open(*newfile)
	if err != nil {
		fmt.Println("не удалось открыть файл:", err)
		return
	}
	defer openNew.Close()

	// Создаем сканер для построчного чтения
	scannerOld := bufio.NewScanner(openOld)
	datamap := make(map[string]struct{})
	var line string
	// Считываем файл построчно
	for scannerOld.Scan() {
		line = strings.TrimSpace(scannerOld.Text()) // Получаем строку до '\n'
		if line == "" {
			continue
		}
		datamap[line] = struct{}{}
	}

	scannerNew := bufio.NewScanner(openNew)
	for scannerNew.Scan() {
		line = strings.TrimSpace(scannerNew.Text())
		if line == "" {
			continue
		}
		if _, ok := datamap[line]; !ok {
			fmt.Println("ADDED", line)
		} else {
			delete(datamap, line)
		}
	}
	for key := range datamap {
		fmt.Println("REMOVED", key)
	}
}
