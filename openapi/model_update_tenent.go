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

// checks if the UpdateTenent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTenent{}

// UpdateTenent struct for UpdateTenent
type UpdateTenent struct {
	Algorithm NullableAlgorithm `json:"algorithm,omitempty"`
	Audience NullableString `json:"audience,omitempty"`
	ClientId NullableString `json:"client_id,omitempty"`
	ExpiresInSeconds NullableInt64 `json:"expires_in_seconds,omitempty"`
	Issuer NullableString `json:"issuer,omitempty"`
	PrivateKey NullableString `json:"private_key,omitempty"`
	PublicKey NullableString `json:"public_key,omitempty"`
	RefreshExpiresInSeconds NullableInt64 `json:"refresh_expires_in_seconds,omitempty"`
}

// NewUpdateTenent instantiates a new UpdateTenent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTenent() *UpdateTenent {
	this := UpdateTenent{}
	return &this
}

// NewUpdateTenentWithDefaults instantiates a new UpdateTenent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTenentWithDefaults() *UpdateTenent {
	this := UpdateTenent{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenent) GetAlgorithm() Algorithm {
	if o == nil || IsNil(o.Algorithm.Get()) {
		var ret Algorithm
		return ret
	}
	return *o.Algorithm.Get()
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenent) GetAlgorithmOk() (*Algorithm, bool) {
	if o == nil {
		return nil, false
	}
	return o.Algorithm.Get(), o.Algorithm.IsSet()
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *UpdateTenent) HasAlgorithm() bool {
	if o != nil && o.Algorithm.IsSet() {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given NullableAlgorithm and assigns it to the Algorithm field.
func (o *UpdateTenent) SetAlgorithm(v Algorithm) {
	o.Algorithm.Set(&v)
}
// SetAlgorithmNil sets the value for Algorithm to be an explicit nil
func (o *UpdateTenent) SetAlgorithmNil() {
	o.Algorithm.Set(nil)
}

// UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
func (o *UpdateTenent) UnsetAlgorithm() {
	o.Algorithm.Unset()
}

// GetAudience returns the Audience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenent) GetAudience() string {
	if o == nil || IsNil(o.Audience.Get()) {
		var ret string
		return ret
	}
	return *o.Audience.Get()
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenent) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Audience.Get(), o.Audience.IsSet()
}

// HasAudience returns a boolean if a field has been set.
func (o *UpdateTenent) HasAudience() bool {
	if o != nil && o.Audience.IsSet() {
		return true
	}

	return false
}

// SetAudience gets a reference to the given NullableString and assigns it to the Audience field.
func (o *UpdateTenent) SetAudience(v string) {
	o.Audience.Set(&v)
}
// SetAudienceNil sets the value for Audience to be an explicit nil
func (o *UpdateTenent) SetAudienceNil() {
	o.Audience.Set(nil)
}

// UnsetAudience ensures that no value is present for Audience, not even an explicit nil
func (o *UpdateTenent) UnsetAudience() {
	o.Audience.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenent) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenent) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *UpdateTenent) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *UpdateTenent) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *UpdateTenent) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *UpdateTenent) UnsetClientId() {
	o.ClientId.Unset()
}

// GetExpiresInSeconds returns the ExpiresInSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenent) GetExpiresInSeconds() int64 {
	if o == nil || IsNil(o.ExpiresInSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiresInSeconds.Get()
}

// GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenent) GetExpiresInSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresInSeconds.Get(), o.ExpiresInSeconds.IsSet()
}

// HasExpiresInSeconds returns a boolean if a field has been set.
func (o *UpdateTenent) HasExpiresInSeconds() bool {
	if o != nil && o.ExpiresInSeconds.IsSet() {
		return true
	}

	return false
}

// SetExpiresInSeconds gets a reference to the given NullableInt64 and assigns it to the ExpiresInSeconds field.
func (o *UpdateTenent) SetExpiresInSeconds(v int64) {
	o.ExpiresInSeconds.Set(&v)
}
// SetExpiresInSecondsNil sets the value for ExpiresInSeconds to be an explicit nil
func (o *UpdateTenent) SetExpiresInSecondsNil() {
	o.ExpiresInSeconds.Set(nil)
}

// UnsetExpiresInSeconds ensures that no value is present for ExpiresInSeconds, not even an explicit nil
func (o *UpdateTenent) UnsetExpiresInSeconds() {
	o.ExpiresInSeconds.Unset()
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenent) GetIssuer() string {
	if o == nil || IsNil(o.Issuer.Get()) {
		var ret string
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenent) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *UpdateTenent) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullableString and assigns it to the Issuer field.
func (o *UpdateTenent) SetIssuer(v string) {
	o.Issuer.Set(&v)
}
// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *UpdateTenent) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *UpdateTenent) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenent) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey.Get()) {
		var ret string
		return ret
	}
	return *o.PrivateKey.Get()
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenent) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKey.Get(), o.PrivateKey.IsSet()
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *UpdateTenent) HasPrivateKey() bool {
	if o != nil && o.PrivateKey.IsSet() {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given NullableString and assigns it to the PrivateKey field.
func (o *UpdateTenent) SetPrivateKey(v string) {
	o.PrivateKey.Set(&v)
}
// SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil
func (o *UpdateTenent) SetPrivateKeyNil() {
	o.PrivateKey.Set(nil)
}

// UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
func (o *UpdateTenent) UnsetPrivateKey() {
	o.PrivateKey.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenent) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenent) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *UpdateTenent) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *UpdateTenent) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *UpdateTenent) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *UpdateTenent) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetRefreshExpiresInSeconds returns the RefreshExpiresInSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenent) GetRefreshExpiresInSeconds() int64 {
	if o == nil || IsNil(o.RefreshExpiresInSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.RefreshExpiresInSeconds.Get()
}

// GetRefreshExpiresInSecondsOk returns a tuple with the RefreshExpiresInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenent) GetRefreshExpiresInSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshExpiresInSeconds.Get(), o.RefreshExpiresInSeconds.IsSet()
}

// HasRefreshExpiresInSeconds returns a boolean if a field has been set.
func (o *UpdateTenent) HasRefreshExpiresInSeconds() bool {
	if o != nil && o.RefreshExpiresInSeconds.IsSet() {
		return true
	}

	return false
}

// SetRefreshExpiresInSeconds gets a reference to the given NullableInt64 and assigns it to the RefreshExpiresInSeconds field.
func (o *UpdateTenent) SetRefreshExpiresInSeconds(v int64) {
	o.RefreshExpiresInSeconds.Set(&v)
}
// SetRefreshExpiresInSecondsNil sets the value for RefreshExpiresInSeconds to be an explicit nil
func (o *UpdateTenent) SetRefreshExpiresInSecondsNil() {
	o.RefreshExpiresInSeconds.Set(nil)
}

// UnsetRefreshExpiresInSeconds ensures that no value is present for RefreshExpiresInSeconds, not even an explicit nil
func (o *UpdateTenent) UnsetRefreshExpiresInSeconds() {
	o.RefreshExpiresInSeconds.Unset()
}

func (o UpdateTenent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTenent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm.IsSet() {
		toSerialize["algorithm"] = o.Algorithm.Get()
	}
	if o.Audience.IsSet() {
		toSerialize["audience"] = o.Audience.Get()
	}
	if o.ClientId.IsSet() {
		toSerialize["client_id"] = o.ClientId.Get()
	}
	if o.ExpiresInSeconds.IsSet() {
		toSerialize["expires_in_seconds"] = o.ExpiresInSeconds.Get()
	}
	if o.Issuer.IsSet() {
		toSerialize["issuer"] = o.Issuer.Get()
	}
	if o.PrivateKey.IsSet() {
		toSerialize["private_key"] = o.PrivateKey.Get()
	}
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	if o.RefreshExpiresInSeconds.IsSet() {
		toSerialize["refresh_expires_in_seconds"] = o.RefreshExpiresInSeconds.Get()
	}
	return toSerialize, nil
}

type NullableUpdateTenent struct {
	value *UpdateTenent
	isSet bool
}

func (v NullableUpdateTenent) Get() *UpdateTenent {
	return v.value
}

func (v *NullableUpdateTenent) Set(val *UpdateTenent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTenent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTenent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTenent(val *UpdateTenent) *NullableUpdateTenent {
	return &NullableUpdateTenent{value: val, isSet: true}
}

func (v NullableUpdateTenent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTenent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

