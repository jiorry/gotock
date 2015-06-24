define(
	'ngEditor',
	['app','gosEditor'],
	function(app){
    	app.directive('ngEditor', [function(){
    		return {
		        restrict: 'A',   
		        replace: false,
		        link: function(scope, elem, attr){
		        	var buttons = $.fn.gosEditor.defaultButtons();
		        	switch(attr.ngEditor){
	        		case 1:
						delete buttons[0];
	        			break;
		        	}

		            elem.gosEditor({
		            	autoHeight:attr.autoHeight? attr.autoHeight : true, 
		            	toolbarPosition:attr.barPosition? attr.barPosition : 'bottom'
		            }, buttons)

		            elem.on('change', function(){
		            	elem.data('gos.editor').val();
		            })

		            if(attr.ngModel){
		        		scope.$watch(attr.ngModel, function(newValue, oldValue){
		        			elem.data('gos.editor').val(newValue);
		        		});
		        	} // attr.ngModel

		        }
		    }
    	}]);  
	}
)