# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VpInfosControllerAddLabel**](VpInfosApi.md#VpInfosControllerAddLabel) | **Put** /api/v1/vp-infos/{id}/label | VP情報へのラベル追加
[**VpInfosControllerConnectUser**](VpInfosApi.md#VpInfosControllerConnectUser) | **Put** /api/v1/vp-infos/{id} | VP情報へのクライアントアカウント紐付け
[**VpInfosControllerCreate**](VpInfosApi.md#VpInfosControllerCreate) | **Post** /api/v1/vp-infos | 新規VP生成
[**VpInfosControllerDisconnectUser**](VpInfosApi.md#VpInfosControllerDisconnectUser) | **Delete** /api/v1/vp-infos/{id}/user | VP情報のクライアントアカウント紐付け解除
[**VpInfosControllerFindAll**](VpInfosApi.md#VpInfosControllerFindAll) | **Get** /api/v1/vp-infos | VP情報一覧取得
[**VpInfosControllerFindOne**](VpInfosApi.md#VpInfosControllerFindOne) | **Get** /api/v1/vp-infos/{id} | VP情報詳細取得
[**VpInfosControllerUpload**](VpInfosApi.md#VpInfosControllerUpload) | **Post** /api/v1/vp-infos/upload | 新規VPアップロード

# **VpInfosControllerAddLabel**
> VpInfoResponseDto VpInfosControllerAddLabel(ctx, body, id)
VP情報へのラベル追加

リクエストパラメータのidで指定された単一のVP情報に対して、任意の管理用ラベルを追加します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddLabelToVpInfoDto**](AddLabelToVpInfoDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**VpInfoResponseDto**](VpInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VpInfosControllerConnectUser**
> VpInfosControllerConnectUser(ctx, body, id)
VP情報へのクライアントアカウント紐付け

VP情報の所有者、関係者などの管理者としてクライアントアカウントを紐付けます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVcInfoDto**](UpdateVcInfoDto.md)|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VpInfosControllerCreate**
> VpInfosControllerCreate(ctx, body)
新規VP生成

新規VPを生成します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVpInfoDto**](CreateVpInfoDto.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VpInfosControllerDisconnectUser**
> VpInfosControllerDisconnectUser(ctx, id)
VP情報のクライアントアカウント紐付け解除

VP情報に紐付いたクライアントアカウントの紐付けを解除します。

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

# **VpInfosControllerFindAll**
> VpInfoResponseDto VpInfosControllerFindAll(ctx, optional)
VP情報一覧取得

アプリケーションが管理するVP情報を一覧として値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VpInfosApiVpInfosControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VpInfosApiVpInfosControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| (必須) ページネーションを表示する際のページ数に当たるクエリパラメーター。全件取得する際は1を指定。(デフォルト: 1) | 
 **limit** | **optional.Float64**| (必須) ページネーションを表示する際のページ数あたりに表示する件数を指定するクエリパラメーター。全件取得する際は0を指定。(デフォルト: 10) | 
 **label** | **optional.String**| (任意) 例: \&quot;vc-for-project1\&quot; | 
 **description** | **optional.String**| (任意)  | 

### Return type

[**VpInfoResponseDto**](VpInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VpInfosControllerFindOne**
> VpInfoDetailDto VpInfosControllerFindOne(ctx, id)
VP情報詳細取得

リクエストパラメータのidで指定された単一のVP情報の詳細情報の値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**VpInfoDetailDto**](VpInfoDetailDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VpInfosControllerUpload**
> VpInfoResponseDto VpInfosControllerUpload(ctx, credentialObject, label, description)
新規VPアップロード

外部で発行された既存VPをアップロードし、本アプリケーションに保存します。。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentialObject** | [**interface{}**](.md)|  | 
  **label** | **string**|  | 
  **description** | **string**|  | 

### Return type

[**VpInfoResponseDto**](VpInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

