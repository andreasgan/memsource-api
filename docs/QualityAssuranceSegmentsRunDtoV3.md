# QualityAssuranceSegmentsRunDtoV3

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobsAndSegments** | [**[]JobPartSegmentsDtoV3**](JobPartSegmentsDtoV3.md) |  | [default to null]
**WarningTypes** | **[]string** | When empty only fast checks run | [optional] [default to null]
**MaxQaWarningsCount** | **int32** | Maximum number of QA warnings in result, default: 100. For efficiency reasons QA warnings are processed with minimum segments chunk size 10, therefore slightly more warnings are returned. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


