# ProjectTemplateEditDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [optional] [default to null]
**TemplateName** | **string** |  | [default to null]
**SourceLang** | **string** |  | [optional] [default to null]
**TargetLangs** | **[]string** |  | [optional] [default to null]
**NotifyProvider** | [***ProjectTemplateNotifyProviderDto**](ProjectTemplateNotifyProviderDto.md) | use to notify assigned providers,         notificationIntervalInMinutes 0 or empty value means immediate notification to all providers | [optional] [default to null]
**WorkFlowSettings** | [**[]WorkflowStepSettingsEditDto**](WorkflowStepSettingsEditDto.md) |  | [optional] [default to null]
**Client** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**CostCenter** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**BusinessUnit** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**Domain** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**SubDomain** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**Note** | **string** |  | [optional] [default to null]
**AssignedTo** | [**[]ProjectTemplateWorkflowSettingsAssignedToDto**](ProjectTemplateWorkflowSettingsAssignedToDto.md) | only use for projects without workflows; otherwise specify in the workflowSettings object | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


