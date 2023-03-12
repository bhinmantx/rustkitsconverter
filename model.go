package main

const (
	ContainerMainBP = "main"
	ContainerWear   = "wear"
	ContainerBelt   = "belt"
)

type Kit struct {
	Name               string    `json:""`
	Description        string    `json:""`
	RequiredPermission string    `json:""`
	MaximumUses        int64     `json:""`
	RequiredAuth       int64     `json:""`
	Cooldown           int64     `json:""`
	Cost               int64     `json:""`
	IsHidden           bool      `json:""`
	CopyPasteFile      string    `json:""`
	KitImage           string    `json:""`
	MainItems          []KitItem `json:""`
	WearItems          []KitItem `json:""`
	BeltItems          []KitItem `json:""`
}

type KitItem struct {
	Shortname          string    `json:"ShortName"`
	Skin               int       `json:"Skin"`
	Amount             int       `json:"Amount"`
	Condition          float64   `json:"Condition"`
	MaxCondition       float64   `json:"MaxCondition"`
	Ammo               int64     `json:"Ammo"`
	Ammotype           string    `json:"Ammotype"`
	Position           int64     `json:"Position"`
	Frequency          int64     `json:"Frequency"`
	BlueprintShortname string    `json:"BlueprintShortname"`
	Contents           []KitItem `json:"Contents"`
}

type OldKit struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Max         float64      `json:"max"`
	Cooldown    float64      `json:"cooldown"`
	Authlevel   float64      `json:"authlevel"`
	Hide        bool         `json:"hide"`
	NpcOnly     bool         `json:"npconly"`
	Permissions string       `json:"permission"`
	Image       string       `json:"image"`
	Building    string       `json:"building"`
	Items       []OldKitItem `json:"items"`
}

type OldKitItem struct {
	ItemID          int    `json:"itemid"`
	Container       string `json:"container"`
	Amount          int    `json:"amount"`
	SkinID          int    `json:"skinid"`
	Weapon          bool   `json:"weapon"`
	BlueprintTarget int    `json:"blueprintTarget"`
	Mods            []int  `json:"mods"`
}
