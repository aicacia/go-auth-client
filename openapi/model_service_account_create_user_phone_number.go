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

// checks if the ServiceAccountCreateUserPhoneNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountCreateUserPhoneNumber{}

// ServiceAccountCreateUserPhoneNumber struct for ServiceAccountCreateUserPhoneNumber
type ServiceAccountCreateUserPhoneNumber struct {
	PhoneNumber string `json:"phone_number"`
	Primary NullableBool `json:"primary,omitempty"`
	Verified NullableBool `json:"verified,omitempty"`
}

type _ServiceAccountCreateUserPhoneNumber ServiceAccountCreateUserPhoneNumber

// NewServiceAccountCreateUserPhoneNumber instantiates a new ServiceAccountCreateUserPhoneNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountCreateUserPhoneNumber(phoneNumber string) *ServiceAccountCreateUserPhoneNumber {
	this := ServiceAccountCreateUserPhoneNumber{}
	this.PhoneNumber = phoneNumber
	return &this
}

// NewServiceAccountCreateUserPhoneNumberWithDefaults instantiates a new ServiceAccountCreateUserPhoneNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountCreateUserPhoneNumberWithDefaults() *ServiceAccountCreateUserPhoneNumber {
	this := ServiceAccountCreateUserPhoneNumber{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *ServiceAccountCreateUserPhoneNumber) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountCreateUserPhoneNumber) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *ServiceAccountCreateUserPhoneNumber) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetPrimary returns the Primary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountCreateUserPhoneNumber) GetPrimary() bool {
	if o == nil || IsNil(o.Primary.Get()) {
		var ret bool
		return ret
	}
	return *o.Primary.Get()
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountCreateUserPhoneNumber) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Primary.Get(), o.Primary.IsSet()
}

// HasPrimary returns a boolean if a field has been set.
func (o *ServiceAccountCreateUserPhoneNumber) HasPrimary() bool {
	if o != nil && o.Primary.IsSet() {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given NullableBool and assigns it to the Primary field.
func (o *ServiceAccountCreateUserPhoneNumber) SetPrimary(v bool) {
	o.Primary.Set(&v)
}
// SetPrimaryNil sets the value for Primary to be an explicit nil
func (o *ServiceAccountCreateUserPhoneNumber) SetPrimaryNil() {
	o.Primary.Set(nil)
}

// UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
func (o *ServiceAccountCreateUserPhoneNumber) UnsetPrimary() {
	o.Primary.Unset()
}

// GetVerified returns the Verified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountCreateUserPhoneNumber) GetVerified() bool {
	if o == nil || IsNil(o.Verified.Get()) {
		var ret bool
		return ret
	}
	return *o.Verified.Get()
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountCreateUserPhoneNumber) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verified.Get(), o.Verified.IsSet()
}

// HasVerified returns a boolean if a field has been set.
func (o *ServiceAccountCreateUserPhoneNumber) HasVerified() bool {
	if o != nil && o.Verified.IsSet() {
		return true
	}

	return false
}

// SetVerified gets a reference to the given NullableBool and assigns it to the Verified field.
func (o *ServiceAccountCreateUserPhoneNumber) SetVerified(v bool) {
	o.Verified.Set(&v)
}
// SetVerifiedNil sets the value for Verified to be an explicit nil
func (o *ServiceAccountCreateUserPhoneNumber) SetVerifiedNil() {
	o.Verified.Set(nil)
}

// UnsetVerified ensures that no value is present for Verified, not even an explicit nil
func (o *ServiceAccountCreateUserPhoneNumber) UnsetVerified() {
	o.Verified.Unset()
}

func (o ServiceAccountCreateUserPhoneNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountCreateUserPhoneNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone_number"] = o.PhoneNumber
	if o.Primary.IsSet() {
		toSerialize["primary"] = o.Primary.Get()
	}
	if o.Verified.IsSet() {
		toSerialize["verified"] = o.Verified.Get()
	}
	return toSerialize, nil
}

func (o *ServiceAccountCreateUserPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone_number",
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

	varServiceAccountCreateUserPhoneNumber := _ServiceAccountCreateUserPhoneNumber{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceAccountCreateUserPhoneNumber)

	if err != nil {
		return err
	}

	*o = ServiceAccountCreateUserPhoneNumber(varServiceAccountCreateUserPhoneNumber)

	return err
}

type NullableServiceAccountCreateUserPhoneNumber struct {
	value *ServiceAccountCreateUserPhoneNumber
	isSet bool
}

func (v NullableServiceAccountCreateUserPhoneNumber) Get() *ServiceAccountCreateUserPhoneNumber {
	return v.value
}

func (v *NullableServiceAccountCreateUserPhoneNumber) Set(val *ServiceAccountCreateUserPhoneNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountCreateUserPhoneNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountCreateUserPhoneNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountCreateUserPhoneNumber(val *ServiceAccountCreateUserPhoneNumber) *NullableServiceAccountCreateUserPhoneNumber {
	return &NullableServiceAccountCreateUserPhoneNumber{value: val, isSet: true}
}

func (v NullableServiceAccountCreateUserPhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountCreateUserPhoneNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


