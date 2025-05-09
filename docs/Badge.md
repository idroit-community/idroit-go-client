# Badge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | (必須) バッジの識別子 | [default to null]
**Label** | **string** | (任意) バッジのラベル名 | [optional] [default to null]
**Filename** | **string** | (任意) バッジのファイル名 | [optional] [default to null]
**Description** | **string** | (任意) バッジの作成目的や用途など任意で設定可能な説明文 | [default to null]
**Status** | **float64** | (必須) バッジの状態(1: アクティブ, 0: 停止中) | [default to null]
**VcSchema** | [***AllOfBadgeVcSchema**](AllOfBadgeVcSchema.md) | (任意) バッジに紐付いたVCスキーマ | [optional] [default to null]
**VcInfo** | [***AllOfBadgeVcInfo**](AllOfBadgeVcInfo.md) | (任意) バッジに紐付き、メタデータとなるVC情報 | [optional] [default to null]
**CreatedAt** | **string** | (必須) グループの作成日時 | [default to null]
**UpdatedAt** | **string** | (必須) グループの最終更新日時 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

