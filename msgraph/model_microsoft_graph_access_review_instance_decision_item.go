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

// MicrosoftGraphAccessReviewInstanceDecisionItem struct for MicrosoftGraphAccessReviewInstanceDecisionItem
type MicrosoftGraphAccessReviewInstanceDecisionItem struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The identifier of the accessReviewInstance parent. Supports $select. Read-only.
	AccessReviewId *string `json:"accessReviewId,omitempty"`
	// The identifier of the user who applied the decision. Read-only.
	AppliedBy AnyOfmicrosoftGraphUserIdentity `json:"appliedBy,omitempty"`
	// The timestamp when the approval decision was applied. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Supports $select. Read-only.
	AppliedDateTime NullableTime `json:"appliedDateTime,omitempty"`
	// The result of applying the decision. Possible values: New, AppliedSuccessfully, AppliedWithUnknownFailure, AppliedSuccessfullyButObjectNotFound and ApplyNotSupported. Supports $select, $orderby, and $filter (eq only). Read-only.
	ApplyResult NullableString `json:"applyResult,omitempty"`
	// Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow. Supports $select, $orderby, and $filter (eq only).
	Decision NullableString `json:"decision,omitempty"`
	// Justification left by the reviewer when they made the decision.
	Justification NullableString `json:"justification,omitempty"`
	// Every decision item in an access review represents a principal's access to a resource. This property represents details of the principal. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is 'Bob' and the resource is 'Sales'. Principals can be of two types - userIdentity and servicePrincipalIdentity. Supports $select. Read-only.
	Principal AnyOfmicrosoftGraphIdentity `json:"principal,omitempty"`
	// A link to the principal object. For example, https://graph.microsoft.com/v1.0/users/a6c7aecb-cbfd-4763-87ef-e91b4bd509d9. Read-only.
	PrincipalLink NullableString `json:"principalLink,omitempty"`
	// A system-generated recommendation for the approval decision based off last interactive sign-in to tenant. Recommend approve if sign-in is within thirty days of start of review. Recommend deny if sign-in is greater than thirty days of start of review. Recommendation not available otherwise. Possible values: Approve, Deny, or NoInfoAvailable. Supports $select, $orderby, and $filter (eq only). Read-only.
	Recommendation NullableString `json:"recommendation,omitempty"`
	// Every decision item in an access review represents a principal's access to a resource. This property represents details of the resource. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is Bob and the resource is 'Sales'. Resources can be of multiple types. See accessReviewInstanceDecisionItemResource. Read-only.
	Resource AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource `json:"resource,omitempty"`
	// A link to the resource. For example, https://graph.microsoft.com/v1.0/servicePrincipals/c86300f3-8695-4320-9f6e-32a2555f5ff8. Supports $select. Read-only.
	ResourceLink NullableString `json:"resourceLink,omitempty"`
	// The identifier of the reviewer. Supports $select. Read-only.
	ReviewedBy AnyOfmicrosoftGraphUserIdentity `json:"reviewedBy,omitempty"`
	// The timestamp when the review decision occurred. Supports $select. Read-only.
	ReviewedDateTime NullableTime `json:"reviewedDateTime,omitempty"`
}

// NewMicrosoftGraphAccessReviewInstanceDecisionItem instantiates a new MicrosoftGraphAccessReviewInstanceDecisionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAccessReviewInstanceDecisionItem() *MicrosoftGraphAccessReviewInstanceDecisionItem {
	this := MicrosoftGraphAccessReviewInstanceDecisionItem{}
	return &this
}

// NewMicrosoftGraphAccessReviewInstanceDecisionItemWithDefaults instantiates a new MicrosoftGraphAccessReviewInstanceDecisionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAccessReviewInstanceDecisionItemWithDefaults() *MicrosoftGraphAccessReviewInstanceDecisionItem {
	this := MicrosoftGraphAccessReviewInstanceDecisionItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetId(v string) {
	o.Id = &v
}

// GetAccessReviewId returns the AccessReviewId field value if set, zero value otherwise.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAccessReviewId() string {
	if o == nil || o.AccessReviewId == nil {
		var ret string
		return ret
	}
	return *o.AccessReviewId
}

// GetAccessReviewIdOk returns a tuple with the AccessReviewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAccessReviewIdOk() (*string, bool) {
	if o == nil || o.AccessReviewId == nil {
		return nil, false
	}
	return o.AccessReviewId, true
}

// HasAccessReviewId returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasAccessReviewId() bool {
	if o != nil && o.AccessReviewId != nil {
		return true
	}

	return false
}

// SetAccessReviewId gets a reference to the given string and assigns it to the AccessReviewId field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetAccessReviewId(v string) {
	o.AccessReviewId = &v
}

// GetAppliedBy returns the AppliedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAppliedBy() AnyOfmicrosoftGraphUserIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUserIdentity
		return ret
	}
	return o.AppliedBy
}

// GetAppliedByOk returns a tuple with the AppliedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAppliedByOk() (*AnyOfmicrosoftGraphUserIdentity, bool) {
	if o == nil || o.AppliedBy == nil {
		return nil, false
	}
	return &o.AppliedBy, true
}

// HasAppliedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasAppliedBy() bool {
	if o != nil && o.AppliedBy != nil {
		return true
	}

	return false
}

// SetAppliedBy gets a reference to the given AnyOfmicrosoftGraphUserIdentity and assigns it to the AppliedBy field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetAppliedBy(v AnyOfmicrosoftGraphUserIdentity) {
	o.AppliedBy = v
}

// GetAppliedDateTime returns the AppliedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAppliedDateTime() time.Time {
	if o == nil || o.AppliedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AppliedDateTime.Get()
}

// GetAppliedDateTimeOk returns a tuple with the AppliedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAppliedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppliedDateTime.Get(), o.AppliedDateTime.IsSet()
}

// HasAppliedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasAppliedDateTime() bool {
	if o != nil && o.AppliedDateTime.IsSet() {
		return true
	}

	return false
}

// SetAppliedDateTime gets a reference to the given NullableTime and assigns it to the AppliedDateTime field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetAppliedDateTime(v time.Time) {
	o.AppliedDateTime.Set(&v)
}
// SetAppliedDateTimeNil sets the value for AppliedDateTime to be an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetAppliedDateTimeNil() {
	o.AppliedDateTime.Set(nil)
}

// UnsetAppliedDateTime ensures that no value is present for AppliedDateTime, not even an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetAppliedDateTime() {
	o.AppliedDateTime.Unset()
}

// GetApplyResult returns the ApplyResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetApplyResult() string {
	if o == nil || o.ApplyResult.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplyResult.Get()
}

// GetApplyResultOk returns a tuple with the ApplyResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetApplyResultOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplyResult.Get(), o.ApplyResult.IsSet()
}

// HasApplyResult returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasApplyResult() bool {
	if o != nil && o.ApplyResult.IsSet() {
		return true
	}

	return false
}

// SetApplyResult gets a reference to the given NullableString and assigns it to the ApplyResult field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetApplyResult(v string) {
	o.ApplyResult.Set(&v)
}
// SetApplyResultNil sets the value for ApplyResult to be an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetApplyResultNil() {
	o.ApplyResult.Set(nil)
}

// UnsetApplyResult ensures that no value is present for ApplyResult, not even an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetApplyResult() {
	o.ApplyResult.Unset()
}

// GetDecision returns the Decision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetDecision() string {
	if o == nil || o.Decision.Get() == nil {
		var ret string
		return ret
	}
	return *o.Decision.Get()
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetDecisionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Decision.Get(), o.Decision.IsSet()
}

// HasDecision returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasDecision() bool {
	if o != nil && o.Decision.IsSet() {
		return true
	}

	return false
}

// SetDecision gets a reference to the given NullableString and assigns it to the Decision field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetDecision(v string) {
	o.Decision.Set(&v)
}
// SetDecisionNil sets the value for Decision to be an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetDecisionNil() {
	o.Decision.Set(nil)
}

// UnsetDecision ensures that no value is present for Decision, not even an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetDecision() {
	o.Decision.Unset()
}

// GetJustification returns the Justification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetJustification() string {
	if o == nil || o.Justification.Get() == nil {
		var ret string
		return ret
	}
	return *o.Justification.Get()
}

// GetJustificationOk returns a tuple with the Justification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetJustificationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Justification.Get(), o.Justification.IsSet()
}

// HasJustification returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasJustification() bool {
	if o != nil && o.Justification.IsSet() {
		return true
	}

	return false
}

// SetJustification gets a reference to the given NullableString and assigns it to the Justification field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetJustification(v string) {
	o.Justification.Set(&v)
}
// SetJustificationNil sets the value for Justification to be an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetJustificationNil() {
	o.Justification.Set(nil)
}

// UnsetJustification ensures that no value is present for Justification, not even an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetJustification() {
	o.Justification.Unset()
}

// GetPrincipal returns the Principal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetPrincipal() AnyOfmicrosoftGraphIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentity
		return ret
	}
	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetPrincipalOk() (*AnyOfmicrosoftGraphIdentity, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return &o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given AnyOfmicrosoftGraphIdentity and assigns it to the Principal field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetPrincipal(v AnyOfmicrosoftGraphIdentity) {
	o.Principal = v
}

// GetPrincipalLink returns the PrincipalLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetPrincipalLink() string {
	if o == nil || o.PrincipalLink.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalLink.Get()
}

// GetPrincipalLinkOk returns a tuple with the PrincipalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetPrincipalLinkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalLink.Get(), o.PrincipalLink.IsSet()
}

// HasPrincipalLink returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasPrincipalLink() bool {
	if o != nil && o.PrincipalLink.IsSet() {
		return true
	}

	return false
}

// SetPrincipalLink gets a reference to the given NullableString and assigns it to the PrincipalLink field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetPrincipalLink(v string) {
	o.PrincipalLink.Set(&v)
}
// SetPrincipalLinkNil sets the value for PrincipalLink to be an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetPrincipalLinkNil() {
	o.PrincipalLink.Set(nil)
}

// UnsetPrincipalLink ensures that no value is present for PrincipalLink, not even an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetPrincipalLink() {
	o.PrincipalLink.Unset()
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetRecommendation() string {
	if o == nil || o.Recommendation.Get() == nil {
		var ret string
		return ret
	}
	return *o.Recommendation.Get()
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetRecommendationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Recommendation.Get(), o.Recommendation.IsSet()
}

// HasRecommendation returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasRecommendation() bool {
	if o != nil && o.Recommendation.IsSet() {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given NullableString and assigns it to the Recommendation field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetRecommendation(v string) {
	o.Recommendation.Set(&v)
}
// SetRecommendationNil sets the value for Recommendation to be an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetRecommendationNil() {
	o.Recommendation.Set(nil)
}

// UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetRecommendation() {
	o.Recommendation.Unset()
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetResource() AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetResourceOk() (*AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return &o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource and assigns it to the Resource field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetResource(v AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource) {
	o.Resource = v
}

// GetResourceLink returns the ResourceLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetResourceLink() string {
	if o == nil || o.ResourceLink.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceLink.Get()
}

// GetResourceLinkOk returns a tuple with the ResourceLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetResourceLinkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceLink.Get(), o.ResourceLink.IsSet()
}

// HasResourceLink returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasResourceLink() bool {
	if o != nil && o.ResourceLink.IsSet() {
		return true
	}

	return false
}

// SetResourceLink gets a reference to the given NullableString and assigns it to the ResourceLink field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetResourceLink(v string) {
	o.ResourceLink.Set(&v)
}
// SetResourceLinkNil sets the value for ResourceLink to be an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetResourceLinkNil() {
	o.ResourceLink.Set(nil)
}

// UnsetResourceLink ensures that no value is present for ResourceLink, not even an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetResourceLink() {
	o.ResourceLink.Unset()
}

// GetReviewedBy returns the ReviewedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetReviewedBy() AnyOfmicrosoftGraphUserIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUserIdentity
		return ret
	}
	return o.ReviewedBy
}

// GetReviewedByOk returns a tuple with the ReviewedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetReviewedByOk() (*AnyOfmicrosoftGraphUserIdentity, bool) {
	if o == nil || o.ReviewedBy == nil {
		return nil, false
	}
	return &o.ReviewedBy, true
}

// HasReviewedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasReviewedBy() bool {
	if o != nil && o.ReviewedBy != nil {
		return true
	}

	return false
}

// SetReviewedBy gets a reference to the given AnyOfmicrosoftGraphUserIdentity and assigns it to the ReviewedBy field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetReviewedBy(v AnyOfmicrosoftGraphUserIdentity) {
	o.ReviewedBy = v
}

// GetReviewedDateTime returns the ReviewedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetReviewedDateTime() time.Time {
	if o == nil || o.ReviewedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ReviewedDateTime.Get()
}

// GetReviewedDateTimeOk returns a tuple with the ReviewedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetReviewedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReviewedDateTime.Get(), o.ReviewedDateTime.IsSet()
}

// HasReviewedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasReviewedDateTime() bool {
	if o != nil && o.ReviewedDateTime.IsSet() {
		return true
	}

	return false
}

// SetReviewedDateTime gets a reference to the given NullableTime and assigns it to the ReviewedDateTime field.
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetReviewedDateTime(v time.Time) {
	o.ReviewedDateTime.Set(&v)
}
// SetReviewedDateTimeNil sets the value for ReviewedDateTime to be an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetReviewedDateTimeNil() {
	o.ReviewedDateTime.Set(nil)
}

// UnsetReviewedDateTime ensures that no value is present for ReviewedDateTime, not even an explicit nil
func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetReviewedDateTime() {
	o.ReviewedDateTime.Unset()
}

func (o MicrosoftGraphAccessReviewInstanceDecisionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AccessReviewId != nil {
		toSerialize["accessReviewId"] = o.AccessReviewId
	}
	if o.AppliedBy != nil {
		toSerialize["appliedBy"] = o.AppliedBy
	}
	if o.AppliedDateTime.IsSet() {
		toSerialize["appliedDateTime"] = o.AppliedDateTime.Get()
	}
	if o.ApplyResult.IsSet() {
		toSerialize["applyResult"] = o.ApplyResult.Get()
	}
	if o.Decision.IsSet() {
		toSerialize["decision"] = o.Decision.Get()
	}
	if o.Justification.IsSet() {
		toSerialize["justification"] = o.Justification.Get()
	}
	if o.Principal != nil {
		toSerialize["principal"] = o.Principal
	}
	if o.PrincipalLink.IsSet() {
		toSerialize["principalLink"] = o.PrincipalLink.Get()
	}
	if o.Recommendation.IsSet() {
		toSerialize["recommendation"] = o.Recommendation.Get()
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.ResourceLink.IsSet() {
		toSerialize["resourceLink"] = o.ResourceLink.Get()
	}
	if o.ReviewedBy != nil {
		toSerialize["reviewedBy"] = o.ReviewedBy
	}
	if o.ReviewedDateTime.IsSet() {
		toSerialize["reviewedDateTime"] = o.ReviewedDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAccessReviewInstanceDecisionItem struct {
	value *MicrosoftGraphAccessReviewInstanceDecisionItem
	isSet bool
}

func (v NullableMicrosoftGraphAccessReviewInstanceDecisionItem) Get() *MicrosoftGraphAccessReviewInstanceDecisionItem {
	return v.value
}

func (v *NullableMicrosoftGraphAccessReviewInstanceDecisionItem) Set(val *MicrosoftGraphAccessReviewInstanceDecisionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAccessReviewInstanceDecisionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAccessReviewInstanceDecisionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAccessReviewInstanceDecisionItem(val *MicrosoftGraphAccessReviewInstanceDecisionItem) *NullableMicrosoftGraphAccessReviewInstanceDecisionItem {
	return &NullableMicrosoftGraphAccessReviewInstanceDecisionItem{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAccessReviewInstanceDecisionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAccessReviewInstanceDecisionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


