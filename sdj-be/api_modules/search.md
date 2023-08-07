 ## [< 返回API导航](../API.md)

## 检索

### 获取检索字段 `GET {base_url}/api/advsearch`
- 要求：使用query传递
- Request
```json
{
  "database":1|11|12|13|14|5|6|7|8,
  //1：公共数据库 11：公共数据库-项目12：公共数据库-图书
  //13：公共数据库-学位论文 14：公共数据库-期刊论文
  //5：学位点 6：特色数据库 7：团队项目 8：全部
}
```
- Response
```json
{
 "searchKeyWords":["title","author","organization",...]
}
```
### 进行高级检索 `PUT {base_url}/api/advsearch`
- 要求：使用body传递
- Request
    ```json
    {
        //均是可选的，某项目未作限定返回全部相关信息
        "category":"1100110011", //传入10个学科信息，不填写默认为1111111111
        "database":1|11|12|13|14|5|6|7|8,
        //1：公共数据库 11：公共数据库-项目12：公共数据库-图书
        //13：公共数据库-学位论文 14：公共数据库-期刊论文
        //5：学位点 6：特色数据库 7：团队项目 8：全部
        "quantity": 3 // keyword的数目
        "keyword": [["title","AND","lk-99"],["author","OR","Kin"],
        ["superintendent","OR","linfuchun"]],
        "key":"relation|time|popularity", 
        "order":"increase|reverse",
        "page": 1,
        "limit":5
    }
  ```
- Response
```json
    {
        "success":true,
        "data":{
            "total":10
            "list":[
                {
                    "id": 1, //此id对应其特定database中的id(可能会重复)
                    "title":"题目",
                    "tablename":"projects"//来自的database
                    "region":"domestic|foreign",
                    "department":"挑战网",
                    "date":"2023-04-03",
                },...
            ]
        }
    }
```

