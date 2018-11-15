package main

import "fmt"
import "strings"

type StatePriorityQueue struct {
  queue []*StateItem
}

func NewStatePriorityQueue() *StatePriorityQueue {
  return &StatePriorityQueue{[]*StateItem{nil}}
}

func (s *StatePriorityQueue) insert(state *State, priority int) {
  stateItem  := &StateItem{state, priority}

  s.queue = append(s.queue, stateItem)

  if len(s.queue) <= 2 {
    return
  }

  var currentLocation = len(s.queue) - 1
  var parentLocation = currentLocation / 2

  for currentLocation > 1 {
    if s.queue[currentLocation].Priority > s.queue[parentLocation].Priority {
      currentItem := s.queue[currentLocation]

      s.queue[currentLocation] = s.queue[parentLocation]
      s.queue[parentLocation] = currentItem

      currentLocation = parentLocation
      parentLocation = currentLocation / 2
    } else {
      break
    }
  }
}

//TODO Need to handle when there is only 1 or 2 items in the queue
func (s *StatePriorityQueue) RemoveMax() *State {
  removedItem := s.queue[1]

  lastItemIndex := len(s.queue) - 1

  // Move the last item in the array to the start of the array
  s.queue[1] = s.queue[lastItemIndex]

  //Remove last item from the end of the queue
  s.queue = append(s.queue[:0], s.queue[:lastItemIndex]...)

  // We moved the last item in the array to position 1
  var lastItemLocation = 1
  var priorityItemLocation = 2

  // Checking to see which index is next in priority
  if s.queue[2].Priority < s.queue[3].Priority {
    priorityItemLocation = 3
  }

  priorityItem := s.queue[priorityItemLocation]

  s.queue[priorityItemLocation] = s.queue[lastItemLocation]
  s.queue[lastItemLocation] = priorityItem

  for {
    lastItemLocation = priorityItemLocation
    priorityItemLocation = lastItemLocation * 2

    if priorityItemLocation >= lastItemIndex {
      break
    }

    if (priorityItemLocation == len(s.queue) -1) {
      // Do nothing
    } else if s.queue[priorityItemLocation].Priority < s.queue[priorityItemLocation + 1].Priority {
      priorityItemLocation = priorityItemLocation + 1
    }

    if s.queue[priorityItemLocation].Priority > s.queue[lastItemLocation].Priority {
      priorityItem = s.queue[priorityItemLocation]

      s.queue[priorityItemLocation] = s.queue[lastItemLocation]
      s.queue[lastItemLocation] = priorityItem
    } else {
      break
    }
  }

  return removedItem.State
}

func (s *StatePriorityQueue) empty() bool {
  return len(s.queue) < 1
}

func (s *StatePriorityQueue) String() string {
  var text = "["

  for _, stateItem := range s.queue {
    if stateItem != nil {
      text += fmt.Sprintf("(State %s, Priority %d), ", stateItem.State.Name, stateItem.Priority)
    }
  }

  return strings.Trim(text, ", ") + "]"
}

type StateItem struct {
  State *State
  Priority int
}
