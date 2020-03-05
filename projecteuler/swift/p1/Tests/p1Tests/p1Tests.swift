import XCTest
@testable import p1

final class p1Tests: XCTestCase {
    func test1000() {
        let x = sumOfMultiples(below: 1000)
        XCTAssertEqual(x, 233168)
    }

    func testSumOfMultiplesPerformance() {
        measure {
            _ = sumOfMultiples(below: 1000)
        }
    }

    static var allTests = [
        ("test1000", test1000),
        ("sumOfMultiplesPerformance", testSumOfMultiplesPerformance)
    ]
}
