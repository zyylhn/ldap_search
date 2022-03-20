package ldapsearch

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"strings"
)

func SearchLdap(conn *ldap.Conn,baseDn,filter,attributeStr string,sizeLimit int) (*ldap.SearchResult,error) {
	var attributeList []string
	if attributeStr!=""{
		attributeList=strings.Split(attributeStr,",")
	}
	sql:=ldap.NewSearchRequest(baseDn, ldap.ScopeWholeSubtree,ldap.NeverDerefAliases,sizeLimit,0,false,filter,attributeList,nil)
	return conn.Search(sql)
}

func OutputResult(attstr string,result *ldap.SearchResult,or bool) interface{} {
	remap:=make(map[string][]string)
	attlist:=strings.Split(attstr,",")
	if len(result.Entries)>0{
		for _,item:=range result.Entries{
			if attstr!=""{
				for _,att:=range attlist{
					cn:=item.GetAttributeValues(att)
					if len(cn)>0{
						for _,i:=range cn{
							remap[item.DN]=append(remap[item.DN],fmt.Sprint("\t"+att+":"+i))
						}
					}
				}
				if !or{
					if len(remap[item.DN])<len(attlist){
						delete(remap,item.DN)
					}
				}
			}
		}
	}
	if len(remap)!=0{
		return remap
	}else {
		return result
	}
}

func GetCommonQueries() map[string]string {
	var cmd map[string]string
	cmd= make(map[string]string)
	cmd["获取所有邮件地址"]="-dn \"CN=users,DC=lhn,DC=com\" -f \"(objectClass=*)\" -a mail"
	cmd["查询域控制器主机名"]="-dn \"OU=Domain Controllers,DC=lhn,DC=com\" -f \"(objectClass=*)\" -a dNSHostName,operatingSystem,operatingSystemVersion"
	cmd["查询域管理员"]="-dn \"CN=Domain Admins,CN=users,DC=lhn,DC=com\" -f \"(member=*)\" -a member"
	cmd["查询所有域用户"]="-dn \"CN=users,DC=lhn,DC=com\" -f \"(sAMAccountType=805306368)\" -a sAMAccountName"
	cmd["查看加入域的所有计算机名(不包括域控)"]="-dn \"CN=Computers,DC=lhn,DC=com\" -f \"(sAMAccountType=805306369)\" -a dNSHostName,operatingSystem,operatingSystemVersion"
	return cmd
}