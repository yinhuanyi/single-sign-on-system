# sso

### 授权码模式
http://localhost:10541/api/v1/authorize?client_id=client_1&response_type=code&scope=all&state=xyz&redirect_uri=http://localhost:10001/api/v1/goods/get
- 请求sso地址，获取code授权码：http://localhost:10541/api/v1/authorize?client_id=client_2&response_type=code&scope=all&state=xyz&redirect_uri=http://localhost:10002/cb
- 获取到code之后，请求应用A：http://localhost:10001/cb?code=MWY3ZDC3YTGTM2Q5ZI0ZYMI5LWJHMJUTNDLHZGJJMTIZNZKX&state=xyz
- 应用A获取code之后，再POST请求sso：http://localhost:10541/api/v1/token (需要basic认证，提供用户名和密码)， 
     post请求参数：grant_type：authorization_code，code：MTLKZTFMODKTZGFKOC0ZMZNJLWEXZMETY2E1MGI4NJZJMTQX，redirect_uri：http://localhost:10001/cb
     获取的返回：此时的数据是再
```
  {
  "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjbGllbnRfMSIsImV4cCI6MTY0MzE3MTYwOSwic3ViIjoiMSJ9.CaA74487Wetiemk_mfQ5vDs3y2DU3sr1pvPNYVsDu7JrdQkuv15FevOr0C4TW_PZUXAznuLv9tywodNGLqte5g",
  "expires_in": 7200,
  "refresh_token": "MTA0NDYXNJGTY2EWYS01Y2Q0LWI4YJUTOGI2NTCXM2MXMTNH",
  "scope": "all",
  "token_type": "Bearer"
  }
```


### implicit
- 请求地址：http://localhost:10541/api/v1/authorize?client_id=client_1&response_type=token&scope=all&state=xyz&redirect_uri=http://localhost:10001/cb

### 密码模式
- 请求