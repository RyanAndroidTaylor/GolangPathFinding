package main
// Node
  // State, the current state and the state at the end of the path
  // Action, the action to get to this state
  // Cost, total cost
  // Parent, the parent node

  //States A -> S -> F
  // Current node would be
    // State = F
    // Action = move to F action
    // Cost = the cost of the move (20)
    // Parent = S the node you came from

// State (city)
  // City name
  // List<Action> actions that can be taken from this state

// called when you perform an action
// fun result(currentState, action): newState

// Action (road)
  // OldState (city the road comes from)
  // NewState (city the road leads to)
  // Cost (distance)

// Path (a list of actions)
  // List<Action>
  // fun pathCost(path): Cost // total cost of the Path

// Need to keep track of the explored states and frontier states

type Tree struct {
  Left *Tree
  State *State
  Right *Tree
}
