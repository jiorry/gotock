define(
	'ngBootstrapSwitch',
	['app'],
	function(app){
    	app.directive('ngBootstrapSwitch', [function(){
    		return {
		        restrict: 'A',
		        replace: false,
		        link: function(scope, elem, attr){
		        	elem.hide();
		        	var isChecked = elem.is(":checked");

		        	var $t = $('<div class="mgb-10 bootstrap-switch bootstrap-switch-wrapper bootstrap-switch-animate '+(isChecked?'bootstrap-switch-on' : 'bootstrap-switch-off')+'">'+
		        					'<div class="bootstrap-switch-container">'+
		        						'<span class="bootstrap-switch-handle-on bootstrap-switch-primary">ON</span>'+
		        						'<label class="bootstrap-switch-label">&nbsp;</label>'+
		        						'<span class="bootstrap-switch-handle-off bootstrap-switch-default">OFF</span>'+
		        					'</div>'+
		        				'</div>');

		        	$t.click(function(){
		        		var val = false;
		        		if($t.hasClass('bootstrap-switch-on')){
		        			$t.removeClass('bootstrap-switch-on').addClass('bootstrap-switch-off');
		        			val = false;
		        		}else{
		        			$t.removeClass('bootstrap-switch-off').addClass('bootstrap-switch-on');
		        			val = true;
		        		}
	        			elem.attr('checked', val);

		        		if(attr.ngModel){
		        			scope.$apply(function(){
		        				scope[attr.ngModel] = val;
		        			})
	        			}
		        	})

		        	elem.after($t)

		        	if(attr.ngModel){
		        		scope[attr.ngModel] = isChecked;

		        		scope.$watch(attr.ngModel, function(newValue, oldValue){
		        			if(newValue){
		        				$t.removeClass('bootstrap-switch-off').addClass('bootstrap-switch-on');
		        			}else{
		        				$t.removeClass('bootstrap-switch-on').addClass('bootstrap-switch-off');
		        			}
		        		});
		        	}
		        }
		    }
    	}]);  
	}
)