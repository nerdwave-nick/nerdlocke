package pokeapi

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
