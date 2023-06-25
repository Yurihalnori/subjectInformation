## [< 返回API导航](../API.md)

## 学界咨询

#### 获取资讯 `GET {base_url}/api/news`
- 要求：使用query传递

- Request
    ```json
    {
        "module":0|1|2|3, // 0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息, 仅支持单值传递
        "page":1,
        "limit":20,
        "category":"1100110011", //传入10个学科信息，默认包含全部
         //默认时间倒序
        "name":"time|popularity", 
        "order":"increase|reverse"
    }
    ```
- Response
    ```json
    {
        "success":true,
        "data":{
            "category":"110011001",
            "module":0,
            "total": 100
            "list":[
                {
                    "id": 1,
                    "title":"题目",
                    "category":"100000000",
                    "module": 0,
                    "region":"domestic|foreign",
                    "department":"挑战网",
                    "date":"2023-04-03",
                    "click": 114514
                }
            ]
        }
    }
    ```
#### 获取单篇文章具体内容（可能有？）`GET {base_url}/api/news/:id`
- 要求：使用params传递

- Request
    ```json
    {
        "id": 1
    }
    ```
- Response
    ```json
    {
        "success":true,
        "data":{
            "id": 1,
            "title":"题目",
            "category":"100000000",
            "module": 0,
            "region":"domestic|foreign",
            "department":"挑战网",
            "date":"2023-04-03",
            "content":"..." //可能包含富文本
        }
    }
    ```
#### 发表文章 `POST {base_url}/api/news/`
- 要求：使用body传递

- Request
    ```json
    {
        "total":10, //传入list长度
        "list":[
            {
            "title":"题目",
            "category":"100000000",
            "module": 0,
            "region":"domestic|foreign",
            "department":"挑战网",
            "date":"2023-04-03",
            "content":"..." //可能包含富文本
        }
        ,...]
    }
        
    ```
- Response
    ```json
    {
        "success":true,
        "data":[{
            "id": 1
        },...]
    }
    ```

#### 修改一篇文章 `PUT {base_url}/api/news/:id`
- 要求：使用body传递, 注意文章id在params中提供

- Request
    ```json
        {
            "title":"题目", //可选
            "category":"100000000", //可选
            "module": 0, //可选
            "region":"domestic|foreign", //可选
            "department":"挑战网", //可选
            "date":"2023-04-03", //可选
            "content":"..." //可选
        }
    ```
- Response
    ```json
    {
        "success":true,
        "data":{
            "id": 1
        }
    }
    ```

#### 删除一篇文章 `DELETE {base_url}/api/news/:id`
- 要求：使用body传递, 注意文章id在params中提供

- Request
    ```json
        {
            "id":1
        }
    ```
- Response
    ```json
    {
        "success":true,
        "data":{
            "id": 1
        }
    }
    ```
