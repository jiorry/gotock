define('loader', ['util', 'jquery'], function(util){
	var $goBack = $('#gos-goback'),
		$header = $('#gos-headerBarR');
	var loader = function(){
		this.defaultViewName = 'home.module';

		var $tViews = $('#gos-Views');
		
		this.getViewName = function ($location){
			var name = $location.routerMatch ? $location.routerMatch.name : '';
			return name.replace(/\//g, '.')
		}

		this.buildHeader = function(m){
			if(m.goBackButton == false){
				$goBack.hide();
			}else{
				$goBack.show();
			}

			if(m.headerHtml){
				$header.html(m.headerHtml);
			}else{
				$header.empty();
			}

   			util.setTitle(m.title);
		}

		this.loadview = function(name, templateUrl, $q, $rootScope){
			var pathConfig = {},
				defer = $q.defer(),
				thisClas = this,
				path = templateUrl.substr(0, templateUrl.indexOf('.html'));

			pathConfig[name] = path;

			require.config({paths: pathConfig});
			require([name], function (m){
				if(m && m.i18n){
					var lang = $rootScope.language;
					lang = lang==''?'en-us':lang;
					window._gos.transData = window._gos.transData || {}

					if(window._gos.transData[name]){
						defer.resolve();
					}else{
						
						var arr = name.split('.'),
							dataUrl = path.substr(0, path.indexOf(arr.join('/'))) + '/' + arr[0] + '/locales/'+lang+'/'+arr[1]+'.json?v='+JsVersion;
						$.getJSON(dataUrl)
							.success(function(result){
								window._gos.transData[name] = result;
								defer.resolve();

							}).fail(function(){
								lang = 'en-us';
								dataUrl = path.substr(0, path.indexOf(arr.join('/'))) + '/' + arr[0] + '/locales/'+lang+'/'+arr[1]+'.json?v='+JsVersion;
								$.getJSON(dataUrl)
									.success(function(result){
										window._gos.transData[name] = result;
										defer.resolve();

									})
							})
					}
				}else{
					defer.resolve();
				}

				thisClas.buildHeader(m);

			}, function(err){
				defer.reject ();
			});

			return defer.promise;
		}
	};

	return new loader();
});



