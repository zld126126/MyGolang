package app

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2 = "<!DOCTYPE html>\r\n<html lang=\"en\">\r\n\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    <title>首页</title>\r\n    <script type=\"text/javascript\" src=\"//cdn.staticfile.org/jquery/2.0.0/jquery.min.js\"></script>\r\n    <script type=\"text/javascript\" src=\"//cdn.staticfile.org/jqueryui/1.10.2/jquery-ui.min.js\"></script>\r\n</head>\r\n\r\n<body>\r\n    <div class=\"container\">\r\n        <div class=\"row clearfix\">\r\n            <div class=\"col-md-12 column\">\r\n                <h3>\r\n                    {{.WEBSITE_TITLE}}\r\n                </h3>\r\n                <p>\r\n                    <em>{{.WEBSITE_TITLE}}</em>由<strong>DongTech</strong>集成开发和维护。\r\n                </p>\r\n            </div>\r\n            <div id=\"up_image\">\r\n                <form action=\"\" method=\"post\" enctype=\"multipart/form-data\">\r\n                    <br>\r\n                    <p>问题名称: <input type=\"text\" name=\"Ask\" value=\"\" id=\"Ask\"/></p><br/>\r\n                    <p>问题答案: <input name=\"Answer\" value=\"\" id=\"Answer\" style=\"width: 1200px;\"></p><br/>\r\n                    <input type=\"button\" value=\"提交\" style=\"width: 100px;color:red;margin-top: 30px;margin-left: 30px;\" onclick=\"GetChatGPTResult()\">\r\n                </form>\r\n            </div>\r\n        </div>\r\n    </div>\r\n\r\n    <script type=\"text/javascript\">\r\n        function GetChatGPTResult(){\r\n            $(\"#Answer\").val(\"\")\r\n            var data = {\r\n                'ask': $(\"#Ask\").val().trim(),\r\n            };\r\n\r\n            if (data.ask == \"\"){\r\n                alert(\"请输入正确的问题~\")\r\n                return\r\n            }\r\n\r\n            $.ajax({\r\n                type: \"POST\",\r\n                url: \"/chatgpt\",\r\n                data: data,\r\n                dataType: 'json',\r\n                async: true,\r\n                success: function (result) {\r\n                    console.log(result.Answer)\r\n                    $(\"#Answer\").val(result.Answer)\r\n                },\r\n                error: function (result) {\r\n                    alert(\"请求失败,请稍后重试...\")\r\n                },\r\n            });\r\n        }\r\n    </script>\r\n</body>\r\n\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"index.html"}}, map[string]*assets.File{
	"/templates/index.html": &assets.File{
		Path:     "/templates/index.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1676016907, 1676016907532833400),
		Data:     []byte(_Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1676023929, 1676023929534779900),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1676016907, 1676016907533829900),
		Data:     nil,
	}}, "")
