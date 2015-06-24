define(
    'ajax',
    
    ['util', 'jquery'],

    function(util) {
        var ajax  = {
            NewClient : function(path, timeout){
                return new Client(path, timeout);
            },

            serverTime : {
                diff : 0,
                set : function(unix){
                    this.diff = unix>0 ? (new Date()).getTime() - unix*1000 : 0;
                },
                time : function(){
                    return (new Date()).getTime() - this.diff;
                }
            },

            getHostName : function(){
                var hostArr = window.location.host.split('.')
                return hostArr[hostArr.length-2]+'.'+hostArr[hostArr.length-1]
            },

            getUrlVars: function(){
                var vars = [], hash;
                var hashes = window.location.href.slice(window.location.href.indexOf('?') + 1).split('&');
                for(var i = 0; i < hashes.length; i++) {
                hash = hashes[i].split('=');
                vars.push(hash[0]);
                vars[hash[0]] = hash[1];
                }
                return vars;
            },

            datasetDecode : function(data){
                if(!data || data.length==0)
                    return data

                var fields=data[0], l = data.length, obj, value, items=[]
                for(var i=1; i < l; i++) {
                    obj = new Object()
                    for(var k=0; k < fields.length; k++)
                        obj[fields[k]] = data[i][k]
                    items.push(obj)
                }
                return items;
            },

        };

        // ajax.setCookie('lang', navigator.language || navigator.userLanguage);
        var Client = function(path, timeout){
            this.path = path || "/api/web";
            this._timeout = timeout || 10000;
            this._button = null;
            this._block = null;
        }
        
        Client.prototype.errorHandler = function(r){console.log(r);};

        Client.prototype.button = function($t){
            this._button = $t;
            return this;
        }
        Client.prototype.block = function($t){
            this._block = $t;
            return this;
        }

        Client.prototype.sendAlone = function(method, args){
            if(this.deferred && this.deferred.state()=='pending'){
                var func = function(){}
                return {done:func, fail:func, always:func};
            }
            
            this.deferred = this.send(method, args);
            return this.deferred;
        };

        Client.prototype.send = function(method, args){
            var pm = {method:method, args:args?args:null};
            var options = {type:'POST', dataType:'json',cache:false, timeout: this._timeout},k = null

            if (this._block){
                this.doBusy(this._block, true);
            }
                
            if(this._button){
                this.doDisable(this._button, true);
            }

            var clas = this;
                deferred = $.ajax({
                    url:        this.path,
                    type:       options['type'],
                    dataType:   options['dataType'],
                    cache:      options['cache'],
                    //jsonp: 'callback',
                    data:       {src: JSON.stringify(pm)},
                    timeout:    options['timeout']
                });

            deferred.always(function(){
                if(clas._block) {
                    clas.doBusy(clas._block, false);
                }

                if(clas._button){
                    clas.doDisable(clas._button, false);
                }

            }).fail(function (jqXHR, textStatus, errorThrown){
                switch(jqXHR.status){
                    case 599:
                        if(clas.errorHandler){
                            try{
                                clas.errorHandler(JSON.parse(jqXHR.responseText))
                            }catch(e){
                                console.log(e)
                            }
                        }
                        break;
                    case 404:
                        alert(textStatus+': api not found!')
                        break;
                    default:
                        alert(textStatus+': '+jqXHR.responseText)
                        break;
                }
            });

            return deferred;
        }

        Client.prototype.doDisable = function(el, sw){
            if(typeof el == "string")
                el  = $(el);

            el.attr('disabled', sw ? 'disabled' : null);
        }

        Client.prototype.doBusy = function(el, sw){
            var container,overlay;
            
            if(typeof el == "string")
                container  = $(el);
            else
                container  = el;

            var options = {
                bgColor         : '#CCCCCC',
                duration        : 200,
                opacity         : 0.5
            }
            container.each(function(){
                var $this = $(this);
                if(sw){
                    overlay = $('<div></div>').css({
                            'background-color': options.bgColor,
                            'opacity':options.opacity,
                            'width':$this.width(),
                            'height':$this.height(),
                            'position':'absolute',
                            'top':'0px',
                            'left':'0px',
                            'z-index':9999
                    })
                    overlay = $('<div class="block-overlay"></div>').css({
                            'position': 'relative'
                    }).append(overlay)

                    $this.prepend(
                        overlay.append('<div class="bloack-ui"></div>').fadeIn(options.duration)
                    );
                }else{
                    overlay = $this.children(".block-overlay");
                    if (overlay.length>0) {
                        overlay.remove();
                    }
                }
            })
        };


        var transitionEnd = function () {
            var el = document.createElement('bootstrap')

            var transEndEventNames = {
              WebkitTransition : 'webkitTransitionEnd',
              MozTransition    : 'transitionend',
              OTransition      : 'oTransitionEnd otransitionend',
              transition       : 'transitionend'
            }

            for (var name in transEndEventNames) {
              if (el.style[name] !== undefined) {
                return { end: transEndEventNames[name] }
              }
            }

            return false // explicit for ie8 (  ._.)
        }

        $.fn.emulateTransitionEnd = function (duration) {
            var called = false
            var $el = this
            $(this).one('bsTransitionEnd', function () { called = true })
            var callback = function () { if (!called) $($el).trigger($.support.transition.end) }
            setTimeout(callback, duration)
            return this
        }

        $(function () {
            $.support.transition = transitionEnd()

            if (!$.support.transition) return;

            $.event.special.bsTransitionEnd = {
                bindType: $.support.transition.end,
                delegateType: $.support.transition.end,
                handle: function (e) {
                    if ($(e.target).is(this)) return e.handleObj.handler.apply(this, arguments)
                }
            }
        })
        
        return ajax;
    }
)
