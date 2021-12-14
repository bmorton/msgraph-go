# AccessReviewInstanceDecisionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewAccessReviewInstanceDecisionItem

`func NewAccessReviewInstanceDecisionItem() *AccessReviewInstanceDecisionItem`

NewAccessReviewInstanceDecisionItem instantiates a new AccessReviewInstanceDecisionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessReviewInstanceDecisionItemWithDefaults

`func NewAccessReviewInstanceDecisionItemWithDefaults() *AccessReviewInstanceDecisionItem`

NewAccessReviewInstanceDecisionItemWithDefaults instantiates a new AccessReviewInstanceDecisionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessReviewId

`func (o *AccessReviewInstanceDecisionItem) GetAccessReviewId() string`

GetAccessReviewId returns the AccessReviewId field if non-nil, zero value otherwise.

### GetAccessReviewIdOk

`func (o *AccessReviewInstanceDecisionItem) GetAccessReviewIdOk() (*string, bool)`

GetAccessReviewIdOk returns a tuple with the AccessReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessReviewId

`func (o *AccessReviewInstanceDecisionItem) SetAccessReviewId(v string)`

SetAccessReviewId sets AccessReviewId field to given value.

### HasAccessReviewId

`func (o *AccessReviewInstanceDecisionItem) HasAccessReviewId() bool`

HasAccessReviewId returns a boolean if a field has been set.

### GetAppliedBy

`func (o *AccessReviewInstanceDecisionItem) GetAppliedBy() AnyOfmicrosoftGraphUserIdentity`

GetAppliedBy returns the AppliedBy field if non-nil, zero value otherwise.

### GetAppliedByOk

`func (o *AccessReviewInstanceDecisionItem) GetAppliedByOk() (*AnyOfmicrosoftGraphUserIdentity, bool)`

GetAppliedByOk returns a tuple with the AppliedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBy

`func (o *AccessReviewInstanceDecisionItem) SetAppliedBy(v AnyOfmicrosoftGraphUserIdentity)`

SetAppliedBy sets AppliedBy field to given value.

### HasAppliedBy

`func (o *AccessReviewInstanceDecisionItem) HasAppliedBy() bool`

HasAppliedBy returns a boolean if a field has been set.

### SetAppliedByNil

`func (o *AccessReviewInstanceDecisionItem) SetAppliedByNil(b bool)`

 SetAppliedByNil sets the value for AppliedBy to be an explicit nil

### UnsetAppliedBy
`func (o *AccessReviewInstanceDecisionItem) UnsetAppliedBy()`

UnsetAppliedBy ensures that no value is present for AppliedBy, not even an explicit nil
### GetAppliedDateTime

`func (o *AccessReviewInstanceDecisionItem) GetAppliedDateTime() time.Time`

GetAppliedDateTime returns the AppliedDateTime field if non-nil, zero value otherwise.

### GetAppliedDateTimeOk

`func (o *AccessReviewInstanceDecisionItem) GetAppliedDateTimeOk() (*time.Time, bool)`

GetAppliedDateTimeOk returns a tuple with the AppliedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDateTime

`func (o *AccessReviewInstanceDecisionItem) SetAppliedDateTime(v time.Time)`

SetAppliedDateTime sets AppliedDateTime field to given value.

### HasAppliedDateTime

`func (o *AccessReviewInstanceDecisionItem) HasAppliedDateTime() bool`

HasAppliedDateTime returns a boolean if a field has been set.

### SetAppliedDateTimeNil

`func (o *AccessReviewInstanceDecisionItem) SetAppliedDateTimeNil(b bool)`

 SetAppliedDateTimeNil sets the value for AppliedDateTime to be an explicit nil

### UnsetAppliedDateTime
`func (o *AccessReviewInstanceDecisionItem) UnsetAppliedDateTime()`

UnsetAppliedDateTime ensures that no value is present for AppliedDateTime, not even an explicit nil
### GetApplyResult

`func (o *AccessReviewInstanceDecisionItem) GetApplyResult() string`

GetApplyResult returns the ApplyResult field if non-nil, zero value otherwise.

### GetApplyResultOk

`func (o *AccessReviewInstanceDecisionItem) GetApplyResultOk() (*string, bool)`

GetApplyResultOk returns a tuple with the ApplyResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyResult

`func (o *AccessReviewInstanceDecisionItem) SetApplyResult(v string)`

SetApplyResult sets ApplyResult field to given value.

### HasApplyResult

`func (o *AccessReviewInstanceDecisionItem) HasApplyResult() bool`

HasApplyResult returns a boolean if a field has been set.

### SetApplyResultNil

`func (o *AccessReviewInstanceDecisionItem) SetApplyResultNil(b bool)`

 SetApplyResultNil sets the value for ApplyResult to be an explicit nil

### UnsetApplyResult
`func (o *AccessReviewInstanceDecisionItem) UnsetApplyResult()`

UnsetApplyResult ensures that no value is present for ApplyResult, not even an explicit nil
### GetDecision

`func (o *AccessReviewInstanceDecisionItem) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *AccessReviewInstanceDecisionItem) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *AccessReviewInstanceDecisionItem) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *AccessReviewInstanceDecisionItem) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### SetDecisionNil

`func (o *AccessReviewInstanceDecisionItem) SetDecisionNil(b bool)`

 SetDecisionNil sets the value for Decision to be an explicit nil

### UnsetDecision
`func (o *AccessReviewInstanceDecisionItem) UnsetDecision()`

UnsetDecision ensures that no value is present for Decision, not even an explicit nil
### GetJustification

`func (o *AccessReviewInstanceDecisionItem) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *AccessReviewInstanceDecisionItem) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *AccessReviewInstanceDecisionItem) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *AccessReviewInstanceDecisionItem) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### SetJustificationNil

`func (o *AccessReviewInstanceDecisionItem) SetJustificationNil(b bool)`

 SetJustificationNil sets the value for Justification to be an explicit nil

### UnsetJustification
`func (o *AccessReviewInstanceDecisionItem) UnsetJustification()`

UnsetJustification ensures that no value is present for Justification, not even an explicit nil
### GetPrincipal

`func (o *AccessReviewInstanceDecisionItem) GetPrincipal() AnyOfmicrosoftGraphIdentity`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *AccessReviewInstanceDecisionItem) GetPrincipalOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *AccessReviewInstanceDecisionItem) SetPrincipal(v AnyOfmicrosoftGraphIdentity)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *AccessReviewInstanceDecisionItem) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *AccessReviewInstanceDecisionItem) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *AccessReviewInstanceDecisionItem) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetPrincipalLink

`func (o *AccessReviewInstanceDecisionItem) GetPrincipalLink() string`

GetPrincipalLink returns the PrincipalLink field if non-nil, zero value otherwise.

### GetPrincipalLinkOk

`func (o *AccessReviewInstanceDecisionItem) GetPrincipalLinkOk() (*string, bool)`

GetPrincipalLinkOk returns a tuple with the PrincipalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalLink

`func (o *AccessReviewInstanceDecisionItem) SetPrincipalLink(v string)`

SetPrincipalLink sets PrincipalLink field to given value.

### HasPrincipalLink

`func (o *AccessReviewInstanceDecisionItem) HasPrincipalLink() bool`

HasPrincipalLink returns a boolean if a field has been set.

### SetPrincipalLinkNil

`func (o *AccessReviewInstanceDecisionItem) SetPrincipalLinkNil(b bool)`

 SetPrincipalLinkNil sets the value for PrincipalLink to be an explicit nil

### UnsetPrincipalLink
`func (o *AccessReviewInstanceDecisionItem) UnsetPrincipalLink()`

UnsetPrincipalLink ensures that no value is present for PrincipalLink, not even an explicit nil
### GetRecommendation

`func (o *AccessReviewInstanceDecisionItem) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *AccessReviewInstanceDecisionItem) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *AccessReviewInstanceDecisionItem) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *AccessReviewInstanceDecisionItem) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### SetRecommendationNil

`func (o *AccessReviewInstanceDecisionItem) SetRecommendationNil(b bool)`

 SetRecommendationNil sets the value for Recommendation to be an explicit nil

### UnsetRecommendation
`func (o *AccessReviewInstanceDecisionItem) UnsetRecommendation()`

UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
### GetResource

`func (o *AccessReviewInstanceDecisionItem) GetResource() AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AccessReviewInstanceDecisionItem) GetResourceOk() (*AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AccessReviewInstanceDecisionItem) SetResource(v AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AccessReviewInstanceDecisionItem) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *AccessReviewInstanceDecisionItem) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *AccessReviewInstanceDecisionItem) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetResourceLink

`func (o *AccessReviewInstanceDecisionItem) GetResourceLink() string`

GetResourceLink returns the ResourceLink field if non-nil, zero value otherwise.

### GetResourceLinkOk

`func (o *AccessReviewInstanceDecisionItem) GetResourceLinkOk() (*string, bool)`

GetResourceLinkOk returns a tuple with the ResourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLink

`func (o *AccessReviewInstanceDecisionItem) SetResourceLink(v string)`

SetResourceLink sets ResourceLink field to given value.

### HasResourceLink

`func (o *AccessReviewInstanceDecisionItem) HasResourceLink() bool`

HasResourceLink returns a boolean if a field has been set.

### SetResourceLinkNil

`func (o *AccessReviewInstanceDecisionItem) SetResourceLinkNil(b bool)`

 SetResourceLinkNil sets the value for ResourceLink to be an explicit nil

### UnsetResourceLink
`func (o *AccessReviewInstanceDecisionItem) UnsetResourceLink()`

UnsetResourceLink ensures that no value is present for ResourceLink, not even an explicit nil
### GetReviewedBy

`func (o *AccessReviewInstanceDecisionItem) GetReviewedBy() AnyOfmicrosoftGraphUserIdentity`

GetReviewedBy returns the ReviewedBy field if non-nil, zero value otherwise.

### GetReviewedByOk

`func (o *AccessReviewInstanceDecisionItem) GetReviewedByOk() (*AnyOfmicrosoftGraphUserIdentity, bool)`

GetReviewedByOk returns a tuple with the ReviewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedBy

`func (o *AccessReviewInstanceDecisionItem) SetReviewedBy(v AnyOfmicrosoftGraphUserIdentity)`

SetReviewedBy sets ReviewedBy field to given value.

### HasReviewedBy

`func (o *AccessReviewInstanceDecisionItem) HasReviewedBy() bool`

HasReviewedBy returns a boolean if a field has been set.

### SetReviewedByNil

`func (o *AccessReviewInstanceDecisionItem) SetReviewedByNil(b bool)`

 SetReviewedByNil sets the value for ReviewedBy to be an explicit nil

### UnsetReviewedBy
`func (o *AccessReviewInstanceDecisionItem) UnsetReviewedBy()`

UnsetReviewedBy ensures that no value is present for ReviewedBy, not even an explicit nil
### GetReviewedDateTime

`func (o *AccessReviewInstanceDecisionItem) GetReviewedDateTime() time.Time`

GetReviewedDateTime returns the ReviewedDateTime field if non-nil, zero value otherwise.

### GetReviewedDateTimeOk

`func (o *AccessReviewInstanceDecisionItem) GetReviewedDateTimeOk() (*time.Time, bool)`

GetReviewedDateTimeOk returns a tuple with the ReviewedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDateTime

`func (o *AccessReviewInstanceDecisionItem) SetReviewedDateTime(v time.Time)`

SetReviewedDateTime sets ReviewedDateTime field to given value.

### HasReviewedDateTime

`func (o *AccessReviewInstanceDecisionItem) HasReviewedDateTime() bool`

HasReviewedDateTime returns a boolean if a field has been set.

### SetReviewedDateTimeNil

`func (o *AccessReviewInstanceDecisionItem) SetReviewedDateTimeNil(b bool)`

 SetReviewedDateTimeNil sets the value for ReviewedDateTime to be an explicit nil

### UnsetReviewedDateTime
`func (o *AccessReviewInstanceDecisionItem) UnsetReviewedDateTime()`

UnsetReviewedDateTime ensures that no value is present for ReviewedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


