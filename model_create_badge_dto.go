/*
 * iDroit Dashboard Admin API
 *
 *      これは[株式会社フォアー](https://www.fore-co.ltd/ja/)が開発するを使ったDecentralizd Identifiers / Verifiable Credentials(DID/VC)に関係する機能を簡単に利用するための REST API です。現在以下のユースケースをサポートしています。これは今後も拡張されていきます。     - DIDの生成:      - グループ管理機能       - (企業/プロジェクトのまとまり)ごとにユーザー、クライアント、証明書(VC)スキーマを紐付けて管理する。          詳細は以下を参照してください。     - [idroit dashboard admin UI](https://admin.idroit-dashboard.com)     - [idroit dashboard公式ホームページ]()     - [idroit dashboard操作マニュアル]()      以下は関連リンクです。     - [Universal Resolver](https://dev.uniresolver.io/)     - [W3C DID Core 1.0](https://www.w3.org/TR/did-core/)     - [Verifiable Credentials Data Model v2.0](https://www.w3.org/TR/vc-data-model-2.0/)   
 *
 * API version: 0.9.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateBadgeDto struct {
	// (任意) バッジのタイトル
	Label string `json:"label,omitempty"`
	// (任意) グループの作成目的や用途など任意で設定可能な説明文
	Description string `json:"description,omitempty"`
	// (必須) VCの発行者の識別子として用いる文字列の値。現在はDIDのみがサポートされていますが、今後のアップデートでDID以外の文字列をサポートする予定です。
	Issuer string `json:"issuer"`
	// (必須) VCの主張内容(クレーム)となる値のオブジェクト型の値
	CredentialSubject *interface{} `json:"credentialSubject"`
	// (必須) 新規発行するVCの元となるVCスキーマの識別子
	VcSchemaId string `json:"vcSchemaId"`
}
