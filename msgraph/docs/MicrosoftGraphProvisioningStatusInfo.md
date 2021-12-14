# MicrosoftGraphProvisioningStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorInformation** | Pointer to [**AnyOfmicrosoftGraphProvisioningErrorInfo**](anyOf&lt;microsoft.graph.provisioningErrorInfo&gt;.md) |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphProvisioningResult**](anyOf&lt;microsoft.graph.provisioningResult&gt;.md) | Possible values are: success, warning, failure, skipped, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphProvisioningStatusInfo

`func NewMicrosoftGraphProvisioningStatusInfo() *MicrosoftGraphProvisioningStatusInfo`

NewMicrosoftGraphProvisioningStatusInfo instantiates a new MicrosoftGraphProvisioningStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphProvisioningStatusInfoWithDefaults

`func NewMicrosoftGraphProvisioningStatusInfoWithDefaults() *MicrosoftGraphProvisioningStatusInfo`

NewMicrosoftGraphProvisioningStatusInfoWithDefaults instantiates a new MicrosoftGraphProvisioningStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorInformation

`func (o *MicrosoftGraphProvisioningStatusInfo) GetErrorInformation() AnyOfmicrosoftGraphProvisioningErrorInfo`

GetErrorInformation returns the ErrorInformation field if non-nil, zero value otherwise.

### GetErrorInformationOk

`func (o *MicrosoftGraphProvisioningStatusInfo) GetErrorInformationOk() (*AnyOfmicrosoftGraphProvisioningErrorInfo, bool)`

GetErrorInformationOk returns a tuple with the ErrorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorInformation

`func (o *MicrosoftGraphProvisioningStatusInfo) SetErrorInformation(v AnyOfmicrosoftGraphProvisioningErrorInfo)`

SetErrorInformation sets ErrorInformation field to given value.

### HasErrorInformation

`func (o *MicrosoftGraphProvisioningStatusInfo) HasErrorInformation() bool`

HasErrorInformation returns a boolean if a field has been set.

### SetErrorInformationNil

`func (o *MicrosoftGraphProvisioningStatusInfo) SetErrorInformationNil(b bool)`

 SetErrorInformationNil sets the value for ErrorInformation to be an explicit nil

### UnsetErrorInformation
`func (o *MicrosoftGraphProvisioningStatusInfo) UnsetErrorInformation()`

UnsetErrorInformation ensures that no value is present for ErrorInformation, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphProvisioningStatusInfo) GetStatus() AnyOfmicrosoftGraphProvisioningResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphProvisioningStatusInfo) GetStatusOk() (*AnyOfmicrosoftGraphProvisioningResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphProvisioningStatusInfo) SetStatus(v AnyOfmicrosoftGraphProvisioningResult)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphProvisioningStatusInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphProvisioningStatusInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphProvisioningStatusInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


