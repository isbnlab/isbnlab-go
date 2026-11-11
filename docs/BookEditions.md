# BookEditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isbn13** | Pointer to **NullableString** |  | [optional] 
**OtherEditions** | Pointer to [**[]EditionRef**](EditionRef.md) |  | [optional] 
**PreviousEdition** | Pointer to [**NullableEditionRef**](EditionRef.md) |  | [optional] 
**NextEdition** | Pointer to [**NullableEditionRef**](EditionRef.md) |  | [optional] 

## Methods

### NewBookEditions

`func NewBookEditions() *BookEditions`

NewBookEditions instantiates a new BookEditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookEditionsWithDefaults

`func NewBookEditionsWithDefaults() *BookEditions`

NewBookEditionsWithDefaults instantiates a new BookEditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsbn13

`func (o *BookEditions) GetIsbn13() string`

GetIsbn13 returns the Isbn13 field if non-nil, zero value otherwise.

### GetIsbn13Ok

`func (o *BookEditions) GetIsbn13Ok() (*string, bool)`

GetIsbn13Ok returns a tuple with the Isbn13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn13

`func (o *BookEditions) SetIsbn13(v string)`

SetIsbn13 sets Isbn13 field to given value.

### HasIsbn13

`func (o *BookEditions) HasIsbn13() bool`

HasIsbn13 returns a boolean if a field has been set.

### SetIsbn13Nil

`func (o *BookEditions) SetIsbn13Nil(b bool)`

 SetIsbn13Nil sets the value for Isbn13 to be an explicit nil

### UnsetIsbn13
`func (o *BookEditions) UnsetIsbn13()`

UnsetIsbn13 ensures that no value is present for Isbn13, not even an explicit nil
### GetOtherEditions

`func (o *BookEditions) GetOtherEditions() []EditionRef`

GetOtherEditions returns the OtherEditions field if non-nil, zero value otherwise.

### GetOtherEditionsOk

`func (o *BookEditions) GetOtherEditionsOk() (*[]EditionRef, bool)`

GetOtherEditionsOk returns a tuple with the OtherEditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherEditions

`func (o *BookEditions) SetOtherEditions(v []EditionRef)`

SetOtherEditions sets OtherEditions field to given value.

### HasOtherEditions

`func (o *BookEditions) HasOtherEditions() bool`

HasOtherEditions returns a boolean if a field has been set.

### SetOtherEditionsNil

`func (o *BookEditions) SetOtherEditionsNil(b bool)`

 SetOtherEditionsNil sets the value for OtherEditions to be an explicit nil

### UnsetOtherEditions
`func (o *BookEditions) UnsetOtherEditions()`

UnsetOtherEditions ensures that no value is present for OtherEditions, not even an explicit nil
### GetPreviousEdition

`func (o *BookEditions) GetPreviousEdition() EditionRef`

GetPreviousEdition returns the PreviousEdition field if non-nil, zero value otherwise.

### GetPreviousEditionOk

`func (o *BookEditions) GetPreviousEditionOk() (*EditionRef, bool)`

GetPreviousEditionOk returns a tuple with the PreviousEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousEdition

`func (o *BookEditions) SetPreviousEdition(v EditionRef)`

SetPreviousEdition sets PreviousEdition field to given value.

### HasPreviousEdition

`func (o *BookEditions) HasPreviousEdition() bool`

HasPreviousEdition returns a boolean if a field has been set.

### SetPreviousEditionNil

`func (o *BookEditions) SetPreviousEditionNil(b bool)`

 SetPreviousEditionNil sets the value for PreviousEdition to be an explicit nil

### UnsetPreviousEdition
`func (o *BookEditions) UnsetPreviousEdition()`

UnsetPreviousEdition ensures that no value is present for PreviousEdition, not even an explicit nil
### GetNextEdition

`func (o *BookEditions) GetNextEdition() EditionRef`

GetNextEdition returns the NextEdition field if non-nil, zero value otherwise.

### GetNextEditionOk

`func (o *BookEditions) GetNextEditionOk() (*EditionRef, bool)`

GetNextEditionOk returns a tuple with the NextEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEdition

`func (o *BookEditions) SetNextEdition(v EditionRef)`

SetNextEdition sets NextEdition field to given value.

### HasNextEdition

`func (o *BookEditions) HasNextEdition() bool`

HasNextEdition returns a boolean if a field has been set.

### SetNextEditionNil

`func (o *BookEditions) SetNextEditionNil(b bool)`

 SetNextEditionNil sets the value for NextEdition to be an explicit nil

### UnsetNextEdition
`func (o *BookEditions) UnsetNextEdition()`

UnsetNextEdition ensures that no value is present for NextEdition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


