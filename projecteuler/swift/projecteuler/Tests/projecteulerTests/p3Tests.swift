import XCTest
@testable import projecteuler

final class p3Tests: XCTestCase {
    func testGetFactors() {
        //let lpf = largestPrimeFactor(600851475143)
        //XCTAssertEquals(lpf, 6857)
        let x = getFactors(64).sorted()
        XCTAssertEqual(x, [1,2,4,8,16,32,64])
    }
    static var allTests = [
        ("testGetFactors", testGetFactors)
    ]
}
