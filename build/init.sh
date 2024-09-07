# Create the tables
cd ./databases/migrate/ \
&& goose mysql "${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up \

# Start the server (前のコマンドを実行すると止まってしまう。)
cd ../../ \
&& air -c ./build/.air.toml