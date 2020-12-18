import XCTest
@testable import projecteuler

final class p25Tests: XCTestCase {
    func testFibonacciGen() {
        var x = FibonacciGen<Int>()
        XCTAssertEqual(x.next(), 1)
        XCTAssertEqual(x.next(), 1)
        XCTAssertEqual(x.next(), 2)
        XCTAssertEqual(x.next(), 3)
        XCTAssertEqual(x.next(), 5)
        XCTAssertEqual(x.next(), 8)
        XCTAssertEqual(x.next(), 13)
        XCTAssertEqual(x.next(), 21)
        XCTAssertEqual(x.next(), 34)
        XCTAssertEqual(x.next(), 55)
        XCTAssertEqual(x.next(), 89)
        XCTAssertEqual(x.next(), 144)
    }
    
    func testFirstFibTermOfNDigits() {
        let (term, _) = firstFibTermOfNDigits(n: 1000)
        XCTAssertEqual(term, 4782)
    }
    
    static var allTests = [
        ("testFibonacciGen", testFibonacciGen),
        ("testFirstFibTermOfNDigits", testFirstFibTermOfNDigits)
    ]
}
