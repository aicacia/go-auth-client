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

// checks if the ServiceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccount{}

// ServiceAccount struct for ServiceAccount
type ServiceAccount struct {
	Active bool `json:"active"`
	ClientId string `json:"client_id"`
	ClientSecret NullableString `json:"client_secret,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Id int64 `json:"id"`
	Name string `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _ServiceAccount ServiceAccount

// NewServiceAccount instantiates a new ServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccount(active bool, clientId string, createdAt time.Time, id int64, name string, updatedAt time.Time) *ServiceAccount {
	this := ServiceAccount{}
	this.Active = active
	this.ClientId = clientId
	this.CreatedAt = createdAt
	this.Id = id
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewServiceAccountWithDefaults instantiates a new ServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountWithDefaults() *ServiceAccount {
	this := ServiceAccount{}
	return &this
}

// GetActive returns the Active field value
func (o *ServiceAccount) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *ServiceAccount) SetActive(v bool) {
	o.Active = v
}

// GetClientId returns the ClientId field value
func (o *ServiceAccount) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ServiceAccount) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccount) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccount) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ServiceAccount) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *ServiceAccount) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}
// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *ServiceAccount) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *ServiceAccount) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServiceAccount) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServiceAccount) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *ServiceAccount) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceAccount) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ServiceAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceAccount) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ServiceAccount) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ServiceAccount) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["client_id"] = o.ClientId
	if o.ClientSecret.IsSet() {
		toSerialize["client_secret"] = o.ClientSecret.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ServiceAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"client_id",
		"created_at",
		"id",
		"name",
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

	varServiceAccount := _ServiceAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceAccount)

	if err != nil {
		return err
	}

	*o = ServiceAccount(varServiceAccount)

	return err
}

type NullableServiceAccount struct {
	value *ServiceAccount
	isSet bool
}

func (v NullableServiceAccount) Get() *ServiceAccount {
	return v.value
}

func (v *NullableServiceAccount) Set(val *ServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccount(val *ServiceAccount) *NullableServiceAccount {
	return &NullableServiceAccount{value: val, isSet: true}
}

func (v NullableServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


