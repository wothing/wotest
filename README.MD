# wotest

#### go语言构建的DSL测试库，用于构建Http测试与断言。所有语句在后缀为.wt的文本文件中撰写。编码采用UTF-8

基础语法：

所有语法关键词均为小写字母。

```
load [file.wt]: 加载依赖测试用例

set [name var]: 设置全局变量
$var: 取出全局变量。注意：$header,$body,$resp相关的前缀为保留前缀，请勿占用

env [name]: 读取系统环境变量
file [path]: 读入文件

echo [var]: 输出，用于调试
pretty [var]: 美化json并输出，用于调试

get [url]: 构建GET请求
post [url]: 构建POST请求
header [key value]: 请求Header
body [var]: 请求Body
form [key value]: 请求form
multipart [key value]: multipart k-v
multifile [name path]: multipart file
ret [code?]: 发起请求，若code不为空则校验http status code。

断言为true则返回T，否则返回F
has [a b]: a是否含有b
eq: [a b]: a是否等于b
neq: [a b]: a不等于b
gt: [a b]: a大于b
gte: [a b]: a大于等于b
lt: [a b]: a小于b
lte [a b]: a小于等于b
```

如果一个参数中含有空格，请使用双引号'"'包裹，例如："Bearer xxxxxx"

如果需要命令嵌套，请使用'\`'包裹，例如：echo \`env GOPATH\`，读取系统环境变量，并输出

支持array；
支持struct;

用例请参见`demo/http.wt`

![demo](./demo.png)

### TODO:

1、regexp 正则匹配

2、条件语句
```
if x then
a
else
b
end
```

### FAQ:
