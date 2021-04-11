/* 223. Rectangle Area
https://leetcode.com/problems/rectangle-area/

Given the coordinates of two rectilinear rectangles in a 2D plane, return the total area covered by the two rectangles.
The first rectangle is defined by its bottom-left corner (A, B) and its top-right corner (C, D).
The second rectangle is defined by its bottom-left corner (E, F) and its top-right corner (G, H).
*/

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area := (C-A)*(D-B) + (G-E)*(H-F)

	if A > E {
		A, B, C, D, E, F, G, H = E, F, G, H, A, B, C, D
	}

	if C <= E {
		return area
	}

	x := min(C, G) - E

	if B > F {
		A, B, C, D, E, F, G, H = E, F, G, H, A, B, C, D
	}

	if D <= F {
		return area
	}

	y := min(D, H) - F

	return area - x*y
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
