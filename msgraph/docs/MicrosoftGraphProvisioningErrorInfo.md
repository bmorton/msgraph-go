# MicrosoftGraphProvisioningErrorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | Pointer to **NullableString** | Additional details in case of error. | [optional] 
**ErrorCategory** | Pointer to [**AnyOfmicrosoftGraphProvisioningStatusErrorCategory**](anyOf&lt;microsoft.graph.provisioningStatusErrorCategory&gt;.md) | Categorizes the error code. Possible values are failure, nonServiceFailure, success, unknownFutureValue | [optional] 
**ErrorCode** | Pointer to **NullableString** | Unique error code if any occurred. Learn more | [optional] 
**Reason** | Pointer to **NullableString** | Summarizes the status and describes why the status happened. | [optional] 
**RecommendedAction** | Pointer to **NullableString** | Provides the resolution for the corresponding error. | [optional] 

## Methods

### NewMicrosoftGraphProvisioningErrorInfo

`func NewMicrosoftGraphProvisioningErrorInfo() *MicrosoftGraphProvisioningErrorInfo`

NewMicrosoftGraphProvisioningErrorInfo instantiates a new MicrosoftGraphProvisioningErrorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphProvisioningErrorInfoWithDefaults

`func NewMicrosoftGraphProvisioningErrorInfoWithDefaults() *MicrosoftGraphProvisioningErrorInfo`

NewMicrosoftGraphProvisioningErrorInfoWithDefaults instantiates a new MicrosoftGraphProvisioningErrorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *MicrosoftGraphProvisioningErrorInfo) GetAdditionalDetails() string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *MicrosoftGraphProvisioningErrorInfo) GetAdditionalDetailsOk() (*string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *MicrosoftGraphProvisioningErrorInfo) SetAdditionalDetails(v string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *MicrosoftGraphProvisioningErrorInfo) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### SetAdditionalDetailsNil

`func (o *MicrosoftGraphProvisioningErrorInfo) SetAdditionalDetailsNil(b bool)`

 SetAdditionalDetailsNil sets the value for AdditionalDetails to be an explicit nil

### UnsetAdditionalDetails
`func (o *MicrosoftGraphProvisioningErrorInfo) UnsetAdditionalDetails()`

UnsetAdditionalDetails ensures that no value is present for AdditionalDetails, not even an explicit nil
### GetErrorCategory

`func (o *MicrosoftGraphProvisioningErrorInfo) GetErrorCategory() AnyOfmicrosoftGraphProvisioningStatusErrorCategory`

GetErrorCategory returns the ErrorCategory field if non-nil, zero value otherwise.

### GetErrorCategoryOk

`func (o *MicrosoftGraphProvisioningErrorInfo) GetErrorCategoryOk() (*AnyOfmicrosoftGraphProvisioningStatusErrorCategory, bool)`

GetErrorCategoryOk returns a tuple with the ErrorCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCategory

`func (o *MicrosoftGraphProvisioningErrorInfo) SetErrorCategory(v AnyOfmicrosoftGraphProvisioningStatusErrorCategory)`

SetErrorCategory sets ErrorCategory field to given value.

### HasErrorCategory

`func (o *MicrosoftGraphProvisioningErrorInfo) HasErrorCategory() bool`

HasErrorCategory returns a boolean if a field has been set.

### SetErrorCategoryNil

`func (o *MicrosoftGraphProvisioningErrorInfo) SetErrorCategoryNil(b bool)`

 SetErrorCategoryNil sets the value for ErrorCategory to be an explicit nil

### UnsetErrorCategory
`func (o *MicrosoftGraphProvisioningErrorInfo) UnsetErrorCategory()`

UnsetErrorCategory ensures that no value is present for ErrorCategory, not even an explicit nil
### GetErrorCode

`func (o *MicrosoftGraphProvisioningErrorInfo) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MicrosoftGraphProvisioningErrorInfo) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MicrosoftGraphProvisioningErrorInfo) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *MicrosoftGraphProvisioningErrorInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *MicrosoftGraphProvisioningErrorInfo) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *MicrosoftGraphProvisioningErrorInfo) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetReason

`func (o *MicrosoftGraphProvisioningErrorInfo) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MicrosoftGraphProvisioningErrorInfo) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MicrosoftGraphProvisioningErrorInfo) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MicrosoftGraphProvisioningErrorInfo) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *MicrosoftGraphProvisioningErrorInfo) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *MicrosoftGraphProvisioningErrorInfo) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRecommendedAction

`func (o *MicrosoftGraphProvisioningErrorInfo) GetRecommendedAction() string`

GetRecommendedAction returns the RecommendedAction field if non-nil, zero value otherwise.

### GetRecommendedActionOk

`func (o *MicrosoftGraphProvisioningErrorInfo) GetRecommendedActionOk() (*string, bool)`

GetRecommendedActionOk returns a tuple with the RecommendedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedAction

`func (o *MicrosoftGraphProvisioningErrorInfo) SetRecommendedAction(v string)`

SetRecommendedAction sets RecommendedAction field to given value.

### HasRecommendedAction

`func (o *MicrosoftGraphProvisioningErrorInfo) HasRecommendedAction() bool`

HasRecommendedAction returns a boolean if a field has been set.

### SetRecommendedActionNil

`func (o *MicrosoftGraphProvisioningErrorInfo) SetRecommendedActionNil(b bool)`

 SetRecommendedActionNil sets the value for RecommendedAction to be an explicit nil

### UnsetRecommendedAction
`func (o *MicrosoftGraphProvisioningErrorInfo) UnsetRecommendedAction()`

UnsetRecommendedAction ensures that no value is present for RecommendedAction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


