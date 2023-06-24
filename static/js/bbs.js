function validate() {
    let name = newmes.name.value
    let content = newmes.content.value
    if (name == "" || content == "") {
        window.alert("输入不能留空")
        return false
    }else {
        return true
    }
}