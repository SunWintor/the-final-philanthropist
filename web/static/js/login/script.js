$(".toggle").on("click", function () {
  $(".container").stop().addClass("active");
});

$(".close").on("click", function () {
  $(".container").stop().removeClass("active");
});

$(function (){
  $('#login').bind('click', function () {
    uname = $("#login_uname").val()
    pwd = $("#login_pwd").val()
    $.ajax({
      type:"post",    // 请求类型
      url:"/auth/login",    // 请求URL
      data:JSON.stringify({"username": uname, "password": pwd}),    // 请求参数 即是 在Servlet中 request.getParement();可以得到的字符串
      dataType:"json",    // 数据返回类型
      cache:false, // 是否缓存
      async:true,    // 默认为true 异步请求
      contentType: "application/json",
      success:function(result){    // 成功返回的结果(响应)
        console.log(JSON.stringify(result))
        if (result["code"] == 0) {
          setTFPLoginCookies(result["data"])
          $(location).attr('href', '/room');
        }
        $("#login_msg").text(JSON.stringify(result["msg"]));
      }
    });
  });
})
$(function (){
  $('#reg').bind('click', function () {
    uname = $("#reg_uname").val()
    pwd = $("#reg_pwd").val()
    repwd = $("#reg_repwd").val()
    if (pwd != repwd) {
      $("#reg_msg").text("两次输入的密码不一致");
      return
    }
    $.ajax({
      type:"post",    // 请求类型
      url:"/auth/register",    // 请求URL
      data:JSON.stringify({"username": uname, "password": pwd}),    // 请求参数 即是 在Servlet中 request.getParement();可以得到的字符串
      dataType:"json",    // 数据返回类型
      cache:false, // 是否缓存
      async:true,    // 默认为true 异步请求
      contentType: "application/json",
      success:function(result){    // 成功返回的结果(响应)
        console.log(JSON.stringify(result))
        if (result["code"] == 0) {
          setTFPLoginCookies(result["data"])
          $(location).attr('href', '/room');
        }
        $("#reg_msg").text(JSON.stringify(result["msg"]));
      }
    });
  });
})

token = getToken()
if (token) {
  $(location).attr('href', '/room');
}