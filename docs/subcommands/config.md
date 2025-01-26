## `config`

タスクリスト，または tt に関わるグローバル設定を表示・変更します．

### Usage

```sh
tt config [--global] [-e] <name> <value>
```

### Flags

- `--global`
    - ログインユーザのグローバル設定を変更します．
- `-e`
    - テキストエディタを利用して変更します
    - `$EDITOR` を利用します．

### 関連項目

- [`.tt` ディレクトリについて](../structs/tt.md)
