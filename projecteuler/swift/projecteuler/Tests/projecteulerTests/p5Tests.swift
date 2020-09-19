import XCTest
@testable import projecteuler

final class p5Tests: XCTestCase {
    func testSmallestMultipleEvenlyDivisible() {
        XCTAssertEqual(smallestMultipleEvenlyDivisible(10), 2520)
        XCTAssertEqual(smallestMultipleEvenlyDivisible(20), 232792560)
    }
    static var allTests = [
        ("testSmallestMultipleEvenlyDivisible", testSmallestMultipleEvenlyDivisible)
    ]
}
