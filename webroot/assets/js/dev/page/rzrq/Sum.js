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
			{targetId: 'chartRzyeBar', method: 'barChart', type: 'rzye', opt:{title:'融资余额差值（亿元）'}},
			{targetId: 'chartRzyeLine', method: 'lineChart', type: 'rzye', opt:{title:'融资余额（亿元）'}},

			{targetId: 'chartRzmreBar', method: 'barChart', type: 'rzmre', opt:{title:'融资买入额差值（亿元）'}},
			{targetId: 'chartRzmreLine', method: 'lineChart', type: 'rzmre', opt:{title:'融资买入额（亿元）'}},

			{targetId: 'chartRqylyeBar', method: 'barChart', type: 'rqylye', opt:{title:'融券余额差值（亿元）', color: '#FF7700'}},
			{targetId: 'chartRqylyeLine', method: 'lineChart', type: 'rqylye', opt:{title:'融券余额（亿元）', color: '#FF7700'}}
		];

		function prepareChart(prefix){
			var i=0,k=0,item;
			for(i=0;i<mapdata.length;i++){
				item = mapdata[i];
				$('#'+item.targetId).empty();
				// remove -1
				for(k=0;k<dataResult.length;k++){
					if(dataResult[k][prefix+'_'+item.type]<0){
						if(k==dataResult.length-1){
							dataResult[k][prefix+'_'+item.type] = dataResult[k-1][prefix+'_'+item.type];
						}else{
							dataResult[k][prefix+'_'+item.type] = dataResult[k+1][prefix+'_'+item.type];
						}
					}
				}
				chart[item.method].call(this, item.targetId, dataResult, prefix+'_'+item.type, item.opt);
			}
		}

		$(':radio[name=optionsRadios]').change(function(){
			prepareChart($(this).filter(':checked').val());
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

		ajax.NewClient("/api/open").send('stock.rzrq.SumData', null)
			.done(function(result){
				dataResult = result
				prepareChart($(':radio[name=optionsRadios]:checked').val());

			}).fail(function(jqXHR){
				var err = JSON.parse(jqXHR.responseText)
				doError(err.message);
			})

});
