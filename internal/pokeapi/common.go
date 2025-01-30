package pokeapi

import "fmt"

func (c *Client) Languages(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("language?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Language(nameOrIdOrUrl string) (*Language, error) {
	return do[Language](c, fmt.Sprintf("language/%s", nameOrIdOrUrl))
}

type NamedAPIResourceList struct {
	// The total number of resources available from this API.
	Count int `json:"count"`
	// The URL for the next page in the list.
	Next string `json:"next"`
	// The URL for the previous page in the list.
	Previous string `json:"previous"`
	// A list of named API resources.
	Results []NamedAPIResource `json:"results"`
}

type APIResourceList struct {
	// The total number of resources available from this API.
	Count int `json:"count"`
	// The URL for the next page in the list.
	Next string `json:"next"`
	// The URL for the previous page in the list.
	Previous string `json:"previous"`
	// A list of named API resources.
	Results []APIResource `json:"results"`
}

// Languages for translations of API resource information.
type Language struct {
	// Identifier for this language
	ID string `json:"id"`
	// The name of this language
	Name string `json:"name"`
	// Whether or not games are published in this language
	Official bool `json:"official"`
	// The two-letter code of the country where this language is spoken (not unique)
	Iso639 string `json:"iso639"`
	// The two-letter code of the language (not unique)
	Iso3166 string `json:"iso3166"`
	// The names of this languages listed in different languages
	Names []Name `json:"names"`
}

type APIResource struct {
	// The URL of the referenced resource
	URL string `json:"url"`
}

type NamedAPIResource struct {
	// The URL of the referenced resource
	URL string `json:"url"`
	// The name of the referenced resource
	Name string `json:"name"`
}

type Description struct {
	// The localized description for an API resource in a specific language
	Description string `json:"description"`
	// The language this description is in
	Language NamedAPIResource `json:"language"`
}

type Effect struct {
	// The localized effect text for an API resource in a specific language.
	Effect string `json:"effect"`
	// The language this effect is in
	Language NamedAPIResource `json:"language"`
}

type Encounter struct {
	// The lowest level the Pokemon could be encountered at
	MinLevel int `json:"min_level"`
	// The highest level the Pokemon could be encountered at
	MaxLevel int `json:"max_level"`
	// A list of condition values that must be in effect for this encounter to occur.
	ConditionValues NamedAPIResource `json:"condition_values"`
	// Percent chance that this encounter will occur.
	Chance int `json:"chance"`
	// The method by which this encounter happens.
	Method NamedAPIResource `json:"method"`
}

type FlavorText struct {
	// The localized flavor text for an API resource in a specific language. Note that this text is left unprocessed as it is found in game files. This means that it contains special characters that one might want to replace with their visible decodable version. Please check out this issue to find out more.
	FlavorText string `json:"flavor_text"`
	// The language this FlavorText is in.
	Language NamedAPIResource `json:"language"`
	// The game version this flavor text is extracted from.
	Version NamedAPIResource `json:"version"`
}

type GenerationGameIndex struct {
	// The internal id of an API resource within game data.
	GameIndex int `json:"game_index"`
	// The generation relevent to this game index.
	Generation NamedAPIResource `json:"generation"`
}

type MachineVersionDetail struct {
	// The machine that teaches a move from an item.
	Machine APIResource `json:"machine"`
	// The version group of this specific machine.
	VersionGroup NamedAPIResource `json:"version_group"`
}

type Name struct {
	// The localized name for an API resource in a specific language.
	Name string `json:"name"`
	// The language this name is in.
	Language NamedAPIResource `json:"language"`
}

type VerboseEffect struct {
	// The localized effect text for an API resource in a specific language.
	Effect string `json:"effect"`
	// The localized effect text in brief.
	ShortEffect string `json:"short_effect"`
	// The language this effect is in.
	Language NamedAPIResource `json:"language"`
}

type VersionEncounterDetail struct {
	// The game version this encounter happens in.
	Version NamedAPIResource `json:"version"`
	// The total percentage of all encounter potential.
	MaxChance int `json:"max_chance"`
	// A list of encounters and their specifics.
	EncounterDetails []Encounter `json:"encounter_details"`
}

type VersionGameIndex struct {
	// The internal id of an API resource within game data.
	GameIndex int `json:"game_index"`
	// The version relevent to this game index.
	Version NamedAPIResource `json:"version"`
}

type VersionGroupFlavorText struct {
	// The localized name for an API resource in a specific language.
	Text string `json:"text"`
	// The language this name is in.
	Language NamedAPIResource `json:"language"`
	// The version group which uses this flavor text.
	VersionGroup NamedAPIResource `json:"version_group"`
}
