# \ProjectReferenceFileApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNoteRef**](ProjectReferenceFileApi.md#CreateNoteRef) | **Post** /api2/v1/projects/{projectUid}/references | Create project reference file
[**DownloadReference**](ProjectReferenceFileApi.md#DownloadReference) | **Get** /api2/v1/projects/{projectUid}/references/{referenceFileId} | Get project reference


# **CreateNoteRef**
> ReferenceFileReference CreateNoteRef(ctx, projectUid, optional)
Create project reference file

Accepts application/octet-stream or application/json.<br>                                                                     <b>application/json</b> - Note will be converted to .txt.<br>                                                                     <b>application/octet-stream</b> - Content-Disposition header is required

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***ProjectReferenceFileApiCreateNoteRefOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectReferenceFileApiCreateNoteRefOpts struct

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

