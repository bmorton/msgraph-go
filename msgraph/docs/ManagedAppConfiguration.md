# ManagedAppConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomSettings** | Pointer to [**[]MicrosoftGraphKeyValuePair**](MicrosoftGraphKeyValuePair.md) | A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service | [optional] 

## Methods

### NewManagedAppConfiguration

`func NewManagedAppConfiguration() *ManagedAppConfiguration`

NewManagedAppConfiguration instantiates a new ManagedAppConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAppConfigurationWithDefaults

`func NewManagedAppConfigurationWithDefaults() *ManagedAppConfiguration`

NewManagedAppConfigurationWithDefaults instantiates a new ManagedAppConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomSettings

`func (o *ManagedAppConfiguration) GetCustomSettings() []MicrosoftGraphKeyValuePair`

GetCustomSettings returns the CustomSettings field if non-nil, zero value otherwise.

### GetCustomSettingsOk

`func (o *ManagedAppConfiguration) GetCustomSettingsOk() (*[]MicrosoftGraphKeyValuePair, bool)`

GetCustomSettingsOk returns a tuple with the CustomSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSettings

`func (o *ManagedAppConfiguration) SetCustomSettings(v []MicrosoftGraphKeyValuePair)`

SetCustomSettings sets CustomSettings field to given value.

### HasCustomSettings

`func (o *ManagedAppConfiguration) HasCustomSettings() bool`

HasCustomSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


