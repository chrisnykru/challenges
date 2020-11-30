import XCTest
@testable import projecteuler

final class p22Tests: XCTestCase {
    func testNameScore() {
        let x = nameScore("COLIN")
        XCTAssertEqual(x, 53)
    }
    static var allTests = [
        ("testNameScore", testNameScore),
    ]
}
