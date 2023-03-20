# lambda-go-template

このリポジトリは、AWS Lambdaを使用したGolangの基本的なテンプレートプロジェクトです。このテンプレートでは、Serverless Frameworkを使用してデプロイし、AWS DynamoDBをバックエンドとしてデータを保存しています。

## 前提条件

このプロジェクトを実行するには、以下が必要です。

- Go 1.x
- Node.js 12.x以上
- Serverless Framework 3.x
- AWS CLI 2.x

## セットアップ

1. リポジトリをクローンしてください。

    ```bash Copy code
    git clone https://github.com/yourusername/lambda-go-template.git
    cd lambda-go-template
    ```

2. 依存関係をインストールします。

    ```bash Copy code
    go mod download
    npm install
    ```

3. `.env.example`ファイルを .env にリネームし、AWSリージョン、DynamoDBテーブル名、およびDynamoDBテーブルARNを設定します。

    ```makefile Copy code
    MY_AWS_REGION=<your-aws-region>
    DYNAMODB_TABLE=<your-dynamodb-table-name>
    DYNAMODB_TABLE_ARN=<your-dynamodb-table-arn>
    ```

## ローカルでの実行

1. ローカルでDynamoDBとLambdaを起動します。

    ```bash Copy code
    make local
    ```

2. `http://localhost:3000/item`に対してGET、POST、PUT、DELETEリクエストを行い、APIの動作を確認します。

## デプロイ

1. アプリケーションをビルドします。

    ```bash Copy code
    make build
    ```

2. アプリケーションをデプロイします。

    ```bash Copy code
    make deploy
    ```

デプロイが完了したら、生成されたエンドポイントURLを使用してAPIをテストできます。

クリーンアップ
アプリケーションを削除するには、以下のコマンドを実行します。

```bash Copy code
sls remove
```

これで、AWS Lambda関数とDynamoDBテーブルが削除されます。
