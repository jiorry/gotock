require.config({
	baseUrl : "/assets/js/",
	paths: {
		'jquery.flot' : 'dev/lib/flot/jquery.flot',
		'jquery.flot.categories' : 'dev/lib/flot/jquery.flot.categories',

		'util' : 'dev/mylib/util',
		'ajax' : 'dev/mylib/ajax',
		'jquery' : 'jquery'

	},
	shim: {
		'jquery.flot' : {
			deps: ['jquery']
		},
		'jquery.flot.categories' : {
			deps: ['jquery.flot']
		}
	}
});

require(
	['ajax', 'util', 'jquery.flot.categories'],
	function (ajax, util){
		function barChart(target, result, s){
			var dataset = [], data;
			for(var i=0;i<36;i++){
				data = [0, 0];
				data[0] = result[i]['date'].substr(5);
				data[1] = (result[i][s]-result[i+1][s])/10000;
				dataset.push(data)
			}
			$.plot(target, [ dataset ], {
				series: {
					bars: {
						show: true,
						barWidth: 0.6,
						fillColor : '#FF9200',
						align: "center"
					}
				},
				xaxis: {
					mode: "categories",
					tickLength: 0,
					transform : function(v){
						return -v
					}
				},
				yaxis: {

				}
			});
		}

		ajax.NewClient("/api/open").send('stock.rzrq.SumData', null)
			.done(function(result){
				barChart('#chartRzyeCanvas', result, 'sm_rzye')
				barChart('#chartRzmreCanvas', result, 'sm_rzmre')
				barChart('#chartRqylyeCanvas', result, 'sm_rqylye')

			}).fail(function(jqXHR){
				var err = JSON.parse(jqXHR.responseText)
				doError(err.message);
			})

});
