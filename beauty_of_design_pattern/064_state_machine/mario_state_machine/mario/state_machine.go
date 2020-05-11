package mario

type StateMachine struct {
	iMario
	score int
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		iMario: smallMario{},
	}
}

func (stateMachine *StateMachine) GetState() string {
	state := stateMachine.getState()
	return stateName[state]
}

func (stateMachine *StateMachine) GetScore() int {
	return stateMachine.score
}

func (stateMachine *StateMachine) ObtainMushRoom() *StateMachine {
	stateMachine.obtainMushRoom(stateMachine)
	return stateMachine
}

func (stateMachine *StateMachine) ObtainCape() *StateMachine {
	stateMachine.obtainCape(stateMachine)
	return stateMachine
}

func (stateMachine *StateMachine) ObtainFireFlower() *StateMachine {
	stateMachine.obtainFireFlower(stateMachine)
	return stateMachine
}

func (stateMachine *StateMachine) MeetMonster() *StateMachine {
	stateMachine.meetMonster(stateMachine)
	return stateMachine
}
