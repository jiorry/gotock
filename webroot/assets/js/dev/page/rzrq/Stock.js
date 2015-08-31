require.config({
	waitSeconds :100,
	baseUrl : "/assets/js/",
	paths: {
		'chart' : MYENV+'/mylib/chart',
		'util' : MYENV+'/mylib/util',
		'ajax' : MYENV+'/mylib/ajax'
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

			if(dataResult.length>0){
				$('#labelStockTitle').text(dataResult[0].code + ' - ' + dataResult[0].name);
			}
		}

		var code = document.title.split('-')[1]
		ajax.NewClient("/api/open").send('stock.rzrq.StockData', {code: code})
			.done(function(result){
				dataResult = result
				prepareChart();

			}).fail(function(jqXHR){
				var err = JSON.parse(jqXHR.responseText)
				$('#labelStockTitle').text(err.message);
			})

		$('#btnRzrqQuery').click(function(){
			var code = $('#txtCode').val();
			var p = /[a-zA-Z]/;
			if(code=='' || p.test(code)){
				alert('股票代码错误，请重新输入！');
				return;
			}
			window.location.href = '/rzrq/stock/' + code;
		});

		$('#txtCode').keypress(function(e){
			if (e.which == 13) {
			   $('#btnRzrqQuery').trigger('click');
			}
		})
});
