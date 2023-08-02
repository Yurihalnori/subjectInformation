## [< 返回API导航](../API.md)

## 团队项目

#### 具体要求为：

项目层级 0国际合作 1国家级 2省部级 3市级 4区县级 5乡镇级

### 接口列表



[用户查看列表信息](#list) `GET {base_url}/api/teamwork`

[管理员添加条目](#add) `POST {base_url}/api/teamwork`

[管理员修改条目](#change) `POST {base_url}/api/teamwork/:id`

[管理员删除条目](#delete) `DELETE {base_url}/api/teamwork/:id`

### 接口详情

<a id="list"></a>

#### 用户查看列表信息 `GET {base_url}/api/teamwork`

+ 要求
  + `page`和`limit`和其他字段均通过body传递

+ Request 
```json
{
  "page":1,
  "limit":10,
  "field":"id",     //根据哪个字段进行排序 默认id
  "order":"asc"     //acs升序 decs降序 默认升序
}
  ```

+ Response
```json
{
  "success":true,
  "data":{
    "total":100,     //总数量
    "areaQuantity":[
      {
        "provinces":"37",     
        "city":"07",        //身份证号前4位确定省市
        "number":"2"        //该地市项目数量
      },
      ......
    ],
    "list":[
      {
        "id":1,
        "title":"基于氢、氮、镥的室温超导体",      //项目名称
        "category":"001001001",                  //专业分类 
        "grade":1,                               //项目层级 
        "direct":"横向/纵向",                //项目类别（横向/纵向）
        "sponsor":"Trump",                  //资助主体
        "authorization":"123456789",        //项目批准号
        "foundTime":"1145-1-4",             //立项时间
        "superintendent":"Rochester",       //项目负责人
        "member":"a,b,c",                   //项目成员
        "area":{                 //所属区域
          "provinces":"37",      //省
          "city":"07"            //市
        },
        "picture":"http://tiaozhan.com",       //调研图片
        "achievement":"balabala"            //项目成果
      },
      ......
    ]
  }
}
```

<a id="add"></a>

#### 管理员添加条目 `POST {base_url}/api/teamwork`

+ Request
  ```json
  {
    "total":10,
    "list":[{
      "id":1,
      "name":"基于氢、氮、镥的室温超导体",      //项目名称
      "category":"001001001",        //专业分类 
      "grade":1,                     //项目层级
      "direction":"横向/纵向",        //项目类别（横向/纵向）
      "sponsor":"Trump",             //资助主体
      "number":"123456789",          //项目批准号
      "time":"1145-1-4",             //立项时间
      "principal":"Rochester",       //项目负责人
      "member":"a,b,c",                   //项目成员
      "region":{                 //所属区域
        "provinces":"37",      //省
        "city":"07",            //市
        "county":"04"          //区
      },
      "picture":"http://tiaozhan.com",   //调研图片
      "achievement":"balabala"           //项目成果
    },...]
  }
  ```
  
+ Response
```json
        {
          "success":true,
          "data":[{
            "id":1,
            "title":"基于氢、氮、镥的室温超导体",      //文章标题
            "time":"2023-02-27"                      //上传时间
          },...]
        }
```
<a id="change"></a>

#### 管理员修改条目 `PUT {base_url}/api/teamwork/:id`

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
        "success":true,
        "data":{
          ......
        }
      }
```

#### 管理员删除条目 `DELETE {base_url}/api/teamwork/:id`
+ 要求
    + 每次仅删除一条
    + id在url中给定
+ Request
```json
        {
          "id":1
        }
```

+ Response 
```json 
        {
          "success":true
        }
```
