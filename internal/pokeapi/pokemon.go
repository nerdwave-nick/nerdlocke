package pokeapi

import "fmt"

func (c *Client) Abilities(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("abilitie?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Ability(nameOrIdOrUrl string) (*Ability, error) {
	return do[Ability](c, fmt.Sprintf("ability/%s", nameOrIdOrUrl))
}

func (c *Client) Characteristics(limit int, offset int) (*APIResourceList, error) {
	return doUncached[APIResourceList](c, fmt.Sprintf("characteristic?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Characteristic(idOrUrl string) (*Characteristic, error) {
	return do[Characteristic](c, fmt.Sprintf("characteristic/%s", idOrUrl))
}

func (c *Client) EggGroups(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("egg-group?limit=%d&offset=%d", limit, offset))
}

func (c *Client) EggGroup(nameOrIdOrUrl string) (*EggGroup, error) {
	return do[EggGroup](c, fmt.Sprintf("egg-group/%s", nameOrIdOrUrl))
}

func (c *Client) Genders(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("gender?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Gender(nameOrIdOrUrl string) (*Gender, error) {
	return do[Gender](c, fmt.Sprintf("gender/%s", nameOrIdOrUrl))
}

func (c *Client) GrowthRates(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("growth-rate?limit=%d&offset=%d", limit, offset))
}

func (c *Client) GrowthRate(nameOrIdOrUrl string) (*GrowthRate, error) {
	return do[GrowthRate](c, fmt.Sprintf("growth-rate/%s", nameOrIdOrUrl))
}

func (c *Client) Natures(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("nature?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Nature(nameOrIdOrUrl string) (*Nature, error) {
	return do[Nature](c, fmt.Sprintf("nature/%s", nameOrIdOrUrl))
}

func (c *Client) PokeathlonStats(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("pokeathlon-stat?limit=%d&offset=%d", limit, offset))
}

func (c *Client) PokeathlonStat(nameOrIdOrUrl string) (*PokeathlonStat, error) {
	return do[PokeathlonStat](c, fmt.Sprintf("pokeathlon-stat/%s", nameOrIdOrUrl))
}

func (c *Client) Pokemons(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("pokemon?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Pokemon(nameOrIdOrUrl string) (*Pokemon, error) {
	return do[Pokemon](c, fmt.Sprintf("pokemon/%s", nameOrIdOrUrl))
}

func (c *Client) PokemonLocationAreas(pokemonIdOrName string) (*[]LocationAreaEncounter, error) {
	return do[[]LocationAreaEncounter](c, fmt.Sprintf("pokemon/%s/encounters", pokemonIdOrName))
}

func (c *Client) PokemonColors(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("pokemon-color?limit=%d&offset=%d", limit, offset))
}

func (c *Client) PokemonColor(nameOrIdOrUrl string) (*PokemonColor, error) {
	return do[PokemonColor](c, fmt.Sprintf("pokemon-color/%s", nameOrIdOrUrl))
}

func (c *Client) PokemonForms(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("pokemon-form?limit=%d&offset=%d", limit, offset))
}

func (c *Client) PokemonForm(nameOrIdOrUrl string) (*PokemonForm, error) {
	return do[PokemonForm](c, fmt.Sprintf("pokemon-form/%s", nameOrIdOrUrl))
}

func (c *Client) PokemonHabitats(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("pokemon-habitat?limit=%d&offset=%d", limit, offset))
}

func (c *Client) PokemonHabitat(nameOrIdOrUrl string) (*PokemonHabitat, error) {
	return do[PokemonHabitat](c, fmt.Sprintf("pokemon-habitat/%s", nameOrIdOrUrl))
}

func (c *Client) PokemonShapes(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("pokemon-shape?limit=%d&offset=%d", limit, offset))
}

func (c *Client) PokemonShape(nameOrIdOrUrl string) (*PokemonShape, error) {
	return do[PokemonShape](c, fmt.Sprintf("pokemon-shape/%s", nameOrIdOrUrl))
}

func (c *Client) PokemonSpecieses(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("pokemon-species?limit=%d&offset=%d", limit, offset))
}

func (c *Client) PokemonSpecies(nameOrIdOrUrl string) (*PokemonSpecies, error) {
	return do[PokemonSpecies](c, fmt.Sprintf("pokemon-species/%s", nameOrIdOrUrl))
}

func (c *Client) Stats(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("stat=%d&offset=%d", limit, offset))
}

func (c *Client) Stat(nameOrIdOrUrl string) (*Stat, error) {
	return do[Stat](c, fmt.Sprintf("stat/%s", nameOrIdOrUrl))
}

func (c *Client) Types(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("type?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Type(nameOrIdOrUrl string) (*Type, error) {
	return do[Type](c, fmt.Sprintf("type/%s", nameOrIdOrUrl))
}

type Ability struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// Whether or not this ability originated in the main series of the video games.
	IsMainSeries bool `json:"is_main_series"`
	// The generation this ability originated in.
	Generation NamedAPIResource `json:"generation"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// The effect of this ability listed in different languages.
	EffectEntries []VerboseEffect `json:"effect_entries"`
	// The list of previous effects this ability has had across version groups.
	EffectChanges []AbilityEffectChange `json:"effect_changes"`
	// The flavor text of this ability listed in different languages.
	FlavorTextEntries []AbilityFlavorText `json:"flavor_text_entries"`
	// A list of Pokémon that could potentially have this ability.
	Pokemon []AbilityPokemon `json:"pokemon"`
}

type AbilityEffectChange struct {
	// The previous effect of this ability listed in different languages.
	EffectEntries []Effect `json:"effect_entries"`
	// The version group in which the previous effect of this ability originated.
	VersionGroup NamedAPIResource `json:"version_group"`
}

type AbilityFlavorText struct {
	// The localized name for an API resource in a specific language.
	FlavorText string `json:"flavor_text"`
	// The language this text resource is in.
	Language NamedAPIResource `json:"language"`
	// The version group that uses this flavor text.
	VersionGroup NamedAPIResource `json:"version_group"`
}

type AbilityPokemon struct {
	// Whether or not this a hidden ability for the referenced Pokémon.
	IsHidden bool `json:"is_hidden"`
	// Pokémon have 3 ability 'slots' which hold references to possible abilities they could have. This is the slot of this ability for the referenced pokemon.
	Slot int `json:"slot"`
	// The Pokémon this ability could belong to.
	Pokemon NamedAPIResource `json:"pokemon"`
}

type Characteristic struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The remainder of the highest stat/IV divided by 5.
	GeneModulo int `json:"gene_modulo"`
	// The possible values of the highest stat that would result in a Pokémon recieving this characteristic when divided by 5.
	PossibleValues []int `json:"possible_values"`
	// The stat which results in this characteristic.
	HighestStat NamedAPIResource `json:"highest_stat"`
	// The descriptions of this characteristic listed in different languages.
	Descriptions []Description `json:"descriptions"`
}

type EggGroup struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of all Pokémon species that are members of this egg group.
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type Gender struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A list of Pokémon species that can be this gender and how likely it is that they will be.
	PokemonSpeciesDetails []PokemonSpeciesGender `json:"pokemon_species_details"`
	// A list of Pokémon species that required this gender in order for a Pokémon to evolve into them.
	RequiredForEvolution []NamedAPIResource `json:"required_for_evolution"`
}

type PokemonSpeciesGender struct {
	// The chance of this Pokémon being female, in eighths; or -1 for genderless.
	Rate int `json:"rate"`
	// A Pokémon species that can be the referenced gender.
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

type GrowthRate struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The formula used to calculate the rate at which the Pokémon species gains level.
	Formula string `json:"formula"`
	// The descriptions of this characteristic listed in different languages.
	Descriptions []Description `json:"descriptions"`
	// A list of levels and the amount of experienced needed to atain them based on this growth rate.
	Levels []GrowthRateExperienceLevel `json:"levels"`
	// A list of Pokémon species that gain levels at this growth rate.
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type GrowthRateExperienceLevel struct {
	// The level gained.
	Level int `json:"level"`
	// The amount of experience required to reach the referenced level.
	Experience int `json:"experience"`
}

type Nature struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The stat decreased by 10% in Pokémon with this nature.
	DecreasedStat NamedAPIResource `json:"decreased_stat"`
	// The stat increased by 10% in Pokémon with this nature.
	IncreasedStat NamedAPIResource `json:"increased_stat"`
	// The flavor hated by Pokémon with this nature.
	HatesFlavor NamedAPIResource `json:"hates_flavor"`
	// The flavor liked by Pokémon with this nature.
	LikesFlavor NamedAPIResource `json:"likes_flavor"`
	// A list of Pokéathlon stats this nature effects and how much it effects them.
	PokeathlonStatChanges []NatureStatChange `json:"pokeathlon_stat_changes"`
	// A list of battle styles and how likely a Pokémon with this nature is to use them in the Battle Palace or Battle Tent.
	MoveBattleStylePreferences []MoveBattleStylePreference `json:"move_battle_style_preferences"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}

type NatureStatChange struct {
	// The amount of change.
	MaxChange int `json:"max_change"`
	// The stat being affected.
	PokeathlonStat NamedAPIResource `json:"pokeathlon_stat"`
}

type MoveBattleStylePreference struct {
	// Chance of using the move, in percent, if HP is under one half.
	LowHpPreference int `json:"low_hp_preference"`
	// Chance of using the move, in percent, if HP is over one half.
	HighHpPreference int `json:"high_hp_preference"`
	// The move battle style.
	MoveBattleStyle NamedAPIResource `json:"move_battle_style"`
}

type PokeathlonStat struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A detail of natures which affect this Pokéathlon stat positively or negatively.
	AffectingNatures NaturePokeathlonStatAffectSets `json:"affecting_natures"`
}

type NaturePokeathlonStatAffectSets struct {
	// A list of natures and how they change the referenced Pokéathlon stat.
	Increase []NaturePokeathlonStatAffect `json:"increase"`
	// A list of natures and how they change the referenced Pokéathlon stat.
	Decrease []NaturePokeathlonStatAffect `json:"decrease"`
}

type NaturePokeathlonStatAffect struct {
	// The maximum amount of change to the referenced Pokéathlon stat.
	MaxChange int `json:"max_change"`
	// The nature causing the change.
	Nature NamedAPIResource `json:"nature"`
}

type Pokemon struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The base experience gained for defeating this Pokémon.
	BaseExperience int `json:"base_experience"`
	// The height of this Pokémon in decimetres.
	Height int `json:"height"`
	// Set for exactly one Pokémon used as the default for each species.
	IsDefault bool `json:"is_default"`
	// Order for sorting. Almost national order, except families are grouped together.
	Order int `json:"order"`
	// The weight of this Pokémon in hectograms.
	Weight int `json:"weight"`
	// A list of abilities this Pokémon could potentially have.
	Abilities []PokemonAbility `json:"abilities"`
	// A list of forms this Pokémon can take on.
	Forms []NamedAPIResource `json:"forms"`
	// A list of game indices relevent to Pokémon item by generation.
	GameIndices []VersionGameIndex `json:"game_indices"`
	// A list of items this Pokémon may be holding when encountered.
	HeldItems []PokemonHeldItem `json:"held_items"`
	// A link to a list of location areas, as well as encounter details pertaining to specific versions.
	LocationAreaEncounters string `json:"location_area_encounters"`
	// A list of moves along with learn methods and level details pertaining to specific version groups.
	Moves []PokemonMove `json:"moves"`
	// A list of details showing types this pokémon had in previous generations
	PastTypes []PokemonTypePast `json:"past_types"`
	// A set of sprites used to depict this Pokémon in the game. A visual representation of the various sprites can be found at PokeAPI/sprites
	Sprites PokemonSprites `json:"sprites"`
	// A set of cries used to depict this Pokémon in the game. A visual representation of the various cries can be found at PokeAPI/cries
	Cries PokemonCries `json:"cries"`
	// The species this Pokémon belongs to.
	Species NamedAPIResource `json:"species"`
	// A list of base stat values for this Pokémon.
	Stats []PokemonStat `json:"stats"`
	// A list of details showing types this Pokémon has.
	Types []PokemonType `json:"types"`
}

type PokemonAbility struct {
	// Whether or not this is a hidden ability.
	IsHidden bool `json:"is_hidden"`
	// The slot this ability occupies in this Pokémon species.
	Slot int `json:"slot"`
	// The ability the Pokémon may have.
	Ability NamedAPIResource `json:"ability"`
}

type PokemonType struct {
	// The order the Pokémon's types are listed in.
	Slot int `json:"slot"`
	// The type the referenced Pokémon has.
	Type NamedAPIResource `json:"type"`
}

type PokemonFormType struct {
	// The order the Pokémon's types are listed in.
	Slot int `json:"slot"`
	// The type the referenced Form has.
	Type NamedAPIResource `json:"type"`
}

type PokemonTypePast struct {
	// The last generation in which the referenced pokémon had the listed types.
	Generation NamedAPIResource `json:"generation"`
	// The types the referenced pokémon had up to and including the listed generation.
	Types []PokemonType `json:"types"`
}

type PokemonHeldItem struct {
	// The item the referenced Pokémon holds.
	Item NamedAPIResource `json:"item"`
	// The details of the different versions in which the item is held.
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	// The version in which the item is held.
	Version NamedAPIResource `json:"version"`
	// How often the item is held.
	Rarity int `json:"rarity"`
}

type PokemonMove struct {
	// The move the Pokémon can learn.
	Move NamedAPIResource `json:"move"`
	// The details of the version in which the Pokémon can learn the move.
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	// The method by which the move is learned.
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	// The version group in which the move is learned.
	VersionGroup NamedAPIResource `json:"version_group"`
	// The minimum level to learn the move.
	LevelLearnedAt int `json:"level_learned_at"`
}

type PokemonStat struct {
	// The stat the Pokémon has.
	Stat NamedAPIResource `json:"stat"`
	// The effort points (EV) the Pokémon has in the stat.
	Effort int `json:"effort"`
	// The base value of the stat.
	BaseStat int `json:"base_stat"`
}

type PokemonSprites struct {
	// The default depiction of this Pokémon from the front in battle.
	FrontDefault string `json:"front_default"`
	// The shiny depiction of this Pokémon from the front in battle.
	FrontShiny string `json:"front_shiny"`
	// The female depiction of this Pokémon from the front in battle.
	FrontFemale string `json:"front_female"`
	// The shiny female depiction of this Pokémon from the front in battle.
	FrontShinyFemale string `json:"front_shiny_female"`
	// The default depiction of this Pokémon from the back in battle.
	BackDefault string `json:"back_default"`
	// The shiny depiction of this Pokémon from the back in battle.
	BackShiny string `json:"back_shiny"`
	// The female depiction of this Pokémon from the back in battle.
	BackFemale string `json:"back_female"`
	// The shiny female depiction of this Pokémon from the back in battle.
	BackShinyFemale string `json:"back_shiny_female"`
}

type PokemonCries struct {
	// The latest depiction of this Pokémon's cry.
	Latest string `json:"latest"`
	// The legacy depiction of this Pokémon's cry.
	Legacy string `json:"legacy"`
}

type LocationAreaEncounter struct {
	// The location area the referenced Pokémon can be encountered in.
	LocationArea NamedAPIResource `json:"location_area"`
	// A list of versions and encounters with the referenced Pokémon that might happen.
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type PokemonColor struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of the Pokémon species that have this color.
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type PokemonForm struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The order in which forms should be sorted within all forms. Multiple forms may have equal order, in which case they should fall back on sorting by name.
	Order int `json:"order"`
	// The order in which forms should be sorted within a species' forms.
	FormOrder int `json:"form_order"`
	// True for exactly one form used as the default for each Pokémon.
	IsDefault bool `json:"is_default"`
	// Whether or not this form can only happen during battle.
	IsBattleOnly bool `json:"is_battle_only"`
	// Whether or not this form requires mega evolution.
	IsMega bool `json:"is_mega"`
	// The name of this form.
	FormName string `json:"form_name"`
	// The Pokémon that can take on this form.
	Pokemon NamedAPIResource `json:"pokemon"`
	// A list of details showing types this Pokémon form has.
	Types []PokemonFormType `json:"types"`
	// A set of sprites used to depict this Pokémon form in the game.
	Sprites PokemonFormSprites `json:"sprites"`
	// The version group this Pokémon form was introduced in.
	VersionGroup NamedAPIResource `json:"version_group"`
	// The form specific full name of this Pokémon form, or empty if the form does not have a specific name.
	Names []Name `json:"names"`
	// The form specific form name of this Pokémon form, or empty if the form does not have a specific name.
	FormNames []Name `json:"form_names"`
}

type PokemonFormSprites struct {
	// The default depiction of this Pokémon form from the front in battle.
	FrontDefault string `json:"front_default"`
	// The shiny depiction of this Pokémon form from the front in battle.
	FrontShiny string `json:"front_shiny"`
	// The default depiction of this Pokémon form from the back in battle.
	BackDefault string `json:"back_default"`
	// The shiny depiction of this Pokémon form from the back in battle.
	BackShiny string `json:"back_shiny"`
}

type PokemonHabitat struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of the Pokémon species that can be found in this habitat.
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type PokemonShape struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The "scientific" name of this Pokémon shape listed in different languages.
	AwesomeNames []AwesomeName `json:"awesome_names"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of the Pokémon species that have this shape.
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type AwesomeName struct {
	// The localized "scientific" name for an API resource in a specific language.
	AwesomeName string `json:"awesome_name"`
	// The language this "scientific" name is in.
	Language NamedAPIResource `json:"language"`
}

type PokemonSpecies struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The order in which species should be sorted. Based on National Dex order, except families are grouped together and sorted by stage.
	Order int `json:"order"`
	// The chance of this Pokémon being female, in eighths; or -1 for genderless.
	GenderRate int `json:"gender_rate"`
	// The base capture rate; up to 255. The higher the number, the easier the catch.
	CaptureRate int `json:"capture_rate"`
	// The happiness when caught by a normal Pokéball; up to 255. The higher the number, the happier the Pokémon.
	BaseHappiness int `json:"base_happiness"`
	// Whether or not this is a baby Pokémon.
	IsBaby bool `json:"is_baby"`
	// Whether or not this is a legendary Pokémon.
	IsLegendary bool `json:"is_legendary"`
	// Whether or not this is a mythical Pokémon.
	IsMythical bool `json:"is_mythical"`
	// Initial hatch counter: one must walk Y × (hatch_counter + 1) steps before this Pokémon's egg hatches, unless utilizing bonuses like Flame Body's. Y varies per generation. In Generations II, III, and VII, Egg cycles are 256 steps long. In Generation IV, Egg cycles are 255 steps long. In Pokémon Brilliant Diamond and Shining Pearl, Egg cycles are also 255 steps long, but are shorter on special dates. In Generations V and VI, Egg cycles are 257 steps long. In Pokémon Sword and Shield, and in Pokémon Scarlet and Violet, Egg cycles are 128 steps long.
	HatchCounter int `json:"hatch_counter"`
	// Whether or not this Pokémon has visual gender differences.
	HasGenderDifferences bool `json:"has_gender_differences"`
	// Whether or not this Pokémon has multiple forms and can switch between them.
	FormsSwitchable bool `json:"forms_switchable"`
	// The rate at which this Pokémon species gains levels.
	GrowthRate NamedAPIResource `json:"growth_rate"`
	// A list of Pokedexes and the indexes reserved within them for this Pokémon species.
	PokedexNumbers []PokemonSpeciesDexEntry `json:"pokedex_numbers"`
	// A list of egg groups this Pokémon species is a member of.
	EggGroups []NamedAPIResource `json:"egg_groups"`
	// The color of this Pokémon for Pokédex search.
	Color NamedAPIResource `json:"color"`
	// The shape of this Pokémon for Pokédex search.
	Shape NamedAPIResource `json:"shape"`
	// The Pokémon species that evolves into this Pokemon_species.
	EvolvesFromSpecies NamedAPIResource `json:"evolves_from_species"`
	// The evolution chain this Pokémon species is a member of.
	EvolutionChain APIResource `json:"evolution_chain"`
	// The habitat this Pokémon species can be encountered in.
	Habitat NamedAPIResource `json:"habitat"`
	// The generation this Pokémon species was introduced in.
	Generation NamedAPIResource `json:"generation"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of encounters that can be had with this Pokémon species in pal park.
	PalParkEncounters []PalParkEncounterArea `json:"pal_park_encounters"`
	// A list of flavor text entries for this Pokémon species.
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
	// Descriptions of different forms Pokémon take on within the Pokémon species.
	FormDescriptions []Description `json:"form_descriptions"`
	// The genus of this Pokémon species listed in multiple languages.
	Genera []Genus `json:"genera"`
	// A list of the Pokémon that exist within this Pokémon species.
	Varieties []PokemonSpeciesVariety `json:"varieties"`
}

type Genus struct {
	// The localized genus for the referenced Pokémon species
	Genus string `json:"genus"`
	// The language this genus is in.
	Language NamedAPIResource `json:"language"`
}

type PokemonSpeciesDexEntry struct {
	// The index number within the Pokédex.
	EntryNumber int `json:"entry_number"`
	// The Pokédex the referenced Pokémon species can be found in.
	Pokedex NamedAPIResource `json:"pokedex"`
}

type PalParkEncounterArea struct {
	// The base score given to the player when the referenced Pokémon is caught during a pal park run.
	BaseScore int `json:"base_score"`
	// The base rate for encountering the referenced Pokémon in this pal park area.
	Rate int `json:"rate"`
	// The pal park area where this encounter happens.
	Area NamedAPIResource `json:"area"`
}

type PokemonSpeciesVariety struct {
	// Whether this variety is the default variety.
	IsDefault bool `json:"is_default"`
	// The Pokémon variety.
	Pokemon NamedAPIResource `json:"pokemon"`
}

type Stat struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// ID the games use for this stat.
	GameIndex int `json:"game_index"`
	// Whether this stat only exists within a battle.
	IsBattleOnly bool `json:"is_battle_only"`
	// A detail of moves which affect this stat positively or negatively.
	AffectingMoves MoveStatAffectSets `json:"affecting_moves"`
	// A detail of natures which affect this stat positively or negatively.
	AffectingNatures NatureStatAffectSets `json:"affecting_natures"`
	// A list of characteristics that are set on a Pokémon when its highest base stat is this stat.
	Characteristics []APIResource `json:"characteristics"`
	// The class of damage this stat is directly related to.
	MoveDamageClass NamedAPIResource `json:"move_damage_class"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}

type MoveStatAffectSets struct {
	// A list of moves and how they change the referenced stat.
	Increase []MoveStatAffect `json:"increase"`
	// A list of moves and how they change the referenced stat.
	Decrease []MoveStatAffect `json:"decrease"`
}

type MoveStatAffect struct {
	// The maximum amount of change to the referenced stat.
	Change int `json:"change"`
	// The move causing the change.
	Move NamedAPIResource `json:"move"`
}

type NatureStatAffectSets struct {
	// A list of natures and how they change the referenced stat.
	Increase []NamedAPIResource `json:"increase"`
	// A list of nature sand how they change the referenced stat.
	Decrease []NamedAPIResource `json:"decrease"`
}

type Type struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A detail of how effective this type is toward others and vice versa.
	DamageRelations TypeRelations `json:"damage_relations"`
	// A list of details of how effective this type was toward others and vice versa in previous generations
	PastDamageRelations []TypeRelationsPast `json:"past_damage_relations"`
	// A list of game indices relevent to this item by generation.
	GameIndices []GenerationGameIndex `json:"game_indices"`
	// The generation this type was introduced in.
	Generation NamedAPIResource `json:"generation"`
	// The class of damage inflicted by this type.
	MoveDamageClass NamedAPIResource `json:"move_damage_class"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of details of Pokémon that have this type.
	Pokemon []TypePokemon `json:"pokemon"`
	// A list of moves that have this type.
	Moves []NamedAPIResource `json:"moves"`
}

type TypePokemon struct {
	// The order the Pokémon's types are listed in.
	Slot int `json:"slot"`
	// The Pokémon that has the referenced type.
	Pokemon NamedAPIResource `json:"pokemon"`
}

type TypeRelations struct {
	// A list of types this type has no effect on.
	NoDamageTo []NamedAPIResource `json:"no_damage_to"`
	// A list of types this type is not very effect against.
	HalfDamageTo []NamedAPIResource `json:"half_damage_to"`
	// A list of types this type is very effect against.
	DoubleDamageTo []NamedAPIResource `json:"double_damage_to"`
	// A list of types that have no effect on this type.
	NoDamageFrom []NamedAPIResource `json:"no_damage_from"`
	// A list of types that are not very effective against this type.
	HalfDamageFrom []NamedAPIResource `json:"half_damage_from"`
	// A list of types that are very effective against this type.
	DoubleDamageFrom []NamedAPIResource `json:"double_damage_from"`
}

type TypeRelationsPast struct {
	// The last generation in which the referenced type had the listed damage relations
	Generation NamedAPIResource `json:"generation"`
	// The damage relations the referenced type had up to and including the listed generation
	DamageRelations TypeRelations `json:"damage_relations"`
}
