# MicrosoftGraphComplianceInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificationControls** | Pointer to [**[]AnyOfmicrosoftGraphCertificationControl**](AnyOfmicrosoftGraphCertificationControl.md) | Collection of the certification controls associated with certification | [optional] 
**CertificationName** | Pointer to **NullableString** | Compliance certification name (for example, ISO 27018:2014, GDPR, FedRAMP, NIST 800-171) | [optional] 

## Methods

### NewMicrosoftGraphComplianceInformation

`func NewMicrosoftGraphComplianceInformation() *MicrosoftGraphComplianceInformation`

NewMicrosoftGraphComplianceInformation instantiates a new MicrosoftGraphComplianceInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphComplianceInformationWithDefaults

`func NewMicrosoftGraphComplianceInformationWithDefaults() *MicrosoftGraphComplianceInformation`

NewMicrosoftGraphComplianceInformationWithDefaults instantiates a new MicrosoftGraphComplianceInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificationControls

`func (o *MicrosoftGraphComplianceInformation) GetCertificationControls() []*AnyOfmicrosoftGraphCertificationControl`

GetCertificationControls returns the CertificationControls field if non-nil, zero value otherwise.

### GetCertificationControlsOk

`func (o *MicrosoftGraphComplianceInformation) GetCertificationControlsOk() (*[]*AnyOfmicrosoftGraphCertificationControl, bool)`

GetCertificationControlsOk returns a tuple with the CertificationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationControls

`func (o *MicrosoftGraphComplianceInformation) SetCertificationControls(v []*AnyOfmicrosoftGraphCertificationControl)`

SetCertificationControls sets CertificationControls field to given value.

### HasCertificationControls

`func (o *MicrosoftGraphComplianceInformation) HasCertificationControls() bool`

HasCertificationControls returns a boolean if a field has been set.

### GetCertificationName

`func (o *MicrosoftGraphComplianceInformation) GetCertificationName() string`

GetCertificationName returns the CertificationName field if non-nil, zero value otherwise.

### GetCertificationNameOk

`func (o *MicrosoftGraphComplianceInformation) GetCertificationNameOk() (*string, bool)`

GetCertificationNameOk returns a tuple with the CertificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationName

`func (o *MicrosoftGraphComplianceInformation) SetCertificationName(v string)`

SetCertificationName sets CertificationName field to given value.

### HasCertificationName

`func (o *MicrosoftGraphComplianceInformation) HasCertificationName() bool`

HasCertificationName returns a boolean if a field has been set.

### SetCertificationNameNil

`func (o *MicrosoftGraphComplianceInformation) SetCertificationNameNil(b bool)`

 SetCertificationNameNil sets the value for CertificationName to be an explicit nil

### UnsetCertificationName
`func (o *MicrosoftGraphComplianceInformation) UnsetCertificationName()`

UnsetCertificationName ensures that no value is present for CertificationName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


