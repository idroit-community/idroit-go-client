# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DidInfosControllerAddLabel**](DidInfosApi.md#DidInfosControllerAddLabel) | **Put** /api/v1/did-infos/{id} | DID情報へのラベル追加
[**DidInfosControllerConnectUser**](DidInfosApi.md#DidInfosControllerConnectUser) | **Put** /api/v1/did-infos/{id}/user/{user_id} | DID情報へのクライアントアカウント紐付け
[**DidInfosControllerCreate**](DidInfosApi.md#DidInfosControllerCreate) | **Post** /api/v1/did-infos | 新規DID生成
[**DidInfosControllerDisconnectUser**](DidInfosApi.md#DidInfosControllerDisconnectUser) | **Delete** /api/v1/did-infos/{id}/user | DID情報のクライアントアカウント紐付け解除
[**DidInfosControllerFindAll**](DidInfosApi.md#DidInfosControllerFindAll) | **Get** /api/v1/did-infos | DID情報一覧取得
[**DidInfosControllerFindOne**](DidInfosApi.md#DidInfosControllerFindOne) | **Get** /api/v1/did-infos/{id} | DID情報詳細取得
[**DidInfosControllerRegister**](DidInfosApi.md#DidInfosControllerRegister) | **Post** /api/v1/did-infos/register | 既存DID登録
[**DidInfosControllerRemove**](DidInfosApi.md#DidInfosControllerRemove) | **Delete** /api/v1/did-infos/{id} | DID情報削除
[**DidInfosControllerResolve**](DidInfosApi.md#DidInfosControllerResolve) | **Post** /api/v1/did-infos/resolver | DID解決

# **DidInfosControllerAddLabel**
> DidInfoResponseDto DidInfosControllerAddLabel(ctx, body, id)
DID情報へのラベル追加

リクエストパラメータのidで指定された単一のDID情報に対して、任意の管理用ラベルを追加します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddLabelToDidInfoDto**](AddLabelToDidInfoDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**DidInfoResponseDto**](DidInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DidInfosControllerConnectUser**
> DidInfoResponseDto DidInfosControllerConnectUser(ctx, id, userId)
DID情報へのクライアントアカウント紐付け

DID情報の所有者、関係者などの管理者としてクライアントアカウントを紐付けます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**DidInfoResponseDto**](DidInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DidInfosControllerCreate**
> DidInfosControllerCreate(ctx, body)
新規DID生成

新規DIDを生成します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDidDto**](CreateDidDto.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DidInfosControllerDisconnectUser**
> DidInfosControllerDisconnectUser(ctx, id)
DID情報のクライアントアカウント紐付け解除

DID情報に紐付いたクライアントアカウントの紐付けを解除します。

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
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DidInfosControllerFindAll**
> DidInfosResponseDto DidInfosControllerFindAll(ctx, optional)
DID情報一覧取得

アプリケーションが管理するDID情報を一覧として値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DidInfosApiDidInfosControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DidInfosApiDidInfosControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| (必須) ページネーションを表示する際のページ数に当たるクエリパラメーター。全件取得する際は1を指定。(デフォルト: 1) | 
 **limit** | **optional.Float64**| (必須) ページネーションを表示する際のページ数あたりに表示する件数を指定するクエリパラメーター。全件取得する際は0を指定。(デフォルト: 10) | 
 **did** | **optional.String**| (任意) 例: \&quot;did:key:z6MkhGeGj7u5htkCYjE4PaQ8HUqjYyTmxpDb6Q1MqUpUDsN7\&quot; | 
 **manageUuid** | **optional.String**| (任意) 例: \&quot;32bad62a-4186-4d04-a26a-fcee79d5824b\&quot; | 
 **label** | **optional.String**| (任意) 例: \&quot;did-for-project1\&quot; | 
 **method** | **optional.String**| (任意) 例: \&quot;did:key | 
 **existPrivateKey** | **optional.Bool**| (任意) 例: true | 
 **description** | **optional.String**| (任意)  | 
 **domainName** | **optional.String**| (任意) 例: \&quot;did:web:idroit-dashboard.com\&quot; | 

### Return type

[**DidInfosResponseDto**](DidInfosResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DidInfosControllerFindOne**
> DidInfoResponseDto DidInfosControllerFindOne(ctx, id)
DID情報詳細取得

リクエストパラメータのidで指定された単一のDID情報の詳細情報の値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DidInfoResponseDto**](DidInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DidInfosControllerRegister**
> DidInfosControllerRegister(ctx, body)
既存DID登録

外部で生成されたDIDを本アプリケーションに取り込みます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegisterDidDto**](RegisterDidDto.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DidInfosControllerRemove**
> DidInfosControllerRemove(ctx, id)
DID情報削除

リクエストパラメータのidで指定された単一のDID情報を削除します。

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
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DidInfosControllerResolve**
> ResolveDidResponseDto DidInfosControllerResolve(ctx, body)
DID解決

DIDを解決した結果であるDID Documentの値を返します。このAPIでは保存などの処理を行いません。生成済みのDIDを保存したい場合、既存DID登録API(/did-infos/register)にリクエストを送信してください。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResolveDidDto**](ResolveDidDto.md)|  | 

### Return type

[**ResolveDidResponseDto**](ResolveDidResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

