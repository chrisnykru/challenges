import XCTest
@testable import projecteuler

final class p27Tests: XCTestCase {
    func testOptimalAB() {
        do {
            let (a, b, primes) = try optimalAB()
            XCTAssertEqual(a, -61)
            XCTAssertEqual(b, 971)
            XCTAssertEqual(primes.count, 71)
        } catch {
            XCTFail()
        }
    }
    
    static var allTests = [
        ("testOptimalAB", testOptimalAB),
    ]
}
