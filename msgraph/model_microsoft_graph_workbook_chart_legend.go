/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// MicrosoftGraphWorkbookChartLegend struct for MicrosoftGraphWorkbookChartLegend
type MicrosoftGraphWorkbookChartLegend struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Boolean value for whether the chart legend should overlap with the main body of the chart.
	Overlay NullableBool `json:"overlay,omitempty"`
	// Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner, Custom.
	Position NullableString `json:"position,omitempty"`
	// A boolean value the represents the visibility of a ChartLegend object.
	Visible *bool `json:"visible,omitempty"`
	// Represents the formatting of a chart legend, which includes fill and font formatting. Read-only.
	Format AnyOfmicrosoftGraphWorkbookChartLegendFormat `json:"format,omitempty"`
}

// NewMicrosoftGraphWorkbookChartLegend instantiates a new MicrosoftGraphWorkbookChartLegend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookChartLegend() *MicrosoftGraphWorkbookChartLegend {
	this := MicrosoftGraphWorkbookChartLegend{}
	return &this
}

// NewMicrosoftGraphWorkbookChartLegendWithDefaults instantiates a new MicrosoftGraphWorkbookChartLegend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookChartLegendWithDefaults() *MicrosoftGraphWorkbookChartLegend {
	this := MicrosoftGraphWorkbookChartLegend{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartLegend) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartLegend) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartLegend) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookChartLegend) SetId(v string) {
	o.Id = &v
}

// GetOverlay returns the Overlay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartLegend) GetOverlay() bool {
	if o == nil || o.Overlay.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Overlay.Get()
}

// GetOverlayOk returns a tuple with the Overlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartLegend) GetOverlayOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Overlay.Get(), o.Overlay.IsSet()
}

// HasOverlay returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartLegend) HasOverlay() bool {
	if o != nil && o.Overlay.IsSet() {
		return true
	}

	return false
}

// SetOverlay gets a reference to the given NullableBool and assigns it to the Overlay field.
func (o *MicrosoftGraphWorkbookChartLegend) SetOverlay(v bool) {
	o.Overlay.Set(&v)
}
// SetOverlayNil sets the value for Overlay to be an explicit nil
func (o *MicrosoftGraphWorkbookChartLegend) SetOverlayNil() {
	o.Overlay.Set(nil)
}

// UnsetOverlay ensures that no value is present for Overlay, not even an explicit nil
func (o *MicrosoftGraphWorkbookChartLegend) UnsetOverlay() {
	o.Overlay.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartLegend) GetPosition() string {
	if o == nil || o.Position.Get() == nil {
		var ret string
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartLegend) GetPositionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartLegend) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableString and assigns it to the Position field.
func (o *MicrosoftGraphWorkbookChartLegend) SetPosition(v string) {
	o.Position.Set(&v)
}
// SetPositionNil sets the value for Position to be an explicit nil
func (o *MicrosoftGraphWorkbookChartLegend) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *MicrosoftGraphWorkbookChartLegend) UnsetPosition() {
	o.Position.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartLegend) GetVisible() bool {
	if o == nil || o.Visible == nil {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartLegend) GetVisibleOk() (*bool, bool) {
	if o == nil || o.Visible == nil {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartLegend) HasVisible() bool {
	if o != nil && o.Visible != nil {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *MicrosoftGraphWorkbookChartLegend) SetVisible(v bool) {
	o.Visible = &v
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartLegend) GetFormat() AnyOfmicrosoftGraphWorkbookChartLegendFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartLegendFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartLegend) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartLegendFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return &o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartLegend) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartLegendFormat and assigns it to the Format field.
func (o *MicrosoftGraphWorkbookChartLegend) SetFormat(v AnyOfmicrosoftGraphWorkbookChartLegendFormat) {
	o.Format = v
}

func (o MicrosoftGraphWorkbookChartLegend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Overlay.IsSet() {
		toSerialize["overlay"] = o.Overlay.Get()
	}
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	if o.Visible != nil {
		toSerialize["visible"] = o.Visible
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookChartLegend struct {
	value *MicrosoftGraphWorkbookChartLegend
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookChartLegend) Get() *MicrosoftGraphWorkbookChartLegend {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookChartLegend) Set(val *MicrosoftGraphWorkbookChartLegend) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookChartLegend) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookChartLegend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookChartLegend(val *MicrosoftGraphWorkbookChartLegend) *NullableMicrosoftGraphWorkbookChartLegend {
	return &NullableMicrosoftGraphWorkbookChartLegend{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookChartLegend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookChartLegend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


