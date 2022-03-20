package main

import (
	"flag"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/zyylhn/ldap_search/ldapsearch"
)
var Host string
var User string
var Password string
var BaseDn	string
var Filter string
var Attribut string
var SizeLimit int
var Or bool
var Listcmd bool


func main() {
	flag.StringVar(&User,"user","","set user")
	flag.StringVar(&Password,"pass","","set user password")
	flag.StringVar(&Host,"host","","set target")
	flag.StringVar(&BaseDn,"dn","","set base dn")
	flag.StringVar(&Filter,"f","","set Attribute filtering")
	flag.StringVar(&Attribut,"a","","Sets the attribute to be queried")
	flag.IntVar(&SizeLimit,"limit",0,"Set the number of data items to return")
	flag.BoolVar(&Or,"or",false,"Multiple attributes return logic（default and）")
	flag.BoolVar(&Listcmd,"listcmd",false,"Listing common commands")
	flag.Parse()
	if Listcmd{
		cmd:=ldapsearch.GetCommonQueries()
		for k,v:=range cmd{
			fmt.Println(k," : ",v)
		}
	}
	if Host==""{
		fmt.Println("must set host")
		return
	}
	if User==""{
		fmt.Println("must set user")
		return
	}
	if Password==""{
		fmt.Println("must set pass")
		return
	}
	conn,err:=LoginBind(User,Password,fmt.Sprintf("%v:%v",Host,389))
	if err!=nil{
		fmt.Println(err)
		return
	}
	result,err:=ldapsearch.SearchLdap(conn,BaseDn,Filter,Attribut,SizeLimit)
	if err!=nil{
		fmt.Println(err)
		return
	}
	re:=ldapsearch.OutputResult(Attribut,result,Or)
	switch re.(type) {
	case *ldap.SearchResult:
		re.(*ldap.SearchResult).PrettyPrint(2)
	case map[string][]string:
		for k,v:=range re.(map[string][]string){
			fmt.Println(k)
			for _,re:=range v{
				fmt.Println(re)
			}
		}
	}
}

func LoginBind(ldapUser, ldapPassword string,addr string) (*ldap.Conn, error) {
	l, err := ldap.Dial("tcp",addr)
	if err != nil {
		return nil, err
	}
	err = l.Bind(ldapUser,ldapPassword)

	if err != nil {
		fmt.Println("ldap password is error: ", ldap.LDAPResultInvalidCredentials)
		return nil, err
	}
	fmt.Println(ldapUser,"登录成功")
	return l, nil
}