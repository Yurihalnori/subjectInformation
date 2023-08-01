## [< 返回API导航](../API.md)

## 学界咨询

Note: module 和 region 永远用int 
```json
"module":0|1|2|3|4, // 0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息,4:全选  仅支持单值传递
"region":0|1 //domestic|foreign
```
#### 获取资讯 `GET {base_url}/api/news`
- 要求：使用query传递

- Request
    ```json
    {
        "module":0|1|2|3|4, // 0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息, 4:全选 仅支持单值传递
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
                    "category":"1000010000",
                    "module": 0,
                    "region":1,
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
            "category":"1000010000",
            "module": 0,
            "region":1,
            "department":"挑战网",
            "date":"2023-04-03",
            "content":"..." //可能包含富文本
        }
    }
    ```
#### 发表文章 `POST {base_url}/api/news/`
- 要求：使用body传递
- 注意：因为会有多篇资讯出现在一天之内，所以date字段请传一个具体时间”2023-1-1T04:05:14Z“
- Request
    ```json
    {
        "total":10, //传入list长度
        "list":[
            {
            "title":"题目",
            "category":"1000010000",
            "module": 0,
            "region":1,
            "department":"挑战网",
            "date":"2006-01-01T04:05:14Z",
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
  注意：因为会有多篇资讯出现在一天之内，所以date字段请传一个具体时间”2023-1-1T04:05:14Z“
- Request
    ```json
        {
            "title":"题目", //可选
            "category":"100000000", //可选
            "module": 0, //可选
            "region":1, //可选
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

#### 搜索 `POST {base_url}/api/news/search`
- 要求：使用body传递
- Request
    ```json
    {
        "content": "是学生",//建议前端直接过滤掉非法字符
        "module":0|1|2|3|4 // 0:行业资讯，1：学术会议，2：学科竞赛，3：招聘信息, 4:全选 仅支持单值传递
        "page":1,
        "limit":20,
        "category":"1100110011", //传入10个学科信息，默认包含全部
         //默认时间倒序
        "name":"relativity|time|popularity",   
        "order":"increase|reverse" //不提供相关性升序
    }
  ```
- Response
  ```json
  {
    "data": {
        "category": "1111001000",
        "list": [
            {
                "id": 18,
                "title": "2023年9月全国计算机等级考试报名通知",
                "module": 3,
                "department": "挑战网",
                "click": 0,
                "date": "2023-06-05",
                "region": 2,
                "category": "1010100011"
            }
        ],
        "module": 3,
        "total": 3
    },
    "success": true
  }
  ```