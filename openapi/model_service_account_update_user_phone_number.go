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

// checks if the ServiceAccountUpdateUserPhoneNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountUpdateUserPhoneNumber{}

// ServiceAccountUpdateUserPhoneNumber struct for ServiceAccountUpdateUserPhoneNumber
type ServiceAccountUpdateUserPhoneNumber struct {
	Primary NullableBool `json:"primary,omitempty"`
	Verified NullableBool `json:"verified,omitempty"`
}

// NewServiceAccountUpdateUserPhoneNumber instantiates a new ServiceAccountUpdateUserPhoneNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountUpdateUserPhoneNumber() *ServiceAccountUpdateUserPhoneNumber {
	this := ServiceAccountUpdateUserPhoneNumber{}
	return &this
}

// NewServiceAccountUpdateUserPhoneNumberWithDefaults instantiates a new ServiceAccountUpdateUserPhoneNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountUpdateUserPhoneNumberWithDefaults() *ServiceAccountUpdateUserPhoneNumber {
	this := ServiceAccountUpdateUserPhoneNumber{}
	return &this
}

// GetPrimary returns the Primary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountUpdateUserPhoneNumber) GetPrimary() bool {
	if o == nil || IsNil(o.Primary.Get()) {
		var ret bool
		return ret
	}
	return *o.Primary.Get()
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountUpdateUserPhoneNumber) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Primary.Get(), o.Primary.IsSet()
}

// HasPrimary returns a boolean if a field has been set.
func (o *ServiceAccountUpdateUserPhoneNumber) HasPrimary() bool {
	if o != nil && o.Primary.IsSet() {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given NullableBool and assigns it to the Primary field.
func (o *ServiceAccountUpdateUserPhoneNumber) SetPrimary(v bool) {
	o.Primary.Set(&v)
}
// SetPrimaryNil sets the value for Primary to be an explicit nil
func (o *ServiceAccountUpdateUserPhoneNumber) SetPrimaryNil() {
	o.Primary.Set(nil)
}

// UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
func (o *ServiceAccountUpdateUserPhoneNumber) UnsetPrimary() {
	o.Primary.Unset()
}

// GetVerified returns the Verified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountUpdateUserPhoneNumber) GetVerified() bool {
	if o == nil || IsNil(o.Verified.Get()) {
		var ret bool
		return ret
	}
	return *o.Verified.Get()
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountUpdateUserPhoneNumber) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verified.Get(), o.Verified.IsSet()
}

// HasVerified returns a boolean if a field has been set.
func (o *ServiceAccountUpdateUserPhoneNumber) HasVerified() bool {
	if o != nil && o.Verified.IsSet() {
		return true
	}

	return false
}

// SetVerified gets a reference to the given NullableBool and assigns it to the Verified field.
func (o *ServiceAccountUpdateUserPhoneNumber) SetVerified(v bool) {
	o.Verified.Set(&v)
}
// SetVerifiedNil sets the value for Verified to be an explicit nil
func (o *ServiceAccountUpdateUserPhoneNumber) SetVerifiedNil() {
	o.Verified.Set(nil)
}

// UnsetVerified ensures that no value is present for Verified, not even an explicit nil
func (o *ServiceAccountUpdateUserPhoneNumber) UnsetVerified() {
	o.Verified.Unset()
}

func (o ServiceAccountUpdateUserPhoneNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountUpdateUserPhoneNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Primary.IsSet() {
		toSerialize["primary"] = o.Primary.Get()
	}
	if o.Verified.IsSet() {
		toSerialize["verified"] = o.Verified.Get()
	}
	return toSerialize, nil
}

type NullableServiceAccountUpdateUserPhoneNumber struct {
	value *ServiceAccountUpdateUserPhoneNumber
	isSet bool
}

func (v NullableServiceAccountUpdateUserPhoneNumber) Get() *ServiceAccountUpdateUserPhoneNumber {
	return v.value
}

func (v *NullableServiceAccountUpdateUserPhoneNumber) Set(val *ServiceAccountUpdateUserPhoneNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountUpdateUserPhoneNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountUpdateUserPhoneNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountUpdateUserPhoneNumber(val *ServiceAccountUpdateUserPhoneNumber) *NullableServiceAccountUpdateUserPhoneNumber {
	return &NullableServiceAccountUpdateUserPhoneNumber{value: val, isSet: true}
}

func (v NullableServiceAccountUpdateUserPhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountUpdateUserPhoneNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


