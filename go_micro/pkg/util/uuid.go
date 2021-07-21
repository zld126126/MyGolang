package util

import (
	uuid "github.com/satori/go.uuid"
)

func CreateUUID() string {
	// Creating UUID Version 4
	u1 := uuid.NewV4().String()
	return u1
	//fmt.Println(u1)
	// Parsing UUID from string input
	//u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	//if err != nil {
	//	fmt.Printf("Something went wrong: %s", err)
	//}
	//fmt.Printf("Successfully parsed: %s", u2)
}
