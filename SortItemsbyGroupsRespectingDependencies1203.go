func topologicalSort(graph map[int][]int, indegree []int)[]int {
    visited, stack := []int{}, []int{}
    for key := range graph {
        if indegree[key] == 0 {
            stack = append(stack, key)
        }
    }
    for len(stack) > 0 {
        current := stack[len(stack)-1]; stack = stack[:len(stack)-1]
        visited = append(visited, current)
        for _, prev := range graph[current] {
            indegree[prev]--
            if indegree[prev] == 0 {
                stack = append(stack,prev)
            }
        }
    }
    if len(visited) == len(graph) {
        return visited
    }
    return []int{}
}


func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
    groupID := m
    for i := 0; i < n; i++{
        if group[i] == -1 {
            group[i] = groupID
            groupID++
        }
    }
    itemGraph := make(map[int][]int)
    itemIndegree := make([]int, n)
    for i := 0; i < n; i++ {
        itemGraph[i] = []int{}
    }
    groupGraph := make(map[int][]int)
    groupIndegree := make([]int, groupID)
    for i := 0; i< groupID; i++{
        groupGraph[i] = []int{}
    }
    for current := 0; current < n; current++{
        for _, prev := range beforeItems[current] {
            itemGraph[prev] = append(itemGraph[prev], current)
            itemIndegree[current]++

            if group[current] != group[prev] {
                groupGraph[group[prev]] = append(groupGraph[group[prev]], group[current])
                groupIndegree[group[current]]++
            }
        }
    }
    itemOrder := topologicalSort(itemGraph, itemIndegree)
    groupOrder := topologicalSort(groupGraph, groupIndegree)
    if len(itemOrder) <= 0 || len(groupOrder) <= 0 {
        return []int{}
    }
    orderedGroups := make(map[int][]int)
    for _, item := range itemOrder {
        orderedGroups[group[item]] = append(orderedGroups[group[item]], item)
    }
    answerList := []int{}
    for _, groupIndex := range groupOrder {
        answerList = append(answerList, orderedGroups[groupIndex]...)
    }
    return answerList
}
