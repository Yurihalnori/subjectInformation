## [< 返回API导航](../API.md)
## 文件
||1|2|3|4|5|6|
| :---: | :----: | :----: | :----: |:----: | :----:| :---:|
|数据库|特色数据库|科研项目|学科图书|学科论文|期刊论文|团队项目|
|表名(table)|unique_database|project|book|dissertation|article|teamwork|
|字段名(field)|attachment|text|text|text|text|text|
|||||data|data|picture|
tip:仅供参考，前端可根据需求改变
### 文件上传
+ Request
```json
{
  //以表单form形式上传
  "file":,
  "table":"unique_database",
  "field":"data",
}
```


