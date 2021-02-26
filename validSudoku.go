var myMap = make(map[string]bool, 9)

func isValidSudoku(board [][]byte) bool {
    return checkRows(board) && checkColumns(board) && checkBoxes(board)
}

func clearGlobalMap() {
    myMap = make(map[string]bool, 9)
}

func checkRows(board [][]byte) bool {
    for i := range board {
        clearGlobalMap()
        for j := range board[i] {
            if !verifyMap(i, j, board) {
                return false
            }
        }
    }
    return true
}

func checkColumns(board [][]byte) bool {
    for i := range board {
        clearGlobalMap()
        for j := range board[i] {
            if !verifyMap(j, i, board) {
                return false
            }
        }
    }
    return true
}

func checkBoxes(board [][]byte) bool {
    for h := 0; h < len(board); h+=3 {
        for g := 0; g < len(board); g += 3 {
            clearGlobalMap()
            for i := h; i < h + 3; i++ {
                for j := g; j < g + 3; j++ {
                    if !verifyMap(i, j, board) {
                        return false
                    }
                }
            }  
        }
    }   
    return true
}

func verifyMap(i int, j int, board [][]byte) bool {
    if myMap[string(board[i][j])] {
        return false
    }
    if board[i][j] != '.' {
        myMap[string(board[i][j])] = true 
    }
    return true
}
