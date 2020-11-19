import XCTest
@testable import projecteuler

final class p17Tests: XCTestCase {
    func testNumLetters() {
        var x = numLetters("three hundred and forty-two")
        XCTAssertEqual(x, 23)
        x = numLetters("one hundred and fifteen")
        XCTAssertEqual(x, 20)
    }
    static var allTests = [
        ("testNumLetters", testNumLetters),
    ]
}

/*
 
 func TestNumToString(t *testing.T) {
         s1 := numToString(342)
         expStr := "three hundred and forty two" // forty two ~= forty-two for our purposes
         if s1 != expStr {
                 t.Errorf("s1 = %v, want %v", s1, expStr)
         }

         s1 = numToString(115)
         expStr = "one hundred and fifteen"
         if s1 != expStr {
                 t.Errorf("s1 = %v, want %v", s1, expStr)
         }
 }

 func TestCount_1_to_1000(t *testing.T) {
         count := letterCount_1_to_1000()
         if count != 21124 {
                 t.Errorf("count = %v, want %v", count, 21124)
         }
 }

 */
