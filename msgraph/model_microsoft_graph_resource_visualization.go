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

// MicrosoftGraphResourceVisualization struct for MicrosoftGraphResourceVisualization
type MicrosoftGraphResourceVisualization struct {
	// A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item.
	ContainerDisplayName NullableString `json:"containerDisplayName,omitempty"`
	// Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
	ContainerType NullableString `json:"containerType,omitempty"`
	// A path leading to the folder in which the item is stored.
	ContainerWebUrl NullableString `json:"containerWebUrl,omitempty"`
	// The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Note that not all Media Mime Types are supported.
	MediaType NullableString `json:"mediaType,omitempty"`
	// A URL leading to the preview image for the item.
	PreviewImageUrl NullableString `json:"previewImageUrl,omitempty"`
	// A preview text for the item.
	PreviewText NullableString `json:"previewText,omitempty"`
	// The item's title text.
	Title NullableString `json:"title,omitempty"`
	// The item's media type. Can be used for filtering for a specific file based on a specific type. See below for supported types.
	Type NullableString `json:"type,omitempty"`
}

// NewMicrosoftGraphResourceVisualization instantiates a new MicrosoftGraphResourceVisualization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphResourceVisualization() *MicrosoftGraphResourceVisualization {
	this := MicrosoftGraphResourceVisualization{}
	return &this
}

// NewMicrosoftGraphResourceVisualizationWithDefaults instantiates a new MicrosoftGraphResourceVisualization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphResourceVisualizationWithDefaults() *MicrosoftGraphResourceVisualization {
	this := MicrosoftGraphResourceVisualization{}
	return &this
}

// GetContainerDisplayName returns the ContainerDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceVisualization) GetContainerDisplayName() string {
	if o == nil || o.ContainerDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContainerDisplayName.Get()
}

// GetContainerDisplayNameOk returns a tuple with the ContainerDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceVisualization) GetContainerDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerDisplayName.Get(), o.ContainerDisplayName.IsSet()
}

// HasContainerDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceVisualization) HasContainerDisplayName() bool {
	if o != nil && o.ContainerDisplayName.IsSet() {
		return true
	}

	return false
}

// SetContainerDisplayName gets a reference to the given NullableString and assigns it to the ContainerDisplayName field.
func (o *MicrosoftGraphResourceVisualization) SetContainerDisplayName(v string) {
	o.ContainerDisplayName.Set(&v)
}
// SetContainerDisplayNameNil sets the value for ContainerDisplayName to be an explicit nil
func (o *MicrosoftGraphResourceVisualization) SetContainerDisplayNameNil() {
	o.ContainerDisplayName.Set(nil)
}

// UnsetContainerDisplayName ensures that no value is present for ContainerDisplayName, not even an explicit nil
func (o *MicrosoftGraphResourceVisualization) UnsetContainerDisplayName() {
	o.ContainerDisplayName.Unset()
}

// GetContainerType returns the ContainerType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceVisualization) GetContainerType() string {
	if o == nil || o.ContainerType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContainerType.Get()
}

// GetContainerTypeOk returns a tuple with the ContainerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceVisualization) GetContainerTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerType.Get(), o.ContainerType.IsSet()
}

// HasContainerType returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceVisualization) HasContainerType() bool {
	if o != nil && o.ContainerType.IsSet() {
		return true
	}

	return false
}

// SetContainerType gets a reference to the given NullableString and assigns it to the ContainerType field.
func (o *MicrosoftGraphResourceVisualization) SetContainerType(v string) {
	o.ContainerType.Set(&v)
}
// SetContainerTypeNil sets the value for ContainerType to be an explicit nil
func (o *MicrosoftGraphResourceVisualization) SetContainerTypeNil() {
	o.ContainerType.Set(nil)
}

// UnsetContainerType ensures that no value is present for ContainerType, not even an explicit nil
func (o *MicrosoftGraphResourceVisualization) UnsetContainerType() {
	o.ContainerType.Unset()
}

// GetContainerWebUrl returns the ContainerWebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceVisualization) GetContainerWebUrl() string {
	if o == nil || o.ContainerWebUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContainerWebUrl.Get()
}

// GetContainerWebUrlOk returns a tuple with the ContainerWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceVisualization) GetContainerWebUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerWebUrl.Get(), o.ContainerWebUrl.IsSet()
}

// HasContainerWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceVisualization) HasContainerWebUrl() bool {
	if o != nil && o.ContainerWebUrl.IsSet() {
		return true
	}

	return false
}

// SetContainerWebUrl gets a reference to the given NullableString and assigns it to the ContainerWebUrl field.
func (o *MicrosoftGraphResourceVisualization) SetContainerWebUrl(v string) {
	o.ContainerWebUrl.Set(&v)
}
// SetContainerWebUrlNil sets the value for ContainerWebUrl to be an explicit nil
func (o *MicrosoftGraphResourceVisualization) SetContainerWebUrlNil() {
	o.ContainerWebUrl.Set(nil)
}

// UnsetContainerWebUrl ensures that no value is present for ContainerWebUrl, not even an explicit nil
func (o *MicrosoftGraphResourceVisualization) UnsetContainerWebUrl() {
	o.ContainerWebUrl.Unset()
}

// GetMediaType returns the MediaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceVisualization) GetMediaType() string {
	if o == nil || o.MediaType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MediaType.Get()
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceVisualization) GetMediaTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MediaType.Get(), o.MediaType.IsSet()
}

// HasMediaType returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceVisualization) HasMediaType() bool {
	if o != nil && o.MediaType.IsSet() {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given NullableString and assigns it to the MediaType field.
func (o *MicrosoftGraphResourceVisualization) SetMediaType(v string) {
	o.MediaType.Set(&v)
}
// SetMediaTypeNil sets the value for MediaType to be an explicit nil
func (o *MicrosoftGraphResourceVisualization) SetMediaTypeNil() {
	o.MediaType.Set(nil)
}

// UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
func (o *MicrosoftGraphResourceVisualization) UnsetMediaType() {
	o.MediaType.Unset()
}

// GetPreviewImageUrl returns the PreviewImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceVisualization) GetPreviewImageUrl() string {
	if o == nil || o.PreviewImageUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviewImageUrl.Get()
}

// GetPreviewImageUrlOk returns a tuple with the PreviewImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceVisualization) GetPreviewImageUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreviewImageUrl.Get(), o.PreviewImageUrl.IsSet()
}

// HasPreviewImageUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceVisualization) HasPreviewImageUrl() bool {
	if o != nil && o.PreviewImageUrl.IsSet() {
		return true
	}

	return false
}

// SetPreviewImageUrl gets a reference to the given NullableString and assigns it to the PreviewImageUrl field.
func (o *MicrosoftGraphResourceVisualization) SetPreviewImageUrl(v string) {
	o.PreviewImageUrl.Set(&v)
}
// SetPreviewImageUrlNil sets the value for PreviewImageUrl to be an explicit nil
func (o *MicrosoftGraphResourceVisualization) SetPreviewImageUrlNil() {
	o.PreviewImageUrl.Set(nil)
}

// UnsetPreviewImageUrl ensures that no value is present for PreviewImageUrl, not even an explicit nil
func (o *MicrosoftGraphResourceVisualization) UnsetPreviewImageUrl() {
	o.PreviewImageUrl.Unset()
}

// GetPreviewText returns the PreviewText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceVisualization) GetPreviewText() string {
	if o == nil || o.PreviewText.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviewText.Get()
}

// GetPreviewTextOk returns a tuple with the PreviewText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceVisualization) GetPreviewTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreviewText.Get(), o.PreviewText.IsSet()
}

// HasPreviewText returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceVisualization) HasPreviewText() bool {
	if o != nil && o.PreviewText.IsSet() {
		return true
	}

	return false
}

// SetPreviewText gets a reference to the given NullableString and assigns it to the PreviewText field.
func (o *MicrosoftGraphResourceVisualization) SetPreviewText(v string) {
	o.PreviewText.Set(&v)
}
// SetPreviewTextNil sets the value for PreviewText to be an explicit nil
func (o *MicrosoftGraphResourceVisualization) SetPreviewTextNil() {
	o.PreviewText.Set(nil)
}

// UnsetPreviewText ensures that no value is present for PreviewText, not even an explicit nil
func (o *MicrosoftGraphResourceVisualization) UnsetPreviewText() {
	o.PreviewText.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceVisualization) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceVisualization) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceVisualization) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *MicrosoftGraphResourceVisualization) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *MicrosoftGraphResourceVisualization) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *MicrosoftGraphResourceVisualization) UnsetTitle() {
	o.Title.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceVisualization) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceVisualization) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceVisualization) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *MicrosoftGraphResourceVisualization) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *MicrosoftGraphResourceVisualization) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *MicrosoftGraphResourceVisualization) UnsetType() {
	o.Type.Unset()
}

func (o MicrosoftGraphResourceVisualization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContainerDisplayName.IsSet() {
		toSerialize["containerDisplayName"] = o.ContainerDisplayName.Get()
	}
	if o.ContainerType.IsSet() {
		toSerialize["containerType"] = o.ContainerType.Get()
	}
	if o.ContainerWebUrl.IsSet() {
		toSerialize["containerWebUrl"] = o.ContainerWebUrl.Get()
	}
	if o.MediaType.IsSet() {
		toSerialize["mediaType"] = o.MediaType.Get()
	}
	if o.PreviewImageUrl.IsSet() {
		toSerialize["previewImageUrl"] = o.PreviewImageUrl.Get()
	}
	if o.PreviewText.IsSet() {
		toSerialize["previewText"] = o.PreviewText.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphResourceVisualization struct {
	value *MicrosoftGraphResourceVisualization
	isSet bool
}

func (v NullableMicrosoftGraphResourceVisualization) Get() *MicrosoftGraphResourceVisualization {
	return v.value
}

func (v *NullableMicrosoftGraphResourceVisualization) Set(val *MicrosoftGraphResourceVisualization) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphResourceVisualization) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphResourceVisualization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphResourceVisualization(val *MicrosoftGraphResourceVisualization) *NullableMicrosoftGraphResourceVisualization {
	return &NullableMicrosoftGraphResourceVisualization{value: val, isSet: true}
}

func (v NullableMicrosoftGraphResourceVisualization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphResourceVisualization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


