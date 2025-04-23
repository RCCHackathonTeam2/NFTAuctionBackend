package types

type NftAttribute struct {
	TraitType        string  `json:"trait_type"`
	TraitValue       string  `json:"trait_value"`
	DisplayType      string  `json:"display_type"`
	RarityPercentage float32 `json:"rarity_percentage"`
}
