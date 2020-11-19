/*

Number letter counts

If the numbers 1 to 5 are written out in words: one, two, three, four, five,
then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out
in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
forty-two) contains 23 letters and 115 (one hundred and fifteen) contains
20 letters. The use of "and" when writing out numbers is in compliance with
British usage.

*/

func numLetters(_ s: String) -> Int {
    var count = 0
    for char in s {
        if char >= "a" && char <= "z" {
            count += 1
        }
    }
    return count
}

func digitToWord(_ x: Int) throws -> String {
    switch x {
    case 9:
        return "nine"
    case 8:
        return "eight"
    case 7:
        return "seven"
    case 6:
        return "six"
    case 5:
        return "five"
    case 4:
        return "four"
    case 3:
        return "three"
    case 2:
        return "two"
    case 1:
        return "one"
    case 0:
        return "zero"
    default:
        throw ProjectEulerError.outOfRange
    }
}

func numToString(_ x: Int) throws -> String {
    guard x >= 1 && x <= 1000 else {
        throw ProjectEulerError.outOfRange
    }
    var s = ""
    var tmp = x
    while tmp > 0 {
        if tmp >= 1000 {
            if s.count > 0 {
                s += " "
            }
            try s += digitToWord(tmp / 1000) + " thousand"
            tmp %= 1000
        }
        
        else if tmp >= 100 {
            if s.count > 0 {
                s += " "
            }
            try s += digitToWord(tmp / 100) + " hundred"
            tmp %= 100
            if tmp > 0 {
                s += " and"
            }
        
        }
        
        // [20,99]
        else if tmp >= 20 {
            var s2 = ""
            switch tmp / 10 {
            case 9:
                s2 = "ninety"
            case 8:
                s2 = "eighty"
            case 7:
                s2 = "seventy"
            case 6:
                s2 = "sixty"
            case 5:
                s2 = "fifty"
            case 4:
                s2 = "forty"
            case 3:
                s2 = "thirty"
            case 2:
                s2 = "twenty"
            default:
                throw ProjectEulerError.internalError
            }
            if s.count > 0 {
                s += " "
            }
            s += s2
            tmp %= 10
        }
        
        // [10,19]
        else if tmp >= 10 {
            var s2 = ""
            switch tmp {
            case 19:
                    s2 = "nineteen"
            case 18:
                    s2 = "eighteen"
            case 17:
                    s2 = "seventeen"
            case 16:
                    s2 = "sixteen"
            case 15:
                    s2 = "fifteen"
            case 14:
                    s2 = "fourteen"
            case 13:
                s2 = "thirteen"
            case 12:
                s2 = "twelve"
            case 11:
                s2 = "eleven"
            case 10:
                s2 = "ten"
            default:
                throw ProjectEulerError.internalError
            }
            if s.count > 0 {
                s += " "
            }
            s += s2
            break
        }
        
        // [1,9]
        else {
            if s.count > 0 {
                s += " "
            }
            try s += digitToWord(tmp)
            break
        }
    }
    return s
}

func letterCount1To1000() throws -> Int {
    var count = 0
    for i in stride(from: 1, through: 1000, by: 1) {
        var s: String
        try s = numToString(i)
        count += numLetters(s)
    }
    return count
}

/*
 
func letterCount_1_to_1000() int {
 count := 0
 for i := 1; i <= 1000; i++ {
         s := numToString(i)
         count += numLetters(s)
 }
 return count
}
 
*/
