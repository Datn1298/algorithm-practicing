/* 70. Climbing Stairs
https://leetcode.com/problems/climbing-stairs/

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top
*/

func climbStairs(n int) int {

	return fibo(n + 1)
}

func fibo(n int) int {
	x, y := 0, 1
	for n > 0 {
		x, y = y, x+y
		n--
	}
	return x
}