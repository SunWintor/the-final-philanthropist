stageMap = {"-2":"未加入房间", "-1":"游戏未开始", "0":"游戏开始", "1":"回合开始", "2":"捐赠", "3":"捐赠结束", "4":"舆论惩罚", "5":"破产判定", "6":"游戏结束"}
currentStage = "-1"
roomId = ""
gameId = ""
roomStatus = 0
playerList = []

tfp = function () {
    console.log("run")
    if (!roomId) {
        return
    }
    if (currentStage === "-1") {
    }
}
setInterval(tfp,2000)


getRoomInfo = function () {
    $.ajax({
        type:"get",    // 请求类型
        url:"/room/info",    // 请求URL
        data:JSON.stringify({"user_id": getUserId()}),    // 请求参数 即是 在Servlet中 request.getParement();可以得到的字符串
        dataType:"json",    // 数据返回类型
        cache:false, // 是否缓存
        async:true,    // 默认为true 异步请求
        contentType: "application/json",
        success:function(result){    // 成功返回的结果(响应)
            console.log(JSON.stringify(result))
            if (result["code"] == 0) {
                console.log("test")
            }
        }
    });
}

$(function (){
    $('#game_start').bind('click', function () {
        console.log("click")
        $.ajax({
            type:"post",    // 请求类型
            url:"/room/join/random",    // 请求URL
            data:JSON.stringify({"user_id": getUserId()}),    // 请求参数 即是 在Servlet中 request.getParement();可以得到的字符串
            dataType:"json",    // 数据返回类型
            cache:false, // 是否缓存
            async:true,    // 默认为true 异步请求
            contentType: "application/json",
            success:function(result){    // 成功返回的结果(响应)
                console.log(JSON.stringify(result))
                if (result["code"] == 0) {
                    console.log("success")
                }
            }
        });
    });
})
