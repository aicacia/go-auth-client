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

// checks if the TenentOAuth2Provider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenentOAuth2Provider{}

// TenentOAuth2Provider struct for TenentOAuth2Provider
type TenentOAuth2Provider struct {
	Active bool `json:"active"`
	AuthUrl string `json:"auth_url"`
	CallbackUrl string `json:"callback_url"`
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	CreatedAt time.Time `json:"created_at"`
	Id int64 `json:"id"`
	Provider string `json:"provider"`
	RedirectUrl string `json:"redirect_url"`
	Scope string `json:"scope"`
	TenentId int64 `json:"tenent_id"`
	TokenUrl string `json:"token_url"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _TenentOAuth2Provider TenentOAuth2Provider

// NewTenentOAuth2Provider instantiates a new TenentOAuth2Provider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenentOAuth2Provider(active bool, authUrl string, callbackUrl string, clientId string, clientSecret string, createdAt time.Time, id int64, provider string, redirectUrl string, scope string, tenentId int64, tokenUrl string, updatedAt time.Time) *TenentOAuth2Provider {
	this := TenentOAuth2Provider{}
	this.Active = active
	this.AuthUrl = authUrl
	this.CallbackUrl = callbackUrl
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.CreatedAt = createdAt
	this.Id = id
	this.Provider = provider
	this.RedirectUrl = redirectUrl
	this.Scope = scope
	this.TenentId = tenentId
	this.TokenUrl = tokenUrl
	this.UpdatedAt = updatedAt
	return &this
}

// NewTenentOAuth2ProviderWithDefaults instantiates a new TenentOAuth2Provider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenentOAuth2ProviderWithDefaults() *TenentOAuth2Provider {
	this := TenentOAuth2Provider{}
	return &this
}

// GetActive returns the Active field value
func (o *TenentOAuth2Provider) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *TenentOAuth2Provider) SetActive(v bool) {
	o.Active = v
}

// GetAuthUrl returns the AuthUrl field value
func (o *TenentOAuth2Provider) GetAuthUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetAuthUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthUrl, true
}

// SetAuthUrl sets field value
func (o *TenentOAuth2Provider) SetAuthUrl(v string) {
	o.AuthUrl = v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *TenentOAuth2Provider) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *TenentOAuth2Provider) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetClientId returns the ClientId field value
func (o *TenentOAuth2Provider) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *TenentOAuth2Provider) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *TenentOAuth2Provider) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *TenentOAuth2Provider) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TenentOAuth2Provider) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TenentOAuth2Provider) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *TenentOAuth2Provider) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TenentOAuth2Provider) SetId(v int64) {
	o.Id = v
}

// GetProvider returns the Provider field value
func (o *TenentOAuth2Provider) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *TenentOAuth2Provider) SetProvider(v string) {
	o.Provider = v
}

// GetRedirectUrl returns the RedirectUrl field value
func (o *TenentOAuth2Provider) GetRedirectUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// SetRedirectUrl sets field value
func (o *TenentOAuth2Provider) SetRedirectUrl(v string) {
	o.RedirectUrl = v
}

// GetScope returns the Scope field value
func (o *TenentOAuth2Provider) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *TenentOAuth2Provider) SetScope(v string) {
	o.Scope = v
}

// GetTenentId returns the TenentId field value
func (o *TenentOAuth2Provider) GetTenentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TenentId
}

// GetTenentIdOk returns a tuple with the TenentId field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetTenentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenentId, true
}

// SetTenentId sets field value
func (o *TenentOAuth2Provider) SetTenentId(v int64) {
	o.TenentId = v
}

// GetTokenUrl returns the TokenUrl field value
func (o *TenentOAuth2Provider) GetTokenUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenUrl, true
}

// SetTokenUrl sets field value
func (o *TenentOAuth2Provider) SetTokenUrl(v string) {
	o.TokenUrl = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TenentOAuth2Provider) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TenentOAuth2Provider) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TenentOAuth2Provider) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o TenentOAuth2Provider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenentOAuth2Provider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["auth_url"] = o.AuthUrl
	toSerialize["callback_url"] = o.CallbackUrl
	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["provider"] = o.Provider
	toSerialize["redirect_url"] = o.RedirectUrl
	toSerialize["scope"] = o.Scope
	toSerialize["tenent_id"] = o.TenentId
	toSerialize["token_url"] = o.TokenUrl
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *TenentOAuth2Provider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"auth_url",
		"callback_url",
		"client_id",
		"client_secret",
		"created_at",
		"id",
		"provider",
		"redirect_url",
		"scope",
		"tenent_id",
		"token_url",
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

	varTenentOAuth2Provider := _TenentOAuth2Provider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenentOAuth2Provider)

	if err != nil {
		return err
	}

	*o = TenentOAuth2Provider(varTenentOAuth2Provider)

	return err
}

type NullableTenentOAuth2Provider struct {
	value *TenentOAuth2Provider
	isSet bool
}

func (v NullableTenentOAuth2Provider) Get() *TenentOAuth2Provider {
	return v.value
}

func (v *NullableTenentOAuth2Provider) Set(val *TenentOAuth2Provider) {
	v.value = val
	v.isSet = true
}

func (v NullableTenentOAuth2Provider) IsSet() bool {
	return v.isSet
}

func (v *NullableTenentOAuth2Provider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenentOAuth2Provider(val *TenentOAuth2Provider) *NullableTenentOAuth2Provider {
	return &NullableTenentOAuth2Provider{value: val, isSet: true}
}

func (v NullableTenentOAuth2Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenentOAuth2Provider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


