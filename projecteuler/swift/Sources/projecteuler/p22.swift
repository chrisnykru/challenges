/*

Names scores

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
containing over five-thousand first names, begin by sorting it into alphabetical
order. Then working out the alphabetical value for each name, multiply this value
by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth
3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain
a score of 938 * 53 = 49714.

What is the total of all the name scores in the file?

*/

import Foundation

func nameScore(_ name: String) -> Int {
    var score = 0
    let capA_Value = "A".utf8.map { $0 }[0]
    for c in name.utf8 {
        score += Int(c) - Int(capA_Value) + 1 // A=1, B=2, etc..
    }
    return score
}

func totalNameScore() throws -> Int {
    let path = Bundle.module.path(forResource: "names", ofType: "txt")!
    let contents = try String(contentsOfFile: path)
    
    var totalScore = 0
    let quoteValue = "\"".utf8.map { $0 }[0]
    var names = contents.split(separator: ",")
    names.sort()
    for i in stride(from: 0, to: names.count, by: 1) {
        let tmp = names[i].utf8.map { $0 }
        if tmp[0] != quoteValue || tmp[tmp.count - 1] != quoteValue {
            throw ProjectEulerError.internalError
        }
        var withoutQuotes = names[i]
        withoutQuotes.removeFirst()
        withoutQuotes.removeLast()
        let score = nameScore(String(withoutQuotes))
        totalScore += (score * (i + 1))
    }
    return totalScore
}
