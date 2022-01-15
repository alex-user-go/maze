# maze

***Breadth First Traversal (BFS)*** on a 2D array.

Given a maze (matrix of size m × n consisting of integers), the taks is to find the **minimum path** out of the maze from starting position (starting from {1,1}). The minimum path is the path that gets to an exit of the maze in the least distance (number of moves). The numbers in the matrix represent objects, some points (objects) are allowed to visit, some are not.

## Example

```go
func main() {
    matrix := [][]int{
                {1, 0, 1, 1}
		{1, 0, 0, 1},
		{1, 1, 2, 1},
		{1, 1, 1, 1},
	},
    //allowed to visit
    openPoints := []int{0,2}
    mz, err := NewMaze(matrix, openPoints)
    if err = nil{
        fmt.Println(err)
        os.Exit()
    }
    //starting position (number 2)
    startColumn := 3
    startRow := 3
    pathLen, err := mz.FindLenExit(startColumn, startRow)
    if err = nil{
        fmt.Println(err)
    }else{
         // 3
        fmt.Println(pathLen)
    }    
}
```
