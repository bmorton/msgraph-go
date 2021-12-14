# MicrosoftGraphSettingSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | Not yet documented | [optional] 
**Id** | Pointer to **NullableString** | Not yet documented | [optional] 
**SourceType** | Pointer to [**AnyOfmicrosoftGraphSettingSourceType**](anyOf&lt;microsoft.graph.settingSourceType&gt;.md) | Not yet documented. Possible values are: deviceConfiguration, deviceIntent. | [optional] 

## Methods

### NewMicrosoftGraphSettingSource

`func NewMicrosoftGraphSettingSource() *MicrosoftGraphSettingSource`

NewMicrosoftGraphSettingSource instantiates a new MicrosoftGraphSettingSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSettingSourceWithDefaults

`func NewMicrosoftGraphSettingSourceWithDefaults() *MicrosoftGraphSettingSource`

NewMicrosoftGraphSettingSourceWithDefaults instantiates a new MicrosoftGraphSettingSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphSettingSource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphSettingSource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphSettingSource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphSettingSource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphSettingSource) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphSettingSource) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetId

`func (o *MicrosoftGraphSettingSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSettingSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSettingSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSettingSource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphSettingSource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphSettingSource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSourceType

`func (o *MicrosoftGraphSettingSource) GetSourceType() AnyOfmicrosoftGraphSettingSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MicrosoftGraphSettingSource) GetSourceTypeOk() (*AnyOfmicrosoftGraphSettingSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MicrosoftGraphSettingSource) SetSourceType(v AnyOfmicrosoftGraphSettingSourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *MicrosoftGraphSettingSource) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *MicrosoftGraphSettingSource) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *MicrosoftGraphSettingSource) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


