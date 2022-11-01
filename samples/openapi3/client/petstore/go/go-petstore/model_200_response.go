/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
)

// Model200Response Model for testing model name starting with number
type Model200Response struct {
	Name *int32 `json:"name,omitempty"`
	Class *string `json:"class,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Model200Response Model200Response

// NewModel200Response instantiates a new Model200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel200Response() *Model200Response {
	this := Model200Response{}
	return &this
}

// NewModel200ResponseWithDefaults instantiates a new Model200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel200ResponseWithDefaults() *Model200Response {
	this := Model200Response{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Model200Response) GetName() int32 {
	if o == nil || isNil(o.Name) {
		var ret int32
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model200Response) GetNameOk() (*int32, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Model200Response) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given int32 and assigns it to the Name field.
func (o *Model200Response) SetName(v int32) {
	o.Name = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *Model200Response) GetClass() string {
	if o == nil || isNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model200Response) GetClassOk() (*string, bool) {
	if o == nil || isNil(o.Class) {
    return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *Model200Response) HasClass() bool {
	if o != nil && !isNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *Model200Response) SetClass(v string) {
	o.Class = &v
}

func (o Model200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Class) {
		toSerialize["class"] = o.Class
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Model200Response) UnmarshalJSON(bytes []byte) (err error) {
	varModel200Response := _Model200Response{}

	if err = json.Unmarshal(bytes, &varModel200Response); err == nil {
		*o = Model200Response(varModel200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "class")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModel200Response struct {
	value *Model200Response
	isSet bool
}

func (v NullableModel200Response) Get() *Model200Response {
	return v.value
}

func (v *NullableModel200Response) Set(val *Model200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel200Response(val *Model200Response) *NullableModel200Response {
	return &NullableModel200Response{value: val, isSet: true}
}

func (v NullableModel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


