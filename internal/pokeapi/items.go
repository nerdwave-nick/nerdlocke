package pokeapi

import "fmt"

func (c *Client) Items(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("item?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Item(nameOrIdOrUrl string) (*Item, error) {
	return do[Item](c, fmt.Sprintf("item/%s", nameOrIdOrUrl))
}

func (c *Client) ItemAttributes(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("item-attribute?limit=%d&offset=%d", limit, offset))
}

func (c *Client) ItemAttribute(nameOrIdOrUrl string) (*ItemAttribute, error) {
	return do[ItemAttribute](c, fmt.Sprintf("item-attribute/%s", nameOrIdOrUrl))
}

func (c *Client) ItemCategories(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("item-category?limit=%d&offset=%d", limit, offset))
}

func (c *Client) ItemCategory(nameOrIdOrUrl string) (*ItemCategory, error) {
	return do[ItemCategory](c, fmt.Sprintf("item-category/%s", nameOrIdOrUrl))
}

func (c *Client) ItemFlingEffects(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("item-fling-effect?limit=%d&offset=%d", limit, offset))
}

func (c *Client) ItemFlingEffect(nameOrIdOrUrl string) (*ItemFlingEffect, error) {
	return do[ItemFlingEffect](c, fmt.Sprintf("item-fling-effect/%s", nameOrIdOrUrl))
}

func (c *Client) ItemPockets(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("item-pocket?limit=%d&offset=%d", limit, offset))
}

func (c *Client) ItemPocket(nameOrIdOrUrl string) (*ItemPocket, error) {
	return do[ItemPocket](c, fmt.Sprintf("item-pocket/%s", nameOrIdOrUrl))
}

type Item struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The price of this item in stores.
	Cost int `json:"cost"`
	// The power of the move Fling when used with this item.
	FlingPower int `json:"fling_power"`
	// The effect of the move Fling when used with this item.
	FlingEffect NamedAPIResource `json:"fling_effect"`
	// A list of attributes this item has.
	Attributes []NamedAPIResource `json:"attributes"`
	// The category of items this item falls into.
	Category NamedAPIResource `json:"category"`
	// The effect of this ability listed in different languages.
	EffectEntries []VerboseEffect `json:"effect_entries"`
	// The flavor text of this ability listed in different languages.
	FlavorText_entries []VersionGroupFlavorText `json:"flavor_text_entries"`
	// A list of game indices relevent to this item by generation.
	GameIndices []GenerationGameIndex `json:"game_indices"`
	// The name of this item listed in different languages.
	Names []Name `json:"names"`
	// A set of sprites used to depict this item in the game.
	Sprites ItemSprites `json:"sprites"`
	// A list of Pokémon that might be found in the wild holding this item.
	HeldByPokemon []ItemHolderPokemon `json:"held_by_pokemon"`
	// An evolution chain this item requires to produce a bay during mating.
	BabyTriggerFor APIResource `json:"baby_trigger_for"`
	// A list of the machines related to this item.
	Machines []MachineVersionDetail `json:"machines"`
}

type ItemSprites struct {
	// The default depiction of this item.
	Default string `json:"default"`
}

type ItemHolderPokemon struct {
	// The Pokémon that holds this item.
	Pokemon NamedAPIResource `json:"pokemon"`
	// The details for the version that this item is held in by the Pokémon.
	VersionDetails []ItemHolderPokemonVersionDetail `json:"version_details"`
}

type ItemHolderPokemonVersionDetail struct {
	// How often this Pokémon holds this item in this version.
	Rarity int `json:"rarity"`
	// The version that this item is held in by the Pokémon.
	Version NamedAPIResource `json:"version"`
}

type ItemAttribute struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A list of items that have this attribute.
	Items []NamedAPIResource `json:"items"`
	// The name of this item attribute listed in different languages.
	Names []Name `json:"names"`
	// The description of this item attribute listed in different languages.
	Descriptions []Description `json:"descriptions"`
}

type ItemCategory struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A list of items that are a part of this category.
	Items []NamedAPIResource `json:"items"`
	// The name of this item category listed in different languages.
	Names []Name `json:"names"`
	// The pocket items in this category would be put in.
	Pocket NamedAPIResource `json:"pocket"`
}

type ItemFlingEffect struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The result of this fling effect listed in different languages.
	EffectEntries []Effect `json:"effect_entries"`
	// A list of items that have this fling effect.
	Items []NamedAPIResource `json:"items"`
}

type ItemPocket struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A list of item categories that are relevant to this item pocket.
	Categories []NamedAPIResource `json:"categories"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}
