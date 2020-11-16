import XCTest
@testable import projecteuler

final class p14Tests: XCTestCase {
    func testP14() {
        let (num, seqLen) = findLargestSeq(1000000)
        XCTAssertEqual(num, 837799)
        XCTAssertEqual(seqLen, 525)
    }
    static var allTests = [
        ("p14", testP14),
    ]
}
