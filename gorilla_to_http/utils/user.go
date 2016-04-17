package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nu7hatch/gouuid"
)

// NewVisitor returns Cookie if the visitor is new
func NewVisitor() *http.Cookie {
	data := InitialModel()
	code := GetCode(data)
	id, _ := uuid.NewV4()
	return CreateCookie(id.String(), data, code)
}

// InitialModel set first model
func InitialModel() string {
	m := Model{[]string{
		"01.jpg",
	}}
	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err", err)
	}
	return string(bs)
}

// AddCookie creates a model and returns the encoded value of the new cookie
func AddCookie(ckValue string, fname string) *http.Cookie {
	// add pic to cookie
	m := GetModel(ckValue)
	m.Pics = append(m.Pics, fname)
	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err", err)
	}

	xs := DecodeSplit(ckValue) // get cookie values in slice
	return CreateCookie(xs[0], string(bs), xs[2])

}

// CreateCookie creates a cookie
func CreateCookie(id, data, code string) *http.Cookie {
	val := base64.StdEncoding.EncodeToString([]byte(id + "|" + data + "|" + code))
	ck := &http.Cookie{
		Name:     "session-id",
		Value:    val,
		HttpOnly: true,
	}
	return ck
}
