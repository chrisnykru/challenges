import XCTest
@testable import projecteuler

final class p4Tests: XCTestCase {
    func testIsPalindrome() {
        XCTAssertEqual(isPalindrome(422), false)
        XCTAssertEqual(isPalindrome(424), true)
        XCTAssertEqual(isPalindrome(42), false)
        XCTAssertEqual(isPalindrome(1221), true)
        XCTAssertEqual(isPalindrome(1222), false)
        
    }
    func testFindLargestPalindrome() {
        XCTAssertEqual(findLargestPalindrome(2), PalindromeProduct(i: 99, j: 91, product: 9009))
        XCTAssertEqual(findLargestPalindrome(3), PalindromeProduct(i: 993, j: 913, product: 906609))
    }
    static var allTests = [
        ("testIsPalindrome", testIsPalindrome),
        ("testFindLargestPalindrome", testFindLargestPalindrome)
    ]
}
