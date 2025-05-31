## 🏗️ システム構成

### バックエンド
- **言語**: Go 1.24
- **フレームワーク**: Gin Web Framework
- **データベース**: PostgreSQL 17
- **ORM**: GORM with Code Generation
- **依存関係注入**: Uber FX
- **API文書化**: Swagger

### フロントエンド
- **フレームワーク**: Next.js 15.3 (App Router)
- **言語**: TypeScript
- **スタイリング**: Tailwind CSS v4
- **状態管理**: Zustand
- **データフェッチング**: TanStack Query (React Query)
- **API型生成**: Orval
- **コンポーネント**: Radix UI + shadcn/ui

### インフラ・開発環境
- **コンテナ**: Docker & Docker Compose
- **データベースマイグレーション**: golang-migrate
- **ホットリロード**: Air (Go), Next.js dev server
- **リンター**: golangci-lint, Biome
- **コード整形**: gofumpt, goimports

## 🎯 主要機能

### 1. 基盤機能
- 認証・権限管理
- マスタデータ管理（災害種別、地域、施設種別等）
- ログ・監査機能

### 2. 被害情報管理
- 被害報告登録・検索
- 被害情報データベース一元管理
- 被害状況分析・可視化

### 3. 申請・手続き管理
- 申請書類自動作成
- 申請進捗管理
- 審査・承認ワークフロー

### 4. 帳票・出力機能
- 定型帳票自動生成
- カスタム帳票作成
- 一括処理

## 🚀 環境構築

### 必要な環境
- Docker & Docker Compose
- Node.js 18+ (フロントエンド開発時)
- Go 1.24+ (バックエンド開発時)
- pnpm (推奨)

### 1. リポジトリのクローン
```bash
git clone <repository-url>
cd fullstack-react-go
```

### 2. 環境変数の設定
```bash
cp .env.example .env
# 必要に応じて .env ファイルを編集
```

### 3. Docker環境の起動
```bash
# 全サービスを起動
docker compose up -d

# ログを確認
make logs
```

### 4. データベースのセットアップ
```bash
# マイグレーション実行
make migrate

# テストデータ投入（スキーマ全体）
make exec-schema
```

### 5. アプリケーションへのアクセス
- **フロントエンド**: http://localhost:3000
- **バックエンドAPI**: http://localhost:8080
- **API文書**: http://localhost:8080/docs
- **ヘルスチェック**: http://localhost:8080/health

## 🛠️ 開発コマンド

### バックエンド開発

```bash
# モデル再生成
make generate-models

# Swagger文書更新
make swag

# コード整形
make fmt

# 静的解析
make lint

# セキュリティチェック
make sec

# テスト実行
make test

# テストカバレッジ
make test-coverage

# 全品質チェック
make quality

# CI相当のチェック
make ci
```

### フロントエンド開発

```bash
cd frontend

# 開発サーバー起動
pnpm dev

# API型定義生成（Swagger JSONから）
pnpm generate

# コード整形・リント
pnpm lint

# ビルド
pnpm build
```

### データベース操作

```bash
# マイグレーション作成
make migrate-create

# マイグレーション実行
make migrate

# マイグレーション巻き戻し
make migrate-down

# マイグレーションバージョン確認
make migrate-version
```

## 📁 プロジェクト構造

```
.
├── backend/                 # Goバックエンド
│   ├── cmd/api/            # メインアプリケーション
│   ├── internal/           # プライベートコード
│   │   ├── domain/         # ドメインモデル・クエリ（自動生成）
│   │   ├── handler/        # HTTPハンドラー
│   │   ├── usecase/        # ビジネスロジック
│   │   ├── infra/          # インフラ層
│   │   └── middleware/     # ミドルウェア
│   ├── migrations/         # DBマイグレーション
│   └── docs/               # Swagger文書
├── frontend/               # Next.jsフロントエンド
│   ├── src/
│   │   ├── app/           # App Router
│   │   ├── components/    # Reactコンポーネント
│   │   ├── api/           # API型定義（自動生成）
│   │   └── lib/           # ユーティリティ
│   └── public/            # 静的ファイル
├── docker-compose.yml     # Docker設定
├── Makefile              # 開発タスク
└── README.md             # このファイル
```

## 🔧 開発ワークフロー

### 1. 新機能開発
1. データベーススキーマ変更（必要に応じて）
   ```bash
   make migrate-create
   # SQLファイルを編集
   make migrate
   ```

2. モデル再生成
   ```bash
   make generate-models
   ```

3. バックエンドAPIの実装
   - Handler → UseCase → Repository の順で実装
   - Swaggerアノテーションを追加

4. Swagger文書更新・API型生成
   ```bash
   make swag
   cd frontend && pnpm generate
   ```

5. フロントエンド実装
   - 生成された型・フックを使用してUI実装

### 2. コード品質管理
```bash
# 開発前の準備
make tools

# コミット前チェック
make fmt
make quality
make test
```

### 3. データベース設計

現在のスキーマ:
- **users**: ユーザー管理
- **regions**: 地域マスタ
- **prefectures**: 都道府県マスタ
- **disasters**: 災害情報（メインテーブル）

## 🔍 API エンドポイント

### 主要エンドポイント
- `GET /health` - ヘルスチェック
- `GET /disasters` - 災害一覧取得
- `GET /docs` - Swagger UI
- `GET /docs/swagger.json` - Swagger JSON

## 🧪 テスト

### バックエンドテスト
```bash
# 全テスト実行
make test

# カバレッジ付きテスト
make test-coverage
```

### 型安全性
- バックエンド: GORMによるコード生成で型安全なDB操作
- フロントエンド: OrvalによるSwagger→TypeScript型生成で型安全なAPI呼び出し

## 🚀 デプロイ

### 本番環境準備
1. 環境変数の設定
2. データベースのセットアップ
3. アプリケーションのビルド・起動

```bash
# プロダクションビルド
cd frontend && pnpm build
cd backend && go build -o main ./cmd/api
```

## 🔒 セキュリティ

- データベース接続の暗号化
- CORS設定
- 入力値バリデーション
- SQLインジェクション対策（GORM使用）
- 脆弱性チェック（govulncheck）

## 📖 参考資料

- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [Next.js Documentation](https://nextjs.org/docs)
- [TanStack Query Documentation](https://tanstack.com/query)
- [Tailwind CSS Documentation](https://tailwindcss.com/docs)

## 🤝 貢献

1. フォークしてブランチを作成
2. 変更を実装
3. テストを追加・実行
4. プルリクエストを作成

## 📄 ライセンス

このプロジェクトは [MIT License](LICENSE) の下で公開されています。
