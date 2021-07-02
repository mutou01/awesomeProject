package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

type ComponentController struct{}

type User231 struct {
	gorm.Model
	Name string `gorm:"varchar(100);not null"`
}

type text_User231 struct {
	gorm.Model
	Name string `gorm:"varchar(100);not null"`
	Alias string `gorm:"varchar(100);not null"`
}

type text_Team struct {
	gorm.Model
	Name string `gorm:"varchar(100);not null"`
	Alias string `gorm:"varchar(100);not null"`
}

type text_team_data_all struct {
	gorm.Model
	Team_id uint //请注意以后的主键都要使用int64，uint的32位和64位存在区别。
	Name string `gorm:"varchar(100);not null"`
	Alias string `gorm:"varchar(100);not null"`
	Description string `gorm:"varchar(255)"`
	Mulity_table int `gorm:"tinyint(2);default:0"`
}

type text_team_fields struct {
	gorm.Model
	All_id uint
	Name string `gorm:"varchar(100);not null"`
	Alias string `gorm:"varchar(100);not null"`
	Description string `gorm:"varchar(255);not null"`
}

type text_team_data_1 struct {
	gorm.Model
	All_id uint `gorm:"not null"`
	Data string `gorm:"text;not null"json:"data"`
}

func CutString(data string) []string{
	var value []string
	data_1 := data+" "
	var j=0
	for i,v := range data_1 {
		if v == ' ' {
			value = append(value,data_1[j:i])
			j=i+1
		}
	}
	return value
}

func (c *ComponentController) PostDataAll(ctx *gin.Context) {

}

func (text_team_data_all) TableName() string {
	return "text_team_data_all"
}

func DataAllPost(ctx *gin.Context) {

}

var  db1 *gorm.DB

func main(){
	db,err :=gorm.Open("mysql","root:Lw^k1+4Ufya?dmWC*Y8@tcp(172.81.212.167:33006)/cde_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SingularTable(true)
	db.AutoMigrate(&text_User231{})
	db.AutoMigrate(&text_Team{})
	db.AutoMigrate(&text_team_data_all{})
	db.AutoMigrate(&text_team_fields{})
	db.AutoMigrate(&text_team_data_1{})

	db1=db

	r := gin.Default()
	//新建总表
	r.POST("/text/data_all/post", AddNewDataAll)//每段后面都加上err检错

	//删除总表
	r.POST("/text/data_all/delete", DeleteDataAll)

	//更新总表
	r.POST("/text/data_all/put", UpdataDataAll)

	//查询总表
	r.GET("/text/data_all/get", GetDataAll)

	//新建字段表
	r.POST("/text/field/post", AddField)

	//删除字段表
	r.POST("/text/field/delete", DeleteField)

	//更新字段表
	r.POST("/text/field/put", UpdataField)

	//查询字段表
	r.GET("/text/field/get", GetField)

	//新建数据库表
	r.POST("/text/data/post", AddData)

	//删除数据库表
	r.POST("/text/data/delete", DeleteData)

	//更新数据库表
	r.POST("text/data/put", UpdataData)

	//查询数据库表
	r.GET("/text/data/get", GetData)

	//创建范例
	r.POST("/text/add",AddText)

	//查询范例
	r.GET("/text/get",GetText)

	defer db1.Close()

	r.Run(":1235")

	/*代码块
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})

	//AsciiJSON	生成具有转义的非ASCII字符的ASCII-only JSON
	//r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context){
		data :=map[string]interface{}{
			"lang": "GO语言",
			"tag": "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})
	r.Run(":1235")
	//{"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}

	fmt.Println("http://localhost:1235/ping")
	r.Run(":1235")
*/

	//db.Create(&text_User231{Name:"chenyun",Alias:"成员"})//增
/*
	var user text_User231
	db.First(&user,5)
	fmt.Println(user)
	var user2 text_User231
	db.First(&user2,"name = ? and alias = ?","chenyun6","成员6")//查询
	fmt.Println(user2)
	var user3 text_User231
	db.Model(&user3).Where("name = ? and Alias = ?","chenyun","成员").Update("Alias","成员7")//改
	fmt.Println(user3)
	var user4 text_User231
	db.Delete(&user3)
	db.Unscoped().Delete(&user4,"alias= ?","成员3")
*/
/*
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context){
		data :=map[string]interface{}{
			"lang": "GO语言",
			"tag": "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.POST("/post_form", func(c *gin.Context) {
		message := c.PostForm( "message")
		nick := c.DefaultPostForm("nick","ni")

		c.JSON(200 , gin.H{
			"status" :	"post",
			"message" : message,
			"nick"	:	nick,
		})
	})

	//url上相应字段的匹配
	///welcome?firstname=li&lastname=daoshuang
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname","Guest")
		lastname := c.Query("lastname")//若无，返回空串

		c.String(http.StatusOK, "Hellow %s %s",firstname,lastname)
	})

	//post获取数据
	//POST http://:1235/post?id=1234&page=1 HTTP/1.1/Content-Type: application/x-www-form-urlencoded/name=manu&message=this_is_great
	r.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	//curl -X POST http://localhost:8080/upload \
	//  -F "file=@/Users/appleboy/test.zip" \
	//  -H "Content-Type: multipart/form-data"
	r.POST("/upload", func(c *gin.Context) {
		// 单文件

		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件到指定的路径
		// c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})


	fmt.Println("http://localhost:1235/")
	r.Run(":1235")
*/
}

func AddNewDataAll(ctx *gin.Context){
	UserId,err:= strconv.Atoi(ctx.PostForm("user_id"))
	if err != nil {
		return
	}
	TeamId,err:= strconv.Atoi(ctx.PostForm("team_id"))
	if err != nil {
		return
	}
	name := ctx.PostForm("name")
	alias := ctx.PostForm("alias")
	description := ctx.PostForm("description")
	MulityTable, err := strconv.Atoi(ctx.PostForm("mulity_table"))
	if err != nil {
		return
	}

	//检查命名，对name和alias跑Check函数

	//验证权限，用user_id和team_id到数据库里面看权限
	if UserId!=TeamId {
		return
	}

	//nsq消息写入

	db1.Create(&text_team_data_all{Team_id: uint(TeamId),Name: name,Alias: alias,Description: description,Mulity_table: MulityTable})
}

func DeleteDataAll(ctx *gin.Context){
	data_all_id := ctx.PostForm("data_all_id")

	//验证权限
	db1.Delete(&text_team_data_all{},data_all_id)
}

func UpdataDataAll(ctx *gin.Context) {
	DataAllId, err := strconv.Atoi(ctx.PostForm("data_all_id"))
	if err != nil {
		return
	}
	name := ctx.PostForm("name")
	alias := ctx.PostForm("alias")
	description := ctx.PostForm("description")
	MulityTable, err := strconv.Atoi(ctx.PostForm("mulity_table"))
	if err != nil {
		return
	}

	//检查命名

	//验证权限

	//写入
	db1.Model(&text_team_data_all{}).Where("id = ?",DataAllId).
		Update(text_team_data_all{Name: name, Alias: alias, Description: description, Mulity_table: MulityTable})
}

func GetDataAll(ctx *gin.Context) {
	UserId,err:= strconv.Atoi(ctx.PostForm("user_id"))
	if err != nil {
		return
	}
	TeamId,err:= strconv.Atoi(ctx.PostForm("team_id"))
	if err != nil {
		return
	}

	//验证权限

	var DataAlls []text_team_data_all ///////////////////

	db1.Where("user_id = ? and team_id = ?",UserId,TeamId).Find(&DataAlls)
	for _,i :=range DataAlls{
		var fields []text_team_fields
		db1.Where("id = ?",i.ID).Find(&fields)
		for _,j :=range fields{
			var datas []text_team_data_1
			db1.Where("id = ?",j.ID).Find(&datas)
			for _,z := range datas {
				fmt.Println(z)
			}
			fmt.Println(fields)
		}
		fmt.Println(DataAlls)
	}
	fmt.Println()
}

func AddField(ctx *gin.Context) {
	name := ctx.PostForm("name")
	alias := ctx.PostForm("alias")
	description := ctx.PostForm("description")
	AllId, err := strconv.Atoi(ctx.PostForm("all_id"))
	if err != nil {
		return
	}

	//检查命名

	//验证权限

	db1.Create(&text_team_fields{Name: name,Alias: alias,Description: description,All_id: uint(AllId)})
}

func DeleteField(ctx *gin.Context) {
	FieldId := ctx.PostForm("id")

	//验证权限

	var delete text_team_fields

	db1.Where("id = ?",FieldId).Find(&delete)

	db1.Delete(&delete)
}

func UpdataField(ctx *gin.Context) {
	field_id := ctx.PostForm("id")
	name := ctx.PostForm("name")
	alias := ctx.PostForm("alias")
	description := ctx.PostForm("description")

	//检查命名

	//验证权限

	db1.Model(&text_team_fields{}).Where("id = ?",field_id).
		Update(text_team_fields{Name: name,Alias: alias,Description: description})
}

func GetField(ctx *gin.Context) {
	DataAllId ,err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		return
	}

	//检查权限

	var v text_team_fields

	db1.Where("id = ?",DataAllId).Find(&v)

	fmt.Println(v)

	ctx.JSON(http.StatusOK,v)
}

func AddData(ctx *gin.Context) {
	data := ctx.PostForm("data")
	AllId, err:= strconv.Atoi(ctx.PostForm("all_id"))
	if err != nil {
		fmt.Println("AllId error")
		return
	}

	//检查格式

	//从总表中根据总表id获取所有字段（按照id降序）。
	var key = []text_team_fields{}
	db1.Where("all_id = ?",AllId).Order("id asc").Find(&key)
	//fmt.Println(key)


	m1 := make(map[uint]interface{})
	//fmt.Println(CutString(data))
	value := CutString(data)
	//for k.name
	for i,k := range key{
		m1[k.ID] = value[i]
	}

	//fmt.Println(m1)

	b,err := json.Marshal(m1)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	//检查权限

	db1.Create(&text_team_data_1{Data: string(b), All_id: uint(AllId)})
	ctx.JSON(http.StatusOK,m1)
}

func DeleteData(ctx *gin.Context) {
	DataId , err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil{
		return
	}

	//检查权限

	var delete text_team_data_1

	db1.Where("id = ?",DataId).Find(&delete)

	db1.Delete(&delete)
}

func UpdataData(ctx *gin.Context) {
	DataId := ctx.PostForm("data_id")
	data := ctx.PostForm("data")
	AllId := ctx.PostForm("all_id")

	//检查权限

	var key = []text_team_fields{}
	db1.Where("all_id = ?",AllId).Order("id desc").Find(&key)

	m1 := make(map[uint]interface{})
	//fmt.Println(CutString(data))
	value := CutString(data)
	//for k.name
	for i,k := range key{
		m1[k.ID] = value[i]
	}

	b,err := json.Marshal(m1)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	db1.Model(&text_team_data_1{}).Where("id = ?", DataId).
		Update(text_team_data_1{Data: string(b)})

	ctx.JSON(http.StatusOK,b)
}

func GetData(ctx *gin.Context) {
	AllId, err := strconv.Atoi(ctx.Query("all_id"))
	if err != nil {
		return
	}


	//检查权限
	var key = []text_team_fields{}
	db1.Where("all_id = ?",AllId).Find(&key)

	var value = []text_team_data_1{} //切片保证返回所有记录
	db1.Where("all_id = ?",AllId).Find(&value)

	//字符串转json，json根据key的k.ID进行选取未被删除的字段数据，然后append给result

	result := []map[uint]interface{}{}
	for _ , v := range value{
		r := make(map[uint]string)
		err := json.Unmarshal([]byte(v.Data),&r)
		if err != nil {
			return
		}
		re := make(map[uint]interface{})
		for _ , k := range key{
			re[k.ID]=r[k.ID]
		}
		result=append(result,re)
	}

	fmt.Println(result)
	ctx.JSON(http.StatusOK,result)
}

func AddText(ctx *gin.Context) {
	teamId, err := strconv.Atoi(ctx.PostForm("team_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	name := ctx.PostForm("name")
	alias := ctx.PostForm("alias")
	description := ctx.PostForm("description")
	mulityTable, err := strconv.Atoi(ctx.PostForm("mulity_table"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	v := text_team_data_all{Team_id: uint(teamId),Name: name,Alias: alias,Description: description,Mulity_table: mulityTable}

	db1.Create(&v)

	ctx.JSON(http.StatusOK,v)
}

func GetText(ctx *gin.Context) {
	id , err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		return
	}

	var v []string

	var value []text_team_data_all

	db1.Where("id = ?",id).Take(&value)
	db1.Where("id = ?",id).First(&value)
	db1.Where("id = ?",id).Last(&value)
	db1.Where("id = ?",id).Find(&value)

	db1.Model(&text_team_data_all{}).Pluck("name",&v)

	ctx.JSON(http.StatusOK,v)
}