require.config({baseUrl:"/assets/js/",paths:{util:MYENV+"/mylib/util",ajax:MYENV+"/mylib/ajax"},shim:{jquery:{exports:"$"}}}),require(["ajax","util","crypto"],function(e,t){function r(e){alert(e)}var n=e.NewClient("/api/open");n.send("public.site.Rsakey",null).done(function(e){rsaData=e}),n.bindClick($("#btn-regist"),function(){var e=$("#inputNick").val(),i=$("#inputPassword").val(),s=$("#inputPasswordConfirm").val();if(i!=s){r("两次输入的密码不匹配");return}n.send("public.sign.Regist",{cipher:t.cipherString(rsaData,e,i)}).done(function(e){window.location.href="/login"}).fail(function(e){var t=JSON.parse(e.responseText);r(t.message)})})});