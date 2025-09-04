package state_transform

type DatasetState int

const (
	DatasetStateInvalid     DatasetState = 0
	DatasetStateDevelopment DatasetState = 1
	DatasetStateReleased    DatasetState = 2
	DatasetStateOutdated    DatasetState = 3
	DatasetStateBanned      DatasetState = 4
)

type DatasetAction int

const (
	DatasetActionDevelop DatasetAction = 0
	DatasetActionPublish DatasetAction = 1
	DatasetActionOutdate DatasetAction = 2
	DatasetActionBan     DatasetAction = 3
	DatasetActionDelete  DatasetAction = 4
	DatasetActionModify  DatasetAction = 5
	DatasetActionRead    DatasetAction = 6
	DatasetActionExec    DatasetAction = 7
)

type User struct {
	ID    string
	Roles []string
}

type StateChecker interface {
	Check() error
}
