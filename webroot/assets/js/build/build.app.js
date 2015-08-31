({
	// optimize: "none",
	paths: {
		'util' : 'mylib/util',
		'ajax' : 'mylib/ajax',
		'chart' : 'mylib/chart',

		'jquery.jqplot' : 'jqplot/jquery.jqplot',
		'jqplot.barRenderer' : 'jqplot/plugins/jqplot.barRenderer',
		'jqplot.pointLabels' : 'jqplot/plugins/jqplot.pointLabels',
		'jqplot.highlighter' : 'jqplot/plugins/jqplot.highlighter',
		'jqplot.canvasTextRenderer' : 'jqplot/plugins/jqplot.canvasTextRenderer',
		'jqplot.canvasAxisTickRenderer' : 'jqplot/plugins/jqplot.canvasAxisTickRenderer',
		'jqplot.canvasAxisLabelRenderer' : 'jqplot/plugins/jqplot.canvasAxisLabelRenderer',
		'jqplot.categoryAxisRenderer' : 'jqplot/plugins/jqplot.categoryAxisRenderer',

		'crypto' : 'empty:',
		'jquery' : 'empty:'
	},
	shim: {
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
	},

	baseUrl : "../dev",
	removeCombined : true,
	modules: [
		{
			name : 'chart',
			create: false,
			include : ['jquery.jqplot', 'jqplot.barRenderer', 'jqplot.categoryAxisRenderer', 'jqplot.pointLabels', 'jqplot.canvasAxisLabelRenderer', 'jqplot.canvasTextRenderer', 'jqplot.canvasAxisTickRenderer', 'jqplot.highlighter']
		},
		{
			name : 'page/rzrq/Stock',
			create: false,
			exclude : ['chart']
		},
		{
			name : 'page/rzrq/Sum',
			create: false,
			exclude : ['chart']
		},
		{
			name : 'page/user/Login',
			create: false,
			exclude : ['crypto']
		},
		{
			name : 'page/user/Regist',
			create: false,
			exclude : ['crypto']
		},
		{
			name : 'page/home/Default',
			create: false,
			exclude : []
		}

	],

	dir : '../pro'
})
