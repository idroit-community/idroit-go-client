# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthControllerGetProfile**](AuthApi.md#AuthControllerGetProfile) | **Get** /api/v1/auth/profile | ログイン済みのアカウント情報取得
[**AuthControllerLogin**](AuthApi.md#AuthControllerLogin) | **Post** /api/v1/auth/login | アカウントログインを実施

# **AuthControllerGetProfile**
> GetProfileResponseDto AuthControllerGetProfile(ctx, )
ログイン済みのアカウント情報取得

ログイン済みの管理者アカウントの情報を返却します。

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetProfileResponseDto**](GetProfileResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthControllerLogin**
> SignInResponseDto AuthControllerLogin(ctx, body)
アカウントログインを実施

アカウントログインを実行し、認証結果に応じてJSON Web Tokenの値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SignInDto**](SignInDto.md)|  | 

### Return type

[**SignInResponseDto**](SignInResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

