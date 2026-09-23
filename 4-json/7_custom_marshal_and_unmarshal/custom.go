package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
)

type Version struct {
	Major, Minor, Patch int64
}

func (v Version) MarshalJSONTo(enc *jsontext.Encoder) error {
	return json.MarshalEncode(enc, fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch))
}

func (v *Version) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	if k := dec.PeekKind(); k != jsontext.KindString {
		return &json.SemanticError{JSONKind: k}
	}

	var s string
	if err := json.UnmarshalDecode(dec, &s); err != nil {
		return err
	}

	_, err := fmt.Sscanf(s, "%d.%d.%d", &v.Major, &v.Minor, &v.Patch)
	return err
}

func main() {
	v := Version{1, 2, 3}

	encoded, err := json.Marshal(v)
	fmt.Printf("Marshaled: %s\nError: %v\n", encoded, err)

	var decoded Version
	err = json.Unmarshal(encoded, &decoded)

	fmt.Printf("Unmarshaled: %+v\nError: %v\n", decoded, err)
}
