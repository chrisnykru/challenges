import XCTest
@testable import projecteuler

final class p10Tests: XCTestCase {
    func testFindSumOfPrimesBelow2Mill() {
        XCTAssertEqual(findSumOfPrimesBelow2Mill(), 142913828922)
    }
    static var allTests = [
        ("findSumOfPrimesBelow2Mill", findSumOfPrimesBelow2Mill)
    ]
}
