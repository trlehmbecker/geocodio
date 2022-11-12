package geocodio

import (
	"errors"
)

type BulkResponse struct {
	Results map[string]BulkResults `json:"results"`
}

type BulkResults struct {
	Response GeocodeResults `json:"response"`
}

type BulkQueryParameterized struct {
	Key        string
	Street     string
	City       string
	State      string
	PostalCode string
	Country    string
}

func (c *Client) GeocodeBulk(queries []BulkQueryParameterized) (data *BulkResponse, err error) {
	if len(queries) == 0 {
		err = errors.New("no queries provided")
		return
	}

	payload := make(map[string]map[string]string)

	for _, element := range queries {
		payload[element.Key] = make(map[string]string, 5)

		if element.Street != "" {
			payload[element.Key]["street"] = element.Street
		}

		if element.City != "" {
			payload[element.Key]["city"] = element.City
		}

		if element.State != "" {
			payload[element.Key]["state"] = element.State
		}

		if element.PostalCode != "" {
			payload[element.Key]["postal_code"] = element.PostalCode
		}

		if element.Country != "" {
			payload[element.Key]["country"] = element.Country
		}
	}

	err = c.post("/geocode", payload, &data)
	if len(data.Results) == 0 {
		err = errors.New("no results")
		return
	}

	return
}
