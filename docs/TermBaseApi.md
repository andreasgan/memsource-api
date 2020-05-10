# \TermBaseApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseTerms**](TermBaseApi.md#BrowseTerms) | **Post** /api2/v1/termBases/{termBaseId}/browse | Browse term base
[**ClearTermBase**](TermBaseApi.md#ClearTermBase) | **Delete** /api2/v1/termBases/{termBaseId}/terms | Clear term base
[**CreateTerm**](TermBaseApi.md#CreateTerm) | **Post** /api2/v1/termBases/{termBaseId}/terms | Create term
[**CreateTermBase**](TermBaseApi.md#CreateTermBase) | **Post** /api2/v1/termBases | Create term base
[**CreateTermByJob**](TermBaseApi.md#CreateTermByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/termBases/createByJob | Create term in job&#39;s term bases
[**DeleteConcept**](TermBaseApi.md#DeleteConcept) | **Delete** /api2/v1/termBases/{termBaseId}/concepts/{conceptId} | Delete concept
[**DeleteConcepts**](TermBaseApi.md#DeleteConcepts) | **Delete** /api2/v1/termBases/{termBaseId}/concepts | Delete concepts
[**DeleteTerm**](TermBaseApi.md#DeleteTerm) | **Delete** /api2/v1/termBases/{termBaseId}/terms/{termId} | Delete term
[**DeleteTermBase**](TermBaseApi.md#DeleteTermBase) | **Delete** /api2/v1/termBases/{termBaseId} | Delete term base
[**ExportTermBase**](TermBaseApi.md#ExportTermBase) | **Get** /api2/v1/termBases/{termBaseId}/export | Export term base
[**GetLastBackgroundTask**](TermBaseApi.md#GetLastBackgroundTask) | **Get** /api2/v1/termBases/{termBaseId}/lastBackgroundTask | Last import status
[**GetProjectTemplateTermBases**](TermBaseApi.md#GetProjectTemplateTermBases) | **Get** /api2/v1/projectTemplates/{projectTemplateId}/termBases | Get term bases
[**GetTerm**](TermBaseApi.md#GetTerm) | **Get** /api2/v1/termBases/{termBaseId}/terms/{termId} | Get term
[**GetTermBase**](TermBaseApi.md#GetTermBase) | **Get** /api2/v1/termBases/{termBaseId} | Get term base
[**GetTermBaseMetadata**](TermBaseApi.md#GetTermBaseMetadata) | **Get** /api2/v1/termBases/{termBaseId}/metadata | Get term base metadata
[**GetTranslationResources**](TermBaseApi.md#GetTranslationResources) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/translationResources | Get translation resources
[**ImportTermBase**](TermBaseApi.md#ImportTermBase) | **Post** /api2/v1/termBases/{termBaseId}/upload | Upload term base
[**ListTermBases**](TermBaseApi.md#ListTermBases) | **Get** /api2/v1/termBases | List term bases
[**ListTermsOfConcept**](TermBaseApi.md#ListTermsOfConcept) | **Get** /api2/v1/termBases/{termBaseId}/concepts/{conceptId}/terms | Get terms of concept
[**SearchTerms**](TermBaseApi.md#SearchTerms) | **Post** /api2/v1/termBases/{termBaseId}/search | Search term base
[**SearchTermsByJob**](TermBaseApi.md#SearchTermsByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/termBases/searchByJob | Search job&#39;s term bases
[**SearchTermsByJob1**](TermBaseApi.md#SearchTermsByJob1) | **Post** /api2/v2/projects/{projectUid}/jobs/{jobUid}/termBases/searchByJob | Search job&#39;s term bases
[**SearchTermsInTextByJob**](TermBaseApi.md#SearchTermsInTextByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/termBases/searchInTextByJob | Search terms in text
[**SearchTermsInTextByJobV2**](TermBaseApi.md#SearchTermsInTextByJobV2) | **Post** /api2/v2/projects/{projectUid}/jobs/{jobUid}/termBases/searchInTextByJob | Search terms in text
[**SetProjectTemplateTermBases**](TermBaseApi.md#SetProjectTemplateTermBases) | **Put** /api2/v1/projectTemplates/{projectTemplateId}/termBases | Edit term bases in project template
[**UpdateTerm**](TermBaseApi.md#UpdateTerm) | **Put** /api2/v1/termBases/{termBaseId}/terms/{termId} | Edit term
[**UpdateTermBase**](TermBaseApi.md#UpdateTermBase) | **Put** /api2/v1/termBases/{termBaseId} | Edit term base


# **BrowseTerms**
> BrowseResponseListDto BrowseTerms(ctx, termBaseId, optional)
Browse term base



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
 **optional** | ***TermBaseApiBrowseTermsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiBrowseTermsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BrowseRequestDto**](BrowseRequestDto.md)|  | 

### Return type

[**BrowseResponseListDto**](BrowseResponseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearTermBase**
> ClearTermBase(ctx, termBaseId)
Clear term base

Deletes all terms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTerm**
> TermDto CreateTerm(ctx, termBaseId, optional)
Create term

Set conceptId to assign the term to an existing concept, otherwise a new concept will be created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
 **optional** | ***TermBaseApiCreateTermOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiCreateTermOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TermCreateDto**](TermCreateDto.md)|  | 

### Return type

[**TermDto**](TermDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTermBase**
> TermBaseDto CreateTermBase(ctx, optional)
Create term base



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TermBaseApiCreateTermBaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiCreateTermBaseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TermBaseEditDto**](TermBaseEditDto.md)|  | 

### Return type

[**TermBaseDto**](TermBaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTermByJob**
> TermPairDto CreateTermByJob(ctx, jobUid, projectUid, optional)
Create term in job's term bases

Create new term in the write term base assigned to the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***TermBaseApiCreateTermByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiCreateTermByJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of CreateTermsDto**](CreateTermsDto.md)|  | 

### Return type

[**TermPairDto**](TermPairDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConcept**
> DeleteConcept(ctx, termBaseId, conceptId)
Delete concept



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
  **conceptId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConcepts**
> DeleteConcepts(ctx, termBaseId, optional)
Delete concepts



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
 **optional** | ***TermBaseApiDeleteConceptsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiDeleteConceptsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ConceptListReference**](ConceptListReference.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTerm**
> DeleteTerm(ctx, termBaseId, termId)
Delete term



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
  **termId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTermBase**
> DeleteTermBase(ctx, termBaseId, optional)
Delete term base



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
 **optional** | ***TermBaseApiDeleteTermBaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiDeleteTermBaseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **optional.Bool**| purge&#x3D;false - the Termbase is can later be restored,                      \&quot;purge&#x3D;true - the Termbase is completely deleted and cannot be restored | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportTermBase**
> ExportTermBase(ctx, termBaseId, optional)
Export term base



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
 **optional** | ***TermBaseApiExportTermBaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiExportTermBaseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | [default to Tbx]
 **charset** | **optional.String**|  | [default to UTF-8]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLastBackgroundTask**
> BackgroundTasksTbDto GetLastBackgroundTask(ctx, termBaseId)
Last import status



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 

### Return type

[**BackgroundTasksTbDto**](BackgroundTasksTbDto.md)

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

# **GetTerm**
> TermDto GetTerm(ctx, termBaseId, termId)
Get term



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
  **termId** | **string**|  | 

### Return type

[**TermDto**](TermDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTermBase**
> TermBaseDto GetTermBase(ctx, termBaseId)
Get term base



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 

### Return type

[**TermBaseDto**](TermBaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTermBaseMetadata**
> MetadataTbDto GetTermBaseMetadata(ctx, termBaseId)
Get term base metadata



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 

### Return type

[**MetadataTbDto**](MetadataTbDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTranslationResources**
> TranslationResourcesDto GetTranslationResources(ctx, projectUid, jobUid)
Get translation resources



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 

### Return type

[**TranslationResourcesDto**](TranslationResourcesDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportTermBase**
> ImportResponse ImportTermBase(ctx, contentDisposition, termBaseId, optional)
Upload term base

 Terms can be imported from XLS/XLSX and TBX file formats into a term base. See <a target=\"_blank\" href=\"https://help.memsource.com/hc/en-us/articles/115003681851-Term-Bases\">Memsource Wiki</a> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentDisposition** | **string**| must match pattern &#x60;((inline|attachment); )?filename\\*&#x3D;UTF-8&#39;&#39;(.+)&#x60; | 
  **termBaseId** | **int64**|  | 
 **optional** | ***TermBaseApiImportTermBaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiImportTermBaseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of InputStream**](InputStream.md)|  | 
 **charset** | **optional.String**|  | [default to UTF-8]
 **strictLangMatching** | **optional.Bool**|  | [default to false]
 **updateTerms** | **optional.Bool**|  | [default to true]

### Return type

[**ImportResponse**](ImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTermBases**
> PageDtoTermBaseDto ListTermBases(ctx, optional)
List term bases



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TermBaseApiListTermBasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiListTermBasesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **lang** | [**optional.Interface of []string**](string.md)| Language of the term base | 
 **clientId** | **optional.String**|  | 
 **domainId** | **optional.String**|  | 
 **subDomainId** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoTermBaseDto**](PageDtoTermBaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTermsOfConcept**
> ConceptDto ListTermsOfConcept(ctx, termBaseId, conceptId)
Get terms of concept



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
  **conceptId** | **string**|  | 

### Return type

[**ConceptDto**](ConceptDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTerms**
> SearchResponseListTbDto SearchTerms(ctx, termBaseId, optional)
Search term base



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
 **optional** | ***TermBaseApiSearchTermsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiSearchTermsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TermBaseSearchRequestDto**](TermBaseSearchRequestDto.md)|  | 

### Return type

[**SearchResponseListTbDto**](SearchResponseListTbDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTermsByJob**
> SearchResponseListTbDto SearchTermsByJob(ctx, jobUid, projectUid, optional)
Search job's term bases

Search all read term bases assigned to the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***TermBaseApiSearchTermsByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiSearchTermsByJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchTbByJobRequestDto**](SearchTbByJobRequestDto.md)|  | 

### Return type

[**SearchResponseListTbDto**](SearchResponseListTbDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTermsByJob1**
> SearchTbResponseListDto SearchTermsByJob1(ctx, jobUid, projectUid, optional)
Search job's term bases

Search all read term bases assigned to the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***TermBaseApiSearchTermsByJob1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiSearchTermsByJob1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchTbByJobRequestDto**](SearchTbByJobRequestDto.md)|  | 

### Return type

[**SearchTbResponseListDto**](SearchTbResponseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTermsInTextByJob**
> SearchInTextResponseListDto SearchTermsInTextByJob(ctx, jobUid, projectUid, optional)
Search terms in text

Search in text in all read term bases assigned to the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***TermBaseApiSearchTermsInTextByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiSearchTermsInTextByJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchInTextByJobRequestDto**](SearchInTextByJobRequestDto.md)|  | 

### Return type

[**SearchInTextResponseListDto**](SearchInTextResponseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTermsInTextByJobV2**
> SearchInTextResponseList2Dto SearchTermsInTextByJobV2(ctx, jobUid, projectUid, optional)
Search terms in text

Search in text in all read term bases assigned to the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***TermBaseApiSearchTermsInTextByJobV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiSearchTermsInTextByJobV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchTbInTextByJobRequestDto**](SearchTbInTextByJobRequestDto.md)|  | 

### Return type

[**SearchInTextResponseList2Dto**](SearchInTextResponseList2Dto.md)

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
 **optional** | ***TermBaseApiSetProjectTemplateTermBasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiSetProjectTemplateTermBasesOpts struct

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

# **UpdateTerm**
> TermDto UpdateTerm(ctx, termBaseId, termId, optional)
Edit term



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
  **termId** | **string**|  | 
 **optional** | ***TermBaseApiUpdateTermOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiUpdateTermOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TermEditDto**](TermEditDto.md)|  | 

### Return type

[**TermDto**](TermDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTermBase**
> TermBaseDto UpdateTermBase(ctx, termBaseId, optional)
Edit term base

It is possible to add new languages only

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **termBaseId** | **int64**|  | 
 **optional** | ***TermBaseApiUpdateTermBaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TermBaseApiUpdateTermBaseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TermBaseEditDto**](TermBaseEditDto.md)|  | 

### Return type

[**TermBaseDto**](TermBaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

