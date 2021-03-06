package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ReadCsvToFloats(path string) [][]float64 {
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return [][]float64{}
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()

	data := make([][]float64, len(records))

	for recordIndex, record := range records {
		data[recordIndex] = StringsToFloats(record)
	}

	return data
}

func StringsToFloats(strings []string) []float64 {
	floats := make([]float64, len(strings))
	for entryIndex, entry := range strings {
		var err error
		floats[entryIndex], err = strconv.ParseFloat(entry, 64)
		if err != nil {
			return []float64{}
		}
	}
	return floats
}

func FloatsToStrings(floats []float64, format string) []string {
	strings := make([]string, len(floats))
	for i, value := range floats {
		strings[i] = fmt.Sprintf(format, value)
	}
	return strings
}
