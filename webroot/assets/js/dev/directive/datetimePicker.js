require.config({
	baseUrl : "/assets/js",
	paths: {
		"bootstrap-datepicker" : "dev/lib/bootstrap-datepicker"
	}
});
define(
	'ngDatetimePicker',
	['app','util', 'jquery', 'bootstrap-datepicker'],
	function(app, util){
    	app.directive('ngDatetimePicker', ['$rootScope', function($rootScope){
    		return {
		        restrict: 'AE',   
		        replace: false,
		        link: function(scope,elem,attr){
	        		var now = new Date(),
		            	startDate = attr.start ? util.str2date(attr.start) : new Date(now.getTime()-util.DATE_DAY*10),
		            	endDate = attr.start ? util.str2date(attr.end) : null,
		            	html, picker, mode, v;

					var $date = $('<input name="'+attr.ngModel+'" type="text" date-type="datepicker" class="form-control">'),
						formName = elem.closest('form[name]');

		        	formName = formName.length==0 ? null : formName.attr('name');

					picker = $date.datepicker({
						language : $rootScope.language?$rootScope.language: 'en_US',
						autoclose : true,
						startDate : startDate,
						endDate : endDate,
					    format: attr.format? attr.format : 'yyyy-mm-dd'
					});

					elem.addClass('hidden');

		        	html = '<em class="fa fa-calendar"></em>';
		        	mode = getMode(attr.ngDatetimePicker);
			        	
		        	if(mode>0){
						html += '<div style="display:inline;"><select class="form-control cmbhour">';
					}
					
					if(mode >= 1){
						for(i=0;i<24;i++){
							v = i.toString();
							html += '<option value="'+v+'">'+v+'</option>';
						}
						html += '</select> h';
					}

					if(mode==2){
						html += '<select class="form-control cmbmin">';
						for(i=0;i<12;i++){
							v = (i*5).toString(); 
							html += '<option value="'+v+'">'+v+'</option>';
						}
						html += '</select> m';
					}
						
					html += '</div>';

					var elemPart = $(html);
					elemPart.find('select').change(function(){
						$date.trigger('changeDate');
					})

					$date.on('changeDate', function(){
						if(formName)
							scope[formName][attr.ngModel].$setViewValue(getDate());
						else{
							scope.$apply(function(){
								scope[attr.ngModel] = getDate();
							})
						}
					})

					scope[attr.ngModel] = now;

					if(attr.rangePicker){
						setValidity(attr.ngModel, false);
					}

					elem.after(elemPart);
					elem.after($date);
					scope.$watch(attr.ngModel, function(newValue, oldValue){
						
						if(!newValue){
							setForm(now);
							return;
						}else{
							setForm(newValue);
						}

						switch(attr.rangePicker){
						case 'begin':
							var $t = elem.parent().siblings('.datepicker').children('[ng-datetime-picker]');
							if(oldValue && newValue.getTime()!=oldValue.getTime()){
								var isOk = newValue.getTime() < $t.data('date-value');
								setValidity(attr.ngModel, isOk);
								setValidity($t.attr('ng-model'), isOk);
							}
							break;

						case 'end':
							var $t = elem.parent().siblings('.datepicker').children('[ng-datetime-picker]');
							if(oldValue && newValue.getTime()!=oldValue.getTime()){
								var isOk = newValue.getTime() > $t.data('date-value');
								setValidity(attr.ngModel, isOk);
								setValidity($t.attr('ng-model'), isOk);
							}
							break;

						}
					});
						
					function setValidity(name, v){
						if(!name)
							return;

						if(formName)
							scope[formName][name].$setValidity(name, v);
						else
							scope[name].$setValidity(name, v);
					}

					function getMode(v){
		        		switch(v){
			        	case 'hour':
			        		return 1;
			        	case 'minute':
			        		return 2
			        	default:
			        		return 0;
			        	}
		        	}

					function getDate(){
						var d = picker.datepicker('getDate');

						switch(mode){
							case 1:
								d.setHours(parseInt(elemPart.find('select.cmbhour').val()));
								break;
							case 2:
								d.setHours(parseInt(elemPart.find('select.cmbhour').val()));
								d.setMinutes(parseInt(elemPart.find('select.cmbmin').val()));
								break;
						}
						return d;
					};// getDate
		        	
					function setForm(d){
						if(!d)
							return;
						elem.data('date-value', d);
						picker.datepicker(d);
						$date.val(util.date2str(d));
						elemPart.find('select.cmbhour').val(d.getHours());
						// minutes = [0 5 10 15 20 25 ... 55]
						elemPart.find('select.cmbmin').val( parseInt(d.getMinutes()/5)*5 );

					};// setDate
				        
		        }
		    }
    	}]);  
	}
)