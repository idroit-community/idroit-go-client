# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersControllerCreate**](UsersApi.md#UsersControllerCreate) | **Post** /api/v1/users | ユーザーアカウント作成
[**UsersControllerExportKey**](UsersApi.md#UsersControllerExportKey) | **Post** /api/v1/users/{id}/keys/{keyId}/mail | アカウントへの鍵のメール送信
[**UsersControllerFindAll**](UsersApi.md#UsersControllerFindAll) | **Get** /api/v1/users | ユーザーアカウント一覧取得
[**UsersControllerFindOne**](UsersApi.md#UsersControllerFindOne) | **Get** /api/v1/users/{id} | ユーザーアカウント詳細取得
[**UsersControllerInvite**](UsersApi.md#UsersControllerInvite) | **Post** /api/v1/users/{id}/invite | ユーザーアカウントへのアカウント有効化メール送信
[**UsersControllerRegisterUserDid**](UsersApi.md#UsersControllerRegisterUserDid) | **Post** /api/v1/users/{id}/dids | Get the count of clients
[**UsersControllerRegistration**](UsersApi.md#UsersControllerRegistration) | **Post** /api/v1/users/registration | ユーザーアカウントへのアカウント有効化、登録フロー
[**UsersControllerRemove**](UsersApi.md#UsersControllerRemove) | **Delete** /api/v1/users/{id} | ユーザーアカウント停止
[**UsersControllerSendDid**](UsersApi.md#UsersControllerSendDid) | **Post** /api/v1/users/{id}/dids/{didInfoId}/mail | アカウントへのDIDのメール送信
[**UsersControllerSendKey**](UsersApi.md#UsersControllerSendKey) | **Post** /api/v1/users/{id}/keys/{didInfoId}/mail | アカウントへのDIDのメール送信
[**UsersControllerSendVc**](UsersApi.md#UsersControllerSendVc) | **Post** /api/v1/users/{id}/vcs/{vcInfoId}/mail | アカウントへのVCのメール送信
[**UsersControllerUpdate**](UsersApi.md#UsersControllerUpdate) | **Put** /api/v1/users/{id} | ユーザーアカウント更新

# **UsersControllerCreate**
> UserResponseDto UsersControllerCreate(ctx, body)
ユーザーアカウント作成

新規ユーザーアカウントを作成します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserDto**](CreateUserDto.md)|  | 

### Return type

[**UserResponseDto**](UserResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerExportKey**
> UsersControllerExportKey(ctx, id, keyId)
アカウントへの鍵のメール送信

クライアントアカウントに紐付いたDIDの鍵情報をメールで送信し、共有します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **keyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerFindAll**
> UsersResponseDto UsersControllerFindAll(ctx, optional)
ユーザーアカウント一覧取得

ユーザーアカウントを一覧として値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| (必須) ページネーションを表示する際のページ数に当たるクエリパラメーター。全件取得する際は1を指定。(デフォルト: 1) | 
 **limit** | **optional.Float64**| (必須) ページネーションを表示する際のページ数あたりに表示する件数を指定するクエリパラメーター。全件取得する際は0を指定。(デフォルト: 10) | 
 **name** | **optional.String**| (任意) 例: \&quot;Jhon Doe\&quot; | 
 **email** | **optional.String**| (任意) 例: \&quot;client1@email.com\&quot; | 
 **status** | **optional.String**| (任意) 例: \&quot;active\&quot; | 
 **role** | **optional.String**| (任意) 例: \&quot;client\&quot; | 
 **needActivateFlow** | **optional.Bool**| (任意) 例: true, false | 

### Return type

[**UsersResponseDto**](UsersResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerFindOne**
> UserResponseDto UsersControllerFindOne(ctx, id)
ユーザーアカウント詳細取得

リクエストパラメータのidで指定された単一の管理者アカウントの詳細情報の値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**UserResponseDto**](UserResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerInvite**
> UsersControllerInvite(ctx, id)
ユーザーアカウントへのアカウント有効化メール送信

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerRegisterUserDid**
> UsersControllerRegisterUserDid(ctx, body, id)
Get the count of clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GenerateDidDto**](GenerateDidDto.md)|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerRegistration**
> UsersControllerRegistration(ctx, body, token)
ユーザーアカウントへのアカウント有効化、登録フロー

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegistrationProcessDto**](RegistrationProcessDto.md)|  | 
  **token** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerRemove**
> UsersControllerRemove(ctx, id)
ユーザーアカウント停止

リクエストパラメータのidで指定された単一のユーザーアカウントを停止します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerSendDid**
> UsersControllerSendDid(ctx, id, didInfoId)
アカウントへのDIDのメール送信

アカウントに紐付いたDIDをメールで送信し、共有します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **didInfoId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerSendKey**
> UsersControllerSendKey(ctx, id, didInfoId)
アカウントへのDIDのメール送信

アカウントに紐付いたDIDをメールで送信し、共有します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **didInfoId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerSendVc**
> UsersControllerSendVc(ctx, id, vcInfoId)
アカウントへのVCのメール送信

アカウントに紐付いたVCをメールで送信し、共有します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **vcInfoId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerUpdate**
> UserResponseDto UsersControllerUpdate(ctx, body, id)
ユーザーアカウント更新

リクエストパラメータのidで指定された単一のユーザーアカウント情報を更新します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserDto**](CreateUserDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**UserResponseDto**](UserResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

