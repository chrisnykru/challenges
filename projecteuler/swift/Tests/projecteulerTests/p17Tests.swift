import XCTest
@testable import projecteuler

final class p17Tests: XCTestCase {
    func testNumLetters() {
        var x = numLetters("three hundred and forty-two")
        XCTAssertEqual(x, 23)
        x = numLetters("one hundred and fifteen")
        XCTAssertEqual(x, 20)
    }
    func testNumToString() {
        do {
            var s = try numToString(342)
            XCTAssertEqual(s, "three hundred and forty two")
            try s = numToString(115)
            XCTAssertEqual(s, "one hundred and fifteen")
        } catch {
            XCTFail()
        }
    }
    static var allTests = [
        ("testNumLetters", testNumLetters),
        ("testNumToString", testNumToString),
    ]
}

/*
 
 func TestCount_1_to_1000(t *testing.T) {
         count := letterCount_1_to_1000()
         if count != 21124 {
                 t.Errorf("count = %v, want %v", count, 21124)
         }
 }

 */
