package main

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readMap() WorldMap {
  data, err := ioutil.ReadFile("map.json")

  check(err)

  var mapData MapData

  dataBytes := []byte(data)

  unmarshalError := json.Unmarshal(dataBytes, &mapData)

  check(unmarshalError)

  states := make(map[string]*State)
  var startingPoint *State

  for _, value := range mapData.States {
    state := &State{value.Name}

    if (startingPoint == nil) {
      startingPoint = state
    }

    states[value.Name] = state
  }

  return WorldMap{startingPoint, states, mapData.Actions}
}

type WorldMap struct {
  StartingPoint *State
  States map[string]*State
  Actions []ActionData
}

func (m *WorldMap) GetActionsFor(stateName string) []*Action {
  var actionsForState []*Action

  for _, value := range m.Actions {
    if value.From == stateName {
      action := &Action{m.States[value.From], m.States[value.To], value.Cost}
      actionsForState = append(actionsForState, action)
    }
  }

  return actionsForState
}

type State struct {
  Name string
}

func (s *State) String() string {
  return s.Name
}

type Action struct {
  From *State
  To *State
  Cost int
}

func (a *Action) String() string {
  return fmt.Sprintf("from: %s, to: %s, cost %d", a.From.String(), a.To.String(), a.Cost)
}

type MapData struct {
  States []StateData `json:"states"`
  Actions []ActionData `json:"actions"`
}

type StateData struct {
  Name string `json:"name"`
}

type ActionData struct {
  From string `json:"current_state"`
  To string `json:"next_state"`
  Cost int `json:"cost"`
}
