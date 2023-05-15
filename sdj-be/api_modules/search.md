 ## [< 返回API导航](../API.md)

## 检索
#### 检索字段获取 `GET {base_url}/api/search/:database`
- 要求：使用params传递
- 注意：因为所有数据均含有专业分类信息，在搜索时总可以传入category字段，故获取检索字段时不会传递category字段
- Request
    ```json
    {
        "database":0|1|2|3|4, //单一传入检索数据库，默认包含全部（0代表总库）,不能多选
        //1总公共数据库 11项目 12学位点 返回字段交集
    }
    ```
- Response
    ```json
    {
        "success":true,
        "data":{
            "text":[
                {
                    "key":"title", //为查询时发送请求字段名
                    "display":"主题" //字段展示名称
                },{
                    "key":"author",
                    "display":"作者"
                },{
                    "key":"department",
                    "display":"主体单位"
                }
                ],
            "selector":[
                {
                    "key":"db",
                    "items":[{
                        "position":0,
                        "display":"学界资讯"
                    },{
                        "position":1,
                        "display":"公共数据库"
                    },...],
                },{
                    "key":"region",
                    "items":[{
                        "position":0,
                        "display":"国内"
                    },{
                        "position":1,
                        "display":"国外"
                    },...]
                }
            ]
        }
    }
    ```


#### 进行高级检索 `PUT {base_url}/api/search`
- 要求：使用body传递

- Request
    ```json
    {
        //均是可选的，某项目未作限定返回全部相关信息
        "category":"1100110011", //传入10个学科信息，不填写默认为1111111111
        "title":"学科信息网站搭建-后端 ",//text选项可传入高级检索字符串
        "database":0|1|2|3|4, //单一传入检索数据库，默认包含全部（0代表总库）,不能多选
        //1总公共数据库 11项目 12学位点 返回字段交集
        "db":"0010",//按照position从右至左传递,0010表示仅包含position为1的数据
        "sort":{ //可选默认时间倒序
            "key":"relation|time|popularity", 
            "order":"increase|reverse"
        }
        ...
    }
    ```
- Response
    ```json
    {
        "success":true,
        "data":{
            "category":"110011001",
            //包含公用字段和归属信息
            "list":[
                {
                    "id": 1, //此id对应其特定database中的id(可能会重复)
                    "title":"题目",
                    "category":"100000000",
                    "db":"0001",
                    "module": 0,
                    "region":"domestic|foreign",
                    "department":"挑战网",
                    "date":"2023-04-03",
                },...
            ]
        }
    }
    ```
