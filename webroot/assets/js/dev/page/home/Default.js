require.config({
	waitSeconds :100,
	baseUrl : "/assets/js/",
	paths: {
		'util' : MYENV+'/mylib/util',
		'ajax' : MYENV+'/mylib/ajax'
	},
	shim: {
        'jquery' : {
        	exports:'$'
        }
	}
});

require(
	['ajax', 'util'],
	function (ajax, util){
		$('#txtMail').text('admin@onqee.com');

		var h,l,v

		var $ph = $('#inputPh'),
			$pl = $('#inputPl'),
			$pv = $('#inputPv'),
			$typ = $('#selectType');

		function setLabel(){
			var power = parseFloat($typ.val());
			$('#labelHeigh').text('^'+(1-power).toString());
			$('#labelLow').text('^'+power.toString());
		}
		setLabel();

		$typ.change(function(){
			$ph.trigger('focusout');
			setLabel();
		})
		$ph.focusout(function(){
			h = parseFloat($ph.val());
			l = parseFloat($pl.val());
			v = parseFloat($pv.val());
			if(isNaN(h) ||h==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$pv.val(0);
			}
			calculate();
		})
		$pl.focusout(function(){
			h = parseFloat($ph.val());
			l = parseFloat($pl.val());
			v = parseFloat($pv.val());
			if(isNaN(l) ||l==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$pv.val(0);
			}
			calculate();
		})
		$pv.focusout(function(){
			h = parseFloat($ph.val());
			l = parseFloat($pl.val());
			v = parseFloat($pv.val());
			if(isNaN(v) || v==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$pl.val(0);
			}
			calculate();
		})

		function calculate(){
			var power = parseFloat($typ.val());
			h = parseFloat($ph.val());
			l = parseFloat($pl.val());
			v = parseFloat($pv.val());

			if(h>0 && l>0){
				v = Math.pow(h, 1-power) * Math.pow(l, power);
				$pv.val(v.toFixed(2));
			}else if(l>0 && v>0){ // l>0 && v>0
			   h = Math.pow(v/Math.pow(l, power), 1/(1-power))
			   $ph.val(h.toFixed(2));
		   	}else if(h>0 && v>0){
			   l = Math.pow(v/Math.pow(h, 1-power), 1/power)
			   $pl.val(l.toFixed(2));
		   	}
		}
		// function calculate(){
		// 	h = parseFloat($ph.val());
		// 	l = parseFloat($pl.val());
		// 	v = parseFloat($pv.val());
		//
		// 	if(h>0 && l>0){
		// 		v = Math.pow(h, 0.382) * Math.pow(l, 0.618);
		// 		$pv.val(v.toFixed(2));
		// 	}else if(l>0 && v>0){ // l>0 && v>0
		// 	   	h = Math.pow(v/Math.pow(l, 0.618), 1/0.382)
		// 	   	$ph.val(h.toFixed(2));
		//    	}else if(h>0 && v>0){
		// 	   	l = Math.pow(v/Math.pow(h, 0.382), 1/0.618)
		// 	   	$pl.val(l.toFixed(2));
		//    	}
		// }

});
