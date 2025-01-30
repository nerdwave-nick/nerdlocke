package pokeapi

import "fmt"

func (c *Client) EvolutionChains(limit int, offset int) (*APIResourceList, error) {
	return doUncached[APIResourceList](c, fmt.Sprintf("evolution-chain?limit=%d&offset=%d", limit, offset))
}

func (c *Client) EvolutionChain(idOrUrl string) (*EvolutionChain, error) {
	return do[EvolutionChain](c, fmt.Sprintf("evolution-chain/%s", idOrUrl))
}

func (c *Client) EvolutionTriggers(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("evolution-trigger?limit=%d&offset=%d", limit, offset))
}

func (c *Client) EvolutionTrigger(nameOrIdOrUrl string) (*EvolutionTrigger, error) {
	return do[EvolutionTrigger](c, fmt.Sprintf("evolution-trigger/%s", nameOrIdOrUrl))
}

type EvolutionChain struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The item that a Pokémon would be holding when mating that would trigger the egg hatching a baby Pokémon rather than a basic Pokémon.
	BabyTriggerItem NamedAPIResource `json:"baby_trigger_item"`
	// The base chain link object. Each link contains evolution details for a Pokémon in the chain. Each link references the next Pokémon in the natural evolution order.
	Chain ChainLink `json:"chain"`
}

type ChainLink struct {
	// Whether or not this link is for a baby Pokémon. This would only ever be true on the base link.
	IsBaby bool `json:"is_baby"`
	// The Pokémon species at this point in the evolution chain.
	Species NamedAPIResource `json:"species"`
	// All details regarding the specific details of the referenced Pokémon species evolution.
	EvolutionDetails []EvolutionDetail `json:"evolution_details"`
	// A List of chain objects.
	EvolvesTo []ChainLink `json:"evolves_to"`
}

type EvolutionDetail struct {
	// The item required to cause evolution this into Pokémon species.
	Item NamedAPIResource `json:"item"`
	// The type of event that triggers evolution into this Pokémon species.
	Trigger NamedAPIResource `json:"trigger"`
	// The id of the gender of the evolving Pokémon species must be in order to evolve into this Pokémon species.
	Gender int `json:"gender"`
	// The item the evolving Pokémon species must be holding during the evolution trigger event to evolve into this Pokémon species.
	HeldItem NamedAPIResource `json:"held_item"`
	// The move that must be known by the evolving Pokémon species during the evolution trigger event in order to evolve into this Pokémon species.
	KnownMove NamedAPIResource `json:"known_move"`
	// The evolving Pokémon species must know a move with this type during the evolution trigger event in order to evolve into this Pokémon species.
	KnownMoveType NamedAPIResource `json:"known_move_type"`
	// The location the evolution must be triggered at.
	Location NamedAPIResource `json:"location"`
	// The minimum required level of the evolving Pokémon species to evolve into this Pokémon species.
	MinLevel int `json:"min_level"`
	// The minimum required level of happiness the evolving Pokémon species to evolve into this Pokémon species.
	MinHappiness int `json:"min_happiness"`
	// The minimum required level of beauty the evolving Pokémon species to evolve into this Pokémon species.
	MinBeauty int `json:"min_beauty"`
	// The minimum required level of affection the evolving Pokémon species to evolve into this Pokémon species.
	MinAffection int `json:"min_affection"`
	// Whether or not it must be raining in the overworld to cause evolution this Pokémon species.
	NeedsOverworldRain bool `json:"needs_overworld_rain"`
	// The Pokémon species that must be in the players party in order for the evolving Pokémon species to evolve into this Pokémon species.
	Party_species NamedAPIResource `json:"party_species"`
	// The player must have a Pokémon of this type in their party during the evolution trigger event in order for the evolving Pokémon species to evolve into this Pokémon species.
	Party_type NamedAPIResource `json:"party_type"`
	// The required relation between the Pokémon's Attack and Defense stats. 1 means Attack > Defense. 0 means Attack = Defense. -1 means Attack < Defense.
	RelativePhysicalStats int `json:"relative_physical_stats"`
	// The required time of day. Day or night.
	TimeOfDay string `json:"time_of_day"`
	// Pokémon species for which this one must be traded.
	TradeSpecies NamedAPIResource `json:"trade_species"`
	// Whether or not the 3DS needs to be turned upside-down as this Pokémon levels up.
	TurnUpsideDown bool `json:"turn_upside_down"`
}

type EvolutionTrigger struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of pokemon species that result from this evolution trigger.
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}
