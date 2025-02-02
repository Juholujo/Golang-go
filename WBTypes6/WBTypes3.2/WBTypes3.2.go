package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os"
)

func main() {
	// Открываем CSV-файл
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Читаем данные в DataFrame
	df := dataframe.ReadCSV(file)
	fmt.Println(df)
	// Удаляем колонку Age
	df = df.Drop("Age")

	// Записываем обновленный DataFrame обратно в CSV
	outFile, err := os.Create("example_without_age_gota.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	if err := df.WriteCSV(outFile); err != nil {
		fmt.Println("Error writing CSV file:", err)
	}

	fmt.Println("CSV file updated successfully with Gota")

	fmt.Println("CSV file updated", df)
}
