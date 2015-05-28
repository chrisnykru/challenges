package triangle

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

func Parse(r io.Reader, expectedRows int) (tri [][]int) {
	csvr := csv.NewReader(r)
	csvr.Comma = ' '
	csvr.FieldsPerRecord = -1
	records, err := csvr.ReadAll()
	if err != nil {
		log.Panic(err.Error())
	}

	if len(records) != expectedRows {
		log.Panicf("got %v rows, want %v", len(records), expectedRows)
	}
	if len(records) < 1 || len(records[0]) != 1 {
		log.Panicf("initial row must have 1 column")
	}

	for i, row := range records {
		if i != 0 && len(row) != len(records[i-1])+1 {
			log.Panicf("got %v columns, want %v", len(row), len(records[i-1])+1)
		}

		triRow := make([]int, len(row))
		for j, col := range row {
			v, err := strconv.ParseInt(col, 10, 0)
			if err != nil {
				log.Panicf("(%v,%v) %s", i, j, err.Error())
			}
			triRow[j] = int(v)
		}
		tri = append(tri, triRow)
	}
	return
}
