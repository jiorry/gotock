require.config({
	baseUrl : "/assets/js/",
	paths: {
		'util' : 'dev/mylib/util',
		'ajax' : 'dev/mylib/ajax',

		'jquery.jqplot.min' : 'jquery.jqplot.min',
		'jqplot.barRenderer' : 'plugins/jqplot.barRenderer',
		'jqplot.pointLabels' : 'plugins/jqplot.pointLabels',
		'jqplot.highlighter' : 'plugins/jqplot.highlighter',
		'jqplot.canvasTextRenderer' : 'plugins/jqplot.canvasTextRenderer',
		'jqplot.canvasAxisTickRenderer' : 'plugins/jqplot.canvasAxisTickRenderer',
		'jqplot.canvasAxisLabelRenderer' : 'plugins/jqplot.canvasAxisLabelRenderer',
		'jqplot.categoryAxisRenderer' : 'plugins/jqplot.categoryAxisRenderer'
	},
	shim: {
		'jqplot.barRenderer' : {
			deps: ['jquery.jqplot.min']
		},
		'jqplot.pointLabels' : {
			deps: ['jquery.jqplot.min']
		},
		'jqplot.highlighter' : {
			deps: ['jquery.jqplot.min']
		},
		'jqplot.canvasAxisLabelRenderer' : {
			deps: ['jquery.jqplot.min']
		},
		'jqplot.canvasTextRenderer' : {
			deps: ['jquery.jqplot.min']
		},
		'jqplot.canvasAxisTickRenderer' : {
			deps: ['jquery.jqplot.min']
		},
		'jqplot.categoryAxisRenderer' : {
			deps: ['jquery.jqplot.min']
		}
	}
});

require(
	['ajax', 'util', 'jqplot.barRenderer', 'jqplot.categoryAxisRenderer', 'jqplot.pointLabels', 'jqplot.canvasAxisLabelRenderer', 'jqplot.canvasTextRenderer', 'jqplot.canvasAxisTickRenderer', 'jqplot.highlighter'],
	function (ajax, util){
		function lineChart(target, result, s, opt){
			var settings = {baseNum : 100000000};
			$.extend(settings, opt);
			var data = [];

			for(var i=0;i<36;i++){
				data.push([
					result[i]['date'].substr(5),
					result[i][s]/settings.baseNum
				]);
			}
			$.jqplot(target, [data], {
				axes:{
					xaxis: {
						renderer: $.jqplot.CategoryAxisRenderer,
						rendererOptions: { reverse: true },
						tickOptions:{
				        	angle: 90
						},
						labelRenderer: $.jqplot.CanvasAxisLabelRenderer,
                		tickRenderer: $.jqplot.CanvasAxisTickRenderer
	                }
				},
				series: [
		            {color: settings.color}
		        ],
				highlighter: {
					show: true,
					sizeAdjust: 7.5
				}
			});
		}
		function barChart(target, result, s, opt){
			var settings = {baseNum : 100000000};
			$.extend(settings, opt);
			var data = [];

			for(var i=0;i<36;i++){
				data.push([
					result[i]['date'].substr(5),
					(result[i][s]-result[i+1][s])/settings.baseNum
				]
			);
			}
			$.jqplot(target, [data], {
	            seriesDefaults:{
	                renderer:$.jqplot.BarRenderer,
	                rendererOptions: {
						fillToZero: true,
						shadowDepth:2,
						shadowOffset:1,
						barWidth: 20
					},
	                pointLabels: { show: false }
	            },
				series: [
		            {color: settings.color}
		        ],
	            axes: {
	                yaxis: {
						tickOptions:{
			            	formatString:'%.1f'
			            }
					},
	                xaxis: {
	                    renderer: $.jqplot.CategoryAxisRenderer,
						rendererOptions: { reverse: true },
						tickOptions:{
				        	angle: 90
						},
						labelRenderer: $.jqplot.CanvasAxisLabelRenderer,
                		tickRenderer: $.jqplot.CanvasAxisTickRenderer
	                    // ticks: labels
	                }
	            },
				highlighter: {
					show: true,
					sizeAdjust: 7.5
				}
	        });
		}

		ajax.NewClient("/api/open").send('stock.rzrq.SumData', null)
			.done(function(result){
				barChart('chartRzyeBar', result, 'sm_rzye');
				lineChart('chartRzyeLine', result, 'sm_rzye');

				barChart('chartRzmreBar', result, 'sm_rzmre');
				lineChart('chartRzmreLine', result, 'sm_rzmre');

				barChart('chartRqylyeBar', result, 'sm_rqylye', {color: '#FF7700'});
				lineChart('chartRqylyeLine', result, 'sm_rqylye', {color: '#FF7700'});

			}).fail(function(jqXHR){
				var err = JSON.parse(jqXHR.responseText)
				doError(err.message);
			})

});
