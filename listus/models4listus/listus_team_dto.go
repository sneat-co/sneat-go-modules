package models4listus

type ListusTeamDto struct {
	ListGroups []*ListGroup `json:"listGroups,omitempty" firestore:"listGroups,omitempty"`
}

func (v ListusTeamDto) Validate() error {
	return nil
}
