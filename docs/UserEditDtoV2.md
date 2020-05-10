# UserEditDtoV2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** |  | [default to null]
**FirstName** | **string** |  | [default to null]
**LastName** | **string** |  | [default to null]
**Email** | **string** |  | [default to null]
**Role** | **string** |  | [default to null]
**Timezone** | **string** |  | [default to null]
**Note** | **string** |  | [optional] [default to null]
**MayEditApprovedTerms** | **bool** | In previous version as terminologist. Default: false | [optional] [default to null]
**MayRejectJobs** | **bool** | Default: false | [optional] [default to null]
**EditorMachineTranslateEnabled** | **bool** | Applies only to Linguist or Guest. Default: true | [optional] [default to null]
**ReceiveNewsletter** | **bool** | Default: true | [optional] [default to null]
**MayEditTranslationMemory** | **bool** | Default: false | [optional] [default to null]
**SourceLangs** | **[]string** |  | [optional] [default to null]
**TargetLangs** | **[]string** |  | [optional] [default to null]
**Active** | **bool** | Default: true | [optional] [default to null]
**WorkflowSteps** | [**[]IdReference**](IdReference.md) |  | [optional] [default to null]
**Clients** | [**[]IdReference**](IdReference.md) |  | [optional] [default to null]
**Domains** | [**[]IdReference**](IdReference.md) |  | [optional] [default to null]
**SubDomains** | [**[]IdReference**](IdReference.md) |  | [optional] [default to null]
**ProjectBusinessUnits** | [**[]IdReference**](IdReference.md) |  | [optional] [default to null]
**AutomationWidgets** | [**[]IdReference**](IdReference.md) | Applies only to Submitters, where it must contain at least one assigned automation widget | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


