
getToken = function () {
    token = $.cookie('x-Token')
    console.log(token)
    if (!token) {
        $(location).attr('href', '/login');
    }
    return token
}

getUserId = function () {
    userId = $.cookie('user_id')
    if (!userId) {
        $(location).attr('href', '/login');
    }
    return Number(userId)
}

setTFPLoginCookies = function (userInfo) {
    $.cookie('x-Token', userInfo["token"], { expires: 30 });
    $.cookie('user_id', userInfo["id"], { expires: 30 });
}