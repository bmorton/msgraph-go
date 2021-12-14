/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphList struct for MicrosoftGraphList
type MicrosoftGraphList struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Identity of the user, device, or application which created the item. Read-only.
	CreatedBy AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Provides a user-visible description of the item. Optional.
	Description NullableString `json:"description,omitempty"`
	// ETag for the item. Read-only.
	ETag NullableString `json:"eTag,omitempty"`
	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the item. Read-write.
	Name NullableString `json:"name,omitempty"`
	// Parent information, if the item has a parent. Read-write.
	ParentReference AnyOfmicrosoftGraphItemReference `json:"parentReference,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl NullableString `json:"webUrl,omitempty"`
	// Identity of the user who created the item. Read-only.
	CreatedByUser AnyOfmicrosoftGraphUser `json:"createdByUser,omitempty"`
	// Identity of the user who last modified the item. Read-only.
	LastModifiedByUser AnyOfmicrosoftGraphUser `json:"lastModifiedByUser,omitempty"`
	// The displayable title of the list.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Provides additional details about the list.
	List AnyOfmicrosoftGraphListInfo `json:"list,omitempty"`
	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds AnyOfmicrosoftGraphSharepointIds `json:"sharepointIds,omitempty"`
	// If present, indicates that this is a system-managed list. Read-only.
	System AnyOfobject `json:"system,omitempty"`
	// The collection of field definitions for this list.
	Columns *[]MicrosoftGraphColumnDefinition `json:"columns,omitempty"`
	// The collection of content types present in this list.
	ContentTypes *[]MicrosoftGraphContentType `json:"contentTypes,omitempty"`
	// Only present on document libraries. Allows access to the list as a [drive][] resource with [driveItems][driveItem].
	Drive AnyOfmicrosoftGraphDrive `json:"drive,omitempty"`
	// All items contained in the list.
	Items *[]MicrosoftGraphListItem `json:"items,omitempty"`
	// The set of subscriptions on the list.
	Subscriptions *[]MicrosoftGraphSubscription `json:"subscriptions,omitempty"`
}

// NewMicrosoftGraphList instantiates a new MicrosoftGraphList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphList() *MicrosoftGraphList {
	this := MicrosoftGraphList{}
	return &this
}

// NewMicrosoftGraphListWithDefaults instantiates a new MicrosoftGraphList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphListWithDefaults() *MicrosoftGraphList {
	this := MicrosoftGraphList{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphList) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphList) SetId(v string) {
	o.Id = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphList) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphList) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphList) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphList) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphList) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphList) UnsetDescription() {
	o.Description.Unset()
}

// GetETag returns the ETag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetETag() string {
	if o == nil || o.ETag.Get() == nil {
		var ret string
		return ret
	}
	return *o.ETag.Get()
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetETagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ETag.Get(), o.ETag.IsSet()
}

// HasETag returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasETag() bool {
	if o != nil && o.ETag.IsSet() {
		return true
	}

	return false
}

// SetETag gets a reference to the given NullableString and assigns it to the ETag field.
func (o *MicrosoftGraphList) SetETag(v string) {
	o.ETag.Set(&v)
}
// SetETagNil sets the value for ETag to be an explicit nil
func (o *MicrosoftGraphList) SetETagNil() {
	o.ETag.Set(nil)
}

// UnsetETag ensures that no value is present for ETag, not even an explicit nil
func (o *MicrosoftGraphList) UnsetETag() {
	o.ETag.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphList) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphList) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphList) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphList) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphList) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphList) UnsetName() {
	o.Name.Unset()
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetParentReference() AnyOfmicrosoftGraphItemReference {
	if o == nil  {
		var ret AnyOfmicrosoftGraphItemReference
		return ret
	}
	return o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetParentReferenceOk() (*AnyOfmicrosoftGraphItemReference, bool) {
	if o == nil || o.ParentReference == nil {
		return nil, false
	}
	return &o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasParentReference() bool {
	if o != nil && o.ParentReference != nil {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.
func (o *MicrosoftGraphList) SetParentReference(v AnyOfmicrosoftGraphItemReference) {
	o.ParentReference = v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetWebUrl() string {
	if o == nil || o.WebUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebUrl.Get()
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetWebUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WebUrl.Get(), o.WebUrl.IsSet()
}

// HasWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasWebUrl() bool {
	if o != nil && o.WebUrl.IsSet() {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given NullableString and assigns it to the WebUrl field.
func (o *MicrosoftGraphList) SetWebUrl(v string) {
	o.WebUrl.Set(&v)
}
// SetWebUrlNil sets the value for WebUrl to be an explicit nil
func (o *MicrosoftGraphList) SetWebUrlNil() {
	o.WebUrl.Set(nil)
}

// UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
func (o *MicrosoftGraphList) UnsetWebUrl() {
	o.WebUrl.Unset()
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetCreatedByUser() AnyOfmicrosoftGraphUser {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetCreatedByUserOk() (*AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return &o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.
func (o *MicrosoftGraphList) SetCreatedByUser(v AnyOfmicrosoftGraphUser) {
	o.CreatedByUser = v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetLastModifiedByUser() AnyOfmicrosoftGraphUser {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetLastModifiedByUserOk() (*AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return &o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser != nil {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.
func (o *MicrosoftGraphList) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser) {
	o.LastModifiedByUser = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphList) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphList) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphList) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetList returns the List field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetList() AnyOfmicrosoftGraphListInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphListInfo
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetListOk() (*AnyOfmicrosoftGraphListInfo, bool) {
	if o == nil || o.List == nil {
		return nil, false
	}
	return &o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasList() bool {
	if o != nil && o.List != nil {
		return true
	}

	return false
}

// SetList gets a reference to the given AnyOfmicrosoftGraphListInfo and assigns it to the List field.
func (o *MicrosoftGraphList) SetList(v AnyOfmicrosoftGraphListInfo) {
	o.List = v
}

// GetSharepointIds returns the SharepointIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds {
	if o == nil  {
		var ret AnyOfmicrosoftGraphSharepointIds
		return ret
	}
	return o.SharepointIds
}

// GetSharepointIdsOk returns a tuple with the SharepointIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetSharepointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool) {
	if o == nil || o.SharepointIds == nil {
		return nil, false
	}
	return &o.SharepointIds, true
}

// HasSharepointIds returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasSharepointIds() bool {
	if o != nil && o.SharepointIds != nil {
		return true
	}

	return false
}

// SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.
func (o *MicrosoftGraphList) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds) {
	o.SharepointIds = v
}

// GetSystem returns the System field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetSystem() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetSystemOk() (*AnyOfobject, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return &o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given AnyOfobject and assigns it to the System field.
func (o *MicrosoftGraphList) SetSystem(v AnyOfobject) {
	o.System = v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *MicrosoftGraphList) GetColumns() []MicrosoftGraphColumnDefinition {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphColumnDefinition
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the Columns field.
func (o *MicrosoftGraphList) SetColumns(v []MicrosoftGraphColumnDefinition) {
	o.Columns = &v
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise.
func (o *MicrosoftGraphList) GetContentTypes() []MicrosoftGraphContentType {
	if o == nil || o.ContentTypes == nil {
		var ret []MicrosoftGraphContentType
		return ret
	}
	return *o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetContentTypesOk() (*[]MicrosoftGraphContentType, bool) {
	if o == nil || o.ContentTypes == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasContentTypes() bool {
	if o != nil && o.ContentTypes != nil {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []MicrosoftGraphContentType and assigns it to the ContentTypes field.
func (o *MicrosoftGraphList) SetContentTypes(v []MicrosoftGraphContentType) {
	o.ContentTypes = &v
}

// GetDrive returns the Drive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphList) GetDrive() AnyOfmicrosoftGraphDrive {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDrive
		return ret
	}
	return o.Drive
}

// GetDriveOk returns a tuple with the Drive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphList) GetDriveOk() (*AnyOfmicrosoftGraphDrive, bool) {
	if o == nil || o.Drive == nil {
		return nil, false
	}
	return &o.Drive, true
}

// HasDrive returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasDrive() bool {
	if o != nil && o.Drive != nil {
		return true
	}

	return false
}

// SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.
func (o *MicrosoftGraphList) SetDrive(v AnyOfmicrosoftGraphDrive) {
	o.Drive = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *MicrosoftGraphList) GetItems() []MicrosoftGraphListItem {
	if o == nil || o.Items == nil {
		var ret []MicrosoftGraphListItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetItemsOk() (*[]MicrosoftGraphListItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MicrosoftGraphListItem and assigns it to the Items field.
func (o *MicrosoftGraphList) SetItems(v []MicrosoftGraphListItem) {
	o.Items = &v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *MicrosoftGraphList) GetSubscriptions() []MicrosoftGraphSubscription {
	if o == nil || o.Subscriptions == nil {
		var ret []MicrosoftGraphSubscription
		return ret
	}
	return *o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetSubscriptionsOk() (*[]MicrosoftGraphSubscription, bool) {
	if o == nil || o.Subscriptions == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasSubscriptions() bool {
	if o != nil && o.Subscriptions != nil {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []MicrosoftGraphSubscription and assigns it to the Subscriptions field.
func (o *MicrosoftGraphList) SetSubscriptions(v []MicrosoftGraphSubscription) {
	o.Subscriptions = &v
}

func (o MicrosoftGraphList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ETag.IsSet() {
		toSerialize["eTag"] = o.ETag.Get()
	}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ParentReference != nil {
		toSerialize["parentReference"] = o.ParentReference
	}
	if o.WebUrl.IsSet() {
		toSerialize["webUrl"] = o.WebUrl.Get()
	}
	if o.CreatedByUser != nil {
		toSerialize["createdByUser"] = o.CreatedByUser
	}
	if o.LastModifiedByUser != nil {
		toSerialize["lastModifiedByUser"] = o.LastModifiedByUser
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.List != nil {
		toSerialize["list"] = o.List
	}
	if o.SharepointIds != nil {
		toSerialize["sharepointIds"] = o.SharepointIds
	}
	if o.System != nil {
		toSerialize["system"] = o.System
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.ContentTypes != nil {
		toSerialize["contentTypes"] = o.ContentTypes
	}
	if o.Drive != nil {
		toSerialize["drive"] = o.Drive
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphList struct {
	value *MicrosoftGraphList
	isSet bool
}

func (v NullableMicrosoftGraphList) Get() *MicrosoftGraphList {
	return v.value
}

func (v *NullableMicrosoftGraphList) Set(val *MicrosoftGraphList) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphList) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphList(val *MicrosoftGraphList) *NullableMicrosoftGraphList {
	return &NullableMicrosoftGraphList{value: val, isSet: true}
}

func (v NullableMicrosoftGraphList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

