require.config({
	baseUrl : "/assets/js/",
	paths: {
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
