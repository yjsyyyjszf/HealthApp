// HttpWeb
//go get -u github.com/labstack/echo/...
/*
DSN数据源名称
  [username[:password]@][protocol[(address)]]/dbname[?param1=value1¶mN=valueN]
  user@unix(/path/to/socket)/dbname
  user:password@tcp(localhost:5555)/dbname?charset=utf8&autocommit=true
  user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?charset=utf8mb4,utf8
  user:password@/dbname
  无数据库: user:password@/

echo web 参考例子
https://blog.csdn.net/Noob_coder_JZ/article/details/83410095
*/

// HTTP状态码的分类

// 分类	分类描述
// 1** 	信息，服务器收到请求，需要请求者继续执行操作
// 2** 	成功，操作被成功接收并处理
// 3** 	重定向，需要进一步的操作以完成请求
// 4**	客户端错误，请求包含语法错误或无法完成请求
// 5** 	服务器错误，服务器在处理请求的过程中发生了错误

package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"strconv"
	"strings"

	// "encoding/json"
	"io/ioutil"

	"./Data"
	"./Units"

	// "fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CustomContext struct {
	echo.Context
}

const (
	empty = ""
	tab   = "\t"
)
const (
	DB_Driver = "root:root@tcp(127.0.0.1:3306)/hit?charset=utf8"
	IMAGE_Dir = "D:/code/C++/HealthApp/bin/win32/DCM_SAVE/Images"
	PAGE_Dir  = "D:/code/C++/HealthApp/Tools/PageWeb"
	PAGE_TEST = "F:/temp/HealthApp/PageWeb"
)

var name string
var maridb_db *sql.DB

func init() {
	flag.StringVar(&name, "name", "default", "log in user")
}

func Healthsystem(c echo.Context) error {
	req := c.Request()
	println("req.URL.Path:" + req.URL.Path)
	println("D:/code/C++/HealthApp/Tools/PageWeb" + req.URL.Path)
	filepath := "D:/code/C++/HealthApp/Tools/PageWeb" + req.URL.Path
	return c.File(filepath)
}

func Login(c echo.Context) error {
	req := c.Request()
	// println("req.URL.Path:" + req.URL.Path)
	// println(PAGE_Dir + req.URL.Path)
	filepath := PAGE_Dir + req.URL.Path
	return c.File(filepath)
}

func LoadImageFile(c echo.Context) error {
	// println("---LoadImageFile---------")
	studyuid := c.FormValue("studyuid")
	image_hash_dir := Units.GetStudyHashDir(studyuid)
	filepath := IMAGE_Dir + image_hash_dir
	filepath += "/"
	filepath += studyuid
	filetype := c.FormValue("type")
	// println(filetype)
	if filetype == "json" {
		filepath += "/"
		filepath += studyuid
		filepath += ".json"
		return c.File(filepath)
	} else {
		seriesuid := c.FormValue("seriesuid")
		sopinstanceuid := c.FormValue("sopinstanceuid")
		filepath += "/"
		filepath += seriesuid
		filepath += "/"
		filepath += sopinstanceuid
		filepath += ".dcm"
		return c.File(filepath)
	}
}

func LoadViewPage(c echo.Context) error {
	req := c.Request()
	filepath := PAGE_Dir + req.URL.Path
	index := strings.Index(req.URL.Path, ".html?")
	if index > -1 {
		filename := req.URL.Path[0 : index+5]
		filepath = PAGE_Dir + filename
		return c.File(filepath)
	} else {
		return c.File(PAGE_Dir + req.URL.Path)
	}
}

func CheckLogin(c echo.Context) error {
	username := c.FormValue("user")
	userpwd := c.FormValue("password")
	println("username:" + username + "/userpwd:" + userpwd)
	// c.Response(http.StatusText())
	//c.SetResponse()
	return c.String(http.StatusOK, "ok")
}

func SaveReportdata(c echo.Context) error {
	var js string
	return c.String(http.StatusOK, string(js))
}

func GetReportdata(c echo.Context) error {
	var reportdata Study.ReportData
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
		err := json.Unmarshal(bodyBytes, &reportdata)
		// reportdata = json.Unmarshal(bodyBytes)
		println("reportdata.ReportIdentity")
		println(reportdata.ReportIdentity)
		if err != nil {
			println(err)
		}

	}
	println(string(bodyBytes))
	//-------------
	var reportDBdata Study.ReportData
	reportDBdata.ReportSave = "false"
	reportDBdata.ReportChange = "false"
	if maridb_db != nil {
		var sqlstr string
		sqlstr = "select  * from h_report where `ReportIdentity`=" + reportdata.ReportIdentity
		println(sqlstr)
		rows, err := maridb_db.Query(sqlstr)
		if err != nil {
			println(err)
		} else {
			for rows.Next() {
				// reportdata.ReportWriterID = sql.NullString{String: "", Valid: false}
				// reportdata.ReportCheckID = sql.NullString{String: "", Valid: false}
				// reportdata.ReportOther = sql.NullString{String: "", Valid: false}
				err = rows.Scan(&reportDBdata.ReportIdentity, &reportDBdata.StudyOrderIdentity,
					&reportDBdata.ReportState, &reportDBdata.ReportTemplate,
					&reportDBdata.ReportCreatDate, &reportDBdata.ReportWriterID,
					&reportDBdata.ReportCheckID, &reportDBdata.ReportCheckDate,
					&reportDBdata.ReportContent, &reportDBdata.ReportOther)
			}
		}
	}
	// 2-------------------------
	// reportdata.ReportContent = "this is json promble"
	// println("reportDBdata.ReportContent")
	// println(reportDBdata.ReportContent)

	// js, err := json.Marshal(reportDBdata)
	// println(string(js))
	// if err != nil {
	// 	println(err)
	// 	return c.String(http.StatusOK, string(js))
	// }
	////////////////////--
	// var data Study.ReportData
	// data.ReportSave = "false"
	// data.ReportChange = "false"
	// err = json.Unmarshal(js, &data)
	// println("data.ReportContent")
	// println(data.ReportContent)
	// var jsonstr string
	// jsonstr = "{\"ReportIdentity\": \"1577160124447\",\"StudyOrderIdentity\": "
	// jsonstr += "\"1577160124447\",\"ReportTemplate\": \"0\",\"ReportCreatDate\": \"2019-12-28 17:13:52\","
	// jsonstr += "\"ReportWriterID\": \"0\",\"ReportCheckID\": \"0\",\"ReportCheckDate\": \"2019-12-28 17:13:52\","
	// jsonstr += "\"ReportContent\": \"this a big box! json\""
	// jsonstr += ",\"ReportState\": \"0\",\"ReportSave\": \"false\",\"ReportChange\": \"false\",\"ReportOther\": \"0\"}"
	// //////////////////////////////////////////
	// println(jsonstr)
	// return c.String(http.StatusOK, string(jsonstr))
	// return c.String(http.StatusOK, string(js))
	////////////////3----------------------
	buf := bytes.NewBuffer([]byte{})
	jsonE := json.NewEncoder(buf)
	jsonE.SetEscapeHTML(false)
	jsonE.SetIndent(empty, tab)
	// reportdata.ReportContent = "this is json promble"
	jsonE.Encode(reportDBdata)
	reportstring := buf.String()
	// reportstring = strings.ReplaceAll(reportstring, "\\\"", "\"")
	reportstring = strings.ReplaceAll(reportstring, ":", " :")
	println("-----buf.String()-----------——")
	println(reportDBdata.ReportContent)
	println(buf.String())
	println(reportstring)
	// return c.String(http.StatusOK, buf.String())
	return c.String(http.StatusOK, reportstring)

	//////////////////
	// jsoncode := &Study.JsonCode{
	// 	Code: 1,
	// 	Msg:  "success",
	// 	Data: reportdata,
	// }
	// js, _ := json.Marshal(jsoncode)
	// println(string(js))
	// return c.String(http.StatusOK, string(js))
}

func GetDBStudyImage(c echo.Context) error {
	//分页查询https://blog.csdn.net/myth_g/article/details/89672722
	startTime := c.FormValue("start")
	endTime := c.FormValue("end")
	page := c.FormValue("page")
	limit := c.FormValue("limit")
	var studyjson Study.StudyDataJson
	if maridb_db != nil {
		var count int
		p, err := strconv.Atoi(page)
		checkErr(err)
		lim, err := strconv.Atoi(limit)
		checkErr(err)
		count = (p - 1) * lim
		var sqlstr string
		sqlstr = "select p.PatientIdentity,p.PatientName,p.PatientID,p.PatientBirthday," +
			" p.PatientSex,s.StudyUID,s.StudyID,s.StudyIdentity,s.StudyDateTime," +
			" s.StudyDescription, s.StudyModality from " +
			" h_patient p, h_study s where p.PatientIdentity = s.PatientIdentity and " +
			" s.StudyDateTime >= " + startTime + " and  s.StudyDateTime <= " + endTime +
			" order by s.PatientIdentity limit " + strconv.Itoa(count) + "," + limit
		println(sqlstr)
		rows, err := maridb_db.Query(sqlstr)
		if err != nil {
			println(err)
		} else {
			studyjson.Code = 0
			studyjson.Msg = ""
			studyjson.Count = 21
			for rows.Next() {
				var data Study.StudyData
				err = rows.Scan(&data.PatientIdentity, &data.PatientName,
					&data.PatientID, &data.PatientBirthday,
					&data.PatientSex, &data.StudyUID, &data.StudyID,
					&data.StudyOrderIdentity, &data.StudyDateTime,
					&data.StudyDescription, &data.StudyModality)
				studyjson.Data = append(studyjson.Data, data)
			}
			sqlstr = "select count(*) count from " +
				"h_patient p, h_study s where p.PatientIdentity = s.PatientIdentity and " +
				" s.StudyDateTime >= " + startTime + " and  s.StudyDateTime <= " + endTime
			rows, err := maridb_db.Query(sqlstr)
			if err == nil {
				for rows.Next() {
					rows.Scan(&studyjson.Count)
				}
			}
		}
	}
	js, err := json.Marshal(studyjson)
	if err != nil {
		println(err)
		return c.String(http.StatusOK, "null")
	}
	println(string(js))
	return c.String(http.StatusOK, string(js))
}

func GetDBStudyData(c echo.Context) error {
	//分页查询https://blog.csdn.net/myth_g/article/details/89672722
	startTime := c.FormValue("start")
	endTime := c.FormValue("end")
	page := c.FormValue("page")
	limit := c.FormValue("limit")
	// var mess strings.Builder
	// mess.WriteString("start:")
	// mess.WriteString(startTime)
	// mess.WriteString("end:")
	// mess.WriteString(endTime)
	// mess.WriteString("page:")
	// mess.WriteString(page)
	// mess.WriteString("limit:")
	// mess.WriteString(limit)
	// // mess := "start:" + startTime + "end:" + end + "page:" + page + "limit:" + limit
	// println(mess.String())
	var studyjson Study.StudyDataJson
	if maridb_db != nil {
		var count int
		p, err := strconv.Atoi(page)
		checkErr(err)
		lim, err := strconv.Atoi(limit)
		checkErr(err)
		count = (p - 1) * lim
		var sqlstr string
		sqlstr = "select p.PatientIdentity,p.PatientName,p.PatientID,p.PatientBirthday,p.PatientSex,p.PatientTelNumber," +
			" p.PatientAddr, p.PatientEmail, p.PatientCarID, s.StudyID ,s.StudyUID,s.StudyDepart,s.CostType," +
			" s.StudyOrderIdentity,s.ScheduledDateTime,s.ScheduledDateTime,s.StudyDescription, s.StudyModality, s.StudyCost, s.StudyState " +
			" from h_patient p, h_order s where p.PatientIdentity = s.PatientIdentity and StudyState > 0 and " +
			" s.ScheduledDateTime>=" + startTime + " and  s.ScheduledDateTime<=" + endTime + " order by s.StudyOrderIdentity " +
			" limit " + strconv.Itoa(count) + "," + limit
		println(sqlstr)
		rows, err := maridb_db.Query(sqlstr)
		if err != nil {
			println(err)
		} else {
			studyjson.Code = 0
			studyjson.Msg = ""
			studyjson.Count = 21
			for rows.Next() {
				var data Study.StudyData
				err = rows.Scan(&data.PatientIdentity, &data.PatientName,
					&data.PatientID, &data.PatientBirthday,
					&data.PatientSex, &data.PatientTelNumber,
					&data.PatientAddr, &data.PatientEmail,
					&data.PatientCarID, &data.StudyID,
					&data.StudyUID, &data.StudyDepart,
					&data.CostType, &data.StudyOrderIdentity,
					&data.ScheduledDateTime, &data.StudyDateTime, &data.StudyDescription,
					&data.StudyModality, &data.StudyCost,
					&data.StudyState)
				studyjson.Data = append(studyjson.Data, data)
			}
			sqlstr = "select count(*) count from " +
				" h_patient p, h_order s where p.PatientIdentity = s.PatientIdentity and StudyState > 0 and" +
				"  s.ScheduledDateTime>= " + startTime + " and  s.ScheduledDateTime <= " + endTime
			rows, err := maridb_db.Query(sqlstr)
			if err == nil {
				for rows.Next() {
					rows.Scan(&studyjson.Count)
				}
			}
		}
	}
	js, err := json.Marshal(studyjson)
	if err != nil {
		println(err)
		return c.String(http.StatusOK, "null")
	}
	println(string(js))
	return c.String(http.StatusOK, string(js))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkMariDB(db *sql.DB) {
	//db, err := OpenDB()//sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hit?charset=utf8")
	rows, err := db.Query("SELECT * FROM h_user")

	for rows.Next() {
		var idd int
		var username string
		var password string
		err = rows.Scan(&idd, &username, &password)
		checkErr(err)
		println(idd)
		println(username)
		println(password)
	}
}

func OpenDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	checkErr(err)
	return isOpen, db
}

func main() {
	if len(os.Args) > 0 {
		arg0 := os.Args[0:]
		println("--arg0--")
		println(arg0)
	}
	// var hash string = Units.GetStudyHashDir("1.2.840.113619.2.55.3.604688119.868.1249343483.504")
	// println(hash)
	maridb_db = nil
	open, db := OpenDB()
	if open == true {
		maridb_db = db
		// test
		//checkMariDB(db)
	}

	// flag.Parse() //暂停获取参数
	e := echo.New()
	//定义日志属性
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "${time_rfc3339_nano} ${remote_ip} ${method} ${uri} ${status}\n",
	// }))
	e.Use(middleware.Logger())
	//login
	e.POST("/login/checkuser", CheckLogin)
	e.GET("/Login/*", Login)
	e.GET("/login/*", Login)

	//ris
	e.GET("/healthsystem/ris/studydata/", GetDBStudyData)
	e.GET("/healthsystem/ris/stduyimage/", GetDBStudyImage)

	//ris/report
	e.POST("/healthsystem/ris/getreportdata", GetReportdata)
	e.POST("/healthsystem/ris/savereportdata", SaveReportdata)

	e.GET("/healthsystem/*", Healthsystem)

	// view dicom
	e.GET("/view/*", LoadViewPage)

	// other
	e.GET("/favicon.ico", func(c echo.Context) error {
		// println("----------favicon.ico--------")
		return c.File("D:/code/C++/HealthApp/Tools/PageWeb/favicon.ico")
	})

	// load image file
	//需要定义/WADO?*路由??
	e.GET("/*", LoadImageFile) //WADO?*
	e.Logger.Fatal(e.Start(":9090"))
}

//参考
func queryMySal() {
	var err error
	connect := "f:f@tcp(192.168.1.0:3306)/lr1"
	db, err := sql.Open("mysql", connect)
	if err != nil {
		println("connect mysql failed, address = "+connect, err)
	} else {
		sqlContent := "select real_nm, sex, birth, ext from member where tenant=? and mem_id=?"
		rows, err := db.Query(sqlContent, "ST", "2017")
		if err != nil {
			println(err)
		} else {
			for rows.Next() {
				nm := sql.NullString{String: "", Valid: false}
				sex := sql.NullString{String: "", Valid: false}
				birth := sql.NullString{String: "", Valid: false}
				ext := sql.NullString{String: "", Valid: false}
				err := rows.Scan(&nm, &sex, &birth, &ext) // birth字段在数据库中是空字段
				if err != nil {
					println("d12 error. ", err)
				}
				println(nm.String)
				println(sex.String)
				println(birth.String)
				println(ext.String)
			}
		}
	}
}

// sqlContent := "select real_nm, sex, birth, ext from member where tenant=? and mem_id=?"
// rows, err := db.Query(sqlContent, "ST", "2017")
// stmt, err := maridb_db.Prepare(sql)
// rows, err := stmt.Exec()//https://www.cnblogs.com/jackylee92/p/6209596.html
// text1 := "abcdefg"
// fmt.Println(text1[n]) //获取字符串索引位置为n的原始字节，比如a为97
// fmt.Println(text1[n:m]) //截取得字符串索引位置为 n 到 m-1 的字符串
// fmt.Println(text1[n:]) //截取得字符串索引位置为 n 到 len(s)-1 的字符串
// fmt.Println(text1[:m]) //截取得字符串索引位置为 0 到 m-1 的字符串
// fmt.Println(len(text1)) //获取字符串的字节数
// fmt.Println(utf8.RuneCountInString(text1)) //获取字符串字符的个数
// fmt.Println([]rune(text1)) // 将字符串的每一个字节转换为码点值，比如这里会输出[97 98 99 100 101 102 103]
// fmt.Println(string(text1[n])) // 获取字符串索引位置为n的字符值
