package pokeapi

import "fmt"

func (c *Client) Machines(limit int, offset int) (*NamedAPIResourceList, error) {
	return doUncached[NamedAPIResourceList](c, fmt.Sprintf("machine?limit=%d&offset=%d", limit, offset))
}

func (c *Client) Machine(idOrUrl string) (*Machine, error) {
	return do[Machine](c, fmt.Sprintf("machine/%s", idOrUrl))
}

type Machine struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The TM or HM item that corresponds to this machine.
	Item NamedAPIResource `json:"item"`
	// The move that is taught by this machine.
	Move NamedAPIResource `json:"move"`
	// The version group that this machine applies to.
	Version_group NamedAPIResource `json:"version_group"`
}
