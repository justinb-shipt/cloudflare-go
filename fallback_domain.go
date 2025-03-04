package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// FallbackDomainResponse represents the response from the get fallback
// domain endpoints.
type FallbackDomainResponse struct {
	Response
	Result []FallbackDomain `json:"result"`
}

// FallbackDomain represents the individual domain struct.
type FallbackDomain struct {
	Suffix       string   `json:"suffix,omitempty"`
	Description  string   `json:"description,omitempty"`
	DNSServer    []string `json:"dns_server,omitempty"`
}

// ListFallbackDomains returns all fallback domains within an account.
//
// API reference: https://api.cloudflare.com/#devices-get-local-domain-fallback-list
func (api *API) ListFallbackDomains(ctx context.Context, accountID string) ([]FallbackDomain, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/fallback_domains", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []FallbackDomain{}, err
	}

	var fallbackDomainResponse FallbackDomainResponse
	err = json.Unmarshal(res, &fallbackDomainResponse)
	if err != nil {
		return []FallbackDomain{}, errors.Wrap(err, errUnmarshalError)
	}

	return fallbackDomainResponse.Result, nil
}

// UpdateFallbackDomain updates the existing fallback domain policy.
//
// API reference: https://api.cloudflare.com/#devices-set-local-domain-fallback-list
func (api *API) UpdateFallbackDomain(ctx context.Context, accountID string, domains []FallbackDomain) ([]FallbackDomain, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/fallback_domains", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, domains)
	if err != nil {
		return []FallbackDomain{}, err
	}

	var fallbackDomainResponse FallbackDomainResponse
	err = json.Unmarshal(res, &fallbackDomainResponse)
	if err != nil {
		return []FallbackDomain{}, errors.Wrap(err, errUnmarshalError)
	}

	return fallbackDomainResponse.Result, nil
}