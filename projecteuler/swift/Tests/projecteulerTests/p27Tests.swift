import XCTest
@testable import projecteuler

final class p27Tests: XCTestCase {
    func testOptimalAB() {
        let (a, b, primes) = optimalAB()
        XCTAssertEqual(a, -61)
        XCTAssertEqual(b, 971)
        XCTAssertEqual(primes.count, 71)
    }
    
    static var allTests = [
        ("testOptimalAB", testOptimalAB),
    ]
}
