/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// OnenotePage struct for OnenotePage
type OnenotePage struct {
	// The page's HTML content.
	Content NullableString `json:"content,omitempty"`
	// The URL for the page's HTML content.  Read-only.
	ContentUrl NullableString `json:"contentUrl,omitempty"`
	// The unique identifier of the application that created the page. Read-only.
	CreatedByAppId NullableString `json:"createdByAppId,omitempty"`
	// The date and time when the page was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// The indentation level of the page. Read-only.
	Level NullableInt32 `json:"level,omitempty"`
	// Links for opening the page. The oneNoteClientURL link opens the page in the OneNote native client if it 's installed. The oneNoteWebUrl link opens the page in OneNote on the web. Read-only.
	Links AnyOfmicrosoftGraphPageLinks `json:"links,omitempty"`
	// The order of the page within its parent section. Read-only.
	Order NullableInt32 `json:"order,omitempty"`
	// The title of the page.
	Title NullableString `json:"title,omitempty"`
	UserTags *[]*string `json:"userTags,omitempty"`
	// The notebook that contains the page.  Read-only.
	ParentNotebook AnyOfmicrosoftGraphNotebook `json:"parentNotebook,omitempty"`
	// The section that contains the page. Read-only.
	ParentSection AnyOfmicrosoftGraphOnenoteSection `json:"parentSection,omitempty"`
}

// NewOnenotePage instantiates a new OnenotePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnenotePage() *OnenotePage {
	this := OnenotePage{}
	return &this
}

// NewOnenotePageWithDefaults instantiates a new OnenotePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnenotePageWithDefaults() *OnenotePage {
	this := OnenotePage{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenotePage) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenotePage) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *OnenotePage) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *OnenotePage) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *OnenotePage) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *OnenotePage) UnsetContent() {
	o.Content.Unset()
}

// GetContentUrl returns the ContentUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenotePage) GetContentUrl() string {
	if o == nil || o.ContentUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentUrl.Get()
}

// GetContentUrlOk returns a tuple with the ContentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenotePage) GetContentUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentUrl.Get(), o.ContentUrl.IsSet()
}

// HasContentUrl returns a boolean if a field has been set.
func (o *OnenotePage) HasContentUrl() bool {
	if o != nil && o.ContentUrl.IsSet() {
		return true
	}

	return false
}

// SetContentUrl gets a reference to the given NullableString and assigns it to the ContentUrl field.
func (o *OnenotePage) SetContentUrl(v string) {
	o.ContentUrl.Set(&v)
}
// SetContentUrlNil sets the value for ContentUrl to be an explicit nil
func (o *OnenotePage) SetContentUrlNil() {
	o.ContentUrl.Set(nil)
}

// UnsetContentUrl ensures that no value is present for ContentUrl, not even an explicit nil
func (o *OnenotePage) UnsetContentUrl() {
	o.ContentUrl.Unset()
}

// GetCreatedByAppId returns the CreatedByAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenotePage) GetCreatedByAppId() string {
	if o == nil || o.CreatedByAppId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedByAppId.Get()
}

// GetCreatedByAppIdOk returns a tuple with the CreatedByAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenotePage) GetCreatedByAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedByAppId.Get(), o.CreatedByAppId.IsSet()
}

// HasCreatedByAppId returns a boolean if a field has been set.
func (o *OnenotePage) HasCreatedByAppId() bool {
	if o != nil && o.CreatedByAppId.IsSet() {
		return true
	}

	return false
}

// SetCreatedByAppId gets a reference to the given NullableString and assigns it to the CreatedByAppId field.
func (o *OnenotePage) SetCreatedByAppId(v string) {
	o.CreatedByAppId.Set(&v)
}
// SetCreatedByAppIdNil sets the value for CreatedByAppId to be an explicit nil
func (o *OnenotePage) SetCreatedByAppIdNil() {
	o.CreatedByAppId.Set(nil)
}

// UnsetCreatedByAppId ensures that no value is present for CreatedByAppId, not even an explicit nil
func (o *OnenotePage) UnsetCreatedByAppId() {
	o.CreatedByAppId.Unset()
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenotePage) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenotePage) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *OnenotePage) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *OnenotePage) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *OnenotePage) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *OnenotePage) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenotePage) GetLevel() int32 {
	if o == nil || o.Level.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenotePage) GetLevelOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *OnenotePage) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableInt32 and assigns it to the Level field.
func (o *OnenotePage) SetLevel(v int32) {
	o.Level.Set(&v)
}
// SetLevelNil sets the value for Level to be an explicit nil
func (o *OnenotePage) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *OnenotePage) UnsetLevel() {
	o.Level.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenotePage) GetLinks() AnyOfmicrosoftGraphPageLinks {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPageLinks
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenotePage) GetLinksOk() (*AnyOfmicrosoftGraphPageLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OnenotePage) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AnyOfmicrosoftGraphPageLinks and assigns it to the Links field.
func (o *OnenotePage) SetLinks(v AnyOfmicrosoftGraphPageLinks) {
	o.Links = v
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenotePage) GetOrder() int32 {
	if o == nil || o.Order.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenotePage) GetOrderOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *OnenotePage) HasOrder() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableInt32 and assigns it to the Order field.
func (o *OnenotePage) SetOrder(v int32) {
	o.Order.Set(&v)
}
// SetOrderNil sets the value for Order to be an explicit nil
func (o *OnenotePage) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *OnenotePage) UnsetOrder() {
	o.Order.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenotePage) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenotePage) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *OnenotePage) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *OnenotePage) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *OnenotePage) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *OnenotePage) UnsetTitle() {
	o.Title.Unset()
}

// GetUserTags returns the UserTags field value if set, zero value otherwise.
func (o *OnenotePage) GetUserTags() []*string {
	if o == nil || o.UserTags == nil {
		var ret []*string
		return ret
	}
	return *o.UserTags
}

// GetUserTagsOk returns a tuple with the UserTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetUserTagsOk() (*[]*string, bool) {
	if o == nil || o.UserTags == nil {
		return nil, false
	}
	return o.UserTags, true
}

// HasUserTags returns a boolean if a field has been set.
func (o *OnenotePage) HasUserTags() bool {
	if o != nil && o.UserTags != nil {
		return true
	}

	return false
}

// SetUserTags gets a reference to the given []*string and assigns it to the UserTags field.
func (o *OnenotePage) SetUserTags(v []*string) {
	o.UserTags = &v
}

// GetParentNotebook returns the ParentNotebook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenotePage) GetParentNotebook() AnyOfmicrosoftGraphNotebook {
	if o == nil  {
		var ret AnyOfmicrosoftGraphNotebook
		return ret
	}
	return o.ParentNotebook
}

// GetParentNotebookOk returns a tuple with the ParentNotebook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenotePage) GetParentNotebookOk() (*AnyOfmicrosoftGraphNotebook, bool) {
	if o == nil || o.ParentNotebook == nil {
		return nil, false
	}
	return &o.ParentNotebook, true
}

// HasParentNotebook returns a boolean if a field has been set.
func (o *OnenotePage) HasParentNotebook() bool {
	if o != nil && o.ParentNotebook != nil {
		return true
	}

	return false
}

// SetParentNotebook gets a reference to the given AnyOfmicrosoftGraphNotebook and assigns it to the ParentNotebook field.
func (o *OnenotePage) SetParentNotebook(v AnyOfmicrosoftGraphNotebook) {
	o.ParentNotebook = v
}

// GetParentSection returns the ParentSection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenotePage) GetParentSection() AnyOfmicrosoftGraphOnenoteSection {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOnenoteSection
		return ret
	}
	return o.ParentSection
}

// GetParentSectionOk returns a tuple with the ParentSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenotePage) GetParentSectionOk() (*AnyOfmicrosoftGraphOnenoteSection, bool) {
	if o == nil || o.ParentSection == nil {
		return nil, false
	}
	return &o.ParentSection, true
}

// HasParentSection returns a boolean if a field has been set.
func (o *OnenotePage) HasParentSection() bool {
	if o != nil && o.ParentSection != nil {
		return true
	}

	return false
}

// SetParentSection gets a reference to the given AnyOfmicrosoftGraphOnenoteSection and assigns it to the ParentSection field.
func (o *OnenotePage) SetParentSection(v AnyOfmicrosoftGraphOnenoteSection) {
	o.ParentSection = v
}

func (o OnenotePage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.ContentUrl.IsSet() {
		toSerialize["contentUrl"] = o.ContentUrl.Get()
	}
	if o.CreatedByAppId.IsSet() {
		toSerialize["createdByAppId"] = o.CreatedByAppId.Get()
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Order.IsSet() {
		toSerialize["order"] = o.Order.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.UserTags != nil {
		toSerialize["userTags"] = o.UserTags
	}
	if o.ParentNotebook != nil {
		toSerialize["parentNotebook"] = o.ParentNotebook
	}
	if o.ParentSection != nil {
		toSerialize["parentSection"] = o.ParentSection
	}
	return json.Marshal(toSerialize)
}

type NullableOnenotePage struct {
	value *OnenotePage
	isSet bool
}

func (v NullableOnenotePage) Get() *OnenotePage {
	return v.value
}

func (v *NullableOnenotePage) Set(val *OnenotePage) {
	v.value = val
	v.isSet = true
}

func (v NullableOnenotePage) IsSet() bool {
	return v.isSet
}

func (v *NullableOnenotePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnenotePage(val *OnenotePage) *NullableOnenotePage {
	return &NullableOnenotePage{value: val, isSet: true}
}

func (v NullableOnenotePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnenotePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


