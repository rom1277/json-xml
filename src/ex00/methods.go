package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"` // переименовываем Recipes в recipes
	Cakes   []Recipe `json:"cake" xml:"cake"`
}

type Recipe struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

type DBReader interface {
	Reader(filiname string) (*Recipes, error)
	Writer(recipes Recipes) (string, error)
}

type JSON struct{}
type XML struct{}

func (r *JSON) Reader(filename string) (*Recipes, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var recipes Recipes
	if err = json.Unmarshal(text, &recipes); err != nil {
		return nil, err
	}
	return &recipes, nil
}

func (r *XML) Reader(filename string) (*Recipes, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var recipes Recipes
	if err = xml.Unmarshal(text, &recipes); err != nil {
		return nil, err
	}
	return &recipes, nil
}

func (w *JSON) Writer(recipes Recipes) (string, error) {

	data, err := xml.MarshalIndent(recipes, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func (w *XML) Writer(recipes Recipes) (string, error) {
	data, err := json.MarshalIndent(recipes, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func GetTypeDB(filename string) (DBReader, error) {
	if strings.HasSuffix(filename, ".json") {
		return &JSON{}, nil
	} else if strings.HasSuffix(filename, ".xml") {
		return &XML{}, nil
	} else {
		return nil, fmt.Errorf("unsupported file extension")
	}
}
