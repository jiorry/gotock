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

		var h,l,v,c;

		var $pA = $('#inputA'),
			$pB = $('#inputB'),
			$pV = $('#inputV'),
			$typ = $('#selectType'),
			$inputPA = $('#inputPA'),
			$inputPB = $('#inputPB');

		function setLabel(){
			var power = parseFloat($typ.val());
			$inputPA.val(1-power);
			$inputPB.val(power);

			$inputHPA.val(parseFloat($typ.val()));
		}


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
			calculateP();
		}).dblclick(function(){
			$pA.val('');
		});

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
			calculateP();
		}).dblclick(function(){
			$pB.val('');
		});

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
			calculateP();
		}).dblclick(function(){
			$pV.val('');
		});

		function calculateP(){
			var powerA = parseFloat($inputPA.val()),
				powerB = parseFloat($inputPB.val());
			h = parseFloat($pA.val());
			l = parseFloat($pB.val());
			v = parseFloat($pV.val());

			if(h>0 && l>0){
				v = Math.pow(h, powerB) * Math.pow(l, powerA);
				$pV.val(v.toFixed(2));
			}else if(l>0 && v>0){ // l>0 && v>0
			   h = Math.pow(v/Math.pow(l, powerA), 1/(powerB))
			   $pA.val(h.toFixed(2));
		   	}else if(h>0 && v>0){
			   l = Math.pow(v/Math.pow(h, powerB), 1/powerA)
			   $pB.val(l.toFixed(2));
		   	}
		}
// ---------------------------------
		var $phA = $('#inputHA'),
			$phB = $('#inputHB'),
			$phC = $('#inputHC'),
			$phV = $('#inputHV'),
			$inputHPA = $('#inputHPA');

		$phA.focusout(calculateH);
		$phB.focusout(calculateH);
		$phC.focusout(calculateH);

		function calculateH(){
			var power = parseFloat($inputHPA.val());
			h = parseFloat($phA.val());
			l = parseFloat($phB.val());
			c = parseFloat($phC.val());
			if(h>0 && l>0 && c>0 && power>0){
				$phV.val((Math.pow(h/l, power)*c).toFixed(2));
			}
		}

		setLabel();
});
