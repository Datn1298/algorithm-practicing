/* 657. Robot Return to Origin
https://leetcode.com/problems/robot-return-to-origin/

There is a robot starting at position (0, 0), the origin, on a 2D plane. Given a sequence of its moves, judge if this robot ends up at (0, 0) after it completes its moves.
The move sequence is represented by a string, and the character moves[i] represents its ith move. Valid moves are R (right), L (left), U (up), and D (down). If the robot returns to the origin after it finishes all of its moves, return true. Otherwise, return false.
Note: The way that the robot is "facing" is irrelevant. "R" will always make the robot move to the right once, "L" will always make it move left, etc. Also, assume that the magnitude of the robot's movement is the same for each move.
*/

func judgeCircle(moves string) bool {

	// m := make(map[string]int)

	if len(moves)%2 != 0 {
		return false
	}

	var h, k int

	for i := 0; i < len(moves); i++ {

		if moves[i] == 'L' {
			h++
		}

		if moves[i] == 'R' {
			h--
		}

		if moves[i] == 'U' {
			k++
		}

		if moves[i] == 'D' {
			k--
		}

	}

	return k == 0 && h == 0

}
