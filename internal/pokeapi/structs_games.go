package pokeapi

type Generation struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A list of abilities that were introduced in this generation.
	Abilities []NamedAPIResource `json:"abilities"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// The main region travelled in this generation.
	MainRegion NamedAPIResource `json:"main_region"`
	// A list of moves that were introduced in this generation.
	Moves []NamedAPIResource `json:"moves"`
	// A list of Pokémon species that were introduced in this generation.
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
	// A list of types that were introduced in this generation.
	Types []NamedAPIResource `json:"types"`
	// A list of version groups that were introduced in this generation.
	VersionGroups []NamedAPIResource `json:"version_groups"`
}

type Pokedex struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// Whether or not this Pokédex originated in the main series of the video games.
	IsMainSeries bool `json:"is_main_series"`
	// The description of this resource listed in different languages.
	Descriptions []Description `json:"descriptions"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of Pokémon catalogued in this Pokédex and their indexes.
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
	// The region this Pokédex catalogues Pokémon for.
	Region NamedAPIResource `json:"region"`
	// A list of version groups this Pokédex is relevant to.
	VersionGroups []NamedAPIResource `json:"version_groups"`
}

type PokemonEntry struct {
	// The index of this Pokémon species entry within the Pokédex.
	EntryNumber int `json:"entry_number"`
	// The Pokémon species being encountered.
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

type Version struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// The version group this version belongs to.
	VersionGroup NamedAPIResource `json:"version_group"`
}

type VersionGroup struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// Order for sorting. Almost by date of release, except similar versions are grouped together.
	Order int `json:"order"`
	// The generation this version was introduced in.
	Generation NamedAPIResource `json:"generation"`
	// A list of methods in which Pokémon can learn moves in this version group.
	MoveLearnMethods []NamedAPIResource `json:"move_learn_methods"`
	// A list of Pokédexes introduces in this version group.
	Pokedexes []NamedAPIResource `json:"pokedexes"`
	// A list of regions that can be visited in this version group.
	Regions []NamedAPIResource `json:"regions"`
	// The versions this version group owns.
	Versions []NamedAPIResource `json:"versions"`
}
