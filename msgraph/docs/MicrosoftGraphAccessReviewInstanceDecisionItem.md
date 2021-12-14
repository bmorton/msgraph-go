# MicrosoftGraphAccessReviewInstanceDecisionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AccessReviewId** | Pointer to **string** | The identifier of the accessReviewInstance parent. Supports $select. Read-only. | [optional] 
**AppliedBy** | Pointer to [**AnyOfmicrosoftGraphUserIdentity**](anyOf&lt;microsoft.graph.userIdentity&gt;.md) | The identifier of the user who applied the decision. Read-only. | [optional] 
**AppliedDateTime** | Pointer to **NullableTime** | The timestamp when the approval decision was applied. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Supports $select. Read-only. | [optional] 
**ApplyResult** | Pointer to **NullableString** | The result of applying the decision. Possible values: New, AppliedSuccessfully, AppliedWithUnknownFailure, AppliedSuccessfullyButObjectNotFound and ApplyNotSupported. Supports $select, $orderby, and $filter (eq only). Read-only. | [optional] 
**Decision** | Pointer to **NullableString** | Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow. Supports $select, $orderby, and $filter (eq only). | [optional] 
**Justification** | Pointer to **NullableString** | Justification left by the reviewer when they made the decision. | [optional] 
**Principal** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | Every decision item in an access review represents a principal&#39;s access to a resource. This property represents details of the principal. For example, if a decision item represents access of User &#39;Bob&#39; to Group &#39;Sales&#39; - The principal is &#39;Bob&#39; and the resource is &#39;Sales&#39;. Principals can be of two types - userIdentity and servicePrincipalIdentity. Supports $select. Read-only. | [optional] 
**PrincipalLink** | Pointer to **NullableString** | A link to the principal object. For example, https://graph.microsoft.com/v1.0/users/a6c7aecb-cbfd-4763-87ef-e91b4bd509d9. Read-only. | [optional] 
**Recommendation** | Pointer to **NullableString** | A system-generated recommendation for the approval decision based off last interactive sign-in to tenant. Recommend approve if sign-in is within thirty days of start of review. Recommend deny if sign-in is greater than thirty days of start of review. Recommendation not available otherwise. Possible values: Approve, Deny, or NoInfoAvailable. Supports $select, $orderby, and $filter (eq only). Read-only. | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource**](anyOf&lt;microsoft.graph.accessReviewInstanceDecisionItemResource&gt;.md) | Every decision item in an access review represents a principal&#39;s access to a resource. This property represents details of the resource. For example, if a decision item represents access of User &#39;Bob&#39; to Group &#39;Sales&#39; - The principal is Bob and the resource is &#39;Sales&#39;. Resources can be of multiple types. See accessReviewInstanceDecisionItemResource. Read-only. | [optional] 
**ResourceLink** | Pointer to **NullableString** | A link to the resource. For example, https://graph.microsoft.com/v1.0/servicePrincipals/c86300f3-8695-4320-9f6e-32a2555f5ff8. Supports $select. Read-only. | [optional] 
**ReviewedBy** | Pointer to [**AnyOfmicrosoftGraphUserIdentity**](anyOf&lt;microsoft.graph.userIdentity&gt;.md) | The identifier of the reviewer. Supports $select. Read-only. | [optional] 
**ReviewedDateTime** | Pointer to **NullableTime** | The timestamp when the review decision occurred. Supports $select. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphAccessReviewInstanceDecisionItem

`func NewMicrosoftGraphAccessReviewInstanceDecisionItem() *MicrosoftGraphAccessReviewInstanceDecisionItem`

NewMicrosoftGraphAccessReviewInstanceDecisionItem instantiates a new MicrosoftGraphAccessReviewInstanceDecisionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAccessReviewInstanceDecisionItemWithDefaults

`func NewMicrosoftGraphAccessReviewInstanceDecisionItemWithDefaults() *MicrosoftGraphAccessReviewInstanceDecisionItem`

NewMicrosoftGraphAccessReviewInstanceDecisionItemWithDefaults instantiates a new MicrosoftGraphAccessReviewInstanceDecisionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccessReviewId

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAccessReviewId() string`

GetAccessReviewId returns the AccessReviewId field if non-nil, zero value otherwise.

### GetAccessReviewIdOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAccessReviewIdOk() (*string, bool)`

GetAccessReviewIdOk returns a tuple with the AccessReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessReviewId

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetAccessReviewId(v string)`

SetAccessReviewId sets AccessReviewId field to given value.

### HasAccessReviewId

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasAccessReviewId() bool`

HasAccessReviewId returns a boolean if a field has been set.

### GetAppliedBy

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAppliedBy() AnyOfmicrosoftGraphUserIdentity`

GetAppliedBy returns the AppliedBy field if non-nil, zero value otherwise.

### GetAppliedByOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAppliedByOk() (*AnyOfmicrosoftGraphUserIdentity, bool)`

GetAppliedByOk returns a tuple with the AppliedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBy

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetAppliedBy(v AnyOfmicrosoftGraphUserIdentity)`

SetAppliedBy sets AppliedBy field to given value.

### HasAppliedBy

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasAppliedBy() bool`

HasAppliedBy returns a boolean if a field has been set.

### SetAppliedByNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetAppliedByNil(b bool)`

 SetAppliedByNil sets the value for AppliedBy to be an explicit nil

### UnsetAppliedBy
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetAppliedBy()`

UnsetAppliedBy ensures that no value is present for AppliedBy, not even an explicit nil
### GetAppliedDateTime

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAppliedDateTime() time.Time`

GetAppliedDateTime returns the AppliedDateTime field if non-nil, zero value otherwise.

### GetAppliedDateTimeOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetAppliedDateTimeOk() (*time.Time, bool)`

GetAppliedDateTimeOk returns a tuple with the AppliedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDateTime

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetAppliedDateTime(v time.Time)`

SetAppliedDateTime sets AppliedDateTime field to given value.

### HasAppliedDateTime

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasAppliedDateTime() bool`

HasAppliedDateTime returns a boolean if a field has been set.

### SetAppliedDateTimeNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetAppliedDateTimeNil(b bool)`

 SetAppliedDateTimeNil sets the value for AppliedDateTime to be an explicit nil

### UnsetAppliedDateTime
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetAppliedDateTime()`

UnsetAppliedDateTime ensures that no value is present for AppliedDateTime, not even an explicit nil
### GetApplyResult

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetApplyResult() string`

GetApplyResult returns the ApplyResult field if non-nil, zero value otherwise.

### GetApplyResultOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetApplyResultOk() (*string, bool)`

GetApplyResultOk returns a tuple with the ApplyResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyResult

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetApplyResult(v string)`

SetApplyResult sets ApplyResult field to given value.

### HasApplyResult

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasApplyResult() bool`

HasApplyResult returns a boolean if a field has been set.

### SetApplyResultNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetApplyResultNil(b bool)`

 SetApplyResultNil sets the value for ApplyResult to be an explicit nil

### UnsetApplyResult
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetApplyResult()`

UnsetApplyResult ensures that no value is present for ApplyResult, not even an explicit nil
### GetDecision

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### SetDecisionNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetDecisionNil(b bool)`

 SetDecisionNil sets the value for Decision to be an explicit nil

### UnsetDecision
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetDecision()`

UnsetDecision ensures that no value is present for Decision, not even an explicit nil
### GetJustification

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### SetJustificationNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetJustificationNil(b bool)`

 SetJustificationNil sets the value for Justification to be an explicit nil

### UnsetJustification
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetJustification()`

UnsetJustification ensures that no value is present for Justification, not even an explicit nil
### GetPrincipal

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetPrincipal() AnyOfmicrosoftGraphIdentity`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetPrincipalOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetPrincipal(v AnyOfmicrosoftGraphIdentity)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetPrincipalLink

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetPrincipalLink() string`

GetPrincipalLink returns the PrincipalLink field if non-nil, zero value otherwise.

### GetPrincipalLinkOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetPrincipalLinkOk() (*string, bool)`

GetPrincipalLinkOk returns a tuple with the PrincipalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalLink

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetPrincipalLink(v string)`

SetPrincipalLink sets PrincipalLink field to given value.

### HasPrincipalLink

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasPrincipalLink() bool`

HasPrincipalLink returns a boolean if a field has been set.

### SetPrincipalLinkNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetPrincipalLinkNil(b bool)`

 SetPrincipalLinkNil sets the value for PrincipalLink to be an explicit nil

### UnsetPrincipalLink
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetPrincipalLink()`

UnsetPrincipalLink ensures that no value is present for PrincipalLink, not even an explicit nil
### GetRecommendation

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### SetRecommendationNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetRecommendationNil(b bool)`

 SetRecommendationNil sets the value for Recommendation to be an explicit nil

### UnsetRecommendation
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetRecommendation()`

UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
### GetResource

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetResource() AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetResourceOk() (*AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetResource(v AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetResourceLink

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetResourceLink() string`

GetResourceLink returns the ResourceLink field if non-nil, zero value otherwise.

### GetResourceLinkOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetResourceLinkOk() (*string, bool)`

GetResourceLinkOk returns a tuple with the ResourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLink

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetResourceLink(v string)`

SetResourceLink sets ResourceLink field to given value.

### HasResourceLink

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasResourceLink() bool`

HasResourceLink returns a boolean if a field has been set.

### SetResourceLinkNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetResourceLinkNil(b bool)`

 SetResourceLinkNil sets the value for ResourceLink to be an explicit nil

### UnsetResourceLink
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetResourceLink()`

UnsetResourceLink ensures that no value is present for ResourceLink, not even an explicit nil
### GetReviewedBy

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetReviewedBy() AnyOfmicrosoftGraphUserIdentity`

GetReviewedBy returns the ReviewedBy field if non-nil, zero value otherwise.

### GetReviewedByOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetReviewedByOk() (*AnyOfmicrosoftGraphUserIdentity, bool)`

GetReviewedByOk returns a tuple with the ReviewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedBy

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetReviewedBy(v AnyOfmicrosoftGraphUserIdentity)`

SetReviewedBy sets ReviewedBy field to given value.

### HasReviewedBy

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasReviewedBy() bool`

HasReviewedBy returns a boolean if a field has been set.

### SetReviewedByNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetReviewedByNil(b bool)`

 SetReviewedByNil sets the value for ReviewedBy to be an explicit nil

### UnsetReviewedBy
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetReviewedBy()`

UnsetReviewedBy ensures that no value is present for ReviewedBy, not even an explicit nil
### GetReviewedDateTime

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetReviewedDateTime() time.Time`

GetReviewedDateTime returns the ReviewedDateTime field if non-nil, zero value otherwise.

### GetReviewedDateTimeOk

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) GetReviewedDateTimeOk() (*time.Time, bool)`

GetReviewedDateTimeOk returns a tuple with the ReviewedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDateTime

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetReviewedDateTime(v time.Time)`

SetReviewedDateTime sets ReviewedDateTime field to given value.

### HasReviewedDateTime

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) HasReviewedDateTime() bool`

HasReviewedDateTime returns a boolean if a field has been set.

### SetReviewedDateTimeNil

`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) SetReviewedDateTimeNil(b bool)`

 SetReviewedDateTimeNil sets the value for ReviewedDateTime to be an explicit nil

### UnsetReviewedDateTime
`func (o *MicrosoftGraphAccessReviewInstanceDecisionItem) UnsetReviewedDateTime()`

UnsetReviewedDateTime ensures that no value is present for ReviewedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


