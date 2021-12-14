# InformationProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bitlocker** | Pointer to [**AnyOfmicrosoftGraphBitlocker**](anyOf&lt;microsoft.graph.bitlocker&gt;.md) |  | [optional] 
**ThreatAssessmentRequests** | Pointer to [**[]MicrosoftGraphThreatAssessmentRequest**](MicrosoftGraphThreatAssessmentRequest.md) |  | [optional] 

## Methods

### NewInformationProtection

`func NewInformationProtection() *InformationProtection`

NewInformationProtection instantiates a new InformationProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInformationProtectionWithDefaults

`func NewInformationProtectionWithDefaults() *InformationProtection`

NewInformationProtectionWithDefaults instantiates a new InformationProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitlocker

`func (o *InformationProtection) GetBitlocker() AnyOfmicrosoftGraphBitlocker`

GetBitlocker returns the Bitlocker field if non-nil, zero value otherwise.

### GetBitlockerOk

`func (o *InformationProtection) GetBitlockerOk() (*AnyOfmicrosoftGraphBitlocker, bool)`

GetBitlockerOk returns a tuple with the Bitlocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitlocker

`func (o *InformationProtection) SetBitlocker(v AnyOfmicrosoftGraphBitlocker)`

SetBitlocker sets Bitlocker field to given value.

### HasBitlocker

`func (o *InformationProtection) HasBitlocker() bool`

HasBitlocker returns a boolean if a field has been set.

### SetBitlockerNil

`func (o *InformationProtection) SetBitlockerNil(b bool)`

 SetBitlockerNil sets the value for Bitlocker to be an explicit nil

### UnsetBitlocker
`func (o *InformationProtection) UnsetBitlocker()`

UnsetBitlocker ensures that no value is present for Bitlocker, not even an explicit nil
### GetThreatAssessmentRequests

`func (o *InformationProtection) GetThreatAssessmentRequests() []MicrosoftGraphThreatAssessmentRequest`

GetThreatAssessmentRequests returns the ThreatAssessmentRequests field if non-nil, zero value otherwise.

### GetThreatAssessmentRequestsOk

`func (o *InformationProtection) GetThreatAssessmentRequestsOk() (*[]MicrosoftGraphThreatAssessmentRequest, bool)`

GetThreatAssessmentRequestsOk returns a tuple with the ThreatAssessmentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatAssessmentRequests

`func (o *InformationProtection) SetThreatAssessmentRequests(v []MicrosoftGraphThreatAssessmentRequest)`

SetThreatAssessmentRequests sets ThreatAssessmentRequests field to given value.

### HasThreatAssessmentRequests

`func (o *InformationProtection) HasThreatAssessmentRequests() bool`

HasThreatAssessmentRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


