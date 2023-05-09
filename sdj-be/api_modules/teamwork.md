## [< 返回API导航](../API.md)

## 团队项目

#### 具体要求为：

1. 无登录状态，可查看简要内容
2. 学生(用户类型1)登录，可查看简要内容、下拉具体内容
3. 老师/管理员(用户类型2)登录，可查看简要内容、下拉具体内容，增删改查数据

### 接口列表

[用户查看列表信息](#list) `GET {base_url}/api/uniqueDatabase/getInfo`

[用户高级搜索](#search) `POST {base_url}/api/uniqueDatabase/search`

[管理员添加条目](#add) `POST {base_url}/api/uniqueDatabase/add`

[管理员修改条目](#change) `POST {base_url}/api/uniqueDatabase/change`

[管理员删除条目](#delete) `DELETE {base_url}/api/uniqueDatabase/delete`

### 接口详情

注意： 未登录只能查看列表，学生可以查看具体信息，管理员可以修改

#### 响应
```json
        {
          "success":"true",
          "data":{
            ......
          }
        }
```

### 错误
```json
        {
          "success":"true",
          "errorMsg":"未找到对应信息"  //错误信息
        }
```

<a id="list"><a>

#### 用户查看列表信息 `GET {base_url}/api/uniqueDatabase/getInfo`

+ 要求

  + 查询参数为页数(page)和每页的限制数量(limit)
  + 返回内容第一部分是全部团队项目在各省市的分布情况
  + 返回内容第二部分是列表信息，形式为一个数组，数组中每个元素包含该条信息的所有信息


+ Request 
  ```json
      // 注意: 这一项的参数不是请求体，而且是url后面的`?page=1&limit=2`部分
      {
        "page": 1, //第几页(从1开始)
        "limit": 10 //一页几条记录
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
              "category":"1,3,8",                      //专业分类 没想好怎么传 可能用二进制传比较方便?(000010110 这种？
              "grade":"1",                        //项目层级 0国际合作 1国家级 2省部级 3市级 4区县级 5乡镇
              "direct":"横向/纵向",               //项目类别（横向/纵向）
              "sponsor":"Trump",                  //资助主体
              "authorization":"123456789",        //项目批准号
              "foundTime":"1145-1-4",             //立项时间
              "superintendent":"Rochester",       //项目负责人
              "member":"a,b,c",                   //项目成员
              "area":{                 //所属区域
                "provinces":"37",      //省
                "city":"07"            //市
              },
              "link":"http://tiaozhan.com",       //调研图片
              "achievement":"balabala"            //项目成果
            },
            ......
          ]
        }
      }
```

<!-- <a id="search"><a>
#### 用户高级搜索 `POST {base_url}/api/uniqueDatabase/search`

> 搜索接口统一？待讨论

+ Request 
```json
      {
        "keyWord":"室温超导",  //搜索框输入内容
        "limitation":[
          {
            "field":"学科分类",  //搜索限定字段
            "msg":"1,3,5"
          },
          ......
        ]
      }
```

+ Response 
```json
        {
          "success":"true",
          "data":{
            "id":"1",
            ......
          }
        }
``` -->

<a id="add"><a>

#### 管理员添加条目 `POST {base_url}/api/uniqueDatabase`

+ Request
```json
{
  "total":10,
  "list":[{
    "id":1,
    "title":"基于氢、氮、镥的室温超导体",      //项目名称
    "category":"1,3,8",                      //专业分类 
    "grade":"1",                        //项目层级 0国际合作 1国家级 2省部级 3市级 4区县级 5乡镇
    "direct":"横向/纵向",               //项目类别（横向/纵向）
    "sponsor":"Trump",                  //资助主体
    "authorization":"123456789",        //项目批准号
    "foundTime":"1145-1-4",             //立项时间
    "superintendent":"Rochester",       //项目负责人
    "member":"a,b,c",                   //项目成员
    "area":{                 //所属区域
      "provinces":"37",      //省
      "city":"07"            //市
    },
    "link":"http://tiaozhan.com",       //调研图片
    "achievement":"balabala"            //项目成果
  },...]
}
      
```
+ Response
```json
        {
          "success":"true",
          "data":[{
            "id":"1",
            "title":"基于氢、氮、镥的室温超导体",      //文章标题
            "updateTime":"2023-02-27"         //上传时间
          },...]
        }
```
<a id="change"><a>

#### 管理员修改条目 `PUT {base_url}/api/uniqueDatabase/:id`

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

#### 管理员删除条目 `DELETE {base_url}/api/uniqueDatabase/:id`
+ 要求
    + 每次仅删除一条
    + id在url中给定
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
