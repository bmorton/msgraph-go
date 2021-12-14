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

// MicrosoftGraphSecureScoreControlProfile struct for MicrosoftGraphSecureScoreControlProfile
type MicrosoftGraphSecureScoreControlProfile struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Control action type (Config, Review, Behavior).
	ActionType NullableString `json:"actionType,omitempty"`
	// URL to where the control can be actioned.
	ActionUrl NullableString `json:"actionUrl,omitempty"`
	// GUID string for tenant ID.
	AzureTenantId *string `json:"azureTenantId,omitempty"`
	// The collection of compliance information associated with secure score control
	ComplianceInformation *[]*AnyOfmicrosoftGraphComplianceInformation `json:"complianceInformation,omitempty"`
	// Control action category (Identity, Data, Device, Apps, Infrastructure).
	ControlCategory NullableString `json:"controlCategory,omitempty"`
	// Flag to indicate where the tenant has marked a control (ignore, thirdParty, reviewed) (supports update).
	ControlStateUpdates *[]*AnyOfmicrosoftGraphSecureScoreControlStateUpdate `json:"controlStateUpdates,omitempty"`
	// Flag to indicate if a control is depreciated.
	Deprecated NullableBool `json:"deprecated,omitempty"`
	// Resource cost of implemmentating control (low, moderate, high).
	ImplementationCost NullableString `json:"implementationCost,omitempty"`
	// Time at which the control profile entity was last modified. The Timestamp type represents date and time
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// max attainable score for the control.
	MaxScore AnyOfnumberstringstring `json:"maxScore,omitempty"`
	// Microsoft's stack ranking of control.
	Rank NullableInt32 `json:"rank,omitempty"`
	// Description of what the control will help remediate.
	Remediation NullableString `json:"remediation,omitempty"`
	// Description of the impact on users of the remediation.
	RemediationImpact NullableString `json:"remediationImpact,omitempty"`
	// Service that owns the control (Exchange, Sharepoint, Azure AD).
	Service NullableString `json:"service,omitempty"`
	// List of threats the control mitigates (accountBreach,dataDeletion,dataExfiltration,dataSpillage,
	Threats *[]*string `json:"threats,omitempty"`
	// Control tier (Core, Defense in Depth, Advanced.)
	Tier NullableString `json:"tier,omitempty"`
	// Title of the control.
	Title NullableString `json:"title,omitempty"`
	// User impact of implementing control (low, moderate, high).
	UserImpact NullableString `json:"userImpact,omitempty"`
	VendorInformation AnyOfmicrosoftGraphSecurityVendorInformation `json:"vendorInformation,omitempty"`
}

// NewMicrosoftGraphSecureScoreControlProfile instantiates a new MicrosoftGraphSecureScoreControlProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSecureScoreControlProfile() *MicrosoftGraphSecureScoreControlProfile {
	this := MicrosoftGraphSecureScoreControlProfile{}
	return &this
}

// NewMicrosoftGraphSecureScoreControlProfileWithDefaults instantiates a new MicrosoftGraphSecureScoreControlProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSecureScoreControlProfileWithDefaults() *MicrosoftGraphSecureScoreControlProfile {
	this := MicrosoftGraphSecureScoreControlProfile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphSecureScoreControlProfile) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetId(v string) {
	o.Id = &v
}

// GetActionType returns the ActionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetActionType() string {
	if o == nil || o.ActionType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActionType.Get()
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetActionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActionType.Get(), o.ActionType.IsSet()
}

// HasActionType returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasActionType() bool {
	if o != nil && o.ActionType.IsSet() {
		return true
	}

	return false
}

// SetActionType gets a reference to the given NullableString and assigns it to the ActionType field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetActionType(v string) {
	o.ActionType.Set(&v)
}
// SetActionTypeNil sets the value for ActionType to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetActionTypeNil() {
	o.ActionType.Set(nil)
}

// UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetActionType() {
	o.ActionType.Unset()
}

// GetActionUrl returns the ActionUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetActionUrl() string {
	if o == nil || o.ActionUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActionUrl.Get()
}

// GetActionUrlOk returns a tuple with the ActionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetActionUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActionUrl.Get(), o.ActionUrl.IsSet()
}

// HasActionUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasActionUrl() bool {
	if o != nil && o.ActionUrl.IsSet() {
		return true
	}

	return false
}

// SetActionUrl gets a reference to the given NullableString and assigns it to the ActionUrl field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetActionUrl(v string) {
	o.ActionUrl.Set(&v)
}
// SetActionUrlNil sets the value for ActionUrl to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetActionUrlNil() {
	o.ActionUrl.Set(nil)
}

// UnsetActionUrl ensures that no value is present for ActionUrl, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetActionUrl() {
	o.ActionUrl.Unset()
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise.
func (o *MicrosoftGraphSecureScoreControlProfile) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) GetAzureTenantIdOk() (*string, bool) {
	if o == nil || o.AzureTenantId == nil {
		return nil, false
	}
	return o.AzureTenantId, true
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId != nil {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetComplianceInformation returns the ComplianceInformation field value if set, zero value otherwise.
func (o *MicrosoftGraphSecureScoreControlProfile) GetComplianceInformation() []*AnyOfmicrosoftGraphComplianceInformation {
	if o == nil || o.ComplianceInformation == nil {
		var ret []*AnyOfmicrosoftGraphComplianceInformation
		return ret
	}
	return *o.ComplianceInformation
}

// GetComplianceInformationOk returns a tuple with the ComplianceInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) GetComplianceInformationOk() (*[]*AnyOfmicrosoftGraphComplianceInformation, bool) {
	if o == nil || o.ComplianceInformation == nil {
		return nil, false
	}
	return o.ComplianceInformation, true
}

// HasComplianceInformation returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasComplianceInformation() bool {
	if o != nil && o.ComplianceInformation != nil {
		return true
	}

	return false
}

// SetComplianceInformation gets a reference to the given []*AnyOfmicrosoftGraphComplianceInformation and assigns it to the ComplianceInformation field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetComplianceInformation(v []*AnyOfmicrosoftGraphComplianceInformation) {
	o.ComplianceInformation = &v
}

// GetControlCategory returns the ControlCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetControlCategory() string {
	if o == nil || o.ControlCategory.Get() == nil {
		var ret string
		return ret
	}
	return *o.ControlCategory.Get()
}

// GetControlCategoryOk returns a tuple with the ControlCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetControlCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ControlCategory.Get(), o.ControlCategory.IsSet()
}

// HasControlCategory returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasControlCategory() bool {
	if o != nil && o.ControlCategory.IsSet() {
		return true
	}

	return false
}

// SetControlCategory gets a reference to the given NullableString and assigns it to the ControlCategory field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetControlCategory(v string) {
	o.ControlCategory.Set(&v)
}
// SetControlCategoryNil sets the value for ControlCategory to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetControlCategoryNil() {
	o.ControlCategory.Set(nil)
}

// UnsetControlCategory ensures that no value is present for ControlCategory, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetControlCategory() {
	o.ControlCategory.Unset()
}

// GetControlStateUpdates returns the ControlStateUpdates field value if set, zero value otherwise.
func (o *MicrosoftGraphSecureScoreControlProfile) GetControlStateUpdates() []*AnyOfmicrosoftGraphSecureScoreControlStateUpdate {
	if o == nil || o.ControlStateUpdates == nil {
		var ret []*AnyOfmicrosoftGraphSecureScoreControlStateUpdate
		return ret
	}
	return *o.ControlStateUpdates
}

// GetControlStateUpdatesOk returns a tuple with the ControlStateUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) GetControlStateUpdatesOk() (*[]*AnyOfmicrosoftGraphSecureScoreControlStateUpdate, bool) {
	if o == nil || o.ControlStateUpdates == nil {
		return nil, false
	}
	return o.ControlStateUpdates, true
}

// HasControlStateUpdates returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasControlStateUpdates() bool {
	if o != nil && o.ControlStateUpdates != nil {
		return true
	}

	return false
}

// SetControlStateUpdates gets a reference to the given []*AnyOfmicrosoftGraphSecureScoreControlStateUpdate and assigns it to the ControlStateUpdates field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetControlStateUpdates(v []*AnyOfmicrosoftGraphSecureScoreControlStateUpdate) {
	o.ControlStateUpdates = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetDeprecated() bool {
	if o == nil || o.Deprecated.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Deprecated.Get()
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetDeprecatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Deprecated.Get(), o.Deprecated.IsSet()
}

// HasDeprecated returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasDeprecated() bool {
	if o != nil && o.Deprecated.IsSet() {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given NullableBool and assigns it to the Deprecated field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetDeprecated(v bool) {
	o.Deprecated.Set(&v)
}
// SetDeprecatedNil sets the value for Deprecated to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetDeprecatedNil() {
	o.Deprecated.Set(nil)
}

// UnsetDeprecated ensures that no value is present for Deprecated, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetDeprecated() {
	o.Deprecated.Unset()
}

// GetImplementationCost returns the ImplementationCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetImplementationCost() string {
	if o == nil || o.ImplementationCost.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImplementationCost.Get()
}

// GetImplementationCostOk returns a tuple with the ImplementationCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetImplementationCostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ImplementationCost.Get(), o.ImplementationCost.IsSet()
}

// HasImplementationCost returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasImplementationCost() bool {
	if o != nil && o.ImplementationCost.IsSet() {
		return true
	}

	return false
}

// SetImplementationCost gets a reference to the given NullableString and assigns it to the ImplementationCost field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetImplementationCost(v string) {
	o.ImplementationCost.Set(&v)
}
// SetImplementationCostNil sets the value for ImplementationCost to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetImplementationCostNil() {
	o.ImplementationCost.Set(nil)
}

// UnsetImplementationCost ensures that no value is present for ImplementationCost, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetImplementationCost() {
	o.ImplementationCost.Unset()
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetMaxScore returns the MaxScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetMaxScore() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetMaxScoreOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.MaxScore == nil {
		return nil, false
	}
	return &o.MaxScore, true
}

// HasMaxScore returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasMaxScore() bool {
	if o != nil && o.MaxScore != nil {
		return true
	}

	return false
}

// SetMaxScore gets a reference to the given AnyOfnumberstringstring and assigns it to the MaxScore field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetMaxScore(v AnyOfnumberstringstring) {
	o.MaxScore = v
}

// GetRank returns the Rank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetRank() int32 {
	if o == nil || o.Rank.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Rank.Get()
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetRankOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Rank.Get(), o.Rank.IsSet()
}

// HasRank returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasRank() bool {
	if o != nil && o.Rank.IsSet() {
		return true
	}

	return false
}

// SetRank gets a reference to the given NullableInt32 and assigns it to the Rank field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetRank(v int32) {
	o.Rank.Set(&v)
}
// SetRankNil sets the value for Rank to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetRankNil() {
	o.Rank.Set(nil)
}

// UnsetRank ensures that no value is present for Rank, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetRank() {
	o.Rank.Unset()
}

// GetRemediation returns the Remediation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetRemediation() string {
	if o == nil || o.Remediation.Get() == nil {
		var ret string
		return ret
	}
	return *o.Remediation.Get()
}

// GetRemediationOk returns a tuple with the Remediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetRemediationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Remediation.Get(), o.Remediation.IsSet()
}

// HasRemediation returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasRemediation() bool {
	if o != nil && o.Remediation.IsSet() {
		return true
	}

	return false
}

// SetRemediation gets a reference to the given NullableString and assigns it to the Remediation field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetRemediation(v string) {
	o.Remediation.Set(&v)
}
// SetRemediationNil sets the value for Remediation to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetRemediationNil() {
	o.Remediation.Set(nil)
}

// UnsetRemediation ensures that no value is present for Remediation, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetRemediation() {
	o.Remediation.Unset()
}

// GetRemediationImpact returns the RemediationImpact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetRemediationImpact() string {
	if o == nil || o.RemediationImpact.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemediationImpact.Get()
}

// GetRemediationImpactOk returns a tuple with the RemediationImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetRemediationImpactOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemediationImpact.Get(), o.RemediationImpact.IsSet()
}

// HasRemediationImpact returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasRemediationImpact() bool {
	if o != nil && o.RemediationImpact.IsSet() {
		return true
	}

	return false
}

// SetRemediationImpact gets a reference to the given NullableString and assigns it to the RemediationImpact field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetRemediationImpact(v string) {
	o.RemediationImpact.Set(&v)
}
// SetRemediationImpactNil sets the value for RemediationImpact to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetRemediationImpactNil() {
	o.RemediationImpact.Set(nil)
}

// UnsetRemediationImpact ensures that no value is present for RemediationImpact, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetRemediationImpact() {
	o.RemediationImpact.Unset()
}

// GetService returns the Service field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetService() string {
	if o == nil || o.Service.Get() == nil {
		var ret string
		return ret
	}
	return *o.Service.Get()
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Service.Get(), o.Service.IsSet()
}

// HasService returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasService() bool {
	if o != nil && o.Service.IsSet() {
		return true
	}

	return false
}

// SetService gets a reference to the given NullableString and assigns it to the Service field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetService(v string) {
	o.Service.Set(&v)
}
// SetServiceNil sets the value for Service to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetServiceNil() {
	o.Service.Set(nil)
}

// UnsetService ensures that no value is present for Service, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetService() {
	o.Service.Unset()
}

// GetThreats returns the Threats field value if set, zero value otherwise.
func (o *MicrosoftGraphSecureScoreControlProfile) GetThreats() []*string {
	if o == nil || o.Threats == nil {
		var ret []*string
		return ret
	}
	return *o.Threats
}

// GetThreatsOk returns a tuple with the Threats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) GetThreatsOk() (*[]*string, bool) {
	if o == nil || o.Threats == nil {
		return nil, false
	}
	return o.Threats, true
}

// HasThreats returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasThreats() bool {
	if o != nil && o.Threats != nil {
		return true
	}

	return false
}

// SetThreats gets a reference to the given []*string and assigns it to the Threats field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetThreats(v []*string) {
	o.Threats = &v
}

// GetTier returns the Tier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetTier() string {
	if o == nil || o.Tier.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tier.Get()
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetTierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Tier.Get(), o.Tier.IsSet()
}

// HasTier returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasTier() bool {
	if o != nil && o.Tier.IsSet() {
		return true
	}

	return false
}

// SetTier gets a reference to the given NullableString and assigns it to the Tier field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetTier(v string) {
	o.Tier.Set(&v)
}
// SetTierNil sets the value for Tier to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetTierNil() {
	o.Tier.Set(nil)
}

// UnsetTier ensures that no value is present for Tier, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetTier() {
	o.Tier.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetTitle() {
	o.Title.Unset()
}

// GetUserImpact returns the UserImpact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetUserImpact() string {
	if o == nil || o.UserImpact.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserImpact.Get()
}

// GetUserImpactOk returns a tuple with the UserImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetUserImpactOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserImpact.Get(), o.UserImpact.IsSet()
}

// HasUserImpact returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasUserImpact() bool {
	if o != nil && o.UserImpact.IsSet() {
		return true
	}

	return false
}

// SetUserImpact gets a reference to the given NullableString and assigns it to the UserImpact field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetUserImpact(v string) {
	o.UserImpact.Set(&v)
}
// SetUserImpactNil sets the value for UserImpact to be an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) SetUserImpactNil() {
	o.UserImpact.Set(nil)
}

// UnsetUserImpact ensures that no value is present for UserImpact, not even an explicit nil
func (o *MicrosoftGraphSecureScoreControlProfile) UnsetUserImpact() {
	o.UserImpact.Unset()
}

// GetVendorInformation returns the VendorInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSecureScoreControlProfile) GetVendorInformation() AnyOfmicrosoftGraphSecurityVendorInformation {
	if o == nil  {
		var ret AnyOfmicrosoftGraphSecurityVendorInformation
		return ret
	}
	return o.VendorInformation
}

// GetVendorInformationOk returns a tuple with the VendorInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSecureScoreControlProfile) GetVendorInformationOk() (*AnyOfmicrosoftGraphSecurityVendorInformation, bool) {
	if o == nil || o.VendorInformation == nil {
		return nil, false
	}
	return &o.VendorInformation, true
}

// HasVendorInformation returns a boolean if a field has been set.
func (o *MicrosoftGraphSecureScoreControlProfile) HasVendorInformation() bool {
	if o != nil && o.VendorInformation != nil {
		return true
	}

	return false
}

// SetVendorInformation gets a reference to the given AnyOfmicrosoftGraphSecurityVendorInformation and assigns it to the VendorInformation field.
func (o *MicrosoftGraphSecureScoreControlProfile) SetVendorInformation(v AnyOfmicrosoftGraphSecurityVendorInformation) {
	o.VendorInformation = v
}

func (o MicrosoftGraphSecureScoreControlProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ActionType.IsSet() {
		toSerialize["actionType"] = o.ActionType.Get()
	}
	if o.ActionUrl.IsSet() {
		toSerialize["actionUrl"] = o.ActionUrl.Get()
	}
	if o.AzureTenantId != nil {
		toSerialize["azureTenantId"] = o.AzureTenantId
	}
	if o.ComplianceInformation != nil {
		toSerialize["complianceInformation"] = o.ComplianceInformation
	}
	if o.ControlCategory.IsSet() {
		toSerialize["controlCategory"] = o.ControlCategory.Get()
	}
	if o.ControlStateUpdates != nil {
		toSerialize["controlStateUpdates"] = o.ControlStateUpdates
	}
	if o.Deprecated.IsSet() {
		toSerialize["deprecated"] = o.Deprecated.Get()
	}
	if o.ImplementationCost.IsSet() {
		toSerialize["implementationCost"] = o.ImplementationCost.Get()
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.MaxScore != nil {
		toSerialize["maxScore"] = o.MaxScore
	}
	if o.Rank.IsSet() {
		toSerialize["rank"] = o.Rank.Get()
	}
	if o.Remediation.IsSet() {
		toSerialize["remediation"] = o.Remediation.Get()
	}
	if o.RemediationImpact.IsSet() {
		toSerialize["remediationImpact"] = o.RemediationImpact.Get()
	}
	if o.Service.IsSet() {
		toSerialize["service"] = o.Service.Get()
	}
	if o.Threats != nil {
		toSerialize["threats"] = o.Threats
	}
	if o.Tier.IsSet() {
		toSerialize["tier"] = o.Tier.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.UserImpact.IsSet() {
		toSerialize["userImpact"] = o.UserImpact.Get()
	}
	if o.VendorInformation != nil {
		toSerialize["vendorInformation"] = o.VendorInformation
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSecureScoreControlProfile struct {
	value *MicrosoftGraphSecureScoreControlProfile
	isSet bool
}

func (v NullableMicrosoftGraphSecureScoreControlProfile) Get() *MicrosoftGraphSecureScoreControlProfile {
	return v.value
}

func (v *NullableMicrosoftGraphSecureScoreControlProfile) Set(val *MicrosoftGraphSecureScoreControlProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSecureScoreControlProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSecureScoreControlProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSecureScoreControlProfile(val *MicrosoftGraphSecureScoreControlProfile) *NullableMicrosoftGraphSecureScoreControlProfile {
	return &NullableMicrosoftGraphSecureScoreControlProfile{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSecureScoreControlProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSecureScoreControlProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


