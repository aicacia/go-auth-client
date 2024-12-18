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

// checks if the UpdateTenentOAuth2Provider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTenentOAuth2Provider{}

// UpdateTenentOAuth2Provider struct for UpdateTenentOAuth2Provider
type UpdateTenentOAuth2Provider struct {
	Active NullableBool `json:"active,omitempty"`
	AuthUrl NullableString `json:"auth_url,omitempty"`
	CallbackUrl NullableString `json:"callback_url,omitempty"`
	ClientId NullableString `json:"client_id,omitempty"`
	ClientSecret NullableString `json:"client_secret,omitempty"`
	RedirectUrl NullableString `json:"redirect_url,omitempty"`
	Scope NullableString `json:"scope,omitempty"`
	TokenUrl NullableString `json:"token_url,omitempty"`
}

// NewUpdateTenentOAuth2Provider instantiates a new UpdateTenentOAuth2Provider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTenentOAuth2Provider() *UpdateTenentOAuth2Provider {
	this := UpdateTenentOAuth2Provider{}
	return &this
}

// NewUpdateTenentOAuth2ProviderWithDefaults instantiates a new UpdateTenentOAuth2Provider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTenentOAuth2ProviderWithDefaults() *UpdateTenentOAuth2Provider {
	this := UpdateTenentOAuth2Provider{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenentOAuth2Provider) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenentOAuth2Provider) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateTenentOAuth2Provider) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *UpdateTenentOAuth2Provider) SetActive(v bool) {
	o.Active.Set(&v)
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *UpdateTenentOAuth2Provider) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *UpdateTenentOAuth2Provider) UnsetActive() {
	o.Active.Unset()
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenentOAuth2Provider) GetAuthUrl() string {
	if o == nil || IsNil(o.AuthUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AuthUrl.Get()
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenentOAuth2Provider) GetAuthUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUrl.Get(), o.AuthUrl.IsSet()
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *UpdateTenentOAuth2Provider) HasAuthUrl() bool {
	if o != nil && o.AuthUrl.IsSet() {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given NullableString and assigns it to the AuthUrl field.
func (o *UpdateTenentOAuth2Provider) SetAuthUrl(v string) {
	o.AuthUrl.Set(&v)
}
// SetAuthUrlNil sets the value for AuthUrl to be an explicit nil
func (o *UpdateTenentOAuth2Provider) SetAuthUrlNil() {
	o.AuthUrl.Set(nil)
}

// UnsetAuthUrl ensures that no value is present for AuthUrl, not even an explicit nil
func (o *UpdateTenentOAuth2Provider) UnsetAuthUrl() {
	o.AuthUrl.Unset()
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenentOAuth2Provider) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CallbackUrl.Get()
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenentOAuth2Provider) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallbackUrl.Get(), o.CallbackUrl.IsSet()
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *UpdateTenentOAuth2Provider) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl.IsSet() {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given NullableString and assigns it to the CallbackUrl field.
func (o *UpdateTenentOAuth2Provider) SetCallbackUrl(v string) {
	o.CallbackUrl.Set(&v)
}
// SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil
func (o *UpdateTenentOAuth2Provider) SetCallbackUrlNil() {
	o.CallbackUrl.Set(nil)
}

// UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
func (o *UpdateTenentOAuth2Provider) UnsetCallbackUrl() {
	o.CallbackUrl.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenentOAuth2Provider) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenentOAuth2Provider) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *UpdateTenentOAuth2Provider) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *UpdateTenentOAuth2Provider) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *UpdateTenentOAuth2Provider) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *UpdateTenentOAuth2Provider) UnsetClientId() {
	o.ClientId.Unset()
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenentOAuth2Provider) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenentOAuth2Provider) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *UpdateTenentOAuth2Provider) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *UpdateTenentOAuth2Provider) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}
// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *UpdateTenentOAuth2Provider) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *UpdateTenentOAuth2Provider) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenentOAuth2Provider) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenentOAuth2Provider) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *UpdateTenentOAuth2Provider) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl.IsSet() {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given NullableString and assigns it to the RedirectUrl field.
func (o *UpdateTenentOAuth2Provider) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}
// SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil
func (o *UpdateTenentOAuth2Provider) SetRedirectUrlNil() {
	o.RedirectUrl.Set(nil)
}

// UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
func (o *UpdateTenentOAuth2Provider) UnsetRedirectUrl() {
	o.RedirectUrl.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenentOAuth2Provider) GetScope() string {
	if o == nil || IsNil(o.Scope.Get()) {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenentOAuth2Provider) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *UpdateTenentOAuth2Provider) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *UpdateTenentOAuth2Provider) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *UpdateTenentOAuth2Provider) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *UpdateTenentOAuth2Provider) UnsetScope() {
	o.Scope.Unset()
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenentOAuth2Provider) GetTokenUrl() string {
	if o == nil || IsNil(o.TokenUrl.Get()) {
		var ret string
		return ret
	}
	return *o.TokenUrl.Get()
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenentOAuth2Provider) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenUrl.Get(), o.TokenUrl.IsSet()
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *UpdateTenentOAuth2Provider) HasTokenUrl() bool {
	if o != nil && o.TokenUrl.IsSet() {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given NullableString and assigns it to the TokenUrl field.
func (o *UpdateTenentOAuth2Provider) SetTokenUrl(v string) {
	o.TokenUrl.Set(&v)
}
// SetTokenUrlNil sets the value for TokenUrl to be an explicit nil
func (o *UpdateTenentOAuth2Provider) SetTokenUrlNil() {
	o.TokenUrl.Set(nil)
}

// UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil
func (o *UpdateTenentOAuth2Provider) UnsetTokenUrl() {
	o.TokenUrl.Unset()
}

func (o UpdateTenentOAuth2Provider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTenentOAuth2Provider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	if o.AuthUrl.IsSet() {
		toSerialize["auth_url"] = o.AuthUrl.Get()
	}
	if o.CallbackUrl.IsSet() {
		toSerialize["callback_url"] = o.CallbackUrl.Get()
	}
	if o.ClientId.IsSet() {
		toSerialize["client_id"] = o.ClientId.Get()
	}
	if o.ClientSecret.IsSet() {
		toSerialize["client_secret"] = o.ClientSecret.Get()
	}
	if o.RedirectUrl.IsSet() {
		toSerialize["redirect_url"] = o.RedirectUrl.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	if o.TokenUrl.IsSet() {
		toSerialize["token_url"] = o.TokenUrl.Get()
	}
	return toSerialize, nil
}

type NullableUpdateTenentOAuth2Provider struct {
	value *UpdateTenentOAuth2Provider
	isSet bool
}

func (v NullableUpdateTenentOAuth2Provider) Get() *UpdateTenentOAuth2Provider {
	return v.value
}

func (v *NullableUpdateTenentOAuth2Provider) Set(val *UpdateTenentOAuth2Provider) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTenentOAuth2Provider) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTenentOAuth2Provider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTenentOAuth2Provider(val *UpdateTenentOAuth2Provider) *NullableUpdateTenentOAuth2Provider {
	return &NullableUpdateTenentOAuth2Provider{value: val, isSet: true}
}

func (v NullableUpdateTenentOAuth2Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTenentOAuth2Provider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


