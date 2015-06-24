define(
	'ngMdEditor',
	['app', 'util'],
	function(app, util){
    	app.directive('ngMdEditor', [function(){
    		return {
		        restrict: 'A',   
		        replace: true,
		        template : '<div><textarea rows=5 class="form-control gos-md-editor"></textarea>' +
		        			'<div class="editor-toolbar" style="margin-top:3px;">'+
								'<a class="btn btn-default single" data-command="href" data-command-value="/" data-command-type="onClick"><i class="fa fa-life-ring"></i> MD HELP</a>'+
								'<div class="btn-group">'+
									'<a class="btn btn-default" data-command="createLink" data-command-type="command"><i class="fa fa-link"></i></a>'+
									'<a class="btn btn-default" data-command="insertImage" data-command-type="command"><i class="fa fa-picture-o"></i></a>'+
								'</div>'+
								'<a class="btn btn-default single" data-command="fullPage" data-command-value="1" data-command-type="onClick"><i class="fa fa-arrows-alt"></i></a>'+
							'</div></div>',
		        link: function(scope, elem, attr){
		        	var $editor = elem.children('textarea.form-control');

		        	var defaultHeight = $editor.height(), lastHeight = defaultHeight, newHeight;

					var commands = [
						{cmd: 'fullPage', value: 1, onClick : function($target, $elem, $editor){
							if(parseInt($target.data('command-value'))===1){
								$elem.addClass('fullPage');
								$target.data('command-value', '0');
								$editor.css('height', '100%');
							}else{
								$elem.removeClass('fullPage')
								$target.data('command-value', '1');
								$editor.css('height', 'auto');
							}
						}},
					];

					$editor.next().on('click', 'a[data-command]', function(){
						var $t = $(this);
						var found = util.objectFind('cmd', $t.data('command'), commands);
						if(!found) return;

						found.onClick.call(this, $t, elem, $editor)
					});

					if(attr.ngModel && attr.ngModel!=''){
						$editor.on('keyup', function(){
							scope[attr.ngModel] = $editor.val();
						});
						$editor.on('focusout', function(){
							scope.$apply(function(){
								scope[attr.ngModel] = $editor.val();
							});

							console.log($editor.html());

							var oo = $('<div>'+$editor.html()+'</div>');

						});
						scope.$watch(attr.ngModel, function(newValue, oldValue){
							if(newValue != oldValue)
		        				$editor.val(newValue);
		        		});
					}

		        }
		    }
    	}]);  
	}
)