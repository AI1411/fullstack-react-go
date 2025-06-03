# データベースマイグレーション

このディレクトリには、データベーススキーマを管理するためのマイグレーションファイルが含まれています。

## マイグレーションファイル一覧

1. `000001_create_users_table.up.sql` - ユーザーテーブルの作成
2. `000002_create_prefectures_tables.up.sql` - 都道府県と地域のマスターテーブルの作成
3. `000003_create_disasters_table.up.sql` - 災害情報テーブルの作成
4. `000004_create_timeline_table.up.sql` - タイムラインテーブルの作成
5. `000005_create_disaster_documents_table.up.sql` - 災害関連書類テーブルの作成
6. `000006_ccreate_support_applications_table.up.sql` - 支援申請テーブルの作成
7. `000007_create_master_tables.up.sql` - マスターデータテーブルの作成
8. `000008_create_assessment_and_gis_tables.up.sql` - 査定・GIS関連テーブルの作成
9. `000009_add_foreign_keys.up.sql` - 外部キー制約の追加

## 作成したテーブルと機能要件の対応

### 1. 基盤機能

#### 1.1 認証・権限管理
- `users` - ユーザー情報の管理
- `roles` - ユーザーの役割（権限）の管理
- `user_roles` - ユーザーと役割の関連付け
- `organizations` - 組織情報の管理
- `user_organizations` - ユーザーと組織の関連付け

#### 1.2 システム管理
- `disaster_types` - 災害種別マスタ
- `regions`, `prefectures` - 地域・行政区画マスタ
- `facility_types` - 施設種別マスタ
- `damage_levels` - 被害程度マスタ

### 2. 被害情報管理機能

#### 2.1 被害報告登録
- `disasters` - 災害の基本情報（発生日時、場所、種別、緊急度など）
- `disaster_documents` - 災害に関連する写真、動画、ドキュメントなどの添付資料

#### 2.2 被害情報データベース
- `disasters` - 被害情報の一元管理
- 検索・絞り込み機能はアプリケーションロジックで実装

#### 2.3 被害状況分析
- `disasters` - 被害状況の集計・分析用データ
- 統計分析、可視化、レポート機能はアプリケーションロジックで実装

### 3. リモート査定機能

#### 3.1 オンライン査定システム
- `assessments` - 査定情報の管理
- `assessment_items` - 査定項目の詳細管理
- `assessment_comments` - 査定に関するコメント・コミュニケーション

#### 3.2 ドローン・IoT連携
- `disaster_documents` - ドローン映像や画像の保存
- `gis_data` - 地理空間情報の管理

#### 3.3 査定結果管理
- `assessments` - 査定結果の記録・保存、承認ワークフロー管理

### 4. 申請・手続き管理機能

#### 4.1 申請書類作成
- `support_applications` - 支援申請情報の管理
- `disaster_documents` - 申請に関連する書類の管理

#### 4.2 申請進捗管理
- `support_applications` - 申請の進捗状況管理
- `notifications` - リマインダーや通知の管理

#### 4.3 審査・承認機能
- `support_applications` - 審査状況、承認情報の管理

### 5. 帳票・出力機能
- 既存のテーブルからデータを取得して帳票を生成（アプリケーションロジックで実装）

### 6. 地図・GIS機能
- `gis_data` - 地理空間情報の管理（被害分布、影響範囲、避難経路など）

### 7. 通知・コミュニケーション機能
- `notifications` - 通知の管理
- `assessment_comments` - コミュニケーション機能の一部

### 8. モバイル対応機能
- 既存のテーブルを使用してモバイルアプリで表示（アプリケーションロジックで実装）

### 9. 外部連携機能
- 既存のテーブルからデータをAPI経由で提供（アプリケーションロジックで実装）

### 10. セキュリティ・運用機能
- `users`, `roles`, `user_roles` - アクセス制御の基盤
- バックアップ、監視などはインフラストラクチャレベルで実装
