package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Открываем существующий CSV-файл
	inputFile, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()
	/*	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {

		}
	}(inputFile)*/

	// Читаем CSV-файл
	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Изменяем название колонки (например, изменяем "City" на "Location")
	if len(records) > 0 {
		for i, header := range records[0] {
			if header == "City" {
				records[0][i] = "Location"
				break
			}
		}
	}

	// Обновляем значения в колонках (например, добавляем 1 год к возрасту)
	for i := 1; i < len(records); i++ {
		if len(records[i]) > 2 {
			age, err := strconv.Atoi(records[i][2])
			if err != nil {
				fmt.Println("Error converting age:", err)
				return
			}
			records[i][2] = strconv.Itoa(age + 1)
		}
	}

	// Добавляем новые значения
	newRecords := [][]string{
		{"11", "Kara", "28", "Houston", "Scientist"},
		{"12", "Liam", "34", "Austin", "Developer"},
	}
	records = append(records, newRecords...)

	// Создаем или перезаписываем CSV-файл
	outputFile, err := os.Create("example_updated.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()
	/*	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {

		}
	}(inputFile)*/

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Записываем обновленные данные в CSV-файл
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to file:", err)
			return
		}
	}

	fmt.Println("CSV file updated successfully")
}
