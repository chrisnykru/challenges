import XCTest
@testable import projecteuler

final class p7Tests: XCTestCase {
    func testGetPrime() {
        do {
            var x = try getPrime(nth: 6)
            XCTAssertEqual(x, 13)
            x = try getPrime(nth: 10001)
            XCTAssertEqual(x, 104743)
        } catch {
            XCTFail()
        }
    }
    static var allTests = [
        ("testGetPrime", testGetPrime)
    ]
}
