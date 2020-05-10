# \UserApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDeletion**](UserApi.md#CancelDeletion) | **Post** /api2/v1/users/{userId}/undelete | Restore user
[**CreateUser**](UserApi.md#CreateUser) | **Post** /api2/v1/users | Create user
[**CreateUserV2**](UserApi.md#CreateUserV2) | **Post** /api2/v2/users | Create user
[**DeleteUser1**](UserApi.md#DeleteUser1) | **Delete** /api2/v1/users/{userId} | Delete user
[**GetListOfUsersFiltered**](UserApi.md#GetListOfUsersFiltered) | **Get** /api2/v1/users | List users
[**GetUser**](UserApi.md#GetUser) | **Get** /api2/v1/users/{userId} | Get user
[**GetUserV2**](UserApi.md#GetUserV2) | **Get** /api2/v2/users/{userId} | Get user
[**ListAssignedProjects**](UserApi.md#ListAssignedProjects) | **Get** /api2/v1/users/{userId}/projects | List assigned projects
[**ListJobs**](UserApi.md#ListJobs) | **Get** /api2/v1/users/{userId}/jobs | List assigned jobs
[**ListTargetLangs**](UserApi.md#ListTargetLangs) | **Get** /api2/v1/users/{userId}/targetLangs | List assigned target languages
[**ListWorkflowSteps**](UserApi.md#ListWorkflowSteps) | **Get** /api2/v1/users/{userId}/workflowSteps | List assigned workflow steps
[**LoginActivity**](UserApi.md#LoginActivity) | **Get** /api2/v1/users/{userId}/loginStatistics | Login statistics
[**SendLoginInfo**](UserApi.md#SendLoginInfo) | **Post** /api2/v1/users/{userId}/emailLoginInformation | Send login information
[**UpdatePassword**](UserApi.md#UpdatePassword) | **Post** /api2/v1/users/{userId}/updatePassword | Update password
[**UpdateUser**](UserApi.md#UpdateUser) | **Put** /api2/v1/users/{userId} | Edit user
[**UpdateUserV2**](UserApi.md#UpdateUserV2) | **Put** /api2/v2/users/{userId} | Edit user
[**UserLastLogins**](UserApi.md#UserLastLogins) | **Get** /api2/v1/users/lastLogins | List last login dates


# **CancelDeletion**
> UserDto CancelDeletion(ctx, userId)
Restore user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> UserDetailsDto CreateUser(ctx, optional)
Create user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiCreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserCreateDto**](UserCreateDto.md)|  | 

### Return type

[**UserDetailsDto**](UserDetailsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserV2**
> UserDetailsDtoV2 CreateUserV2(ctx, optional)
Create user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiCreateUserV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateUserV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserCreateDtoV2**](UserCreateDtoV2.md)|  | 

### Return type

[**UserDetailsDtoV2**](UserDetailsDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser1**
> DeleteUser1(ctx, userId)
Delete user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListOfUsersFiltered**
> PageDtoUserDto GetListOfUsersFiltered(ctx, optional)
List users



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiGetListOfUsersFilteredOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetListOfUsersFilteredOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstName** | **optional.String**| Filter for first name, that starts with value | 
 **lastName** | **optional.String**| Filter for last name, that starts with value | 
 **name** | **optional.String**| Filter for last name or first name, that starts with value | 
 **userName** | **optional.String**|  | 
 **email** | **optional.String**|  | 
 **role** | [**optional.Interface of []string**](string.md)|  | 
 **includeDeleted** | **optional.Bool**|  | [default to false]
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]
 **sort** | [**optional.Interface of []string**](string.md)|  | 
 **order** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**PageDtoUserDto**](PageDtoUserDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> UserDetailsDto GetUser(ctx, userId)
Get user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 

### Return type

[**UserDetailsDto**](UserDetailsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserV2**
> UserDetailsDtoV2 GetUserV2(ctx, userId)
Get user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 

### Return type

[**UserDetailsDtoV2**](UserDetailsDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
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
 **optional** | ***UserApiListAssignedProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListAssignedProjectsOpts struct

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

# **ListJobs**
> PageDtoAssignedJobDto ListJobs(ctx, userId, optional)
List assigned jobs



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***UserApiListJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListJobsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**optional.Interface of []string**](string.md)|  | 
 **projectUid** | **optional.String**|  | 
 **targetLang** | [**optional.Interface of []string**](string.md)|  | 
 **workflowStepId** | **optional.Int64**|  | 
 **dueInHours** | **optional.Int32**| -1 for jobs that are overdue | 
 **filename** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]

### Return type

[**PageDtoAssignedJobDto**](PageDtoAssignedJobDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTargetLangs**
> PageDtoString ListTargetLangs(ctx, userId, optional)
List assigned target languages



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***UserApiListTargetLangsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListTargetLangsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**optional.Interface of []string**](string.md)|  | 
 **projectUid** | **optional.String**|  | 
 **workflowStepId** | **optional.Int64**|  | 
 **dueInHours** | **optional.Int32**| -1 for jobs that are overdue | 
 **filename** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]

### Return type

[**PageDtoString**](PageDtoString.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWorkflowSteps**
> PageDtoWorkflowStepReference ListWorkflowSteps(ctx, userId, optional)
List assigned workflow steps



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***UserApiListWorkflowStepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListWorkflowStepsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**optional.Interface of []string**](string.md)|  | 
 **projectUid** | **optional.String**|  | 
 **targetLang** | [**optional.Interface of []string**](string.md)|  | 
 **dueInHours** | **optional.Int32**| -1 for jobs that are overdue | 
 **filename** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]

### Return type

[**PageDtoWorkflowStepReference**](PageDtoWorkflowStepReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoginActivity**
> UserStatisticsListDto LoginActivity(ctx, userId)
Login statistics



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 

### Return type

[**UserStatisticsListDto**](UserStatisticsListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendLoginInfo**
> SendLoginInfo(ctx, userId)
Send login information



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePassword**
> UpdatePassword(ctx, userId, optional)
Update password



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***UserApiUpdatePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdatePasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserPasswordEditDto**](UserPasswordEditDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UserDetailsDto UpdateUser(ctx, userId, optional)
Edit user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***UserApiUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserEditDto**](UserEditDto.md)|  | 

### Return type

[**UserDetailsDto**](UserDetailsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserV2**
> UserDetailsDtoV2 UpdateUserV2(ctx, userId, optional)
Edit user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***UserApiUpdateUserV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateUserV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserEditDtoV2**](UserEditDtoV2.md)|  | 

### Return type

[**UserDetailsDtoV2**](UserDetailsDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLastLogins**
> PageDtoLastLoginDto UserLastLogins(ctx, optional)
List last login dates



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiUserLastLoginsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserLastLoginsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **optional.String**|  | 
 **role** | [**optional.Interface of []string**](string.md)|  | 
 **sort** | [**optional.Interface of []string**](string.md)|  | 
 **order** | [**optional.Interface of []string**](string.md)|  | 
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 100, default 100 | [default to 100]

### Return type

[**PageDtoLastLoginDto**](PageDtoLastLoginDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

