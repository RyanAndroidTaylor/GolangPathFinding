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

  var step = 10

  startingPath := &Path{worldMap.StartingPoint, nil, 0, nil}

  queue.Insert(startingPath, step)

  var queueLength = queue.Len()

  for queueLength > 0 {
    pathItem := queue.RemoveMax()
    currentPath := pathItem.Path

    step = pathItem.Priority - 1

    actions := worldMap.GetActionsFor(currentPath.State.Name)

    explored[currentPath.State.Name] = currentPath

    for _, action := range actions {
      if (explored[action.To.Name] == nil && !queue.ContainsPathWithState(action.To)) {
        newPath := &Path{action.To, action, action.Cost, currentPath}

        queue.Insert(newPath, step)
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
