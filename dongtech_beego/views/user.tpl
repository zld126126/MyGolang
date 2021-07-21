<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <header>
    <h1 class="logo">Hello World</h1>
    <div class="description">
      This is {{.Name}} demo.
    </div>
  </header>
  <footer>
    <div class="author">
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
