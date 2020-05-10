# SearchRequestDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | [default to null]
**SourceLang** | **string** |  | [default to null]
**TargetLangs** | **[]string** |  | [optional] [default to null]
**PreviousSegment** | **string** |  | [optional] [default to null]
**NextSegment** | **string** |  | [optional] [default to null]
**TagMetadata** | [**[]TagMetadataDto**](TagMetadataDto.md) |  | [optional] [default to null]
**TrimQuery** | **bool** | Remove leading and trailing whitespace from query. Default: true | [optional] [default to null]
**PhraseQuery** | **bool** | Return both wildcard and exact search results. Default: true | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


