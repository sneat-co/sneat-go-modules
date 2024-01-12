package api4contactus

import (
	"context"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/contactus/dal4contactus"
	"github.com/sneat-co/sneat-go-modules/contactus/facade4contactus"
	"net/http"
)

var createMember = facade4contactus.CreateMember

// httpPostCreateMember is an API endpoint that adds a members to a team.
// While is very similar to contactus/api4contactus/http_create_contact.go, it's not the same.
func httpPostCreateMember(w http.ResponseWriter, r *http.Request) {
	var request dal4contactus.CreateMemberRequest
	handler := func(ctx context.Context, userCtx facade.User) (interface{}, error) {
		return createMember(ctx, userCtx, request)
	}
	apicore.HandleAuthenticatedRequestWithBody(w, r, &request, handler, http.StatusCreated, verify.DefaultJsonWithAuthRequired)
}
