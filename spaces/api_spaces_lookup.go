package spaces

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/spaces/types"
)

const (
	SpacesLookupIDEndpoint           = "https://api.twitter.com/2/spaces/:id"
	SpacesLookupEndpoint             = "https://api.twitter.com/2/spaces"
	SpacesLookupByCreatorIDsEndpoint = "https://api.twitter.com/2/spaces/by/creator_ids"
)

// Returns a variety of information about a single Space specified by the requested ID.
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-id
func SpacesLookupID(ctx context.Context, c *gotwi.GotwiClient, p *types.SpacesLookupIDParams) (*types.SpacesLookupIDResponse, error) {
	res := &types.SpacesLookupIDResponse{}
	if err := c.CallAPI(ctx, SpacesLookupIDEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns details about multiple Spaces. Up to 100 comma-separated Spaces IDs can be looked up using this endpoint
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces
func SpacesLookup(ctx context.Context, c *gotwi.GotwiClient, p *types.SpacesLookupParams) (*types.SpacesLookupResponse, error) {
	res := &types.SpacesLookupResponse{}
	if err := c.CallAPI(ctx, SpacesLookupEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns live or scheduled Spaces created by the specified user IDs.
// Up to 100 comma-separated IDs can be looked up using this endpoint.
// https://developer.twitter.com/en/docs/twitter-api/spaces/lookup/api-reference/get-spaces-by-creator-ids#Optional
func SpacesLookupByCreatorIDs(ctx context.Context, c *gotwi.GotwiClient, p *types.SpacesLookupByCreatorIDsParams) (*types.SpacesLookupByCreatorIDsResponse, error) {
	res := &types.SpacesLookupByCreatorIDsResponse{}
	if err := c.CallAPI(ctx, SpacesLookupByCreatorIDsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
