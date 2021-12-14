# MicrosoftGraphParticipantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **NullableString** | The ISO 3166-1 Alpha-2 country code of the participant&#39;s best estimated physical location at the start of the call. Read-only. | [optional] 
**EndpointType** | Pointer to [**AnyOfmicrosoftGraphEndpointType**](anyOf&lt;microsoft.graph.endpointType&gt;.md) | The type of endpoint the participant is using. Possible values are: default, skypeForBusiness, or skypeForBusinessVoipPhone. Read-only. | [optional] 
**Identity** | Pointer to [**MicrosoftGraphIdentitySet**](MicrosoftGraphIdentitySet.md) |  | [optional] 
**LanguageId** | Pointer to **NullableString** | The language culture string. Read-only. | [optional] 
**ParticipantId** | Pointer to **NullableString** | The participant ID of the participant. Read-only. | [optional] 
**Region** | Pointer to **NullableString** | The home region of the participant. This can be a country, a continent, or a larger geographic region. This does not change based on the participant&#39;s current physical location. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphParticipantInfo

`func NewMicrosoftGraphParticipantInfo() *MicrosoftGraphParticipantInfo`

NewMicrosoftGraphParticipantInfo instantiates a new MicrosoftGraphParticipantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphParticipantInfoWithDefaults

`func NewMicrosoftGraphParticipantInfoWithDefaults() *MicrosoftGraphParticipantInfo`

NewMicrosoftGraphParticipantInfoWithDefaults instantiates a new MicrosoftGraphParticipantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *MicrosoftGraphParticipantInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *MicrosoftGraphParticipantInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *MicrosoftGraphParticipantInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *MicrosoftGraphParticipantInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *MicrosoftGraphParticipantInfo) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *MicrosoftGraphParticipantInfo) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetEndpointType

`func (o *MicrosoftGraphParticipantInfo) GetEndpointType() AnyOfmicrosoftGraphEndpointType`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *MicrosoftGraphParticipantInfo) GetEndpointTypeOk() (*AnyOfmicrosoftGraphEndpointType, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *MicrosoftGraphParticipantInfo) SetEndpointType(v AnyOfmicrosoftGraphEndpointType)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *MicrosoftGraphParticipantInfo) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### SetEndpointTypeNil

`func (o *MicrosoftGraphParticipantInfo) SetEndpointTypeNil(b bool)`

 SetEndpointTypeNil sets the value for EndpointType to be an explicit nil

### UnsetEndpointType
`func (o *MicrosoftGraphParticipantInfo) UnsetEndpointType()`

UnsetEndpointType ensures that no value is present for EndpointType, not even an explicit nil
### GetIdentity

`func (o *MicrosoftGraphParticipantInfo) GetIdentity() MicrosoftGraphIdentitySet`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *MicrosoftGraphParticipantInfo) GetIdentityOk() (*MicrosoftGraphIdentitySet, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *MicrosoftGraphParticipantInfo) SetIdentity(v MicrosoftGraphIdentitySet)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *MicrosoftGraphParticipantInfo) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLanguageId

`func (o *MicrosoftGraphParticipantInfo) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *MicrosoftGraphParticipantInfo) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *MicrosoftGraphParticipantInfo) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *MicrosoftGraphParticipantInfo) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *MicrosoftGraphParticipantInfo) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *MicrosoftGraphParticipantInfo) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetParticipantId

`func (o *MicrosoftGraphParticipantInfo) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *MicrosoftGraphParticipantInfo) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *MicrosoftGraphParticipantInfo) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *MicrosoftGraphParticipantInfo) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### SetParticipantIdNil

`func (o *MicrosoftGraphParticipantInfo) SetParticipantIdNil(b bool)`

 SetParticipantIdNil sets the value for ParticipantId to be an explicit nil

### UnsetParticipantId
`func (o *MicrosoftGraphParticipantInfo) UnsetParticipantId()`

UnsetParticipantId ensures that no value is present for ParticipantId, not even an explicit nil
### GetRegion

`func (o *MicrosoftGraphParticipantInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MicrosoftGraphParticipantInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MicrosoftGraphParticipantInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MicrosoftGraphParticipantInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *MicrosoftGraphParticipantInfo) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *MicrosoftGraphParticipantInfo) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


