package pokeapi

type LocationArea struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type LocationAreasRes struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []LocationArea `json:"results"`
}
