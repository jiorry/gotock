window._gos = window._gos || {};

require.config({
	baseUrl : "/assets/js/",
	paths: {
		'util' : 'dev/mylib/util',
		'ajax' : 'dev/mylib/ajax'
	},
	shim: {
        'jquery' : {
        	exports:'$'
        }
	}
});

require(
	['ajax', 'util'],
	function (ajax, util, app){


});
