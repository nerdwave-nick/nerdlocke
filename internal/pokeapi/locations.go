package pokeapi

import "fmt"

func (c *Client) Locations(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("location?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Location(nameOrIdOrUrl string) (*Location, error) {
	return do[Location](c, fmt.Sprintf("location/%s", nameOrIdOrUrl))
}

func (c *Client) LocationAreas(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("location-area?limit=%d&offset=%d", limit, offset))
}

func (c *Client) LocationArea(nameOrIdOrUrl string) (*LocationArea, error) {
	return do[LocationArea](c, fmt.Sprintf("location-area/%s", nameOrIdOrUrl))
}

func (c *Client) PalParkAreas(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("pal-park-area?limit=%d&offset=%d", limit, offset))
}

func (c *Client) PalParkArea(nameOrIdOrUrl string) (*PalParkArea, error) {
	return do[PalParkArea](c, fmt.Sprintf("pal-park-area/%s", nameOrIdOrUrl))
}

func (c *Client) Regions(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("region?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Region(nameOrIdOrUrl string) (*Region, error) {
	return do[Region](c, fmt.Sprintf("region/%s", nameOrIdOrUrl))
}

type Location struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The region this location can be found in.
	Region NamedAPIResource `json:"region"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of game indices relevent to this location by generation.
	GameIndices []GenerationGameIndex `json:"game_indices"`
	// Areas that can be found within this location.
	Areas []NamedAPIResource `json:"areas"`
}

type LocationArea struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The internal id of an API resource within game data.
	GameIndex int `json:"game_index"`
	// A list of methods in which Pokémon may be encountered in this area and how likely the method will occur depending on the version of the game.
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	// The region this location area can be found in.
	Location NamedAPIResource `json:"location"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of Pokémon that can be encountered in this area along with version specific details about the encounter.
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	// The method in which Pokémon may be encountered in an area..
	EncounterMethod NamedAPIResource `json:"encounter_method"`
	// The chance of the encounter to occur on a version of the game.
	VersionDetails []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	// The chance of an encounter to occur.
	Rate int `json:"rate"`
	// The version of the game in which the encounter can occur with the given chance.
	Version NamedAPIResource `json:"version"`
}

type PokemonEncounter struct {
	// The Pokémon being encountered.
	Pokemon NamedAPIResource `json:"pokemon"`
	// A list of versions and encounters with Pokémon that might happen in the referenced location area.
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type PalParkArea struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of Pokémon encountered in thi pal park area along with details.
	PokemonEncounters []PalParkEncounterSpecies `json:"pokemon_encounters"`
}

type PalParkEncounterSpecies struct {
	// The base score given to the player when this Pokémon is caught during a pal park run.
	Base_score int `json:"base_score"`
	// The base rate for encountering this Pokémon in this pal park area.
	Rate int `json:"rate"`
	// The Pokémon species being encountered.
	Pokemon_species NamedAPIResource `json:"pokemon_species"`
}

type Region struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// A list of locations that can be found in this region.
	Locations []NamedAPIResource `json:"locations"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// The generation this region was introduced in.
	MainGeneration NamedAPIResource `json:"main_generation"`
	// A list of pokédexes that catalogue Pokémon in this region.
	Pokedexes []NamedAPIResource `json:"pokedexes"`
	// A list of version groups where this region can be visited.
	VersionGroups []NamedAPIResource `json:"version_groups"`
}
