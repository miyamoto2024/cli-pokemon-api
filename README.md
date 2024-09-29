### 必須
- go

### 利用方法
例: go run main.go get --id 10


### ディレクトリ構成
```
.
├── aws
│   └── cloudformation
│       └── rds
├── cmd
│   ├── get.go
│   └── root.go
├── controller
│   └── pokemon.go
├── di
│   └── pokemon.go
├── domain
│   ├── model
│   │   └── pokemon.go
│   └── repository
│       └── pokemon.go
├── driver
│   └── db.go
├── infrastructure
│   └── pokemon_rds.go
└── usecase
    └── pokemon.go
```

### aws
- AWSリソースを作成するための設定。
- `cloudformation`: AWS CloudFormationテンプレートを管理します。
- `rds`: RDSインスタンスの設定を含みます。

### cmd
- CLIコマンドのエントリーポイント。
- `get.go`: ポケモンデータを取得するためのコマンドを定義します。
- `root.go`: CLIアプリケーションのルートコマンドを定義します。

### controller
- HTTPリクエストとレスポンスを処理するロジックを含みます。
- `pokemon.go`: ポケモン関連のエンドポイントを管理します。

### di
- 依存性注入の設定。
- `pokemon.go`: ポケモン関連のコンポーネントの依存関係を設定します。

### domain
- コアビジネスロジックとエンティティ。
- `model/pokemon.go`: ポケモンエンティティを定義します。
- `repository/pokemon.go`: ポケモンデータの永続化のためのインターフェース。

### driver
- データベースドライバーと接続管理。
- `db.go`: データベースの接続と設定を管理します。

### infrastructure
- 外部サービスとデータソースの実装。
- `pokemon_rds.go`: ポケモンデータのためのRDSデータベースとのやり取りを処理します。

### usecase
- アプリケーション固有のビジネスルール。
- `pokemon.go`: ポケモンデータを管理するためのユースケース。