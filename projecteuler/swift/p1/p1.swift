/*

Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

*/

import XCTest


// 
func sumOfMultiples(below belowNumber: Int) -> Int { 
  var sum = 0 
  for n in 1...belowNumber-1 { 
    if n % 3 == 0 || n % 5 == 0 { 
      sum += n 
    } 
  } 
  return sum 
}

final class SumOfMultiplesTests: XCTestCase {
  func test1000() {
    let x = sumOfMultiples(below: 1000)
    XCTAssertEqual(x, 233168)
  }
}

print(sumOfMultiples(below: 1000))