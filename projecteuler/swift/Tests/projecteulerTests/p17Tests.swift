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
    func testCount1To1000() {
        do {
            var x: Int
            try x = letterCount1To1000()
            XCTAssertEqual(x, 21124)
        } catch {
            XCTFail()
        }
    }
    static var allTests = [
        ("testNumLetters", testNumLetters),
        ("testNumToString", testNumToString),
        ("testCount1To1000", testCount1To1000),
    ]
}
