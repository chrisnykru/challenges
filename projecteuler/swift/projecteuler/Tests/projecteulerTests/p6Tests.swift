import XCTest
@testable import projecteuler

final class p6Tests: XCTestCase {
    func testSumSquareAndSquareOfSumDifference() {
        XCTAssertEqual(sumSquareAndSquareOfSumDifference(10), 25164150)
    }
    static var allTests = [
        ("testSumSquareAndSquareOfSumDifference", testSumSquareAndSquareOfSumDifference)
    ]
}
