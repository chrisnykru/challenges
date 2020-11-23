import XCTest
@testable import projecteuler

final class p20Tests: XCTestCase {
    func testSumOfDigitsOfFactorial() {
        var x = sumOfDigitsOfFactorial(10)
        XCTAssertEqual(x, 27)
        x = sumOfDigitsOfFactorial(100)
        XCTAssertEqual(x, 648)
    }
    static var allTests = [
        ("testSumOfDigitsOfFactorial", testSumOfDigitsOfFactorial),
    ]
}
