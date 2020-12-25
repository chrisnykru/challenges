import XCTest
@testable import projecteuler

final class p26Tests: XCTestCase {
    func testCycle() {
        XCTAssertEqual(cycle(2), 0)
        XCTAssertEqual(cycle(3), 1)
        XCTAssertEqual(cycle(4), 0)
        XCTAssertEqual(cycle(6), 1)
        XCTAssertEqual(cycle(7), 6)
        XCTAssertEqual(cycle(13), 6)
        XCTAssertEqual(cycle(990), 2)
    }
    
    func testFindLongestCycle() {
        let (d, d_cycleLen) = findLongestCycle(1000)
        XCTAssertEqual(d, 983)
        XCTAssertEqual(d_cycleLen, 981)
    }
    
    static var allTests = [
        ("testCycle", testCycle),
        ("testFindLongestCycle", testFindLongestCycle)
    ]
}
