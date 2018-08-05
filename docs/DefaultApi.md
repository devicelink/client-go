# \DefaultApi

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicelinkAddDeviceToGroup**](DefaultApi.md#DevicelinkAddDeviceToGroup) | **Put** /v1/groups/{group_id}/device/{id} | Adds device to group.
[**DevicelinkCreateGroup**](DefaultApi.md#DevicelinkCreateGroup) | **Post** /v1/groups | Creates a device group.
[**DevicelinkDeleteGroup**](DefaultApi.md#DevicelinkDeleteGroup) | **Delete** /v1/groups/{id} | Deletes a group.
[**DevicelinkDoAction**](DefaultApi.md#DevicelinkDoAction) | **Post** /v1/device/{id}/actions | Prompts a device action.
[**DevicelinkGetActions**](DefaultApi.md#DevicelinkGetActions) | **Get** /v1/device/{id}/actions | Gets device actions.
[**DevicelinkGetData**](DefaultApi.md#DevicelinkGetData) | **Get** /v1/device/{id}/properties | Gets device properties.
[**DevicelinkGetDeviceEvent**](DefaultApi.md#DevicelinkGetDeviceEvent) | **Get** /v1/device/{id}/events | Gets device events.
[**DevicelinkGetGroupInfo**](DefaultApi.md#DevicelinkGetGroupInfo) | **Get** /v1/groups/{id} | Gets a device group.
[**DevicelinkGetGroups**](DefaultApi.md#DevicelinkGetGroups) | **Get** /v1/groups | Gets device groups.
[**DevicelinkGetSpecificAction**](DefaultApi.md#DevicelinkGetSpecificAction) | **Get** /v1/device/{id}/actions/{action-id} | Gets a device action.
[**DevicelinkRemoveDeviceFromGroup**](DefaultApi.md#DevicelinkRemoveDeviceFromGroup) | **Delete** /v1/groups/{group_id}/device/{id} | Removes device from a group.
[**DevicelinkUpdateGroupMetadata**](DefaultApi.md#DevicelinkUpdateGroupMetadata) | **Patch** /v1/groups/{id} | Updates group information.


# **DevicelinkAddDeviceToGroup**
> DevicelinkAddDeviceToGroup(ctx, groupId, id)
Adds device to group.

Adds a device to a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | [**string**](.md)| uuid of group | 
  **id** | [**string**](.md)| uuid of device | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkCreateGroup**
> InlineResponse201 DevicelinkCreateGroup(ctx, optional)
Creates a device group.

Creates virtual group for devices. Request can include list of devices as request-body or attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicelinkCreateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicelinkCreateGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body**](Body.md)| List of devices | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkDeleteGroup**
> DevicelinkDeleteGroup(ctx, id)
Deletes a group.

Deletes group and its metadata.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| uuid of group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkDoAction**
> InlineResponse2011 DevicelinkDoAction(ctx, id, body2)
Prompts a device action.

Prompts an action on a device.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| uuid of device | 
  **body2** | [**Body2**](Body2.md)| Action that shall be executed | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkGetActions**
> map[string]interface{} DevicelinkGetActions(ctx, id, optional)
Gets device actions.

Returns device actions within a time range.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| uuid of device | 
 **optional** | ***DevicelinkGetActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicelinkGetActionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.Int32**| Range start in seconds since epoch | 
 **to** | **optional.Int32**| Range end in seconds since epoch | 
 **start** | **optional.Int32**| The start key of the first action object to return, | 
 **sort** | **optional.String**| If set to asc (default), action objects are returned in time ascending order (i.e. starting with from or start and ending with stop). If set to desc, action objects are returned in time descending order (i.e. starting with to or start and ending with from). | [default to desc]
 **count** | **optional.Int32**| Limits the number of returned action objects | [default to 100]

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkGetData**
> DeviceData DevicelinkGetData(ctx, id, optional)
Gets device properties.

Returns properties for a device.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| uuid of device | 
 **optional** | ***DevicelinkGetDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicelinkGetDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.Int32**| Start time in seconds since epoch | 
 **to** | **optional.Int32**| End time in seconds since epoch | 
 **start** | **optional.Int32**| The start key of the first property object to return | 
 **sort** | **optional.String**| Order in which roperty objects are returned | [default to desc]
 **count** | **optional.Int32**| Limits the number of returned property objects | [default to 100]
 **propertyKey** | **optional.Int32**| Filters returned property objects by property-key | 
 **index** | **optional.Int32**| Filters returned property objects by property-index (can only be used in combination with property-key parameter) | 

### Return type

[**DeviceData**](DeviceData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkGetDeviceEvent**
> map[string]interface{} DevicelinkGetDeviceEvent(ctx, id)
Gets device events.

Opens event-stream for a device.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| uuid of device | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkGetGroupInfo**
> GroupInfo DevicelinkGetGroupInfo(ctx, id)
Gets a device group.

Returns a group, members, and metadata.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| uuid of group | 

### Return type

[**GroupInfo**](GroupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkGetGroups**
> GroupInfo DevicelinkGetGroups(ctx, )
Gets device groups.

Returns all groups with devices and metadata.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GroupInfo**](GroupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkGetSpecificAction**
> InlineResponse200 DevicelinkGetSpecificAction(ctx, id, actionId)
Gets a device action.

Returns a action of device.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| uuid of device | 
  **actionId** | **int32**| Action key | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkRemoveDeviceFromGroup**
> DevicelinkRemoveDeviceFromGroup(ctx, id2, id)
Removes device from a group.

Removes a device from a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id2** | [**string**](.md)| uuid of group | 
  **id** | [**string**](.md)| uuid of device | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicelinkUpdateGroupMetadata**
> DevicelinkUpdateGroupMetadata(ctx, id, optional)
Updates group information.

Updates group metadata.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| uuid of group | 
 **optional** | ***DevicelinkUpdateGroupMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicelinkUpdateGroupMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body1** | [**optional.Interface of Body1**](Body1.md)| New metadata of group | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

