require.config({
	baseUrl : "/assets/js",
	paths: {
		"bootstrap-datepicker" : "dev/lib/bootstrap-datepicker"
	}
});

define(
	'datetimePick',
	['util', 'bootstrap-datepicker'],

	function(util){
		$.fn.datepicker.dates['zh-CN'] = {
			days: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"],
			daysShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六", "周日"],
			daysMin:  ["日", "一", "二", "三", "四", "五", "六", "日"],
			months: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
			monthsShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
			today: "今日",
			format: "yyyy年mm月dd日",
			weekStart: 1
		};

		var GosDatePicker = function(el){
			this.element = el;

			this.init = function(date, startDate, endDate, opts){
				var defaultOpts = {
					mode:0, 
					language:'en', 
					format: 'yyyy-mm-dd',
					onChange : null
				}
				this.options = $.extend(defaultOpts, opts)

				var $this = $(this.element),
					clas = this;
				this.picker = $this.datepicker({
					language : this.options.language,
					autoclose : true,
					startDate : startDate,
					endDate : endDate,
				    format: this.options.format
				});

				$this.val(util.date2str(date));

				var i,v,html = '<em class="fa fa-calendar"></em>';

				if(this.options.mode>0){
					html += '<div style="display:inline;"><select class="form-control cmbhour">';
				}
				
				if(this.options.mode >= 1){
					for(i=0;i<24;i++){
						v = i.toString();
						html += '<option'+(i==date.getHours() ? ' selected="selected"' : '')+' value="'+v+'">'+v+'</option>';
					}
					html += '</select>时';
				}

				if(this.options.mode==2){
					html += '<select class="form-control cmbmin">';
					for(i=0;i<12;i++){
						v = (i*5).toString(); 
						html += '<option'+(i*5<=date.getMinutes() && date.getMinutes()<(i+1)*5 ? ' selected="selected"' : '')+' value="'+v+'">'+v+'</option>';
					}
					html += '</select>分';
				}
					
				html += '</div>';
				var part = $(html);

				if(this.options['onChange']){
					this.picker.on('changeDate', function(){
						clas.options.onChange.call($this)
					})
					part.find('select').change(function(){
						clas.options.onChange.call($this)
					})
				}

				$this.after(part);

				return this;
			}
			
			this.getDate = function(){
				var $this = $(this.element),
					d = this.picker.datepicker('getDate'),
					$box = $this.next().next();

				switch(this.options.mode){
					case 1:
						d.setHours(parseInt($box.find('select.cmbhour').val()));
						break;
					case 2:
						d.setHours(parseInt($box.find('select.cmbhour').val()));
						d.setMinutes(parseInt($box.find('select.cmbmin').val()));
						break;
				}
				return d;
			}
		}

		$.fn.gosDatePicker = function(date, startDate, endDate, opts) {
			if(!date){
				date = new Date();
			}
			return this.each(function() {
				var $this = $(this),
					data = $this.data('gos.datepicker');

				if(data) return;
				$this.data('gos.datepicker', (new GosDatePicker(this)).init(date, startDate, endDate, opts));
			})
		}

	}
)
