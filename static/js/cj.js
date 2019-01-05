document.oncontextmenu = new Function("event.returnValue=false");
document.onselectstart = new Function("event.returnValue=false");

var runTimes = 0, levelNum = 4;

var xinm = new Array();
var phone = new Array();
var auth = $("#auth").val();

function saveDraw(row,level) {
	$.ajax({
		type: 'POST',
		url: '/auth/api/saveDraw',
		async:false,
		dataType: 'json',
		beforeSend: function (xhr) {
			xhr.setRequestHeader("auth", auth);
		},
		data: {"row": row, "level": level},
		success: function (data) {
			console.log(data)
		}
	});
}

function initX() {
	$.ajax({
		type: 'POST',
		url: '/auth/api/queryDraw',
		async:false,
		dataType: 'json',
		beforeSend: function (xhr) {
			xhr.setRequestHeader("auth", auth);
		},
		data: {"page": 1, "pagesize": 99999999},
		success: function (data) {
			console.log(data)
			if(data.code === 200){
			$.each(data.data.List,function (index,v) {
				xinm[index] = v.Row
				phone[index] = " "
			})
			}else {
				swal("请重新登录!", {
					button: false,
					icon:"error"
				});
			}
		}
	});
}
initX()
$.each(xinm,function (index,v) {
	console.log(v)
})
console.log("--------------------")
$.each(phone,function (index,v) {
	console.log(v)
})
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
	if(zjnum.length == xinm.length){
		swal("无法抽奖!", {
			button: false,
			icon:"error"
		});
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
	num = Math.floor(Math.random() * xinm.length);
	var i_num = 0, hasNum = false;
	for (var a = 0; a < xinm.length; a++) {
	    if (xinm[a] != "" && xinm[a] != undefined) { hasNum = true; break;}
	}
	if (!hasNum) { swal("奖池号码已使用完毕!", {
		button: false,
		icon:"error"
	}); return false;}
	while (xinm[num] == "") {
	    num = Math.floor(Math.random() * xinm.length);
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
	saveDraw(xinm[num],levelNum)
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
	    else if (levelNum == 2) { $("#levelImg").attr("src", "/static/img/1_02b.png");/************指定中奖名单************/ }
	    else if (levelNum == 1) { $("#levelImg").attr("src", "/static/img/1_02c.png");}
	}

})

$(function(){
	qs();
  });
function qs(){
	$(".lucknum span:last").addClass("span");
}





