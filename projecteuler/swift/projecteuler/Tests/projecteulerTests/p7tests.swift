import XCTest
@testable import projecteuler

final class p7Tests: XCTestCase {
    func testGetPrime() {
        XCTAssertEqual(getPrime(6), 13)
        //XXX XCTAssertEqual(getPrime(10001), 104743)
    }
    static var allTests = [
        ("testGetPrime", testGetPrime)
    ]
}
