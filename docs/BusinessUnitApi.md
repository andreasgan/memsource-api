# \BusinessUnitApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBusinessUnit**](BusinessUnitApi.md#CreateBusinessUnit) | **Post** /api2/v1/businessUnits | Create business unit
[**DeleteBusinessUnit**](BusinessUnitApi.md#DeleteBusinessUnit) | **Delete** /api2/v1/businessUnits/{businessUnitId} | Delete business unit
[**GetBusinessUnit**](BusinessUnitApi.md#GetBusinessUnit) | **Get** /api2/v1/businessUnits/{businessUnitId} | Get business unit
[**ListBusinessUnits**](BusinessUnitApi.md#ListBusinessUnits) | **Get** /api2/v1/businessUnits | List business units
[**UpdateBusinessUnit**](BusinessUnitApi.md#UpdateBusinessUnit) | **Put** /api2/v1/businessUnits/{businessUnitId} | Edit business unit


# **CreateBusinessUnit**
> BusinessUnitDto CreateBusinessUnit(ctx, optional)
Create business unit



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BusinessUnitApiCreateBusinessUnitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessUnitApiCreateBusinessUnitOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BusinessUnitEditDto**](BusinessUnitEditDto.md)|  | 

### Return type

[**BusinessUnitDto**](BusinessUnitDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBusinessUnit**
> DeleteBusinessUnit(ctx, businessUnitId)
Delete business unit



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **businessUnitId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBusinessUnit**
> BusinessUnitDto GetBusinessUnit(ctx, businessUnitId)
Get business unit



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **businessUnitId** | **int64**|  | 

### Return type

[**BusinessUnitDto**](BusinessUnitDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBusinessUnits**
> PageDtoBusinessUnitDto ListBusinessUnits(ctx, optional)
List business units



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BusinessUnitApiListBusinessUnitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessUnitApiListBusinessUnitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Unique name of the business unit | 
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoBusinessUnitDto**](PageDtoBusinessUnitDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBusinessUnit**
> BusinessUnitDto UpdateBusinessUnit(ctx, businessUnitId, optional)
Edit business unit



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **businessUnitId** | **int64**|  | 
 **optional** | ***BusinessUnitApiUpdateBusinessUnitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessUnitApiUpdateBusinessUnitOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BusinessUnitEditDto**](BusinessUnitEditDto.md)|  | 

### Return type

[**BusinessUnitDto**](BusinessUnitDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

