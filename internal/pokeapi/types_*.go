package pokeapi

type LocationAreaResp struct {
	PokemonEncounters []struct {
		Pokemon struct{
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonResp struct {
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`

	PokemonStats []struct {
		BaseStat int `json:"base_stat"`
		Stat struct{
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	PokemonTypes []struct {
		Type struct{
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}	