import XCTest
@testable import projecteuler

final class p15Tests: XCTestCase {
    func testP15() {
        var x = numRoutes(squareSide: 2)
        XCTAssertEqual(x, 6)
        x = numRoutes(squareSide: 20)
        XCTAssertEqual(x, 137846528820)
    }
    static var allTests = [
        ("p15", testP15),
    ]
}

/*

func Test2(t *testing.T) {
        x := numRoutes(2)
        expected_x := int64(6)
        if x != expected_x {
                t.Errorf("x = %v, want %v", x, expected_x)
        }
}

func Test20(t *testing.T) {
        x := numRoutes(20)
        expected_x := int64(137846528820)
        if x != expected_x {
                t.Errorf("x = %v, want %v", x, expected_x)
        }
}
../../projecteuler/go/15/15_test.go (END)

*/
