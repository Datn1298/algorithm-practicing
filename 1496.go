/*
1496. Path Crossing https://leetcode.com/problems/path-crossing/
	Given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing moving one unit north, south, east, or west, respectively.
	You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.
	Return True if the path crosses itself at any point, that is, if at any time you are on a location you've previously visited. Return False otherwise.
*/

type Vertex struct {
	x int
	y int
}

func isPathCrossing(path string) bool {
	if len(path) == 1 {
		return false
	}

	//     for i:=1; i< len(path);i++{
	//         if(string(path[i])=="N" && string(path[i-1])=="S"){
	//             return true
	//         }
	//         if(string(path[i])=="S" && string(path[i-1])=="N"){
	//             return true
	//         }
	//         if(string(path[i])=="W" && string(path[i-1])=="E"){
	//             return true
	//         }
	//         if(string(path[i])=="E" && string(path[i-1])=="W"){
	//             return true
	//         }

	//     }

	//     sum := 0;

	//     for i:=0; i<len(path);i++{
	//         if(string(path[i])=="N"){
	//             sum += 1;
	//         } else if(string(path[i])=="S"){
	//             sum -= 1;
	//         } else if(string(path[i])=="E"){
	//             sum += 10000;
	//         } else if(string(path[i])=="W"){
	//             sum -= 10000;
	//         }

	//         if(sum==0) {
	//             return true
	//         }
	//     }

	//     return false

	m := make(map[Vertex]int)

	root := Vertex{0, 0}

	m[root] = 1

	var preVer Vertex
	preVer = root

	for i := 0; i < len(path); i++ {
		if string(path[i]) == "N" {
			preVer.y++
		} else if string(path[i]) == "S" {
			preVer.y--
		} else if string(path[i]) == "E" {
			preVer.x++
		} else if string(path[i]) == "W" {
			preVer.x--
		}
		if n, ok := m[preVer]; ok {
			m[preVer] = n + 1
		} else {
			m[preVer] = 1
		}
	}

	for _, value := range m {
		if value >= 2 {
			return true
		}
	}
	return false
}


