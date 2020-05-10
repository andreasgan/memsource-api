# UserDetailsDtoV2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**UserName** | **string** |  | [optional] [default to null]
**FirstName** | **string** |  | [optional] [default to null]
**LastName** | **string** |  | [optional] [default to null]
**Email** | **string** |  | [optional] [default to null]
**DateCreated** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**DateDeleted** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CreatedBy** | [***UserReference**](UserReference.md) |  | [optional] [default to null]
**Role** | **string** |  | [optional] [default to null]
**Timezone** | **string** |  | [optional] [default to null]
**Note** | **string** |  | [optional] [default to null]
**MayEditApprovedTerms** | **bool** | Default: false | [optional] [default to null]
**MayRejectJobs** | **bool** | Default: false | [optional] [default to null]
**EditorMachineTranslateEnabled** | **bool** | Default: true | [optional] [default to null]
**ReceiveNewsletter** | **bool** | Default: true | [optional] [default to null]
**MayEditTranslationMemory** | **bool** | Default: false | [optional] [default to null]
**SourceLangs** | **[]string** |  | [optional] [default to null]
**TargetLangs** | **[]string** |  | [optional] [default to null]
**WorkflowSteps** | [**[]WorkflowStepReference**](WorkflowStepReference.md) |  | [optional] [default to null]
**Clients** | [**[]ClientReference**](ClientReference.md) |  | [optional] [default to null]
**Domains** | [**[]DomainReference**](DomainReference.md) |  | [optional] [default to null]
**SubDomains** | [**[]SubDomainReference**](SubDomainReference.md) |  | [optional] [default to null]
**ProjectBusinessUnits** | [**[]BusinessUnitReference**](BusinessUnitReference.md) |  | [optional] [default to null]
**Organization** | [***IdReference**](IdReference.md) |  | [optional] [default to null]
**PriceList** | [***PriceListReference**](PriceListReference.md) |  | [optional] [default to null]
**NetRateScheme** | [***DiscountSchemeReference**](DiscountSchemeReference.md) |  | [optional] [default to null]
**AutomationWidgets** | [**[]JobWidgetReference**](JobWidgetReference.md) |  | [optional] [default to null]
**Active** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


