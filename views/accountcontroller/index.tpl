<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<!-- saved from url=(0013)http://u.com/ -->
<html xmlns="http://www.w3.org/1999/xhtml"><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">

<title>蜜芽宝贝用户管理系统 - 首页</title>
<meta http-equiv="X-UA-Compatible" content="IE=EmulateIE7">
<meta http-equiv="pragma" content="no-cache"> 
<meta http-equiv="cache-control" content="no-cache, must-revalidate">
<meta http-equiv="expires" content="0">
<link href="/static/css/style.css" type="text/css" rel="stylesheet">
<script type="text/javascript" src="/static/js/jquery-1.5.2.min.js"></script>
<script type="text/javascript" src="/static/js/main.js"></script>
<script type="text/javascript">
$(document).ready(function(){
	$("input[name='username']").focus();
	var ie6=($.browser.msie&&($.browser.version=="6.0")&&!$.support.style);
	if(ie6){
	    $(".intext").hover(function(){
	        $(this).addClass("intextHover");
	    },function(){
	        $(this).removeClass("intextHover");
	    });
	    $(".intext").focus(function(){
	        $(this).addClass("intextFocus");
	    }).blur(function(){
	        $(this).removeClass("intextFocus");
	    });
	    $(".inputBtn").hover(function(){
	        $(this).addClass("inputBtnHover");
	    },function(){
	        $(this).removeClass("inputBtnHover");
	    });
	}
});
</script>
<style type="text/css">
html{
	height:100%;
	background-color:#355370;
	background-image:-moz-linear-gradient(center top,#3a5976,#2a4763);
	background-image:-webkit-gradient(linear,left top,left bottom,from(#3a5976),to(#2a4763));
	filter:progid:DXImageTransform.Microsoft.gradient(startColorstr='#3a5976',endColorstr='#2a4763');
}
.loginWrapper{
	width:312px;
	height:465px;
	margin:0 auto;
	position:relative;
/*	border:1px dashed #ccc;*/
}
.loginForm{
	color:#fff;
	padding-left:2px;
	padding-top:145px;
	_padding-top:138px;
}
.loginForm p{
	margin-top:20px;
	font-size:20px;
	color:#fff;
	line-height:35px;
	font-family:"微软雅黑","宋体";
}
.intext{
    border:2px solid #d0d0d0;
    background:#d0d0d0;
    width:271px;
    height:28px;
    line-height:28px;
    color:#333;
    font-size:16px;
    margin-right:16px;
    padding-left:8px;
    z-index:1;
}
.intext:hover,.intextHover{
    border:2px solid #fefefe;
    background:#fefefe;
}
.intext:focus,.intextFocus{
    border:2px solid #fefefe;
    background:#fefefe;
}
.inputBtn{
	font-size:18px;
	display:inline-block;
	padding:7px 36px 6px 36px;
	*padding:7px 25px 5px 25px;
	_padding:8px 25px 3px 25px;
	cursor:pointer;
    color:#fff;
    background-color:#258;
    border:2px solid #fff;
}
.inputBtn:hover,.inputBtnHover{
    background-color:#036;
}
.loginSubmit{
	margin-top:40px;
}
#loginTitleFg{
	position:absolute;
	top:60px;
	_top:57px;
	left:-20px;
	z-index:10;
	font:bold 50px/50px "微软雅黑","黑体","宋体";
	color:#fff;
}
#loginTitleBg{
	position:absolute;
	top:63px;
	_top:60px;
	left:-20px;
	z-index:9;
	font:bold 50px/50px "微软雅黑","黑体","宋体";
	color:#2c2c2c;
}
#loginTitleEnFg{
	position:absolute;
	top:60px;
	_top:57px;
	left:190px;
	z-index:10;
	font:bold 52px/52px "Arial";
	color:#f0b310;
}
#loginTitleEnBg{
	position:absolute;
	top:63px;
	_top:60px;
	left:190px;
	z-index:9;
	font:bold 52px/52px "Arial";
	color:#222;
}
.loginWrapper h2{
	margin-top:150px;
	margin-left:2px;
	font:normal 34px/34px "微软雅黑","宋体";
	_font-weight:bold;
	color:#fff;
}
img{
	width: 100px;
	height: 100px;
}
</style>
</head>
<body>
<div class="loginWrapper">
<div id="loginTitleFg">蜜芽MIA</div>
<div id="loginTitleBg">蜜芽MIA</div>
<div id="loginTitleEnFg" title="统一管理系统 Universal Manage System">UMS</div>
<div id="loginTitleEnBg">UMS</div>
<div class="clear"></div>
<div class="loginForm">
	<p>用户名</p>
	<input type="text" class="intext" id="username" name="username">
	<p>密码</p>
	<input type="password" class="intext" id="password" name="password" maxlength="15">
	{{if .isVarified}}
		<p class="captche">验证码<input type="text" class="intext" id="captche" name="captche"></p>
		<img src="/static/img/1.jpg" />
	{{end}}
	<!-- {{ .xsrfdata }} -->
	<div class="loginSubmit">
		<input type="submit" class="inputBtn" onclick="ajaxLogin();" value="登录">
	</div>
</div>
</div>
{{template "common/footer.tpl" .}}