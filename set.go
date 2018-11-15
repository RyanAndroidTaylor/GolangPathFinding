package main

type StateSet struct {
  set map[string]*State
}

func (s *StateSet) add(state *State) {
  s.set[state.Name] = state
}

func (s *StateSet) remove(state *State) {
  s.set[state.Name] = nil
}

func (s *StateSet) values() map[string]*State {
  return s.set
}

func NewStateSet() *StateSet {
  return &StateSet{make(map[string]*State)}
}
