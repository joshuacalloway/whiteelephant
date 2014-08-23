package main

import (
	"fmt"
	"math/rand"
)

type assignments struct {
	whiteelephant string
	chosen string
}

func (a assignments) String() string {
	return fmt.Sprintf("{white elephant: %s, chosen: %s}", a.whiteelephant, a.chosen)
}

// Pick a chosen from hat.  Return that chosen and a list
// of participants without the picked chosen
func pickfromHat(participants []string) (string, []string) {
	var random = rand.Intn(len(participants))
	var chosen = participants[random]
	return chosen,append(participants[0:random], participants[random+1:len(participants)]...)
}

// return tuple.  participants without lookingfor.
// and return true if lookingfor was removed.  else false if not found
func removeParticipant(lookingfor string, participants []string) (bool, []string) {
	ret := []string{}
    found := false
	for _, person := range participants {
		if string(person) != lookingfor {
		  ret = append(ret, person)
		  found = true
		}
	}
	return found, ret
}

func exists(item string, list []string) bool {
	for _, b := range list {
		if b == item {
			return true
		}
	}
	return false
}

func generateAssignments(participants ...string) ([]assignments) {
  ret := []assignments{}
  var hat []string = participants
  chosenlist := []string{}

  for _, person := range participants {
	_, hat = removeParticipant(string(person), hat)
	var chosen string
	chosen, hat = pickfromHat(hat)
	chosenlist = append(chosenlist, chosen)
    ret = append(ret, assignments{person, chosen})
	if !exists(string(person), chosenlist) {
	  hat = append(hat, person)
	}
  }
  return ret
}

func main() {
  fmt.Println(generateAssignments("joshua", "alicia", "tom", "tony", "sarah", "angie"))
}
