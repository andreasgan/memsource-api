# CreateAnalyseListAsyncDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]UidReference**](UidReference.md) |  | [default to null]
**Type_** | **string** | default: PreAnalyse | [optional] [default to null]
**IncludeFuzzyRepetitions** | **bool** | Default: true | [optional] [default to null]
**IncludeConfirmedSegments** | **bool** | Default: true | [optional] [default to null]
**IncludeNumbers** | **bool** | Default: true | [optional] [default to null]
**IncludeLockedSegments** | **bool** | Default: true | [optional] [default to null]
**CountSourceUnits** | **bool** | Default: true | [optional] [default to null]
**IncludeTransMemory** | **bool** | Default: true | [optional] [default to null]
**IncludeNonTranslatables** | **bool** | Default: false. Works only for type&#x3D;PreAnalyse. | [optional] [default to null]
**IncludeMachineTranslationMatches** | **bool** | Default: false. Works only for type&#x3D;PreAnalyse. | [optional] [default to null]
**TransMemoryPostEditing** | **bool** | Default: false. Works only for type&#x3D;PostAnalyse. | [optional] [default to null]
**NonTranslatablePostEditing** | **bool** | Default: false. Works only for type&#x3D;PostAnalyse. | [optional] [default to null]
**MachineTranslatePostEditing** | **bool** | Default: false. Works only for type&#x3D;PostAnalyse. | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**NetRateScheme** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**CompareWorkflowLevel** | **int32** | Required for type&#x3D;Compare | [optional] [default to null]
**UseProjectAnalysisSettings** | **bool** | Default: false. Use default project settings. Will be overwritten with setting sent         in the API call. | [optional] [default to null]
**CallbackUrl** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


