import XCTest
@testable import projecteuler

final class p24Tests: XCTestCase {
    func testPermGen() {
        do {
            var x = try PermGen([0,1,2])
            var (p, last) = x.next()
            XCTAssertEqual(p, [0,1,2])
            XCTAssertEqual(last, false)
            
            (p, last) = x.next()
            XCTAssertEqual(p, [0,2,1])
            XCTAssertEqual(last, false)
            
            (p, last) = x.next()
            XCTAssertEqual(p, [1,0,2])
            XCTAssertEqual(last, false)
            
            (p, last) = x.next()
            XCTAssertEqual(p, [1,2,0])
            XCTAssertEqual(last, false)
            
            (p, last) = x.next()
            XCTAssertEqual(p, [2,0,1])
            XCTAssertEqual(last, false)
            
            (p, last) = x.next()
            XCTAssertEqual(p, [2,1,0])
            XCTAssertEqual(last, true)
            
            // should repeat
            (p, last) = x.next()
            XCTAssertEqual(p, [2,1,0])
            XCTAssertEqual(last, true)
            
            
            /*
             unc TestPermGen_012(t *testing.T) {
                 gen := New([]int{0, 1, 2})
                 expected := [][]int{
                     []int{0, 1, 2},
                     []int{0, 2, 1},
                     []int{1, 0, 2},
                     []int{1, 2, 0},
                     []int{2, 0, 1},
                     []int{2, 1, 0},
                 }
                 testPermGen(t, gen, expected)
             }
             */
            //XCTAssertEqual(x, 4179871)
        } catch {
            XCTFail()
        }
    }
    static var allTests = [
        ("testPermGen", testPermGen),
    ]
}
