require.config({
	waitSeconds :100,
	baseUrl : "/assets/js/",
	paths: {
		'util' : MYENV+'/mylib/util',
		'ajax' : MYENV+'/mylib/ajax'
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

		client.bindClick($('#btn-login'), function(){
			var nick = $('#inputNick').val(),
				password = $('#inputPassword').val();

			client.send('public.sign.UserLogin', {cipher : util.cipherString(rsaData, nick, password)})
				.done(function(result){
					window.location.href = "/"

				}).fail(function(jqXHR){
					var err = JSON.parse(jqXHR.responseText)
					doError(err.message);
				})
		})

		$('#inputPassword').keypress(function(e){
			if (e.which == 13) {
			   $('#btn-login').trigger('click');
			}
		})

});
