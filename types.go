package geocodio

type Address struct {
	FormattedAddress string   `json:"formatted_address"`
	Location         Location `json:"location"`
	Accuracy         float64  `json:"accuracy"`
	AccuracyType     string   `json:"accuracy_type"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type GeocodeResults struct {
	Results []Address `json:"results"`
}
