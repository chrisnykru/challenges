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

import Foundation

func nameScore(_ name: String) -> Int {
    var score = 0
    let capA_Value = "A".utf8.map { $0 }[0]
    for c in name.utf8 {
        //print(c, capA_Value)
        // A needs to equal 1, B needs to equal 2, etc..
        score += Int(c) - Int(capA_Value) + 1
    }
    return score
}

/*
 
 func parseTriangle(_ s: String) throws -> [[Int]] {
     var tri: [[Int]] = []
     let rows = s.split(separator: "\n")
     var lastNumCols = 0
     for i in stride(from: 0, to: rows.count, by: 1) {
         let rowCols = rows[i].split(separator: " ")
         guard rowCols.count == lastNumCols + 1 else {
             throw ProjectEulerError.internalError // not our desired triangle shape
         }
         lastNumCols = rowCols.count
         tri.append(rowCols.map({ (v: Substring.SubSequence) -> Int in
             return Int(v)!
         }))
     }
     return tri
 }
 */

/*
func totalNameScore() throws -> Int {
    let path = Bundle.module.path(forResource: "names", ofType: "txt")!
    print("resource path:", path)
    let contents = try String(contentsOfFile: path)
    
    var totalScore = 0
    let names = contents.split(separator: ",")
    for n in names {
        
        
    }
    return totalScore
}
 */


/*

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
