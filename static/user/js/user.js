$(document).ready(function(){
	const URL='/GetUserMessage?'
	$.ajax({
		url: URL,
		type: "GET",
		success: function (result) {
			$("#loginuser").text(result.nick_name)
		}
	})
});