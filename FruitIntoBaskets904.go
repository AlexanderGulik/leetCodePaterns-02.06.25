func totalFruit(fruits []int) int {
  basket := make(map[int]int)
  left, maxFruit := 0, 0
  for right, fruit := range fruits{
    basket[fruit]++
    for len(basket) > 2 {
      leftFruit := fruits[left]
      basket[leftFruit]--
      if basket[leftFruit] == 0 {
          delete(basket, leftFruit)
      }
      left++
    }
    if right - left +1 > maxFruit{
      maxFruit = right - left +1
    }
  }
  return maxFruit
}
