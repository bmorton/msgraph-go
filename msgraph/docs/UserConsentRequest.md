# UserConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **NullableString** | The user&#39;s justification for requiring access to the app. Supports $filter (eq only) and $orderby. | [optional] 
**Approval** | Pointer to [**AnyOfmicrosoftGraphApproval**](anyOf&lt;microsoft.graph.approval&gt;.md) | Approval decisions associated with a request. | [optional] 

## Methods

### NewUserConsentRequest

`func NewUserConsentRequest() *UserConsentRequest`

NewUserConsentRequest instantiates a new UserConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConsentRequestWithDefaults

`func NewUserConsentRequestWithDefaults() *UserConsentRequest`

NewUserConsentRequestWithDefaults instantiates a new UserConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *UserConsentRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UserConsentRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UserConsentRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UserConsentRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *UserConsentRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *UserConsentRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetApproval

`func (o *UserConsentRequest) GetApproval() AnyOfmicrosoftGraphApproval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *UserConsentRequest) GetApprovalOk() (*AnyOfmicrosoftGraphApproval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *UserConsentRequest) SetApproval(v AnyOfmicrosoftGraphApproval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *UserConsentRequest) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### SetApprovalNil

`func (o *UserConsentRequest) SetApprovalNil(b bool)`

 SetApprovalNil sets the value for Approval to be an explicit nil

### UnsetApproval
`func (o *UserConsentRequest) UnsetApproval()`

UnsetApproval ensures that no value is present for Approval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


