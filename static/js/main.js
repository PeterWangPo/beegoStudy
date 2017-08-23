//login
function ajaxLogin(){
	var username = $("#username").val();
	var password = $("#password").val();
	if(!username || !password){
		return;
	}
	$.post('/account/login',{username:username,password:password},function(res){
		if(res.code == 1){
			// console.log(res);
			window.location.href = res.url;
		}else{
			alert(res.msg);
		}
	},'json');
}