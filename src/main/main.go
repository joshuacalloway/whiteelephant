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

func removeParticipant(lookingfor string, participants []string) ([]string) {
	ret := []string{}

	for _, person := range participants {
		if string(person) != lookingfor {
		  ret = append(ret, person)
		}
	}
	return ret
}

func generateAssignments(participants ...string) ([]assignments) {
  ret := []assignments{}
  var hat []string = participants

  for _, person := range participants {
	hat = removeParticipant(string(person), hat)
	var chosen string
	chosen, hat = pickfromHat(hat)
    ret = append(ret, assignments{person, chosen})
	hat = append(hat, person)
  }
  return ret
}

func main() {
  fmt.Println(generateAssignments("joshua", "alicia", "tom", "tony", "sarah", "angie"))
}
