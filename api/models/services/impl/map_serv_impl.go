package impl

import (
	_map "be/models/requests/map"
	_mapRes "be/models/resources/map"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type MapServImpl struct {
	ApiEndpoint string
	ApiKey      string
}

func NewMapServImpl(apiEndpoint string, apiKey string) *MapServImpl {
	return &MapServImpl{
		ApiEndpoint: apiEndpoint,
		ApiKey:      apiKey,
	}
}

// In backend, create helper function
func (serv *MapServImpl) makeAuthenticatedRequest(endpoint string, params url.Values) ([]byte, error) {
	baseURL := fmt.Sprintf("%s%s", serv.ApiEndpoint, endpoint)

	// API key stays in query param but only server-side
	params.Add("key", serv.ApiKey)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (serv *MapServImpl) GetAddress(request _map.GetAddressRequest) (*_mapRes.ReverseResource, error) {
	params := url.Values{}
	params.Add("lat", fmt.Sprintf("%f", request.Lat))
	params.Add("lon", fmt.Sprintf("%f", request.Lon))
	params.Add("format", "json")

	body, err := serv.makeAuthenticatedRequest("reverse", params)
	if err != nil {
		return nil, err
	}

	return _mapRes.ToReverseResource(body)
}

func (serv *MapServImpl) AutoComplete(request _map.AutocompleteRequest) ([]_mapRes.AutocompleteResource, error) {
	params := url.Values{}
	params.Add("q", request.Q)
	params.Add("limit", fmt.Sprintf("%d", request.Limit))

	body, err := serv.makeAuthenticatedRequest("autocomplete", params)
	if err != nil {
		return nil, err
	}

	return _mapRes.ToAutocompleteResources(body)
}
