$(document).ready(function(){
	const URL='/UserLogin?'
	$("#login").click(function(event) {
		userManageUrl = URL + "account=" + $("#username").val() + "&password=" + $("#password").val()
		$.ajax({
			url: userManageUrl,
			type: "GET",
			success: function (result) {
				if(result === "Login Success") {
					window.location.href="/UserManage"
				}
			}
		})
	});
});