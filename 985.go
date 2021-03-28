/* 985. Sum of Even Numbers After Queries
https://leetcode.com/problems/sum-of-even-numbers-after-queries/

We have an array A of integers, and an array queries of queries.

For the i-th query val = queries[i][0], index = queries[i][1], we add val to A[index].  Then, the answer to the i-th query is the sum of the even values of A.

(Here, the given index = queries[i][1] is a 0-based index, and each query permanently modifies the array A.)

Return the answer to all queries.  Your answer array should have answer[i] as the answer to the i-th query.

*/

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		res[i] = 0
		A[queries[i][1]] += queries[i][0]
		for j := 0; j < len(A); j++ {
			if A[j]%2 == 0 {
				res[i] += A[j]
			}
		}
	}

	return res
}