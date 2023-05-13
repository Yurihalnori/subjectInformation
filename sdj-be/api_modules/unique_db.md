## [< 返回API导航](../API.md)
## 特色数据库

#### 具体要求为：

每个库包含九个子学科，以九位二进制000000000表示查询的模块，1表示包括 ，0表示不包括

### 接口列表

[用户查看列表信息](#list) `GET {base_url}/api/uniqueDatabase/getInfo`

[管理员添加条目](#add) `POST {base_url}/api/uniqueDatabase`

[管理员修改条目](#change) `PUT {base_url}/api/uniqueDatabase/:id`

[管理员删除条目](#delete) `DELETE {base_url}/api/uniqueDatabase/:id`

### 接口详情

<a id="list"></a>

#### 用户查看列表信息 `GET {base_url}/api/uniqueDatabase/getInfo`

+ 要求
  + page和limit通过Querys传递
  + 返回内容第一部分是九种学科总数以及每种的个数，返回形式为一个整数和一个数组，用于右侧显示(待定)
  + 返回内容第二部分是列表信息，形式为一个数组，数组中每个元素包含该条信息的所有信息

+ Request 
  ```json
      {
        "page": 1, 
        "limit": 10
      }
  ```

+ Response
```json
      {
        "success":true,
        "data":{
          "total":100,
          "categoryNumber":[1,2,3,4,5,6,7,8,9],  //每个学科的条目数量
          "list":[
            {
              "id":1,
              "title":"区块链助力平台发展",      //文章标题
              "author":"于XX",                  //整理者
              "category":"001001001",           //学科分类
              "keyWord":"区块链,银行,商业",      //关键词 可以存在多个
              "introduction":"balabala……",      //简介 点进去可看？
              "updateTime":"2023-02-27",        //上传时间
              "link":"http://tiaozhan.com",     //附件
              "click":"37",                     //点击数
              "download":"12"                   //下载数
            },
            ......
          ]
        }
      }
```

<a id="add"></a>

#### 管理员添加条目 `POST {base_url}/api/uniqueDatabase`

+ Request
```json
{
  "total":10,
  "list":[
    {
        "title":"区块链助力平台发展",      //文章标题
        "author":"于XX",                  //整理者
        "category":"001001001",           //学科分类
        "keyWord":"区块链,银行,商业",      //关键词
        "introduction":"balabala……",      //简介
        "link":"http://tiaozhan.com"      //附件
    },...
  ]
}
      
```
+ Response
```json
        {
          "success":true,
          "data":[{
            "id":1,
            "title":"区块链助力平台发展",      //文章标题
            "updateTime":"2023-02-27"         //上传时间
          },...]
        }
```
<a id="change"></a>

#### 管理员修改条目 `PUT {base_url}/api/uniqueDatabase/:id`

+ 要求
    + 每次仅修改一条
    + id在url中给定

+ Request
```json
      {
        "title":"",           //文章标题
        "author":"",          //整理者
        "category":"",        //学科分类
        "keyWord":"",         //关键词 可以存在多个
        "introduction":"",    //简介 点进去可看？
        "link":"",            //附件
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

<a id="delete"></a>

#### 管理员删除条目 `DELETE {base_url}/api/uniqueDatabase/:id`
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
