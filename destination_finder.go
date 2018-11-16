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
  explored := make(map[string]*State)
  queue := NewStatePriorityQueue()

  var step = 10

  queue.Insert(worldMap.StartingPoint, step)

  var queueLength = queue.Len()

  for queueLength > 0 {
    stateItem := queue.RemoveMax()
    item := stateItem.State

    step = stateItem.Priority - 1

    actions := worldMap.GetActionsFor(item.Name)

    explored[item.Name] = item

    for _, action := range actions {
      var actionInQueue = false

      for _, state := range queue.queue {
          if state != nil && state.State == action.To {
            actionInQueue = true
          }
      }

      if (explored[action.To.Name] == nil && !actionInQueue) {
        queue.Insert(action.To, step)
      }
    }

    queueLength = queue.Len()

    fmt.Printf("\nitem %s\n", item.Name)
    fmt.Printf("queue length %d\n", queueLength)
    fmt.Printf("queue ", queue)
  }
}

type Node struct {
  State *State
  Action *Action
  Parent * Node
}
