package internal

import (
	"log"
	"os"

	"encoding/csv"
)

func Write(cities []City, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	emptyCity := City{}
	var rows [][]string
	rows = append(rows, emptyCity.Header())

	for _, city := range cities {
		if err := w.Write(city.ToRow()); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
