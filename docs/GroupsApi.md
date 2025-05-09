# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsControllerConnectDidInfos**](GroupsApi.md#GroupsControllerConnectDidInfos) | **Put** /api/v1/groups/{id}/did-infos | グループへのDID情報紐付け
[**GroupsControllerConnectUsers**](GroupsApi.md#GroupsControllerConnectUsers) | **Put** /api/v1/groups/{id}/users | グループへのユーザーアカウント紐付け
[**GroupsControllerConnectVcInfos**](GroupsApi.md#GroupsControllerConnectVcInfos) | **Put** /api/v1/groups/{id}/vc-infos | グループへのVC情報紐付け
[**GroupsControllerConnectVcSchema**](GroupsApi.md#GroupsControllerConnectVcSchema) | **Put** /api/v1/groups/{id}/vc-schemas | グループへのVCスキーマ紐付け
[**GroupsControllerConnectVpInfos**](GroupsApi.md#GroupsControllerConnectVpInfos) | **Put** /api/v1/groups/{id}/vp-infos | グループへのVP情報紐付け
[**GroupsControllerCreate**](GroupsApi.md#GroupsControllerCreate) | **Post** /api/v1/groups | グループ作成
[**GroupsControllerDisconnectDidInfo**](GroupsApi.md#GroupsControllerDisconnectDidInfo) | **Delete** /api/v1/groups/{id}/did-info/{did_info_id} | グループのDID情報紐付け解除
[**GroupsControllerDisconnectUser**](GroupsApi.md#GroupsControllerDisconnectUser) | **Delete** /api/v1/groups/{id}/user/{user_id} | グループのユーザーアカウント紐付け解除
[**GroupsControllerDisconnectVcInfo**](GroupsApi.md#GroupsControllerDisconnectVcInfo) | **Delete** /api/v1/groups/{id}/vc-info/{vc_info_id} | グループのVC情報紐付け解除
[**GroupsControllerDisconnectVcSchema**](GroupsApi.md#GroupsControllerDisconnectVcSchema) | **Delete** /api/v1/groups/{id}/vc-schema/{vc_schema_id} | グループのVCスキーマ紐付け解除
[**GroupsControllerDisconnectVpInfo**](GroupsApi.md#GroupsControllerDisconnectVpInfo) | **Delete** /api/v1/groups/{id}/vp-info/{vp_info_id} | グループのVP情報紐付け解除
[**GroupsControllerFindAll**](GroupsApi.md#GroupsControllerFindAll) | **Get** /api/v1/groups | グループ一覧取得
[**GroupsControllerFindOne**](GroupsApi.md#GroupsControllerFindOne) | **Get** /api/v1/groups/{id} | グループ詳細取得
[**GroupsControllerRemove**](GroupsApi.md#GroupsControllerRemove) | **Delete** /api/v1/groups/{id} | グループ停止
[**GroupsControllerUpdate**](GroupsApi.md#GroupsControllerUpdate) | **Put** /api/v1/groups/{id} | グループ更新

# **GroupsControllerConnectDidInfos**
> GroupResponseDto GroupsControllerConnectDidInfos(ctx, body, id)
グループへのDID情報紐付け

グループにDID情報を紐付けます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectDidInfosDto**](ConnectDidInfosDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerConnectUsers**
> GroupResponseDto GroupsControllerConnectUsers(ctx, body, id)
グループへのユーザーアカウント紐付け

グループの所有者、関係者などの管理者としてユーザーアカウントを紐付けます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectUsersDto**](ConnectUsersDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerConnectVcInfos**
> GroupResponseDto GroupsControllerConnectVcInfos(ctx, body, id)
グループへのVC情報紐付け

グループにVC情報を紐付けます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectVcInfosDto**](ConnectVcInfosDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerConnectVcSchema**
> GroupResponseDto GroupsControllerConnectVcSchema(ctx, body, id)
グループへのVCスキーマ紐付け

グループにVCスキーマを紐付けます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectVcSchemasDto**](ConnectVcSchemasDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerConnectVpInfos**
> GroupResponseDto GroupsControllerConnectVpInfos(ctx, body, id)
グループへのVP情報紐付け

グループにVP情報を紐付けます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectVpInfosDto**](ConnectVpInfosDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerCreate**
> GroupResponseDto GroupsControllerCreate(ctx, body)
グループ作成

グループ作成に成功しました。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateGroupDto**](CreateGroupDto.md)|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerDisconnectDidInfo**
> GroupResponseDto GroupsControllerDisconnectDidInfo(ctx, id, didInfoId)
グループのDID情報紐付け解除

グループに紐付いたDID情報の紐付けを解除します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **didInfoId** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerDisconnectUser**
> GroupResponseDto GroupsControllerDisconnectUser(ctx, id, userId)
グループのユーザーアカウント紐付け解除

グループに紐付いたユーザーアカウントの紐付けを解除します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerDisconnectVcInfo**
> GroupResponseDto GroupsControllerDisconnectVcInfo(ctx, id, vcInfoId)
グループのVC情報紐付け解除

グループに紐付いたVC情報の紐付けを解除します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **vcInfoId** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerDisconnectVcSchema**
> GroupResponseDto GroupsControllerDisconnectVcSchema(ctx, id, vcSchemaId)
グループのVCスキーマ紐付け解除

グループに紐付いたVCスキーマの紐付けを解除します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **vcSchemaId** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerDisconnectVpInfo**
> GroupResponseDto GroupsControllerDisconnectVpInfo(ctx, id, vpInfoId)
グループのVP情報紐付け解除

グループに紐付いたVP情報の紐付けを解除します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **vpInfoId** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerFindAll**
> GroupsResponseDto GroupsControllerFindAll(ctx, optional)
グループ一覧取得

グループを一覧として値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| (必須) ページネーションを表示する際のページ数に当たるクエリパラメーター。全件取得する際は1を指定。(デフォルト: 1) | 
 **limit** | **optional.Float64**| (必須) ページネーションを表示する際のページ数あたりに表示する件数を指定するクエリパラメーター。全件取得する際は0を指定。(デフォルト: 10) | 
 **name** | **optional.String**| (任意)グループの名前。 例: \&quot;Group Project1\&quot; | 
 **status** | **optional.String**| (任意)グループのステータス。 例: \&quot;active\&quot; | 

### Return type

[**GroupsResponseDto**](GroupsResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerFindOne**
> GroupResponseDto GroupsControllerFindOne(ctx, id)
グループ詳細取得

リクエストパラメータのidで指定された単一の管理者アカウントの詳細情報の値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsControllerRemove**
> GroupsControllerRemove(ctx, id)
グループ停止

リクエストパラメータのidで指定された単一のグループを停止します。

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

# **GroupsControllerUpdate**
> GroupResponseDto GroupsControllerUpdate(ctx, id)
グループ更新

リクエストパラメータのidで指定された単一のグループ情報を更新します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**GroupResponseDto**](GroupResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

