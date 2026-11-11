# BookCover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isbn13** | Pointer to **NullableString** |  | [optional] 
**Covers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBookCover

`func NewBookCover() *BookCover`

NewBookCover instantiates a new BookCover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookCoverWithDefaults

`func NewBookCoverWithDefaults() *BookCover`

NewBookCoverWithDefaults instantiates a new BookCover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsbn13

`func (o *BookCover) GetIsbn13() string`

GetIsbn13 returns the Isbn13 field if non-nil, zero value otherwise.

### GetIsbn13Ok

`func (o *BookCover) GetIsbn13Ok() (*string, bool)`

GetIsbn13Ok returns a tuple with the Isbn13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn13

`func (o *BookCover) SetIsbn13(v string)`

SetIsbn13 sets Isbn13 field to given value.

### HasIsbn13

`func (o *BookCover) HasIsbn13() bool`

HasIsbn13 returns a boolean if a field has been set.

### SetIsbn13Nil

`func (o *BookCover) SetIsbn13Nil(b bool)`

 SetIsbn13Nil sets the value for Isbn13 to be an explicit nil

### UnsetIsbn13
`func (o *BookCover) UnsetIsbn13()`

UnsetIsbn13 ensures that no value is present for Isbn13, not even an explicit nil
### GetCovers

`func (o *BookCover) GetCovers() []string`

GetCovers returns the Covers field if non-nil, zero value otherwise.

### GetCoversOk

`func (o *BookCover) GetCoversOk() (*[]string, bool)`

GetCoversOk returns a tuple with the Covers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovers

`func (o *BookCover) SetCovers(v []string)`

SetCovers sets Covers field to given value.

### HasCovers

`func (o *BookCover) HasCovers() bool`

HasCovers returns a boolean if a field has been set.

### SetCoversNil

`func (o *BookCover) SetCoversNil(b bool)`

 SetCoversNil sets the value for Covers to be an explicit nil

### UnsetCovers
`func (o *BookCover) UnsetCovers()`

UnsetCovers ensures that no value is present for Covers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


