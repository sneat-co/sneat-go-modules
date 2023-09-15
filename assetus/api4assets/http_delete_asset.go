package api4assets

import (
	"context"
	"fmt"
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-core/modules/teamus/dal4teamus"
	"github.com/sneat-co/sneat-go-modules/assetus/facade4assets"
	"net/http"
)

var deleteAsset = facade4assets.DeleteAsset

// httpDeleteAsset deletes assets
func httpDeleteAsset(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var request dal4teamus.TeamItemRequest
	request.TeamID = q.Get("team")
	request.ID = q.Get("id")
	handler := func(ctx context.Context, userCtx facade.User) (interface{}, error) {
		if err := deleteAsset(ctx, userCtx, request); err != nil {
			return nil, fmt.Errorf("failed to delete asset: %w", err)
		}
		return nil, nil
	}
	apicore.HandleAuthenticatedRequestWithBody(w, r, &request, handler, http.StatusNoContent,
		endpoints.VerifyRequest{
			MinContentLength: 0,
			MaxContentLength: 0,
			AuthRequired:     true,
		},
	)
}
