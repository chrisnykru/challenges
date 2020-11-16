import XCTest
@testable import projecteuler

final class p15Tests: XCTestCase {
    func testP15() {
        var x = numRoutes(squareSide: 2)
        XCTAssertEqual(x, 6)
        x = numRoutes(squareSide: 20)
        XCTAssertEqual(x, 137846528820)
    }
    static var allTests = [
        ("p15", testP15),
    ]
}
