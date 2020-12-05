import XCTest
@testable import projecteuler

final class p23Tests: XCTestCase {
    func testSumAllPosIntNotWriteableAsSumOfTwoAbundantNum() {
        do {
            let x = try sumAllPosIntNotWriteableAsSumOfTwoAbundantNum()
            XCTAssertEqual(x, 4179871)
        } catch {
            XCTFail()
        }
    }
    static var allTests = [
        ("testSumAllPosIntNotWriteableAsSumOfTwoAbundantNum", testSumAllPosIntNotWriteableAsSumOfTwoAbundantNum),
    ]
}
