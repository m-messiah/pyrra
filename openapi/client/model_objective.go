/*
 * Athene
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Objective struct for Objective
type Objective struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Target      float64    `json:"target"`
	Window      int64      `json:"window"`
	Config      string     `json:"config"`
	Indicator   *Indicator `json:"indicator,omitempty"`
}

// NewObjective instantiates a new Objective object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjective(name string, description string, target float64, window int64, config string) *Objective {
	this := Objective{}
	this.Name = name
	this.Description = description
	this.Target = target
	this.Window = window
	this.Config = config
	return &this
}

// NewObjectiveWithDefaults instantiates a new Objective object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectiveWithDefaults() *Objective {
	this := Objective{}
	return &this
}

// GetName returns the Name field value
func (o *Objective) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Objective) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Objective) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Objective) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Objective) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Objective) SetDescription(v string) {
	o.Description = v
}

// GetTarget returns the Target field value
func (o *Objective) GetTarget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *Objective) GetTargetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *Objective) SetTarget(v float64) {
	o.Target = v
}

// GetWindow returns the Window field value
func (o *Objective) GetWindow() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *Objective) GetWindowOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *Objective) SetWindow(v int64) {
	o.Window = v
}

// GetConfig returns the Config field value
func (o *Objective) GetConfig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *Objective) GetConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *Objective) SetConfig(v string) {
	o.Config = v
}

// GetIndicator returns the Indicator field value if set, zero value otherwise.
func (o *Objective) GetIndicator() Indicator {
	if o == nil || o.Indicator == nil {
		var ret Indicator
		return ret
	}
	return *o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Objective) GetIndicatorOk() (*Indicator, bool) {
	if o == nil || o.Indicator == nil {
		return nil, false
	}
	return o.Indicator, true
}

// HasIndicator returns a boolean if a field has been set.
func (o *Objective) HasIndicator() bool {
	if o != nil && o.Indicator != nil {
		return true
	}

	return false
}

// SetIndicator gets a reference to the given Indicator and assigns it to the Indicator field.
func (o *Objective) SetIndicator(v Indicator) {
	o.Indicator = &v
}

func (o Objective) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["window"] = o.Window
	}
	if true {
		toSerialize["config"] = o.Config
	}
	if o.Indicator != nil {
		toSerialize["indicator"] = o.Indicator
	}
	return json.Marshal(toSerialize)
}

type NullableObjective struct {
	value *Objective
	isSet bool
}

func (v NullableObjective) Get() *Objective {
	return v.value
}

func (v *NullableObjective) Set(val *Objective) {
	v.value = val
	v.isSet = true
}

func (v NullableObjective) IsSet() bool {
	return v.isSet
}

func (v *NullableObjective) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjective(val *Objective) *NullableObjective {
	return &NullableObjective{value: val, isSet: true}
}

func (v NullableObjective) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjective) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
