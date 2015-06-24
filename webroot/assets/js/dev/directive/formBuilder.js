define(
	'ngFormBuilder',
	['app'],
	function(app){
    	app.directive('ngFormBuilder', ['$rootScope', '$compile', '$animate', function($rootScope, $compile, $animate){
    		function compileFunc(element, attr, linker){
    			
				return function(scope, $element, attr, ctrl, $transclude) {
					var changeCounter = 0,
						previousElement,
						currentElement;

					var cleanupLastIncludeContent = function() {
						if (previousElement) {
							previousElement.remove();
							previousElement = null;
						}
						
						if (currentElement) {
							$animate.leave(currentElement).then(function() {
								previousElement = null;
							});
							previousElement = currentElement;
							currentElement = null;
						}
					};

					var builder = {
						contact : function(d, n){
							var labStr = trans(d.label),
								fieldName = d.name;
							return '<div class="form-group">'+
									(labStr!=''?'<label>'+labStr+': </label>': '')+
									'<div ng-contact-select="'+fieldName+'" name="'+fieldName+'" ng-model="'+fieldName+'" '+(d.require?'ng-required="true"':'')+'></div>'+
								'</div>';
						},

						checkbox : function(d, n){
						
						},

						date : function(d, n, typ){
							var labStr = trans(d.label),
								fieldName = d.name,
								typ = typ ? typ : 'date';
							return '<div class="form-group datepicker" ng-class="{\'has-error\':'+n+'.'+fieldName+'.$invalid && !'+n+'.'+fieldName+'.$pristine}">'+
								'<label>'+labStr+'</label>'+
								'<input name="'+fieldName+'" ng-model="'+fieldName+'" type="text" ng-datetime-picker="'+typ+'" class="form-control" '+(d.range? 'range-picker="'+d.range+'"' : '')+'/>'+
							'</div>'
						},

						datetime : function(d, n){
							return builder.date(d, n, 'hour');
						},

						editor : function(d, n){
							var labStr = trans(d.label),
								fieldName = d.name;
							return '<div class="form-group">'+
										(labStr!=''?'<label>'+labStr+': </label>': '')+
										'<textarea name="'+fieldName+'" ng-model="'+fieldName+'" class="form-control editor" ng-editor="'+d.mode+'" '+(d.require?'ng-required="true"':'')+' placeholder="'+(d.placeholder?d.placeholder:'')+'" rows="8"></textarea>'+
									'</div>';
						},
						textarea : function(d, n){
							var labStr = trans(d.label),
								fieldName = d.name;
							return '<div class="form-group">'+
										(labStr!=''?'<label>'+labStr+': </label>': '')+
										'<textarea name="'+fieldName+'" ng-model="'+fieldName+'" class="form-control editor" '+(d.require?'ng-required="true"':'')+' placeholder="'+(d.placeholder?d.placeholder:'')+'" rows="8"></textarea>'+
									'</div>';
						},
						text : function(d, n){
							var labStr = trans(d.label),
								fieldName = d.name;
							return '<div class="form-group">'+
										(labStr!=''?'<label>'+labStr+': </label>': '')+
										'<strong>{{'+fieldName+'}}</strong>'
									'</div>';
						},
						input : function(d, n){
							var labStr = trans(d.label),
								fieldName = d.name;
							return '<div class="form-group">'+
										(labStr!=''?'<label>'+labStr+': </label>': '')+
										'<input name="'+fieldName+'" ng-model="'+fieldName+'" '+(d.require?'ng-required="true"':'')+' type="text" class="form-control" placeholder="'+(d.placeholder?d.placeholder:'')+'">'+
									'</div>';

						},
						buttonSave : function(d, n){
							var labStrSubmit = trans(d.submit);
							var labStrCancel = trans(d.cancel);
							return '<div class="btn-group mgt-50" style="width:100%">'+
								(labStrCancel!=''?'<a class="cancel btn btn-default col-xs-2">&lt;-'+labStrCancel+'</a>':'')+
								'<a ng-click="submit($event)" ng-disabled="'+n+'.$invalid" class="btn btn-primary btn-submit col-xs-10">'+labStrSubmit+'</a>'+
							'</div>';
						}
					};

					var template;

					scope.$on('$runFormBuilder', update);

					var newScope = scope.$new();
					// update view
					function update(e, opt, name){
						if(!opt)
							return;

						template = '<form name="'+name+'" onsubmit="return false;" novalidate>';
						template+= '<header class="step-title"><em>:</em>' +trans(opt.title)+'</header>'
							
						var i,func,o;
						for(var i=0;i<opt.items.length;i++){
							o = opt.items[i];
							func = builder[o['type']];
							if(func)
								template += func(o, name)
						}
						template += '</form>';
							
						if(currentElement){
							currentElement.remove();
						}

						var clone = $transclude(newScope, function(clone){
							$element.after(clone);
						});
						clone.html(template);
						$compile(clone.contents())(newScope);
						
						currentElement = clone;

						clone.on('click', 'a.btn-submit', function(){
							cleanupLastIncludeContent();
							scope.$emit('$formSubmit', clone, newScope);

						}).on('click', 'a.cancel', function(){
							scope.$emit('$formCancel', clone, newScope);

						})

						scope.$emit('$buildFormSuccess', clone, newScope);
					  
					};

					function trans(o){
						switch(typeof o){
						case 'string':
							return o;

						case 'object':
							var labStr = o[$rootScope.language];
							return labStr? labStr : '';

						default:
							return '';
						}
					};
				}
			};

    		return {
    			scope : false,
    			restrict: 'A',
			    priority: 400,
			    terminal: true,
			    transclude: 'element',
			    controller: angular.noop,
		        compile : compileFunc
		    }
    	}]);  
	}
)
