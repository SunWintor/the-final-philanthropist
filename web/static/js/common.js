
getToken = function() {
    token = $.cookie('x-Token')
    console.log(token)
    if (!token) {
        $(location).attr('href', '/login');
    }
    return token
}

getUserId = function() {
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

logout = function() {
    $.removeCookie('x-Token', { path: '/' });
    $.removeCookie('user_id', { path: '/' });
    $(location).attr('href', '/login');
}

checkTFPResult = function(result, showDialog = true) {
    if (result["code"] == 0) {
        return true
    }
    if (showDialog) {
        showErrorDialog(result["msg"])
        return false
    }
}

var m_DialogGroup = [],
m_Start = false;
showErrorDialog = function(msg){
    showDialogBox(msg, "error", 250, 60);
    if (m_Start == false) {
        m_Start = true;
        removeDialog();
    }
}

//计时器，按着队列一个一个消失
function removeDialog() {
    setInterval(() => {
        if ($(".dialog_err").length > 0) {
            m_DialogGroup.shift();
            $(".dialog_err")[0].remove();
            Array.from($(".dialog_err")).forEach(element => {
                element.style.top = (parseInt(element.style.top) - 80) + 'px';
            })
            for (let i = 0; i < m_DialogGroup.length; i++) {
                m_DialogGroup[i] -= 80;
            }
        }
    }, 2000);
}

//消息提示组件
function showDialogBox(msg, type, width, height) {
    m_DialogGroup[0] = -80;
    let m_Top = m_DialogGroup.slice(-1)[0] + 80;
    m_DialogGroup.push(m_Top);
    let m_Dialog = document.createElement("div");
    m_Dialog.style.height = height + "px";
    m_Dialog.style.width = width + "px";
    m_Dialog.className = "dialog_err";
    m_Dialog.innerText = msg;
    m_Dialog.setAttribute("type", "error");
    $("body").append(m_Dialog);
    setTimeout(() => {
        m_Dialog.style.opacity = 1;
        m_Dialog.style.top = m_Top + 'px';
    }, 0);
}