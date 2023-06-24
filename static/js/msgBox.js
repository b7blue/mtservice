function delMsg(form) {
    formData = new FormData(form);
    let id = escape(JSON.stringify(formData.getAll("id")))
    let xsrf = formData.get("_xsrf")
    window.alert(id);


    if (id.length != 0) {
        var xmlhttp;
        if (window.XMLHttpRequest) {
            xmlhttp = new XMLHttpRequest();
        }
        else {
            xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
        }
        data = "_xsrf=" + xsrf + "&id=" + id
        xmlhttp.open("POST", "/msgBox", true);
        xmlhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xmlhttp.send(data);


        xmlhttp.onreadystatechange = function () {
            if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                result = xmlhttp.responseText
                if (result == "登陆状态已过期，请求失败，请重新登录") {
                    jumpAfter5()
                } else {
                    window.alert(result);
                }
            }
        }

    } else {
        window.alert("你还没有选中要删除的消息呢")
    }
}

function jumpAfter5() {
    var addSubResult = document.getElementById("addSubResult");
    var time = 5;
    timer();
    setInterval(timer, 1000)

    function timer() {
        if (time == 0) {
            location.href = '/login'

        } else {
            addSubResult.innerHTML = "登陆状态已过期，请求失败，请重新登录！将在" + time + '秒之后跳转到登陆页面...';
            time--;
        }
    }
}