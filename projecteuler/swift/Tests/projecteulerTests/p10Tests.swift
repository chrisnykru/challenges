import XCTest
@testable import projecteuler

final class p10Tests: XCTestCase {
    func testFindSumOfPrimesBelow2Mill() {
        XCTAssertEqual(findSumOfPrimesBelow2Mill(), 142913828922)
    }
    static var allTests = [
        ("testFindSumOfPrimesBelow2Mill", testFindSumOfPrimesBelow2Mill)
    ]
}
