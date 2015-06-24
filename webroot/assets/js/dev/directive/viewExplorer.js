define('ngViewExplorer', ['app', 'loader'], function(app, loader){
	function urlPath(s, isFull){
		var n = s.indexOf('#');
		if(n>0)
			s = s.substr(0, n);

		if(isFull){
			n = s.indexOf('?');
			if(n>0)
				s = s.substr(0, n);
		}
		return s
	}

	$(document).on('click', 'a[href]', function(){
		window._gos.openMode = $(this).attr('open-mode');
	})

	app.service('Viewstack',function(){
		this.getRouter = function(){
			return window._gos.viewstackRouter;
		}

		this.route = function(url){
			var vr = window._gos.viewstackRouter,
				item, i, m, k, fields,
				reg = new RegExp(':([a-zA-Z0-9_]+)', 'g');

			for (i = vr.length - 1; i >= 0; i--) {
				item = $.extend({}, vr[i]);
				item.search = {}
				fields = item.router.match(reg);

				if(item.router.indexOf(':')>0 && fields && fields.length>0){
					item.router = item.router.replace(/\/:\w+/g, '/(\\w+)');
					
					m = (new RegExp(item.router, 'g')).exec(url);
					
					if(!m) continue;

					for(k=1; m.length>k; k++){
						item.search[fields[k-1].substr(1)] = m[k];
					}

					return item;
				}else{
					if(item.router==urlPath(url))
						return item;
				}
			};

			return null;

		}

	});

	app.directive("ngViewExplorer",
		['$rootScope', '$q', '$animate', '$http', '$compile', '$controller', '$location', '$templateRequest', '$sce', 'Viewstack',
		function($rootScope, $q, $animate, $http, $compile, $controller, $location, $templateRequest, $sce, Viewstack){
			function compileFunc(element, attr, linker){
					return function(scope, $element, attr) {
						var currentElement,match,template,controller,viewname,
							ngGroup = attr.ngViewExplorer,
							newScope = scope.$new();

						$rootScope.$on('$locationChangeSuccess', function(evt, newUrl, oldUrl){
							if(newUrl==oldUrl)
								oldUrl = '';
							
							if(urlPath(newUrl)==urlPath(oldUrl)){
								return;
							}

							// require load js
							var url;
							if($location.$$html5){
								url = $location.$$path;
							}else{
								url = newUrl.match(/#(\/.*)/);
							}
							match = Viewstack.route(url);

							// router matched
							if (!match){
								console.log('error: no matched router:', url)
								return false;
							}
							viewname = match.name;

							$location.router = match;
							
							loader.loadview(viewname, match.templateUrl, $q, $rootScope)
								.then(function(){
									if(checkFunc(evt, newUrl, oldUrl)){
										console.log(viewname, 'loaded successfull');
										update(evt, newUrl, oldUrl);
									}
								}, function(){
									console.log(viewname, 'failed loaded');
								})
						});

						function checkFunc(evt, newUrl, oldUrl){
							var router = Viewstack.getRouter();
							if(!newUrl || !router || newUrl==oldUrl){
								return false;
							}

							match.group = match.group || '';
							if(ngGroup != match.group){
								return;
							}

							controller = match.controller;

							if(currentElement){
								var isFound = false;
								currentElement.children().each(function(){
									var $t =$(this);
									if(newUrl!=$t.data('view-url'))
										return;

									$t.siblings('.in').removeClass('in').one('bsTransitionEnd', function(){
										$(this).addClass('hidden');
										$t.removeClass('hidden').addClass('in');
									}).emulateTransitionEnd(150);
									isFound = true;									
								})

								if(isFound){
									return false;
								}
							}

							return true;
						}

						// update view
						function update(evt, newUrl, oldUrl){
							defer = $q.defer();
							defer.promise.then(function(){
								if(template){
									var newController = controller,
										arr = viewname.split('.'), 
										openMode = window._gos.openMode;

									linker(newScope, function(clone){
										var $t = $('<div class="view-item fade" ng-controller="'+newController+'">'+template+'</div>');
										$t.data('$ngController', newController);
										$t.data('view-url', newUrl);
										clone.append($t);

										if(currentElement){
											currentElement.children('.in').removeClass('in').one('bsTransitionEnd', function(){
												var $this = $(this);

												switch(openMode){
													case 'self':
														$this.after($t)
														$this.remove();
														break;
													case '':
													case undefined:
														$this.addClass('hidden').nextAll().remove()
														$this.remove();
														currentElement.append($t);
														break;
													case 'blank':
														$this.addClass('hidden').nextAll().remove()
														currentElement.append($t);
														break;
												}
													
												$t.addClass('in');
											}).emulateTransitionEnd(150);

										}else{
											// first load view
											$element.after(clone);
											$t.addClass('in');
											currentElement = clone;
										}

										$compile(clone.contents())(newScope);

									});

								}else{
									//cleanup last view
								}
							});


							if(match['template']){
								template = match.template;
								defer.resolve();
							}else{
								templateUrl = match.templateUrl;
								templateUrl = $sce.getTrustedResourceUrl(templateUrl);
								if (angular.isDefined(templateUrl)) {
									template = $templateRequest(templateUrl);
								}
								$http.get(templateUrl+'?v='+JsVersion).success(function(data) {
									template = data
									defer.resolve();
								});
							}
						}
					}
				};

			return {
				scope : false,
				terminal: true,
				priority: 300,
				transclude: 'element',
				compile : compileFunc
			}
		}]);

})
