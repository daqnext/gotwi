package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type SpacesLookupIDParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	Expansions  fields.ExpansionList
	SpaceFields fields.SpaceFieldList
	UserFields  fields.UserFieldList
}

var SpacesLookupIDQueryParams = map[string]struct{}{
	"expansions":   {},
	"space.fields": {},
	"user.fields":  {},
}

func (p *SpacesLookupIDParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SpacesLookupIDParams) AccessToken() string {
	return p.accessToken
}

func (p *SpacesLookupIDParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, SpacesLookupIDQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *SpacesLookupIDParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *SpacesLookupIDParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.SpaceFields, p.UserFields)
	return m
}

// SpacesLookupParams is struct of parameters
// for request GET /2/spaces
type SpacesLookupParams struct {
	accessToken string

	// Query parameters
	IDs         []string
	Expansions  fields.ExpansionList
	SpaceFields fields.SpaceFieldList
	UserFields  fields.UserFieldList
}

var SpacesLookupQueryParams = map[string]struct{}{
	"ids":          {},
	"expansions":   {},
	"space.fields": {},
	"user.fields":  {},
}

func (p *SpacesLookupParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SpacesLookupParams) AccessToken() string {
	return p.accessToken
}

func (p *SpacesLookupParams) ResolveEndpoint(endpointBase string) string {
	if p.IDs == nil || len(p.IDs) == 0 {
		return ""
	}

	endpoint := endpointBase

	pm := p.ParameterMap()
	qs := util.QueryString(pm, SpacesLookupQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *SpacesLookupParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *SpacesLookupParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m["ids"] = util.QueryValue(p.IDs)
	m = fields.SetFieldsParams(m, p.Expansions, p.SpaceFields, p.UserFields)
	return m
}

// SpacesLookupByCreatorIDsParams is struct of parameters
// for request GET /2/spaces/by/creator_ids
type SpacesLookupByCreatorIDsParams struct {
	accessToken string

	// Query parameters
	UserIDs     []string
	Expansions  fields.ExpansionList
	SpaceFields fields.SpaceFieldList
	UserFields  fields.UserFieldList
}

var SpacesLookupByCreatorIDsQueryParams = map[string]struct{}{
	"user_ids":     {},
	"expansions":   {},
	"space.fields": {},
	"user.fields":  {},
}

func (p *SpacesLookupByCreatorIDsParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SpacesLookupByCreatorIDsParams) AccessToken() string {
	return p.accessToken
}

func (p *SpacesLookupByCreatorIDsParams) ResolveEndpoint(endpointBase string) string {
	if p.UserIDs == nil || len(p.UserIDs) == 0 {
		return ""
	}

	endpoint := endpointBase

	pm := p.ParameterMap()
	qs := util.QueryString(pm, SpacesLookupByCreatorIDsQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *SpacesLookupByCreatorIDsParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *SpacesLookupByCreatorIDsParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m["user_ids"] = util.QueryValue(p.UserIDs)
	m = fields.SetFieldsParams(m, p.Expansions, p.SpaceFields, p.UserFields)
	return m
}
