# \ProjectApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTargetLanguageToProject**](ProjectApi.md#AddTargetLanguageToProject) | **Post** /api2/v1/projects/{projectUid}/targetLangs | Add target languages
[**AddWorkflowSteps**](ProjectApi.md#AddWorkflowSteps) | **Post** /api2/v1/projects/{projectUid}/workflowSteps | Add workflow steps
[**AssignLinguistsFromTemplate**](ProjectApi.md#AssignLinguistsFromTemplate) | **Post** /api2/v1/projects/{projectUid}/applyTemplate/{templateId}/assignProviders | Assigns providers from template
[**AssignLinguistsFromTemplateToJobParts**](ProjectApi.md#AssignLinguistsFromTemplateToJobParts) | **Post** /api2/v1/projects/{projectUid}/applyTemplate/{templateId}/assignProviders/forJobParts | Assigns providers from template (specific jobs)
[**AssignVendorToProject**](ProjectApi.md#AssignVendorToProject) | **Post** /api2/v1/projects/{projectUid}/assignVendor | Assign vendor
[**AssignableTemplates**](ProjectApi.md#AssignableTemplates) | **Get** /api2/v1/projects/{projectUid}/assignableTemplates | List assignable templates
[**CloneProject**](ProjectApi.md#CloneProject) | **Post** /api2/v1/projects/{projectUid}/clone | Clone project
[**CreateNoteRef**](ProjectApi.md#CreateNoteRef) | **Post** /api2/v1/projects/{projectUid}/references | Create project reference file
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /api2/v1/projects | Create project
[**CreateProjectFromTemplate**](ProjectApi.md#CreateProjectFromTemplate) | **Post** /api2/v1/projects/applyTemplate/{templateId} | Create project from template
[**CreateProjectFromTemplateV2**](ProjectApi.md#CreateProjectFromTemplateV2) | **Post** /api2/v2/projects/applyTemplate/{templateId} | Create project from template
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /api2/v1/projects/{projectUid} | Delete project
[**DownloadReference**](ProjectApi.md#DownloadReference) | **Get** /api2/v1/projects/{projectUid}/references/{referenceFileId} | Get project reference
[**EditProject**](ProjectApi.md#EditProject) | **Put** /api2/v1/projects/{projectUid} | Edit project
[**EditProjectAccessSettings**](ProjectApi.md#EditProjectAccessSettings) | **Put** /api2/v1/projects/{projectUid}/accessSettings | Edit access and security settings
[**EditProjectQualityAssuranceSettings**](ProjectApi.md#EditProjectQualityAssuranceSettings) | **Put** /api2/v1/projects/{projectUid}/qaSettings | Edit quality assurance settings
[**EnabledQualityChecks**](ProjectApi.md#EnabledQualityChecks) | **Get** /api2/v1/projects/{projectUid}/qaSettingsChecks | Get QA checks
[**GetAnalyseSettingsForProject**](ProjectApi.md#GetAnalyseSettingsForProject) | **Get** /api2/v1/projects/{projectUid}/analyseSettings | Get analyse settings
[**GetFinancialSettings**](ProjectApi.md#GetFinancialSettings) | **Get** /api2/v1/projects/{projectUid}/financialSettings | Get financial settings
[**GetImportSettings2**](ProjectApi.md#GetImportSettings2) | **Get** /api2/v1/projects/{projectUid}/importSettings | Get projects&#39;s default import settings
[**GetMtSettingsForProject**](ProjectApi.md#GetMtSettingsForProject) | **Get** /api2/v1/projects/{projectUid}/mtSettings | Get machine translate settings
[**GetPreTranslateSettingsForProject**](ProjectApi.md#GetPreTranslateSettingsForProject) | **Get** /api2/v1/projects/{projectUid}/preTranslateSettings | Get Pre-translate settings
[**GetProject**](ProjectApi.md#GetProject) | **Get** /api2/v1/projects/{projectUid} | Get project
[**GetProjectAccessSettings**](ProjectApi.md#GetProjectAccessSettings) | **Get** /api2/v1/projects/{projectUid}/accessSettings | Get access and security settings
[**GetProjectAssignments**](ProjectApi.md#GetProjectAssignments) | **Get** /api2/v1/projects/{projectUid}/providers | List project providers
[**GetProjectQASettings**](ProjectApi.md#GetProjectQASettings) | **Get** /api2/v1/projects/{projectUid}/qaSettings | Get quality assurance settings
[**GetProjectQASettingsV2**](ProjectApi.md#GetProjectQASettingsV2) | **Get** /api2/v2/projects/{projectUid}/qaSettings | Get quality assurance settings
[**GetProjectSettings**](ProjectApi.md#GetProjectSettings) | **Get** /api2/v1/projects/{projectUid}/lqaSettings | Get LQA settings
[**GetProjectTermBases**](ProjectApi.md#GetProjectTermBases) | **Get** /api2/v1/projects/{projectUid}/termBases | Get term bases
[**GetProjectTransMemories**](ProjectApi.md#GetProjectTransMemories) | **Get** /api2/v1/projects/{projectUid}/transMemories | Get translation memories
[**GetProjectWorkflowSteps**](ProjectApi.md#GetProjectWorkflowSteps) | **Get** /api2/v1/projects/{projectUid}/workflowSteps | Get workflow steps
[**GetQuotesForProject**](ProjectApi.md#GetQuotesForProject) | **Get** /api2/v1/projects/{projectUid}/quotes | List quotes
[**ListAssignedProjects**](ProjectApi.md#ListAssignedProjects) | **Get** /api2/v1/users/{userId}/projects | List assigned projects
[**ListByProject**](ProjectApi.md#ListByProject) | **Get** /api2/v1/projects/{projectUid}/analyses | List analyses by project
[**ListByProjectV2**](ProjectApi.md#ListByProjectV2) | **Get** /api2/v2/projects/{projectUid}/analyses | List analyses by project
[**ListProjects**](ProjectApi.md#ListProjects) | **Get** /api2/v1/projects | List projects
[**ListProviders**](ProjectApi.md#ListProviders) | **Post** /api2/v1/projects/{projectUid}/providers/suggest | Get suggested providers
[**SearchSegment**](ProjectApi.md#SearchSegment) | **Post** /api2/v1/projects/{projectUid}/transMemories/searchSegmentInProject | Search translation memory for segment in the project
[**SetFinancialSettings**](ProjectApi.md#SetFinancialSettings) | **Put** /api2/v1/projects/{projectUid}/financialSettings | Edit financial settings
[**SetMtSettingsForProject**](ProjectApi.md#SetMtSettingsForProject) | **Put** /api2/v1/projects/{projectUid}/mtSettings | Edit machine translate settings
[**SetMtSettingsPerLanguageForProject**](ProjectApi.md#SetMtSettingsPerLanguageForProject) | **Put** /api2/v1/projects/{projectUid}/mtSettingsPerLanguage | Edit machine translate settings per language
[**SetProjectQASettingsV2**](ProjectApi.md#SetProjectQASettingsV2) | **Put** /api2/v2/projects/{projectUid}/qaSettings | Edit quality assurance settings
[**SetProjectStatus**](ProjectApi.md#SetProjectStatus) | **Post** /api2/v1/projects/{projectUid}/setStatus | Edit project status
[**SetProjectTermBases**](ProjectApi.md#SetProjectTermBases) | **Put** /api2/v1/projects/{projectUid}/termBases | Edit term bases
[**SetProjectTransMemories**](ProjectApi.md#SetProjectTransMemories) | **Put** /api2/v1/projects/{projectUid}/transMemories | Edit translation memories
[**SetProjectTransMemoriesV2**](ProjectApi.md#SetProjectTransMemoriesV2) | **Put** /api2/v2/projects/{projectUid}/transMemories | Edit translation memories


# **AddTargetLanguageToProject**
> AddTargetLanguageToProject(ctx, projectUid, optional)
Add target languages

Add target languages to project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiAddTargetLanguageToProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiAddTargetLanguageToProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AddTargetLangDto**](AddTargetLangDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddWorkflowSteps**
> AddWorkflowSteps(ctx, projectUid, optional)
Add workflow steps



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiAddWorkflowStepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiAddWorkflowStepsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AddWorkflowStepsDto**](AddWorkflowStepsDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignLinguistsFromTemplate**
> JobPartsDto AssignLinguistsFromTemplate(ctx, templateId, projectUid)
Assigns providers from template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**|  | 
  **projectUid** | **string**|  | 

### Return type

[**JobPartsDto**](JobPartsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignLinguistsFromTemplateToJobParts**
> JobPartsDto AssignLinguistsFromTemplateToJobParts(ctx, templateId, projectUid, optional)
Assigns providers from template (specific jobs)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiAssignLinguistsFromTemplateToJobPartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiAssignLinguistsFromTemplateToJobPartsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of JobPartReferences**](JobPartReferences.md)|  | 

### Return type

[**JobPartsDto**](JobPartsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignVendorToProject**
> AssignVendorToProject(ctx, projectUid, optional)
Assign vendor

 To unassign Vendor from Project, use empty body: ``` {} ```     

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiAssignVendorToProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiAssignVendorToProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AssignVendorDto**](AssignVendorDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignableTemplates**
> AssignableTemplatesDto AssignableTemplates(ctx, projectUid)
List assignable templates



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**AssignableTemplatesDto**](AssignableTemplatesDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloneProject**
> AbstractProjectDto CloneProject(ctx, projectUid, optional)
Clone project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiCloneProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCloneProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CloneProjectDto**](CloneProjectDto.md)|  | 

### Return type

[**AbstractProjectDto**](AbstractProjectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNoteRef**
> ReferenceFileReference CreateNoteRef(ctx, projectUid, optional)
Create project reference file

Accepts application/octet-stream or application/json.<br>                                                                     <b>application/json</b> - Note will be converted to .txt.<br>                                                                     <b>application/octet-stream</b> - Content-Disposition header is required

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiCreateNoteRefOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCreateNoteRefOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateReferenceFileNoteDto**](CreateReferenceFileNoteDto.md)|  | 
 **contentDisposition** | **optional.String**| &lt;b&gt;Required&lt;/b&gt; for application/octet-stream.&lt;br&gt; Example: filename*&#x3D;UTF-8&#39;&#39;YourFileName.txt | 

### Return type

[**ReferenceFileReference**](ReferenceFileReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> AbstractProjectDto CreateProject(ctx, optional)
Create project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectApiCreateProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCreateProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateProjectDto**](CreateProjectDto.md)|  | 

### Return type

[**AbstractProjectDto**](AbstractProjectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProjectFromTemplate**
> AbstractProjectDto CreateProjectFromTemplate(ctx, templateId, optional)
Create project from template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**|  | 
 **optional** | ***ProjectApiCreateProjectFromTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCreateProjectFromTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateProjectFromTemplateDto**](CreateProjectFromTemplateDto.md)|  | 

### Return type

[**AbstractProjectDto**](AbstractProjectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProjectFromTemplateV2**
> AbstractProjectDtoV2 CreateProjectFromTemplateV2(ctx, templateId, optional)
Create project from template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**|  | 
 **optional** | ***ProjectApiCreateProjectFromTemplateV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCreateProjectFromTemplateV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateProjectFromTemplateV2Dto**](CreateProjectFromTemplateV2Dto.md)|  | 

### Return type

[**AbstractProjectDtoV2**](AbstractProjectDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject(ctx, projectUid, optional)
Delete project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiDeleteProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiDeleteProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadReference**
> DownloadReference(ctx, projectUid, referenceFileId)
Get project reference



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **referenceFileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditProject**
> AbstractProjectDto EditProject(ctx, projectUid, optional)
Edit project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiEditProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiEditProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EditProjectDto**](EditProjectDto.md)|  | 

### Return type

[**AbstractProjectDto**](AbstractProjectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditProjectAccessSettings**
> ProjectSecuritySettingsDto EditProjectAccessSettings(ctx, projectUid, optional)
Edit access and security settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiEditProjectAccessSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiEditProjectAccessSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EditProjectSecuritySettingsDto**](EditProjectSecuritySettingsDto.md)|  | 

### Return type

[**ProjectSecuritySettingsDto**](ProjectSecuritySettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditProjectQualityAssuranceSettings**
> QaSettingsDto EditProjectQualityAssuranceSettings(ctx, projectUid, optional)
Edit quality assurance settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiEditProjectQualityAssuranceSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiEditProjectQualityAssuranceSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EditQaSettingsDto**](EditQaSettingsDto.md)|  | 

### Return type

[**QaSettingsDto**](QASettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnabledQualityChecks**
> EnabledQualityChecksDto EnabledQualityChecks(ctx, projectUid)
Get QA checks

Returns enabled quality assurance settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**EnabledQualityChecksDto**](EnabledQualityChecksDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnalyseSettingsForProject**
> AnalyseSettingsDto GetAnalyseSettingsForProject(ctx, projectUid)
Get analyse settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**AnalyseSettingsDto**](AnalyseSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFinancialSettings**
> FinancialSettingsDto GetFinancialSettings(ctx, projectUid)
Get financial settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**FinancialSettingsDto**](FinancialSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportSettings2**
> FileImportSettingsDto GetImportSettings2(ctx, projectUid)
Get projects's default import settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**FileImportSettingsDto**](FileImportSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMtSettingsForProject**
> MtSettingsPerLanguageListDto GetMtSettingsForProject(ctx, projectUid)
Get machine translate settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**MtSettingsPerLanguageListDto**](MTSettingsPerLanguageListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPreTranslateSettingsForProject**
> PreTranslateSettingsDto GetPreTranslateSettingsForProject(ctx, projectUid)
Get Pre-translate settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**PreTranslateSettingsDto**](PreTranslateSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> AbstractProjectDto GetProject(ctx, projectUid)
Get project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**AbstractProjectDto**](AbstractProjectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectAccessSettings**
> ProjectSecuritySettingsDto GetProjectAccessSettings(ctx, projectUid)
Get access and security settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**ProjectSecuritySettingsDto**](ProjectSecuritySettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectAssignments**
> PageDtoProviderReference GetProjectAssignments(ctx, projectUid, optional)
List project providers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiGetProjectAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetProjectAssignmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerName** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoProviderReference**](PageDtoProviderReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectQASettings**
> QaSettingsDto GetProjectQASettings(ctx, projectUid)
Get quality assurance settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**QaSettingsDto**](QASettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectQASettingsV2**
> QaSettingsDtoV2 GetProjectQASettingsV2(ctx, projectUid)
Get quality assurance settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**QaSettingsDtoV2**](QASettingsDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectSettings**
> LqaSettingsDto GetProjectSettings(ctx, projectUid, optional)
Get LQA settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiGetProjectSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetProjectSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowLevel** | **optional.Int32**|  | [default to 1]

### Return type

[**LqaSettingsDto**](LqaSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectTermBases**
> ProjectTermBaseListDto GetProjectTermBases(ctx, projectUid)
Get term bases



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**ProjectTermBaseListDto**](ProjectTermBaseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectTransMemories**
> ProjectTransMemoryListDto GetProjectTransMemories(ctx, projectUid)
Get translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**ProjectTransMemoryListDto**](ProjectTransMemoryListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectWorkflowSteps**
> ProjectWorkflowStepListDto GetProjectWorkflowSteps(ctx, projectUid)
Get workflow steps



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**ProjectWorkflowStepListDto**](ProjectWorkflowStepListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotesForProject**
> PageDtoQuoteDto GetQuotesForProject(ctx, projectUid, optional)
List quotes



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiGetQuotesForProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiGetQuotesForProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoQuoteDto**](PageDtoQuoteDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAssignedProjects**
> PageDtoProjectReference ListAssignedProjects(ctx, userId, optional)
List assigned projects



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***ProjectApiListAssignedProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiListAssignedProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**optional.Interface of []string**](string.md)|  | 
 **targetLang** | [**optional.Interface of []string**](string.md)|  | 
 **workflowStepId** | **optional.Int64**|  | 
 **dueInHours** | **optional.Int32**| -1 for jobs that are overdue | 
 **filename** | **optional.String**|  | 
 **projectName** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]

### Return type

[**PageDtoProjectReference**](PageDtoProjectReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByProject**
> PageDtoAnalyseDto ListByProject(ctx, projectUid, optional)
List analyses by project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiListByProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiListByProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoAnalyseDto**](PageDtoAnalyseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByProjectV2**
> PageDtoAnalyseV2Dto ListByProjectV2(ctx, projectUid, optional)
List analyses by project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiListByProjectV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiListByProjectV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoAnalyseV2Dto**](PageDtoAnalyseV2Dto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProjects**
> PageDtoAbstractProjectDto ListProjects(ctx, optional)
List projects



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectApiListProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiListProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **clientId** | **optional.Int64**|  | 
 **clientName** | **optional.String**|  | 
 **businessUnitName** | **optional.String**|  | 
 **statuses** | [**optional.Interface of []string**](string.md)|  | 
 **targetLangs** | [**optional.Interface of []string**](string.md)|  | 
 **domainName** | **optional.String**|  | 
 **subDomainName** | **optional.String**|  | 
 **costCenterId** | **optional.Int64**|  | 
 **costCenterName** | **optional.String**|  | 
 **dueInHours** | **optional.Int32**| -1 for projects that are overdue | 
 **createdInLastHours** | **optional.Int32**|  | 
 **sourceLangs** | [**optional.Interface of []string**](string.md)|  | 
 **ownerId** | **optional.Int64**|  | 
 **jobStatuses** | [**optional.Interface of []string**](string.md)| Allowed for linguists only | 
 **jobStatusGroup** | **optional.String**| Allowed for linguists only | 
 **buyerId** | **optional.Int64**|  | 
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoAbstractProjectDto**](PageDtoAbstractProjectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProviders**
> ProviderListDto ListProviders(ctx, projectUid)
Get suggested providers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**ProviderListDto**](ProviderListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSegment**
> SearchResponseListTmDto SearchSegment(ctx, projectUid, optional)
Search translation memory for segment in the project

Returns at most <i>maxSegments</i>             records with <i>score >= scoreThreshold</i> and at most <i>maxSubsegments</i> records which are subsegment,             i.e. the source text is substring of the query text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiSearchSegmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiSearchSegmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SearchTmRequestDto**](SearchTmRequestDto.md)|  | 

### Return type

[**SearchResponseListTmDto**](SearchResponseListTmDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetFinancialSettings**
> FinancialSettingsDto SetFinancialSettings(ctx, projectUid, optional)
Edit financial settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiSetFinancialSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiSetFinancialSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SetFinancialSettingsDto**](SetFinancialSettingsDto.md)|  | 

### Return type

[**FinancialSettingsDto**](FinancialSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMtSettingsForProject**
> MtSettingsPerLanguageListDto SetMtSettingsForProject(ctx, projectUid, optional)
Edit machine translate settings

This will erase all mtSettings per language for project.         To remove all machine translate settings from project call without a machineTranslateSettings parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiSetMtSettingsForProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiSetMtSettingsForProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EditProjectMtSettingsDto**](EditProjectMtSettingsDto.md)|  | 

### Return type

[**MtSettingsPerLanguageListDto**](MTSettingsPerLanguageListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMtSettingsPerLanguageForProject**
> MtSettingsPerLanguageListDto SetMtSettingsPerLanguageForProject(ctx, projectUid, optional)
Edit machine translate settings per language

This will erase mtSettings for project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiSetMtSettingsPerLanguageForProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiSetMtSettingsPerLanguageForProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EditProjectMtSettPerLangListDto**](EditProjectMtSettPerLangListDto.md)|  | 

### Return type

[**MtSettingsPerLanguageListDto**](MTSettingsPerLanguageListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetProjectQASettingsV2**
> QaSettingsDtoV2 SetProjectQASettingsV2(ctx, projectUid, optional)
Edit quality assurance settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiSetProjectQASettingsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiSetProjectQASettingsV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EditQaSettingsDtoV2**](EditQaSettingsDtoV2.md)|  | 

### Return type

[**QaSettingsDtoV2**](QASettingsDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetProjectStatus**
> SetProjectStatus(ctx, projectUid, optional)
Edit project status



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiSetProjectStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiSetProjectStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SetProjectStatusDto**](SetProjectStatusDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetProjectTermBases**
> ProjectTermBaseListDto SetProjectTermBases(ctx, projectUid, optional)
Edit term bases



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiSetProjectTermBasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiSetProjectTermBasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SetTermBaseDto**](SetTermBaseDto.md)|  | 

### Return type

[**ProjectTermBaseListDto**](ProjectTermBaseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetProjectTransMemories**
> ProjectTransMemoryListDto SetProjectTransMemories(ctx, projectUid, optional)
Edit translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiSetProjectTransMemoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiSetProjectTransMemoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SetProjectTransMemoriesDto**](SetProjectTransMemoriesDto.md)|  | 

### Return type

[**ProjectTransMemoryListDto**](ProjectTransMemoryListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetProjectTransMemoriesV2**
> ProjectTransMemoryListDtoV2 SetProjectTransMemoriesV2(ctx, projectUid, optional)
Edit translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectApiSetProjectTransMemoriesV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiSetProjectTransMemoriesV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SetProjectTransMemoriesV2Dto**](SetProjectTransMemoriesV2Dto.md)|  | 

### Return type

[**ProjectTransMemoryListDtoV2**](ProjectTransMemoryListDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

