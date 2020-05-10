# \FileApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUrlFile**](FileApi.md#CreateUrlFile) | **Post** /api2/v1/files | Upload file
[**DeletesFile**](FileApi.md#DeletesFile) | **Delete** /api2/v1/files/{fileUid} | Delete file
[**GetFileJson**](FileApi.md#GetFileJson) | **Get** /api2/v1/files/{fileUid} | Get file
[**GetFiles**](FileApi.md#GetFiles) | **Get** /api2/v1/files | List files


# **CreateUrlFile**
> UploadedFileDto CreateUrlFile(ctx, body, contentDisposition)
Upload file

Accepts multipart/form-data, application/octet-stream or application/json.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoteUploadedFileDto**](RemoteUploadedFileDto.md)| file | 
  **contentDisposition** | **string**| must match pattern &#x60;((inline|attachment); )?filename\\*&#x3D;UTF-8&#39;&#39;(.+)&#x60; | 

### Return type

[**UploadedFileDto**](UploadedFileDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/octet-stream, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletesFile**
> DeletesFile(ctx, fileUid)
Delete file



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileUid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileJson**
> UploadedFileDto GetFileJson(ctx, fileUid)
Get file

Get uploaded file as <b>octet-stream</b> or as <b>json</b> based on 'Accept' header

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileUid** | **string**|  | 

### Return type

[**UploadedFileDto**](UploadedFileDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFiles**
> PageDtoUploadedFileDto GetFiles(ctx, optional)
List files



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileApiGetFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileApiGetFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]
 **name** | **optional.String**|  | 
 **types** | [**optional.Interface of []string**](string.md)|  | 
 **createdBy** | **optional.Int64**|  | 
 **biggerThan** | **optional.Int64**| Size in bytes | 

### Return type

[**PageDtoUploadedFileDto**](PageDtoUploadedFileDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

