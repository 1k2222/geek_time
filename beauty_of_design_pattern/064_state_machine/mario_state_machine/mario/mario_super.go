package mario

type superMario struct{}

func (s superMario) getState() state {
	return super
}

func (s superMario) obtainMushRoom(*StateMachine) {}

func (s superMario) obtainCape(stateMachine *StateMachine) {
	stateMachine.iMario = capeMario{}
	stateMachine.score += 200
}

func (s superMario) obtainFireFlower(stateMachine *StateMachine) {
	stateMachine.iMario = fireMario{}
	stateMachine.score += 300
}

func (s superMario) meetMonster(stateMachine *StateMachine) {
	stateMachine.iMario = smallMario{}
	stateMachine.score -= 100
}
