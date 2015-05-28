package main

/*

most numbers in range 1-googol are bouncy
easier to count not-bouncy numbers

////////

[0,99] are bouncy


not bouncy = ..., 122, 123, 124, ...

1222, 1223, 1224, ...


----

Key critical concept:
For every digit we add, we need to thin the herd on a percentage-basis...
Don't go depth first, it'll take too damn long

Not-bouncy numbers:
[1,10) --> [1,2,3,4,5,6,7,8,9]
  note: 100% not-bouncy

[10,100) --> ...
  note: some percentage < 100% will be not-bouncy

[100, 1000) --> 100, 110, 111, 112, ..., 244, 245, 246, ..., 980, 981, 982, ..., 998, 999
  again, not-bouncy percentage should be LESS than previous power of ten...

[1000, 10000) --> we derive subset of possibilities from previous answers:
  e.g., check intervals: [100*10, 110*10) + [110*10, 111*10] + ...
	now, from check intervals, we derive the new not-bouncy dudes for [1000, 10000) range





*/
