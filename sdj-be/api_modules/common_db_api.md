## [< 返回API导航](../API.md)
## 公共数据库

公共数据库包含五个子模块：包括科研项目、学位点、学科图书、学位论文、期刊论文

以五位二进制00000表示查询的模块，1表示包括 ，0 表示不包括

默认的检索是不包括学位点的跨模块检索即 10111

当检索学位点时只检索学位点 即01000

后端返回时用一位数字（0|1|2|3|4）表示该条目的所属模块

|0|1|2|3|4|
| :----: | :----: | :----: |:----: | :----:|
|科研项目|学位点|学科图书|学位论文|期刊论文|

排序方式目前用1位数字表示 
|0|1|2|3|
| :----: | :----: | :----: |:----: |
|相关度| 时间倒序| 时间正序| 点击量| 

#### 进行检索 `PUT {base_url}/api/search`
> 搜索接口统一？待讨论
- 要求：使用body传递
- Request
    ```json
    {
        //均是可选的，某项目未作限定返回全部相关信息
        "module":"10111", // 查询的模块，默认为10111
        "category":"110011001", //传入9个学科信息，不填写默认为111111111
        "title":"学科信息网站搭建-后端 ",//text选项可传入高级检索字符串
        "database":"0010",//按照position从右至左传递,0010表示仅包含position为1的数据
        "rank":0, // 默认相关度排序 (0)
        "page":1,
        "limit":20,// 默认20
        ...
    }
    ```
- Response
    ```json
    {
        "success":true,
        "data":{
            "category":"110011001",
            "visual": "",//此处会返回用于可视化的数据，形式等待确定
            
            //包含公用字段和归属信息 即对于不同模块的数据，具体的返回属性可能不同，但是可以确保返回某些属性，例如id title database module date writer click以作展示 （有可能为空
            "total":10086,
            "list":[
                {
                    "id": 1,
                    "title":"题目",
                    "category":"100000000",
                    "database":1,
                    "module": 0,
                    "region":"domestic|foreign",
                    "department":"挑战网",
                    "date":"2023-04-03",
                    "click":114,
                    "link": "file:///",
                    "download":514,
                },
                { // 特殊的，对于学位点，这里会返回所有属于该学位点的导师信息
                  "id":1, 
                  "module":1,
                  "location": "xian",
                  "type": 1, // 学硕 专硕 博士
                  "list":[
                    {
                      "id":1,
                      "teachername":"xinzhou",
                      "category":"100000000",
                      "region":0,//国内 国外
                    },...
                  ]
                },...
            ]
        }
    }
    ```


#### 管理员添加条目 `POST {base_url}/api/commonDatabase`
- 要求
  + id在url中给定
+ Request
```json
  {
        "total":10,
        "list":[{
          "title":"今天吃什么",      //项目名称
          "category":"100101010",   //专业分类  
          ...
      },...]
  }
      
```
+ Response
```json
        {
          "success":"true",
          "data":[{
              "check":"", // 查重状态，形式等待进一步确定
              "id":"1",
              "title":"今天吃什么",      //文章标题  
            },...],
            
        }
```

#### 管理员修改条目 `PUT {base_url}/api/commonDatabase/:id`

+ 要求
    + 每次仅修改一条
    + id在url中给定
+ Request
```json
      {
        "id":1,
        字段同上
      }
```
+ Response 
```json
      {
        "success":"true",
        "data":{
          ......
        }
      }
```

#### 管理员删除条目 `DELETE {base_url}/api/commonDatabase/:id`

+ Request
```json
        {
          "id":"1"
        }
```

+ Response 
```json 
        {
          "success":"true"
        }
```