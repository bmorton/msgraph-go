# ColumnDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewColumnDefinition

`func NewColumnDefinition() *ColumnDefinition`

NewColumnDefinition instantiates a new ColumnDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnDefinitionWithDefaults

`func NewColumnDefinitionWithDefaults() *ColumnDefinition`

NewColumnDefinitionWithDefaults instantiates a new ColumnDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoolean

`func (o *ColumnDefinition) GetBoolean() AnyOfobject`

GetBoolean returns the Boolean field if non-nil, zero value otherwise.

### GetBooleanOk

`func (o *ColumnDefinition) GetBooleanOk() (*AnyOfobject, bool)`

GetBooleanOk returns a tuple with the Boolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolean

`func (o *ColumnDefinition) SetBoolean(v AnyOfobject)`

SetBoolean sets Boolean field to given value.

### HasBoolean

`func (o *ColumnDefinition) HasBoolean() bool`

HasBoolean returns a boolean if a field has been set.

### SetBooleanNil

`func (o *ColumnDefinition) SetBooleanNil(b bool)`

 SetBooleanNil sets the value for Boolean to be an explicit nil

### UnsetBoolean
`func (o *ColumnDefinition) UnsetBoolean()`

UnsetBoolean ensures that no value is present for Boolean, not even an explicit nil
### GetCalculated

`func (o *ColumnDefinition) GetCalculated() AnyOfmicrosoftGraphCalculatedColumn`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *ColumnDefinition) GetCalculatedOk() (*AnyOfmicrosoftGraphCalculatedColumn, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculated

`func (o *ColumnDefinition) SetCalculated(v AnyOfmicrosoftGraphCalculatedColumn)`

SetCalculated sets Calculated field to given value.

### HasCalculated

`func (o *ColumnDefinition) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### SetCalculatedNil

`func (o *ColumnDefinition) SetCalculatedNil(b bool)`

 SetCalculatedNil sets the value for Calculated to be an explicit nil

### UnsetCalculated
`func (o *ColumnDefinition) UnsetCalculated()`

UnsetCalculated ensures that no value is present for Calculated, not even an explicit nil
### GetChoice

`func (o *ColumnDefinition) GetChoice() AnyOfmicrosoftGraphChoiceColumn`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *ColumnDefinition) GetChoiceOk() (*AnyOfmicrosoftGraphChoiceColumn, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *ColumnDefinition) SetChoice(v AnyOfmicrosoftGraphChoiceColumn)`

SetChoice sets Choice field to given value.

### HasChoice

`func (o *ColumnDefinition) HasChoice() bool`

HasChoice returns a boolean if a field has been set.

### SetChoiceNil

`func (o *ColumnDefinition) SetChoiceNil(b bool)`

 SetChoiceNil sets the value for Choice to be an explicit nil

### UnsetChoice
`func (o *ColumnDefinition) UnsetChoice()`

UnsetChoice ensures that no value is present for Choice, not even an explicit nil
### GetColumnGroup

`func (o *ColumnDefinition) GetColumnGroup() string`

GetColumnGroup returns the ColumnGroup field if non-nil, zero value otherwise.

### GetColumnGroupOk

`func (o *ColumnDefinition) GetColumnGroupOk() (*string, bool)`

GetColumnGroupOk returns a tuple with the ColumnGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnGroup

`func (o *ColumnDefinition) SetColumnGroup(v string)`

SetColumnGroup sets ColumnGroup field to given value.

### HasColumnGroup

`func (o *ColumnDefinition) HasColumnGroup() bool`

HasColumnGroup returns a boolean if a field has been set.

### SetColumnGroupNil

`func (o *ColumnDefinition) SetColumnGroupNil(b bool)`

 SetColumnGroupNil sets the value for ColumnGroup to be an explicit nil

### UnsetColumnGroup
`func (o *ColumnDefinition) UnsetColumnGroup()`

UnsetColumnGroup ensures that no value is present for ColumnGroup, not even an explicit nil
### GetContentApprovalStatus

`func (o *ColumnDefinition) GetContentApprovalStatus() AnyOfobject`

GetContentApprovalStatus returns the ContentApprovalStatus field if non-nil, zero value otherwise.

### GetContentApprovalStatusOk

`func (o *ColumnDefinition) GetContentApprovalStatusOk() (*AnyOfobject, bool)`

GetContentApprovalStatusOk returns a tuple with the ContentApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentApprovalStatus

`func (o *ColumnDefinition) SetContentApprovalStatus(v AnyOfobject)`

SetContentApprovalStatus sets ContentApprovalStatus field to given value.

### HasContentApprovalStatus

`func (o *ColumnDefinition) HasContentApprovalStatus() bool`

HasContentApprovalStatus returns a boolean if a field has been set.

### SetContentApprovalStatusNil

`func (o *ColumnDefinition) SetContentApprovalStatusNil(b bool)`

 SetContentApprovalStatusNil sets the value for ContentApprovalStatus to be an explicit nil

### UnsetContentApprovalStatus
`func (o *ColumnDefinition) UnsetContentApprovalStatus()`

UnsetContentApprovalStatus ensures that no value is present for ContentApprovalStatus, not even an explicit nil
### GetCurrency

`func (o *ColumnDefinition) GetCurrency() AnyOfmicrosoftGraphCurrencyColumn`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ColumnDefinition) GetCurrencyOk() (*AnyOfmicrosoftGraphCurrencyColumn, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ColumnDefinition) SetCurrency(v AnyOfmicrosoftGraphCurrencyColumn)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ColumnDefinition) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ColumnDefinition) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ColumnDefinition) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetDateTime

`func (o *ColumnDefinition) GetDateTime() AnyOfmicrosoftGraphDateTimeColumn`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *ColumnDefinition) GetDateTimeOk() (*AnyOfmicrosoftGraphDateTimeColumn, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *ColumnDefinition) SetDateTime(v AnyOfmicrosoftGraphDateTimeColumn)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *ColumnDefinition) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTimeNil

`func (o *ColumnDefinition) SetDateTimeNil(b bool)`

 SetDateTimeNil sets the value for DateTime to be an explicit nil

### UnsetDateTime
`func (o *ColumnDefinition) UnsetDateTime()`

UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil
### GetDefaultValue

`func (o *ColumnDefinition) GetDefaultValue() AnyOfmicrosoftGraphDefaultColumnValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ColumnDefinition) GetDefaultValueOk() (*AnyOfmicrosoftGraphDefaultColumnValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ColumnDefinition) SetDefaultValue(v AnyOfmicrosoftGraphDefaultColumnValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ColumnDefinition) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ColumnDefinition) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ColumnDefinition) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDescription

`func (o *ColumnDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ColumnDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ColumnDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ColumnDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ColumnDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ColumnDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *ColumnDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ColumnDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ColumnDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ColumnDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ColumnDefinition) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ColumnDefinition) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEnforceUniqueValues

`func (o *ColumnDefinition) GetEnforceUniqueValues() bool`

GetEnforceUniqueValues returns the EnforceUniqueValues field if non-nil, zero value otherwise.

### GetEnforceUniqueValuesOk

`func (o *ColumnDefinition) GetEnforceUniqueValuesOk() (*bool, bool)`

GetEnforceUniqueValuesOk returns a tuple with the EnforceUniqueValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUniqueValues

`func (o *ColumnDefinition) SetEnforceUniqueValues(v bool)`

SetEnforceUniqueValues sets EnforceUniqueValues field to given value.

### HasEnforceUniqueValues

`func (o *ColumnDefinition) HasEnforceUniqueValues() bool`

HasEnforceUniqueValues returns a boolean if a field has been set.

### SetEnforceUniqueValuesNil

`func (o *ColumnDefinition) SetEnforceUniqueValuesNil(b bool)`

 SetEnforceUniqueValuesNil sets the value for EnforceUniqueValues to be an explicit nil

### UnsetEnforceUniqueValues
`func (o *ColumnDefinition) UnsetEnforceUniqueValues()`

UnsetEnforceUniqueValues ensures that no value is present for EnforceUniqueValues, not even an explicit nil
### GetGeolocation

`func (o *ColumnDefinition) GetGeolocation() AnyOfobject`

GetGeolocation returns the Geolocation field if non-nil, zero value otherwise.

### GetGeolocationOk

`func (o *ColumnDefinition) GetGeolocationOk() (*AnyOfobject, bool)`

GetGeolocationOk returns a tuple with the Geolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocation

`func (o *ColumnDefinition) SetGeolocation(v AnyOfobject)`

SetGeolocation sets Geolocation field to given value.

### HasGeolocation

`func (o *ColumnDefinition) HasGeolocation() bool`

HasGeolocation returns a boolean if a field has been set.

### SetGeolocationNil

`func (o *ColumnDefinition) SetGeolocationNil(b bool)`

 SetGeolocationNil sets the value for Geolocation to be an explicit nil

### UnsetGeolocation
`func (o *ColumnDefinition) UnsetGeolocation()`

UnsetGeolocation ensures that no value is present for Geolocation, not even an explicit nil
### GetHidden

`func (o *ColumnDefinition) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ColumnDefinition) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ColumnDefinition) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ColumnDefinition) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *ColumnDefinition) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *ColumnDefinition) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
### GetHyperlinkOrPicture

`func (o *ColumnDefinition) GetHyperlinkOrPicture() AnyOfmicrosoftGraphHyperlinkOrPictureColumn`

GetHyperlinkOrPicture returns the HyperlinkOrPicture field if non-nil, zero value otherwise.

### GetHyperlinkOrPictureOk

`func (o *ColumnDefinition) GetHyperlinkOrPictureOk() (*AnyOfmicrosoftGraphHyperlinkOrPictureColumn, bool)`

GetHyperlinkOrPictureOk returns a tuple with the HyperlinkOrPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperlinkOrPicture

`func (o *ColumnDefinition) SetHyperlinkOrPicture(v AnyOfmicrosoftGraphHyperlinkOrPictureColumn)`

SetHyperlinkOrPicture sets HyperlinkOrPicture field to given value.

### HasHyperlinkOrPicture

`func (o *ColumnDefinition) HasHyperlinkOrPicture() bool`

HasHyperlinkOrPicture returns a boolean if a field has been set.

### SetHyperlinkOrPictureNil

`func (o *ColumnDefinition) SetHyperlinkOrPictureNil(b bool)`

 SetHyperlinkOrPictureNil sets the value for HyperlinkOrPicture to be an explicit nil

### UnsetHyperlinkOrPicture
`func (o *ColumnDefinition) UnsetHyperlinkOrPicture()`

UnsetHyperlinkOrPicture ensures that no value is present for HyperlinkOrPicture, not even an explicit nil
### GetIndexed

`func (o *ColumnDefinition) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *ColumnDefinition) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *ColumnDefinition) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *ColumnDefinition) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### SetIndexedNil

`func (o *ColumnDefinition) SetIndexedNil(b bool)`

 SetIndexedNil sets the value for Indexed to be an explicit nil

### UnsetIndexed
`func (o *ColumnDefinition) UnsetIndexed()`

UnsetIndexed ensures that no value is present for Indexed, not even an explicit nil
### GetIsDeletable

`func (o *ColumnDefinition) GetIsDeletable() bool`

GetIsDeletable returns the IsDeletable field if non-nil, zero value otherwise.

### GetIsDeletableOk

`func (o *ColumnDefinition) GetIsDeletableOk() (*bool, bool)`

GetIsDeletableOk returns a tuple with the IsDeletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletable

`func (o *ColumnDefinition) SetIsDeletable(v bool)`

SetIsDeletable sets IsDeletable field to given value.

### HasIsDeletable

`func (o *ColumnDefinition) HasIsDeletable() bool`

HasIsDeletable returns a boolean if a field has been set.

### SetIsDeletableNil

`func (o *ColumnDefinition) SetIsDeletableNil(b bool)`

 SetIsDeletableNil sets the value for IsDeletable to be an explicit nil

### UnsetIsDeletable
`func (o *ColumnDefinition) UnsetIsDeletable()`

UnsetIsDeletable ensures that no value is present for IsDeletable, not even an explicit nil
### GetIsReorderable

`func (o *ColumnDefinition) GetIsReorderable() bool`

GetIsReorderable returns the IsReorderable field if non-nil, zero value otherwise.

### GetIsReorderableOk

`func (o *ColumnDefinition) GetIsReorderableOk() (*bool, bool)`

GetIsReorderableOk returns a tuple with the IsReorderable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReorderable

`func (o *ColumnDefinition) SetIsReorderable(v bool)`

SetIsReorderable sets IsReorderable field to given value.

### HasIsReorderable

`func (o *ColumnDefinition) HasIsReorderable() bool`

HasIsReorderable returns a boolean if a field has been set.

### SetIsReorderableNil

`func (o *ColumnDefinition) SetIsReorderableNil(b bool)`

 SetIsReorderableNil sets the value for IsReorderable to be an explicit nil

### UnsetIsReorderable
`func (o *ColumnDefinition) UnsetIsReorderable()`

UnsetIsReorderable ensures that no value is present for IsReorderable, not even an explicit nil
### GetIsSealed

`func (o *ColumnDefinition) GetIsSealed() bool`

GetIsSealed returns the IsSealed field if non-nil, zero value otherwise.

### GetIsSealedOk

`func (o *ColumnDefinition) GetIsSealedOk() (*bool, bool)`

GetIsSealedOk returns a tuple with the IsSealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSealed

`func (o *ColumnDefinition) SetIsSealed(v bool)`

SetIsSealed sets IsSealed field to given value.

### HasIsSealed

`func (o *ColumnDefinition) HasIsSealed() bool`

HasIsSealed returns a boolean if a field has been set.

### SetIsSealedNil

`func (o *ColumnDefinition) SetIsSealedNil(b bool)`

 SetIsSealedNil sets the value for IsSealed to be an explicit nil

### UnsetIsSealed
`func (o *ColumnDefinition) UnsetIsSealed()`

UnsetIsSealed ensures that no value is present for IsSealed, not even an explicit nil
### GetLookup

`func (o *ColumnDefinition) GetLookup() AnyOfmicrosoftGraphLookupColumn`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *ColumnDefinition) GetLookupOk() (*AnyOfmicrosoftGraphLookupColumn, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *ColumnDefinition) SetLookup(v AnyOfmicrosoftGraphLookupColumn)`

SetLookup sets Lookup field to given value.

### HasLookup

`func (o *ColumnDefinition) HasLookup() bool`

HasLookup returns a boolean if a field has been set.

### SetLookupNil

`func (o *ColumnDefinition) SetLookupNil(b bool)`

 SetLookupNil sets the value for Lookup to be an explicit nil

### UnsetLookup
`func (o *ColumnDefinition) UnsetLookup()`

UnsetLookup ensures that no value is present for Lookup, not even an explicit nil
### GetName

`func (o *ColumnDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ColumnDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ColumnDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ColumnDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ColumnDefinition) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ColumnDefinition) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNumber

`func (o *ColumnDefinition) GetNumber() AnyOfmicrosoftGraphNumberColumn`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ColumnDefinition) GetNumberOk() (*AnyOfmicrosoftGraphNumberColumn, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ColumnDefinition) SetNumber(v AnyOfmicrosoftGraphNumberColumn)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ColumnDefinition) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *ColumnDefinition) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *ColumnDefinition) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetPersonOrGroup

`func (o *ColumnDefinition) GetPersonOrGroup() AnyOfmicrosoftGraphPersonOrGroupColumn`

GetPersonOrGroup returns the PersonOrGroup field if non-nil, zero value otherwise.

### GetPersonOrGroupOk

`func (o *ColumnDefinition) GetPersonOrGroupOk() (*AnyOfmicrosoftGraphPersonOrGroupColumn, bool)`

GetPersonOrGroupOk returns a tuple with the PersonOrGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonOrGroup

`func (o *ColumnDefinition) SetPersonOrGroup(v AnyOfmicrosoftGraphPersonOrGroupColumn)`

SetPersonOrGroup sets PersonOrGroup field to given value.

### HasPersonOrGroup

`func (o *ColumnDefinition) HasPersonOrGroup() bool`

HasPersonOrGroup returns a boolean if a field has been set.

### SetPersonOrGroupNil

`func (o *ColumnDefinition) SetPersonOrGroupNil(b bool)`

 SetPersonOrGroupNil sets the value for PersonOrGroup to be an explicit nil

### UnsetPersonOrGroup
`func (o *ColumnDefinition) UnsetPersonOrGroup()`

UnsetPersonOrGroup ensures that no value is present for PersonOrGroup, not even an explicit nil
### GetPropagateChanges

`func (o *ColumnDefinition) GetPropagateChanges() bool`

GetPropagateChanges returns the PropagateChanges field if non-nil, zero value otherwise.

### GetPropagateChangesOk

`func (o *ColumnDefinition) GetPropagateChangesOk() (*bool, bool)`

GetPropagateChangesOk returns a tuple with the PropagateChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagateChanges

`func (o *ColumnDefinition) SetPropagateChanges(v bool)`

SetPropagateChanges sets PropagateChanges field to given value.

### HasPropagateChanges

`func (o *ColumnDefinition) HasPropagateChanges() bool`

HasPropagateChanges returns a boolean if a field has been set.

### SetPropagateChangesNil

`func (o *ColumnDefinition) SetPropagateChangesNil(b bool)`

 SetPropagateChangesNil sets the value for PropagateChanges to be an explicit nil

### UnsetPropagateChanges
`func (o *ColumnDefinition) UnsetPropagateChanges()`

UnsetPropagateChanges ensures that no value is present for PropagateChanges, not even an explicit nil
### GetReadOnly

`func (o *ColumnDefinition) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ColumnDefinition) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ColumnDefinition) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ColumnDefinition) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *ColumnDefinition) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *ColumnDefinition) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetRequired

`func (o *ColumnDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ColumnDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ColumnDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ColumnDefinition) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ColumnDefinition) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ColumnDefinition) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetTerm

`func (o *ColumnDefinition) GetTerm() AnyOfmicrosoftGraphTermColumn`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *ColumnDefinition) GetTermOk() (*AnyOfmicrosoftGraphTermColumn, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *ColumnDefinition) SetTerm(v AnyOfmicrosoftGraphTermColumn)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *ColumnDefinition) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### SetTermNil

`func (o *ColumnDefinition) SetTermNil(b bool)`

 SetTermNil sets the value for Term to be an explicit nil

### UnsetTerm
`func (o *ColumnDefinition) UnsetTerm()`

UnsetTerm ensures that no value is present for Term, not even an explicit nil
### GetText

`func (o *ColumnDefinition) GetText() AnyOfmicrosoftGraphTextColumn`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ColumnDefinition) GetTextOk() (*AnyOfmicrosoftGraphTextColumn, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ColumnDefinition) SetText(v AnyOfmicrosoftGraphTextColumn)`

SetText sets Text field to given value.

### HasText

`func (o *ColumnDefinition) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ColumnDefinition) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ColumnDefinition) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetThumbnail

`func (o *ColumnDefinition) GetThumbnail() AnyOfobject`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *ColumnDefinition) GetThumbnailOk() (*AnyOfobject, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *ColumnDefinition) SetThumbnail(v AnyOfobject)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *ColumnDefinition) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### SetThumbnailNil

`func (o *ColumnDefinition) SetThumbnailNil(b bool)`

 SetThumbnailNil sets the value for Thumbnail to be an explicit nil

### UnsetThumbnail
`func (o *ColumnDefinition) UnsetThumbnail()`

UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
### GetType

`func (o *ColumnDefinition) GetType() AnyOfmicrosoftGraphColumnTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ColumnDefinition) GetTypeOk() (*AnyOfmicrosoftGraphColumnTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ColumnDefinition) SetType(v AnyOfmicrosoftGraphColumnTypes)`

SetType sets Type field to given value.

### HasType

`func (o *ColumnDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ColumnDefinition) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ColumnDefinition) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValidation

`func (o *ColumnDefinition) GetValidation() AnyOfmicrosoftGraphColumnValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *ColumnDefinition) GetValidationOk() (*AnyOfmicrosoftGraphColumnValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *ColumnDefinition) SetValidation(v AnyOfmicrosoftGraphColumnValidation)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *ColumnDefinition) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *ColumnDefinition) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *ColumnDefinition) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetSourceColumn

`func (o *ColumnDefinition) GetSourceColumn() AnyOfmicrosoftGraphColumnDefinition`

GetSourceColumn returns the SourceColumn field if non-nil, zero value otherwise.

### GetSourceColumnOk

`func (o *ColumnDefinition) GetSourceColumnOk() (*AnyOfmicrosoftGraphColumnDefinition, bool)`

GetSourceColumnOk returns a tuple with the SourceColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceColumn

`func (o *ColumnDefinition) SetSourceColumn(v AnyOfmicrosoftGraphColumnDefinition)`

SetSourceColumn sets SourceColumn field to given value.

### HasSourceColumn

`func (o *ColumnDefinition) HasSourceColumn() bool`

HasSourceColumn returns a boolean if a field has been set.

### SetSourceColumnNil

`func (o *ColumnDefinition) SetSourceColumnNil(b bool)`

 SetSourceColumnNil sets the value for SourceColumn to be an explicit nil

### UnsetSourceColumn
`func (o *ColumnDefinition) UnsetSourceColumn()`

UnsetSourceColumn ensures that no value is present for SourceColumn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


