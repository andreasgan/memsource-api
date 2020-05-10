# CreateProjectDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**SourceLang** | **string** |  | [default to null]
**TargetLangs** | **[]string** |  | [default to null]
**Client** | [***IdReference**](IdReference.md) | Client referenced by id | [optional] [default to null]
**BusinessUnit** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**Domain** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**SubDomain** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**UseDefaultProjectSettings** | **bool** | Default: false | [optional] [default to null]
**PurchaseOrder** | **string** |  | [optional] [default to null]
**WorkflowSteps** | [**[]IdReference**](IdReference.md) |  | [optional] [default to null]
**DateDue** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Note** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


