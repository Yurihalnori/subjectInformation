# 学科信息外包


## 快速预览

```
go run .
```

## 数据库

数据库sql文件为 `subject.sql` ，请设置环境变量以指定数据库

## 环境变量

环境变量名及其默认值在 `util/config.go` 中定义

- 项目实际上线时， `APP_PROD` 应设置为任意非空字符串，以开启生产模式
- 项目实际上线时， `APP_SECRET` 应设置为各应用互不相同的字符串并保密

## 日志

在生产模式下，日志会输出到 `log` 目录下
