require.config({
	baseUrl : "/assets/js/",
	paths: {
		'chart' : 'dev/mylib/chart',
		'util' : 'dev/mylib/util',
		'ajax' : 'dev/mylib/ajax'
	}
});

require(
	['ajax', 'util', 'chart'],
	function (ajax, util, chart){
		var dataResult;
		var mapdata = [
			{targetId: 'chartRzyeBar', method: 'barChart', type: 'rzye', opt:{title:'融资余额差值（万元）',baseNum : 10000}},
			{targetId: 'chartRzyeLine', method: 'lineChart', type: 'rzye', opt:{title:'融资余额（万元）',baseNum : 10000}},

			{targetId: 'chartRzmreBar', method: 'barChart', type: 'rzmre', opt:{title:'融资买入额差值（万元）',baseNum : 10000}},
			{targetId: 'chartRzmreLine', method: 'lineChart', type: 'rzmre', opt:{title:'融资买入额（万元）',baseNum : 10000}},

			{targetId: 'chartRqyeBar', method: 'barChart', type: 'rqye', opt:{title:'融券余额差值（万元）', color: '#FF7700',baseNum : 10000}},
			{targetId: 'chartRqyeLine', method: 'lineChart', type: 'rqye', opt:{title:'融券余额（万元）', color: '#FF7700',baseNum : 10000}}
		];

		function prepareChart(){
			var i=0,item;
			for(i=0;i<mapdata.length;i++){
				item = mapdata[i];
				$('#'+item.targetId).empty();
				chart[item.method].call(this, item.targetId, dataResult, item.type, item.opt);
			}
		}

		var code = document.title.split('-')[1]
		ajax.NewClient("/api/open").send('stock.rzrq.StockData', {code: code})
			.done(function(result){
				dataResult = result
				prepareChart();

			}).fail(function(jqXHR){
				var err = JSON.parse(jqXHR.responseText)
				doError(err.message);
			})
});
