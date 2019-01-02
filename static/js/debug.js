$(function(){
    initAlert();
});

function initAlert() {
    var  msg = $("#debug").val()
    var msgType = $("#msgType").val()
    console.log(msgType+"--"+msg)
    var type = "error"
    if(parseInt(msgType) == 1){
        type = "success"
    }
    if("debug" != msg.toString()){
        console.log(1)
        swal(msg, {
            button: false,
            icon:type
        });
    }
}

