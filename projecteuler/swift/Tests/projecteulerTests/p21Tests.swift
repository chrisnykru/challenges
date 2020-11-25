import XCTest
@testable import projecteuler

final class p21Tests: XCTestCase {
    func test_d_func() {
        do {
            var x = try d(284)
            XCTAssertEqual(x, 220)
            x = try d(220)
            XCTAssertEqual(x, 284)
        } catch {
            XCTFail()
        }
    }
    func testSumOfAmicableNumbers() {
        do {
            let x = try sumOfAmicableNumbers(under: 10000)
            XCTAssertEqual(x, 31626)
        } catch {
            XCTFail()
        }
        
    }
    static var allTests = [
        ("test_d_func", test_d_func),
        ("testSumOfAmicableNumbers", testSumOfAmicableNumbers),
    ]
}
