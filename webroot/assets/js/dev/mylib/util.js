define('util', ['jquery'], function(){
	var util = {}

	util.language = function(){
		var lang = util.getCookie('lang');
		if(!lang){
			lang = navigator.language || navigator.userLanguage;
		}
		return lang.replace('_', '-').toLowerCase();
	}

	util.DATE_DAY = 86400000;
	util.DATE_HOUR = 3600000;
	util.humanTime = function (b, e){
	    if (typeof(b)=='string')
	        b = util.Str2date(b);
	    if (typeof(e)=='string')
	        e = util.Str2date(e);
	    
	    var n,s,str,diff = Math.abs(e.getTime() - b.getTime())/1000;
	    if (diff >= 31536000) {
	        n = diff / 31536000;
	        s = Math.round((n - Math.floor(n)) * 12);
	        str = parseInt(n) + '年' + (s==0?'':s +'个月') ;
	    }else if (diff >= 2592000) {
	        n = diff / 2592000;
	        s = Math.round((n - Math.floor(n)) * 30);
	        str = parseInt(n) + '个月' + (s==0?'':s +'天') ;
	    }else if (diff >= 86400) {
	        n = diff / 86400;
	        s = Math.round((n - Math.floor(n)) * 24);
	        str = parseInt(n) + '天' + (s==0?'':s +'小时') ;
	    } else if (diff >= 3600) {
	        n = diff / 3600;
	        s = Math.round((n - Math.floor(n)) * 60);
	        str = parseInt(n) + '小时'+(s==0?'':s+'分钟');
	    } else if (diff >= 60) {
	        n = diff / 60;
	        str = parseInt(n) + '分钟';
	    } else if (diff < 60) {
	        str = parseInt(diff).toString() + '秒';
	    }
	    return str;
	}

	util.drawingActiveStatus = function(item, appData){
		var userId = appData.user.id;
  		var obj = {
  			is_js:false, is_xmjl:false, is_xmgl:false, is_zt:false, 
  			finish_js:false, finish_xmjl:false, finish_xmgl:false, finish_zt:false};

  		if(item.xmjl_id ==userId){
  			obj.finish_xmjl = item.is_xmjl_sign;
  			obj.is_xmjl = true;
  		}

  		if(appData.draw_js_user_ids.indexOf(userId.toString()) > -1){
  			obj.finish_js = item.js_sign_by>0;
  			obj.is_js = true;
  		}

  		if(appData.draw_sw_user_ids.indexOf(userId.toString()) > -1){
  			obj.finish_sw = item.sw_sign_by>0;
  			obj.is_sw = true;
  		}

  		if(appData.draw_xmgl_user_ids.indexOf(userId.toString()) > -1){
  			obj.finish_xmgl = item.xmgl_sign_by>0;
  			obj.is_xmgl = true;
  		}

  		if(appData.draw_zt_user_ids.indexOf(userId.toString()) > -1){
  			obj.finish_zt = item.zt_sign_by>0;
  			obj.is_zt = true;
  		}

  		return obj;
  	}

	util.userAvatar = function (s){
		if(!s || s==''){
			s = '/assets/img/avatar_empty.png';
		}else{
			s = '/upload/avatar/' + s
		}
		return s;
	}

	util.setTitle = function (title){
		if(!title){
			title = ''
		}
		$('#gos-headerTitle').text(title);
		document.title = title;
	}

	util.str2date = function (str){
		str = str.replace(/[A-Za-z]/g, ' ').substr(0,19)
	    // var d = new Date(Date.parse(str));
	    // d.setTime(d.getTime() - d.getTimezoneOffset()*60000)
	    return new Date(Date.parse(str));;
	}
	
	util.date2str = function(time, ctype){
		if(!time){
			return '';
		}
		switch(typeof(time)){
			case 'number':
				time = new Date(time);
				break;
			case 'string':
				time = this.str2date(time);
				break;
		}

		if(ctype && ctype=='time')
			return time.getFullYear()+'-'+this.lpad(time.getMonth()+1, '0', 2)+'-'+this.lpad(time.getDate(), '0', 2)+' '+this.lpad(time.getHours(), '0', 2)+':'+this.lpad(time.getMinutes(), '0', 2)
		else
	    	return time.getFullYear()+'-'+this.lpad(time.getMonth()+1, '0', 2)+'-'+this.lpad(time.getDate(), '0', 2)
	}

	util.getUrlParameter = function (sParam){
	    var sPageURL = window.location.search.substring(1);
	    var sURLVariables = sPageURL.split('&');
	    for (var i = 0; i < sURLVariables.length; i++) 
	    {
	        var sParameterName = sURLVariables[i].split('=');
	        if (sParameterName[0] == sParam) 
	        {
	            return sParameterName[1];
	        }
	    }
	};
	
	util.cipherString = function(rsaData, nick, pwd){
		var rsa = new RSAKey(),
			ts = Server.getTime().toString(),
			userkey = CryptoJS.MD5( ts + nick )
		rsa.setPublic(rsaData.hex, '10001');
		
		var cipher = rsa.encrypt(util.lpad(ts, '0', 16)+userkey.toString(CryptoJS.enc.Base64)),
			text = nick + "|" +pwd;
		
		var aesCipher = util.aesEncrypto(text, ts, userkey);

		var s = rsaData.keyid.toString()+"|"+
				CryptoJS.enc.Hex.parse(cipher.toString()).toString(CryptoJS.enc.Base64)+"|"+
				aesCipher.toString();

		return s;
	};

	util.aesEncrypto = function(text, ts, key){
		ts = ts.toString()
	    var iv  = CryptoJS.MD5(util.lpad(ts, '0', 16)),
	    	encrypted = CryptoJS.AES.encrypt(text, key, { iv: iv })
	    
		return encrypted.ciphertext.toString(CryptoJS.enc.Base64)
	};
	
	util.aesDecrypto = function(src, ts, key){
		ts = ts.toString()
	    var iv  = CryptoJS.MD5(util.lpad(ts, '0', 16)),
	    	obj = {
				ciphertext: CryptoJS.enc.Base64.parse(src),
				salt: ""
			}
	    	decrypted = CryptoJS.AES.decrypt(obj, key, { iv: iv })
		return decrypted.toString(CryptoJS.enc.Utf8)
	};

	util.lpad = function(str, padString, l) {
	    while (str.toString().length < l)
	        str = padString + str;
	    return str;
	};
	 
	//pads right
	util.rpad = function(str, padString, l) {
	    while (str.toString().length < l)
	        str = str + padString;
	    return str;
	};

    util.objectFind = function(field, value, data){
		if(!data)
			return null;

		var i,len = data.length
		for(i=0;i<len;i++)
			if(data[i] && data[i][field] == value){
				return data[i];
			}
		return null;
	}

	util.objectFindIndex = function(field, value, data){
		if(!data)
			return -1;

		var i,len = data.length
		for(i=0;i<len;i++)
			if(data[i] && data[i][field] == value){
				return i;
			}
		return -1;
	}

    util.setCookie = function (name,value,days) {
        if (days) {
            var date = new Date();
            date.setTime(date.getTime()+(days*86400000));
            var expires = "; expires="+date.toGMTString();
        }
        else var expires = "";
        document.cookie = name+"="+value+expires+"; path=/";
    };

	util.getCookie = function (name) {
        var nameEQ = name + "=";
        var ca = document.cookie.split(';');
        for(var i=0;i < ca.length;i++) {
            var c = ca[i];
            while (c.charAt(0)==' ') c = c.substring(1,c.length);
            if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
        }
        return null;
    };

    util.deleteCookie = function (name) {
        this.SetCookie(name,"",-1);
    };

    util.getNick = function(){
        if(this.currentNick)
            return this.currentNick

        var s = decodeURIComponent(this.getCookie('PUB_gosauth')),
            arr = s.split('|');
            
        if(arr.length>1){
            this.currentNick = arr[0]
        }else{
            this.currentNick = s
        }
        return this.currentNick;
    };

    util.getSecret = function(){
        return decodeURIComponent(this.getCookie('secret'))
    };

	util.scopeFormData = function(s){
		var item, o, oo, data = {};
		for(o in s){
			if(o.charAt(0)=='$') continue;
			
			item = s[o];
			if(!item) continue;

			if(typeof item.$$parentForm != 'undefined' && item.$valid){
				for(oo in item){
					if(oo.charAt(0)=='$') continue;

					if(item[oo] && typeof item[oo].$modelValue != 'undefined')
						data[oo] = item[oo].$modelValue;
					else
						data[oo] = item[oo].$viewValue;
				}
			}
		}
		return data;
	}

	return util;
});



