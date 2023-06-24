function validateForm(href) {
    if (href == "") {
        window.alert("输入不可留空！")
        return false
    }else {
        return true
    }
    // 多加一个非法字符检测
}

function post2TXT(form) {
    formData = new FormData(form);
    let href = escape(formData.get("href"))
    let xsrf = formData.get("_xsrf")

    if (validateForm(href) == true) {
        var xmlhttp;
        if (window.XMLHttpRequest) {
            xmlhttp = new XMLHttpRequest();
        }
        else {
            xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
        }
        data = "_xsrf="+xsrf+"&href="+href
        xmlhttp.open("POST", "/txtdownload", true);
        xmlhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xmlhttp.send(data);


        xmlhttp.onreadystatechange = function () {
            if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                result = JSON.parse(xmlhttp.responseText)
                if (result.err == "") {
                    // 将下载框可视化，填充文件名
                    document.getElementById("downloadbox").style.display = "block";
                    document.getElementById("filename").innerHTML = result.filename;
                    document.getElementById("filenameinput").value = result.filename;
                }else {
                    // 弹窗显示错误原因
                    window.alert("发生错误:"+result.err)
                }
            }
        }

    }

}

