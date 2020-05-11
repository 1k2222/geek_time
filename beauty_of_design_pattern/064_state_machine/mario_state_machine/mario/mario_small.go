package mario

type smallMario struct{}

func (s smallMario) getState() state {
	return small
}

func (s smallMario) obtainMushRoom(stateMachine *StateMachine) {
	stateMachine.iMario = superMario{}
	stateMachine.score += 100
}

func (s smallMario) obtainCape(stateMachine *StateMachine) {
	stateMachine.iMario = capeMario{}
	stateMachine.score += 200
}

func (s smallMario) obtainFireFlower(stateMachine *StateMachine) {
	stateMachine.iMario = fireMario{}
	stateMachine.score += 300
}

func (s smallMario) meetMonster(*StateMachine) {}
