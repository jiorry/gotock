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
		var h,l,v
		// ------- p2
		var $p2h = $('#inputP2h'),
			$p2l = $('#inputP2l'),
			$p2v = $('#inputP2v');
		$p2h.focusout(function(){
			h = parseFloat($p2h.val());
			l = parseFloat($p2l.val());
			v = parseFloat($p2v.val());
			if(isNaN(h) ||h==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$p2v.val(0);
			}
			calculateP2();
		})
		$p2l.focusout(function(){
			h = parseFloat($p2h.val());
			l = parseFloat($p2l.val());
			v = parseFloat($p2v.val());
			if(isNaN(l) ||l==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$p2v.val(0);
			}
			calculateP2();
		})
		$p2v.focusout(function(){
			h = parseFloat($p2h.val());
			l = parseFloat($p2l.val());
			v = parseFloat($p2v.val());
			if(isNaN(v) || v==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$p2l.val(0);
			}
			calculateP2();
		})

		// ------- p3
		var $p3h = $('#inputP3h'),
			$p3l = $('#inputP3l'),
			$p3v = $('#inputP3v');

		$p3h.focusout(function(){
			h = parseFloat($p3h.val());
			l = parseFloat($p3l.val());
			v = parseFloat($p3v.val());
			if(isNaN(h) ||h==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$p3v.val(0);
			}
			calculateP3();
		})
		$p3l.focusout(function(){
			h = parseFloat($p3h.val());
			l = parseFloat($p3l.val());
			v = parseFloat($p3v.val());
			if(isNaN(l) ||l==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$p3v.val(0);
			}
			calculateP3();
		})
		$p3v.focusout(function(){
			h = parseFloat($p3h.val());
			l = parseFloat($p3l.val());
			v = parseFloat($p3v.val());
			if(isNaN(v) || v==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$p3l.val(0);
			}
			calculateP3();
		})

		function calculateP2(){
			h = parseFloat($p2h.val());
			l = parseFloat($p2l.val());
			v = parseFloat($p2v.val());

			if(h>0 && l>0){
				v = Math.sqrt(h*l);
				$p2v.val(v.toFixed(2));
			}else if(l>0 && v>0){ // l>0 && v>0
			   h = Math.pow(v,2)/l;
			   $p2h.val(h.toFixed(2));
		   	}else if(h>0 && v>0){
			   l = Math.pow(v,2)/h;
			   $p2l.val(l.toFixed(2));
		   	}
		}
		function calculateP3(){
			h = parseFloat($p3h.val());
			l = parseFloat($p3l.val());
			v = parseFloat($p3v.val());

			if(h>0 && l>0){
				v = Math.pow(h, 0.382) * Math.pow(l, 0.618);
				$p3v.val(v.toFixed(2));
			}else if(l>0 && v>0){ // l>0 && v>0
			   h = Math.pow(v/Math.pow(l, 0.618), 1/0.382)
			   $p3h.val(h.toFixed(2));
		   	}else if(h>0 && v>0){
			   l = Math.pow(v/Math.pow(h, 0.382), 1/0.618)
			   $p3l.val(l.toFixed(2));
		   	}
		}

});
