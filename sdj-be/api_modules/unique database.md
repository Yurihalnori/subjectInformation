## [< 返回API导航](../API.md)
## 特色数据库

#### 具体要求为：

1. 无登录状态，可查看简要内容
2. 学生(用户类型1)登录，可查看简要内容、下拉具体内容
3. 老师/管理员(用户类型2)登录，可查看简要内容、下拉具体内容，增删改查数据

### 接口列表

[用户查看列表信息](#list) `GET {base_url}/api/uniqueDatabase/getInfo`

[用户高级搜索](#search) `POST {base_url}/api/uniqueDatabase/search`

[管理员添加条目](#add) `POST {base_url}/api/uniqueDatabase/add`

[管理员修改条目](#change) `PUT {base_url}/api/uniqueDatabase/:id`

[管理员删除条目](#delete) `DELETE {base_url}/api/uniqueDatabase/:id`

### 接口详情

注意： 未登录只能查看列表，学生可以查看具体信息，管理员可以修改

<!-- #### 响应
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
``` -->

<a id="list"><a>

#### 用户查看列表信息 `GET {base_url}/api/uniqueDatabase/getInfo`

+ 要求

  + 查询参数为页数(page)和每页的限制数量(limit)
  + 返回内容第一部分是九种学科总数以及每种的个数，返回形式为一个整数和一个数组
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
          "total":100, //总数量
          "categoryNumber":[1,2,3,4,5,6,7,8,9],//每个学科的条目数量
          "list":[
            {
              "id":1,
              "title":"区块链助力平台发展",      //文章标题
              "author":"于XX",                  //整理者
              "category":"1,3,8",               //学科分类 没想好怎么传 可能用二进制传比较方便?(000010110 这种？
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

<!-- <a id="search"><a>

#### 用户高级搜索 `POST {base_url}/api/uniqueDatabase/search`
> 搜索接口统一？待讨论
+ Request 
```json
      {
        "keyWord":"区块链",  //搜索框输入内容
        "limitation":[
          {
            "field":"学科分类",
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
  "list":[
    {
        "title":"区块链助力平台发展",      //文章标题
        "author":"于XX",                  //整理者
        "category":"1,3,8",               //学科分类
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
          "success":"true",
          "data":[{
            "id":"1",
            "title":"区块链助力平台发展",      //文章标题
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
        "success":"true",
        "data":{
          ......
        }
      }
```

#### 管理员删除条目 `DELETE {base_url}/api/uniqueDatabase/:id`
+ 要求
    + 每次仅修改一条
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
