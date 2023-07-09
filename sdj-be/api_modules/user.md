## [< 返回API导航](../API.md)
管理员、一般用户
## 一般用户模块
> 具体用户管理模块api还需要讨论
### 用户注册 `POST {base_url}/api/user/register`

- Request
    ```json
    {
        "username":"fusion",
        "password":"adcd1234",
    }
    ```
- Response
    ```json
    {
        "success":true,
        "data":{
            "userId":1,
            "role":0
        }
    }
    ```
### 用户登录 `POST {base_url}/api/user/login`
- Request
    ```json
    {    
        "role": 0,
        "username":"fusion",
        "password":"adcd1234",
    }
    ```
    普通用户选择管理员登录则无法登录，管理员选择普通登录仍为管理员
- Response
    ```json
    {
        "success":true,
        "data":{
            "userId":1,
            "role":0
        }
    }
    ```
### 用户登出 `DELETE {base_url}/api/user/logout`
- Request
   无
- Response
    ```json
    {
        "success":true,
        "data":"登出成功"
    }
    ```

### 用户状态 `GET {base_url}/api/user/:id`
- Request
   无
- Response
    ```json
    {
        "success":true,
        "data":{
            "userId":1,
            "role":0
        }
    }
    ```
## 管理员用户(老师)

### 身份修改 `POST {base_url}/api/admin/change`
- Request
    ```json
    {
        "userId":1,
        "role":0 //1 是管理员
    }
    ```
- Response
    ```json
    {
        "success":true,
        "data":{
            "userId":1,
            "role":0
        }
    }
    ```