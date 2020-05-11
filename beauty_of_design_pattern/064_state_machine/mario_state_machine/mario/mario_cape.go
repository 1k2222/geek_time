package mario

type capeMario struct{}

func (c capeMario) getState() state {
	return cape
}

func (c capeMario) obtainMushRoom(*StateMachine) {}

func (c capeMario) obtainCape(*StateMachine) {}

func (c capeMario) obtainFireFlower(*StateMachine) {}

func (c capeMario) meetMonster(stateMachine *StateMachine) {
	stateMachine.iMario = smallMario{}
	stateMachine.score -= 200
}
