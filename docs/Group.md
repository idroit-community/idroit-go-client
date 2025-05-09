# Group

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | (必須) グループの識別子 | [default to null]
**Name** | **string** | (必須) グループの名前。(例: プロジェクト名や企業名など) | [default to null]
**Status** | **string** | (必須) グループの状態(1: アクティブ, 0: 停止中) | [default to null]
**Description** | **string** | (任意) グループの作成目的や用途など任意で設定可能な説明文 | [default to null]
**Users** | [**[]User**](User.md) | (任意) グループに紐付いたアカウントの配列 | [optional] [default to null]
**CreatedBy** | [***AllOfGroupCreatedBy**](AllOfGroupCreatedBy.md) | (必須) グループを作成したユーザーアカウント。 | [default to null]
**DidInfos** | [**[]DidInfo**](DidInfo.md) | (任意) グループに紐付いたDID情報の配列 | [optional] [default to null]
**VcInfos** | [**[]VcInfo**](VcInfo.md) | (任意) グループに紐付いたVC情報の配列 | [optional] [default to null]
**VpInfos** | [**[]VpInfo**](VpInfo.md) | (任意) グループに紐付いたVP情報の配列 | [optional] [default to null]
**VcSchemas** | [**[]VcSchema**](VcSchema.md) | (任意) グループに紐付いたVCスキーマの配列 | [optional] [default to null]
**CreatedAt** | **string** | (必須) グループの作成日時 | [default to null]
**UpdatedAt** | **string** | (必須) グループの最終更新日時 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

