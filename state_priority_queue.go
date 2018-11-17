package main

import "fmt"
import "strings"

type PathPriorityQueue struct {
  queue []*PathItem
  set map[string]*Path
}

func NewPathPriorityQueue() *PathPriorityQueue {
  return &PathPriorityQueue{[]*PathItem{nil}, make(map[string]*Path)}
}

func (p *PathPriorityQueue) Insert(path *Path, priority int) {
  pathItem := &PathItem{path, priority}

  p.queue = append(p.queue, pathItem)
  p.set[path.State.Name] = path

  if len(p.queue) <= 2 {
    return
  }

  var currentLocation = len(p.queue) - 1
  var parentLocation = currentLocation / 2

  for currentLocation > 1 {
    if p.queue[currentLocation].Priority > p.queue[parentLocation].Priority {
      currentItem := p.queue[currentLocation]

      p.queue[currentLocation] = p.queue[parentLocation]
      p.queue[parentLocation] = currentItem

      currentLocation = parentLocation
      parentLocation = currentLocation / 2
    } else {
      break
    }
  }
}

//TODO Need to handle when there is only 1 or 2 items in the queue
func (p *PathPriorityQueue) RemoveMax() *PathItem {
  if len(p.queue) <= 1 {
    return nil
  }

  removedItem := p.queue[1]

  lastItemIndex := len(p.queue) - 1

  // Move the last item in the array to the start of the array
  p.queue[1] = p.queue[lastItemIndex]
  p.set[removedItem.Path.State.Name] = nil

  //Remove last item from the end of the queue
  p.queue = append(p.queue[:0], p.queue[:lastItemIndex]...)

  // We moved the last item in the array to position 1
  var lastItemLocation = 1
  var priorityItemLocation = 2

  queueLength := len(p.queue)

  // Checking to see which index is next in priority
  if queueLength <= 2 {
    return removedItem
  } else if queueLength == 3 {
    priorityItemLocation = 2
  } else if p.queue[2].Priority < p.queue[3].Priority {
    priorityItemLocation = 3
  }

  priorityItem := p.queue[priorityItemLocation]

  p.queue[priorityItemLocation] = p.queue[lastItemLocation]
  p.queue[lastItemLocation] = priorityItem

  for {
    lastItemLocation = priorityItemLocation
    priorityItemLocation = lastItemLocation * 2

    if priorityItemLocation >= lastItemIndex {
      break
    }

    if (priorityItemLocation == len(p.queue) -1) {
      // Do nothing
    } else if p.queue[priorityItemLocation].Priority < p.queue[priorityItemLocation + 1].Priority {
      priorityItemLocation = priorityItemLocation + 1
    }

    if p.queue[priorityItemLocation].Priority > p.queue[lastItemLocation].Priority {
      priorityItem = p.queue[priorityItemLocation]

      p.queue[priorityItemLocation] = p.queue[lastItemLocation]
      p.queue[lastItemLocation] = priorityItem
    } else {
      break
    }
  }

  return removedItem
}

func (p *PathPriorityQueue) Len() int {
  return len(p.queue) - 1
}

func (p *PathPriorityQueue) Empty() bool {
  return len(p.queue) < 1
}

func (p *PathPriorityQueue) ContainsPathWithState(state *State) bool {
  return p.set[state.Name] != nil
}

func (p *PathPriorityQueue) String() string {
  var text = "["

  for _, pathItem := range p.queue {
    if pathItem != nil {
      text += fmt.Sprintf("(State %s, Priority %d), ", pathItem.Path.State.Name, pathItem.Priority)
    }
  }

  return strings.Trim(text, ", ") + "]"
}

type PathItem struct {
  Path *Path
  Priority int
}
