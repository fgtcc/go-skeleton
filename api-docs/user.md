[toc]

### 登录

* 描述：用户登录，返回token
* url：/api/v1/login
* 请求方式：post

请求参数

|  参数名   | 必选 |  类型  |    说明    |
| :-------: | :--: | :----: | :--------: |
| loginName |  是  | string | 用户登录名 |
| password  |  是  | string |  用户密码  |

请求实例

```json
{
    "loginName":"admin",
    "password":"123456"
}
```

返回参数

| 参数名 |  类型  |      说明       |
| :----: | :----: | :-------------: |
| token  | string | 登录令牌（jwt） |

返回实例

```json
{
    "code": 200,
    "result": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbl9uYW1lIjoiYnV6aGkxIiwiZXhwIjoxNjAzMDc1MjM2LCJpc3MiOiJ0YWxsc2FmZSJ9.M7dB2TdfIKBN6zsi8uTfECfa1dKqRBzAPEBfMGecVlU"
    },
    "message": "OK"
}
```

### 添加用户

* 描述：添加用户
* url：/api/v1/user/add
* 请求方式：post

请求参数

|  参数名   | 必选 |  类型  |    说明    |
| :-------: | :--: | :----: | :--------: |
| username  |  是  | string |  用户名称  |
| loginName |  是  | string | 用户登录名 |
| password  |  是  | string |  用户密码  |

请求实例

```json
{
    "username":"用户02",
    "loginName":"user002",
    "password":"123456"
}
```

返回实例

```json
{
    "code": 200,
    "result": null,
    "message": "OK"
}
```

### 获取用户信息

* 描述：获取用户信息
* url：/api/v1/user/getInfo
* 请求方式：post

请求参数

|  参数名   | 必选 |  类型  |    说明    |
| :-------: | :--: | :----: | :--------: |
|    id     |  --  |  int   |   用户id   |
| loginName |  --  | string | 用户登录名 |

> 参数 id 与 loginName二选一

请求实例

```json
{
   "id":30
}
```

返回参数

|  参数名   |  类型  |           说明           |
| :-------: | :----: | :----------------------: |
|    id     |  int   |          用户id          |
| loginName | string |        用户登录名        |
| username  | string |         用户名称         |
|  status   |  int   | 用户状态，0-禁用，1-启用 |

返回实例

```json
{
    "code": 200,
    "result": {
        "id": 30,
        "loginName": "user011",
        "username": "用户011",
        "status": 1
    },
    "message": "OK"
}
```

### 获取用户列表

* 描述：获取用户列表，支持根据用户状态筛选
* url：/api/v1/user/getList
* 请求方式：post

请求参数

|  参数名  | 必选 | 类型 |          说明           |
| :------: | :--: | :--: | :---------------------: |
| status | 否 | int | 用户状态， |
|  pageNo  |  否  | int  |      用户状态，默认值-1表示全部查询      |
| pageSize |  否  | int  | 每页条数， 默认为10 |

请求实例

```json
{
    "pageNo":1,
    "pageSize":6
}
```

返回参数

|   参数名   | 类型 |     说明     |
| :--------: | :--: | :----------: |
| totalCount | int  |     总数     |
| totalPage  | int  |    总页数    |
|   pageNo   | int  |     页码     |
|  pageSize  | int  | 每页条数 |
|    data    | array |   用户列表   |

返回实例

```json
{
    "code": 200,
    "result": {
        "totalCount": 18,
        "totalPage": 3,
        "pageNo": 1,
        "pageSize": 6,
        "data": [
            {
                "id": 1,
                "loginName": "admin",
                "username": "管理员",
                "status": 1
            },
            {
                "id": 13,
                "loginName": "user002",
                "username": "用户0002",
                "status": 1
            },
            {
                "id": 14,
                "loginName": "user003",
                "username": "用户003",
                "status": 1
            },
            {
                "id": 16,
                "loginName": "user005",
                "username": "用户005",
                "status": 1
            },
            {
                "id": 17,
                "loginName": "user006",
                "username": "用户006",
                "status": 1
            },
            {
                "id": 18,
                "loginName": "user007",
                "username": "用户007",
                "status": 1
            }
        ]
    },
    "message": "OK"
}
```

### 编辑用户

* 描述：编辑用户信息
* url：/api/v1/user/update
* 请求方式：post

请求参数

|  参数名  | 必选 |  类型  |      说明      |
| :------: | :--: | :----: | :------------: |
|    id    |  是  |  int   |     用户id     |
| username |  否  | string | 修改后的用户名 |
| password |  否  | string |  修改后的密码  |

请求实例

```json
{
    "id":9,
    "username":"张三"
}
```

返回实例

```json
{
    "code": 200,
    "result": null,
    "message": "OK"
}
```

### 批量删除用户

* 描述：编辑用户信息
* url：/api/v1/user/delete
* 请求方式：post

请求参数

| 参数名 | 必选 | 类型 |    说明    |
| :----: | :--: | :--: | :--------: |
| idList |  是  | array | 用户id列表 |

请求实例

```json
{
    "idList": [4,5]
}
```

返回实例

```json
{
    "code": 200,
    "result": null,
    "message": "OK"
}
```

### 批量更新用户状态

* 描述：批量更新用户状态，实现用户的禁用和启用
* url：/api/v1/user/batchUpdateStatus
* 请求方式：post

请求参数

| 参数名 | 必选 | 类型  |             说明              |
| :----: | :--: | :---: | :---------------------------: |
| idList |  是  | array |          用户id列表           |
| status |  否  |  int  | 状态，0-禁用，1-启用，默认为1 |

请求实例

```json
{
    "idList":[28,29],
    "status":0
}
```

返回实例

```json
{
    "code": 200,
    "result": null,
    "message": "OK"
}
```

