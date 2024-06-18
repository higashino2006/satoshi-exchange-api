# Satoshi Exchange API

取引所のサンプルです。  
手持ちのJPYとSatoshiを売買します。

具体的な挙動の例は(./docs/behavior_videos)参照

## セットアップ

```
cp .env.example .env # 初回だけ
make up
```

## 実装項目

- エラーハンドリング
- モデル,コントローラー,サービスなど要素の分解
  - [src/internal](./src/internal)参照
- トランザクション/ロールバック
  - [src/internal/services/tx_service.go](./src/internal/services/tx_service.go)参照
- 単体テスト
  - [src/internal/tests/unit](./src/internal/tests/unit)参照
- 統合テスト(一部APIのみ)
  - [src/internal/tests/integration](./src/internal/tests/integration)参照

### 単体テスト

TEST_PATHに/tests/unit以下のパスを指定

例
```
make unit_test TEST_PATH=/lib/common_test.go
```

### 統合テスト

TEST_PATHに/tests/unit以下のパスを指定。  
事前にテスト用のdocker composeを立ち上げる必要がある。

例
```
make up_test
make integration_test TEST_PATH=/controllers/trade_controller_test.go
```
