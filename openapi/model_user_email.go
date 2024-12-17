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

// checks if the UserEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserEmail{}

// UserEmail struct for UserEmail
type UserEmail struct {
	CreatedAt time.Time `json:"created_at"`
	Email string `json:"email"`
	Id int64 `json:"id"`
	Primary bool `json:"primary"`
	UpdatedAt time.Time `json:"updated_at"`
	Verified bool `json:"verified"`
}

type _UserEmail UserEmail

// NewUserEmail instantiates a new UserEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEmail(createdAt time.Time, email string, id int64, primary bool, updatedAt time.Time, verified bool) *UserEmail {
	this := UserEmail{}
	this.CreatedAt = createdAt
	this.Email = email
	this.Id = id
	this.Primary = primary
	this.UpdatedAt = updatedAt
	this.Verified = verified
	return &this
}

// NewUserEmailWithDefaults instantiates a new UserEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEmailWithDefaults() *UserEmail {
	this := UserEmail{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserEmail) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserEmail) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserEmail) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEmail returns the Email field value
func (o *UserEmail) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserEmail) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserEmail) SetEmail(v string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *UserEmail) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserEmail) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserEmail) SetId(v int64) {
	o.Id = v
}

// GetPrimary returns the Primary field value
func (o *UserEmail) GetPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *UserEmail) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *UserEmail) SetPrimary(v bool) {
	o.Primary = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UserEmail) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UserEmail) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UserEmail) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVerified returns the Verified field value
func (o *UserEmail) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *UserEmail) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *UserEmail) SetVerified(v bool) {
	o.Verified = v
}

func (o UserEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["email"] = o.Email
	toSerialize["id"] = o.Id
	toSerialize["primary"] = o.Primary
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["verified"] = o.Verified
	return toSerialize, nil
}

func (o *UserEmail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"email",
		"id",
		"primary",
		"updated_at",
		"verified",
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

	varUserEmail := _UserEmail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserEmail)

	if err != nil {
		return err
	}

	*o = UserEmail(varUserEmail)

	return err
}

type NullableUserEmail struct {
	value *UserEmail
	isSet bool
}

func (v NullableUserEmail) Get() *UserEmail {
	return v.value
}

func (v *NullableUserEmail) Set(val *UserEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEmail(val *UserEmail) *NullableUserEmail {
	return &NullableUserEmail{value: val, isSet: true}
}

func (v NullableUserEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


