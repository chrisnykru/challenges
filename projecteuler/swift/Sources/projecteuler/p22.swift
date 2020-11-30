/*

Names scores

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
containing over five-thousand first names, begin by sorting it into alphabetical
order. Then working out the alphabetical value for each name, multiply this value
by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth
3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain
a score of 938  53 = 49714.

What is the total of all the name scores in the file?

*/

func nameScore(_ name: String) -> Int {
    var score = 0
    let capA_Value = "A".unicodeScalars.map { $0.value }[0]
    for c in name.utf8 {
        print(c, capA_Value)
        score += Int(c) - Int(capA_Value) + 1 // A needs to equal 1, B needs to equal 2, etc..
    }
    return score
}


/*
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "sort"
)

func nameScore(name string) int64 {
    score := int64(0)
    for _, v := range name {
        letterScore := int64(v - 'A' + 1)
        score += letterScore
    }
    return score
}

func totalNameScores(r io.Reader) int64 {
    csv := csv.NewReader(r)
    records, err := csv.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    if len(records) != 1 {
        log.Fatal(len(records))
    }

    // sort
    sort.Strings(records[0])

    totalScore := int64(0)
    for i, col := range records[0] {
        totalScore += int64(i+1) * nameScore(col)
    }
    return totalScore
}

func main() {
    f, err := os.Open("names.txt")
    if err != nil {
        log.Fatal(err)
    }
    totalScore := totalNameScores(f)
    fmt.Printf("total score = %v\n", totalScore)
}
*/
