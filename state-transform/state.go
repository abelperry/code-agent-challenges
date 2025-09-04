package state_transform

type StateTransform struct {
	creatorID string
	operator  User
	state     DatasetState
	action    DatasetAction
}

func NewStateTransform(creatorID string, operator User, state DatasetState, action DatasetAction) StateChecker {
	return &StateTransform{
		creatorID: creatorID,
		operator:  operator,
		state:     state,
		action:    action,
	}
}

func (s *StateTransform) Check() error {
	return nil
}
