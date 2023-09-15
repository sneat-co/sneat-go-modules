package api4assets

import (
	"context"
	"errors"
	"fmt"
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-modules/assetus/const4assets"
	"github.com/sneat-co/sneat-go-modules/assetus/facade4assets"
	"github.com/sneat-co/sneat-go-modules/assetus/models4assets"
	"net/http"
)

// httpPostCreateAsset creates an asset
func httpPostCreateAsset(w http.ResponseWriter, r *http.Request) {
	var request facade4assets.CreateAssetRequest
	assetCategory := r.URL.Query().Get("assetCategory")
	switch assetCategory {
	case const4assets.AssetCategoryVehicle:
		asset := models4assets.NewVehicleAssetDbData()
		asset.Title = asset.GenerateTitle()
		request.Asset = asset.VehicleAssetMainData
		request.DbData = asset
	case const4assets.AssetCategoryRealEstate:
		asset := models4assets.NewDwellingAssetDbData()
		request.Asset = asset.DwellingAssetMainDto
		request.DbData = asset
	case const4assets.AssetCategoryDocument:
		asset := models4assets.NewDocumentDbData()
		request.Asset = asset.DocumentMainData
		request.DbData = asset
	case "":
		apicore.ReturnError(r.Context(), w, r, errors.New("GET parameter 'assetCategory' is required"))
		return
	default:
		apicore.ReturnError(r.Context(), w, r, fmt.Errorf("unsupported asset category: %s", assetCategory))
		return
	}
	handler := func(ctx context.Context, userCtx facade.User) (interface{}, error) {
		asset, err := facade4assets.CreateAsset(ctx, userCtx, request)
		if err != nil {
			return nil, fmt.Errorf("failed to create asset: %w", err)
		}
		if asset.ID == "" {
			return nil, errors.New("asset created by facade does not have an ContactID")
		}
		if asset.Data == nil {
			return nil, errors.New("asset created by facade does not have a DTO")
		}
		if err = asset.Data.Validate(); err != nil {
			err = fmt.Errorf("asset created by facade is not valid: %w", err)
			return asset, err
		}
		return asset, nil
	}
	apicore.HandleAuthenticatedRequestWithBody(w, r, &request, handler, http.StatusCreated,
		endpoints.VerifyRequest{
			MinContentLength: apicore.MinJSONRequestSize,
			MaxContentLength: 10 * apicore.KB,
			AuthRequired:     true,
		},
	)
}
