import XCTest
@testable import projecteuler

final class p24Tests: XCTestCase {
    func testPermGenHelper(_ expected: [[Int]]) throws {
        // extract 'set' from expected output
        // as it's the first result
        var x = try PermGen(expected[0])
        
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
            try testPermGenHelper(expected)
            
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
            try testPermGenHelper(expected)
            
            expected = [
                [5, 6, 6, 6],
                [6, 5, 6, 6],
                [6, 6, 5, 6],
                [6, 6, 6, 5],
            ]
            try testPermGenHelper(expected)
            
            expected = [
                [5, 5, 6, 6],
                [5, 6, 5, 6],
                [5, 6, 6, 5],
                [6, 5, 5, 6],
                [6, 5, 6, 5],
                [6, 6, 5, 5],
            ]
            try testPermGenHelper(expected)
                
        } catch {
            XCTFail()
        }
    }
    
    func testPerm1e6() {
        do {
            let x = try perm1e6()
            XCTAssertEqual(x, [2, 7, 8, 3, 9, 1, 5, 4, 6, 0])
        } catch {
            XCTFail()
        }
    }
    
    static var allTests = [
        ("testPermGen", testPermGen),
        ("testPerm1e6", testPerm1e6),
    ]
}
