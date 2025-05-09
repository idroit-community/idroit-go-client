# CreateUserDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | (必須) ユーザーアカウントの名前 | [default to null]
**Email** | **string** | (必須) ユーザーアカウントのメールアドレス | [default to null]
**Password** | **string** | (任意) ユーザーアカウントのパスワード(8文字以上20字未満)。\&quot;needActivateFlow\&quot;が\&quot;false\&quot;の場合は必須。 | [optional] [default to null]
**Role** | **string** | (必須) アカウントのタイプ | [default to null]
**NeedActivateFlow** | **bool** | (任意) アカウントの有効化フローが必要か否か。デフォルト値ではfalse。(true: 必要 /false: 不要) | [optional] [default to null]
**GroupIds** | **[]string** | (任意) ユーザーアカウントに紐付けるグループのIDの配列 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

