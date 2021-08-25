# learn golang

## menu

- [x] golang
    - [x] gin
    - [x] gorm
- [x] func testing
- [ ] api client
    - [ ] line bot
    - [ ] discord bot
- [x] api server
    - [x] restAPI
    - [x] grpc
    - [x] jwt
- [x] database
    - [x] mysql
    - [ ] redis
- [x] http status code
    - [x] 100
    - [x] 200
    - [x] 300
    - [x] 400
    - [x] 500
- [x] http 
    - [x] https tls
    - [x] http1.1 http2(沒人用) http3(quic)
    - [x] web socket
    - [x] CDN
    - [x] DNS

### 08/09

- Https
- TLS
- CDN
- DNS

### 08/10

- refactor code
- code review with @sky
- repository pattern

### 08/11

- migrate 實作
- transation 實作
- 資料庫正規化概念學習

### 08/12
graceful shutdown

### 08/14
gin: download

### 08/15
implement gin/grpc: graceful shutdown
goroutine & context
pub/sub (redis)
cache (redis)

### 8/16
sync.Mutex, sync.RWMutex
db deadlock

### 8/19
N+1 query problem

### 8/21
join & preload

### 8/25
sql index

### homework

哪個客人買最多A商品
- order table (id, purchaser_id, product_id, created_at)
- order_detail (id, order_id, product_id, product_amount)
- product table (id, product_name, product_cost, created_at)
- discount table (id, product_id, percentage, start_at, end_at)

mysql百萬級query沒啥問題
訂單一定成功

sql preload
sql explain
sql partition
Pt online schema change
elastic search (不要懂太多)
水平拆分、垂直拆分

刪除db時也要刪除cache(redis)

#### learn

hashids
notify
sql injection
regex
websocket
dotenv

- frontend
    - React
    - Vue
    - Angular
    - html
    - javascript
    - jQuery(js)
    - Typescript
    - CSS (sass)
    - npm
    - nodejs (express framework)

