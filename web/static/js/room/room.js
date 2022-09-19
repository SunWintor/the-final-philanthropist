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
    console.time("a")
    getRoomInfo()
    console.timeEnd("a")
}

flushPage = function() {
    stageMap["-1"] = "房间编号:" + roomId
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
        username = username + "(R:" + roomUser.ranking
        if (roomUser.ranking_up) {
            username = username + rankingTemplate.signMix(roomUser.ranking_up>0?"red":"green", roomUser.ranking_up>0?"+"+roomUser.ranking_up:roomUser.ranking_up)
        }
        username = username + ")"
        hero_name = roomUser.hero_name?roomUser.hero_name:"未选择英雄"
        current_money = roomUser.current_money?roomUser.current_money:0
        punishment_money = roomUser.punishment_money?"<p style='color:red;font-weight: bold;'>" + roomUser.punishment_money + "</p>":0
        donated_money = roomUser.donated_money?roomUser.donated_money:0
        donated_money = donated_money >=0?donated_money:"-"
        donated_money = "<p style='font-weight: bold;'>" + donated_money + "</p>"
        html += userTemplate.signMix("/game/tfp/static/image/room/silenthouse-black.png", username,
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

function syncGameInfo(gameInfo) {
    if (gameInfo == null || !gameInfo) {
        return
    }
    if (Number(gameInfo["status"]) === 3 && roomStatus !== 2) {
        return
    }
    gameId = gameInfo["game_id"]
    roomId = gameInfo["room_id"]
    roomStatus = Number(gameInfo["status"])
    roundInfo["round_no"] = Number(gameInfo["game_info"]["round_info"]["round_no"])
    roundInfo["public_opinion"] = Number(gameInfo["game_info"]["round_info"]["public_opinion"])
    currentStage = gameInfo["game_info"]["round_info"]["stage"]["stage"]
    playerId = gameInfo["game_info"]["player_game_info"]["player_id"]
    startTimeUnix = Number(gameInfo["game_info"]["round_info"]["stage"]["start_time"])
    durationUnix = Number(gameInfo["game_info"]["round_info"]["stage"]["duration"]) * 1000
    roundInfo["stage_end_time"] = startTimeUnix + durationUnix

    playerList = []
    for (let roomUser of gameInfo["game_info"]["round_info"]["player_info_list"]) {
        playerList.push({
            player_id: roomUser.player_id,
            username: roomUser.username,
            hero_name: roomUser.hero_name,
            current_money:roomUser.current_money,
            donated_money:roomUser.donated_money,
            punishment_money:roomUser.punishment_money,
            bankrupt:roomUser.bankrupt,
            ranking: roomUser.ranking,
            ranking_up: roomUser.ranking_up,
        })
    }

    playerHistoryList = []
    for (let roomUser of gameInfo["game_info"]["round_history_list"]) {
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
}

function syncCharts() {
    index = 0
    playerMoneyChartList = []
    playerHistoryChartList = []
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

function syncRoomInfoJoin(result) {
    if (result["code"] === 10004) {
        init()
        flushPage()
        return
    }
    if (checkTFPResult(result, true)) {
        syncRoomInfo(result)
    }
}

function syncRoomInfo(result) {
    if (result["code"] === 10004) {
        init()
        flushPage()
        return
    }
    if (checkTFPResult(result, false)) {
        gameId = result["data"]["game_id"]
        roomId = result["data"]["room_id"]
        roomStatus = Number(result["data"]["status"])
        playerList = []
        for (let roomUser of result["data"]["room_user_list"]) {
            playerList.push({
                player_id: roomUser.player_id,
                user_id: roomUser.user_id,
                username: roomUser.username,
                is_ready: roomUser.is_ready,
                ranking: roomUser.ranking,
                ranking_up: roomUser.ranking_up,
            })
        }
        playerList.sort(comparePlayer)
        syncGameInfo(result["data"]["game_info"])
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
            success:syncRoomInfoJoin
        });
    });
})

$(function(){
    $('#join_new').bind('click', function () {
        $.ajax({
            type:"post",
            url:"/game/tfp/room/join/new",
            data:JSON.stringify({"user_id": getUserId()}),
            dataType:"json",
            cache:false,
            async:true,
            contentType: "application/json",
            success:syncRoomInfoJoin
        });
    });
})

$(function(){
    $('#join_room_id').bind('click', function () {
        $.ajax({
            type:"post",
            url:"/game/tfp/room/join/room_id",
            data:JSON.stringify({"user_id": getUserId(), "room_id": $("#room_id").val()}),
            dataType:"json",
            cache:false,
            async:true,
            contentType: "application/json",
            success:syncRoomInfoJoin
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

setInterval(tfp,2000)
setInterval(flushTime,500)
tfp()
