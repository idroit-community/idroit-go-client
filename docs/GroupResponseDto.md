# GroupResponseDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | グループの識別子 | [default to null]
**Name** | **string** |  グループの名前 | [default to null]
**Status** | **string** | グループの状態(1: アクティブ, 0: 停止中) | [default to null]
**Description** | **string** | グループの作成目的や用途など任意で設定可能な説明文 | [default to null]
**Users** | [**[]User**](User.md) | グループに紐付いたユーザーアカウントの配列 | [default to null]
**CreatedBy** | [***AllOfGroupResponseDtoCreatedBy**](AllOfGroupResponseDtoCreatedBy.md) | グループを作成したユーザーアカウント | [default to null]
**DidInfos** | [**[]DidInfo**](DidInfo.md) | グループに紐付いたVC情報の配列 | [default to null]
**VcInfos** | [**[]VcInfo**](VcInfo.md) | グループに紐付いたVC情報の配列 | [default to null]
**VpInfos** | [**[]VpInfo**](VpInfo.md) | グループに紐付いたVP情報の配列 | [default to null]
**VcSchemas** | [**[]VcSchema**](VcSchema.md) | グループに紐付いたVCスキーマの配列 | [default to null]
**CreatedAt** | **string** | グループの作成日時 | [default to null]
**UpdatedAt** | **string** | グループの最終更新日時 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

