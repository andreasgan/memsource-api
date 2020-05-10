# \ConnectorApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnector**](ConnectorApi.md#GetConnector) | **Get** /api2/v1/connectors/{connectorId} | Get a connector
[**GetConnectorList**](ConnectorApi.md#GetConnectorList) | **Get** /api2/v1/connectors | List connectors
[**GetFile**](ConnectorApi.md#GetFile) | **Get** /api2/v1/connectors/{connectorId}/folders/{folder}/files/{file} | Download file
[**GetFolder**](ConnectorApi.md#GetFolder) | **Get** /api2/v1/connectors/{connectorId}/folders/{folder} | List files in a subfolder
[**GetRootFolder**](ConnectorApi.md#GetRootFolder) | **Get** /api2/v1/connectors/{connectorId}/folders | List files in root
[**UploadFile**](ConnectorApi.md#UploadFile) | **Post** /api2/v1/connectors/{connectorId}/folders/{folder} | Upload a file to a subfolder of the selected connector


# **GetConnector**
> ConnectorDto GetConnector(ctx, connectorId)
Get a connector



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorId** | **string**|  | 

### Return type

[**ConnectorDto**](ConnectorDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorList**
> ConnectorListDto GetConnectorList(ctx, optional)
List connectors



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectorApiGetConnectorListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorApiGetConnectorListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**|  | 

### Return type

[**ConnectorListDto**](ConnectorListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFile**
> InputStreamLength GetFile(ctx, connectorId, folder, file)
Download file

Download a file from a subfolder of the selected connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorId** | **string**|  | 
  **folder** | **string**|  | 
  **file** | **string**|  | 

### Return type

[**InputStreamLength**](InputStreamLength.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolder**
> FileListDto GetFolder(ctx, connectorId, folder, optional)
List files in a subfolder

List files in a subfolder of the selected connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorId** | **string**|  | 
  **folder** | **string**|  | 
 **optional** | ***ConnectorApiGetFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorApiGetFolderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileType** | **optional.String**|  | [default to ALL]
 **sort** | **optional.String**|  | [default to NAME]
 **direction** | **optional.String**|  | [default to ASCENDING]

### Return type

[**FileListDto**](FileListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRootFolder**
> FileListDto GetRootFolder(ctx, connectorId, optional)
List files in root

List files in a root folder of the selected connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorId** | **string**|  | 
 **optional** | ***ConnectorApiGetRootFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorApiGetRootFolderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileType** | **optional.String**|  | [default to ALL]
 **sort** | **optional.String**|  | [default to NAME]
 **direction** | **optional.String**|  | [default to ASCENDING]

### Return type

[**FileListDto**](FileListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFile**
> UploadResultDto UploadFile(ctx, connectorId, folder, contentType, file, optional)
Upload a file to a subfolder of the selected connector

Upload a file to a subfolder of the selected connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorId** | **string**|  | 
  **folder** | **string**|  | 
  **contentType** | **string**|  | 
  **file** | ***os.File**| Translated file to upload | 
 **optional** | ***ConnectorApiUploadFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorApiUploadFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **sourceFileName** | **optional.String**| Name or ID of the original file | 
 **subfolderName** | **optional.String**| Optional subfolder to upload the file to | 
 **mimeType** | **optional.String**| Mime type of the file to upload | 
 **commitMessage** | **optional.String**| Commit message for upload to Git, etc. | 

### Return type

[**UploadResultDto**](UploadResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

