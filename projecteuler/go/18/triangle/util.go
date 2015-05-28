package triangle

func MergeUpAndReduce(tri [][]int) (largestSum int) {
	for len(tri) > 1 {
		lastRow := tri[len(tri)-1]

		for i := 0; i < len(lastRow)-1; i++ {
			// merge up
			if lastRow[i] >= lastRow[i+1] {
				tri[len(tri)-2][i] += lastRow[i]
			} else {
				tri[len(tri)-2][i] += lastRow[i+1]
			}
		}

		// reduce
		tri = tri[0 : len(tri)-1]
	}
	return tri[0][0]
}
