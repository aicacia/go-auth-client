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

// checks if the TokenRequestOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRequestOneOf{}

// TokenRequestOneOf struct for TokenRequestOneOf
type TokenRequestOneOf struct {
	GrantType string `json:"grant_type"`
	Password string `json:"password"`
	Scope NullableString `json:"scope,omitempty"`
	Username string `json:"username"`
}

type _TokenRequestOneOf TokenRequestOneOf

// NewTokenRequestOneOf instantiates a new TokenRequestOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRequestOneOf(grantType string, password string, username string) *TokenRequestOneOf {
	this := TokenRequestOneOf{}
	this.GrantType = grantType
	this.Password = password
	this.Username = username
	return &this
}

// NewTokenRequestOneOfWithDefaults instantiates a new TokenRequestOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRequestOneOfWithDefaults() *TokenRequestOneOf {
	this := TokenRequestOneOf{}
	return &this
}

// GetGrantType returns the GrantType field value
func (o *TokenRequestOneOf) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *TokenRequestOneOf) GetGrantTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *TokenRequestOneOf) SetGrantType(v string) {
	o.GrantType = v
}

// GetPassword returns the Password field value
func (o *TokenRequestOneOf) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *TokenRequestOneOf) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *TokenRequestOneOf) SetPassword(v string) {
	o.Password = v
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenRequestOneOf) GetScope() string {
	if o == nil || IsNil(o.Scope.Get()) {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenRequestOneOf) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *TokenRequestOneOf) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *TokenRequestOneOf) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *TokenRequestOneOf) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *TokenRequestOneOf) UnsetScope() {
	o.Scope.Unset()
}

// GetUsername returns the Username field value
func (o *TokenRequestOneOf) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *TokenRequestOneOf) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *TokenRequestOneOf) SetUsername(v string) {
	o.Username = v
}

func (o TokenRequestOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRequestOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grant_type"] = o.GrantType
	toSerialize["password"] = o.Password
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *TokenRequestOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grant_type",
		"password",
		"username",
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

	varTokenRequestOneOf := _TokenRequestOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenRequestOneOf)

	if err != nil {
		return err
	}

	*o = TokenRequestOneOf(varTokenRequestOneOf)

	return err
}

type NullableTokenRequestOneOf struct {
	value *TokenRequestOneOf
	isSet bool
}

func (v NullableTokenRequestOneOf) Get() *TokenRequestOneOf {
	return v.value
}

func (v *NullableTokenRequestOneOf) Set(val *TokenRequestOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRequestOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRequestOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRequestOneOf(val *TokenRequestOneOf) *NullableTokenRequestOneOf {
	return &NullableTokenRequestOneOf{value: val, isSet: true}
}

func (v NullableTokenRequestOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRequestOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


