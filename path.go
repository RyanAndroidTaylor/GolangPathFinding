package main

import "strconv"

type Path struct {
  State *State
  Action *Action
  Cost int
  Parent *Path
}

func (p *Path) Copy() *Path {
  var parent *Path = nil

  if (p.Parent != nil) {
    parent = p.Parent.Copy()
  }

  return &Path{p.State, p.Action, p.Cost, parent}
}

func (p *Path) String() string {
  if (p.Parent != nil) {
    var cost = 0

    if p.Action != nil {
      cost = p.Action.Cost
    }

    return p.Parent.String() + " -> " + strconv.Itoa(cost) + " -> " + p.State.Name
  } else {
    return p.State.Name
  }
}
