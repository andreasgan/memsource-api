# AdminProjectManager

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | [optional] [default to null]
**InternalId** | **int32** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**DateCreated** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Domain** | [***DomainReference**](DomainReference.md) |  | [optional] [default to null]
**SubDomain** | [***SubDomainReference**](SubDomainReference.md) |  | [optional] [default to null]
**Owner** | [***UserReference**](UserReference.md) |  | [optional] [default to null]
**SourceLang** | **string** |  | [optional] [default to null]
**TargetLangs** | **[]string** |  | [optional] [default to null]
**References** | [**[]ReferenceFileReference**](ReferenceFileReference.md) |  | [optional] [default to null]
**UserRole** | **string** | Response differs based on user&#39;s role | [optional] [default to null]
**Shared** | **bool** | Default: false | [optional] [default to null]
**Progress** | [***ProgressDto**](ProgressDto.md) |  | [optional] [default to null]
**Client** | [***ClientReference**](ClientReference.md) |  | [optional] [default to null]
**CostCenter** | [***CostCenterReference**](CostCenterReference.md) |  | [optional] [default to null]
**BusinessUnit** | [***BusinessUnitReference**](BusinessUnitReference.md) |  | [optional] [default to null]
**DateDue** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**PurchaseOrder** | **string** |  | [optional] [default to null]
**IsPublishedOnJobBoard** | **bool** | Default: false | [optional] [default to null]
**Note** | **string** |  | [optional] [default to null]
**CreatedBy** | [***UserReference**](UserReference.md) |  | [optional] [default to null]
**QualityAssuranceSettings** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**WorkflowSteps** | [**[]ProjectWorkflowStepDto**](ProjectWorkflowStepDto.md) |  | [optional] [default to null]
**AnalyseSettings** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**AccessSettings** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**FinancialSettings** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**MtSettingsPerLanguageList** | [**[]MtSettingsPerLanguageReference**](MTSettingsPerLanguageReference.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


