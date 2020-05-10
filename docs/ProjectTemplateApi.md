# \ProjectTemplateApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignLinguistsFromTemplate**](ProjectTemplateApi.md#AssignLinguistsFromTemplate) | **Post** /api2/v1/projects/{projectUid}/applyTemplate/{templateId}/assignProviders | Assigns providers from template
[**AssignLinguistsFromTemplateToJobParts**](ProjectTemplateApi.md#AssignLinguistsFromTemplateToJobParts) | **Post** /api2/v1/projects/{projectUid}/applyTemplate/{templateId}/assignProviders/forJobParts | Assigns providers from template (specific jobs)
[**AssignableTemplates**](ProjectTemplateApi.md#AssignableTemplates) | **Get** /api2/v1/projects/{projectUid}/assignableTemplates | List assignable templates
[**CreateProjectFromTemplate**](ProjectTemplateApi.md#CreateProjectFromTemplate) | **Post** /api2/v1/projects/applyTemplate/{templateId} | Create project from template
[**CreateProjectFromTemplateV2**](ProjectTemplateApi.md#CreateProjectFromTemplateV2) | **Post** /api2/v2/projects/applyTemplate/{templateId} | Create project from template
[**CreateProjectTemplate**](ProjectTemplateApi.md#CreateProjectTemplate) | **Post** /api2/v1/projectTemplates | Create project template
[**DeleteProjectTemplate**](ProjectTemplateApi.md#DeleteProjectTemplate) | **Delete** /api2/v1/projectTemplates/{projectTemplateId} | Delete project template
[**EditProjectTemplate**](ProjectTemplateApi.md#EditProjectTemplate) | **Put** /api2/v1/projectTemplates/{projectTemplateId} | Edit project template
[**GetAnalyseSettingsForProjectTemplate**](ProjectTemplateApi.md#GetAnalyseSettingsForProjectTemplate) | **Get** /api2/v1/projectTemplates/{projectTemplateId}/analyseSettings | Get analyse settings
[**GetMachineTranslateSettingsForProjectTemplate**](ProjectTemplateApi.md#GetMachineTranslateSettingsForProjectTemplate) | **Get** /api2/v1/projectTemplates/{projectTemplateId}/mtSettings | Get machine translate settings
[**GetPreTranslateSettingsForProjectTemplate**](ProjectTemplateApi.md#GetPreTranslateSettingsForProjectTemplate) | **Get** /api2/v1/projectTemplates/{projectTemplateId}/preTranslateSettings | Get Pre-translate settings
[**GetProjectTemplate**](ProjectTemplateApi.md#GetProjectTemplate) | **Get** /api2/v1/projectTemplates/{projectTemplateId} | Get project template
[**GetProjectTemplateTermBases**](ProjectTemplateApi.md#GetProjectTemplateTermBases) | **Get** /api2/v1/projectTemplates/{projectTemplateId}/termBases | Get term bases
[**GetProjectTemplateTransMemories**](ProjectTemplateApi.md#GetProjectTemplateTransMemories) | **Get** /api2/v1/projectTemplates/{projectTemplateId}/transMemories | Get translation memories
[**GetProjectTemplates**](ProjectTemplateApi.md#GetProjectTemplates) | **Get** /api2/v1/projectTemplates | List project templates
[**SetProjectTemplateTermBases**](ProjectTemplateApi.md#SetProjectTemplateTermBases) | **Put** /api2/v1/projectTemplates/{projectTemplateId}/termBases | Edit term bases in project template
[**SetProjectTemplateTransMemories**](ProjectTemplateApi.md#SetProjectTemplateTransMemories) | **Put** /api2/v1/projectTemplates/{projectTemplateId}/transMemories | Edit translation memories


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
 **optional** | ***ProjectTemplateApiAssignLinguistsFromTemplateToJobPartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectTemplateApiAssignLinguistsFromTemplateToJobPartsOpts struct

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

# **CreateProjectFromTemplate**
> AbstractProjectDto CreateProjectFromTemplate(ctx, templateId, optional)
Create project from template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**|  | 
 **optional** | ***ProjectTemplateApiCreateProjectFromTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectTemplateApiCreateProjectFromTemplateOpts struct

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
 **optional** | ***ProjectTemplateApiCreateProjectFromTemplateV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectTemplateApiCreateProjectFromTemplateV2Opts struct

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

# **CreateProjectTemplate**
> ProjectTemplateDto CreateProjectTemplate(ctx, body)
Create project template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectTemplateCreateActionDto**](ProjectTemplateCreateActionDto.md)|  | 

### Return type

[**ProjectTemplateDto**](ProjectTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectTemplate**
> DeleteProjectTemplate(ctx, projectTemplateId)
Delete project template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditProjectTemplate**
> ProjectTemplateDto EditProjectTemplate(ctx, projectTemplateId, body)
Edit project template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **int64**|  | 
  **body** | [**ProjectTemplateEditDto**](ProjectTemplateEditDto.md)|  | 

### Return type

[**ProjectTemplateDto**](ProjectTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnalyseSettingsForProjectTemplate**
> AnalyseSettingsDto GetAnalyseSettingsForProjectTemplate(ctx, projectTemplateId)
Get analyse settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **string**|  | 

### Return type

[**AnalyseSettingsDto**](AnalyseSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMachineTranslateSettingsForProjectTemplate**
> MtSettingsPerLanguageListDto GetMachineTranslateSettingsForProjectTemplate(ctx, projectTemplateId)
Get machine translate settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **string**|  | 

### Return type

[**MtSettingsPerLanguageListDto**](MTSettingsPerLanguageListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPreTranslateSettingsForProjectTemplate**
> PreTranslateSettingsDto GetPreTranslateSettingsForProjectTemplate(ctx, projectTemplateId)
Get Pre-translate settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **string**|  | 

### Return type

[**PreTranslateSettingsDto**](PreTranslateSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectTemplate**
> ProjectTemplateDto GetProjectTemplate(ctx, projectTemplateId)
Get project template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **int64**|  | 

### Return type

[**ProjectTemplateDto**](ProjectTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectTemplateTermBases**
> ProjectTemplateTermBaseListDto GetProjectTemplateTermBases(ctx, projectTemplateId)
Get term bases



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **string**|  | 

### Return type

[**ProjectTemplateTermBaseListDto**](ProjectTemplateTermBaseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectTemplateTransMemories**
> ProjectTemplateTransMemoryListDto GetProjectTemplateTransMemories(ctx, projectTemplateId)
Get translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **string**|  | 

### Return type

[**ProjectTemplateTransMemoryListDto**](ProjectTemplateTransMemoryListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectTemplates**
> PageDtoProjectTemplateReference GetProjectTemplates(ctx, optional)
List project templates



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectTemplateApiGetProjectTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectTemplateApiGetProjectTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **clientId** | **optional.Int64**|  | 
 **clientName** | **optional.String**|  | 
 **ownerUid** | **optional.String**|  | 
 **domainName** | **optional.String**|  | 
 **subDomainName** | **optional.String**|  | 
 **costCenterId** | **optional.Int64**|  | 
 **costCenterName** | **optional.String**|  | 
 **businessUnitName** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoProjectTemplateReference**](PageDtoProjectTemplateReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetProjectTemplateTermBases**
> ProjectTemplateTermBaseListDto SetProjectTemplateTermBases(ctx, projectTemplateId, optional)
Edit term bases in project template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **string**|  | 
 **optional** | ***ProjectTemplateApiSetProjectTemplateTermBasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectTemplateApiSetProjectTemplateTermBasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SetProjectTemplateTermBaseDto**](SetProjectTemplateTermBaseDto.md)|  | 

### Return type

[**ProjectTemplateTermBaseListDto**](ProjectTemplateTermBaseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetProjectTemplateTransMemories**
> ProjectTemplateTransMemoryListDto SetProjectTemplateTransMemories(ctx, projectTemplateId, optional)
Edit translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **string**|  | 
 **optional** | ***ProjectTemplateApiSetProjectTemplateTransMemoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectTemplateApiSetProjectTemplateTransMemoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SetProjectTemplateTransMemoriesDto**](SetProjectTemplateTransMemoriesDto.md)|  | 

### Return type

[**ProjectTemplateTransMemoryListDto**](ProjectTemplateTransMemoryListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

