import XCTest
@testable import projecteuler

final class p22Tests: XCTestCase {
    func testTotalNameScore() {
        do {
            let x = try totalNameScore()
            XCTAssertEqual(x, 871198282)
        } catch {
            XCTFail()
        }
    }
    func testNameScore() {
        let x = nameScore("COLIN")
        XCTAssertEqual(x, 53)
    }
    static var allTests = [
        ("testNameScore", testNameScore),
        ("testTotalNameScore", testTotalNameScore)
    ]
}
