package pokeapi

import "fmt"

func (c *Client) EncounterMethods(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("encounter-method?limit=%d&offset=%d", limit, offset))
}

func (c *Client) EncounterMethod(nameOrIdOrUrl string) (*EncounterMethod, error) {
	return do[EncounterMethod](c, fmt.Sprintf("encounter-method/%s", nameOrIdOrUrl))
}

func (c *Client) EncounterConditions(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("encounter-condition?limit=%d&offset=%d", limit, offset))
}

func (c *Client) EncounterCondition(nameOrIdOrUrl string) (*EncounterCondition, error) {
	return do[EncounterCondition](c, fmt.Sprintf("encounter-condition/%s", nameOrIdOrUrl))
}

func (c *Client) EncounterConditionValues(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("encounter-condition-value?limit=%d&offset=%d", limit, offset))
}

func (c *Client) EncounterConditionValue(nameOrIdOrUrl string) (*EncounterConditionValue, error) {
	return do[EncounterConditionValue](c, fmt.Sprintf("encounter-condition-value/%s", nameOrIdOrUrl))
}

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
