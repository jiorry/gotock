require.config({
	baseUrl : "/assets/js/",
	paths: {
		'jquery.jqplot' : MYENV+'/jqplot/jquery.jqplot',
		'jqplot.barRenderer' : MYENV+'/jqplot/plugins/jqplot.barRenderer',
		'jqplot.pointLabels' : MYENV+'/jqplot/plugins/jqplot.pointLabels',
		'jqplot.highlighter' : MYENV+'/jqplot/plugins/jqplot.highlighter',
		'jqplot.canvasTextRenderer' : MYENV+'/jqplot/plugins/jqplot.canvasTextRenderer',
		'jqplot.canvasAxisTickRenderer' : MYENV+'/jqplot/plugins/jqplot.canvasAxisTickRenderer',
		'jqplot.canvasAxisLabelRenderer' : MYENV+'/jqplot/plugins/jqplot.canvasAxisLabelRenderer',
		'jqplot.categoryAxisRenderer' : MYENV+'/jqplot/plugins/jqplot.categoryAxisRenderer'
	},
	shim: {
		'jqplot.jqplot' : {
			deps: ['jquery']
		},
		'jqplot.barRenderer' : {
			deps: ['jquery.jqplot']
		},
		'jqplot.pointLabels' : {
			deps: ['jquery.jqplot']
		},
		'jqplot.highlighter' : {
			deps: ['jquery.jqplot']
		},
		'jqplot.canvasAxisLabelRenderer' : {
			deps: ['jquery.jqplot']
		},
		'jqplot.canvasTextRenderer' : {
			deps: ['jquery.jqplot']
		},
		'jqplot.canvasAxisTickRenderer' : {
			deps: ['jquery.jqplot']
		},
		'jqplot.categoryAxisRenderer' : {
			deps: ['jquery.jqplot']
		}
	}
});
define(
	'chart',
	['jqplot.barRenderer', 'jqplot.categoryAxisRenderer', 'jqplot.pointLabels', 'jqplot.canvasAxisLabelRenderer', 'jqplot.canvasTextRenderer', 'jqplot.canvasAxisTickRenderer', 'jqplot.highlighter'],

	function() {
		var chart = {};

        chart.lineChart = function(target, result, s, opt){
			var settings = {title:'',baseNum : 100000000};
			$.extend(settings, opt);
			var data = [];

			for(var i=0;i<36;i++){
				data.push([
					result[i]['date'].substr(5,5),
					result[i][s]/settings.baseNum
				]);
			}
			$.jqplot(target, [data], {
				title: settings.title,
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
		chart.barChart = function (target, result, s, opt){
			var settings = {title:'',baseNum : 100000000};
			$.extend(settings, opt);
			var data = [];

			for(var i=0;i<36;i++){
				data.push([
					result[i]['date'].substr(5,5),
					(result[i][s]-result[i+1][s])/settings.baseNum
				]);
			}

			$.jqplot(target, [data], {
				title: settings.title,
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

		return chart;
	}
);
