

define(
    "websock",
    
    ['util', 'crypto', 'jquery'],

    function(util) {
        var TYPE_CALL                 = 10,
            TYPE_MESSAGE_TO_USER      = 1,
            TYPE_MESSAGE_TO_USERS     = 2,
            TYPE_MESSAGE_TO_GROUP     = 3

        var sendFunc = function(conn, json){
                try{
                    conn.send(json)
                }catch(e) {
                    console.log(e);
                }
            },

            aesKey = function(ts){
                return CryptoJS.MD5(util.getSecret() + util.getNick() + util.lpad(ts, '0', 16))
            },

            messageTo = function(itype, to, message){
                var ts = Server.getTime(),
                    json = JSON.stringify({to:to, from:[util.getNick()], message:message}),
                    obj = ["1",itype.toString(),util.aesEncrypto(json, ts, aesKey(ts)),ts.toString()]
                return JSON.stringify(obj)
            },

            callJson = function(method, args){
                var ts = Server.getTime(),
                    json = JSON.stringify({method:method,args:args}),
                    obj = ["1",TYPE_CALL.toString(),util.aesEncrypto(json, ts, aesKey(ts)),ts.toString()]
                
                return JSON.stringify(obj)
            }

        var ws = {
            conn :null,
            open : function(url,callback){
                try{
                    var clas = this
                    this.conn = new WebSocket(url)
                    
                    this.conn.onopen = function(e){
                        callback(e);
                        clas.afterOpen(e);
                    }
                    this.conn.onclose = function(e){
                        clas.onClose(e);
                    }
                    this.conn.onerror = function(e){
                        clas.onError(e);
                    }
                    this.conn.onmessage = function(e){
                        var socketData = JSON.parse(e.data),
                            itype = parseInt(socketData[1]),
                            ts = socketData[3],
                            jsonStr='';

                        if(socketData[0]=="1"){
                            jsonStr = util.aesDecrypto(socketData[2], ts, aesKey(ts));
                        }else{
                            jsonStr = socketData[2];
                        }
                        if(jsonStr==''){
                            alert("返回数据为空！")
                            return
                        }
                        switch(itype){
                            case TYPE_CALL:
                                var callData = JSON.parse(jsonStr)
                                $(document).trigger('socket.'+callData.method, callData.args);
                                // clas.resultFunc[callData.method].call(clas.resultFunc, callData.args)
                                break;
                            case TYPE_MESSAGE_TO_USER:
                            case TYPE_MESSAGE_TO_USERS:
                                var messageData = JSON.parse(jsonStr)
                                console.log(messageData.message)
                                break;
                        }
                    }
                }catch(e) {
                    alert(e);
                }
            },
            afterOpen : function(e){
                console.log('initDo')
            },
            sendMessage : function(nick, message){
                sendFunc(this.conn, messageTo(TYPE_MESSAGE_TO_USER, [nick], message))
            },
            sendMessageToUsers : function(nicks, message){
                sendFunc(this.conn, messageTo(TYPE_MESSAGE_TO_USERS, nicks, message))
            },
            sendMessageToGroup : function(group, message){
                sendFunc(this.conn, messageTo(TYPE_MESSAGE_TO_GROUP, [group], message))
            },
            callErrorCount : 0,
            sendCall :function(method, args){
                console.log('sendCall', method, this.conn.readyState)
                if(this.conn.readyState == 0){
                    if(this.callErrorCount>100){
                        return;
                    }
                    var clas = this;
                    window.setTimeout(function(){
                        clas.callErrorCount++;
                        sendFunc(clas.conn, callJson(method, args));
                    }, 200)
                }else if(this.conn.readyState == 1){
                    this.callErrorCount = 0;
                    sendFunc(this.conn, callJson(method, args));
                }
            },

            close :function(){
                this.conn.close()
            },
            onClose : function(e){
                console.log('onClose')
                console.log(e)
            },
            onError : function(e){
                console.log('onError')
                console.log(e)
            }
            
        }

        ws.prepare = function(url, callback){
            if(ws.conn && ws.conn.readyState==1)
                return;
            else
                ws.open(url, callback)
        }

        return ws
    }
)
