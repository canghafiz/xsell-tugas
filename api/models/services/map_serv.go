package services

import _map "be/models/requests/map"
import _mapRes "be/models/resources/map"

type MapServ interface {
	GetAddress(request _map.GetAddressRequest) (*_mapRes.ReverseResource, error)
	AutoComplete(request _map.AutocompleteRequest) ([]_mapRes.AutocompleteResource, error)
}
