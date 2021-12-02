### ApolloKV

ApolloKV is a lightweight KV store written in GoLang.

ApolloKV can store Key Value pairs in memory for short term use.


You can test out ApolloKV by doing the following:

```bash
docker pull wdoogz/apollokv:main
docker run -d -p 8080:10000 wdoogz/apollokv:main
curl -X POST localhost:8080/write -d '{"ApolloKV":"is awesome"}'
curl -X POST localhost:8080/write -d '{"Doogz":"Rules"}'
curl -X GET localhost:8080/kv
```

To delete a KV you can do the following, your body should just be the key to the key value pair:
```bash
curl -X DELETE localhost:8080/write -d 'Doogz'
```
