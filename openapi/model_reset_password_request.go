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

// checks if the ResetPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetPasswordRequest{}

// ResetPasswordRequest struct for ResetPasswordRequest
type ResetPasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	Password string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type _ResetPasswordRequest ResetPasswordRequest

// NewResetPasswordRequest instantiates a new ResetPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetPasswordRequest(currentPassword string, password string, passwordConfirmation string) *ResetPasswordRequest {
	this := ResetPasswordRequest{}
	this.CurrentPassword = currentPassword
	this.Password = password
	this.PasswordConfirmation = passwordConfirmation
	return &this
}

// NewResetPasswordRequestWithDefaults instantiates a new ResetPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetPasswordRequestWithDefaults() *ResetPasswordRequest {
	this := ResetPasswordRequest{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value
func (o *ResetPasswordRequest) GetCurrentPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value
// and a boolean to check if the value has been set.
func (o *ResetPasswordRequest) GetCurrentPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPassword, true
}

// SetCurrentPassword sets field value
func (o *ResetPasswordRequest) SetCurrentPassword(v string) {
	o.CurrentPassword = v
}

// GetPassword returns the Password field value
func (o *ResetPasswordRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ResetPasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ResetPasswordRequest) SetPassword(v string) {
	o.Password = v
}

// GetPasswordConfirmation returns the PasswordConfirmation field value
func (o *ResetPasswordRequest) GetPasswordConfirmation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordConfirmation
}

// GetPasswordConfirmationOk returns a tuple with the PasswordConfirmation field value
// and a boolean to check if the value has been set.
func (o *ResetPasswordRequest) GetPasswordConfirmationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordConfirmation, true
}

// SetPasswordConfirmation sets field value
func (o *ResetPasswordRequest) SetPasswordConfirmation(v string) {
	o.PasswordConfirmation = v
}

func (o ResetPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetPasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current_password"] = o.CurrentPassword
	toSerialize["password"] = o.Password
	toSerialize["password_confirmation"] = o.PasswordConfirmation
	return toSerialize, nil
}

func (o *ResetPasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"current_password",
		"password",
		"password_confirmation",
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

	varResetPasswordRequest := _ResetPasswordRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResetPasswordRequest)

	if err != nil {
		return err
	}

	*o = ResetPasswordRequest(varResetPasswordRequest)

	return err
}

type NullableResetPasswordRequest struct {
	value *ResetPasswordRequest
	isSet bool
}

func (v NullableResetPasswordRequest) Get() *ResetPasswordRequest {
	return v.value
}

func (v *NullableResetPasswordRequest) Set(val *ResetPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResetPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResetPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetPasswordRequest(val *ResetPasswordRequest) *NullableResetPasswordRequest {
	return &NullableResetPasswordRequest{value: val, isSet: true}
}

func (v NullableResetPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

