# sns-app

私は神ではない。

## Local Development

###### Frontend

[http://localhost:3000](http://localhost:3000)

###### Backend

[http://localhost:1323](http://localhost:1323)

###### Database GUI

[http://localhost:8888](http://localhost:8888)

### Requirements

- [Docker](https://www.docker.com)
- [Volta](https://volta.sh)

### Getting Started

#### Setting Up

アプリケーションの準備

```shell
make build
```

アプリケーションの起動

```shell
make up
```

データベースのマイグレーション

```shell
make migrate
```

初期データの投入

```shell
make seed
```

#### Shutting Down

データベースを初期化

```shell
make migrate-down
```

アプリケーションの停止

```shell
make down
```

#### Help

その他のコマンドは以下を実行して参照！

```shell
make help
```
