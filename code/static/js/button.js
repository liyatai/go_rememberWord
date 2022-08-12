//下一个单词
function next(){
    location.href="/word/next"
}

// 选择单词组
function select(){

    var obj=document.getElementById("so")
    var key=obj.value
     location.href="/word/select?key="+key
}
//显示单词组的中文
function show(){
    document.getElementById("chinese").removeAttribute("hidden")
}
//删除单词
function del(){
    var s=document.getElementById("english").innerText
    location.href="/word/delete?eng="+s
}
//播放音频
function musicPlay(){
    //获取英文原文
    var eng=document.getElementById("english").innerText
    //获取音频组件对象
    var sss=document.getElementById("music");
    var str="http://dict.youdao.com/dictvoice?audio="+eng+"&type=2";
    //重定向
    sss=new Audio(str);
    //播放
    sss.play();
}
//添加后台记录
function add(){
    var name =document.getElementById("so").value
    var level =document.getElementById("level").value
    location.href="/admin/add?name="+name+"&&level="+level 
}
//后台记录排序
function order(){
    location.href="/admin/order"
}
//返回主页
function exit(){
    location.href="/word/index"
}
//跳转后台
function jumpAdmin(){
    location.href="/admin/index"
}