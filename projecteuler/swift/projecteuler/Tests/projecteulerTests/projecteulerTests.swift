import XCTest
@testable import projecteuler

final class projecteulerTests: XCTestCase {
    func testExample() {
        // This is an example of a functional test case.
        // Use XCTAssert and related functions to verify your tests produce the correct
        // results.
        XCTAssertEqual(projecteuler().text, "Hello, World!")
    }

    static var allTests = [
        ("testExample", testExample),
    ]
}
