import XCTest
@testable import projecteuler

final class p2Tests: XCTestCase {
    func testSumFibonacciUpTo() {
        let x = sumFibonacciUpTo(4000000)
        XCTAssertEqual(x, 4613732)
        
    }
    static var allTests = [
        ("testSumFibonnaciUpTo", testSumFibonacciUpTo),
    ]
}
