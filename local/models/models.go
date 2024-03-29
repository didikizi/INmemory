package models

type Users struct {
	Account int     `json:"account"`
	Name    string  `json:"name"`
	Value   float64 `json:"value"`
	Pass    string  `json:"-"`
}

type Filter struct {
	Account int     `query:"Account"`
	Name    string  `query:"Name"`
	Value   float64 `query:"Value"`
}

type SearchName struct {
	Account int
	Value   float64
}

type SearchAccount struct {
	Name  string
	Value float64
}

type SearchValue struct {
	Account int
	Name    string
}
