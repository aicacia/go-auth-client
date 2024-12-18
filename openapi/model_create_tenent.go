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

// checks if the CreateTenent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTenent{}

// CreateTenent struct for CreateTenent
type CreateTenent struct {
	Algorithm NullableAlgorithm `json:"algorithm,omitempty"`
	Audience string `json:"audience"`
	ClientId NullableString `json:"client_id,omitempty"`
	ExpiresInSeconds NullableInt64 `json:"expires_in_seconds,omitempty"`
	Issuer string `json:"issuer"`
	PrivateKey string `json:"private_key"`
	PublicKey NullableString `json:"public_key,omitempty"`
	RefreshExpiresInSeconds NullableInt64 `json:"refresh_expires_in_seconds,omitempty"`
}

type _CreateTenent CreateTenent

// NewCreateTenent instantiates a new CreateTenent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTenent(audience string, issuer string, privateKey string) *CreateTenent {
	this := CreateTenent{}
	this.Audience = audience
	this.Issuer = issuer
	this.PrivateKey = privateKey
	return &this
}

// NewCreateTenentWithDefaults instantiates a new CreateTenent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTenentWithDefaults() *CreateTenent {
	this := CreateTenent{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTenent) GetAlgorithm() Algorithm {
	if o == nil || IsNil(o.Algorithm.Get()) {
		var ret Algorithm
		return ret
	}
	return *o.Algorithm.Get()
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTenent) GetAlgorithmOk() (*Algorithm, bool) {
	if o == nil {
		return nil, false
	}
	return o.Algorithm.Get(), o.Algorithm.IsSet()
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *CreateTenent) HasAlgorithm() bool {
	if o != nil && o.Algorithm.IsSet() {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given NullableAlgorithm and assigns it to the Algorithm field.
func (o *CreateTenent) SetAlgorithm(v Algorithm) {
	o.Algorithm.Set(&v)
}
// SetAlgorithmNil sets the value for Algorithm to be an explicit nil
func (o *CreateTenent) SetAlgorithmNil() {
	o.Algorithm.Set(nil)
}

// UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
func (o *CreateTenent) UnsetAlgorithm() {
	o.Algorithm.Unset()
}

// GetAudience returns the Audience field value
func (o *CreateTenent) GetAudience() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *CreateTenent) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *CreateTenent) SetAudience(v string) {
	o.Audience = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTenent) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTenent) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *CreateTenent) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *CreateTenent) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *CreateTenent) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *CreateTenent) UnsetClientId() {
	o.ClientId.Unset()
}

// GetExpiresInSeconds returns the ExpiresInSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTenent) GetExpiresInSeconds() int64 {
	if o == nil || IsNil(o.ExpiresInSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiresInSeconds.Get()
}

// GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTenent) GetExpiresInSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresInSeconds.Get(), o.ExpiresInSeconds.IsSet()
}

// HasExpiresInSeconds returns a boolean if a field has been set.
func (o *CreateTenent) HasExpiresInSeconds() bool {
	if o != nil && o.ExpiresInSeconds.IsSet() {
		return true
	}

	return false
}

// SetExpiresInSeconds gets a reference to the given NullableInt64 and assigns it to the ExpiresInSeconds field.
func (o *CreateTenent) SetExpiresInSeconds(v int64) {
	o.ExpiresInSeconds.Set(&v)
}
// SetExpiresInSecondsNil sets the value for ExpiresInSeconds to be an explicit nil
func (o *CreateTenent) SetExpiresInSecondsNil() {
	o.ExpiresInSeconds.Set(nil)
}

// UnsetExpiresInSeconds ensures that no value is present for ExpiresInSeconds, not even an explicit nil
func (o *CreateTenent) UnsetExpiresInSeconds() {
	o.ExpiresInSeconds.Unset()
}

// GetIssuer returns the Issuer field value
func (o *CreateTenent) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *CreateTenent) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *CreateTenent) SetIssuer(v string) {
	o.Issuer = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *CreateTenent) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *CreateTenent) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *CreateTenent) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTenent) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTenent) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *CreateTenent) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *CreateTenent) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *CreateTenent) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *CreateTenent) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetRefreshExpiresInSeconds returns the RefreshExpiresInSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTenent) GetRefreshExpiresInSeconds() int64 {
	if o == nil || IsNil(o.RefreshExpiresInSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.RefreshExpiresInSeconds.Get()
}

// GetRefreshExpiresInSecondsOk returns a tuple with the RefreshExpiresInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTenent) GetRefreshExpiresInSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshExpiresInSeconds.Get(), o.RefreshExpiresInSeconds.IsSet()
}

// HasRefreshExpiresInSeconds returns a boolean if a field has been set.
func (o *CreateTenent) HasRefreshExpiresInSeconds() bool {
	if o != nil && o.RefreshExpiresInSeconds.IsSet() {
		return true
	}

	return false
}

// SetRefreshExpiresInSeconds gets a reference to the given NullableInt64 and assigns it to the RefreshExpiresInSeconds field.
func (o *CreateTenent) SetRefreshExpiresInSeconds(v int64) {
	o.RefreshExpiresInSeconds.Set(&v)
}
// SetRefreshExpiresInSecondsNil sets the value for RefreshExpiresInSeconds to be an explicit nil
func (o *CreateTenent) SetRefreshExpiresInSecondsNil() {
	o.RefreshExpiresInSeconds.Set(nil)
}

// UnsetRefreshExpiresInSeconds ensures that no value is present for RefreshExpiresInSeconds, not even an explicit nil
func (o *CreateTenent) UnsetRefreshExpiresInSeconds() {
	o.RefreshExpiresInSeconds.Unset()
}

func (o CreateTenent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTenent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm.IsSet() {
		toSerialize["algorithm"] = o.Algorithm.Get()
	}
	toSerialize["audience"] = o.Audience
	if o.ClientId.IsSet() {
		toSerialize["client_id"] = o.ClientId.Get()
	}
	if o.ExpiresInSeconds.IsSet() {
		toSerialize["expires_in_seconds"] = o.ExpiresInSeconds.Get()
	}
	toSerialize["issuer"] = o.Issuer
	toSerialize["private_key"] = o.PrivateKey
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	if o.RefreshExpiresInSeconds.IsSet() {
		toSerialize["refresh_expires_in_seconds"] = o.RefreshExpiresInSeconds.Get()
	}
	return toSerialize, nil
}

func (o *CreateTenent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"audience",
		"issuer",
		"private_key",
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

	varCreateTenent := _CreateTenent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTenent)

	if err != nil {
		return err
	}

	*o = CreateTenent(varCreateTenent)

	return err
}

type NullableCreateTenent struct {
	value *CreateTenent
	isSet bool
}

func (v NullableCreateTenent) Get() *CreateTenent {
	return v.value
}

func (v *NullableCreateTenent) Set(val *CreateTenent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTenent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTenent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTenent(val *CreateTenent) *NullableCreateTenent {
	return &NullableCreateTenent{value: val, isSet: true}
}

func (v NullableCreateTenent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTenent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


