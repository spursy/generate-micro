package template

var (
	SvcSRV = `### redis 配置
redis:
  host: tcp:127.0.0.1:6379
  database: 0
  maxOpenConn: 10
  maxIdleConn: 2

### pg 配置
news:
  username: "postgres"
  password: "nbwkNZV2tI8owhsIL3aQ"
  address: "10.0.0.152:5432"
  dbName: "lb-news-dev"
  sslmode: "disable"
  logMode: true

### mysql 配置
search:
  address: "10.0.0.144"
  dbName: "bi_search_support"
  password: "ydoA8mpKhYjc2xlZNQzX"
  username: "root"
`
)



