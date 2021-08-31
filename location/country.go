package location

type Country struct {
	Alpha2  string `json:"alpha_2"`
	Alpha3  string `json:"alpha_3"`
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Numeric uint8  `json:"numeric"`
}
