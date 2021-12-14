# MicrosoftGraphUserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ContributionToContentDiscoveryAsOrganizationDisabled** | Pointer to **bool** | Reflects the Office Delve organization level setting. When set to true, the organization doesn&#39;t have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center. | [optional] 
**ContributionToContentDiscoveryDisabled** | Pointer to **bool** | When set to true, documents in the user&#39;s Office Delve are disabled. Users can control this setting in Office Delve. | [optional] 
**ShiftPreferences** | Pointer to [**AnyOfmicrosoftGraphShiftPreferences**](anyOf&lt;microsoft.graph.shiftPreferences&gt;.md) | The shift preferences for the user. | [optional] 

## Methods

### NewMicrosoftGraphUserSettings

`func NewMicrosoftGraphUserSettings() *MicrosoftGraphUserSettings`

NewMicrosoftGraphUserSettings instantiates a new MicrosoftGraphUserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUserSettingsWithDefaults

`func NewMicrosoftGraphUserSettingsWithDefaults() *MicrosoftGraphUserSettings`

NewMicrosoftGraphUserSettingsWithDefaults instantiates a new MicrosoftGraphUserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUserSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUserSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUserSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUserSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContributionToContentDiscoveryAsOrganizationDisabled

`func (o *MicrosoftGraphUserSettings) GetContributionToContentDiscoveryAsOrganizationDisabled() bool`

GetContributionToContentDiscoveryAsOrganizationDisabled returns the ContributionToContentDiscoveryAsOrganizationDisabled field if non-nil, zero value otherwise.

### GetContributionToContentDiscoveryAsOrganizationDisabledOk

`func (o *MicrosoftGraphUserSettings) GetContributionToContentDiscoveryAsOrganizationDisabledOk() (*bool, bool)`

GetContributionToContentDiscoveryAsOrganizationDisabledOk returns a tuple with the ContributionToContentDiscoveryAsOrganizationDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributionToContentDiscoveryAsOrganizationDisabled

`func (o *MicrosoftGraphUserSettings) SetContributionToContentDiscoveryAsOrganizationDisabled(v bool)`

SetContributionToContentDiscoveryAsOrganizationDisabled sets ContributionToContentDiscoveryAsOrganizationDisabled field to given value.

### HasContributionToContentDiscoveryAsOrganizationDisabled

`func (o *MicrosoftGraphUserSettings) HasContributionToContentDiscoveryAsOrganizationDisabled() bool`

HasContributionToContentDiscoveryAsOrganizationDisabled returns a boolean if a field has been set.

### GetContributionToContentDiscoveryDisabled

`func (o *MicrosoftGraphUserSettings) GetContributionToContentDiscoveryDisabled() bool`

GetContributionToContentDiscoveryDisabled returns the ContributionToContentDiscoveryDisabled field if non-nil, zero value otherwise.

### GetContributionToContentDiscoveryDisabledOk

`func (o *MicrosoftGraphUserSettings) GetContributionToContentDiscoveryDisabledOk() (*bool, bool)`

GetContributionToContentDiscoveryDisabledOk returns a tuple with the ContributionToContentDiscoveryDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributionToContentDiscoveryDisabled

`func (o *MicrosoftGraphUserSettings) SetContributionToContentDiscoveryDisabled(v bool)`

SetContributionToContentDiscoveryDisabled sets ContributionToContentDiscoveryDisabled field to given value.

### HasContributionToContentDiscoveryDisabled

`func (o *MicrosoftGraphUserSettings) HasContributionToContentDiscoveryDisabled() bool`

HasContributionToContentDiscoveryDisabled returns a boolean if a field has been set.

### GetShiftPreferences

`func (o *MicrosoftGraphUserSettings) GetShiftPreferences() AnyOfmicrosoftGraphShiftPreferences`

GetShiftPreferences returns the ShiftPreferences field if non-nil, zero value otherwise.

### GetShiftPreferencesOk

`func (o *MicrosoftGraphUserSettings) GetShiftPreferencesOk() (*AnyOfmicrosoftGraphShiftPreferences, bool)`

GetShiftPreferencesOk returns a tuple with the ShiftPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftPreferences

`func (o *MicrosoftGraphUserSettings) SetShiftPreferences(v AnyOfmicrosoftGraphShiftPreferences)`

SetShiftPreferences sets ShiftPreferences field to given value.

### HasShiftPreferences

`func (o *MicrosoftGraphUserSettings) HasShiftPreferences() bool`

HasShiftPreferences returns a boolean if a field has been set.

### SetShiftPreferencesNil

`func (o *MicrosoftGraphUserSettings) SetShiftPreferencesNil(b bool)`

 SetShiftPreferencesNil sets the value for ShiftPreferences to be an explicit nil

### UnsetShiftPreferences
`func (o *MicrosoftGraphUserSettings) UnsetShiftPreferences()`

UnsetShiftPreferences ensures that no value is present for ShiftPreferences, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


