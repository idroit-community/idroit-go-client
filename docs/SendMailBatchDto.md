# SendMailBatchDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | (必須) メールの件名 | [default to null]
**ForEveryClient** | **bool** | (必須) 全てのクライアントアカウントに対して送信するか (true: 全てのクライアントアカウントに送信, false: 全てのクライアントアカウントに送信しない) | [default to null]
**ForEveryUser** | **bool** | (必須) 全てのユーザーアカウントに対して送信するか (true: 全てのユーザーアカウントに送信, false: 全てのユーザーアカウントに送信しない) | [default to null]
**UserIds** | **[]string** | (任意) メールの送信対象のユーザーアカウント識別子の配列 | [optional] [default to null]
**GroupIds** | **[]string** | (任意) 送信対象のグループ識別子の配列。グループを指定した場合グループに紐付くユーザーアカウントとクライアントアカウント全てにメールが送信される | [optional] [default to null]
**Content** | **string** | (必須) メールの内容 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

