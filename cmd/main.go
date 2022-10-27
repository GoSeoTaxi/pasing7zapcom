package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"pasing7zapcom/internal/constData"
	"pasing7zapcom/internal/initApp"
	"pasing7zapcom/internal/makeListUrl"
	"time"
)

const urlListDirectory = "https://7zap.com/ru/catalog/cars/Audi/A4/Audi%20A4%2FS4%2FAvant%2Fquattro%20%282016%20-%202021%29/UjBRVnVEYlRvdHRBT2pGMThabzFEUT09--/"
const urlListDirectory2 = "https://7zap.com/ru/catalog/cars/Audi/Q7/Audi%20Q7%20%282016%20-%202021%29/bmo1S0FpZzlhK1RDWmdaazdDUzV1Zz09--/"

var end string

func main() {

	fmt.Println(`Запуск`)
	time.Sleep(constData.TimeSleepStart * time.Second)

	records := initApp.ReadCsvFile(constData.InputCSV)
	//	fmt.Println(records)

	//	Старый способ передачи
	//	StartingParsint(ReadLines(constData.FileRead))

	StartingParsint(records)

	fmt.Println("Нажми Enter")
	fmt.Scanf("%s\n", &end)

}

func StartingParsint(lines [][]string) {

	for value := range lines {

		makeListUrl.MakeList(lines[value])

	}

	//for i := 0; i < len(lines); i++ {
	//	fmt.Println(lines[i][0])
	//	fmt.Println(`++`)
	//	fmt.Println(lines[i])
	//	fmt.Println(lines[i][1])
	//}

	/*	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
		makeListUrl.MakeList(lines[i])
	}*/

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
