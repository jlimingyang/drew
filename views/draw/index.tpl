<!doctype html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
<meta name="apple-mobile-web-app-capable" content="yes" />
<meta name="renderer" content="webkit">
<meta http-equiv="X-UA-Compatible" content="IE=Edge,chrome=1">
<title>抽奖</title>

<link rel="stylesheet" href="/static/css/css.css" type="text/css">
  <script src="https://unpkg.com/sweetalert/dist/sweetalert.min.js"></script>


</head>
<body>
<div class="bg"></div>
<div class="dog"><img src="/static/img/1_01.png"></div>
<div class="reward"><img id="levelImg" src="/static/img/1_02.png"><!-- 四等奖1_02.png;  三等奖1_02a.png;  二等奖1_02b.png;  一等奖1_02c.png --></div>
<div class="turntable">
	<img src="/static/img/1_03.png" class="img">
	<div class="jz"><span class="name"><img src="/static/img/1_05.png"></span></div>
</div>
<div class="btn">
  	<div class="start" id="btntxt" onclick="start()">开始</div>
	<div class="start btn_none" id="btnqr" onclick="">确认名单</div>
</div>
<div class="confirmbox">
	<div class="bg_img"><img src="/static/img/1_04.png"></div>
    <div class="box">
    	<div class="title">获奖号码公布</div>
        <div class="lucknum">
        	<span>00</span>
            <span>00</span>
            <span>00</span>
            <span>00</span>
            <div class="clear"></div>
        </div>
        <div class="conbox"></div>
    </div>
</div>
<div class="zjmd">
  	<div class="title">获奖名单公布栏</div>
    <div class="div1">
    	<div class="p1a">一等奖：</div>
        <div class="p1a">二等奖：</div>
        <div class="p1a">三等奖：</div>
        <div class="p1a">四等奖：</div>
        <div class="div2">
            <div class="p1">??</div>
            <div class="p2">??</div>
            <div class="p3">??</div>
            <div class="p4">??</div>
            <div class="p5">??</div>
            <div class="p6">??</div>
            <div class="p7">??</div>
            <div class="p8">??</div>
            <div class="p9">??</div>
            <div class="p10">??</div>
        </div>
        <div id="content">
            <div class="zjmd_bt_xy"></div>
        </div>
    </div>
</div>
    <script type="text/javascript" src="http://code.jquery.com/jquery-1.12.1.min.js"></script>
<script type="text/javascript" src="/static/js/adaptive-version2.js"></script>
<script type="text/javascript" src="/static/js/cj.js"></script>


</body>
</html>


