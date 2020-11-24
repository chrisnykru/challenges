import XCTest
@testable import projecteuler

final class p67Tests: XCTestCase {
    func testParseAndMergeMassive() {
        do {
            let x = try tri67()
            XCTAssertEqual(x, 7273)
        } catch {
            XCTFail()
        }
    }
    static var allTests = [
        ("testParseAndMergeMassive", testParseAndMergeMassive),
    ]
}
