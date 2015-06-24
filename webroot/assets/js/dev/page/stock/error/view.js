define('error.view', ['app'], function(app, util){
	app.controller('ErrorViewCtrl', ['$scope',  function($scope){
	  	app.setTitle('error');

	}]);

	return {i18n : false};
})

