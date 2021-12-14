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

// MicrosoftGraphAccessPackageCatalog struct for MicrosoftGraphAccessPackageCatalog
type MicrosoftGraphAccessPackageCatalog struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue.
	CatalogType AnyOfmicrosoftGraphAccessPackageCatalogType `json:"catalogType,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// The description of the access package catalog.
	Description NullableString `json:"description,omitempty"`
	// The display name of the access package catalog.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Whether the access packages in this catalog can be requested by users outside of the tenant.
	IsExternallyVisible NullableBool `json:"isExternallyVisible,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ModifiedDateTime NullableTime `json:"modifiedDateTime,omitempty"`
	// Has the value published if the access packages are available for management. The possible values are: unpublished, published, unknownFutureValue.
	State AnyOfmicrosoftGraphAccessPackageCatalogState `json:"state,omitempty"`
	// The access packages in this catalog. Read-only. Nullable.
	AccessPackages *[]MicrosoftGraphAccessPackage `json:"accessPackages,omitempty"`
}

// NewMicrosoftGraphAccessPackageCatalog instantiates a new MicrosoftGraphAccessPackageCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAccessPackageCatalog() *MicrosoftGraphAccessPackageCatalog {
	this := MicrosoftGraphAccessPackageCatalog{}
	return &this
}

// NewMicrosoftGraphAccessPackageCatalogWithDefaults instantiates a new MicrosoftGraphAccessPackageCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAccessPackageCatalogWithDefaults() *MicrosoftGraphAccessPackageCatalog {
	this := MicrosoftGraphAccessPackageCatalog{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphAccessPackageCatalog) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAccessPackageCatalog) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageCatalog) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphAccessPackageCatalog) SetId(v string) {
	o.Id = &v
}

// GetCatalogType returns the CatalogType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageCatalog) GetCatalogType() AnyOfmicrosoftGraphAccessPackageCatalogType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAccessPackageCatalogType
		return ret
	}
	return o.CatalogType
}

// GetCatalogTypeOk returns a tuple with the CatalogType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageCatalog) GetCatalogTypeOk() (*AnyOfmicrosoftGraphAccessPackageCatalogType, bool) {
	if o == nil || o.CatalogType == nil {
		return nil, false
	}
	return &o.CatalogType, true
}

// HasCatalogType returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageCatalog) HasCatalogType() bool {
	if o != nil && o.CatalogType != nil {
		return true
	}

	return false
}

// SetCatalogType gets a reference to the given AnyOfmicrosoftGraphAccessPackageCatalogType and assigns it to the CatalogType field.
func (o *MicrosoftGraphAccessPackageCatalog) SetCatalogType(v AnyOfmicrosoftGraphAccessPackageCatalogType) {
	o.CatalogType = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageCatalog) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageCatalog) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageCatalog) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphAccessPackageCatalog) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphAccessPackageCatalog) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphAccessPackageCatalog) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageCatalog) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageCatalog) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageCatalog) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphAccessPackageCatalog) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphAccessPackageCatalog) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphAccessPackageCatalog) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageCatalog) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageCatalog) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageCatalog) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphAccessPackageCatalog) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphAccessPackageCatalog) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphAccessPackageCatalog) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetIsExternallyVisible returns the IsExternallyVisible field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageCatalog) GetIsExternallyVisible() bool {
	if o == nil || o.IsExternallyVisible.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsExternallyVisible.Get()
}

// GetIsExternallyVisibleOk returns a tuple with the IsExternallyVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageCatalog) GetIsExternallyVisibleOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsExternallyVisible.Get(), o.IsExternallyVisible.IsSet()
}

// HasIsExternallyVisible returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageCatalog) HasIsExternallyVisible() bool {
	if o != nil && o.IsExternallyVisible.IsSet() {
		return true
	}

	return false
}

// SetIsExternallyVisible gets a reference to the given NullableBool and assigns it to the IsExternallyVisible field.
func (o *MicrosoftGraphAccessPackageCatalog) SetIsExternallyVisible(v bool) {
	o.IsExternallyVisible.Set(&v)
}
// SetIsExternallyVisibleNil sets the value for IsExternallyVisible to be an explicit nil
func (o *MicrosoftGraphAccessPackageCatalog) SetIsExternallyVisibleNil() {
	o.IsExternallyVisible.Set(nil)
}

// UnsetIsExternallyVisible ensures that no value is present for IsExternallyVisible, not even an explicit nil
func (o *MicrosoftGraphAccessPackageCatalog) UnsetIsExternallyVisible() {
	o.IsExternallyVisible.Unset()
}

// GetModifiedDateTime returns the ModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageCatalog) GetModifiedDateTime() time.Time {
	if o == nil || o.ModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDateTime.Get()
}

// GetModifiedDateTimeOk returns a tuple with the ModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageCatalog) GetModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedDateTime.Get(), o.ModifiedDateTime.IsSet()
}

// HasModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageCatalog) HasModifiedDateTime() bool {
	if o != nil && o.ModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetModifiedDateTime gets a reference to the given NullableTime and assigns it to the ModifiedDateTime field.
func (o *MicrosoftGraphAccessPackageCatalog) SetModifiedDateTime(v time.Time) {
	o.ModifiedDateTime.Set(&v)
}
// SetModifiedDateTimeNil sets the value for ModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphAccessPackageCatalog) SetModifiedDateTimeNil() {
	o.ModifiedDateTime.Set(nil)
}

// UnsetModifiedDateTime ensures that no value is present for ModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphAccessPackageCatalog) UnsetModifiedDateTime() {
	o.ModifiedDateTime.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageCatalog) GetState() AnyOfmicrosoftGraphAccessPackageCatalogState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAccessPackageCatalogState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageCatalog) GetStateOk() (*AnyOfmicrosoftGraphAccessPackageCatalogState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageCatalog) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphAccessPackageCatalogState and assigns it to the State field.
func (o *MicrosoftGraphAccessPackageCatalog) SetState(v AnyOfmicrosoftGraphAccessPackageCatalogState) {
	o.State = v
}

// GetAccessPackages returns the AccessPackages field value if set, zero value otherwise.
func (o *MicrosoftGraphAccessPackageCatalog) GetAccessPackages() []MicrosoftGraphAccessPackage {
	if o == nil || o.AccessPackages == nil {
		var ret []MicrosoftGraphAccessPackage
		return ret
	}
	return *o.AccessPackages
}

// GetAccessPackagesOk returns a tuple with the AccessPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAccessPackageCatalog) GetAccessPackagesOk() (*[]MicrosoftGraphAccessPackage, bool) {
	if o == nil || o.AccessPackages == nil {
		return nil, false
	}
	return o.AccessPackages, true
}

// HasAccessPackages returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageCatalog) HasAccessPackages() bool {
	if o != nil && o.AccessPackages != nil {
		return true
	}

	return false
}

// SetAccessPackages gets a reference to the given []MicrosoftGraphAccessPackage and assigns it to the AccessPackages field.
func (o *MicrosoftGraphAccessPackageCatalog) SetAccessPackages(v []MicrosoftGraphAccessPackage) {
	o.AccessPackages = &v
}

func (o MicrosoftGraphAccessPackageCatalog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CatalogType != nil {
		toSerialize["catalogType"] = o.CatalogType
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.IsExternallyVisible.IsSet() {
		toSerialize["isExternallyVisible"] = o.IsExternallyVisible.Get()
	}
	if o.ModifiedDateTime.IsSet() {
		toSerialize["modifiedDateTime"] = o.ModifiedDateTime.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.AccessPackages != nil {
		toSerialize["accessPackages"] = o.AccessPackages
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAccessPackageCatalog struct {
	value *MicrosoftGraphAccessPackageCatalog
	isSet bool
}

func (v NullableMicrosoftGraphAccessPackageCatalog) Get() *MicrosoftGraphAccessPackageCatalog {
	return v.value
}

func (v *NullableMicrosoftGraphAccessPackageCatalog) Set(val *MicrosoftGraphAccessPackageCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAccessPackageCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAccessPackageCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAccessPackageCatalog(val *MicrosoftGraphAccessPackageCatalog) *NullableMicrosoftGraphAccessPackageCatalog {
	return &NullableMicrosoftGraphAccessPackageCatalog{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAccessPackageCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAccessPackageCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

