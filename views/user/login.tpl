<!DOCTYPE html>
<html lang="zh_cn">
<head>
    <meta charset="UTF-8">
    <title>用户登录</title>

     <script src="/static/js/jquery.min.js"></script>
  <link href="/static/css/bs3/bootstrap.css" rel="stylesheet">
  <script src="https://unpkg.com/sweetalert/dist/sweetalert.min.js"></script>
     <script src="/static/js/bs3/bootstrap.js"></script>
     <script src="/static/js/debug.js"></script>
    <link rel="stylesheet" href="/static/css/user/style.css">
</head>
<body>
    <div class="container">
        <div class="form row">
            <div class="form-horizontal col-md-offset-3" id="login_form">
            <input  id="debug" type="hidden" value = "{{.msg}}"/>
            <input  id="msgType" type="hidden" value = "{{.msgType}}"/>
                <h3 class="form-title">用户登录</h3>
                <form role="form" action="/user/login.py" method="post">
                <div class="col-md-9">
                    <div class="form-group">
                        <i class="fa fa-user fa-lg"></i>
                        <input class="form-control required" type="text" placeholder="用户名" id="username" value="{{if .username}}{{.username}}{{else}}{{end}}" name="username" autofocus="autofocus" maxlength="20"/>
                    </div>
                    <div class="form-group">
                            <i class="fa fa-lock fa-lg"></i>
                            <input class="form-control required" type="password" placeholder="密码" id="password" name="password" value="{{if .password}}{{.password}}{{else}}{{end}}" maxlength="8"/>
                    </div>
                    <div class="form-group col-md-offset-9">
                        <button type="button" class="btn btn-error pull-left" onclick="location='/registe.page'">去注册</button>
                        <button type="submit" class="btn btn-success pull-right" name="submit">登录</button>
                    </div>
                </div>
                </form>
            </div>
        </div>
    </div>
</body>
</html>