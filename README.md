# Satoshi Exchange API

取引所のサンプルです。  
手持ちのJPYとSatoshiを売買します。

## セットアップ

```
cp .env.example .env # 初回だけ
docker compose up --build
```

## 実装項目

- エラーハンドリング
- モデル,コントローラー,サービスなど要素の分解
- トランザクション/ロールバック
