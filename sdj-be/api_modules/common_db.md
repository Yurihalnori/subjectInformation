## [< 返回API导航](../API.md)
## 公共数据库

公共数据库包含五个子模块,对应科研项目、学位点、导师、学科图书、学位论文、期刊论文
|0|1|2|3|4|5|
| :----: | :----: | :----: |:----: | :----:| :---:|
|科研项目|学位点|导师|学科图书|学位论文|期刊论文|
|Project|Institute|Tutor|Book|Dissertation|Article|

## 重点！ 在所有的url后面加对应module名，代表哪个子模块，如:
+ `POST {base_url}/api/commonDatabase/project`
+ `PUT {base_url}/api/commonDatabase/institute/:id`
+ `DELETE {base_url}/api/commonDatabase/book/:id`

注意 所有的时间形式写成 2023-01-02T14:14:14Z 
### 接口列表

<!-- [用户查看列表信息](#list) `GET {base_url}/api/commonDatabase/getInfo` -->

[管理员添加条目](#add) `POST {base_url}/api/commonDatabase`

[管理员修改条目](#change) `PUT {base_url}/api/commonDatabase/:id`

[管理员删除条目](#delete) `DELETE {base_url}/api/commonDatabase/:id`


<a id="add"></a>

#### 管理员添加条目 `POST {base_url}/api/commonDatabase`
- 要求
  + 每次请求只能对一个子模块进行请求，module名在url最后延伸，如
    + `POST {base_url}/api/commonDatabase/project`
    + `POST {base_url}/api/commonDatabase/dissertation`
  + List列表中字段尽量按照model写全
  + 需要返回字段未定？是否需要补充？
+ Request
```json
  {
        "total":10,
        "list":[{
          "title":"今天吃什么",      //项目名称
          "category":"100101010",    //专业分类  
          ...
      },...]
  }
      
```
+ Response
```json
        {
          "success":true,
          "data":[{
              "check":"", // 查重状态，形式等待进一步确定
              "id":1,   //返回重复的id
              "title":"今天吃什么",      //文章标题  
            },...],
            
        }
```

<a id="change"></a>

#### 管理员修改条目 `PUT {base_url}/api/commonDatabase/:id`

+ 要求
  + 每次请求只能对一个子模块进行请求，module名在url最后延伸，如
    + `PUT {base_url}/api/commonDatabase/project/:id`
    + `PUT {base_url}/api/commonDatabase/dissertation/:id`
  + 每次仅修改一条
  + id在url中给定
  + module在url中给定
  + 字段可任意数目给定
  + 需要返回字段未定？
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
          "id":1,
        }
      }
```

<a id="delete"></a>

#### 管理员删除条目 `DELETE {base_url}/api/commonDatabase/:id`

+ 要求
  + 每次请求只能对一个子模块进行请求，module名在url最后延伸，如
    + `DELETE {base_url}/api/commonDatabase/project/:id`
    + `DELETE {base_url}/api/commonDatabase/dissertation/:id`

+ Request
```json
        {
          "id":2,
        }
```

+ Response 
```json 
        {
          "success":true
        }
```