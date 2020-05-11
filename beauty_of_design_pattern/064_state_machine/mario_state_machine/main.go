package main

import (
	"./mario"
	"fmt"
)

var cases = []func() *mario.StateMachine{
	func() *mario.StateMachine {
		return mario.NewStateMachine().ObtainMushRoom().MeetMonster().ObtainFireFlower()
	},
	func() *mario.StateMachine {
		return mario.NewStateMachine().MeetMonster().ObtainCape().ObtainCape().ObtainFireFlower()
	},
	func() *mario.StateMachine {
		return mario.NewStateMachine().ObtainFireFlower().MeetMonster()
	},
}

func main() {
	for caseID, cas := range cases {
		machine := cas()
		fmt.Printf("Case %d, state=%s score=%d\n", caseID, machine.GetState(), machine.GetScore())
	}
}
