# dbgo

通过数据库表结构,生成go实体文件

## 安装

```bash
go install github.com/wslky/dbgo
```
## 使用

```bash
dbgo -u root -p 123456. -d mall -url 127.0.0.1:3306
```

cli参数说明

```bash
-u 数据库用户名
-p 数据库密码
-d 数据库名
-url 数据库地址
-pn 包名
-fn 文件名 支持 goStyle | go_style | gostyle
-jt json tag  支持 goStyle | go_style | gostyle
```
