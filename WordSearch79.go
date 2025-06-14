func exit(board [][]byte, word string) bool {
  for i := 0; i < len(board); i++{
    for j := 0; j < len(board[0]); j++{
      if backstep(board, word, i, j, 0) {
        return true
      }
    }
  }
  return false
}

func backstep(board [][]byte, word string, i, j, index int) bool {
  if len(word) == index {
    return true
  }

  if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || word[index] != board[i][j]{
    return false
  }
    temp := board[i][j]
    board[i][j] = '#'
  found := backstep(board, word, i + 1, j, index + 1) || 
  backstep(board, word, i - 1, j, index + 1) ||
  backstep(board, word, i, j + 1, index + 1) ||
  backstep(board, word, i, j - 1, index +1)

  board[i][j] = temp
  return found
}
