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

  priorityQueue := NewStatePriorityQueue()

  priorityQueue.insert(worldMap.States["Arad"], 1)
  priorityQueue.insert(worldMap.States["Zerind"], 2)
  priorityQueue.insert(worldMap.States["Oradea"], 3)
  priorityQueue.insert(worldMap.States["Timisoara"], 4)
  priorityQueue.insert(worldMap.States["Lugoj"], 5)
  priorityQueue.insert(worldMap.States["Sibiu"], 6)
  priorityQueue.insert(worldMap.States["Urziceni"], 15)

  fmt.Println(priorityQueue)

  removedState := priorityQueue.RemoveMax()

  fmt.Println(priorityQueue)
  fmt.Printf("RemovedState %s\n", removedState.Name)
}

type Explorer struct {
  explored *StateSet
}

func (e *Explorer) Explore(worldMap WorldMap) {
  explored := make(map[string]*State)
  queue := []*State{worldMap.StartingPoint}

  var queueLength = len(queue)

  for queueLength > 0 {
    item := queue[queueLength-1]

    // Remove item from end of queue
    queue = append(queue[:0], queue[:queueLength-1]...)

    actions := worldMap.GetActionsFor(item.Name)

    explored[item.Name] = item

    for _, action := range actions {
      var actionInQueue = false

      for _, state := range queue {
          if (state == action.To) {
            actionInQueue = true
          }
      }

      if (explored[action.To.Name] == nil && !actionInQueue) {
        queue = append(queue, action.To)
      }
    }

    queueLength = len(queue)

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
