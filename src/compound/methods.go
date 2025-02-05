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
func Сompare(menu1, menu2 *Recipes) {
	addedcakes, removedcakes, matchescakes := ComparingNames(menu1.Cakes, menu2.Cakes)
	for _, cake := range addedcakes {
		fmt.Printf("ADDED cake \"%s\"\n", cake)
	}
	for _, cake := range removedcakes {
		fmt.Printf("REMOVED cake \"%s\"\n", cake)
	}
	var time1, time2 string
	var ingreds1, ingreds2 []Ingredient
	i := 0
	for i < len(matchescakes) {
		for _, cakes := range menu1.Cakes {
			if cakes.Name == matchescakes[i] {
				time1 = cakes.Time
				ingreds1 = cakes.Ingredients
				break
			}
		}
		for _, cakes := range menu2.Cakes {
			if cakes.Name == matchescakes[i] {
				time2 = cakes.Time
				ingreds2 = cakes.Ingredients
				break
			}
		}

		if time1 != time2 {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", matchescakes[i], time2, time1)
		}
		fmt.Print(ComparingIngredient(matchescakes[i], ingreds1, ingreds2))
		i++
		if i != len(matchescakes) {
			fmt.Println()
		}
	}
}

// сравнение игнредиентов одних и тех же тортов
func ComparingIngredient(cake string, ingredients1, ingredients2 []Ingredient) string {

	igr1 := make(map[string]Ingredient, len(ingredients1))
	igr2 := make(map[string]Ingredient, len(ingredients2))
	ingrhave := make(map[string]int)
	var str string
	var result strings.Builder
	for _, val := range ingredients1 {
		igr1[val.Name] = val
		ingrhave[val.Name] = 1
	}

	for _, val := range ingredients2 {
		igr2[val.Name] = val
		ingrhave[val.Name] += 2
	}
	for name, val := range ingrhave {
		if val == 1 {
			str = fmt.Sprintf("REMOVED ingredient \"%s\" for cake \"%s\"\n", name, cake)
			result.WriteString(str)
		}
		if val == 2 {
			str = fmt.Sprintf("ADDED ingredient \"%s\" for cake \"%s\"\n", name, cake)
			result.WriteString(str)
		}

	}
	for name := range igr1 {
		//ИЗМЕНЕНА единица измерения ингредиента
		if igr1[name].Unit != igr2[name].Unit && igr2[name].Unit != "" {
			str = fmt.Sprintf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, cake, igr2[name].Unit, igr1[name].Unit)
			result.WriteString(str)
		}
		//ИЗМЕНЕНО количество единиц измерения
		if igr1[name].Count != igr2[name].Count && igr2[name].Count != "" {
			str = fmt.Sprintf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, cake, igr2[name].Count, igr1[name].Count)
			result.WriteString(str)
		}
		//УДАЛЕНО единица измерения
		if igr1[name].Unit != igr2[name].Unit && igr2[name].Unit == "" && igr2[name].Name != "" {
			str = fmt.Sprintf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", igr1[name].Unit, name, cake)
			result.WriteString(str)
		}
	}
	return result.String()
}

func ComparingNames(recipes1, recipes2 []Recipe) (added, removed, matches []string) {
	set1 := make(map[string]bool, len(recipes1))
	set2 := make(map[string]bool, len(recipes2))

	for _, item := range recipes1 {
		set1[item.Name] = true
	}
	for _, item := range recipes2 {
		set2[item.Name] = true
	}
	// Находим добавленные элементы (есть в recipes2, но нет в recipes1)
	for item := range set2 {
		if !set1[item] {
			added = append(added, item)
		}
	}
	// Находим удалённые элементы (есть в recipes1, но нет в recipes2)
	for item := range set1 {
		if !set2[item] {
			removed = append(removed, item)
		}
	}
	for item := range set1 {
		if set1[item] == set2[item] {
			matches = append(matches, item)
		}
	}
	return added, removed, matches
}
