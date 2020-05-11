package mario

type state int

const (
	small state = iota
	super
	fire
	cape
)

var stateName = map[state]string{
	small: "small",
	super: "super",
	fire:  "fire",
	cape:  "cape",
}

type iMario interface {
	getState() state
	obtainMushRoom(stateMachine *StateMachine)
	obtainCape(stateMachine *StateMachine)
	obtainFireFlower(stateMachine *StateMachine)
	meetMonster(MarioStateMachine *StateMachine)
}
