# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsControllerFindAll**](SettingsApi.md#SettingsControllerFindAll) | **Get** /api/v1/settings | Get all admins
[**SettingsControllerFindOne**](SettingsApi.md#SettingsControllerFindOne) | **Get** /api/v1/settings/{key} | Get a specific admin

# **SettingsControllerFindAll**
> SettingListResponseDto SettingsControllerFindAll(ctx, )
Get all admins

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingListResponseDto**](SettingListResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsControllerFindOne**
> SettingResponseDto SettingsControllerFindOne(ctx, key)
Get a specific admin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 

### Return type

[**SettingResponseDto**](SettingResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

