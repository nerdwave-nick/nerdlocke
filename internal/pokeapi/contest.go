package pokeapi

import "fmt"

func (c *Client) ContestTypes(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("contest-type?limit=%d&offset=%d", limit, offset))
}

func (c *Client) ContestType(nameOrIdOrUrl string) (*ContestType, error) {
	return do[ContestType](c, fmt.Sprintf("contest-type/%s", nameOrIdOrUrl))
}

func (c *Client) ContestEffects(limit int, offset int) (*APIResourceList, error) {
	return doUncached[APIResourceList](c, fmt.Sprintf("contest-effect?limit=%d&offset=%d", limit, offset))
}

func (c *Client) ContestEffect(idOrUrl string) (*ContestEffect, error) {
	return do[ContestEffect](c, fmt.Sprintf("contest-effect/%s", idOrUrl))
}

func (c *Client) SuperContestEffects(limit int, offset int) (*APIResourceList, error) {
	return doUncached[APIResourceList](c, fmt.Sprintf("super-contest-effect?limit=%d&offset=%d", limit, offset))
}

func (c *Client) SuperContestEffect(idOrUrl string) (*SuperContestEffect, error) {
	return do[SuperContestEffect](c, fmt.Sprintf("super-contest-effect/%s", idOrUrl))
}

type ContestType struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The berry flavor that correlates with this contest type.
	BerryFlavor NamedAPIResource `json:"berry_flavor"`
	// The name of this contest type listed in different languages.
	Names []ContestName `json:"names"`
}

type ContestName struct {
	// The name for this contest.
	Name string `json:"name"`
	// The color associated with this contest's name.
	Color string `json:"color"`
	// The language that this name is in.
	Language NamedAPIResource `json:"language"`
}

type ContestEffect struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The base number of hearts the user of this move gets.
	Appeal int `json:"appeal"`
	// The base number of hearts the user's opponent loses.
	Jam int `json:"jam"`
	// The result of this contest effect listed in different languages.
	EffectEntries []Effect `json:"effect_entries"`
	// The flavor text of this contest effect listed in different languages.
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
}

type SuperContestEffect struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The level of appeal this super contest effect has.
	Appeal int `json:"appeal"`
	// The flavor text of this super contest effect listed in different languages.
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
	// A list of moves that have the effect when used in super contests.
	Moves []NamedAPIResource `json:"moves"`
}
