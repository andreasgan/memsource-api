# SetProjectTransMemoriesDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadTransMemories** | [**[]IdReference**](IdReference.md) |  | [optional] [default to null]
**WriteTransMemories** | [**[]IdReference**](IdReference.md) | Write translation memory must be included in the read translation memories, too; max 2 write TMs allowed | [optional] [default to null]
**Penalties** | **[]float64** | A list of penalties for each read translation memory | [optional] [default to null]
**TargetLang** | **string** | Set translation memories only for the specific project target language | [optional] [default to null]
**WorkflowStep** | [***IdReference**](IdReference.md) | set translation memories only for the specific workflow step | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


