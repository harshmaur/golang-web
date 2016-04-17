package utils

import (
	"encoding/json"
	"fmt"
)

// Model is
type Model struct {
	Pics []string
}

// GetModel returns Model
func GetModel(ckValue string) Model {

	xs := DecodeSplit(ckValue)
	uPics := xs[1]
	var m Model
	err := json.Unmarshal([]byte(uPics), &m)
	if err != nil {
		fmt.Println("err unmarshal", err)
	}
	return m
}
