import XCTest
@testable import projecteuler

final class p12Tests: XCTestCase {
    func testTriangleNumGen() {
        let g = TriangleNumGen()
        XCTAssertEqual(g.next(), 1)
        XCTAssertEqual(g.next(), 3)
        XCTAssertEqual(g.next(), 6)
        XCTAssertEqual(g.next(), 10)
        XCTAssertEqual(g.next(), 15)
        XCTAssertEqual(g.next(), 21)
        XCTAssertEqual(g.next(), 28)
        XCTAssertEqual(g.next(), 36)
        XCTAssertEqual(g.next(), 45)
        XCTAssertEqual(g.next(), 55)
    }
    static var allTests = [
        ("testTriangleNumGen", testTriangleNumGen)
    ]
}
