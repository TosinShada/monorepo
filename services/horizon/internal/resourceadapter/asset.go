package resourceadapter

import (
	"context"

	protocol "github.com/TosinShada/monorepo/protocols/horizon"
	"github.com/TosinShada/monorepo/xdr"
)

func PopulateAsset(ctx context.Context, dest *protocol.Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
