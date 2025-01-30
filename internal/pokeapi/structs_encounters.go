package pokeapi

type EncounterMethod struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A good value for sorting.
	Order int `json:"order"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}

type EncounterCondition struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
	// A list of possible values for this encounter condition.
	Values []NamedAPIResource `json:"values"`
}

type EncounterConditionValue struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// The condition this encounter condition value pertains to.
	Condition NamedAPIResource `json:"condition"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}
