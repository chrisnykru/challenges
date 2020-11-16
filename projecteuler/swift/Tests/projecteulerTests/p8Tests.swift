import XCTest
@testable import projecteuler

final class p8Tests: XCTestCase {
    func testP8_greatestProduct5ConsecutiveDigits() {
        XCTAssertEqual(p8_greatestProduct5ConsecutiveDigits(), 40824)
    }
    static var allTests = [
        ("testP8_greatestProduct5ConsecutiveDigits", testP8_greatestProduct5ConsecutiveDigits)
    ]
}
