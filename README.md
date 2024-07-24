## Local Development

<table>
	<tbody>
		<tr>
			<td>FRONEND</td>
			<td><a href="http://localhost:3000">http://localhost:3000</a></td>
		</tr>
		<tr>
			<td>BACKEND</td>
			<td><a href="http://localhost:3001">http://localhost:3001</a></td>
		</tr>
</table>

### Requirements

- [Docker](https://www.docker.com)
- [Volta](https://volta.sh)

### Getting Started

セットアップ

```shell
make setup
```

アプリケーションの起動

```shell
make up
```

初期データの投入

```shell
make seed
```

アプリケーションの停止

```shell
make down
```

### Database

データベース内部を GUI から確認

```shell
make view-db
```

データベースのスキーマ更新

```shell
make migrate name={更新名}
```
