# \SCIMApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserSCIM**](SCIMApi.md#CreateUserSCIM) | **Post** /api2/v1/scim/Users | Create user using SCIM
[**DeleteUser**](SCIMApi.md#DeleteUser) | **Delete** /api2/v1/scim/Users/{userId} | Delete user using SCIM
[**EditUser**](SCIMApi.md#EditUser) | **Put** /api2/v1/scim/Users/{userId} | Edit user using SCIM
[**GetResourceTypes**](SCIMApi.md#GetResourceTypes) | **Get** /api2/v1/scim/ResourceTypes | List the types of SCIM Resources available
[**GetSchemaByUrn**](SCIMApi.md#GetSchemaByUrn) | **Get** /api2/v1/scim/Schemas/{schemaUrn} | Get supported SCIM Schema by urn
[**GetSchemas**](SCIMApi.md#GetSchemas) | **Get** /api2/v1/scim/Schemas | Get supported SCIM Schemas
[**GetScimUser**](SCIMApi.md#GetScimUser) | **Get** /api2/v1/scim/Users/{userId} | Get user
[**GetServiceProviderConfigDto**](SCIMApi.md#GetServiceProviderConfigDto) | **Get** /api2/v1/scim/ServiceProviderConfig | Retrieve the Service Provider&#39;s Configuration
[**PatchUser**](SCIMApi.md#PatchUser) | **Patch** /api2/v1/scim/Users/{userId} | Patch user using SCIM
[**SearchUsers**](SCIMApi.md#SearchUsers) | **Get** /api2/v1/scim/Users | Search users


# **CreateUserSCIM**
> ScimUserCoreDto CreateUserSCIM(ctx, optional)
Create user using SCIM

 Supported schema: `\"urn:ietf:params:scim:schemas:core:2.0:User\"`  Create active user: ``` {     \"schemas\": [         \"urn:ietf:params:scim:schemas:core:2.0:User\"     ],     \"active\": true,     \"userName\": \"john.doe\",     \"emails\": [         {             \"primary\": true,             \"value\": \"john.doe@example.com\",             \"type\": \"work\"         }     ],     \"name\": {         \"givenName\": \"John\",         \"familyName\": \"Doe\"     } } ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SCIMApiCreateUserSCIMOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SCIMApiCreateUserSCIMOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ScimUserCoreDto**](ScimUserCoreDto.md)|  | 
 **authorization** | **optional.String**|  | 

### Return type

[**ScimUserCoreDto**](ScimUserCoreDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/scim+json
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, userId, optional)
Delete user using SCIM



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***SCIMApiDeleteUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SCIMApiDeleteUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditUser**
> ScimUserCoreDto EditUser(ctx, userId, optional)
Edit user using SCIM



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***SCIMApiEditUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SCIMApiEditUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ScimUserCoreDto**](ScimUserCoreDto.md)|  | 
 **authorization** | **optional.String**|  | 

### Return type

[**ScimUserCoreDto**](ScimUserCoreDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/scim+json
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceTypes**
> ScimResourceTypeSchema GetResourceTypes(ctx, )
List the types of SCIM Resources available



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ScimResourceTypeSchema**](ScimResourceTypeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchemaByUrn**
> ScimResourceSchema GetSchemaByUrn(ctx, schemaUrn)
Get supported SCIM Schema by urn



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schemaUrn** | **string**|  | 

### Return type

[**ScimResourceSchema**](ScimResourceSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchemas**
> ScimResourceSchema GetSchemas(ctx, )
Get supported SCIM Schemas



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ScimResourceSchema**](ScimResourceSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScimUser**
> ScimUserCoreDto GetScimUser(ctx, userId, optional)
Get user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***SCIMApiGetScimUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SCIMApiGetScimUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**|  | 

### Return type

[**ScimUserCoreDto**](ScimUserCoreDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceProviderConfigDto**
> ServiceProviderConfigDto GetServiceProviderConfigDto(ctx, )
Retrieve the Service Provider's Configuration



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceProviderConfigDto**](ServiceProviderConfigDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchUser**
> ScimUserCoreDto PatchUser(ctx, userId, optional)
Patch user using SCIM



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***SCIMApiPatchUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SCIMApiPatchUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 
 **authorization** | **optional.String**|  | 

### Return type

[**ScimUserCoreDto**](ScimUserCoreDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchUsers**
> ScimResourceDto SearchUsers(ctx, optional)
Search users

 This operation supports <a href=\"http://ldapwiki.com/wiki/SCIM%20Filtering\" target=\"_blank\">SCIM Filter</a>,  <a href=\"http://ldapwiki.com/wiki/SCIM%20Search%20Request\" target=\"_blank\">SCIM attributes</a> and  <a href=\"http://ldapwiki.com/wiki/SCIM%20Sorting\" target=\"_blank\">SCIM sort</a>  Supported attributes:   - `id`   - `active`   - `userName`   - `name.givenName`   - `name.familyName`   - `emails.value`   - `meta.created` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SCIMApiSearchUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SCIMApiSearchUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **optional.String**|  | 
 **filter** | **optional.String**| See method description | 
 **attributes** | **optional.String**| See method description | 
 **sortBy** | **optional.String**| See method description | 
 **sortOrder** | **optional.String**| See method description | [default to ascending]
 **startIndex** | **optional.Int32**| The 1-based index of the first search result. Default 1 | [default to 1]
 **count** | **optional.Int32**| Non-negative Integer. Specifies the desired maximum number of search results per page; e.g., 10. | [default to 50]

### Return type

[**ScimResourceDto**](ScimResourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

