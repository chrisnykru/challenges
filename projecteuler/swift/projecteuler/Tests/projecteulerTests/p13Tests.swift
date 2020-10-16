import XCTest
@testable import projecteuler

final class p13Tests: XCTestCase {
    func testP13() {
        XCTAssertEqual(p13(), 5537376230)
    }
    static var allTests = [
        ("p13", p13),
    ]
}
