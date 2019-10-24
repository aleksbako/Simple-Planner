var xmlhttp = new XMLHttpRequest();   // new HttpRequest instance 
var theUrl = "localhost:8080/event";
xmlhttp.open("POST", theUrl);
xmlhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

