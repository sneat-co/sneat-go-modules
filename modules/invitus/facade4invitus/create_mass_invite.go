package facade4invitus

import (
	"context"
	"github.com/sneat-co/sneat-go-modules/modules/invitus/models4invitus"
)

// CreateMassInviteRequest parameters for creating a mass invite
type CreateMassInviteRequest struct {
	Invite models4invitus.MassInvite `json:"invite"`
}

// Validate validates parameters for creating a mass invite
func (request *CreateMassInviteRequest) Validate() error {
	return request.Invite.Validate()
}

// CreateMassInviteResponse creating a mass invite
type CreateMassInviteResponse struct {
	ID string `json:"id"`
}

// CreateMassInvite creates a mass invite
func CreateMassInvite(_ context.Context, _ CreateMassInviteRequest) (response CreateMassInviteResponse, err error) {
	//request.InviteDto.TeamIDs.InviteID
	response.ID = ""
	return
}
