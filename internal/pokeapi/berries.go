package pokeapi

import "fmt"

func (c *Client) Berry(nameOrId string) (*Berry, error) {
	return do[Berry](c, fmt.Sprintf("berry/%s", nameOrId))
}

func (c *Client) BerryFirmness(nameOrId string) (*BerryFirmness, error) {
	return do[BerryFirmness](c, fmt.Sprintf("berry-firmness/%s", nameOrId))
}

func (c *Client) BerryFlavor(nameOrId string) (*BerryFlavor, error) {
	return do[BerryFlavor](c, fmt.Sprintf("berry-flavor/%s", nameOrId))
}

type Berry struct {
	// The identifier for this berry
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// Time it takes the tree to grow one stage, in hours. Berry trees go through four of these growth stages before they can be picked.
	GrowthTime int `json:"growth_time"`
	// The maximum number of these berries that can grow on one tree in Generation IV.
	MaxHarvest int `json:"max_harvest"`
	// The power of the move "Natural Gift" when used with this Berry.
	NaturalGiftPower int `json:"natural_gift_power"`
	// The size of this Berry, in millimeters.
	Size int `json:"size"`
	// The smoothness of this Berry, used in making Pokéblocks or Poffins.
	Smoothness int `json:"smoothness"`
	// The speed at which this Berry dries out the soil as it grows. A higher rate means the soil dries more quickly.
	SoilDryness int `json:"soil_dryness"`
	// The firmness of this berry, used in making Pokéblocks or Poffins.
	Firmness NamedAPIResource `json:"firmness"`
	// A list of references to each flavor a berry can have and the potency of each of those flavors in regard to this berry.
	Flavors []BerryFlavorMap `json:"flavors"`
	// Berries are actually items. This is a reference to the item specific data for this berry.
	Item NamedAPIResource `json:"item"`
	// The type inherited by "Natural Gift" when used with this Berry.
	NaturalGiftType NamedAPIResource `json:"natural_gift_type"`
}

type BerryFlavorMap struct {
	// How powerful the referenced flavor is for this berry.
	Potency int `json:"potency"`
	// The referenced berry flavor.
	Flavor NamedAPIResource `json:"flavor"`
}

type BerryFirmness struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A list of the berries with this firmness.
	Berries []NamedAPIResource `json:"berries"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}

type BerryFlavor struct {
	// The identifier for this resource.
	ID int `json:"id"`
	// The name for this resource.
	Name string `json:"name"`
	// A list of the berries with this flavor.
	Berries []FlavorBerryMap `json:"berries"`
	// The contest type that correlates with this berry flavor.
	ContestType NamedAPIResource `json:"contest_type"`
	// The name of this resource listed in different languages.
	Names []Name `json:"names"`
}

type FlavorBerryMap struct {
	// How powerful the referenced flavor is for this berry.
	Potency int `json:"potency"`
	// The berry with the referenced flavor.
	Berry NamedAPIResource `json:"berry"`
}
