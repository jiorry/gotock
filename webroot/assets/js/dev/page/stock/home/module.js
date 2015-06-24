define('home.module', ['app'], function(app){
	app.controller('HomeModuleCtrl', ['$scope', '$element', function($scope, $element){
	  	$scope.login = 'jiorry#lDdj2niqPnS';

	}]);

	return {
		title : '项目列表',
		goBackButton : false,
		headerHtml : '<a href="/project/form/new" open-mode="blank">+添加项目</a>'
	};
})

