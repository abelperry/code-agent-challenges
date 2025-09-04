package state_transform

import (
	"testing"
)

type CheckerTest struct {
	creatorID  string
	operator   User
	stateTrans [5][8]int // 5 states, 8 actions, 0 means no error, 1 means error
}

func TestStateChecker(t *testing.T) {
	creator := User{
		ID:    "creator",
		Roles: []string{"developer"},
	}

	admin := User{
		ID:    "admin",
		Roles: []string{"admin"},
	}

	normalUser := User{
		ID:    "normalUser",
		Roles: []string{"developer"},
	}

	tests := []CheckerTest{
		// creator
		{
			creatorID: creator.ID,
			operator:  creator,
			stateTrans: [5][8]int{
				// dev pub dep ban del modify read exec
				{0, 1, 1, 1, 1, 1, 1, 1}, // invalid
				{0, 0, 0, 0, 0, 0, 0, 0}, // dev
				{0, 0, 0, 0, 0, 0, 0, 0}, // release
				{0, 1, 0, 0, 0, 0, 0, 0}, // outdate
				{1, 1, 1, 0, 0, 1, 0, 1}, // ban
			},
		},

		// admin
		{
			creatorID: creator.ID,
			operator:  admin,
			stateTrans: [5][8]int{
				// prv pre stab ban del modify read exec
				{0, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 1, 0, 1},
			},
		},

		// normal user
		{
			creatorID: creator.ID,
			operator:  normalUser,
			stateTrans: [5][8]int{
				// prv pre stab ban modify del read exec
				{0, 1, 1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 1, 1, 1},
				{1, 0, 1, 1, 1, 1, 0, 0},
				{1, 1, 0, 1, 1, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
			},
		},
	}

	for i, test := range tests {
		for state := 0; state < 5; state++ {
			for action := 0; action < 8; action++ {
				checker := NewStateTransform(test.creatorID, test.operator, DatasetState(state), DatasetAction(action))
				err := checker.Check()
				if test.stateTrans[state][action] == 1 && err == nil {
					t.Fatalf("test %d: expected error at state %d, action %d, but got nil", i, state, action)
				}
				if test.stateTrans[state][action] == 0 && err != nil {
					t.Fatalf("test %d: unexpected error at state %d, action %d: %v", i, state, action, err)
				}
			}
		}
	}
}
