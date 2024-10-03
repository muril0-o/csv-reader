package src_reader

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Reader() {
	file, err := os.Open("/home/murilooliveira/Projetos/csv-reader/people-10000.csv")
	if err != nil {
		log.Fatalf("failed to read CSV file: %s", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read CSV file: %s", err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
