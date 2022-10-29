package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"pasing7zapcom/internal/constData"
	"pasing7zapcom/internal/endApp"
	"pasing7zapcom/internal/initApp"
	"pasing7zapcom/internal/makeListUrl"
	"time"
)

func main() {

	fmt.Println(`Запуск`)
	time.Sleep(constData.TimeSleepStart * time.Second)
	records, err := initApp.ReadCsvFile(constData.InputCSV)
	if err != nil {
		endApp.Fin()
	}

	StartingParsint(records)
	endApp.Fin()
}

func StartingParsint(lines [][]string) {
	for value := range lines {
		makeListUrl.MakeList(lines[value])
	}
}

func ReadLines(path string) (lines []string) {

	var abcd []byte
	lineFiles, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) || lineFiles.Size() == 0 {
		ioutil.WriteFile(path, abcd, 0644)
		log.Fatal(`Пустой файл с задачами`)
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
			log.Fatal(err)
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines

}

func WriteLines(path string, lines []string) error {

	file, err := os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
			return err
		}
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		if line != "" {
			fmt.Fprintln(w, line)
		}
	}
	return w.Flush()

}
