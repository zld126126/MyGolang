package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

var (
	IndexHtml = `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>auto-api exe生成工具</title>
	<script src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
</head>
<body>
<p>一个简单的后端API .exe生成助手</p><br>
PORT端口:<br>
<input id="port" type="text" name="port" value="9090">
<br>
METHOD类型:<br>
<input id="method" type="text" name="method" value="GET">
<br>
API接口:<br>
<input id="api" type="text" name="api" value="/test">
<br>
JSON数据:<br>
<input id="json" type="text" name="json" value='{"data":"test"}'>
<br><br>
<button type="button" onclick="submit()">提交</button>
<script>
	function submit(){
		var port = document.getElementById("port").value;
		var api = document.getElementById("api").value;
		var json = document.getElementById("json").value;
		var method = document.getElementById("method").value;
		var url = "http://localhost:10086/add"+"?"+"port="+port+"&api="+api+"&json="+json+"&method="+method
		$.ajax({
    		contentType: 'application/json; charset=utf-8',
			url:url,
			async: true,
			success: function(data) {
				if (data.success){
					alert("创建成功,文件路径:"+data.path);
				}else{
					alert("创建失败");
				}
			},
		 	error:function(){
				console.log('服务器繁忙,稍后重试');
			}
		});
	} 
</script>
</body>
</html>`

	Dir, _ = os.Getwd()
	PORT   = 10086
)

type H map[string]interface{}

func htmlExample() {
	// http://localhost:10086
	index := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(NowTime(), ",请求主页")

		// 扩展html原生写法
		//w.Header().Set("Content-Type", "text/html")
		//t, _ := template.ParseFiles(Dir + "/templates/index.html")
		//t.Execute(w, "")

		w.Header().Set("Content-Type", "text/html")
		t := template.New("index")
		t.Parse(IndexHtml)
		t.Execute(w, nil)
	}

	add := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(NowTime(), ",请求创建exe")

		values, _ := url.ParseQuery(r.URL.RawQuery)
		fmt.Println(values)

		port, ok1 := values["port"]
		if !ok1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("参数port有误"))
		}
		if len(port) == 0 || port[0] == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("参数port有误"))
		}
		p, err := strconv.Atoi(port[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("参数port有误"))
		}
		if p <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("参数port有误"))
		}

		api, ok2 := values["api"]
		if !ok2 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("参数api有误"))
		}
		if len(api) == 0 || api[0] == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("参数api有误"))
		}

		method, ok3 := values["method"]
		if !ok3 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("参数method有误"))
		}
		if len(method) == 0 || method[0] == "" || (method[0] != "POST" && method[0] != "GET") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("参数method有误"))
		}

		data, ok4 := values["json"]
		if !ok4 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("参数json有误"))
		}
		if len(data) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("参数json有误"))
		}

		fileName := RandFileName() + ".exe"
		goFileName := RandFileName() + ".go"

		w.Header().Set("content-type", "text/json")
		resp := make(H)
		resp["success"] = false
		resp["path"] = ""

		flag := Add(p, api[0], data[0], method[0], goFileName)
		if flag {
			DelUseless(goFileName)
			resp["success"] = true
			resp["path"] = Dir + "\\" + fileName
		}

		result, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

	http.HandleFunc("/index", index)
	http.HandleFunc("/", index)
	http.HandleFunc("/add", add)
	fmt.Println("server is run :", PORT)
	http.ListenAndServe(":"+fmt.Sprint(PORT), nil)
}

func NowTime() string {
	return time.Now().Local().String()
}

func RandFileName() string {
	tick64 := func() int64 {
		return time.Now().Local().UnixNano() / 1e6
	}

	formatTime := func(t time.Time) string {
		return t.Format("20060102")
	}

	date := formatTime(time.Now())
	r := rand.Intn(1000)
	tick := tick64()
	fileName := fmt.Sprintf("%s%d%03d", date, tick, r)

	return fileName
}

var (
	goTemplate1 = `
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	t := func() string {
		return time.Now().Local().String()
	}

	Api := func(w http.ResponseWriter, r *http.Request) {
		if "%s" != r.Method{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("请求Method有误"))
			return
		}
		fmt.Println(t(), ",请求api:", r.Host+r.RequestURI,r.Method)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
	goTemplate2 = `))
	}

	http.HandleFunc("%s", Api)
	fmt.Println("server is run",%d)
	http.ListenAndServe(":"+fmt.Sprint(%d), nil)
}`
)

func Add(port int, api string, data string, method string, fileName string) bool {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Open file err =", err)
		return false
	}
	defer file.Close()

	content := fmt.Sprintf(goTemplate1, method) + "`" + data + "`" + fmt.Sprintf(goTemplate2, api, port, port)
	_, err = file.Write([]byte(content))
	if err != nil {
		fmt.Println("Write file err =", err)
		return false
	}

	cmd := exec.Command("cmd", "/c", "go build "+fileName)
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Build exe err =", err)
			return false
		}

		fmt.Println("Write & Build file success")
		return true
	}

	return false
}

func DelUseless(fileName string) {
	p := Dir + "//" + fileName
	err := os.Remove(p)
	if err != nil {
		fmt.Println(fileName + "删除失败")
	}
}

var (
	IndexJson = `{"port":9090,"api":"/test","method":"GET","json":{"data":"test"}}`
	JsonPath  = `config.json`
)

type JsonFile struct {
	Port   int                    `json:"port"`
	Method string                 `json:"method"`
	Api    string                 `json:"api"`
	Data   map[string]interface{} `json:"json"`
}

func jsonExample() {
	// 判断文件或文件夹是否存在
	chkExist := func(path string) bool {
		_, err := os.Stat(path)
		if err != nil {
			if os.IsExist(err) {
				return true
			}
			if os.IsNotExist(err) {
				return false
			}
			fmt.Println(err)
			return false
		}
		return true
	}

	// 默认json文件
	upsertJsonFile := func() bool {
		path := Dir + "\\" + JsonPath
		if !chkExist(path) {
			file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				fmt.Println("Open file err =", err)
				return false
			}
			defer file.Close()

			_, err = file.Write([]byte(IndexJson))
			if err != nil {
				fmt.Println("Write file err =", err)
				return false
			}
		}
		return true
	}

	// 读取默认json
	readJsonFile := func() string {
		path := Dir + "\\" + JsonPath
		file, err := os.OpenFile(path, os.O_RDONLY, 0666)
		if err != nil {
			fmt.Println("Open file err =", err)
			return ""
		}
		defer file.Close()

		b, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Open file err =", err)
			return ""
		}
		return string(b)
	}

	flag := upsertJsonFile()
	if !flag {
		return
	}

	jsonContent := readJsonFile()
	if jsonContent == "" {
		return
	}

	var jf *JsonFile
	err := json.Unmarshal([]byte(jsonContent), &jf)
	if err != nil {
		fmt.Println("Read Config err =", err)
		return
	}

	if jf == nil {
		fmt.Println("json文件有误")
		return
	}

	if jf.Port <= 0 {
		fmt.Println("参数Port有误")
		return
	}

	if jf.Api == "" {
		fmt.Println("参数Api有误")
		return
	}

	t := func() string {
		return time.Now().Local().String()
	}

	Api := func(w http.ResponseWriter, r *http.Request) {
		if jf.Method != r.Method {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("请求Method有误"))
			return
		}
		fmt.Println(t(), ",请求api:", r.Host+r.RequestURI, r.Method)

		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(jf.Data)
		w.Write(b)
	}

	http.HandleFunc(jf.Api, Api)
	fmt.Println("server is run :", jf.Port)
	http.ListenAndServe(":"+fmt.Sprint(jf.Port), nil)
}

func main() {
	// html请求
	// http://localhost:10086
	htmlExample()

	// json格式
	// http://localhost:9090/test
	jsonExample()
}
