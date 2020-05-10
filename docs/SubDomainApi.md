# \SubDomainApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubDomain**](SubDomainApi.md#CreateSubDomain) | **Post** /api2/v1/subDomains | Create subdomain
[**DeleteSubDomain**](SubDomainApi.md#DeleteSubDomain) | **Delete** /api2/v1/subDomains/{subDomainId} | Delete subdomain
[**GetSubDomain**](SubDomainApi.md#GetSubDomain) | **Get** /api2/v1/subDomains/{subDomainId} | Get subdomain
[**ListSubDomains**](SubDomainApi.md#ListSubDomains) | **Get** /api2/v1/subDomains | List subdomains
[**UpdateSubDomain**](SubDomainApi.md#UpdateSubDomain) | **Put** /api2/v1/subDomains/{subDomainId} | Edit subdomain


# **CreateSubDomain**
> SubDomainDto CreateSubDomain(ctx, optional)
Create subdomain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubDomainApiCreateSubDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubDomainApiCreateSubDomainOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SubDomainEditDto**](SubDomainEditDto.md)|  | 

### Return type

[**SubDomainDto**](SubDomainDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubDomain**
> DeleteSubDomain(ctx, subDomainId)
Delete subdomain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subDomainId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubDomain**
> SubDomainDto GetSubDomain(ctx, subDomainId)
Get subdomain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subDomainId** | **int64**|  | 

### Return type

[**SubDomainDto**](SubDomainDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubDomains**
> PageDtoSubDomainDto ListSubDomains(ctx, optional)
List subdomains



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubDomainApiListSubDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubDomainApiListSubDomainsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoSubDomainDto**](PageDtoSubDomainDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubDomain**
> SubDomainDto UpdateSubDomain(ctx, subDomainId, optional)
Edit subdomain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subDomainId** | **int64**|  | 
 **optional** | ***SubDomainApiUpdateSubDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubDomainApiUpdateSubDomainOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SubDomainEditDto**](SubDomainEditDto.md)|  | 

### Return type

[**SubDomainDto**](SubDomainDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

