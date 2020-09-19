import XCTest
@testable import projecteuler

final class p6Tests: XCTestCase {
    func testSumSquareAndSquareOfSumDifference() {
        // test naive
        XCTAssertEqual(sumSquareAndSquareOfSumDifferenceNaive(10), 2640)
        XCTAssertEqual(sumSquareAndSquareOfSumDifferenceNaive(100), 25164150)
        // test optimized
        XCTAssertEqual(sumSquareAndSquareOfSumDifference(10), 2640)
        XCTAssertEqual(sumSquareAndSquareOfSumDifference(100), 25164150)
    }
    static var allTests = [
        ("testSumSquareAndSquareOfSumDifference", testSumSquareAndSquareOfSumDifference)
    ]
}
