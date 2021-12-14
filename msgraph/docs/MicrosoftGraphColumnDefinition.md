# MicrosoftGraphColumnDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Boolean** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | This column stores boolean values. | [optional] 
**Calculated** | Pointer to [**AnyOfmicrosoftGraphCalculatedColumn**](anyOf&lt;microsoft.graph.calculatedColumn&gt;.md) | This column&#39;s data is calculated based on other columns. | [optional] 
**Choice** | Pointer to [**AnyOfmicrosoftGraphChoiceColumn**](anyOf&lt;microsoft.graph.choiceColumn&gt;.md) | This column stores data from a list of choices. | [optional] 
**ColumnGroup** | Pointer to **NullableString** | For site columns, the name of the group this column belongs to. Helps organize related columns. | [optional] 
**ContentApprovalStatus** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | This column stores content approval status. | [optional] 
**Currency** | Pointer to [**AnyOfmicrosoftGraphCurrencyColumn**](anyOf&lt;microsoft.graph.currencyColumn&gt;.md) | This column stores currency values. | [optional] 
**DateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeColumn**](anyOf&lt;microsoft.graph.dateTimeColumn&gt;.md) | This column stores DateTime values. | [optional] 
**DefaultValue** | Pointer to [**AnyOfmicrosoftGraphDefaultColumnValue**](anyOf&lt;microsoft.graph.defaultColumnValue&gt;.md) | The default value for this column. | [optional] 
**Description** | Pointer to **NullableString** | The user-facing description of the column. | [optional] 
**DisplayName** | Pointer to **NullableString** | The user-facing name of the column. | [optional] 
**EnforceUniqueValues** | Pointer to **NullableBool** | If true, no two list items may have the same value for this column. | [optional] 
**Geolocation** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | This column stores a geolocation. | [optional] 
**Hidden** | Pointer to **NullableBool** | Specifies whether the column is displayed in the user interface. | [optional] 
**HyperlinkOrPicture** | Pointer to [**AnyOfmicrosoftGraphHyperlinkOrPictureColumn**](anyOf&lt;microsoft.graph.hyperlinkOrPictureColumn&gt;.md) | This column stores hyperlink or picture values. | [optional] 
**Indexed** | Pointer to **NullableBool** | Specifies whether the column values can be used for sorting and searching. | [optional] 
**IsDeletable** | Pointer to **NullableBool** | Indicates whether this column can be deleted. | [optional] 
**IsReorderable** | Pointer to **NullableBool** | Indicates whether values in the column can be reordered. Read-only. | [optional] 
**IsSealed** | Pointer to **NullableBool** | Specifies whether the column can be changed. | [optional] 
**Lookup** | Pointer to [**AnyOfmicrosoftGraphLookupColumn**](anyOf&lt;microsoft.graph.lookupColumn&gt;.md) | This column&#39;s data is looked up from another source in the site. | [optional] 
**Name** | Pointer to **NullableString** | The API-facing name of the column as it appears in the [fields][] on a [listItem][]. For the user-facing name, see displayName. | [optional] 
**Number** | Pointer to [**AnyOfmicrosoftGraphNumberColumn**](anyOf&lt;microsoft.graph.numberColumn&gt;.md) | This column stores number values. | [optional] 
**PersonOrGroup** | Pointer to [**AnyOfmicrosoftGraphPersonOrGroupColumn**](anyOf&lt;microsoft.graph.personOrGroupColumn&gt;.md) | This column stores Person or Group values. | [optional] 
**PropagateChanges** | Pointer to **NullableBool** | If &#39;true&#39;, changes to this column will be propagated to lists that implement the column. | [optional] 
**ReadOnly** | Pointer to **NullableBool** | Specifies whether the column values can be modified. | [optional] 
**Required** | Pointer to **NullableBool** | Specifies whether the column value isn&#39;t optional. | [optional] 
**Term** | Pointer to [**AnyOfmicrosoftGraphTermColumn**](anyOf&lt;microsoft.graph.termColumn&gt;.md) | This column stores taxonomy terms. | [optional] 
**Text** | Pointer to [**AnyOfmicrosoftGraphTextColumn**](anyOf&lt;microsoft.graph.textColumn&gt;.md) | This column stores text values. | [optional] 
**Thumbnail** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | This column stores thumbnail values. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphColumnTypes**](anyOf&lt;microsoft.graph.columnTypes&gt;.md) | For site columns, the type of column. Read-only. | [optional] 
**Validation** | Pointer to [**AnyOfmicrosoftGraphColumnValidation**](anyOf&lt;microsoft.graph.columnValidation&gt;.md) | This column stores validation formula and message for the column. | [optional] 
**SourceColumn** | Pointer to [**AnyOfmicrosoftGraphColumnDefinition**](anyOf&lt;microsoft.graph.columnDefinition&gt;.md) | The source column for the content type column. | [optional] 

## Methods

### NewMicrosoftGraphColumnDefinition

`func NewMicrosoftGraphColumnDefinition() *MicrosoftGraphColumnDefinition`

NewMicrosoftGraphColumnDefinition instantiates a new MicrosoftGraphColumnDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphColumnDefinitionWithDefaults

`func NewMicrosoftGraphColumnDefinitionWithDefaults() *MicrosoftGraphColumnDefinition`

NewMicrosoftGraphColumnDefinitionWithDefaults instantiates a new MicrosoftGraphColumnDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphColumnDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphColumnDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphColumnDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphColumnDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBoolean

`func (o *MicrosoftGraphColumnDefinition) GetBoolean() AnyOfobject`

GetBoolean returns the Boolean field if non-nil, zero value otherwise.

### GetBooleanOk

`func (o *MicrosoftGraphColumnDefinition) GetBooleanOk() (*AnyOfobject, bool)`

GetBooleanOk returns a tuple with the Boolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolean

`func (o *MicrosoftGraphColumnDefinition) SetBoolean(v AnyOfobject)`

SetBoolean sets Boolean field to given value.

### HasBoolean

`func (o *MicrosoftGraphColumnDefinition) HasBoolean() bool`

HasBoolean returns a boolean if a field has been set.

### SetBooleanNil

`func (o *MicrosoftGraphColumnDefinition) SetBooleanNil(b bool)`

 SetBooleanNil sets the value for Boolean to be an explicit nil

### UnsetBoolean
`func (o *MicrosoftGraphColumnDefinition) UnsetBoolean()`

UnsetBoolean ensures that no value is present for Boolean, not even an explicit nil
### GetCalculated

`func (o *MicrosoftGraphColumnDefinition) GetCalculated() AnyOfmicrosoftGraphCalculatedColumn`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *MicrosoftGraphColumnDefinition) GetCalculatedOk() (*AnyOfmicrosoftGraphCalculatedColumn, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculated

`func (o *MicrosoftGraphColumnDefinition) SetCalculated(v AnyOfmicrosoftGraphCalculatedColumn)`

SetCalculated sets Calculated field to given value.

### HasCalculated

`func (o *MicrosoftGraphColumnDefinition) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### SetCalculatedNil

`func (o *MicrosoftGraphColumnDefinition) SetCalculatedNil(b bool)`

 SetCalculatedNil sets the value for Calculated to be an explicit nil

### UnsetCalculated
`func (o *MicrosoftGraphColumnDefinition) UnsetCalculated()`

UnsetCalculated ensures that no value is present for Calculated, not even an explicit nil
### GetChoice

`func (o *MicrosoftGraphColumnDefinition) GetChoice() AnyOfmicrosoftGraphChoiceColumn`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *MicrosoftGraphColumnDefinition) GetChoiceOk() (*AnyOfmicrosoftGraphChoiceColumn, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *MicrosoftGraphColumnDefinition) SetChoice(v AnyOfmicrosoftGraphChoiceColumn)`

SetChoice sets Choice field to given value.

### HasChoice

`func (o *MicrosoftGraphColumnDefinition) HasChoice() bool`

HasChoice returns a boolean if a field has been set.

### SetChoiceNil

`func (o *MicrosoftGraphColumnDefinition) SetChoiceNil(b bool)`

 SetChoiceNil sets the value for Choice to be an explicit nil

### UnsetChoice
`func (o *MicrosoftGraphColumnDefinition) UnsetChoice()`

UnsetChoice ensures that no value is present for Choice, not even an explicit nil
### GetColumnGroup

`func (o *MicrosoftGraphColumnDefinition) GetColumnGroup() string`

GetColumnGroup returns the ColumnGroup field if non-nil, zero value otherwise.

### GetColumnGroupOk

`func (o *MicrosoftGraphColumnDefinition) GetColumnGroupOk() (*string, bool)`

GetColumnGroupOk returns a tuple with the ColumnGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnGroup

`func (o *MicrosoftGraphColumnDefinition) SetColumnGroup(v string)`

SetColumnGroup sets ColumnGroup field to given value.

### HasColumnGroup

`func (o *MicrosoftGraphColumnDefinition) HasColumnGroup() bool`

HasColumnGroup returns a boolean if a field has been set.

### SetColumnGroupNil

`func (o *MicrosoftGraphColumnDefinition) SetColumnGroupNil(b bool)`

 SetColumnGroupNil sets the value for ColumnGroup to be an explicit nil

### UnsetColumnGroup
`func (o *MicrosoftGraphColumnDefinition) UnsetColumnGroup()`

UnsetColumnGroup ensures that no value is present for ColumnGroup, not even an explicit nil
### GetContentApprovalStatus

`func (o *MicrosoftGraphColumnDefinition) GetContentApprovalStatus() AnyOfobject`

GetContentApprovalStatus returns the ContentApprovalStatus field if non-nil, zero value otherwise.

### GetContentApprovalStatusOk

`func (o *MicrosoftGraphColumnDefinition) GetContentApprovalStatusOk() (*AnyOfobject, bool)`

GetContentApprovalStatusOk returns a tuple with the ContentApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentApprovalStatus

`func (o *MicrosoftGraphColumnDefinition) SetContentApprovalStatus(v AnyOfobject)`

SetContentApprovalStatus sets ContentApprovalStatus field to given value.

### HasContentApprovalStatus

`func (o *MicrosoftGraphColumnDefinition) HasContentApprovalStatus() bool`

HasContentApprovalStatus returns a boolean if a field has been set.

### SetContentApprovalStatusNil

`func (o *MicrosoftGraphColumnDefinition) SetContentApprovalStatusNil(b bool)`

 SetContentApprovalStatusNil sets the value for ContentApprovalStatus to be an explicit nil

### UnsetContentApprovalStatus
`func (o *MicrosoftGraphColumnDefinition) UnsetContentApprovalStatus()`

UnsetContentApprovalStatus ensures that no value is present for ContentApprovalStatus, not even an explicit nil
### GetCurrency

`func (o *MicrosoftGraphColumnDefinition) GetCurrency() AnyOfmicrosoftGraphCurrencyColumn`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MicrosoftGraphColumnDefinition) GetCurrencyOk() (*AnyOfmicrosoftGraphCurrencyColumn, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MicrosoftGraphColumnDefinition) SetCurrency(v AnyOfmicrosoftGraphCurrencyColumn)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MicrosoftGraphColumnDefinition) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *MicrosoftGraphColumnDefinition) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *MicrosoftGraphColumnDefinition) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetDateTime

`func (o *MicrosoftGraphColumnDefinition) GetDateTime() AnyOfmicrosoftGraphDateTimeColumn`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *MicrosoftGraphColumnDefinition) GetDateTimeOk() (*AnyOfmicrosoftGraphDateTimeColumn, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *MicrosoftGraphColumnDefinition) SetDateTime(v AnyOfmicrosoftGraphDateTimeColumn)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *MicrosoftGraphColumnDefinition) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTimeNil

`func (o *MicrosoftGraphColumnDefinition) SetDateTimeNil(b bool)`

 SetDateTimeNil sets the value for DateTime to be an explicit nil

### UnsetDateTime
`func (o *MicrosoftGraphColumnDefinition) UnsetDateTime()`

UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil
### GetDefaultValue

`func (o *MicrosoftGraphColumnDefinition) GetDefaultValue() AnyOfmicrosoftGraphDefaultColumnValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MicrosoftGraphColumnDefinition) GetDefaultValueOk() (*AnyOfmicrosoftGraphDefaultColumnValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *MicrosoftGraphColumnDefinition) SetDefaultValue(v AnyOfmicrosoftGraphDefaultColumnValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *MicrosoftGraphColumnDefinition) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *MicrosoftGraphColumnDefinition) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *MicrosoftGraphColumnDefinition) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphColumnDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphColumnDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphColumnDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphColumnDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphColumnDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphColumnDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphColumnDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphColumnDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphColumnDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphColumnDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphColumnDefinition) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphColumnDefinition) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEnforceUniqueValues

`func (o *MicrosoftGraphColumnDefinition) GetEnforceUniqueValues() bool`

GetEnforceUniqueValues returns the EnforceUniqueValues field if non-nil, zero value otherwise.

### GetEnforceUniqueValuesOk

`func (o *MicrosoftGraphColumnDefinition) GetEnforceUniqueValuesOk() (*bool, bool)`

GetEnforceUniqueValuesOk returns a tuple with the EnforceUniqueValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUniqueValues

`func (o *MicrosoftGraphColumnDefinition) SetEnforceUniqueValues(v bool)`

SetEnforceUniqueValues sets EnforceUniqueValues field to given value.

### HasEnforceUniqueValues

`func (o *MicrosoftGraphColumnDefinition) HasEnforceUniqueValues() bool`

HasEnforceUniqueValues returns a boolean if a field has been set.

### SetEnforceUniqueValuesNil

`func (o *MicrosoftGraphColumnDefinition) SetEnforceUniqueValuesNil(b bool)`

 SetEnforceUniqueValuesNil sets the value for EnforceUniqueValues to be an explicit nil

### UnsetEnforceUniqueValues
`func (o *MicrosoftGraphColumnDefinition) UnsetEnforceUniqueValues()`

UnsetEnforceUniqueValues ensures that no value is present for EnforceUniqueValues, not even an explicit nil
### GetGeolocation

`func (o *MicrosoftGraphColumnDefinition) GetGeolocation() AnyOfobject`

GetGeolocation returns the Geolocation field if non-nil, zero value otherwise.

### GetGeolocationOk

`func (o *MicrosoftGraphColumnDefinition) GetGeolocationOk() (*AnyOfobject, bool)`

GetGeolocationOk returns a tuple with the Geolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocation

`func (o *MicrosoftGraphColumnDefinition) SetGeolocation(v AnyOfobject)`

SetGeolocation sets Geolocation field to given value.

### HasGeolocation

`func (o *MicrosoftGraphColumnDefinition) HasGeolocation() bool`

HasGeolocation returns a boolean if a field has been set.

### SetGeolocationNil

`func (o *MicrosoftGraphColumnDefinition) SetGeolocationNil(b bool)`

 SetGeolocationNil sets the value for Geolocation to be an explicit nil

### UnsetGeolocation
`func (o *MicrosoftGraphColumnDefinition) UnsetGeolocation()`

UnsetGeolocation ensures that no value is present for Geolocation, not even an explicit nil
### GetHidden

`func (o *MicrosoftGraphColumnDefinition) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MicrosoftGraphColumnDefinition) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *MicrosoftGraphColumnDefinition) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *MicrosoftGraphColumnDefinition) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *MicrosoftGraphColumnDefinition) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *MicrosoftGraphColumnDefinition) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
### GetHyperlinkOrPicture

`func (o *MicrosoftGraphColumnDefinition) GetHyperlinkOrPicture() AnyOfmicrosoftGraphHyperlinkOrPictureColumn`

GetHyperlinkOrPicture returns the HyperlinkOrPicture field if non-nil, zero value otherwise.

### GetHyperlinkOrPictureOk

`func (o *MicrosoftGraphColumnDefinition) GetHyperlinkOrPictureOk() (*AnyOfmicrosoftGraphHyperlinkOrPictureColumn, bool)`

GetHyperlinkOrPictureOk returns a tuple with the HyperlinkOrPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperlinkOrPicture

`func (o *MicrosoftGraphColumnDefinition) SetHyperlinkOrPicture(v AnyOfmicrosoftGraphHyperlinkOrPictureColumn)`

SetHyperlinkOrPicture sets HyperlinkOrPicture field to given value.

### HasHyperlinkOrPicture

`func (o *MicrosoftGraphColumnDefinition) HasHyperlinkOrPicture() bool`

HasHyperlinkOrPicture returns a boolean if a field has been set.

### SetHyperlinkOrPictureNil

`func (o *MicrosoftGraphColumnDefinition) SetHyperlinkOrPictureNil(b bool)`

 SetHyperlinkOrPictureNil sets the value for HyperlinkOrPicture to be an explicit nil

### UnsetHyperlinkOrPicture
`func (o *MicrosoftGraphColumnDefinition) UnsetHyperlinkOrPicture()`

UnsetHyperlinkOrPicture ensures that no value is present for HyperlinkOrPicture, not even an explicit nil
### GetIndexed

`func (o *MicrosoftGraphColumnDefinition) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *MicrosoftGraphColumnDefinition) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *MicrosoftGraphColumnDefinition) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *MicrosoftGraphColumnDefinition) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### SetIndexedNil

`func (o *MicrosoftGraphColumnDefinition) SetIndexedNil(b bool)`

 SetIndexedNil sets the value for Indexed to be an explicit nil

### UnsetIndexed
`func (o *MicrosoftGraphColumnDefinition) UnsetIndexed()`

UnsetIndexed ensures that no value is present for Indexed, not even an explicit nil
### GetIsDeletable

`func (o *MicrosoftGraphColumnDefinition) GetIsDeletable() bool`

GetIsDeletable returns the IsDeletable field if non-nil, zero value otherwise.

### GetIsDeletableOk

`func (o *MicrosoftGraphColumnDefinition) GetIsDeletableOk() (*bool, bool)`

GetIsDeletableOk returns a tuple with the IsDeletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletable

`func (o *MicrosoftGraphColumnDefinition) SetIsDeletable(v bool)`

SetIsDeletable sets IsDeletable field to given value.

### HasIsDeletable

`func (o *MicrosoftGraphColumnDefinition) HasIsDeletable() bool`

HasIsDeletable returns a boolean if a field has been set.

### SetIsDeletableNil

`func (o *MicrosoftGraphColumnDefinition) SetIsDeletableNil(b bool)`

 SetIsDeletableNil sets the value for IsDeletable to be an explicit nil

### UnsetIsDeletable
`func (o *MicrosoftGraphColumnDefinition) UnsetIsDeletable()`

UnsetIsDeletable ensures that no value is present for IsDeletable, not even an explicit nil
### GetIsReorderable

`func (o *MicrosoftGraphColumnDefinition) GetIsReorderable() bool`

GetIsReorderable returns the IsReorderable field if non-nil, zero value otherwise.

### GetIsReorderableOk

`func (o *MicrosoftGraphColumnDefinition) GetIsReorderableOk() (*bool, bool)`

GetIsReorderableOk returns a tuple with the IsReorderable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReorderable

`func (o *MicrosoftGraphColumnDefinition) SetIsReorderable(v bool)`

SetIsReorderable sets IsReorderable field to given value.

### HasIsReorderable

`func (o *MicrosoftGraphColumnDefinition) HasIsReorderable() bool`

HasIsReorderable returns a boolean if a field has been set.

### SetIsReorderableNil

`func (o *MicrosoftGraphColumnDefinition) SetIsReorderableNil(b bool)`

 SetIsReorderableNil sets the value for IsReorderable to be an explicit nil

### UnsetIsReorderable
`func (o *MicrosoftGraphColumnDefinition) UnsetIsReorderable()`

UnsetIsReorderable ensures that no value is present for IsReorderable, not even an explicit nil
### GetIsSealed

`func (o *MicrosoftGraphColumnDefinition) GetIsSealed() bool`

GetIsSealed returns the IsSealed field if non-nil, zero value otherwise.

### GetIsSealedOk

`func (o *MicrosoftGraphColumnDefinition) GetIsSealedOk() (*bool, bool)`

GetIsSealedOk returns a tuple with the IsSealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSealed

`func (o *MicrosoftGraphColumnDefinition) SetIsSealed(v bool)`

SetIsSealed sets IsSealed field to given value.

### HasIsSealed

`func (o *MicrosoftGraphColumnDefinition) HasIsSealed() bool`

HasIsSealed returns a boolean if a field has been set.

### SetIsSealedNil

`func (o *MicrosoftGraphColumnDefinition) SetIsSealedNil(b bool)`

 SetIsSealedNil sets the value for IsSealed to be an explicit nil

### UnsetIsSealed
`func (o *MicrosoftGraphColumnDefinition) UnsetIsSealed()`

UnsetIsSealed ensures that no value is present for IsSealed, not even an explicit nil
### GetLookup

`func (o *MicrosoftGraphColumnDefinition) GetLookup() AnyOfmicrosoftGraphLookupColumn`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *MicrosoftGraphColumnDefinition) GetLookupOk() (*AnyOfmicrosoftGraphLookupColumn, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *MicrosoftGraphColumnDefinition) SetLookup(v AnyOfmicrosoftGraphLookupColumn)`

SetLookup sets Lookup field to given value.

### HasLookup

`func (o *MicrosoftGraphColumnDefinition) HasLookup() bool`

HasLookup returns a boolean if a field has been set.

### SetLookupNil

`func (o *MicrosoftGraphColumnDefinition) SetLookupNil(b bool)`

 SetLookupNil sets the value for Lookup to be an explicit nil

### UnsetLookup
`func (o *MicrosoftGraphColumnDefinition) UnsetLookup()`

UnsetLookup ensures that no value is present for Lookup, not even an explicit nil
### GetName

`func (o *MicrosoftGraphColumnDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphColumnDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphColumnDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphColumnDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphColumnDefinition) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphColumnDefinition) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNumber

`func (o *MicrosoftGraphColumnDefinition) GetNumber() AnyOfmicrosoftGraphNumberColumn`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *MicrosoftGraphColumnDefinition) GetNumberOk() (*AnyOfmicrosoftGraphNumberColumn, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *MicrosoftGraphColumnDefinition) SetNumber(v AnyOfmicrosoftGraphNumberColumn)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *MicrosoftGraphColumnDefinition) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *MicrosoftGraphColumnDefinition) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *MicrosoftGraphColumnDefinition) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetPersonOrGroup

`func (o *MicrosoftGraphColumnDefinition) GetPersonOrGroup() AnyOfmicrosoftGraphPersonOrGroupColumn`

GetPersonOrGroup returns the PersonOrGroup field if non-nil, zero value otherwise.

### GetPersonOrGroupOk

`func (o *MicrosoftGraphColumnDefinition) GetPersonOrGroupOk() (*AnyOfmicrosoftGraphPersonOrGroupColumn, bool)`

GetPersonOrGroupOk returns a tuple with the PersonOrGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonOrGroup

`func (o *MicrosoftGraphColumnDefinition) SetPersonOrGroup(v AnyOfmicrosoftGraphPersonOrGroupColumn)`

SetPersonOrGroup sets PersonOrGroup field to given value.

### HasPersonOrGroup

`func (o *MicrosoftGraphColumnDefinition) HasPersonOrGroup() bool`

HasPersonOrGroup returns a boolean if a field has been set.

### SetPersonOrGroupNil

`func (o *MicrosoftGraphColumnDefinition) SetPersonOrGroupNil(b bool)`

 SetPersonOrGroupNil sets the value for PersonOrGroup to be an explicit nil

### UnsetPersonOrGroup
`func (o *MicrosoftGraphColumnDefinition) UnsetPersonOrGroup()`

UnsetPersonOrGroup ensures that no value is present for PersonOrGroup, not even an explicit nil
### GetPropagateChanges

`func (o *MicrosoftGraphColumnDefinition) GetPropagateChanges() bool`

GetPropagateChanges returns the PropagateChanges field if non-nil, zero value otherwise.

### GetPropagateChangesOk

`func (o *MicrosoftGraphColumnDefinition) GetPropagateChangesOk() (*bool, bool)`

GetPropagateChangesOk returns a tuple with the PropagateChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagateChanges

`func (o *MicrosoftGraphColumnDefinition) SetPropagateChanges(v bool)`

SetPropagateChanges sets PropagateChanges field to given value.

### HasPropagateChanges

`func (o *MicrosoftGraphColumnDefinition) HasPropagateChanges() bool`

HasPropagateChanges returns a boolean if a field has been set.

### SetPropagateChangesNil

`func (o *MicrosoftGraphColumnDefinition) SetPropagateChangesNil(b bool)`

 SetPropagateChangesNil sets the value for PropagateChanges to be an explicit nil

### UnsetPropagateChanges
`func (o *MicrosoftGraphColumnDefinition) UnsetPropagateChanges()`

UnsetPropagateChanges ensures that no value is present for PropagateChanges, not even an explicit nil
### GetReadOnly

`func (o *MicrosoftGraphColumnDefinition) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *MicrosoftGraphColumnDefinition) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *MicrosoftGraphColumnDefinition) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *MicrosoftGraphColumnDefinition) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *MicrosoftGraphColumnDefinition) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *MicrosoftGraphColumnDefinition) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetRequired

`func (o *MicrosoftGraphColumnDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *MicrosoftGraphColumnDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *MicrosoftGraphColumnDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *MicrosoftGraphColumnDefinition) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *MicrosoftGraphColumnDefinition) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *MicrosoftGraphColumnDefinition) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetTerm

`func (o *MicrosoftGraphColumnDefinition) GetTerm() AnyOfmicrosoftGraphTermColumn`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *MicrosoftGraphColumnDefinition) GetTermOk() (*AnyOfmicrosoftGraphTermColumn, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *MicrosoftGraphColumnDefinition) SetTerm(v AnyOfmicrosoftGraphTermColumn)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *MicrosoftGraphColumnDefinition) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### SetTermNil

`func (o *MicrosoftGraphColumnDefinition) SetTermNil(b bool)`

 SetTermNil sets the value for Term to be an explicit nil

### UnsetTerm
`func (o *MicrosoftGraphColumnDefinition) UnsetTerm()`

UnsetTerm ensures that no value is present for Term, not even an explicit nil
### GetText

`func (o *MicrosoftGraphColumnDefinition) GetText() AnyOfmicrosoftGraphTextColumn`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MicrosoftGraphColumnDefinition) GetTextOk() (*AnyOfmicrosoftGraphTextColumn, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MicrosoftGraphColumnDefinition) SetText(v AnyOfmicrosoftGraphTextColumn)`

SetText sets Text field to given value.

### HasText

`func (o *MicrosoftGraphColumnDefinition) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *MicrosoftGraphColumnDefinition) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *MicrosoftGraphColumnDefinition) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetThumbnail

`func (o *MicrosoftGraphColumnDefinition) GetThumbnail() AnyOfobject`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *MicrosoftGraphColumnDefinition) GetThumbnailOk() (*AnyOfobject, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *MicrosoftGraphColumnDefinition) SetThumbnail(v AnyOfobject)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *MicrosoftGraphColumnDefinition) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### SetThumbnailNil

`func (o *MicrosoftGraphColumnDefinition) SetThumbnailNil(b bool)`

 SetThumbnailNil sets the value for Thumbnail to be an explicit nil

### UnsetThumbnail
`func (o *MicrosoftGraphColumnDefinition) UnsetThumbnail()`

UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
### GetType

`func (o *MicrosoftGraphColumnDefinition) GetType() AnyOfmicrosoftGraphColumnTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphColumnDefinition) GetTypeOk() (*AnyOfmicrosoftGraphColumnTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphColumnDefinition) SetType(v AnyOfmicrosoftGraphColumnTypes)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphColumnDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphColumnDefinition) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphColumnDefinition) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValidation

`func (o *MicrosoftGraphColumnDefinition) GetValidation() AnyOfmicrosoftGraphColumnValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *MicrosoftGraphColumnDefinition) GetValidationOk() (*AnyOfmicrosoftGraphColumnValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *MicrosoftGraphColumnDefinition) SetValidation(v AnyOfmicrosoftGraphColumnValidation)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *MicrosoftGraphColumnDefinition) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *MicrosoftGraphColumnDefinition) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *MicrosoftGraphColumnDefinition) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetSourceColumn

`func (o *MicrosoftGraphColumnDefinition) GetSourceColumn() AnyOfmicrosoftGraphColumnDefinition`

GetSourceColumn returns the SourceColumn field if non-nil, zero value otherwise.

### GetSourceColumnOk

`func (o *MicrosoftGraphColumnDefinition) GetSourceColumnOk() (*AnyOfmicrosoftGraphColumnDefinition, bool)`

GetSourceColumnOk returns a tuple with the SourceColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceColumn

`func (o *MicrosoftGraphColumnDefinition) SetSourceColumn(v AnyOfmicrosoftGraphColumnDefinition)`

SetSourceColumn sets SourceColumn field to given value.

### HasSourceColumn

`func (o *MicrosoftGraphColumnDefinition) HasSourceColumn() bool`

HasSourceColumn returns a boolean if a field has been set.

### SetSourceColumnNil

`func (o *MicrosoftGraphColumnDefinition) SetSourceColumnNil(b bool)`

 SetSourceColumnNil sets the value for SourceColumn to be an explicit nil

### UnsetSourceColumn
`func (o *MicrosoftGraphColumnDefinition) UnsetSourceColumn()`

UnsetSourceColumn ensures that no value is present for SourceColumn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


