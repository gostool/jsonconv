package jsonconv

import "encoding/json"

type JsonCamelCase struct {
	Value interface{}
}

func (c JsonCamelCase) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			resKey := LowerFirst(Case2Camel(key))
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}
