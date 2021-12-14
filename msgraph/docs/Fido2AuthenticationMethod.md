# Fido2AuthenticationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AaGuid** | Pointer to **NullableString** | Authenticator Attestation GUID, an identifier that indicates the type (e.g. make and model) of the authenticator. | [optional] 
**AttestationCertificates** | Pointer to **[]string** | The attestation certificate(s) attached to this security key. | [optional] 
**AttestationLevel** | Pointer to [**AnyOfmicrosoftGraphAttestationLevel**](anyOf&lt;microsoft.graph.attestationLevel&gt;.md) | The attestation level of this FIDO2 security key. Possible values are: attested, or notAttested. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The timestamp when this key was registered to the user. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the key as given by the user. | [optional] 
**Model** | Pointer to **NullableString** | The manufacturer-assigned model of the FIDO2 security key. | [optional] 

## Methods

### NewFido2AuthenticationMethod

`func NewFido2AuthenticationMethod() *Fido2AuthenticationMethod`

NewFido2AuthenticationMethod instantiates a new Fido2AuthenticationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFido2AuthenticationMethodWithDefaults

`func NewFido2AuthenticationMethodWithDefaults() *Fido2AuthenticationMethod`

NewFido2AuthenticationMethodWithDefaults instantiates a new Fido2AuthenticationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAaGuid

`func (o *Fido2AuthenticationMethod) GetAaGuid() string`

GetAaGuid returns the AaGuid field if non-nil, zero value otherwise.

### GetAaGuidOk

`func (o *Fido2AuthenticationMethod) GetAaGuidOk() (*string, bool)`

GetAaGuidOk returns a tuple with the AaGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaGuid

`func (o *Fido2AuthenticationMethod) SetAaGuid(v string)`

SetAaGuid sets AaGuid field to given value.

### HasAaGuid

`func (o *Fido2AuthenticationMethod) HasAaGuid() bool`

HasAaGuid returns a boolean if a field has been set.

### SetAaGuidNil

`func (o *Fido2AuthenticationMethod) SetAaGuidNil(b bool)`

 SetAaGuidNil sets the value for AaGuid to be an explicit nil

### UnsetAaGuid
`func (o *Fido2AuthenticationMethod) UnsetAaGuid()`

UnsetAaGuid ensures that no value is present for AaGuid, not even an explicit nil
### GetAttestationCertificates

`func (o *Fido2AuthenticationMethod) GetAttestationCertificates() []*string`

GetAttestationCertificates returns the AttestationCertificates field if non-nil, zero value otherwise.

### GetAttestationCertificatesOk

`func (o *Fido2AuthenticationMethod) GetAttestationCertificatesOk() (*[]*string, bool)`

GetAttestationCertificatesOk returns a tuple with the AttestationCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationCertificates

`func (o *Fido2AuthenticationMethod) SetAttestationCertificates(v []*string)`

SetAttestationCertificates sets AttestationCertificates field to given value.

### HasAttestationCertificates

`func (o *Fido2AuthenticationMethod) HasAttestationCertificates() bool`

HasAttestationCertificates returns a boolean if a field has been set.

### GetAttestationLevel

`func (o *Fido2AuthenticationMethod) GetAttestationLevel() AnyOfmicrosoftGraphAttestationLevel`

GetAttestationLevel returns the AttestationLevel field if non-nil, zero value otherwise.

### GetAttestationLevelOk

`func (o *Fido2AuthenticationMethod) GetAttestationLevelOk() (*AnyOfmicrosoftGraphAttestationLevel, bool)`

GetAttestationLevelOk returns a tuple with the AttestationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationLevel

`func (o *Fido2AuthenticationMethod) SetAttestationLevel(v AnyOfmicrosoftGraphAttestationLevel)`

SetAttestationLevel sets AttestationLevel field to given value.

### HasAttestationLevel

`func (o *Fido2AuthenticationMethod) HasAttestationLevel() bool`

HasAttestationLevel returns a boolean if a field has been set.

### SetAttestationLevelNil

`func (o *Fido2AuthenticationMethod) SetAttestationLevelNil(b bool)`

 SetAttestationLevelNil sets the value for AttestationLevel to be an explicit nil

### UnsetAttestationLevel
`func (o *Fido2AuthenticationMethod) UnsetAttestationLevel()`

UnsetAttestationLevel ensures that no value is present for AttestationLevel, not even an explicit nil
### GetCreatedDateTime

`func (o *Fido2AuthenticationMethod) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Fido2AuthenticationMethod) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Fido2AuthenticationMethod) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Fido2AuthenticationMethod) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *Fido2AuthenticationMethod) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *Fido2AuthenticationMethod) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDisplayName

`func (o *Fido2AuthenticationMethod) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Fido2AuthenticationMethod) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Fido2AuthenticationMethod) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Fido2AuthenticationMethod) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Fido2AuthenticationMethod) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Fido2AuthenticationMethod) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetModel

`func (o *Fido2AuthenticationMethod) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Fido2AuthenticationMethod) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Fido2AuthenticationMethod) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Fido2AuthenticationMethod) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *Fido2AuthenticationMethod) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *Fido2AuthenticationMethod) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


