/*
auth

auth

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateUsername type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUsername{}

// UpdateUsername struct for UpdateUsername
type UpdateUsername struct {
	Username NullableString `json:"username,omitempty"`
}

// NewUpdateUsername instantiates a new UpdateUsername object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUsername() *UpdateUsername {
	this := UpdateUsername{}
	return &this
}

// NewUpdateUsernameWithDefaults instantiates a new UpdateUsername object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUsernameWithDefaults() *UpdateUsername {
	this := UpdateUsername{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUsername) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUsername) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateUsername) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *UpdateUsername) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *UpdateUsername) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *UpdateUsername) UnsetUsername() {
	o.Username.Unset()
}

func (o UpdateUsername) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUsername) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	return toSerialize, nil
}

type NullableUpdateUsername struct {
	value *UpdateUsername
	isSet bool
}

func (v NullableUpdateUsername) Get() *UpdateUsername {
	return v.value
}

func (v *NullableUpdateUsername) Set(val *UpdateUsername) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUsername) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUsername) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUsername(val *UpdateUsername) *NullableUpdateUsername {
	return &NullableUpdateUsername{value: val, isSet: true}
}

func (v NullableUpdateUsername) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUsername) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


