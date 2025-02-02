package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Заголовки колонок
	headers := []string{"ID", "Name", "Age", "City", "Occupation"}
	if err := writer.Write(headers); err != nil {
		fmt.Println("Error writing headers to file:", err)
		return
	}

	// Данные
	records := [][]string{
		{"1", "Alice", "30", "New York", "Engineer"},
		{"2", "Bob", "25", "San Francisco", "Designer"},
		{"3", "Charlie", "35", "Los Angeles", "Teacher"},
		{"4", "David", "28", "Chicago", "Doctor"},
		{"5", "Eva", "32", "Boston", "Artist"},
		{"6", "Frank", "45", "Seattle", "Manager"},
		{"7", "Grace", "29", "Miami", "Nurse"},
		{"8", "Hank", "33", "Dallas", "Technician"},
		{"9", "Ivy", "26", "Denver", "Lawyer"},
		{"10", "Jack", "31", "Atlanta", "Chef"},
	}

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to file:", err)
			return
		}
	}

	fmt.Println("CSV file created successfully")
}
