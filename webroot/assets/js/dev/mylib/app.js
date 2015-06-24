window._gos = window._gos || {};

define(
	'app',
	['util', 'loader', 'angular', 'angular-route', 'angular-animate', 'angular-resource', 'jquery'],
	function(util, loader){
    	var app = angular.module('app',['ngResource', 'ngRoute', 'ngAnimate']);

    	app.setTitle = function(title){
    		util.setTitle(title);
    	}

		window._gos.viewstackRouter = window._gos.viewstackRouter || [];

		app.gosRouter = function(r, name, ctr, group){
			window._gos.viewstackRouter.push(
				{
					router : r,
					name : name,
			        templateUrl : window._gos.srcPath + name.split('.').join('/')+'.html',
					group : group,
			        controller : ctr
			    }
			);
		}

		window._gos.srcPath = '/assets/js/'+MYENV+'/page/stock/';		

		app.gosRouter('/', 'home.module', 'HomeModuleCtrl');
		app.gosRouter('/home/module', 'home.module', 'HomeModuleCtrl');

		// app.gosRouter('/git/:nick/:id', 'git.view', 'GitViewCtrl');

		app.gosRouter('/error/view', 'error.view', 'ErrorViewCtrl');

    	return app;
	}
)