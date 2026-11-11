# Series

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**Identifier** | Pointer to **NullableString** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Number** | Pointer to **NullableInt64** |  | [optional] 
**Kind** | Pointer to **NullableString** |  | [optional] 
**Parent** | Pointer to [**NullableSeriesParent**](SeriesParent.md) |  | [optional] 
**OtherSeries** | Pointer to [**[]OtherSeries**](OtherSeries.md) |  | [optional] 
**PreviousBook** | Pointer to [**NullableSeriesBook**](SeriesBook.md) |  | [optional] 
**NextBook** | Pointer to [**NullableSeriesBook**](SeriesBook.md) |  | [optional] 

## Methods

### NewSeries

`func NewSeries() *Series`

NewSeries instantiates a new Series object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesWithDefaults

`func NewSeriesWithDefaults() *Series`

NewSeriesWithDefaults instantiates a new Series object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Series) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Series) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Series) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Series) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Series) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Series) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentifier

`func (o *Series) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Series) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Series) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Series) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *Series) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *Series) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetSlug

`func (o *Series) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Series) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Series) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Series) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *Series) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *Series) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetName

`func (o *Series) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Series) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Series) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Series) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Series) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Series) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNumber

`func (o *Series) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Series) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Series) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Series) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *Series) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *Series) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetKind

`func (o *Series) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Series) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Series) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Series) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *Series) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *Series) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetParent

`func (o *Series) GetParent() SeriesParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Series) GetParentOk() (*SeriesParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Series) SetParent(v SeriesParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Series) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *Series) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *Series) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetOtherSeries

`func (o *Series) GetOtherSeries() []OtherSeries`

GetOtherSeries returns the OtherSeries field if non-nil, zero value otherwise.

### GetOtherSeriesOk

`func (o *Series) GetOtherSeriesOk() (*[]OtherSeries, bool)`

GetOtherSeriesOk returns a tuple with the OtherSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSeries

`func (o *Series) SetOtherSeries(v []OtherSeries)`

SetOtherSeries sets OtherSeries field to given value.

### HasOtherSeries

`func (o *Series) HasOtherSeries() bool`

HasOtherSeries returns a boolean if a field has been set.

### SetOtherSeriesNil

`func (o *Series) SetOtherSeriesNil(b bool)`

 SetOtherSeriesNil sets the value for OtherSeries to be an explicit nil

### UnsetOtherSeries
`func (o *Series) UnsetOtherSeries()`

UnsetOtherSeries ensures that no value is present for OtherSeries, not even an explicit nil
### GetPreviousBook

`func (o *Series) GetPreviousBook() SeriesBook`

GetPreviousBook returns the PreviousBook field if non-nil, zero value otherwise.

### GetPreviousBookOk

`func (o *Series) GetPreviousBookOk() (*SeriesBook, bool)`

GetPreviousBookOk returns a tuple with the PreviousBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBook

`func (o *Series) SetPreviousBook(v SeriesBook)`

SetPreviousBook sets PreviousBook field to given value.

### HasPreviousBook

`func (o *Series) HasPreviousBook() bool`

HasPreviousBook returns a boolean if a field has been set.

### SetPreviousBookNil

`func (o *Series) SetPreviousBookNil(b bool)`

 SetPreviousBookNil sets the value for PreviousBook to be an explicit nil

### UnsetPreviousBook
`func (o *Series) UnsetPreviousBook()`

UnsetPreviousBook ensures that no value is present for PreviousBook, not even an explicit nil
### GetNextBook

`func (o *Series) GetNextBook() SeriesBook`

GetNextBook returns the NextBook field if non-nil, zero value otherwise.

### GetNextBookOk

`func (o *Series) GetNextBookOk() (*SeriesBook, bool)`

GetNextBookOk returns a tuple with the NextBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBook

`func (o *Series) SetNextBook(v SeriesBook)`

SetNextBook sets NextBook field to given value.

### HasNextBook

`func (o *Series) HasNextBook() bool`

HasNextBook returns a boolean if a field has been set.

### SetNextBookNil

`func (o *Series) SetNextBookNil(b bool)`

 SetNextBookNil sets the value for NextBook to be an explicit nil

### UnsetNextBook
`func (o *Series) UnsetNextBook()`

UnsetNextBook ensures that no value is present for NextBook, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


