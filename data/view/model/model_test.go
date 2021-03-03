package model

import (
	"encoding/json"
	"fmt"
	"github.com/xxjwxc/gormt/data/config"
	"testing"
)

func TestTypeName(t *testing.T) {
	fmt.Println(getTypeName("tinyint", true))
}

func TestTools(t *testing.T) {
	str := `{"DbName":"oauth_db","PackageName":"model","TabList":[{"Name":"user_account_tbl","Notes":"用户账号","SQLBuildStr":"","Em":[{"Name":"id","Notes":"","Type":"int(11)","Index":[{"Key":1,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"account","Notes":"","Type":"varchar(64)","Index":[{"Key":2,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"password","Notes":"","Type":"varchar(64)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"account_type","Notes":"帐号类型:0手机号，1邮件","Type":"int(11)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"app_key","Notes":"authbucket_oauth2_client表的id","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"user_info_tbl_id","Notes":"","Type":"int(11)","Index":[{"Key":3,"KeyName":"user_info_id"}],"IsNull":false,"ForeignKeyList":[{"TableName":"user_info_tbl","ColumnName":"id"}]},{"Name":"reg_time","Notes":"","Type":"datetime","Index":null,"IsNull":true,"ForeignKeyList":null},{"Name":"reg_ip","Notes":"","Type":"varchar(15)","Index":[{"Key":3,"KeyName":"user_info_id"}],"IsNull":true,"ForeignKeyList":null},{"Name":"bundle_id","Notes":"","Type":"varchar(255)","Index":null,"IsNull":true,"ForeignKeyList":null},{"Name":"describ","Notes":"","Type":"varchar(255)","Index":null,"IsNull":true,"ForeignKeyList":null}]},{"Name":"user_info_tbl","Notes":"用户信息","SQLBuildStr":"","Em":[{"Name":"","Notes":"","Type":"gorm.Model","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"nickname","Notes":"","Type":"varchar(32)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"headurl","Notes":"","Type":"varchar(255)","Index":null,"IsNull":true,"ForeignKeyList":null}]},{"Name":"oauth2_access_token","Notes":"token认证","SQLBuildStr":"","Em":[{"Name":"id","Notes":"","Type":"int(11)","Index":[{"Key":1,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"access_token","Notes":"","Type":"varchar(255)","Index":[{"Key":2,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"token_type","Notes":"","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"app_key","Notes":"key","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"username","Notes":"用户名","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"expires","Notes":"过期时间","Type":"datetime","Index":null,"IsNull":false,"ForeignKeyList":null}]},{"Name":"user","Notes":"","SQLBuildStr":"","Em":[{"Name":"userId","Notes":"","Type":"int(11)","Index":[{"Key":1,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"name","Notes":"","Type":"varchar(30)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"sex","Notes":"","Type":"int(11)","Index":[{"Key":3,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"job","Notes":"","Type":"int(11)","Index":null,"IsNull":false,"ForeignKeyList":null}]},{"Name":"organ","Notes":"","SQLBuildStr":"","Em":[{"Name":"id","Notes":"","Type":"int(11)","Index":[{"Key":1,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"userId","Notes":"","Type":"int(11)","Index":[{"Key":3,"KeyName":""}],"IsNull":true,"ForeignKeyList":[{"TableName":"user","ColumnName":"sex"}]},{"Name":"type","Notes":"","Type":"int(11)","Index":null,"IsNull":true,"ForeignKeyList":null},{"Name":"score","Notes":"","Type":"int(11)","Index":null,"IsNull":true,"ForeignKeyList":null}]},{"Name":"sign_client_tbl","Notes":"","SQLBuildStr":"","Em":[{"Name":"id","Notes":"","Type":"int(11)","Index":[{"Key":1,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"app_key","Notes":"","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"app_secret","Notes":"","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"expire_time","Notes":"超时时间","Type":"datetime","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"strict_sign","Notes":"是否强制验签:0：用户自定义，1：强制","Type":"int(255)","Index":null,"IsNull":true,"ForeignKeyList":null},{"Name":"strict_verify","Notes":"是否强制验证码:0：用户自定义，1：强制","Type":"int(11)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"token_expire_time","Notes":"token过期时间","Type":"int(11)","Index":null,"IsNull":false,"ForeignKeyList":null}]},{"Name":"user_paybill_order","Notes":"","SQLBuildStr":"","Em":[{"Name":"id","Notes":"","Type":"int(11)","Index":[{"Key":1,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"paybill_id","Notes":"二次账单id","Type":"bigint(20)","Index":[{"Key":3,"KeyName":"order_id"}],"IsNull":false,"ForeignKeyList":null},{"Name":"order_id_mysql","Notes":"MySql中的订单Id","Type":"bigint(20)","Index":[{"Key":3,"KeyName":"order_id"}],"IsNull":false,"ForeignKeyList":null},{"Name":"order_item_id_mysql","Notes":"MySql中的订单ItemId","Type":"bigint(20)","Index":[{"Key":3,"KeyName":"order_id"}],"IsNull":false,"ForeignKeyList":null},{"Name":"order_id_mssql","Notes":"MsSql中的订单Id","Type":"bigint(20)","Index":null,"IsNull":false,"ForeignKeyList":null}]},{"Name":"oauth2_client_tbl","Notes":"client key 信息","SQLBuildStr":"","Em":[{"Name":"id","Notes":"","Type":"int(11)","Index":[{"Key":1,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"app_key","Notes":"","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"app_secret","Notes":"","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"expire_time","Notes":"超时时间","Type":"datetime","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"strict_sign","Notes":"是否强制验签:0：用户自定义，1：强制","Type":"int(255)","Index":null,"IsNull":true,"ForeignKeyList":null},{"Name":"strict_verify","Notes":"是否强制验证码:0：用户自定义，1：强制","Type":"int(11)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"token_expire_time","Notes":"token过期时间","Type":"int(11)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"aaa","Notes":"","Type":"json","Index":null,"IsNull":true,"ForeignKeyList":null}]},{"Name":"oauth2_refresh_token","Notes":"刷新token","SQLBuildStr":"","Em":[{"Name":"id","Notes":"","Type":"int(11)","Index":[{"Key":1,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"refresh_token","Notes":"","Type":"varchar(255)","Index":[{"Key":2,"KeyName":""}],"IsNull":false,"ForeignKeyList":null},{"Name":"token_type","Notes":"","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"app_key","Notes":"","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"username","Notes":"","Type":"varchar(255)","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"expires","Notes":"","Type":"datetime","Index":null,"IsNull":false,"ForeignKeyList":null},{"Name":"token_expire_time","Notes":"","Type":"int(11)","Index":null,"IsNull":false,"ForeignKeyList":null}]}]}`
	var pkg DBInfo
	json.Unmarshal([]byte(str), &pkg)
	// out, _ := json.Marshal(pkg)
	// tools.WriteFile("test.txt", []string{string(out)}, true)

	config.SetIsWEBTag(true)
	config.SetIsOutFunc(false)
	list, _ := Generate(pkg)
	fmt.Println(list)

	config.SetForeignKey(true)
	list, _ = Generate(pkg)
	fmt.Println(list)
}
