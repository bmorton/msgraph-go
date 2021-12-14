# MicrosoftGraphPersistentBrowserSessionControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **NullableBool** | Specifies whether the session control is enabled. | [optional] 
**Mode** | Pointer to [**AnyOfmicrosoftGraphPersistentBrowserSessionMode**](anyOf&lt;microsoft.graph.persistentBrowserSessionMode&gt;.md) | Possible values are: always, never. | [optional] 

## Methods

### NewMicrosoftGraphPersistentBrowserSessionControl

`func NewMicrosoftGraphPersistentBrowserSessionControl() *MicrosoftGraphPersistentBrowserSessionControl`

NewMicrosoftGraphPersistentBrowserSessionControl instantiates a new MicrosoftGraphPersistentBrowserSessionControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPersistentBrowserSessionControlWithDefaults

`func NewMicrosoftGraphPersistentBrowserSessionControlWithDefaults() *MicrosoftGraphPersistentBrowserSessionControl`

NewMicrosoftGraphPersistentBrowserSessionControlWithDefaults instantiates a new MicrosoftGraphPersistentBrowserSessionControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *MicrosoftGraphPersistentBrowserSessionControl) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphPersistentBrowserSessionControl) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphPersistentBrowserSessionControl) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphPersistentBrowserSessionControl) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *MicrosoftGraphPersistentBrowserSessionControl) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *MicrosoftGraphPersistentBrowserSessionControl) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetMode

`func (o *MicrosoftGraphPersistentBrowserSessionControl) GetMode() AnyOfmicrosoftGraphPersistentBrowserSessionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MicrosoftGraphPersistentBrowserSessionControl) GetModeOk() (*AnyOfmicrosoftGraphPersistentBrowserSessionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MicrosoftGraphPersistentBrowserSessionControl) SetMode(v AnyOfmicrosoftGraphPersistentBrowserSessionMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MicrosoftGraphPersistentBrowserSessionControl) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *MicrosoftGraphPersistentBrowserSessionControl) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *MicrosoftGraphPersistentBrowserSessionControl) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


