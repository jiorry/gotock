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

		var $pA = $('#inputA'),
			$pB = $('#inputB'),
			$pV = $('#inputV'),
			$typ = $('#selectType'),
			$inputPA = $('#inputPA'),
			$inputPB = $('#inputPB');

		$inputPA.change(function(){
			$inputPB.val(1-parseFloat($inputPA.val()));
		})
		$inputPB.change(function(){
			$inputPA.val(1-parseFloat($inputPB.val()));
		})
		function setLabel(){
			var power = parseFloat($typ.val());
			$inputPA.val(1-power);
			$inputPB.val(power);
		}
		setLabel();

		$typ.change(function(){
			$pA.trigger('focusout');
			setLabel();
		})
		$pA.focusout(function(){
			h = parseFloat($pA.val());
			l = parseFloat($pB.val());
			v = parseFloat($pV.val());
			if(isNaN(h) ||h==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$pV.val(0);
			}
			calculate();
		})
		$pB.focusout(function(){
			h = parseFloat($pA.val());
			l = parseFloat($pB.val());
			v = parseFloat($pV.val());
			if(isNaN(l) ||l==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$pV.val(0);
			}
			calculate();
		})
		$pV.focusout(function(){
			h = parseFloat($pA.val());
			l = parseFloat($pB.val());
			v = parseFloat($pV.val());
			if(isNaN(v) || v==0){
				return
			}

			if(h>0&&l>0&&v>0){
				$pB.val(0);
			}
			calculate();
		})

		function calculate(){
			var power = parseFloat($typ.val());
			h = parseFloat($pA.val());
			l = parseFloat($pB.val());
			v = parseFloat($pV.val());

			if(h>0 && l>0){
				v = Math.pow(h, 1-power) * Math.pow(l, power);
				$pV.val(v.toFixed(2));
			}else if(l>0 && v>0){ // l>0 && v>0
			   h = Math.pow(v/Math.pow(l, power), 1/(1-power))
			   $pA.val(h.toFixed(2));
		   	}else if(h>0 && v>0){
			   l = Math.pow(v/Math.pow(h, 1-power), 1/power)
			   $pB.val(l.toFixed(2));
		   	}
		}

});
