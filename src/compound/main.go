package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("f", "", "path to the file")
	oldfile := flag.String("old", "", "path to the file")
	newfile := flag.String("new", "", "path to the file")
	flag.Parse()
	if *fileName != "" && *oldfile == "" && *newfile == "" { // ex00
		var reader DBReader
		reader, err := GetTypeDB(*fileName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		openfile, err := os.Open(*fileName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer openfile.Close()
		recipes, err := reader.Reader(*fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		writer, err := reader.Writer(*recipes)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(writer)
	} else if *oldfile != "" && *newfile != "" && *fileName == "" { // ex01ex02
		// открываем и закрываем файлы
		openOld, err := os.Open(*oldfile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer openOld.Close()

		openNew, err := os.Open(*newfile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer openNew.Close()
		if strings.HasSuffix(*oldfile, ".xml") && strings.HasSuffix(*newfile, ".json") { //ex01
			// fmt.Println("names:", *oldfile, *newfile, fileName)
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
			menu1, err := file1.Reader(*oldfile)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			menu2, err := file2.Reader(*newfile)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			Сompare(menu1, menu2)
		} else { //ex02
			// сканер для построчного чтения
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
	} else {
		fmt.Println("Error: Incorrect use of flags with files.")
		return
	}
}
