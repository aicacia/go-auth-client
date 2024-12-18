/*
auth

auth

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PaginationTenent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationTenent{}

// PaginationTenent struct for PaginationTenent
type PaginationTenent struct {
	HasMore bool `json:"has_more"`
	Items []PaginationTenentItemsInner `json:"items"`
}

type _PaginationTenent PaginationTenent

// NewPaginationTenent instantiates a new PaginationTenent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationTenent(hasMore bool, items []PaginationTenentItemsInner) *PaginationTenent {
	this := PaginationTenent{}
	this.HasMore = hasMore
	this.Items = items
	return &this
}

// NewPaginationTenentWithDefaults instantiates a new PaginationTenent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationTenentWithDefaults() *PaginationTenent {
	this := PaginationTenent{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *PaginationTenent) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *PaginationTenent) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *PaginationTenent) SetHasMore(v bool) {
	o.HasMore = v
}

// GetItems returns the Items field value
func (o *PaginationTenent) GetItems() []PaginationTenentItemsInner {
	if o == nil {
		var ret []PaginationTenentItemsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PaginationTenent) GetItemsOk() ([]PaginationTenentItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PaginationTenent) SetItems(v []PaginationTenentItemsInner) {
	o.Items = v
}

func (o PaginationTenent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationTenent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_more"] = o.HasMore
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *PaginationTenent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"has_more",
		"items",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaginationTenent := _PaginationTenent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginationTenent)

	if err != nil {
		return err
	}

	*o = PaginationTenent(varPaginationTenent)

	return err
}

type NullablePaginationTenent struct {
	value *PaginationTenent
	isSet bool
}

func (v NullablePaginationTenent) Get() *PaginationTenent {
	return v.value
}

func (v *NullablePaginationTenent) Set(val *PaginationTenent) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationTenent) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationTenent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationTenent(val *PaginationTenent) *NullablePaginationTenent {
	return &NullablePaginationTenent{value: val, isSet: true}
}

func (v NullablePaginationTenent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationTenent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


