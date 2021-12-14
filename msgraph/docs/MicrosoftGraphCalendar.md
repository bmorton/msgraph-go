# MicrosoftGraphCalendar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AllowedOnlineMeetingProviders** | Pointer to [**[]AnyOfmicrosoftGraphOnlineMeetingProviderType**](AnyOfmicrosoftGraphOnlineMeetingProviderType.md) | Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness. | [optional] 
**CanEdit** | Pointer to **NullableBool** | true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access. | [optional] 
**CanShare** | Pointer to **NullableBool** | true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it. | [optional] 
**CanViewPrivateItems** | Pointer to **NullableBool** | true if the user can read calendar items that have been marked private, false otherwise. | [optional] 
**ChangeKey** | Pointer to **NullableString** | Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only. | [optional] 
**Color** | Pointer to [**AnyOfmicrosoftGraphCalendarColor**](anyOf&lt;microsoft.graph.calendarColor&gt;.md) | Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor. | [optional] 
**DefaultOnlineMeetingProvider** | Pointer to [**AnyOfmicrosoftGraphOnlineMeetingProviderType**](anyOf&lt;microsoft.graph.onlineMeetingProviderType&gt;.md) | The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness. | [optional] 
**HexColor** | Pointer to **NullableString** | The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is empty. Read-only. | [optional] 
**IsDefaultCalendar** | Pointer to **NullableBool** | true if this is the default calendar where new events are created by default, false otherwise. | [optional] 
**IsRemovable** | Pointer to **NullableBool** | Indicates whether this user calendar can be deleted from the user mailbox. | [optional] 
**IsTallyingResponses** | Pointer to **NullableBool** | Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users&#39; primary calendars support tracking of meeting responses. | [optional] 
**Name** | Pointer to **NullableString** | The calendar name. | [optional] 
**Owner** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) | If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user. | [optional] 
**CalendarPermissions** | Pointer to [**[]MicrosoftGraphCalendarPermission**](MicrosoftGraphCalendarPermission.md) | The permissions of the users with whom the calendar is shared. | [optional] 
**CalendarView** | Pointer to [**[]MicrosoftGraphEvent**](MicrosoftGraphEvent.md) | The calendar view for the calendar. Navigation property. Read-only. | [optional] 
**Events** | Pointer to [**[]MicrosoftGraphEvent**](MicrosoftGraphEvent.md) | The events in the calendar. Navigation property. Read-only. | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md) | The collection of multi-value extended properties defined for the calendar. Read-only. Nullable. | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md) | The collection of single-value extended properties defined for the calendar. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphCalendar

`func NewMicrosoftGraphCalendar() *MicrosoftGraphCalendar`

NewMicrosoftGraphCalendar instantiates a new MicrosoftGraphCalendar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCalendarWithDefaults

`func NewMicrosoftGraphCalendarWithDefaults() *MicrosoftGraphCalendar`

NewMicrosoftGraphCalendarWithDefaults instantiates a new MicrosoftGraphCalendar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphCalendar) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphCalendar) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphCalendar) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphCalendar) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllowedOnlineMeetingProviders

`func (o *MicrosoftGraphCalendar) GetAllowedOnlineMeetingProviders() []*AnyOfmicrosoftGraphOnlineMeetingProviderType`

GetAllowedOnlineMeetingProviders returns the AllowedOnlineMeetingProviders field if non-nil, zero value otherwise.

### GetAllowedOnlineMeetingProvidersOk

`func (o *MicrosoftGraphCalendar) GetAllowedOnlineMeetingProvidersOk() (*[]*AnyOfmicrosoftGraphOnlineMeetingProviderType, bool)`

GetAllowedOnlineMeetingProvidersOk returns a tuple with the AllowedOnlineMeetingProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOnlineMeetingProviders

`func (o *MicrosoftGraphCalendar) SetAllowedOnlineMeetingProviders(v []*AnyOfmicrosoftGraphOnlineMeetingProviderType)`

SetAllowedOnlineMeetingProviders sets AllowedOnlineMeetingProviders field to given value.

### HasAllowedOnlineMeetingProviders

`func (o *MicrosoftGraphCalendar) HasAllowedOnlineMeetingProviders() bool`

HasAllowedOnlineMeetingProviders returns a boolean if a field has been set.

### GetCanEdit

`func (o *MicrosoftGraphCalendar) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *MicrosoftGraphCalendar) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *MicrosoftGraphCalendar) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *MicrosoftGraphCalendar) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### SetCanEditNil

`func (o *MicrosoftGraphCalendar) SetCanEditNil(b bool)`

 SetCanEditNil sets the value for CanEdit to be an explicit nil

### UnsetCanEdit
`func (o *MicrosoftGraphCalendar) UnsetCanEdit()`

UnsetCanEdit ensures that no value is present for CanEdit, not even an explicit nil
### GetCanShare

`func (o *MicrosoftGraphCalendar) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *MicrosoftGraphCalendar) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *MicrosoftGraphCalendar) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.

### HasCanShare

`func (o *MicrosoftGraphCalendar) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### SetCanShareNil

`func (o *MicrosoftGraphCalendar) SetCanShareNil(b bool)`

 SetCanShareNil sets the value for CanShare to be an explicit nil

### UnsetCanShare
`func (o *MicrosoftGraphCalendar) UnsetCanShare()`

UnsetCanShare ensures that no value is present for CanShare, not even an explicit nil
### GetCanViewPrivateItems

`func (o *MicrosoftGraphCalendar) GetCanViewPrivateItems() bool`

GetCanViewPrivateItems returns the CanViewPrivateItems field if non-nil, zero value otherwise.

### GetCanViewPrivateItemsOk

`func (o *MicrosoftGraphCalendar) GetCanViewPrivateItemsOk() (*bool, bool)`

GetCanViewPrivateItemsOk returns a tuple with the CanViewPrivateItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewPrivateItems

`func (o *MicrosoftGraphCalendar) SetCanViewPrivateItems(v bool)`

SetCanViewPrivateItems sets CanViewPrivateItems field to given value.

### HasCanViewPrivateItems

`func (o *MicrosoftGraphCalendar) HasCanViewPrivateItems() bool`

HasCanViewPrivateItems returns a boolean if a field has been set.

### SetCanViewPrivateItemsNil

`func (o *MicrosoftGraphCalendar) SetCanViewPrivateItemsNil(b bool)`

 SetCanViewPrivateItemsNil sets the value for CanViewPrivateItems to be an explicit nil

### UnsetCanViewPrivateItems
`func (o *MicrosoftGraphCalendar) UnsetCanViewPrivateItems()`

UnsetCanViewPrivateItems ensures that no value is present for CanViewPrivateItems, not even an explicit nil
### GetChangeKey

`func (o *MicrosoftGraphCalendar) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphCalendar) GetChangeKeyOk() (*string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKey

`func (o *MicrosoftGraphCalendar) SetChangeKey(v string)`

SetChangeKey sets ChangeKey field to given value.

### HasChangeKey

`func (o *MicrosoftGraphCalendar) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKeyNil

`func (o *MicrosoftGraphCalendar) SetChangeKeyNil(b bool)`

 SetChangeKeyNil sets the value for ChangeKey to be an explicit nil

### UnsetChangeKey
`func (o *MicrosoftGraphCalendar) UnsetChangeKey()`

UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
### GetColor

`func (o *MicrosoftGraphCalendar) GetColor() AnyOfmicrosoftGraphCalendarColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MicrosoftGraphCalendar) GetColorOk() (*AnyOfmicrosoftGraphCalendarColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *MicrosoftGraphCalendar) SetColor(v AnyOfmicrosoftGraphCalendarColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *MicrosoftGraphCalendar) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *MicrosoftGraphCalendar) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *MicrosoftGraphCalendar) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDefaultOnlineMeetingProvider

`func (o *MicrosoftGraphCalendar) GetDefaultOnlineMeetingProvider() AnyOfmicrosoftGraphOnlineMeetingProviderType`

GetDefaultOnlineMeetingProvider returns the DefaultOnlineMeetingProvider field if non-nil, zero value otherwise.

### GetDefaultOnlineMeetingProviderOk

`func (o *MicrosoftGraphCalendar) GetDefaultOnlineMeetingProviderOk() (*AnyOfmicrosoftGraphOnlineMeetingProviderType, bool)`

GetDefaultOnlineMeetingProviderOk returns a tuple with the DefaultOnlineMeetingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOnlineMeetingProvider

`func (o *MicrosoftGraphCalendar) SetDefaultOnlineMeetingProvider(v AnyOfmicrosoftGraphOnlineMeetingProviderType)`

SetDefaultOnlineMeetingProvider sets DefaultOnlineMeetingProvider field to given value.

### HasDefaultOnlineMeetingProvider

`func (o *MicrosoftGraphCalendar) HasDefaultOnlineMeetingProvider() bool`

HasDefaultOnlineMeetingProvider returns a boolean if a field has been set.

### SetDefaultOnlineMeetingProviderNil

`func (o *MicrosoftGraphCalendar) SetDefaultOnlineMeetingProviderNil(b bool)`

 SetDefaultOnlineMeetingProviderNil sets the value for DefaultOnlineMeetingProvider to be an explicit nil

### UnsetDefaultOnlineMeetingProvider
`func (o *MicrosoftGraphCalendar) UnsetDefaultOnlineMeetingProvider()`

UnsetDefaultOnlineMeetingProvider ensures that no value is present for DefaultOnlineMeetingProvider, not even an explicit nil
### GetHexColor

`func (o *MicrosoftGraphCalendar) GetHexColor() string`

GetHexColor returns the HexColor field if non-nil, zero value otherwise.

### GetHexColorOk

`func (o *MicrosoftGraphCalendar) GetHexColorOk() (*string, bool)`

GetHexColorOk returns a tuple with the HexColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexColor

`func (o *MicrosoftGraphCalendar) SetHexColor(v string)`

SetHexColor sets HexColor field to given value.

### HasHexColor

`func (o *MicrosoftGraphCalendar) HasHexColor() bool`

HasHexColor returns a boolean if a field has been set.

### SetHexColorNil

`func (o *MicrosoftGraphCalendar) SetHexColorNil(b bool)`

 SetHexColorNil sets the value for HexColor to be an explicit nil

### UnsetHexColor
`func (o *MicrosoftGraphCalendar) UnsetHexColor()`

UnsetHexColor ensures that no value is present for HexColor, not even an explicit nil
### GetIsDefaultCalendar

`func (o *MicrosoftGraphCalendar) GetIsDefaultCalendar() bool`

GetIsDefaultCalendar returns the IsDefaultCalendar field if non-nil, zero value otherwise.

### GetIsDefaultCalendarOk

`func (o *MicrosoftGraphCalendar) GetIsDefaultCalendarOk() (*bool, bool)`

GetIsDefaultCalendarOk returns a tuple with the IsDefaultCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultCalendar

`func (o *MicrosoftGraphCalendar) SetIsDefaultCalendar(v bool)`

SetIsDefaultCalendar sets IsDefaultCalendar field to given value.

### HasIsDefaultCalendar

`func (o *MicrosoftGraphCalendar) HasIsDefaultCalendar() bool`

HasIsDefaultCalendar returns a boolean if a field has been set.

### SetIsDefaultCalendarNil

`func (o *MicrosoftGraphCalendar) SetIsDefaultCalendarNil(b bool)`

 SetIsDefaultCalendarNil sets the value for IsDefaultCalendar to be an explicit nil

### UnsetIsDefaultCalendar
`func (o *MicrosoftGraphCalendar) UnsetIsDefaultCalendar()`

UnsetIsDefaultCalendar ensures that no value is present for IsDefaultCalendar, not even an explicit nil
### GetIsRemovable

`func (o *MicrosoftGraphCalendar) GetIsRemovable() bool`

GetIsRemovable returns the IsRemovable field if non-nil, zero value otherwise.

### GetIsRemovableOk

`func (o *MicrosoftGraphCalendar) GetIsRemovableOk() (*bool, bool)`

GetIsRemovableOk returns a tuple with the IsRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemovable

`func (o *MicrosoftGraphCalendar) SetIsRemovable(v bool)`

SetIsRemovable sets IsRemovable field to given value.

### HasIsRemovable

`func (o *MicrosoftGraphCalendar) HasIsRemovable() bool`

HasIsRemovable returns a boolean if a field has been set.

### SetIsRemovableNil

`func (o *MicrosoftGraphCalendar) SetIsRemovableNil(b bool)`

 SetIsRemovableNil sets the value for IsRemovable to be an explicit nil

### UnsetIsRemovable
`func (o *MicrosoftGraphCalendar) UnsetIsRemovable()`

UnsetIsRemovable ensures that no value is present for IsRemovable, not even an explicit nil
### GetIsTallyingResponses

`func (o *MicrosoftGraphCalendar) GetIsTallyingResponses() bool`

GetIsTallyingResponses returns the IsTallyingResponses field if non-nil, zero value otherwise.

### GetIsTallyingResponsesOk

`func (o *MicrosoftGraphCalendar) GetIsTallyingResponsesOk() (*bool, bool)`

GetIsTallyingResponsesOk returns a tuple with the IsTallyingResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTallyingResponses

`func (o *MicrosoftGraphCalendar) SetIsTallyingResponses(v bool)`

SetIsTallyingResponses sets IsTallyingResponses field to given value.

### HasIsTallyingResponses

`func (o *MicrosoftGraphCalendar) HasIsTallyingResponses() bool`

HasIsTallyingResponses returns a boolean if a field has been set.

### SetIsTallyingResponsesNil

`func (o *MicrosoftGraphCalendar) SetIsTallyingResponsesNil(b bool)`

 SetIsTallyingResponsesNil sets the value for IsTallyingResponses to be an explicit nil

### UnsetIsTallyingResponses
`func (o *MicrosoftGraphCalendar) UnsetIsTallyingResponses()`

UnsetIsTallyingResponses ensures that no value is present for IsTallyingResponses, not even an explicit nil
### GetName

`func (o *MicrosoftGraphCalendar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphCalendar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphCalendar) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphCalendar) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphCalendar) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphCalendar) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwner

`func (o *MicrosoftGraphCalendar) GetOwner() AnyOfmicrosoftGraphEmailAddress`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphCalendar) GetOwnerOk() (*AnyOfmicrosoftGraphEmailAddress, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MicrosoftGraphCalendar) SetOwner(v AnyOfmicrosoftGraphEmailAddress)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MicrosoftGraphCalendar) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *MicrosoftGraphCalendar) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *MicrosoftGraphCalendar) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetCalendarPermissions

`func (o *MicrosoftGraphCalendar) GetCalendarPermissions() []MicrosoftGraphCalendarPermission`

GetCalendarPermissions returns the CalendarPermissions field if non-nil, zero value otherwise.

### GetCalendarPermissionsOk

`func (o *MicrosoftGraphCalendar) GetCalendarPermissionsOk() (*[]MicrosoftGraphCalendarPermission, bool)`

GetCalendarPermissionsOk returns a tuple with the CalendarPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarPermissions

`func (o *MicrosoftGraphCalendar) SetCalendarPermissions(v []MicrosoftGraphCalendarPermission)`

SetCalendarPermissions sets CalendarPermissions field to given value.

### HasCalendarPermissions

`func (o *MicrosoftGraphCalendar) HasCalendarPermissions() bool`

HasCalendarPermissions returns a boolean if a field has been set.

### GetCalendarView

`func (o *MicrosoftGraphCalendar) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *MicrosoftGraphCalendar) GetCalendarViewOk() (*[]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarView

`func (o *MicrosoftGraphCalendar) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView sets CalendarView field to given value.

### HasCalendarView

`func (o *MicrosoftGraphCalendar) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### GetEvents

`func (o *MicrosoftGraphCalendar) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MicrosoftGraphCalendar) GetEventsOk() (*[]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MicrosoftGraphCalendar) SetEvents(v []MicrosoftGraphEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *MicrosoftGraphCalendar) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphCalendar) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphCalendar) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphCalendar) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphCalendar) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphCalendar) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphCalendar) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphCalendar) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphCalendar) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


