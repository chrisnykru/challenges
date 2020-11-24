/*

Maximum path sum II

By starting at the top of the triangle below and moving to adjacent numbers
on the row below, the maximum total from top to bottom is 23.

   3
  7 4
 2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and
'Save Link/Target As...'), a 15K text file containing a triangle with
one-hundred rows.

NOTE: This is a much more difficult version of Problem 18. It is not possible
to try every route to solve this problem, as there are 299 altogether! If you
could check one trillion (1012) routes every second it would take over twenty
billion years to check them all. There is an efficient algorithm to solve it.
;o)

*/

/*
func Test(t *testing.T) {
    sum := tri67()
    if sum != 7273 {
        t.Errorf("sum = %v, want %v", sum, 7273)
    }
}
*/

import Foundation

func tri67() throws -> Int {
    //let stringPath = Bundle.main.path(forResource: "input", ofType: "txt")
    //let urlPath = Bundle.main.url(forResource: "input", withExtension: "txt")
    
    // Get URL of config.json
    let path = Bundle.module.path(forResource: "triangle", ofType: "txt")!
    print(path)
    
    let contents = try String(contentsOfFile: path)
    //print(contents)
    let tri = try parseTriangle(contents)
    return mergeUpAndReduce(tri)
}
