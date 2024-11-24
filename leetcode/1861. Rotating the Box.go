package leetcode

func RotateTheBox(box [][]byte) [][]byte {
    m, n := len(box), len(box[0])

    // Step 1: Simulate gravity for each row
    for i := 0; i < m; i++ {
        writePointer := n - 1
        for j := n - 1; j >= 0; j-- {
            if box[i][j] == '#' {
                box[i][writePointer] = '#' // Place the stone
                if writePointer != j {
                    box[i][j] = '.' // Clear the original position
                }
                writePointer--
            } else if box[i][j] == '*' {
                writePointer = j - 1 // Reset writePointer above the obstacle
            }
        }
    }

    // Step 2: Rotate the box 90 degrees clockwise
    rotatedBox := make([][]byte, n)
    for i := 0; i < n; i++ {
        rotatedBox[i] = make([]byte, m)
        for j := 0; j < m; j++ {
            rotatedBox[i][j] = box[m-j-1][i]
        }
    }

    return rotatedBox
}
