# Calendar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewCalendar

`func NewCalendar() *Calendar`

NewCalendar instantiates a new Calendar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarWithDefaults

`func NewCalendarWithDefaults() *Calendar`

NewCalendarWithDefaults instantiates a new Calendar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedOnlineMeetingProviders

`func (o *Calendar) GetAllowedOnlineMeetingProviders() []*AnyOfmicrosoftGraphOnlineMeetingProviderType`

GetAllowedOnlineMeetingProviders returns the AllowedOnlineMeetingProviders field if non-nil, zero value otherwise.

### GetAllowedOnlineMeetingProvidersOk

`func (o *Calendar) GetAllowedOnlineMeetingProvidersOk() (*[]*AnyOfmicrosoftGraphOnlineMeetingProviderType, bool)`

GetAllowedOnlineMeetingProvidersOk returns a tuple with the AllowedOnlineMeetingProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOnlineMeetingProviders

`func (o *Calendar) SetAllowedOnlineMeetingProviders(v []*AnyOfmicrosoftGraphOnlineMeetingProviderType)`

SetAllowedOnlineMeetingProviders sets AllowedOnlineMeetingProviders field to given value.

### HasAllowedOnlineMeetingProviders

`func (o *Calendar) HasAllowedOnlineMeetingProviders() bool`

HasAllowedOnlineMeetingProviders returns a boolean if a field has been set.

### GetCanEdit

`func (o *Calendar) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *Calendar) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *Calendar) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *Calendar) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### SetCanEditNil

`func (o *Calendar) SetCanEditNil(b bool)`

 SetCanEditNil sets the value for CanEdit to be an explicit nil

### UnsetCanEdit
`func (o *Calendar) UnsetCanEdit()`

UnsetCanEdit ensures that no value is present for CanEdit, not even an explicit nil
### GetCanShare

`func (o *Calendar) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *Calendar) GetCanShareOk() (*bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanShare

`func (o *Calendar) SetCanShare(v bool)`

SetCanShare sets CanShare field to given value.

### HasCanShare

`func (o *Calendar) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### SetCanShareNil

`func (o *Calendar) SetCanShareNil(b bool)`

 SetCanShareNil sets the value for CanShare to be an explicit nil

### UnsetCanShare
`func (o *Calendar) UnsetCanShare()`

UnsetCanShare ensures that no value is present for CanShare, not even an explicit nil
### GetCanViewPrivateItems

`func (o *Calendar) GetCanViewPrivateItems() bool`

GetCanViewPrivateItems returns the CanViewPrivateItems field if non-nil, zero value otherwise.

### GetCanViewPrivateItemsOk

`func (o *Calendar) GetCanViewPrivateItemsOk() (*bool, bool)`

GetCanViewPrivateItemsOk returns a tuple with the CanViewPrivateItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewPrivateItems

`func (o *Calendar) SetCanViewPrivateItems(v bool)`

SetCanViewPrivateItems sets CanViewPrivateItems field to given value.

### HasCanViewPrivateItems

`func (o *Calendar) HasCanViewPrivateItems() bool`

HasCanViewPrivateItems returns a boolean if a field has been set.

### SetCanViewPrivateItemsNil

`func (o *Calendar) SetCanViewPrivateItemsNil(b bool)`

 SetCanViewPrivateItemsNil sets the value for CanViewPrivateItems to be an explicit nil

### UnsetCanViewPrivateItems
`func (o *Calendar) UnsetCanViewPrivateItems()`

UnsetCanViewPrivateItems ensures that no value is present for CanViewPrivateItems, not even an explicit nil
### GetChangeKey

`func (o *Calendar) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *Calendar) GetChangeKeyOk() (*string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKey

`func (o *Calendar) SetChangeKey(v string)`

SetChangeKey sets ChangeKey field to given value.

### HasChangeKey

`func (o *Calendar) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKeyNil

`func (o *Calendar) SetChangeKeyNil(b bool)`

 SetChangeKeyNil sets the value for ChangeKey to be an explicit nil

### UnsetChangeKey
`func (o *Calendar) UnsetChangeKey()`

UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
### GetColor

`func (o *Calendar) GetColor() AnyOfmicrosoftGraphCalendarColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Calendar) GetColorOk() (*AnyOfmicrosoftGraphCalendarColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Calendar) SetColor(v AnyOfmicrosoftGraphCalendarColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *Calendar) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *Calendar) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *Calendar) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDefaultOnlineMeetingProvider

`func (o *Calendar) GetDefaultOnlineMeetingProvider() AnyOfmicrosoftGraphOnlineMeetingProviderType`

GetDefaultOnlineMeetingProvider returns the DefaultOnlineMeetingProvider field if non-nil, zero value otherwise.

### GetDefaultOnlineMeetingProviderOk

`func (o *Calendar) GetDefaultOnlineMeetingProviderOk() (*AnyOfmicrosoftGraphOnlineMeetingProviderType, bool)`

GetDefaultOnlineMeetingProviderOk returns a tuple with the DefaultOnlineMeetingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOnlineMeetingProvider

`func (o *Calendar) SetDefaultOnlineMeetingProvider(v AnyOfmicrosoftGraphOnlineMeetingProviderType)`

SetDefaultOnlineMeetingProvider sets DefaultOnlineMeetingProvider field to given value.

### HasDefaultOnlineMeetingProvider

`func (o *Calendar) HasDefaultOnlineMeetingProvider() bool`

HasDefaultOnlineMeetingProvider returns a boolean if a field has been set.

### SetDefaultOnlineMeetingProviderNil

`func (o *Calendar) SetDefaultOnlineMeetingProviderNil(b bool)`

 SetDefaultOnlineMeetingProviderNil sets the value for DefaultOnlineMeetingProvider to be an explicit nil

### UnsetDefaultOnlineMeetingProvider
`func (o *Calendar) UnsetDefaultOnlineMeetingProvider()`

UnsetDefaultOnlineMeetingProvider ensures that no value is present for DefaultOnlineMeetingProvider, not even an explicit nil
### GetHexColor

`func (o *Calendar) GetHexColor() string`

GetHexColor returns the HexColor field if non-nil, zero value otherwise.

### GetHexColorOk

`func (o *Calendar) GetHexColorOk() (*string, bool)`

GetHexColorOk returns a tuple with the HexColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexColor

`func (o *Calendar) SetHexColor(v string)`

SetHexColor sets HexColor field to given value.

### HasHexColor

`func (o *Calendar) HasHexColor() bool`

HasHexColor returns a boolean if a field has been set.

### SetHexColorNil

`func (o *Calendar) SetHexColorNil(b bool)`

 SetHexColorNil sets the value for HexColor to be an explicit nil

### UnsetHexColor
`func (o *Calendar) UnsetHexColor()`

UnsetHexColor ensures that no value is present for HexColor, not even an explicit nil
### GetIsDefaultCalendar

`func (o *Calendar) GetIsDefaultCalendar() bool`

GetIsDefaultCalendar returns the IsDefaultCalendar field if non-nil, zero value otherwise.

### GetIsDefaultCalendarOk

`func (o *Calendar) GetIsDefaultCalendarOk() (*bool, bool)`

GetIsDefaultCalendarOk returns a tuple with the IsDefaultCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultCalendar

`func (o *Calendar) SetIsDefaultCalendar(v bool)`

SetIsDefaultCalendar sets IsDefaultCalendar field to given value.

### HasIsDefaultCalendar

`func (o *Calendar) HasIsDefaultCalendar() bool`

HasIsDefaultCalendar returns a boolean if a field has been set.

### SetIsDefaultCalendarNil

`func (o *Calendar) SetIsDefaultCalendarNil(b bool)`

 SetIsDefaultCalendarNil sets the value for IsDefaultCalendar to be an explicit nil

### UnsetIsDefaultCalendar
`func (o *Calendar) UnsetIsDefaultCalendar()`

UnsetIsDefaultCalendar ensures that no value is present for IsDefaultCalendar, not even an explicit nil
### GetIsRemovable

`func (o *Calendar) GetIsRemovable() bool`

GetIsRemovable returns the IsRemovable field if non-nil, zero value otherwise.

### GetIsRemovableOk

`func (o *Calendar) GetIsRemovableOk() (*bool, bool)`

GetIsRemovableOk returns a tuple with the IsRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemovable

`func (o *Calendar) SetIsRemovable(v bool)`

SetIsRemovable sets IsRemovable field to given value.

### HasIsRemovable

`func (o *Calendar) HasIsRemovable() bool`

HasIsRemovable returns a boolean if a field has been set.

### SetIsRemovableNil

`func (o *Calendar) SetIsRemovableNil(b bool)`

 SetIsRemovableNil sets the value for IsRemovable to be an explicit nil

### UnsetIsRemovable
`func (o *Calendar) UnsetIsRemovable()`

UnsetIsRemovable ensures that no value is present for IsRemovable, not even an explicit nil
### GetIsTallyingResponses

`func (o *Calendar) GetIsTallyingResponses() bool`

GetIsTallyingResponses returns the IsTallyingResponses field if non-nil, zero value otherwise.

### GetIsTallyingResponsesOk

`func (o *Calendar) GetIsTallyingResponsesOk() (*bool, bool)`

GetIsTallyingResponsesOk returns a tuple with the IsTallyingResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTallyingResponses

`func (o *Calendar) SetIsTallyingResponses(v bool)`

SetIsTallyingResponses sets IsTallyingResponses field to given value.

### HasIsTallyingResponses

`func (o *Calendar) HasIsTallyingResponses() bool`

HasIsTallyingResponses returns a boolean if a field has been set.

### SetIsTallyingResponsesNil

`func (o *Calendar) SetIsTallyingResponsesNil(b bool)`

 SetIsTallyingResponsesNil sets the value for IsTallyingResponses to be an explicit nil

### UnsetIsTallyingResponses
`func (o *Calendar) UnsetIsTallyingResponses()`

UnsetIsTallyingResponses ensures that no value is present for IsTallyingResponses, not even an explicit nil
### GetName

`func (o *Calendar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Calendar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Calendar) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Calendar) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Calendar) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Calendar) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwner

`func (o *Calendar) GetOwner() AnyOfmicrosoftGraphEmailAddress`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Calendar) GetOwnerOk() (*AnyOfmicrosoftGraphEmailAddress, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Calendar) SetOwner(v AnyOfmicrosoftGraphEmailAddress)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Calendar) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Calendar) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Calendar) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetCalendarPermissions

`func (o *Calendar) GetCalendarPermissions() []MicrosoftGraphCalendarPermission`

GetCalendarPermissions returns the CalendarPermissions field if non-nil, zero value otherwise.

### GetCalendarPermissionsOk

`func (o *Calendar) GetCalendarPermissionsOk() (*[]MicrosoftGraphCalendarPermission, bool)`

GetCalendarPermissionsOk returns a tuple with the CalendarPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarPermissions

`func (o *Calendar) SetCalendarPermissions(v []MicrosoftGraphCalendarPermission)`

SetCalendarPermissions sets CalendarPermissions field to given value.

### HasCalendarPermissions

`func (o *Calendar) HasCalendarPermissions() bool`

HasCalendarPermissions returns a boolean if a field has been set.

### GetCalendarView

`func (o *Calendar) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *Calendar) GetCalendarViewOk() (*[]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarView

`func (o *Calendar) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView sets CalendarView field to given value.

### HasCalendarView

`func (o *Calendar) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### GetEvents

`func (o *Calendar) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Calendar) GetEventsOk() (*[]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Calendar) SetEvents(v []MicrosoftGraphEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Calendar) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetMultiValueExtendedProperties

`func (o *Calendar) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *Calendar) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *Calendar) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *Calendar) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetSingleValueExtendedProperties

`func (o *Calendar) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *Calendar) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *Calendar) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *Calendar) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


