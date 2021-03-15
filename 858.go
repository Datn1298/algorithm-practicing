/* 858. Mirror Reflection
https://leetcode.com/problems/mirror-reflection/
There is a special square room with mirrors on each of the four walls.  Except for the southwest corner, there are receptors on each of the remaining corners, numbered 0, 1, and 2.

The square room has walls of length p, and a laser ray from the southwest corner first meets the east wall at a distance q from the 0th receptor.

Return the number of the receptor that the ray meets first.  (It is guaranteed that the ray will meet a receptor eventually.)
*/

func mirrorReflection(p int, q int) int {

	gcd := GCD(p, q)

	p, q = p/gcd, q/gcd

	lcm := LCM(p, q)

	m := lcm / q
	if p == q {
		return 1
	}

	if m%2 == 0 {
		return 2
	}
	if (lcm/m)%2 == 0 {
		return 0
	} else {
		return 1
	}
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b

	}
	return a
}

func LCM(a, b int) int {
	result := a * b / GCD(a, b)
	return result
}

 