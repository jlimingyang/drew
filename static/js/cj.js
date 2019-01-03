document.oncontextmenu = new Function("event.returnValue=false");
document.onselectstart = new Function("event.returnValue=false");

var runTimes = 0, levelNum = 4;

var xinm = new Array();
var phone = new Array();

xinm[0]="01"
xinm[1]="02"
xinm[2]="03"
xinm[3]="04"
xinm[4]="05"
xinm[5]="06"
xinm[6]="07"
xinm[7]="08"
xinm[8]="09"
xinm[9]="10"
xinm[10]="11"
xinm[11]="12"
xinm[12]="13"
xinm[13]="14"
xinm[14]="15"
xinm[15]="16"
xinm[16]="17"
xinm[17]="18"
xinm[18]="19"
xinm[19]="20"
xinm[20]="21"
xinm[21]="22"
xinm[22]="23"
xinm[23]="24"
xinm[24]="25"
xinm[25]="26"
xinm[26]="27"
xinm[27]="28"
xinm[28]="29"
xinm[29]="30"
xinm[30]="31"
xinm[31]="32"
xinm[32]="33"
xinm[33]="34"
xinm[34]="35"
xinm[35]="36"
xinm[36]="37"
xinm[37]="38"
xinm[38]="39"
xinm[39]="40"

phone[0]="18111229182"
phone[1]="18111229182"
phone[2]="18111229182"
phone[3]="18111229182"
phone[4]="18111229182"
phone[5]="18111229182"
phone[6]="18111229182"
phone[7]="18111229182"
phone[8]="18111229182"
phone[9]="18111229182"
phone[10]="18111229182"
phone[11]="18111229182"
phone[12]="18111229182"
phone[13]="18111229182"
phone[14]="18111229182"
phone[15]="18111229182"
phone[16]="18111229182"
phone[17]="18111229182"
phone[18]="18111229182"
phone[19]="18111229182"
phone[20]="18111229182"
phone[21]="18111229182"
phone[22]="18111229182"
phone[23]="18111229182"
phone[24]="18111229182"
phone[25]="18111229182"
phone[26]="18111229182"
phone[27]="18111229182"
phone[28]="18111229182"
phone[29]="18111229182"
phone[30]="18111229182"
phone[31]="18111229182"
phone[32]="18111229182"
phone[33]="18111229182"
phone[34]="18111229182"
phone[35]="18111229182"
phone[36]="18111229182"
phone[37]="18111229182"
phone[38]="33333"
phone[39]="444444"

var nametxt = $('.name');
var phonetxt = $('.phone');
var pcount = xinm.length;//参加人数
var runing = true;  
var num = 0; //随机数存储
var to = 0;//从0开始
var numr = 10;//每次抽取幸运奖人数
var t;//循环调用
var lucknum = 0;
var pdnum = pcount;//参加人数判断是否抽取完

//大奖开始停止
function start() {
    if ($("#btntxt").hasClass("btn_none")) {  swal("错误操作，请先确认名单!", {
		button: false,
		icon:"error"
	});return false; }/************判断开始按钮是否可用************/
	var zjnum = $('.list').find('p');
	if(zjnum.length == pdnum){
		alert('无法抽奖');
	}else{
		if (runing) {
			runing = false;
			$('#btntxt').removeClass('start').addClass('stop');
			$('#btntxt').html('停止');
			startNum();
			$(".turntable .img").css("animation","5s linear 0s normal none infinite rotate");
			$(".turntable .img").css("animation-play-state","running");
		} else {
			runing = true;
			$('#btntxt').removeClass('stop').addClass('start');
			$('#btntxt').html('抽奖');
			stop();
	        bzd();//中奖函数
	        $('#btnqx').css('display','block');
	        $('.lucknum').css('display','none');
			$(".turntable .img").css("animation-play-state","paused");
	    }
	}
}

//循环参加名单
function startNum() {
	num = Math.floor(Math.random() * pcount);
	var i_num = 0, hasNum = false;
	for (var a = 0; a < pcount; a++) {
	    if (xinm[a] != "") { hasNum = true; break;}
	}
	if (!hasNum) { swal("奖池号码已使用完毕!", {
		button: false,
		icon:"error"
	}); return false;}
	while (xinm[num] == "") {
	    num = Math.floor(Math.random() * pcount);
	}
	nametxt.html(xinm[num]);
	phonetxt.html(phone[num]);
	t = setTimeout(startNum, 0);
}

//停止跳动
function stop() {
	clearInterval(t);
	t = 0;
}

//打印中奖名单
function bzd() {
	//获取中奖人数
	var zjnum = $('.list').find('p');
	console.log(xinm[num])
	console.log(phone[num])
	//打印中奖者名单
	$('.conbox').prepend("<p>"+xinm[num]+"   "+phone[num]+"</p>");
	$(".lucknum span:last,.conbox p:last").addClass("span");
	$('.confirmbox').show();
    //将已中奖者从数组中"删除",防止二次中奖
	xinm[$.inArray(xinm[num], xinm)] = "";
	phone[$.inArray(phone[num], phone)] = "";
	runTimes++;
	if (runTimes >= levelNum) {
	    $("#btntxt").addClass("btn_none");/************修改开始按钮为不可用************/
	    $("#btnqr").removeClass("btn_none");
	    levelNum--;
    }
}

//确认名单
$('#btnqr').on('click', function () {
    if ($("#btnqr").hasClass("btn_none")) {swal("错误操作，请先抽奖!", {
		button: false,
		icon:"error"
	}); return false; }/************判断确认名单按钮是否可用************/
	var cp = $('.conbox').find('p').removeAttr('style').clone();
	$('.zjmd_bt_xy').prepend(cp);
	$('.zjmd_bt_xy p:nth-last-child(1)').addClass("p1")
	$('.zjmd_bt_xy p:nth-last-child(2)').addClass("p2")
	$('.zjmd_bt_xy p:nth-last-child(3)').addClass("p3")
	$('.zjmd_bt_xy p:nth-last-child(4)').addClass("p4")
	$('.zjmd_bt_xy p:nth-last-child(5)').addClass("p5")
	$('.zjmd_bt_xy p:nth-last-child(6)').addClass("p6")
	$('.zjmd_bt_xy p:nth-last-child(7)').addClass("p7")
	$('.zjmd_bt_xy p:nth-last-child(8)').addClass("p8")
	$('.zjmd_bt_xy p:nth-last-child(9)').addClass("p9")
	$('.zjmd_bt_xy p:nth-last-child(10)').addClass("p10")
	$('.conbox').empty();
	$('.confirmbox').show();
	$('.lucknum').show();
	$("#btnqr").addClass("btn_none");
	if (levelNum > 0) {
	    runTimes = 0;
	    $("#btntxt").removeClass("btn_none")
	    /************修改抽奖轮次图标************/
	    if (levelNum == 3) { $("#levelImg").attr("src", "/static/img/1_02a.png");}
	    else if (levelNum == 2) { $("#levelImg").attr("src", "/static/img/1_02b.png"); pcount = 30/************指定中奖名单************/ }
	    else if (levelNum == 1) { $("#levelImg").attr("src", "/static/img/1_02c.png");}
	}
	
})

$(function(){	    
	qs();
  });
function qs(){
	$(".lucknum span:last").addClass("span");
}





