# Mail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | (必須) 送信済みメールの識別子 | [default to null]
**Title** | **string** | (必須) 送信済みメールの件名 | [default to null]
**Content** | **string** | (必須) 送信済みメールの内容 | [default to null]
**ForEveryUser** | **bool** | (必須) 全てのuser権限アカウントに対して送信するか。(true: 全てのuser権限アカウントに送信, false: 全てのuser権限アカウントに送信しない) | [default to null]
**ForEveryClient** | **bool** | (必須) 全てのclient権限アカウントに対して送信するか。(true: 全てのclient権限アカウントに送信, false: 全てのclient権限アカウントに送信しない) | [default to null]
**Users** | [**[]User**](User.md) | (任意) メールの送信先となるユーザーアカウントの配列 | [optional] [default to null]
**Groups** | [**[]Group**](Group.md) | (任意) メールの送信先となるグループの配列 | [optional] [default to null]
**CreatedBy** | [**[]User**](User.md) | (必須) メール送信操作を行ったAdmin権限アカウント | [optional] [default to null]
**CreatedAt** | **string** | (必須) メールの送信日時 | [default to null]
**UpdatedAt** | **string** | (必須) メールの送信日時 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

