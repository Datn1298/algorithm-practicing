/*650. 2 Keys Keyboard
https://leetcode.com/problems/2-keys-keyboard/
Initially on a notepad only one character 'A' is present. You can perform two operations on this notepad for each step:

Copy All: You can copy all the characters present on the notepad (partial copy is not allowed).
Paste: You can paste the characters which are copied last time.


Given a number n. You have to get exactly n 'A' on the notepad by performing the minimum number of steps permitted. Output the minimum number of steps to get n 'A'.
*/

func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	if checkPrime(n) == true {
		return n
	} else {
		return findSteps(n)
	}
	return 1
}

func checkPrime(numb int) bool {
	for i := 2; i <= int(math.Sqrt(float64(numb))); i++ {
		if numb%i == 0 {
			return false
		}
	}

	return true
}

func findSL(numb int) int {
	for i := 2; i < numb; i++ {
		if numb%i == 0 {
			return i
			break
		}
	}

	return numb
}

func findSteps(numb int) int {

	if numb == 1 {
		return 0
	}

	lb := findSL(numb)

	return findSteps(numb/lb) + lb

}