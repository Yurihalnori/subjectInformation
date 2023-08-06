# 学科信息外包
News & Database of Journalism
> API文档导航，临时版



## [学界资讯](./api_modules/news.md)
## [公共数据库](./api_modules/common_db.md)
## [特色数据库](./api_modules/unique_db.md) 
## [团队项目](./api_modules/teamwork.md)
## [高级搜索接口](./api_modules/search.md)
## [用户管理](./api_modules/user.md)
## [文件管理](./api_modules/file.md)

## 返回示例
### 正确
```json
        {
          "success":true,
          "data":{
            ......
          }
        }
```

### 错误
```json
        {
          "success":false,
          "message":"未找到对应信息",  //错误信息
          "code":0  // 1: "内部错误",
                    // 2: "公开错误",
                    // 3: "参数错误",
                    // 4: "系统错误",
                    // 5: "操作错误",
                    // 6: "鉴权错误",
                    // 7: "权限错误",
        }
```