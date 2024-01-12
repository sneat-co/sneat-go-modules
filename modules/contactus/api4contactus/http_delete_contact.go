package api4contactus

import (
	"context"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/apicore/verify"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/modules/contactus/dto4contactus"
	"github.com/sneat-co/sneat-go-modules/modules/contactus/facade4contactus"
	"net/http"
)

// httpDeleteContact deletes contact
func httpDeleteContact(w http.ResponseWriter, r *http.Request) {
	var request dto4contactus.ContactRequest
	handler := func(ctx context.Context, userCtx facade.User) (interface{}, error) {
		return nil, facade4contactus.DeleteContact(ctx, userCtx, request)
	}
	apicore.HandleAuthenticatedRequestWithBody(w, r, &request, handler, http.StatusCreated, verify.DefaultJsonWithAuthRequired)
}
