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

// checks if the Token type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Token{}

// Token struct for Token
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	IdToken NullableString `json:"id_token,omitempty"`
	IssuedTokenType NullableString `json:"issued_token_type,omitempty"`
	RefreshToken NullableString `json:"refresh_token,omitempty"`
	RefreshTokenExpiresIn NullableInt64 `json:"refresh_token_expires_in,omitempty"`
	Scope NullableString `json:"scope,omitempty"`
	TokenType string `json:"token_type"`
}

type _Token Token

// NewToken instantiates a new Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToken(accessToken string, expiresIn int64, tokenType string) *Token {
	this := Token{}
	this.AccessToken = accessToken
	this.ExpiresIn = expiresIn
	this.TokenType = tokenType
	return &this
}

// NewTokenWithDefaults instantiates a new Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWithDefaults() *Token {
	this := Token{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *Token) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *Token) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *Token) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *Token) GetExpiresIn() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *Token) GetExpiresInOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *Token) SetExpiresIn(v int64) {
	o.ExpiresIn = v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Token) GetIdToken() string {
	if o == nil || IsNil(o.IdToken.Get()) {
		var ret string
		return ret
	}
	return *o.IdToken.Get()
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetIdTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdToken.Get(), o.IdToken.IsSet()
}

// HasIdToken returns a boolean if a field has been set.
func (o *Token) HasIdToken() bool {
	if o != nil && o.IdToken.IsSet() {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given NullableString and assigns it to the IdToken field.
func (o *Token) SetIdToken(v string) {
	o.IdToken.Set(&v)
}
// SetIdTokenNil sets the value for IdToken to be an explicit nil
func (o *Token) SetIdTokenNil() {
	o.IdToken.Set(nil)
}

// UnsetIdToken ensures that no value is present for IdToken, not even an explicit nil
func (o *Token) UnsetIdToken() {
	o.IdToken.Unset()
}

// GetIssuedTokenType returns the IssuedTokenType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Token) GetIssuedTokenType() string {
	if o == nil || IsNil(o.IssuedTokenType.Get()) {
		var ret string
		return ret
	}
	return *o.IssuedTokenType.Get()
}

// GetIssuedTokenTypeOk returns a tuple with the IssuedTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetIssuedTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuedTokenType.Get(), o.IssuedTokenType.IsSet()
}

// HasIssuedTokenType returns a boolean if a field has been set.
func (o *Token) HasIssuedTokenType() bool {
	if o != nil && o.IssuedTokenType.IsSet() {
		return true
	}

	return false
}

// SetIssuedTokenType gets a reference to the given NullableString and assigns it to the IssuedTokenType field.
func (o *Token) SetIssuedTokenType(v string) {
	o.IssuedTokenType.Set(&v)
}
// SetIssuedTokenTypeNil sets the value for IssuedTokenType to be an explicit nil
func (o *Token) SetIssuedTokenTypeNil() {
	o.IssuedTokenType.Set(nil)
}

// UnsetIssuedTokenType ensures that no value is present for IssuedTokenType, not even an explicit nil
func (o *Token) UnsetIssuedTokenType() {
	o.IssuedTokenType.Unset()
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Token) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *Token) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *Token) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *Token) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *Token) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

// GetRefreshTokenExpiresIn returns the RefreshTokenExpiresIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Token) GetRefreshTokenExpiresIn() int64 {
	if o == nil || IsNil(o.RefreshTokenExpiresIn.Get()) {
		var ret int64
		return ret
	}
	return *o.RefreshTokenExpiresIn.Get()
}

// GetRefreshTokenExpiresInOk returns a tuple with the RefreshTokenExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetRefreshTokenExpiresInOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshTokenExpiresIn.Get(), o.RefreshTokenExpiresIn.IsSet()
}

// HasRefreshTokenExpiresIn returns a boolean if a field has been set.
func (o *Token) HasRefreshTokenExpiresIn() bool {
	if o != nil && o.RefreshTokenExpiresIn.IsSet() {
		return true
	}

	return false
}

// SetRefreshTokenExpiresIn gets a reference to the given NullableInt64 and assigns it to the RefreshTokenExpiresIn field.
func (o *Token) SetRefreshTokenExpiresIn(v int64) {
	o.RefreshTokenExpiresIn.Set(&v)
}
// SetRefreshTokenExpiresInNil sets the value for RefreshTokenExpiresIn to be an explicit nil
func (o *Token) SetRefreshTokenExpiresInNil() {
	o.RefreshTokenExpiresIn.Set(nil)
}

// UnsetRefreshTokenExpiresIn ensures that no value is present for RefreshTokenExpiresIn, not even an explicit nil
func (o *Token) UnsetRefreshTokenExpiresIn() {
	o.RefreshTokenExpiresIn.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Token) GetScope() string {
	if o == nil || IsNil(o.Scope.Get()) {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *Token) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *Token) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *Token) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *Token) UnsetScope() {
	o.Scope.Unset()
}

// GetTokenType returns the TokenType field value
func (o *Token) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *Token) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *Token) SetTokenType(v string) {
	o.TokenType = v
}

func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Token) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["expires_in"] = o.ExpiresIn
	if o.IdToken.IsSet() {
		toSerialize["id_token"] = o.IdToken.Get()
	}
	if o.IssuedTokenType.IsSet() {
		toSerialize["issued_token_type"] = o.IssuedTokenType.Get()
	}
	if o.RefreshToken.IsSet() {
		toSerialize["refresh_token"] = o.RefreshToken.Get()
	}
	if o.RefreshTokenExpiresIn.IsSet() {
		toSerialize["refresh_token_expires_in"] = o.RefreshTokenExpiresIn.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	toSerialize["token_type"] = o.TokenType
	return toSerialize, nil
}

func (o *Token) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
		"expires_in",
		"token_type",
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

	varToken := _Token{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToken)

	if err != nil {
		return err
	}

	*o = Token(varToken)

	return err
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

