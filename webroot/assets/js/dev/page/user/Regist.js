require.config({
	waitSeconds :100,
	baseUrl : "/assets/js/",
	paths: {
		'util' : MYENV+'/mylib/util',
		'ajax' : MYENV+'/mylib/ajax',
	},
	shim: {
        'jquery' : {
        	exports:'$'
        }
	}
});

require(
	['ajax', 'util', 'crypto'],
	function (ajax, util){
		var client = ajax.NewClient("/api/open");
		client.send('public.site.Rsakey', null)
			.done(function(result){
				rsaData = result;
			})

		function doError(s){
			alert(s);
		}

		client.bindClick($('#btn-regist'), function(){
			var nick = $('#inputNick').val(),
				password = $('#inputPassword').val(),
				confirm = $('#inputPasswordConfirm').val();

			if(password != confirm){
				doError('两次输入的密码不匹配');
				return
			}

			client.send('public.sign.Regist', {cipher : util.cipherString(rsaData, nick, password)})
				.done(function(result){
					window.location.href = "/login"

				}).fail(function(jqXHR){
					var err = JSON.parse(jqXHR.responseText)
					doError(err.message);
				})
		})

});
