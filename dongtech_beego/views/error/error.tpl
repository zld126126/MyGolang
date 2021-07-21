<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>404</title>
      <style type="text/css">
          body.error_page
          {
              background-color: #00629f;
              background-image: url(https://lxxybucket.oss-cn-hangzhou.aliyuncs.com/EXILE/error.png);
              background-position: center top;
              background-repeat: no-repeat;
          }
          #error
          {
              color: #FFF;
              width: 100%;
              text-align: center;
              margin: 22% auto;
          }
          #error span,#error a
          {
              color:Yellow;
          }
          #error a:hover
          {
              color:#FFF;
              text-decoration: underline;
              cursor: pointer;
          }
         .error_number{
             font-size: 8rem;
         }
      </style>

</head>
<body class="error_page">
<h1 id="error">
    <span>{{.content}}</span>

</h1>
</body>
</html>
<script>
    window.onload=function () {
        $width=document.body.clientWidth;
        if($width>1024){
            document.getElementById("error").style.margin="22% auto";
        }else{
            document.getElementById("error").style.margin="42% auto";
        }
    }
</script>
