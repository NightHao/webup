package gsheet

type ValueRange struct {
	Values []Value `json:"values"`
}

type Value []string
