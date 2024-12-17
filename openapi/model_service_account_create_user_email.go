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

// checks if the ServiceAccountCreateUserEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountCreateUserEmail{}

// ServiceAccountCreateUserEmail struct for ServiceAccountCreateUserEmail
type ServiceAccountCreateUserEmail struct {
	Email string `json:"email"`
	Primary NullableBool `json:"primary,omitempty"`
	Verified NullableBool `json:"verified,omitempty"`
}

type _ServiceAccountCreateUserEmail ServiceAccountCreateUserEmail

// NewServiceAccountCreateUserEmail instantiates a new ServiceAccountCreateUserEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountCreateUserEmail(email string) *ServiceAccountCreateUserEmail {
	this := ServiceAccountCreateUserEmail{}
	this.Email = email
	return &this
}

// NewServiceAccountCreateUserEmailWithDefaults instantiates a new ServiceAccountCreateUserEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountCreateUserEmailWithDefaults() *ServiceAccountCreateUserEmail {
	this := ServiceAccountCreateUserEmail{}
	return &this
}

// GetEmail returns the Email field value
func (o *ServiceAccountCreateUserEmail) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountCreateUserEmail) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ServiceAccountCreateUserEmail) SetEmail(v string) {
	o.Email = v
}

// GetPrimary returns the Primary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountCreateUserEmail) GetPrimary() bool {
	if o == nil || IsNil(o.Primary.Get()) {
		var ret bool
		return ret
	}
	return *o.Primary.Get()
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountCreateUserEmail) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Primary.Get(), o.Primary.IsSet()
}

// HasPrimary returns a boolean if a field has been set.
func (o *ServiceAccountCreateUserEmail) HasPrimary() bool {
	if o != nil && o.Primary.IsSet() {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given NullableBool and assigns it to the Primary field.
func (o *ServiceAccountCreateUserEmail) SetPrimary(v bool) {
	o.Primary.Set(&v)
}
// SetPrimaryNil sets the value for Primary to be an explicit nil
func (o *ServiceAccountCreateUserEmail) SetPrimaryNil() {
	o.Primary.Set(nil)
}

// UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
func (o *ServiceAccountCreateUserEmail) UnsetPrimary() {
	o.Primary.Unset()
}

// GetVerified returns the Verified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountCreateUserEmail) GetVerified() bool {
	if o == nil || IsNil(o.Verified.Get()) {
		var ret bool
		return ret
	}
	return *o.Verified.Get()
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountCreateUserEmail) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verified.Get(), o.Verified.IsSet()
}

// HasVerified returns a boolean if a field has been set.
func (o *ServiceAccountCreateUserEmail) HasVerified() bool {
	if o != nil && o.Verified.IsSet() {
		return true
	}

	return false
}

// SetVerified gets a reference to the given NullableBool and assigns it to the Verified field.
func (o *ServiceAccountCreateUserEmail) SetVerified(v bool) {
	o.Verified.Set(&v)
}
// SetVerifiedNil sets the value for Verified to be an explicit nil
func (o *ServiceAccountCreateUserEmail) SetVerifiedNil() {
	o.Verified.Set(nil)
}

// UnsetVerified ensures that no value is present for Verified, not even an explicit nil
func (o *ServiceAccountCreateUserEmail) UnsetVerified() {
	o.Verified.Unset()
}

func (o ServiceAccountCreateUserEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountCreateUserEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if o.Primary.IsSet() {
		toSerialize["primary"] = o.Primary.Get()
	}
	if o.Verified.IsSet() {
		toSerialize["verified"] = o.Verified.Get()
	}
	return toSerialize, nil
}

func (o *ServiceAccountCreateUserEmail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varServiceAccountCreateUserEmail := _ServiceAccountCreateUserEmail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceAccountCreateUserEmail)

	if err != nil {
		return err
	}

	*o = ServiceAccountCreateUserEmail(varServiceAccountCreateUserEmail)

	return err
}

type NullableServiceAccountCreateUserEmail struct {
	value *ServiceAccountCreateUserEmail
	isSet bool
}

func (v NullableServiceAccountCreateUserEmail) Get() *ServiceAccountCreateUserEmail {
	return v.value
}

func (v *NullableServiceAccountCreateUserEmail) Set(val *ServiceAccountCreateUserEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountCreateUserEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountCreateUserEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountCreateUserEmail(val *ServiceAccountCreateUserEmail) *NullableServiceAccountCreateUserEmail {
	return &NullableServiceAccountCreateUserEmail{value: val, isSet: true}
}

func (v NullableServiceAccountCreateUserEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountCreateUserEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


