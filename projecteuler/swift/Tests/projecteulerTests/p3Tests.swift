import XCTest
@testable import projecteuler

final class p3Tests: XCTestCase {
    func testFactors() {
        do {
            let x = try factors(64).sorted()
            XCTAssertEqual(x, [1,2,4,8,16,32,64])
        } catch {
            XCTFail()
        }
    }
    func testFactorsThrow() {
        XCTAssertThrowsError(try factors(-1))
    }
    func testPrimeFactors() {
        do {
            let x = try primeFactors(64).sorted()
            XCTAssertEqual(x, [1,2])
        } catch {
            XCTFail()
        }
    }
    func testLargestPrimeFactor() {
        do {
            let x = try largestPrimeFactor(600851475143)
            XCTAssertEqual(x, 6857)
        } catch {
            XCTFail()
        }
    }
    static var allTests = [
        ("testFactors", testFactors),
        ("testFactorsThrow", testFactorsThrow),
        ("testPrimeFactors", testPrimeFactors),
        ("testLargestPrimeFactor", testLargestPrimeFactor)
    ]
}
