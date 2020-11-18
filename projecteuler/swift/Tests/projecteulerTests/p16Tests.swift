import XCTest
@testable import projecteuler

final class p16Tests: XCTestCase {
    func testP16() {
        var x = sumOfDigits(base: 2, exponent: 15)
        XCTAssertEqual(x, 26)
        x = sumOfDigits(base: 2, exponent: 1000)
        XCTAssertEqual(x, 1366)
    }
    static var allTests = [
        ("p16", testP16),
    ]
}
