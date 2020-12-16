import XCTest
@testable import projecteuler

final class p24Tests: XCTestCase {
    func testPermGenHelper(_ perm: [Int], _ expected: [[Int]]) throws {
        var x = try PermGen(perm)
        for i in stride(from: 0, to: expected.count, by: 1) {
            let (p, last) = x.next()
            XCTAssertEqual(p, expected[i])
            XCTAssertEqual(last, i == expected.count - 1)
        }
        // hold last permutation
        let (p, last) = x.next()
        XCTAssertEqual(p, expected[expected.count - 1])
        XCTAssertEqual(last, true)
    }
    
    func testPermGen() {
        do {
            var expected = [
                [0,1,2],
                [0,2,1],
                [1,0,2],
                [1,2,0],
                [2,0,1],
                [2,1,0],
            ]
            try testPermGenHelper(expected[0], expected)
            
            expected = [
                [5, 5, 6, 9],
                [5, 5, 9, 6],
                [5, 6, 5, 9],
                [5, 6, 9, 5],
                [5, 9, 5, 6],
                [5, 9, 6, 5],
                [6, 5, 5, 9],
                [6, 5, 9, 5],
                [6, 9, 5, 5],
                [9, 5, 5, 6],
                [9, 5, 6, 5],
                [9, 6, 5, 5],
            ]
            try testPermGenHelper(expected[0], expected)
            
            expected = [
                [5, 6, 6, 6],
                [6, 5, 6, 6],
                [6, 6, 5, 6],
                [6, 6, 6, 5],
            ]
            try testPermGenHelper(expected[0], expected)
            
            expected = [
                [5, 5, 6, 6],
                [5, 6, 5, 6],
                [5, 6, 6, 5],
                [6, 5, 5, 6],
                [6, 5, 6, 5],
                [6, 6, 5, 5],
            ]
            try testPermGenHelper(expected[0], expected)
                
        } catch {
            XCTFail()
        }
    }
    
    /*
    func testPermGen__() {
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
            
            
            //XCTAssertEqual(x, 4179871)
        } catch {
            XCTFail()
        }
    }
    */
    static var allTests = [
        ("testPermGen", testPermGen),
    ]
}
