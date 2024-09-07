# データベース更新方法(goose を利用)

> [!CAUTION]
> からなず[.env.local](../build/.env.local)ファイル名を.env ファイル名に変更すること

## migration ファイルの生成

```bash
$ make b
root$ cd ./databases/migrate
root$ goose create ./add_table_worker sql
```

> [!TIP]
>
> ### 命名規則
>
> - migration の名前は[操作]_[粒度]_[テーブル名]
> - 操作　= add(追加) delete(削除) edit(編集)
> - 粒度 = table(テーブル) column(カラム)
> - テーブル名

## migration の中身を記載

- 20240810085323_add_table_user.sql ファイルが生成されるため、そのファイルに内容を記載。　記載方法は、以下のサイトを参考にする。

## テーブルを更新する

```bash
goose mysql "${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up
```
