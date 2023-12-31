# gin_mall

项目结构
- api/



# 通用返回结构

| 参数名  | 类型   | 说明                               |
| ------- | ------ | ---------------------------------- |
| status  | int    | 业务状态码                         |
| message | string | 业务状态短语                       |
| data    | Object | 错误返回NILL；成功返回数据包括null |


# 用户路由

## 用户注册

| 地址           | 方法   | 请求格式 |说明                   |
| ---------------- | ------ | -------|--------------- |
| `/user/register` | `POST` |`JSON`| 通过该接口，注册新用户 |



### 请求参数

| 参数名 | 必要性 | 参数位置 | 类型|说明 |
| ---- | ------ |---| ---- | ---- |
| username | Y | body | string | 用户名 |
| nickname | N | body | string | 昵称 |
| password | Y | body | string | 登录密码 |



### 返回结果

| 参数名 | 类型 | 说明                   |
| ------ | ---- | ---------------------- |
| data   | NULL | 注册成功和错误返回NILL |



### 示例

```shell
curl http://127.0.0.1:8000/api/v1/user/register -X POST -H "Content-Type: application/json" -d '{"username":"2233","password":"233"}'
```







## 用户登录

| 地址          | 方法   | 请求格式 | 说明     |
| ------------- | ------ | -------- | -------- |
| `/user/login` | `POST` | `JSON`   | 用户登录 |



### 请求参数

| 参数名   | 必要性 | 参数位置 | 类型   | 说明     |
| -------- | ------ | -------- | ------ | -------- |
| username | Y      | body     | string | 用户名   |
| password | Y      | body     | string | 登录密码 |



### 返回结果

| 参数名  | 类型   | 说明                         |
| ------- | ------ | ---------------------------- |
| data    | Object | 成功返回Object；错误返回NILL |

> **成功 data 数据**
>
> | 参数名 | 类型   | 说明     |
> | ------ | ------ | -------- |
> | token  | string | 登录凭证 |
> | user   | Object | 用户信息 |
>
> > **用户信息** 
> >
> > | 参数名    | 类型   | 说明     |
> > | --------- | ------ | -------- |
> > | id        | int    | 用户ID   |
> > | user_name | string | 用户名   |
> > | nickname  | string | 昵称     |
> > | email     | string | 邮箱     |
> > | status    | string | 用户状态 |
> > | avatar    | string | 头像     |
> > | create_at | int    | 创建时间 |



### 示例

```shell
curl http://127.0.0.1:8000/api/v1/user/login -X POST -H "Content-Type: application/json" -d '{"username":"2233","password":"233"}'
```





## 修改用户信息

| 地址    | 方法  | 请求格式 | 说明         |
| ------- | ----- | -------- | ------------ |
| `/user` | `PUT` | `JSON`   | 修改用户信息 |



### 请求参数

| 参数名        | 必要性 | 参数位置 | 类型   | 说明         |
| ------------- | ------ | -------- | ------ | ------------ |
| Authorization | Y      | header   | string | 用户登录凭证 |
| username      | N      | body     | string | 用户名       |
| nickname      | N      | body     | string | 昵称         |
| email         | N      | body     | string | 邮箱         |
| avatar        | N      | body     | string | 头像地址     |



### 返回结果

| 参数名  | 类型   | 说明                         |
| ------- | ------ | ---------------------------- |
| data    | Object | 成功返回Object；错误返回NILL |

> **成功 data 数据**
>
> | 参数名 | 类型   | 说明     |
> | ------ | ------ | -------- |
> | user   | Object | 用户信息 |
>
> > **用户信息** 
> >
> > | 参数名    | 类型   | 说明     |
> > | --------- | ------ | -------- |
> > | id        | int    | 用户ID   |
> > | user_name | string | 用户名   |
> > | nickname  | string | 昵称     |
> > | email     | string | 邮箱     |
> > | status    | string | 用户状态 |
> > | avatar    | string | 头像     |
> > | create_at | int    | 创建时间 |



### 示例

```shell
curl http://127.0.0.1:8000/api/v1/user 
```



## 上传用户头像

| 地址      | 方法   | 请求格式    | 说明         |
| --------- | ------ | ----------- | ------------ |
| `/avatar` | `POST` | `form-data` | 修改用户信息 |



### 请求参数

| 参数名        | 必要性 | 参数位置 | 类型   | 说明         |
| ------------- | ------ | -------- | ------ | ------------ |
| Authorization | Y      | header   | string | 用户登录凭证 |
| file          | Y      | form     | file   | 头像文件     |



### 返回结果

| 参数名  | 类型   | 说明                         |
| ------- | ------ | ---------------------------- |
| data    | Object | 成功返回Object；错误返回NILL |

> **成功 data 数据**
>
> | 参数名 | 类型   | 说明     |
> | ------ | ------ | -------- |
> | user   | Object | 用户信息 |
>
> > **用户信息** 
> >
> > | 参数名    | 类型   | 说明     |
> > | --------- | ------ | -------- |
> > | id        | int    | 用户ID   |
> > | user_name | string | 用户名   |
> > | nickname  | string | 昵称     |
> > | email     | string | 邮箱     |
> > | status    | string | 用户状态 |
> > | avatar    | string | 头像     |
> > | create_at | int    | 创建时间 |







## 显示金额

| 地址     | 方法   | 请求格式 | 说明         |
| -------- | ------ | -------- | ------------ |
| `/money` | `POST` | `any`    | 修改用户信息 |



### 请求参数

| 参数名        | 必要性 | 参数位置 | 类型   | 说明         |
| ------------- | ------ | -------- | ------ | ------------ |
| Authorization | Y      | header   | string | 用户登录凭证 |



### 返回结果

| 参数名  | 类型   | 说明                         |
| ------- | ------ | ---------------------------- |
| status  | int    | 业务状态码                   |
| message | string | 业务状态短语                 |
| data    | Object | 成功返回Object；错误返回NILL |

> **成功 data 数据**
>
> | 参数名 | 类型   | 说明     |
> | ------ | ------ | -------- |
> | money  | number | 用户金额 |





# 地址相关路由

## 添加地址

| 地址         | 方法   | 请求格式 | 说明             |
| ------------ | ------ | -------- | ---------------- |
| `/addresses` | `POST` | `JSON`   | 添加商品收货地址 |

### 请求参数

| 参数名        | 必要性 | 参数位置 | 类型   | 说明       |
| ------------- | ------ | -------- | ------ | ---------- |
| Authorization | Y      | header   | string | 登录凭证   |
| name          | Y      | body     | string | 收货人名称 |
| address       | Y      | body     | string | 详细地址   |
| phone         | Y      | body     | string | 手机号码   |



### 返回结果

| 参数名 | 类型   | 说明                   |
| ------ | ------ | ---------------------- |
| data   | Object | 成功添加的地址详细信息 |

> **成功 data 数据**
>
> | 参数名  | 类型   | 说明             |
> | ------- | ------ | ---------------- |
> | aid     | int    | 当前地址的字段ID |
> | name    | string | 收货人名称       |
> | phone   | string | 手机号码         |
> | address | string | 地址             |

### 示例

```shell
curl http://127.0.0.1:8000/api/v1/user/register -X POST -H "Content-Type: application/json" -d '{"username":"2233","password":"233"}'
```



## 删除地址

| 地址           | 方法     | 请求格式 | 说明     |
| -------------- | -------- | -------- | -------- |
| `/address/:id` | `DELETE` | `JSON`   | 删除地址 |

### 请求参数

| 参数名        | 必要性 | 参数位置 | 类型   | 说明     |
| ------------- | ------ | -------- | ------ | -------- |
| id            | Y      | path     | int    | 地址ID   |
| Authorization | Y      | header   | string | 登录凭证 |

### 返回结果

| 参数名 | 类型 | 说明               |
| ------ | ---- | ------------------ |
| data   | NULL | 成功和错误返回NILL |

### 示例

```shell
curl http://127.0.0.1:8000/api/v1/user/register -X POST -H "Content-Type: application/json" -d '{"username":"2233","password":"233"}'
```



## 更新地址

| 地址       | 方法  | 请求格式 | 说明     |
| ---------- | ----- | -------- | -------- |
| `/address` | `PUT` | `JSON`   | 更新地址 |

### 请求参数

| 参数名        | 必要性 | 参数位置 | 类型   | 说明       |
| ------------- | ------ | -------- | ------ | ---------- |
| Authorization | Y      | header   | string | 登录凭证   |
| aid           | Y      | body     | int    | 地址ID     |
| name          | N      | body     | string | 收货人名称 |
| phone         | N      | body     | string | 手机号码   |
| address       | N      | body     | string | 地址       |

### 返回结果

| 参数名 | 类型   | 说明                       |
| ------ | ------ | -------------------------- |
| data   | Object | 更新成功返回地址的详细信息 |

> **成功 data 数据**
>
> | 参数名  | 类型   | 说明             |
> | ------- | ------ | ---------------- |
> | aid     | int    | 当前地址的字段ID |
> | name    | string | 收货人名称       |
> | phone   | string | 手机号码         |
> | address | string | 地址             |

### 示例 

```shell
curl http://127.0.0.1:8000/api/v1/user/register -X POST -H "Content-Type: application/json" -d '{"username":"2233","password":"233"}'
```



## 地址列表

| 地址       | 方法  | 请求格式 | 说明                       |
| ---------- | ----- | -------- | -------------------------- |
| `/address` | `GET` | `JSON`   | 获取当前用户的所有地址信息 |

### 请求参数

| 参数名        | 必要性 | 参数位置 | 类型   | 说明     |
| ------------- | ------ | -------- | ------ | -------- |
| Authorization | Y      | header   | string | 登录凭证 |

### 返回结果

| 参数名 | 类型          | 说明                 |
| ------ | ------------- | -------------------- |
| data   | Array[Object] | 成功返回所有地址信息 |

> **成功 item 数据**
>
> | 参数名  | 类型   | 说明             |
> | ------- | ------ | ---------------- |
> | aid     | int    | 当前地址的字段ID |
> | name    | string | 收货人名称       |
> | phone   | string | 手机号码         |
> | address | string | 地址             |

### 示例

```shell
curl http://127.0.0.1:8000/api/v1/user/register -X POST -H "Content-Type: application/json" -d '{"username":"2233","password":"233"}'
```



## 地址详情

| 地址           | 方法   | 请求格式 | 说明     |
| -------------- | ------ | -------- | -------- |
| `/address/:id` | `POST` | `JSON`   | 地址详情 |



### 请求参数

| 参数名        | 必要性 | 参数位置 | 类型   | 说明     |
| ------------- | ------ | -------- | ------ | -------- |
| Authorization | Y      | header   | string | 登录凭证 |
| id            | Y      | path     | int    | 地址ID   |



### 返回结果

| 参数名 | 类型   | 说明                   |
| ------ | ------ | ---------------------- |
| data   | Object | 成功返回地址的详细信息 |

> **成功 data 数据**
>
> | 参数名  | 类型   | 说明             |
> | ------- | ------ | ---------------- |
> | aid     | int    | 当前地址的字段ID |
> | name    | string | 收货人名称       |
> | phone   | string | 手机号码         |
> | address | string | 地址             |

### 示例

```shell
curl http://127.0.0.1:8000/api/v1/user/register -X POST -H "Content-Type: application/json" -d '{"username":"2233","password":"233"}'
```







# 轮播图

## 获取轮播图列表

| 地址        | 方法  | 请求格式 | 说明           |
| ----------- | ----- | -------- | -------------- |
| `/carousel` | `GET` |          | 获取轮播图列表 |



### 返回结果

| 参数名 | 类型 | 说明                   |
| ------ | ---- | ---------------------- |
| data   | NULL | 注册成功和错误返回NILL |





# 上传文件

## 上传文件

| 地址        | 方法   | 请求格式 | 说明           |
| ----------- | ------ | -------- | -------------- |
| `/carousel` | `POST` | form     | 获取轮播图列表 |



### 请求参数

| 参数名        | 必要性 | 参数位置 | 类型   | 说明               |
| ------------- | ------ | -------- | ------ | ------------------ |
| Authorization | Y      | header   | string | 登录凭证           |
| type          | N      | query    | string | 上传图片的用途类型 |
| file          | Y      | form     | file   | 文件               |





### 返回结果

| 参数名 | 类型   | 说明        |
| ------ | ------ | ----------- |
| data   | Object | 成功返回url |







































































































































































> **data 数据**
>
> 
>
> | 参数名 | 类型 | 说明 |
> | ------ | ---- | ---- |
> |        |      |      |
>
> - object 详写k-v对
> - array 元素信息
> - 为 string number bool null 不用详写





### 示例

```shell
curl http://127.0.0.1:8000/api/v1/user/register -X POST -H "Content-Type: application/json" -d '{"email":"2233@qq.com","pwd":"233"}'
```

















































































































































---
像是淘宝之类的电商平台，商品在数据库中的表结构会是这样的？
ChatGPT

在电商平台中，商品表的结构通常包含以下字段：

id：商品唯一标识符，通常是自增的整数型。

name：商品名称，通常是字符串类型。

description：商品描述，通常是字符串类型。

price：商品价格，通常是浮点型。

stock：商品库存量，通常是整数型。

category_id：商品分类ID，通常是整数型。

brand_id：品牌ID，通常是整数型。

status：商品状态，通常是整数型，用于表示商品是否上架、下架、删除等状态。

image_url：商品图片的URL地址，通常是字符串类型。

created_at：商品创建时间，通常是时间戳类型。

updated_at：商品最后更新时间，通常是时间戳类型。

具体的表结构可能会因电商平台的不同而有所不同，但通常都会包含以上字段。此外，可能还会有其他的字段，如商品规格、销量、评价等信息。

我

商品的详情页，有没有必要单独存放在一张表中？
ChatGPT

对于电商平台来说，商品的详情页通常包含大量的文本、图片、视频等信息，而且这些信息往往是用户可编辑的，因此将商品详情页单独存放在一张表中是有必要的。

通常情况下，商品详情页会单独存储在一个名为 product_detail 的表中，该表会包含以下字段：

id：商品详情页唯一标识符，通常是自增的整数型。

product_id：商品ID，通常是整数型，表示该详情页所属的商品。

content：商品详情页的文本内容，通常是字符串类型。

images：商品详情页所包含的图片地址，通常是字符串类型或JSON格式的字符串。

videos：商品详情页所包含的视频地址，通常是字符串类型或JSON格式的字符串。

created_at：商品详情页创建时间，通常是时间戳类型。

updated_at：商品详情页最后更新时间，通常是时间戳类型。

将商品详情页单独存放在一个表中，有助于提高数据的可读性和可维护性，同时也方便对商品详情页进行管理和维护。

我

详情页想要在多个终端展示，用什么格式存储，用markdown可以吗
ChatGPT

如果商品详情页需要在多个终端上展示，通常需要将其存储为一种通用的格式，比如 HTML 或 JSON 格式。Markdown 格式在展示效果上比较简洁，但是不同的终端对其的解析效果可能存在差异，因此不太适合作为跨终端展示的格式。

将商品详情页存储为 HTML 或 JSON 格式，可以确保在不同的终端上展示效果基本一致。如果需要在多个终端上展示商品详情页，可以将商品详情页的文本内容存储在数据库中，并使用 HTML 或 JSON 格式对其进行格式化。在需要展示时，可以将格式化后的内容传递给前端，然后在前端进行展示。

如果你希望使用 Markdown 格式来存储商品详情页，也可以尝试将 Markdown 格式的文本内容转换为 HTML 格式，然后将其存储在数据库中。在需要展示时，可以将 HTML 格式的内容传递给前端，然后在前端进行展示。这样可以确保在不同的终端上展示效果基本一致，并且可以兼顾 Markdown 格式的简洁性。