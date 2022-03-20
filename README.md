# LDAP查询工具

可以通过各种简单高级语法进行ldap查询，并集成在在[zscan](https://github.com/zyylhn/zscan)的exploit的ldap利用模块中，使用基础下面稍为讲一下

## 使用方式

```shell
Usage of ./ldap_search:
  -a string
        Sets the attribute to be queried    (要查询的属性)
  -dn string
        set base dn						（要查询的dn）
  -f string
        set Attribute filtering    （查询时筛选的属性）
  -host string
        set target   （目标）
  -limit int
        Set the number of data items to return  （返回的最大条数）
  -or
        Multiple attributes return logic（default and） （要查询的属性为多个值的时候的查询逻辑）
  -pass string
        set user password   （密码）
  -user string		（用户名）
        set user

```

## 免责声明

本工具仅面向**合法授权**的企业安全建设行为，如您需要测试本工具的可用性，请自行搭建靶机环境。

在使用本工具进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。**请勿对非授权目标进行扫描。**

如您在使用本工具的过程中存在任何非法行为，您需自行承担相应后果，我们将不承担任何法律及连带责任。

在安装并使用本工具前，请您**务必审慎阅读、充分理解各条款内容**，限制、免责条款或者其他涉及您重大权益的条款可能会以加粗、加下划线等形式提示您重点注意。 除非您已充分阅读、完全理解并接受本协议所有条款，否则，请您不要安装并使用本工具。您的使用行为或者您以其他任何明示或者默示方式表示接受本协议的，即视为您已阅读并同意本协议的约束。