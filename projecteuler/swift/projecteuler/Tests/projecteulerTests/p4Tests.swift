import XCTest
@testable import projecteuler

final class p4Tests: XCTestCase {
    func testIsPalindrome() {
        XCTAssertEqual(isPalindrome(x: 422), false)
        XCTAssertEqual(isPalindrome(x: 424), true)
        XCTAssertEqual(isPalindrome(x: 42), false)
        XCTAssertEqual(isPalindrome(x: 1221), true)
        XCTAssertEqual(isPalindrome(x: 1222), false)
        
    }
    static var allTests = [
        ("testIsPalindrome", testIsPalindrome)
    ]
}
