package searcher

import (
	"fmt"
	"reflect"
	"strings"
)

// Maximum matches to look for in each country
const maxMatches = 3

// CountryList is the array of country struct
type CountryList []Country

// Country is the the ultimate Struct that json file loads into
type Country struct {
	Name           string    `json:"name"`
	TopLevelDomain []string  `json:"topLevelDomain"`
	Alpha2Code     string    `json:"alpha2Code"`
	Alpha3Code     string    `json:"alpha3Code"`
	CallingCodes   []string  `json:"callingCodes"`
	Capital        string    `json:"capital"`
	AltSpellings   []string  `json:"altSpellings"`
	Region         string    `json:"region"`
	Subregion      string    `json:"subregion"`
	Population     int       `json:"population"`
	Latlng         []float64 `json:"latlng"`
	Demonym        string    `json:"demonym"`
	Area           float64   `json:"area"`
	Gini           float64   `json:"gini"`
	Timezones      []string  `json:"timezones"`
	Borders        []string  `json:"borders"`
	NativeName     string    `json:"nativeName"`
	NumericCode    string    `json:"numericCode"`
	Currencies     []struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Languages []struct {
		Iso6391    string `json:"iso639_1"`
		Iso6392    string `json:"iso639_2"`
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	} `json:"languages"`
	Translations struct {
		De string `json:"de"`
		Es string `json:"es"`
		Fr string `json:"fr"`
		Ja string `json:"ja"`
		It string `json:"it"`
		Br string `json:"br"`
		Pt string `json:"pt"`
		Nl string `json:"nl"`
		Hr string `json:"hr"`
		Fa string `json:"fa"`
	} `json:"translations"`
	Flag          string `json:"flag"`
	RegionalBlocs []struct {
		Acronym       string        `json:"acronym"`
		Name          string        `json:"name"`
		OtherAcronyms []interface{} `json:"otherAcronyms"`
		OtherNames    []interface{} `json:"otherNames"`
	} `json:"regionalBlocs"`
	Cioc string `json:"cioc"`
}

// ListItem single struct to create LisItemArray struct
type ListItem struct {
	Name  string        `json:"name"`
	Code  string        `json:"code"`
	Match []MatchSingle `json:"match"`
}

// MatchSingle holds single matches key val
type MatchSingle struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

// ListItemArray is the array of ListItem struct that the frontend is expecting at /search url
type ListItemArray []ListItem

// Search is a method with pointer receiver that searches for a specific search string
func (d *CountryList) Search(searchText string) (ListItemArray, error) {
	var listArray ListItemArray
	// Looping through each country in the CountryList (Array of struct)
	for _, element := range *d {
		var matches []MatchSingle
		// Don't do seach if the search string's length is less than 2. Display all the countries
		if len(searchText) >= 2 {
			matches = structDecoder(matches, element, searchText, "")
		}
		// If the lenth of searchText is less than 2. we still need name and code of country.
		if len(matches) > 0 || len(searchText) <= 1 {
			listArray = append(listArray, ListItem{element.Name, element.Alpha2Code, matches})

		}
	}
	return listArray, nil
}

// Loop through the element of a struct. In our case all elements of single country struct. i.e Country primarily.
// This method also has a recursive call that happens if the element of struct is a stuct
func structDecoder(matches []MatchSingle, structFields interface{}, searchText string, recTracker string) []MatchSingle {

	t := reflect.TypeOf(structFields)
	v := reflect.ValueOf(structFields)

	for i := 0; i < v.NumField(); i++ {
		if len(matches) >= maxMatches {
			return matches
		}
		switch v.Field(i).Kind() {
		case reflect.String:
			if ok := containsI(v.Field(i).Interface().(string), searchText); ok {
				matches = append(matches, MatchSingle{recTracker + "-" + t.Field(i).Name, v.Field(i).Interface().(string)})
			}
		case reflect.Float64:
			val := fmt.Sprintf("%f", v.Field(i).Interface().(float64))
			if ok := containsI(val, searchText); ok {
				matches = append(matches, MatchSingle{recTracker + "-" + t.Field(i).Name, val})
			}

		case reflect.Int:
			val := fmt.Sprintf("%d", v.Field(i).Interface().(int))
			if ok := containsI(val, searchText); ok {
				matches = append(matches, MatchSingle{recTracker + "-" + t.Field(i).Name, val})
			}
		case reflect.Struct:
			// Recursive call to the same function if the element of parent struct is also struct
			matches = structDecoder(matches, v.Field(i).Interface(), searchText, t.Field(i).Name)
		case reflect.Slice:
			s := v.Field(i)

			for j := 0; j < s.Len(); j++ {
				currentVal := s.Index(j)
				// If the element of struct is a slice of structs. Then we do recursive call for each struct.
				// Else we consider it as a slice of (int float64 or string)
				if currentVal.Kind() == reflect.Struct {
					matches = structDecoder(matches, currentVal.Interface(), searchText, t.Field(i).Name)
				} else {
					matches = sliceCase(matches, currentVal.Interface(), searchText, t.Field(i).Name)
				}
			}

		default:
			// Do nothing for not
		}

	}
	return matches
}

// checks if a contains b. case insensitive
func containsI(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}

// Check the type of the data and append accordingly.
func sliceCase(matches []MatchSingle, structFields interface{}, searchText string, recTracker string) []MatchSingle {
	v := reflect.ValueOf(structFields)
	if len(matches) >= maxMatches {
		return matches
	}
	switch v.Kind() {
	case reflect.String:
		if ok := containsI(v.Interface().(string), searchText); ok {
			matches = append(matches, MatchSingle{recTracker, v.Interface().(string)})
		}
	case reflect.Float64:
		val := fmt.Sprintf("%f", v.Interface().(float64))
		if ok := containsI(val, searchText); ok {
			matches = append(matches, MatchSingle{recTracker, val})
		}

	case reflect.Int:
		val := fmt.Sprintf("%d", v.Interface().(int))
		if ok := containsI(val, searchText); ok {
			matches = append(matches, MatchSingle{recTracker, val})
		}

	default:
		// Do nothing for not
	}

	return matches
}
