# ProjectTemplateDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**TemplateName** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**SourceLang** | **string** |  | [optional] [default to null]
**TargetLangs** | **[]string** |  | [optional] [default to null]
**Note** | **string** |  | [optional] [default to null]
**Owner** | [***UserReference**](UserReference.md) |  | [optional] [default to null]
**Client** | [***ClientReference**](ClientReference.md) |  | [optional] [default to null]
**Domain** | [***DomainReference**](DomainReference.md) |  | [optional] [default to null]
**SubDomain** | [***SubDomainReference**](SubDomainReference.md) |  | [optional] [default to null]
**CreatedBy** | [***UserReference**](UserReference.md) |  | [optional] [default to null]
**DateCreated** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**WorkflowSteps** | [**[]WorkflowStepDto**](WorkflowStepDto.md) |  | [optional] [default to null]
**WorkflowSettings** | [**[]WorkflowStepSettingsDto**](WorkflowStepSettingsDto.md) |  | [optional] [default to null]
**BusinessUnit** | [***BusinessUnitReference**](BusinessUnitReference.md) |  | [optional] [default to null]
**NotifyProviders** | [***ProjectTemplateNotifyProviderDto**](ProjectTemplateNotifyProviderDto.md) |  | [optional] [default to null]
**AssignedTo** | [**[]AssignmentPerTargetLangDto**](AssignmentPerTargetLangDto.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


