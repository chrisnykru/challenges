/*

Even Fibonacci numbers

Each new term in the Fibonacci sequence is generated by adding the previous two
terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.

*/

func sumFibonacciUpTo(_ doNotExceed: Int) -> Int {
    /*
     
     1,1,2,3,5,8,13,...
     */
    
    var x = 1
    var y = 1
    var sum = 0 // first two terms are odd
    
    // x=2,y=1,sum=4
    // x=3,y=2,sum=7
    // x=5,y=3,sum=12
    while x <= doNotExceed {
        let tmp = x
        x = x + y
        y = tmp
        if x % 2 == 0 {
            sum += x
        }
    }
    return sum
}
