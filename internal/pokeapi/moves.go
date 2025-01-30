package pokeapi

import "fmt"

func (c *Client) Moves(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("move?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Move(nameOrIdOrUrl string) (*Move, error) {
	return do[Move](c, fmt.Sprintf("move/%s", nameOrIdOrUrl))
}

func (c *Client) MoveAilments(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("move-ailment?limit=%d&offset=%d", limit, offset))
}

func (c *Client) MoveAilment(nameOrIdOrUrl string) (*MoveAilment, error) {
	return do[MoveAilment](c, fmt.Sprintf("move-ailment/%s", nameOrIdOrUrl))
}

func (c *Client) MoveBattleStyles(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("move-battle-style?limit=%d&offset=%d", limit, offset))
}

func (c *Client) MoveBattleStyle(nameOrIdOrUrl string) (*MoveBattleStyle, error) {
	return do[MoveBattleStyle](c, fmt.Sprintf("move-battle-style/%s", nameOrIdOrUrl))
}

func (c *Client) MoveCategories(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("move-category?limit=%d&offset=%d", limit, offset))
}

func (c *Client) MoveCategory(nameOrIdOrUrl string) (*MoveCategory, error) {
	return do[MoveCategory](c, fmt.Sprintf("move-category/%s", nameOrIdOrUrl))
}

func (c *Client) MoveDamageClasses(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("move-damage-class?limit=%d&offset=%d", limit, offset))
}

func (c *Client) MoveDamageClass(nameOrIdOrUrl string) (*MoveDamageClass, error) {
	return do[MoveDamageClass](c, fmt.Sprintf("move-damage-class/%s", nameOrIdOrUrl))
}

func (c *Client) MoveLearnMethods(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("move-learn-method?limit=%d&offset=%d", limit, offset))
}

func (c *Client) MoveLearnMethod(nameOrIdOrUrl string) (*MoveLearnMethod, error) {
	return do[MoveLearnMethod](c, fmt.Sprintf("move-learn-method/%s", nameOrIdOrUrl))
}

func (c *Client) MoveTargets(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("move-target?limit=%d&offset=%d", limit, offset))
}

func (c *Client) MoveTarget(nameOrIdOrUrl string) (*MoveTarget, error) {
	return do[MoveTarget](c, fmt.Sprintf("move-target/%s", nameOrIdOrUrl))
}

type Move struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The percent value of how likely this move is to be successful.
	Accuracy int `json:"accuracy"`
	// The percent value of how likely it is this moves effect will happen.
	EffectChance int `json:"effect_chance"`
	// Power points. The number of times this move can be used.
	PP int `json:"pp"`
	// A value between -8 and 8. Sets the order in which moves are executed during battle. See Bulbapedia for greater detail.
	Priority int `json:"priority"`
	// The base power of this move with a value of 0 if it does not have a base power.
	Power int `json:"power"`
	// A detail of normal and super contest combos that require this move.
	ContestCombos ContestComboSets `json:"contest_combos"`
	// The type of appeal this move gives a Pokémon when used in a contest.
	ContestType NamedAPIResource `json:"contest_type"`
	// The effect the move has when used in a contest.
	ContestEffect APIResource `json:"contest_effect"`
	// The type of damage the move inflicts on the target, e.g. physical.
	DamageClass NamedAPIResource `json:"damage_class"`
	// The effect of this move listed in different languages.
	EffectEntries []VerboseEffect `json:"effect_entries"`
	// The list of previous effects this move has had across version groups of the games.
	EffectChanges []AbilityEffectChange `json:"effect_changes"`
	// List of Pokemon that can learn the move
	LearnedByPokemon []NamedAPIResource `json:"learned_by_pokemon"`
	// The flavor text of this move listed in different languages.
	FlavorTextEntries []MoveFlavorText `json:"flavor_text_entries"`
	// The generation in which this move was introduced.
	Generation NamedAPIResource `json:"generation"`
	// A list of the machines that teach this move.
	Machines []MachineVersionDetail `json:"machines"`
	// Metadata about this move
	Meta MoveMetaData `json:"meta"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of move resource value changes across version groups of the game.
	PastValues []PastMoveStatValues `json:"past_values"`
	// A list of stats this moves effects and how much it effects them.
	StatChanges []MoveStatChange `json:"stat_changes"`
	// The effect the move has when used in a super contest.
	SuperContest_effect APIResource `json:"super_contest_effect"`
	// The type of target that will receive the effects of the attack.
	Target NamedAPIResource `json:"target"`
	// The elemental type of this move.
	Type NamedAPIResource `json:"type"`
}

type ContestComboSets struct {
	// A detail of moves this move can be used before or after, granting additional appeal points in contests.
	Normal ContestComboDetail `json:"normal"`
	// A detail of moves this move can be used before or after, granting additional appeal points in super contests.
	Super ContestComboDetail `json:"super"`
}

type ContestComboDetail struct {
	// A list of moves to use before this move.
	UseBefore []NamedAPIResource `json:"use_before"`
	// A list of moves to use after this move.
	UseAfter []NamedAPIResource `json:"use_after"`
}

type MoveFlavorText struct {
	// The localized flavor text for an api resource in a specific language.
	FlavorText string `json:"flavor_text"`
	// The language this name is in.
	Language NamedAPIResource `json:"language"`
	// The version group that uses this flavor text.
	VersionGroup NamedAPIResource `json:"version_group"`
}

type MoveMetaData struct {
	// The status ailment this move inflicts on its target.
	Ailment NamedAPIResource `json:"ailment"`
	// The category of move this move falls under, e.g. damage or ailment.
	Category NamedAPIResource `json:"category"`
	// The minimum number of times this move hits. Null if it always only hits once.
	MinHits int `json:"min_hits"`
	// The maximum number of times this move hits. Null if it always only hits once.
	MaxHits int `json:"max_hits"`
	// The minimum number of turns this move continues to take effect. Null if it always only lasts one turn.
	MinTurns int `json:"min_turns"`
	// The maximum number of turns this move continues to take effect. Null if it always only lasts one turn.
	MaxTurns int `json:"max_turns"`
	// HP drain (if positive) or Recoil damage (if negative), in percent of damage done.
	Drain int `json:"drain"`
	// The amount of hp gained by the attacking Pokemon, in percent of it's maximum HP.
	Healing int `json:"healing"`
	// Critical hit rate bonus.
	CritRate int `json:"crit_rate"`
	// The likelihood this attack will cause an ailment.
	AilmentChance int `json:"ailment_chance"`
	// The likelihood this attack will cause the target Pokémon to flinch.
	FlinchChance int `json:"flinch_chance"`
	// The likelihood this attack will cause a stat change in the target Pokémon.
	StatChance int `json:"stat_chance"`
}

type MoveStatChange struct {
	// The amount of change.
	Change int `json:"change"`
	// The stat being affected.
	Stat NamedAPIResource `json:"stat"`
}

type PastMoveStatValues struct {
	// The percent value of how likely this move is to be successful.
	Accuracy int `json:"accuracy"`
	// The percent value of how likely it is this moves effect will take effect.
	EffectChance int `json:"effect_chance"`
	// The base power of this move with a value of 0 if it does not have a base power.
	Power int `json:"power"`
	// Power points. The number of times this move can be used.
	Pp int `json:"pp"`
	// The effect of this move listed in different languages.
	EffectEntries []VerboseEffect `json:"effect_entries"`
	// The elemental type of this move.
	Type NamedAPIResource `json:"type"`
	// The version group in which these move stat values were in effect.
	VersionGroup NamedAPIResource `json:"version_group"`
}

type MoveAilment struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A list of moves that cause this ailment.
	Moves []NamedAPIResource `json:"moves"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}

type MoveBattleStyle struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}

type MoveCategory struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A list of moves that fall into this category.
	Moves []NamedAPIResource `json:"moves"`
	// The description of this resource listed in different languages.
	Descriptions []Description `json:"descriptions"`
}

type MoveDamageClass struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The description of this resource listed in different languages.
	Descriptions []Description `json:"descriptions"`
	// A list of moves that fall into this damage class.
	Moves []NamedAPIResource `json:"moves"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}

type MoveLearnMethod struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The description of this resource listed in different languages.
	Descriptions []Description `json:"descriptions"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of version groups where moves can be learned through this method.
	VersionGroups []NamedAPIResource `json:"version_groups"`
}

type MoveTarget struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The description of this resource listed in different languages.
	Descriptions []Description `json:"descriptions"`
	// A list of moves that that are directed at this target.
	Moves []NamedAPIResource `json:"moves"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}
