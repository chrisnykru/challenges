import XCTest
@testable import projecteuler

final class p19Tests: XCTestCase {
    func testSundaysOnFirstOfMonth() {
        do {
            let x = try sundaysOnFirstOfMonth()
            XCTAssertEqual(x, 171)
        } catch {
            XCTFail()
        }
    }
    static var allTests = [
        ("testSundaysOnFirstOfMonth", testSundaysOnFirstOfMonth),
    ]
}
