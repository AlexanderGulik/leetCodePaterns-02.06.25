func canCompleteCircuit(gas []int, cost []int) int {
  totalGas, start, tank := 0,0,0
  for i := 0; i < len(gas); i++{
    currentgas := gas[i] - cost[i]
    totalGas += currentgas 
    tank += currentgas
    if tank < 0 {
      tank = 0
      start = i+1
    }
  }
  if totalGas < 0 {
    return -1
  }
  return start
}
