import XCTest
@testable import projecteuler

final class p1Tests: XCTestCase {
    func test1000() {
        let x = sumOfMultiples(below: 1000)
        XCTAssertEqual(x, 233168)
        let x2 = fasterSumOfMultiples(below: 1000)
        XCTAssertEqual(x2, 233168)
    }

    // testing at 1000 does not reveal performance differences, so bump up to 1e6
    func testSumOfMultiplesPerformance() {
        measure {
            _ = sumOfMultiples(below: 1_000_000)
        }
    }

    func testFasterSumOfMultiplesPerformance() {
        measure {
            _ = fasterSumOfMultiples(below: 1_000_000)
        }
    }

    static var allTests = [
        ("test1000", test1000),
        ("sum performance", testSumOfMultiplesPerformance),
        ("faster sum performance", testFasterSumOfMultiplesPerformance)
    ]
}
