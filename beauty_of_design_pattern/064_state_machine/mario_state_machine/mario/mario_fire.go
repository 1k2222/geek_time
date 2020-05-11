package mario

type fireMario struct{}

func (c fireMario) getState() state {
	return fire
}

func (c fireMario) obtainMushRoom(*StateMachine) {}

func (c fireMario) obtainCape(*StateMachine) {}

func (c fireMario) obtainFireFlower(*StateMachine) {}

func (c fireMario) meetMonster(stateMachine *StateMachine) {
	stateMachine.iMario = smallMario{}
	stateMachine.score -= 300
}
