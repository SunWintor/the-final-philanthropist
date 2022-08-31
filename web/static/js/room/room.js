stageMap = {"-2":"未加入房间", "-1":"游戏未开始", "0":"游戏开始", "1":"回合开始", "2":"捐赠", "3":"捐赠结束", "4":"舆论惩罚", "5":"破产判定", "99":"回合结束", "6":"游戏结束"}
statusMap = {"0":notReadyTemplate, "1":readyTemplate, "2":lifeTemplate, "3":bankruptcyTemplate}
colorList = ["#2563eb", "#eab308", "#dc2626", "#475569",
    "#fef9c3", "#f0abfc", "#0f172a", "#d6d3d1" ]
currentStage = "-2"
roomId = ""
gameId = ""
playerId = ""
// 1：游戏未开始，2：游戏已开始，3：游戏已结束
roomStatus = 0
//{player_id: '', user_id: 1, username: 'sdf', is_ready: false}
playerList = [];
playerHistoryList = []
roundInfo = {}

playerMoneyChartList = []
playerHistoryChartList = []
exampleHtml = ""

init = function() {
    currentStage = "-2"
    roomId = ""
    gameId = ""
    playerId = ""
    roomStatus = 0
    playerList = []
    roundInfo = {}
    moneyDatasets = []
    historyDatasets = []
    playerHistoryList = []
    playerMoneyChartList = []
    playerHistoryChartList = []
    exampleHtml = ""
    clearTfpChara()
}

tfp = function() {
    switch (roomStatus) {
        case 0:
        case 1:
            getRoomInfo()
            break
        case 2:
            getGameInfo()
            break
        case 3:
    }
    flushPage()
}

flushPage = function() {
    flushStage()
    flushPlayer()
    flushRoundInfo()
}

flushPlayer = function() {
    html = ""
    for (let roomUser of playerList) {
        if (roomUser.is_ready) {
            playerStatus = "1"
        } else {
            playerStatus = "0"
        }
        if (roomStatus === 2) {
            if (roomUser.bankrupt) {
                playerStatus = "3"
            } else {
                playerStatus = "2"
            }
        }
        username = roomUser.username?roomUser.username:"SILENT HOUSE"
        hero_name = roomUser.hero_name?roomUser.hero_name:"未选择英雄"
        current_money = roomUser.current_money?roomUser.current_money:0
        punishment_money = roomUser.punishment_money?"<p style='color:red;font-weight: bold;'>" + roomUser.punishment_money + "</p>":0
        donated_money = roomUser.donated_money?roomUser.donated_money:0
        donated_money = donated_money >=0?donated_money:"-"
        donated_money = "<p style='font-weight: bold;'>" + donated_money + "</p>"
        html += userTemplate.signMix("/game/tfp/static/image/room/silenthouse-black.png", roomUser.username,
            hero_name, current_money, statusMap[playerStatus], donated_money, punishment_money)
    }
    $('#player_list').html(html)
}

flushStage = function() {
    if (roomStatus === 1) {
        if (roomId) {
            currentStage = "-1"
        } else {
            currentStage = "-2"
        }
    }
    if (roomStatus === 3) {
        currentStage = "6"
    }
    $('#stage_msg').html(stageMap[currentStage])
}

flushRoundInfo = function() {
    $('#round_no').html(roundInfo["round_no"])
    $('#public_opinion').html(roundInfo["public_opinion"])
}

flushTime = function () {
    if (roomStatus !== 2) {
        return
    }
    timeRemaining = (roundInfo["stage_end_time"] - new Date().getTime()) / 1000
    if (timeRemaining < 0) {
        tfp()
        timeRemaining = 0
    }
    if (!timeRemaining) {
        timeRemaining = 0
    }
    $('#time_remaining').html(Math.round(timeRemaining) + "秒")
}

getGameInfo = function() {
    $.ajax({
        type:"get",
        url:"/game/tfp/game/info?user_id=" + getUserId(),
        dataType:"json",
        cache:false,
        async:true,
        success:syncGameInfo
    });
}

function syncGameInfo(result) {
    if (checkTFPResult(result, false)) {
        gameId = result["data"]["game_id"]
        roomId = result["data"]["room_id"]
        roomStatus = Number(result["data"]["status"])
        roundInfo["round_no"] = Number(result["data"]["game_info"]["round_info"]["round_no"])
        roundInfo["public_opinion"] = Number(result["data"]["game_info"]["round_info"]["public_opinion"])
        currentStage = result["data"]["game_info"]["round_info"]["stage"]["stage"]
        playerId = result["data"]["game_info"]["player_game_info"]["player_id"]
        startTimeUnix = Number(result["data"]["game_info"]["round_info"]["stage"]["start_time"])
        durationUnix = Number(result["data"]["game_info"]["round_info"]["stage"]["duration"]) * 1000
        roundInfo["stage_end_time"] = startTimeUnix + durationUnix

        playerList = []
        for (let roomUser of result["data"]["game_info"]["round_info"]["player_info_list"]) {
            playerList.push({
                player_id: roomUser.player_id,
                username: roomUser.username,
                hero_name: roomUser.hero_name,
                current_money:roomUser.current_money,
                donated_money:roomUser.donated_money,
                punishment_money:roomUser.punishment_money,
                bankrupt:roomUser.bankrupt,
            })
        }

        playerHistoryList = []
        for (let roomUser of result["data"]["game_info"]["round_history_list"]) {
            playerHistoryList.push({
                player_id: roomUser.player_id,
                username: roomUser.username,
                hero_name: roomUser.hero_name,
                current_money_list:roomUser.current_money_list,
                donated_money_list:roomUser.donated_money_list,
                punishment_money_list:roomUser.punishment_money_list,
                bankrupt_list:roomUser.bankrupt_list,
            })
        }
        syncCharts()
        flushPage()
    }
}

function syncCharts() {
    index = 0
    for (let player of playerHistoryList) {
        tempMoney = {
            label: player.username,
            backgroundColor: colorList[index],
            borderColor: colorList[index],
            data: player.current_money_list,
            fill: false,
            index: index,
        }
        tempHistory = {
            label: player.username,
            backgroundColor: colorList[index],
            borderColor: colorList[index],
            data: player.donated_money_list,
            fill: false,
            index: index,
        }
        playerMoneyChartList.push(tempMoney)
        playerHistoryChartList.push(tempHistory)
        index++
    }
}

function syncRoomInfo(result) {
    // {"code":0,"data":{"game_id":"","room_id":"R18047856779410","status":1,"room_user_list":[{"player_id":"","user_id":1,"is_ready":false}]},"msg":""}
    if (checkTFPResult(result, false)) {
        gameId = result["data"]["game_id"]
        roomId = result["data"]["room_id"]
        roomStatus = Number(result["data"]["status"])
        playerList = []
        for (let roomUser of result["data"]["room_user_list"]) {
            //{player_id: '', user_id: 1, username: 'sdf', is_ready: false}
            playerList.push({
                player_id: roomUser.player_id,
                user_id: roomUser.user_id,
                username: roomUser.username,
                is_ready: roomUser.is_ready
            })
        }
        playerList.sort(comparePlayer)
        flushPage()
    }
}

getRoomInfo = function() {
    $.ajax({
        type:"get",
        url:"/game/tfp/room/info?user_id=" + getUserId(),
        dataType:"json",
        cache:false,
        async:true,
        success:syncRoomInfo
    });
}

$(function(){
    $('#game_start').bind('click', function () {
        $.ajax({
            type:"post",
            url:"/game/tfp/room/join/random",
            data:JSON.stringify({"user_id": getUserId()}),
            dataType:"json",
            cache:false,
            async:true,
            contentType: "application/json",
            success:syncRoomInfo
        });
    });
})

$(function(){
    $('#ready').bind('click', function () {
        $.ajax({
            type:"post",
            url:"/game/tfp/room/ready",
            data:JSON.stringify({"user_id": getUserId()}),
            dataType:"json",
            cache:false,
            async:true,
            contentType: "application/json",
            success:function (result) {
                if (checkTFPResult(result)) {
                    tfp()
                }
            }
        });
    });
})

$(function(){
    $('#ready_cancel').bind('click', function () {
        $.ajax({
            type:"post",
            url:"/game/tfp/room/ready/cancel",
            data:JSON.stringify({"user_id": getUserId()}),
            dataType:"json",
            cache:false,
            async:true,
            contentType: "application/json",
            success:function (result) {
                if (checkTFPResult(result)) {
                    tfp()
                }
            }
        });
    });
})

$(function(){
    $('#exit_room').bind('click', function () {
        if (roomStatus === 2) {
            showErrorDialog("游戏进行中")
            return
        }
        $.ajax({
            type:"post",
            url:"/game/tfp/room/exit",
            data:JSON.stringify({"user_id": getUserId()}),
            dataType:"json",
            cache:false,
            async:true,
            contentType: "application/json",
            success:function (result) {
                if (checkTFPResult(result)) {
                    init()
                    tfp()
                }
            }
        });
    });
})

$(function(){
    $('#donated_btn').bind('click', function () {
        $.ajax({
            type:"post",
            url:"/game/tfp/game/donated",
            data:JSON.stringify({"user_id": getUserId(), "round_no": Number(roundInfo["round_no"]), "player_id": playerId, "donated": Number($("#donated_money").val())}),
            dataType:"json",
            cache:false,
            async:true,
            contentType: "application/json",
            success:function (result) {
                if (checkTFPResult(result)) {
                    tfp()
                }
            }
        });
    });
})

$(function(){
    $('#logout').bind('click', function () {
        logout()
    });
})

tfp()
setInterval(tfp,2000)
setInterval(flushTime,500)
