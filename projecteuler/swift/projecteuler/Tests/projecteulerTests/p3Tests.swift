import XCTest
@testable import projecteuler

final class p3Tests: XCTestCase {
    func testFactors() {
        //let lpf = largestPrimeFactor(600851475143)
        //XCTAssertEquals(lpf, 6857)
        let x = factors(64).sorted()
        XCTAssertEqual(x, [1,2,4,8,16,32,64])
    }
    func testPrimeFactors() {
        let x = primeFactors(64).sorted()
        XCTAssertEqual(x, [1,2])
    }
    func testLargestPrimeFactor() {
        let x = largestPrimeFactor(600851475143)
        XCTAssertEqual(x, 6857)
    }
    static var allTests = [
        ("testFactors", testFactors),
        ("testPrimeFactors", testPrimeFactors),
        ("testLargestPrimeFactor", testLargestPrimeFactor)
    ]
}
