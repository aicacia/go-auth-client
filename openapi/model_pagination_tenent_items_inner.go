/*
auth

auth

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the PaginationTenentItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationTenentItemsInner{}

// PaginationTenentItemsInner struct for PaginationTenentItemsInner
type PaginationTenentItemsInner struct {
	Algorithm string `json:"algorithm"`
	Audience NullableString `json:"audience,omitempty"`
	ClientId string `json:"client_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresInSeconds int64 `json:"expires_in_seconds"`
	Id int64 `json:"id"`
	Issuer string `json:"issuer"`
	Oauth2Providers []TenentOAuth2Provider `json:"oauth2_providers"`
	PrivateKey string `json:"private_key"`
	PublicKey NullableString `json:"public_key,omitempty"`
	RefreshExpiresInSeconds int64 `json:"refresh_expires_in_seconds"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _PaginationTenentItemsInner PaginationTenentItemsInner

// NewPaginationTenentItemsInner instantiates a new PaginationTenentItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationTenentItemsInner(algorithm string, clientId string, createdAt time.Time, expiresInSeconds int64, id int64, issuer string, oauth2Providers []TenentOAuth2Provider, privateKey string, refreshExpiresInSeconds int64, updatedAt time.Time) *PaginationTenentItemsInner {
	this := PaginationTenentItemsInner{}
	this.Algorithm = algorithm
	this.ClientId = clientId
	this.CreatedAt = createdAt
	this.ExpiresInSeconds = expiresInSeconds
	this.Id = id
	this.Issuer = issuer
	this.Oauth2Providers = oauth2Providers
	this.PrivateKey = privateKey
	this.RefreshExpiresInSeconds = refreshExpiresInSeconds
	this.UpdatedAt = updatedAt
	return &this
}

// NewPaginationTenentItemsInnerWithDefaults instantiates a new PaginationTenentItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationTenentItemsInnerWithDefaults() *PaginationTenentItemsInner {
	this := PaginationTenentItemsInner{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *PaginationTenentItemsInner) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *PaginationTenentItemsInner) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *PaginationTenentItemsInner) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetAudience returns the Audience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationTenentItemsInner) GetAudience() string {
	if o == nil || IsNil(o.Audience.Get()) {
		var ret string
		return ret
	}
	return *o.Audience.Get()
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationTenentItemsInner) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Audience.Get(), o.Audience.IsSet()
}

// HasAudience returns a boolean if a field has been set.
func (o *PaginationTenentItemsInner) HasAudience() bool {
	if o != nil && o.Audience.IsSet() {
		return true
	}

	return false
}

// SetAudience gets a reference to the given NullableString and assigns it to the Audience field.
func (o *PaginationTenentItemsInner) SetAudience(v string) {
	o.Audience.Set(&v)
}
// SetAudienceNil sets the value for Audience to be an explicit nil
func (o *PaginationTenentItemsInner) SetAudienceNil() {
	o.Audience.Set(nil)
}

// UnsetAudience ensures that no value is present for Audience, not even an explicit nil
func (o *PaginationTenentItemsInner) UnsetAudience() {
	o.Audience.Unset()
}

// GetClientId returns the ClientId field value
func (o *PaginationTenentItemsInner) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *PaginationTenentItemsInner) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *PaginationTenentItemsInner) SetClientId(v string) {
	o.ClientId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaginationTenentItemsInner) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaginationTenentItemsInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaginationTenentItemsInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiresInSeconds returns the ExpiresInSeconds field value
func (o *PaginationTenentItemsInner) GetExpiresInSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresInSeconds
}

// GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field value
// and a boolean to check if the value has been set.
func (o *PaginationTenentItemsInner) GetExpiresInSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresInSeconds, true
}

// SetExpiresInSeconds sets field value
func (o *PaginationTenentItemsInner) SetExpiresInSeconds(v int64) {
	o.ExpiresInSeconds = v
}

// GetId returns the Id field value
func (o *PaginationTenentItemsInner) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaginationTenentItemsInner) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaginationTenentItemsInner) SetId(v int64) {
	o.Id = v
}

// GetIssuer returns the Issuer field value
func (o *PaginationTenentItemsInner) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *PaginationTenentItemsInner) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *PaginationTenentItemsInner) SetIssuer(v string) {
	o.Issuer = v
}

// GetOauth2Providers returns the Oauth2Providers field value
func (o *PaginationTenentItemsInner) GetOauth2Providers() []TenentOAuth2Provider {
	if o == nil {
		var ret []TenentOAuth2Provider
		return ret
	}

	return o.Oauth2Providers
}

// GetOauth2ProvidersOk returns a tuple with the Oauth2Providers field value
// and a boolean to check if the value has been set.
func (o *PaginationTenentItemsInner) GetOauth2ProvidersOk() ([]TenentOAuth2Provider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Oauth2Providers, true
}

// SetOauth2Providers sets field value
func (o *PaginationTenentItemsInner) SetOauth2Providers(v []TenentOAuth2Provider) {
	o.Oauth2Providers = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *PaginationTenentItemsInner) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *PaginationTenentItemsInner) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *PaginationTenentItemsInner) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationTenentItemsInner) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationTenentItemsInner) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PaginationTenentItemsInner) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *PaginationTenentItemsInner) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *PaginationTenentItemsInner) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *PaginationTenentItemsInner) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetRefreshExpiresInSeconds returns the RefreshExpiresInSeconds field value
func (o *PaginationTenentItemsInner) GetRefreshExpiresInSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RefreshExpiresInSeconds
}

// GetRefreshExpiresInSecondsOk returns a tuple with the RefreshExpiresInSeconds field value
// and a boolean to check if the value has been set.
func (o *PaginationTenentItemsInner) GetRefreshExpiresInSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshExpiresInSeconds, true
}

// SetRefreshExpiresInSeconds sets field value
func (o *PaginationTenentItemsInner) SetRefreshExpiresInSeconds(v int64) {
	o.RefreshExpiresInSeconds = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PaginationTenentItemsInner) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PaginationTenentItemsInner) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PaginationTenentItemsInner) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PaginationTenentItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationTenentItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algorithm"] = o.Algorithm
	if o.Audience.IsSet() {
		toSerialize["audience"] = o.Audience.Get()
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["expires_in_seconds"] = o.ExpiresInSeconds
	toSerialize["id"] = o.Id
	toSerialize["issuer"] = o.Issuer
	toSerialize["oauth2_providers"] = o.Oauth2Providers
	toSerialize["private_key"] = o.PrivateKey
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	toSerialize["refresh_expires_in_seconds"] = o.RefreshExpiresInSeconds
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *PaginationTenentItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"algorithm",
		"client_id",
		"created_at",
		"expires_in_seconds",
		"id",
		"issuer",
		"oauth2_providers",
		"private_key",
		"refresh_expires_in_seconds",
		"updated_at",
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

	varPaginationTenentItemsInner := _PaginationTenentItemsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginationTenentItemsInner)

	if err != nil {
		return err
	}

	*o = PaginationTenentItemsInner(varPaginationTenentItemsInner)

	return err
}

type NullablePaginationTenentItemsInner struct {
	value *PaginationTenentItemsInner
	isSet bool
}

func (v NullablePaginationTenentItemsInner) Get() *PaginationTenentItemsInner {
	return v.value
}

func (v *NullablePaginationTenentItemsInner) Set(val *PaginationTenentItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationTenentItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationTenentItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationTenentItemsInner(val *PaginationTenentItemsInner) *NullablePaginationTenentItemsInner {
	return &NullablePaginationTenentItemsInner{value: val, isSet: true}
}

func (v NullablePaginationTenentItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationTenentItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

