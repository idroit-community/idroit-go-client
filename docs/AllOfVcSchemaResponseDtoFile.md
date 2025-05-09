# AllOfVcSchemaResponseDtoFile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | (必須) ファイルの識別子 | [default to null]
**Filename** | **string** | (任意) システム内で管理される際にシステムによって名付けられるユニークなファイル名 | [default to null]
**Originalname** | **string** | (必須) システムへアップロード時の元のファイル名 | [default to null]
**Type_** | **string** | (必須) ファイル形式 | [default to null]
**FileData** | **string** | (必須) ファイルのバイナリデータ | [default to null]
**Executed** | **bool** | (任意) ファイルがCSVの場合、アカウントデータ生成実行に使用されたか。(true: 実行済み, false: 未使用) | [default to null]
**Status** | **float64** | (必須) ファイルのステータス | [default to null]
**Group** | [**[]Group**](Group.md) | (任意) ファイルを紐づけたグループの配列 | [default to null]
**VcSchema** | [***Object**](.md) | (任意) ファイルに紐づいたVCスキーマ | [optional] [default to null]
**CreatedBy** | [***Object**](.md) | (必須) ファイルをアップロードしたユーザーアカウント。 | [default to null]
**CreatedAt** | **string** | (必須) ファイルの作成日時 | [default to null]
**UpdatedAt** | **string** | (必須) グループの最終更新日時 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

