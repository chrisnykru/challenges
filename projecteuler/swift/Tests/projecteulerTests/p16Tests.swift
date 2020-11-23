import XCTest
@testable import projecteuler

final class p16Tests: XCTestCase {
    func testSumOfDigitsBaseExponent() {
        var x = sumOfDigitsBaseExponent(base: 2, exponent: 15)
        XCTAssertEqual(x, 26)
        x = sumOfDigitsBaseExponent(base: 2, exponent: 1000)
        XCTAssertEqual(x, 1366)
    }
    static var allTests = [
        ("testSumOfDigitsBaseExponent", testSumOfDigitsBaseExponent),
    ]
}
