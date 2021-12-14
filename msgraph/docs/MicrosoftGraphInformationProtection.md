# MicrosoftGraphInformationProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Bitlocker** | Pointer to [**AnyOfmicrosoftGraphBitlocker**](anyOf&lt;microsoft.graph.bitlocker&gt;.md) |  | [optional] 
**ThreatAssessmentRequests** | Pointer to [**[]MicrosoftGraphThreatAssessmentRequest**](MicrosoftGraphThreatAssessmentRequest.md) |  | [optional] 

## Methods

### NewMicrosoftGraphInformationProtection

`func NewMicrosoftGraphInformationProtection() *MicrosoftGraphInformationProtection`

NewMicrosoftGraphInformationProtection instantiates a new MicrosoftGraphInformationProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphInformationProtectionWithDefaults

`func NewMicrosoftGraphInformationProtectionWithDefaults() *MicrosoftGraphInformationProtection`

NewMicrosoftGraphInformationProtectionWithDefaults instantiates a new MicrosoftGraphInformationProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphInformationProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphInformationProtection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphInformationProtection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphInformationProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBitlocker

`func (o *MicrosoftGraphInformationProtection) GetBitlocker() AnyOfmicrosoftGraphBitlocker`

GetBitlocker returns the Bitlocker field if non-nil, zero value otherwise.

### GetBitlockerOk

`func (o *MicrosoftGraphInformationProtection) GetBitlockerOk() (*AnyOfmicrosoftGraphBitlocker, bool)`

GetBitlockerOk returns a tuple with the Bitlocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitlocker

`func (o *MicrosoftGraphInformationProtection) SetBitlocker(v AnyOfmicrosoftGraphBitlocker)`

SetBitlocker sets Bitlocker field to given value.

### HasBitlocker

`func (o *MicrosoftGraphInformationProtection) HasBitlocker() bool`

HasBitlocker returns a boolean if a field has been set.

### SetBitlockerNil

`func (o *MicrosoftGraphInformationProtection) SetBitlockerNil(b bool)`

 SetBitlockerNil sets the value for Bitlocker to be an explicit nil

### UnsetBitlocker
`func (o *MicrosoftGraphInformationProtection) UnsetBitlocker()`

UnsetBitlocker ensures that no value is present for Bitlocker, not even an explicit nil
### GetThreatAssessmentRequests

`func (o *MicrosoftGraphInformationProtection) GetThreatAssessmentRequests() []MicrosoftGraphThreatAssessmentRequest`

GetThreatAssessmentRequests returns the ThreatAssessmentRequests field if non-nil, zero value otherwise.

### GetThreatAssessmentRequestsOk

`func (o *MicrosoftGraphInformationProtection) GetThreatAssessmentRequestsOk() (*[]MicrosoftGraphThreatAssessmentRequest, bool)`

GetThreatAssessmentRequestsOk returns a tuple with the ThreatAssessmentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatAssessmentRequests

`func (o *MicrosoftGraphInformationProtection) SetThreatAssessmentRequests(v []MicrosoftGraphThreatAssessmentRequest)`

SetThreatAssessmentRequests sets ThreatAssessmentRequests field to given value.

### HasThreatAssessmentRequests

`func (o *MicrosoftGraphInformationProtection) HasThreatAssessmentRequests() bool`

HasThreatAssessmentRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


