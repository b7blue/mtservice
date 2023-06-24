function validateForm(email, pw, pw1, Vericode) {
    
    if (email == "" || pw == "" || pw1 == "" || Vericode == "") {
        window.alert("输入不可留空！")
        return false
    } else if (pw1 != pw) {
        window.alert("两次输入的密码不一致，请重新确认！")
        return false
    } else {
        return true
    }
    // 多加一个非法字符检测
}

function signUp(form) {
    formData = new FormData(form);
    let email = formData.get("email")
    let pw = formData.get("pw")
    let pw1 = formData.get("pw1")
    let vericode = formData.get("vericode")
    let xsrf = formData.get("_xsrf")

    if (validateForm(email, pw, pw1, vericode) == true) {
        var xmlhttp;
        if (window.XMLHttpRequest) {
            xmlhttp = new XMLHttpRequest();
        }
        else {
            xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
        }
        data = "_xsrf="+xsrf+"&email="+email+"&pw="+pw+"&vericode="+vericode
        xmlhttp.open("POST", "/signup", true);
        xmlhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xmlhttp.send(data);


        xmlhttp.onreadystatechange = function () {
            if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                result = xmlhttp.responseText
                if (result == "OK") {
                    jumpAfter5()
                }else {
                    document.getElementById("signupResult").innerHTML = result;
                }
            }
        }

    }

}

function jumpAfter5() {
    var signupResult = document.getElementById("signupResult");
    var time = 5;
    timer();
    setInterval(timer, 1000)

    function timer() {
        if (time == 0) {
            location.href = '/login'

        } else {
            signupResult.innerHTML = "注册成功！将在" + time + '秒之后跳转到登陆页面...';
            time--;
        }
    }
}


function getVeriCode() {
    var xmlhttp;
    if (window.XMLHttpRequest) {
        // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
        xmlhttp = new XMLHttpRequest();
    }
    else {
        // IE6, IE5 浏览器执行代码
        xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
    }
    var param = "_xsrf=" + document.forms["myForm"]["_xsrf"].value + "&email=" + document.forms["myForm"]["email"].value
    xmlhttp.open("POST", "/sendVeriCode", true);
    xmlhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xmlhttp.send(param);


    xmlhttp.onreadystatechange = function () {
        if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
            window.alert(xmlhttp.responseText)
            // document.getElementById("getVeriCodeResult").innerHTML = xmlhttp.responseText;
        }
    }

}

// function to_stu_info() {

//     document.getElementById("base_info_form").id = "base_info_form_folded"
//     document.getElementById("stu_info_form_folded").id = "stu_info_form"
// }
