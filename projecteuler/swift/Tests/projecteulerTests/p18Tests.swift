import XCTest
@testable import projecteuler

final class p18Tests: XCTestCase {
    func testParseAndMerge() {
        do {
            var x = try small()
            XCTAssertEqual(x, 23)
            x = try large()
            XCTAssertEqual(x, 1074)
        } catch {
            XCTFail()
        }
    }
    static var allTests = [
        ("testParseAndMerge", testParseAndMerge),
    ]
}
