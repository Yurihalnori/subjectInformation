# 学科信息外包

<!-- ## API有较大改动，尤其是公共数据库请求相关，如有更好解决办法一定要来call我，我实在无能为力了(哭 -->

~~tenzor也镇不住了，我要跑路力~~
## API

[详情请查看API](./sdj-be/API.md)

## 测试

目前仅支持在本地数据库中测试
base_url为 `127.0.0.1:8080` 

---

## 数据库

数据库sql文件为 `subject.sql` ，端口为 `3306` ,已开启自动迁移

+ 公共数据库
  + 一共七张表，每个都比较抽象，都在model里写了备注

+ 特色数据库
  + 名称`unique_databses`

---

## Postman

详情请导入后查看`subject.postman_collection.json`
格式为collect v2.1

---

## 已实现功能

#### 用户
 + 注册登录登出
#### 特色数据库

+ 管理员添加条目 `POST {base_url}/api/uniqueDatabase`
  + 要求全部字段均required

+ 管理员修改条目 `POST {base_url}/api/uniqueDatabase/:id`
  + 返回信息待进一步确定

+ 管理员删除条目 `DELETE {base_url}/api/uniqueDatabase/:id`

