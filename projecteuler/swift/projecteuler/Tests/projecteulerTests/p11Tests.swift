import XCTest
@testable import projecteuler

final class p11Tests: XCTestCase {
    func testGreatestProductOfFourAdjacentNums() {
        XCTAssertEqual(greatestProductOfFourAdjacentNums(), 70600674)
    }
    static var allTests = [
        ("testGreatestProductOfFourAdjacentNums", testGreatestProductOfFourAdjacentNums)
    ]
}
