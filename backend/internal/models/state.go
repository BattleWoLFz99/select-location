package models

type State struct {
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
}
