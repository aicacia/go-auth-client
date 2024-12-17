/*
auth

auth

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// Algorithm jsonwebtoken::Algorithm
type Algorithm string

// List of Algorithm
const (
	HS256 Algorithm = "HS256"
	HS384 Algorithm = "HS384"
	HS512 Algorithm = "HS512"
	ES256 Algorithm = "ES256"
	ES384 Algorithm = "ES384"
	RS256 Algorithm = "RS256"
	RS384 Algorithm = "RS384"
	RS512 Algorithm = "RS512"
	PS256 Algorithm = "PS256"
	PS384 Algorithm = "PS384"
	PS512 Algorithm = "PS512"
	ED_DSA Algorithm = "EdDSA"
)

// All allowed values of Algorithm enum
var AllowedAlgorithmEnumValues = []Algorithm{
	"HS256",
	"HS384",
	"HS512",
	"ES256",
	"ES384",
	"RS256",
	"RS384",
	"RS512",
	"PS256",
	"PS384",
	"PS512",
	"EdDSA",
}

func (v *Algorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Algorithm(value)
	for _, existing := range AllowedAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Algorithm", value)
}

// NewAlgorithmFromValue returns a pointer to a valid Algorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlgorithmFromValue(v string) (*Algorithm, error) {
	ev := Algorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Algorithm: valid values are %v", v, AllowedAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Algorithm) IsValid() bool {
	for _, existing := range AllowedAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Algorithm value
func (v Algorithm) Ptr() *Algorithm {
	return &v
}

type NullableAlgorithm struct {
	value *Algorithm
	isSet bool
}

func (v NullableAlgorithm) Get() *Algorithm {
	return v.value
}

func (v *NullableAlgorithm) Set(val *Algorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgorithm(val *Algorithm) *NullableAlgorithm {
	return &NullableAlgorithm{value: val, isSet: true}
}

func (v NullableAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

