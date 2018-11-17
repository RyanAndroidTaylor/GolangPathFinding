package main

import "fmt"

func main() {
  worldMap := readMap()

  aradActions := worldMap.GetActionsFor("Arad")

  fmt.Printf("aradActions length %d\n", len(aradActions))

  for i := 0; i < len(aradActions); i++ {
    fmt.Println(aradActions[i])
  }

  explorer := Explorer{NewStateSet()}

  explorer.Explore(worldMap)

  fmt.Println(explorer)
}

type Explorer struct {
  explored *StateSet
}

func (e *Explorer) Explore(worldMap WorldMap) {
  explored := make(map[string]*Path)
  queue := NewPathPriorityQueue()

  startingPath := &Path{worldMap.StartingPoint, nil, 0, nil}

  queue.Insert(startingPath, 0)

  var queueLength = queue.Len()

  for queueLength > 0 {
    pathItem := queue.RemoveMax()
    currentPath := pathItem.Path

    actions := worldMap.GetActionsFor(currentPath.State.Name)

    explored[currentPath.State.Name] = currentPath

    for _, action := range actions {
      // Use && !queue.ContainsPathWithState(action.To) if you are trying to get to the location in the fewest moves
      if (explored[action.To.Name] == nil) {
        newPath := &Path{action.To, action, action.Cost, currentPath}

        queue.Insert(newPath, pathItem.Priority + action.Cost)
      }
    }

    queueLength = queue.Len()

    fmt.Printf("\nitem %s\n", pathItem.Path.State.Name)
    fmt.Printf("queue length %d\n", queueLength)
    fmt.Printf("queue ", queue)

    if pathItem.Path.State.Name == "Bucharest" {
      fmt.Println("\nBucharest found", pathItem)

      break
    }
  }
}
