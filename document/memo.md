# 設計メモ

## 定義

### ロール

|ロール名|説明|
|:--|:--|
|admin|管理者|
|user|ユーザー|

- admin ⊂ user

## 機能

### ユーザ

|機能名|ロール|範囲|説明|エンドポイント候補|
|:--|:--|:--|:--|:--|
|新規登録|!user|本人|ユーザの新規登録|POST /api/v1/user|
|プロフィール閲覧|user|全員|ユーザのプロフィール閲覧|GET /api/v1/user/:id|
|プロフィール編集|user|本人|ユーザのプロフィール編集|POST /api/v1/user/:id|
|退会|user, admin|本人, 全員|ユーザの退会|DELETE /api/v1/user/:id|
|ユーザ一覧表示|admin|全員|ユーザの一覧表示|GET /api/v1/users|

### 飲用状況

|機能名|ロール|範囲|説明|エンドポイント候補|
|:--|:--|:--|:--|:--|
|コーラ実績報告|user|全員|コーラを飲んだことを登録|POST /api/v1/cola|
|コーラ実績閲覧|user|全員|コーラを飲んだ実績詳細表示|GET /api/v1/cola/:id|
|コーラ実績編集|user|本人|コーラを飲んだ実績編集|POST /api/v1/cola/:id|
|コーラ実績削除|user|本人|コーラを飲んだ実績削除|DELETE /api/v1/cola/:id|
|全コーラ実績|user|全員|コーラを飲んだ実績のサマリ|GET /api/v1/colas|
|ユーザのコーラ実績|user|全員|ユーザごとのコーラを飲んだ実績|GET /api/v1/user/colas|