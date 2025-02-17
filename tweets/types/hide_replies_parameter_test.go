package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_HideRepliesParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.HideRepliesParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.HideRepliesParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: has no required parameter",
			params: &types.HideRepliesParams{
				Hidden: gotwi.Bool(true),
			},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_HideRepliesParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.HideRepliesParams
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.HideRepliesParams{
				ID:     "test-id",
				Hidden: gotwi.Bool(true),
			},
			expect: strings.NewReader(`{"hidden":true}`),
		},
		{
			name: "ok: has both of path and json parameters",
			params: &types.HideRepliesParams{
				ID:     "test-id",
				Hidden: gotwi.Bool(false),
			},
			expect: strings.NewReader(`{"hidden":false}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.HideRepliesParams{ID: "id"},
			expect: strings.NewReader(`{}`),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, r)
		})
	}
}
