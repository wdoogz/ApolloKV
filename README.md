### ApolloKV

ApolloKV is a lightweight KV store written in GoLang.

ApolloKV can store Key Value pairs in memory for short term use.


You can test out ApolloKV by doing the following:

```bash
docker pull wdoogz/apollokv
docker run -d -p 8080:10000 wdoogz/apollokv
curl -X POST localhost:8080/write -d '{"ApolloKV":"is awesome"}'
curl -X POST localhost:8080/write -d '{"Doogz":"Rules"}'
curl -X GET localhost:8080/kv
```
