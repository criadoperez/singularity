# データソース再スキャン

デフォルトでは、データソースは一度だけスキャンされます。ただし、Singularityには、一定の間隔ごとにデータソースを再スキャンするオプションが用意されています。

```sh
singularity datasource add <type> --rescan-interval value
```

また、手動で再スキャンをトリガーすることもできます。

```sh
singularity datasource rescan datasource_id
```

再スキャンでは、新しいファイルは準備のためにキューに追加されます。削除されたファイルは無視されます。

#### ファイルのバージョニング

変更されたファイルについては、新しいバージョンのファイルが準備のためにキューに追加され、ディレクトリのCIDは最新バージョンのファイルを使用するように更新されます。

ファイルの変更があるかどうかのロジックは、以下の手順で決定されます：

1. データソースがファイルのハッシュ値（例：Etag）を提供している場合、ハッシュ値が変更された場合に新しいバージョンが作成されます。
2. それ以外の場合、データソースがファイルの最終更新日時を提供している場合、その値が変更されたか、ファイルサイズが変更された場合に新しいバージョンが作成されます。
3. それ以外の場合、ファイルサイズを使用して新しいバージョンが作成されるかどうかが判断されます。

再スキャンの間に同じファイルが複数回上書きされると、一部のファイルバージョンが見落とされる可能性があります。すべてのファイルバージョンがキャプチャされるようにするには、ユーザーは [push-and-upload.md](push-and-upload.md "mention") を使用して、ファイルが更新されるたびにSingularityに通知する必要があります。