/*

Path sum: two ways

In the 5 by 5 matrix below, the minimal path sum from the top left to the bottom
right, by only moving to the right and down, is indicated in bold red and is
equal to 2427.


(131)  673   234   103   18
(201)  (96) (342)  965   150
 630   803  (746) (422)  111
 537   699   497  (121)  956
 805   732   524  (37)  (331)

Find the minimal path sum, in matrix.txt (right click and
'Save Link/Target As...'), a 31K text file containing a 80 by 80 matrix, from
the top left to the bottom right by only moving right and down.

-------------------------------------------------------------------------------

Strategy:

If we treat as permutation problem w/ down moves and right moves
(e.g., d1,d2,d3,d4,d5,r1,r2,r3,r4,r5)

5x5 grid = 10*9*...*1 = 10! = 3,628,800 permutations

80x80 grid = 160! = 4.7e284.. insane.


Easier and faster method is to on-the-fly matrix reduction.
e.g., take the top left section:

  131  673
  201  96

Reduce:

  ___      673+131
  201+131  96

Note that 201+131 < 673+131

  ___      ___
  ___      96+201+131

-------------------------------------------------------------------------------

*/
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type matrix [][]int

func (m matrix) String() string {
	var b bytes.Buffer
	maxColLen := 0

	for _, row := range m {
		for _, col := range row {
			if colLen := len(fmt.Sprintf("%d", col)); colLen > maxColLen {
				maxColLen = colLen
			}
		}
	}
	for _, row := range m {
		for j, col := range row {
			s := fmt.Sprintf("%d", col)
			spacing := maxColLen - len(s)
			if j != 0 {
				spacing += 2 // from prev. col
			}
			b.WriteString(fmt.Sprintf("%s%s", strings.Repeat(" ", spacing), s))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func newMatrix(r io.Reader, rows, cols int) (m matrix) {
	csvr := csv.NewReader(r)

	records, err := csvr.ReadAll()
	if err != nil {
		log.Panic(err.Error())
	}
	if len(records) != rows {
		log.Panicf("got %d rows, want %d", len(records), rows)
	}

	for i, rec := range records {
		if len(rec) != cols {
			log.Panicf("row %d has %d columns, want %d", i, len(rec), cols)
		}

		row := make([]int, len(rec))
		for j, col := range rec {
			v, err := strconv.ParseInt(col, 10, 0)
			if err != nil {
				log.Panicf("(%v,%v) %s", i, j, err.Error())
			}
			row[j] = int(v)
		}
		m = append(m, row)
	}
	return
}

func matrix5x5() matrix {
	s := "131,673,234,103,18\n"
	s += "201,96,342,965,150\n"
	s += "630,803,746,422,111\n"
	s += "537,699,497,121,956\n"
	s += "805,732,524,37,331\n"
	return newMatrix(bytes.NewBufferString(s), 5, 5)
}

func matrix80x80() matrix {
	f, err := os.Open("matrix.txt")
	if err != nil {
		log.Panic(err.Error())
	}
	return newMatrix(f, 80, 80)
}

func (m matrix) merge(row, col int) {
	lesser := m[row-1][col]
	if m[row][col-1] < lesser {
		lesser = m[row][col-1]
	}
	m[row][col] += lesser
}

// Matrix must be square!
func reduce(m matrix) int {
	// propagate along row 0 and column 0 (simple case--no merging)
	for i := 1; i < len(m); i++ {
		m[0][i] += m[0][i-1]
		m[i][0] += m[i-1][0]
	}

	merged := make([][]bool, len(m))
	for i, row := range m {
		if len(row) != len(m) {
			panic("matrix is not square!")
		}
		merged[i] = make([]bool, len(m))
	}

	// suboptimal iteration for sure
	for row := 1; row < len(m); row++ {
		for col := 1; col < len(m); col++ {
			if merged[row][col] {
				continue
			}
			m.merge(row, col)
			merged[row][col] = true
		}
	}

	lastRow := m[len(m)-1]
	return lastRow[len(lastRow)-1]
}

func main() {
	fmt.Println("reduce:", reduce(matrix80x80()))
}
