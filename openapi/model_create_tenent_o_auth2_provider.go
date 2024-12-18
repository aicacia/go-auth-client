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

// checks if the CreateTenentOAuth2Provider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTenentOAuth2Provider{}

// CreateTenentOAuth2Provider struct for CreateTenentOAuth2Provider
type CreateTenentOAuth2Provider struct {
	AuthUrl NullableString `json:"auth_url,omitempty"`
	CallbackUrl NullableString `json:"callback_url,omitempty"`
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Provider string `json:"provider"`
	RedirectUrl string `json:"redirect_url"`
	Scope NullableString `json:"scope,omitempty"`
	TokenUrl NullableString `json:"token_url,omitempty"`
}

type _CreateTenentOAuth2Provider CreateTenentOAuth2Provider

// NewCreateTenentOAuth2Provider instantiates a new CreateTenentOAuth2Provider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTenentOAuth2Provider(clientId string, clientSecret string, provider string, redirectUrl string) *CreateTenentOAuth2Provider {
	this := CreateTenentOAuth2Provider{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.Provider = provider
	this.RedirectUrl = redirectUrl
	return &this
}

// NewCreateTenentOAuth2ProviderWithDefaults instantiates a new CreateTenentOAuth2Provider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTenentOAuth2ProviderWithDefaults() *CreateTenentOAuth2Provider {
	this := CreateTenentOAuth2Provider{}
	return &this
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTenentOAuth2Provider) GetAuthUrl() string {
	if o == nil || IsNil(o.AuthUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AuthUrl.Get()
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTenentOAuth2Provider) GetAuthUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUrl.Get(), o.AuthUrl.IsSet()
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *CreateTenentOAuth2Provider) HasAuthUrl() bool {
	if o != nil && o.AuthUrl.IsSet() {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given NullableString and assigns it to the AuthUrl field.
func (o *CreateTenentOAuth2Provider) SetAuthUrl(v string) {
	o.AuthUrl.Set(&v)
}
// SetAuthUrlNil sets the value for AuthUrl to be an explicit nil
func (o *CreateTenentOAuth2Provider) SetAuthUrlNil() {
	o.AuthUrl.Set(nil)
}

// UnsetAuthUrl ensures that no value is present for AuthUrl, not even an explicit nil
func (o *CreateTenentOAuth2Provider) UnsetAuthUrl() {
	o.AuthUrl.Unset()
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTenentOAuth2Provider) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CallbackUrl.Get()
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTenentOAuth2Provider) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallbackUrl.Get(), o.CallbackUrl.IsSet()
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *CreateTenentOAuth2Provider) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl.IsSet() {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given NullableString and assigns it to the CallbackUrl field.
func (o *CreateTenentOAuth2Provider) SetCallbackUrl(v string) {
	o.CallbackUrl.Set(&v)
}
// SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil
func (o *CreateTenentOAuth2Provider) SetCallbackUrlNil() {
	o.CallbackUrl.Set(nil)
}

// UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
func (o *CreateTenentOAuth2Provider) UnsetCallbackUrl() {
	o.CallbackUrl.Unset()
}

// GetClientId returns the ClientId field value
func (o *CreateTenentOAuth2Provider) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *CreateTenentOAuth2Provider) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *CreateTenentOAuth2Provider) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *CreateTenentOAuth2Provider) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *CreateTenentOAuth2Provider) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *CreateTenentOAuth2Provider) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetProvider returns the Provider field value
func (o *CreateTenentOAuth2Provider) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CreateTenentOAuth2Provider) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CreateTenentOAuth2Provider) SetProvider(v string) {
	o.Provider = v
}

// GetRedirectUrl returns the RedirectUrl field value
func (o *CreateTenentOAuth2Provider) GetRedirectUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
func (o *CreateTenentOAuth2Provider) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// SetRedirectUrl sets field value
func (o *CreateTenentOAuth2Provider) SetRedirectUrl(v string) {
	o.RedirectUrl = v
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTenentOAuth2Provider) GetScope() string {
	if o == nil || IsNil(o.Scope.Get()) {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTenentOAuth2Provider) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *CreateTenentOAuth2Provider) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *CreateTenentOAuth2Provider) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *CreateTenentOAuth2Provider) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *CreateTenentOAuth2Provider) UnsetScope() {
	o.Scope.Unset()
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTenentOAuth2Provider) GetTokenUrl() string {
	if o == nil || IsNil(o.TokenUrl.Get()) {
		var ret string
		return ret
	}
	return *o.TokenUrl.Get()
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTenentOAuth2Provider) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenUrl.Get(), o.TokenUrl.IsSet()
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *CreateTenentOAuth2Provider) HasTokenUrl() bool {
	if o != nil && o.TokenUrl.IsSet() {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given NullableString and assigns it to the TokenUrl field.
func (o *CreateTenentOAuth2Provider) SetTokenUrl(v string) {
	o.TokenUrl.Set(&v)
}
// SetTokenUrlNil sets the value for TokenUrl to be an explicit nil
func (o *CreateTenentOAuth2Provider) SetTokenUrlNil() {
	o.TokenUrl.Set(nil)
}

// UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil
func (o *CreateTenentOAuth2Provider) UnsetTokenUrl() {
	o.TokenUrl.Unset()
}

func (o CreateTenentOAuth2Provider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTenentOAuth2Provider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthUrl.IsSet() {
		toSerialize["auth_url"] = o.AuthUrl.Get()
	}
	if o.CallbackUrl.IsSet() {
		toSerialize["callback_url"] = o.CallbackUrl.Get()
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret
	toSerialize["provider"] = o.Provider
	toSerialize["redirect_url"] = o.RedirectUrl
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	if o.TokenUrl.IsSet() {
		toSerialize["token_url"] = o.TokenUrl.Get()
	}
	return toSerialize, nil
}

func (o *CreateTenentOAuth2Provider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_id",
		"client_secret",
		"provider",
		"redirect_url",
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

	varCreateTenentOAuth2Provider := _CreateTenentOAuth2Provider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTenentOAuth2Provider)

	if err != nil {
		return err
	}

	*o = CreateTenentOAuth2Provider(varCreateTenentOAuth2Provider)

	return err
}

type NullableCreateTenentOAuth2Provider struct {
	value *CreateTenentOAuth2Provider
	isSet bool
}

func (v NullableCreateTenentOAuth2Provider) Get() *CreateTenentOAuth2Provider {
	return v.value
}

func (v *NullableCreateTenentOAuth2Provider) Set(val *CreateTenentOAuth2Provider) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTenentOAuth2Provider) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTenentOAuth2Provider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTenentOAuth2Provider(val *CreateTenentOAuth2Provider) *NullableCreateTenentOAuth2Provider {
	return &NullableCreateTenentOAuth2Provider{value: val, isSet: true}
}

func (v NullableCreateTenentOAuth2Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTenentOAuth2Provider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


