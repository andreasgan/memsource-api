# \VendorApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVendor**](VendorApi.md#GetVendor) | **Get** /api2/v1/vendors/{vendorId} | Get vendor
[**ListDiscovered**](VendorApi.md#ListDiscovered) | **Get** /api2/v1/discoveredVendors | List discovered vendors
[**ListVendors**](VendorApi.md#ListVendors) | **Get** /api2/v1/vendors | List vendors


# **GetVendor**
> VendorDto GetVendor(ctx, vendorId)
Get vendor



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorId** | **int64**|  | 

### Return type

[**VendorDto**](VendorDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDiscovered**
> PageDtoDiscoveredVendorDto ListDiscovered(ctx, optional)
List discovered vendors



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VendorApiListDiscoveredOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VendorApiListDiscoveredOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoDiscoveredVendorDto**](PageDtoDiscoveredVendorDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVendors**
> PageDtoVendorDto ListVendors(ctx, optional)
List vendors



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VendorApiListVendorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VendorApiListVendorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoVendorDto**](PageDtoVendorDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

