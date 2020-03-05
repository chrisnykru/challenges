/*

Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

*/

func sumOfMultiples(below belowNumber: Int) -> Int { 
  var sum = 0 
  for n in 1...belowNumber-1 { 
    if n % 3 == 0 || n % 5 == 0 { 
      sum += n 
    } 
  } 
  return sum 
}

func fasterSumOfMultiples(below belowNumber: Int) -> Int {
    var sum = 0
    var i = 3
    while i < belowNumber {
        sum += i
        i += 3
    }
    i = 5
    while i < belowNumber {
        sum += i
        i += 5
    }
    // deal with overcounting common factors
    i = 15
    while i < belowNumber {
        sum -= i
        i += 15
    }
    return sum
}