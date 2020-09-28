import XCTest
@testable import projecteuler

final class p9Tests: XCTestCase {
    func testpythagoreanTripletSumEquals1000() {
        XCTAssertEqual(pythagoreanTripletSumEquals1000(), 31875000)
    }
    static var allTests = [
        ("pythagoreanTripletSumEquals1000", pythagoreanTripletSumEquals1000)
    ]
}
