(function($) {
	var GosEditor = function(el, buttons, options){
		this.target = el;
		this.isMarkdown = false;
		this.options = options;

		var $target = $(el),
			h = $target.height(),
			bkColor = '';

		if(this.options.autoHeight){
			h = -1;
		}else{
			h = h<0 ? parseInt($target.attr('rows')) * parseFloat($target.css('line-height')) : h;
		}

		if(this.options['bkColor'] && this.options['bkColor']!=''){
			bkColor = 'background-color:'+ this.options['bkColor'];
		}

		var style = '<style>.btn.gos-btn-sm{padding: 2px 10px;}.gos-editor-container{position:relative;max-width:'+options.maxWidth+'}.gos-editor-container.toolbar-bottom{margin-bottom:20px;}.gos-editor-container.menu-bottom .dropdown ul{border-top-style:solid;border-bottom-style:none;top:auto;bottom:12px;}.gos-editor-container.toolbar-bottom .editor-toolbar{z-index:10;margin-bottom: 3px;position: absolute;bottom: -27px;}.gos-editor-container.fullPage{position: fixed;background-color: #FFFFFF;height: 100%;width: 100%;z-index:90;left: 0;top:0;padding: 20px;}.gos-editor-container.fullPage>.editor-toolbar, .gos-editor-container.fullPage>.editor-parent{margin-left:auto;margin-right:auto;padding-left:30px;padding-right:30px;}.gos-editor-container.fullPage>.editor-parent{height: 100%;padding-bottom: 30px;}.gos-editor{overflow: auto;}.single{margin-right:3px}.dropdown {position: relative;}.dropdown:hover ul {display: block;}.dropdown ul{display: none;list-style-type: none;position:absolute;z-index:2;top:20px;right:-1px;left:-1px;border:1px solid #adadad;border-top-style:none;background:#ebebeb;overflow: hidden;padding:0;}.dropdown li:first-child{margin-top:5px;}.dropdown li a{color:inherit}.gos-placeholder{color:#999999;position:absolute;left:10px;top:8px;display:none}.empty>.gos-placeholder{display:block}.editor-parent{position:relative}.btn-group>a[data-command]:last-child{margin-right:3px;}</style>';

		var toolbarPosition = '';
		if(this.options.toolbarPosition=='bottom'){
			toolbarPosition = ' toolbar-bottom'
		}
		this.$editor = $('<div class="gos-editor-container'+toolbarPosition+' '+this.options.class+'">'+style+'<div class="editor-parent clearfix empty"><div tabindex="'+this.options.tabindex+'" class="form-control gos-editor context" contenteditable="true" style="'+(h>0 ? ("height:"+h.toString()+"px:"): 'height:auto;')+ bkColor +'">'+
			'</div><div class="gos-placeholder"></div></div></div>');

		this.$editor.prepend(this.buildToolbar(buttons));

		var thisClas = this;
		this.$editor.find('div.editor-toolbar').on('click', 'a[data-command]', function(e){
			var $this = $(this),
				ctype = $this.data('command-type');

			switch(ctype){
				case 'onClick':
					var option;
					if($this.data('drowdown-item')){
						option = $this.closest('div.dropdown').data('option')
					}else{
						option = $this.data('option');
					}
					if(option.onClick){
						option.onClick.call(this, thisClas.$editor, thisClas)
					}
					break;
				default:
					thisClas.exec($this.data('command'), $this.data('command-value'));
					break;
			}

		})
		
		$target.addClass('hidden').after(this.$editor)
		this.setPlaceholder($target.attr('placeholder')).click(function(e){
			$(this).prev().focus();
		});

		this.$editor.find('div[contenteditable]').focusin(function(e){
			$(this).parent().removeClass('empty');
			$target.trigger('focusin', e);

		}).focusout(function(e) {
			var $this = $(this);
			
			if($this.text()==='')
				$this.parent().addClass('empty');
			else
				$this.parent().removeClass('empty');

			var html = $this.html();
			if(html != $target.html()){
				$target.val(html);
				$target.trigger('change', e);
			}

			$target.trigger('focusout', e);

		}).keypress(function(e){
			$target.trigger('focusout', e);

		});

		return this
	}

	GosEditor.prototype.setPlaceholder = function(txt){
		return this.$editor.find('div.gos-placeholder').text(txt);
	}

	GosEditor.prototype.enable = function(){
		this.$editor.find('div.gos-editor').attr('contenteditable', 'true');
		this.$editor.find('div.editor-toolbar').show()
	}

	GosEditor.prototype.disable = function(){
		this.$editor.find('div.gos-editor').attr('contenteditable', 'false');
		this.$editor.find('div.editor-toolbar').hide()
	}

	GosEditor.prototype.copyToTextArea = function(){
		var html = this.$editor.find('div[contenteditable]').html();
		if(this.isMarkdown){
			if(html.substr(0,1)!="\2")
				html = "\2"+html;
		}else{
			if(html.substr(0,1)=="\2")
				html = html.substr(1)
		}

		this.$editor.prev().html(html);
	}

	GosEditor.prototype.setMarkdownMode = function(v){
		if(v)
			this.$editor.find('div.editor-toolbar a[data-command=mode][data-command-value=0]')
				.trigger('click', this.$editor, this);
		else
			this.$editor.find('div.editor-toolbar a[data-command=mode][data-command-value=1]')
				.trigger('click', this.$editor, this);
		
	}

	GosEditor.prototype.val = function(v){
		if(v){
			var isMarkdown = v.charAt(0)=="\2";
			if(this.isMarkdown != isMarkdown){
				this.setMarkdownMode(isMarkdown);
				this.isMarkdown = isMarkdown;
			}
			$(this.target).val(v);
			this.$editor.find('div[contenteditable]').html(v);
		}else{
			this.copyToTextArea();
			return $(this.target).val();
		}
	}

	GosEditor.prototype.exec = function(action, value){
		if($.fn.gosEditor.safeActions.indexOf(action)===-1)
			return;
		
		if (this.state(action)) {
			document.execCommand(action, false, null);
		} else {
			document.execCommand(action, false, value);
		}
	}

	GosEditor.prototype.state = function(action) {
		return document.queryCommandState(action) === true;
	}

	GosEditor.prototype.buildToolbar = function(data){
		var i,k,$group = '', $toolbar = $('<div class="editor-toolbar" style="margin-bottom:3px;"></div>'), arr;
		for (i=0; i < data.length; i++) {
			if($.isArray(data[i])){
				$group = $('<div class="btn-group"></div>');
				for (k = 0; k<data[i].length; k++) {
					$group.append(this.buildButton(data[i][k]));
				};
				$toolbar.append($group);

			}else{
				$toolbar.append(this.buildButton(data[i], true));
			}
		};
		return $toolbar;
	}

	GosEditor.prototype.buildButton = function(item, lonly){
		if(!item)
			return '';

		var html,$node,btnClass=this.options.btnClass,
			cmd = item['cmd'] ? ' data-command="'+item['cmd']+'"' : '',
			val = item['value'] ? ' data-command-value="'+String(item['value'])+'"' : '',
			single = lonly ? ' single':'',
			tagclass = item['class'] ? ' '+item['class'] : '',
			ctype = 'command', drowdownItem;

		if(this.options.disabledButtons.indexOf(item['cmd'])>-1){
			return null;
		}

		if(item.dialog)
			ctype = 'dialog';
		else if(item.onClick)
			ctype = 'onClick';

		switch(item.type){
			case 'dropdown':
				html = '<div data-group="true" data-command="'+item.cmd+'" class="'+btnClass+tagclass+' dropdown command-active'+single+'"'+(item.value?' data-value="'+item.value+'"': '')+'><span class="content">'+item.content+'</span> <span class="caret"></span><ul>';

				for (var i = item['dropdown'].length - 1; i >= 0; i--) {
					drowdownItem = item['dropdown'][i];
					tagclass = drowdownItem['class'] ? ' class="'+drowdownItem['class']+'" ' : '';
					switch(item.cmd){
						case 'foreColor':
							html += '<li><a href="#"'+tagclass+' style="display:block;width:auto;height:12px;margin:3px;background-color:'+drowdownItem['value']+';" data-command="'+drowdownItem['cmd']+'" data-command-value="'+drowdownItem['value']+'" data-command-type="'+ctype+'" data-drowdown-item="true"></a></li>'
							break;
						default:
							html += '<li><a href="#"'+tagclass+' style="margin-bottom:3px;" data-command="'+drowdownItem['cmd']+'" data-command-value="'+drowdownItem['value']+'" data-command-type="'+ctype+'" data-drowdown-item="true">'+drowdownItem['content']+'</a></li>'
							break;
					}
				};
				html += '</ul></div>'
				$node = $(html);

				break;
			default:
				html = '<a href="#" class="'+btnClass+single+tagclass+'"'+cmd+val+' data-command-type="'+ctype+'">'+item['content']+'</a>';

				$node = $(html);
				break;
		}

		if(ctype!== 'command')
			$node.data('option', item);

		return $node;
	}
	

	$.fn.gosEditor = function(options, buttons) {
		this.each(function() {

			buttons = buttons || $.fn.gosEditor.defaultButtons();
			var opts = {},o;
			for(o in $.fn.gosEditor.defaultOptions){
				opts[o] = $.fn.gosEditor.defaultOptions[o]
			}
			options = $.extend(opts, options);

			$(this).data('gos.editor',  new GosEditor(this, buttons, options));
		})

		return this;
	}

	$.fn.gosEditor.Constructor = GosEditor;

	$.fn.gosEditor.safeActions = ['bold', 'italic', 'underline', 'strikethrough', 'insertunorderedlist', 'insertorderedlist', 'blockquote', 'pre', 'foreColor'];

	$.fn.gosEditor.defaultOptions = {
		btnClass : 'btn btn-default gos-btn-sm',
		toolbarPosition: 'top',
		maxWidth : 'none',
		class : '',
		tabindex : 10,
		disabledButtons : [],
		autoHeight : false
	}

	$.fn.gosEditor.defaultButtons = function(){
		return [
			{cmd: 'mode', value:1, content:'Rich Text', type:'dropdown', 
				dropdown : [
					{cmd: 'mode', value: 1, content: 'Rich Text', class : 'hidden'},
					{cmd: 'mode', value: 0, content: 'MarkDown '}
				],
				onClick : function($editor, clas){
					var $this = $(this),
						$parent = $this.closest('div.dropdown'),
						common = ['mode', 'fullPage', 'insertImage', 'createLink'],
						setVisibleFunc = function(v){
							$editor.find('a[data-command],div[data-command]').each(function(){
								$thisSub = $(this);
								if($thisSub.data('drowdown-item'))
									return;
								if(common.indexOf($thisSub.data('command'))>-1)
									return;
								if(v)
									$thisSub.show();
								else
									$thisSub.hide();
							})
						};
					
					if(parseInt($this.data('command-value'))===1){
						// rich text
						$parent.data('command-value', 1)
						setVisibleFunc(true);
						clas.isMarkdown = false;
					}else{
						// markdown
						$parent.data('command-value', '0')
						setVisibleFunc(false);
						clas.isMarkdown = true;
					}
					
					$parent.find('span.content').text($this.text())
					$parent.find('.hidden').removeClass('hidden');
					$this.addClass('hidden');
				}
			},
			[	
				{cmd: 'bold', content:'<i class="fa fa-bold"></i>'},
				{cmd: 'italic', content:'<i class="fa fa-italic"></i>'}
			],

			[	
				{cmd: 'justifyLeft', content:'<i class="fa fa-align-left"></i>'},
				{cmd: 'justifyCenter', content:'<i class="fa fa-align-center"></i>'},
				{cmd: 'justifyRight', content:'<i class="fa fa-align-right"></i>'}
			],
			
			{cmd: 'foreColor', content:'<i class="fa fa-font"></i>', type:'dropdown', dropdown : [
				{cmd: 'foreColor', value: 'black', content: '<span class="background-color:black">color</span>'}, 
				{cmd: 'foreColor', value: 'silver', content: '<span class="background-color:silver">color</span>'}, 
				{cmd: 'foreColor', value: 'gray', content: '<span class="background-color:gray">color</span>'}, 
				{cmd: 'foreColor', value: 'purple', content: '<span class="background-color:purple">color</span>'}, 
				{cmd: 'foreColor', value: 'fuchsia', content: '<span class="background-color:fuchsia">color</span>'}, 
				{cmd: 'foreColor', value: 'darkOrchid', content: '<span class="background-color:darkOrchid">color</span>'}, 
				{cmd: 'foreColor', value: 'olive', content: '<span class="background-color:olive">color</span>'}, 
				{cmd: 'foreColor', value: 'coral', content: '<span class="background-color:coral">color</span>'}, 
				{cmd: 'foreColor', value: 'maroon', content: '<span class="background-color:maroon">color</span>'}, 
				{cmd: 'foreColor', value: 'crimson', content: '<span class="background-color:crimson">color</span>'}, 
				{cmd: 'foreColor', value: 'navy', content: '<span class="background-color:navy">color</span>'}, 
				{cmd: 'foreColor', value: 'darkGreen', content: '<span class="background-color:darkGreen">color</span>'}, 
				{cmd: 'foreColor', value: 'teal', content: '<span class="background-color:teal">color</span>'}, 
				{cmd: 'foreColor', value: 'dodgerBlue', content: '<span class="background-color:dodgerBlue">color</span>'}, 
				{cmd: 'foreColor', value: 'slateBlue', content: '<span class="background-color:slateBlue">color</span>'}, 
				{cmd: 'foreColor', value: 'steelBlue', content: '<span class="background-color:steelBlue">color</span>'}
			]},
			
			[	
				{cmd: 'insertUnorderedList', content:'<i class="fa fa-list"></i>'},
				{cmd: 'insertOrderedList', content:'<i class="fa fa-list-ol"></i>'}
			],

			[	
				{cmd: 'createLink', content:'<i class="fa fa-link"></i>'},
				{cmd: 'insertImage', content:'<i class="fa fa-picture-o"></i>'}
			],

			{cmd: 'fullPage', value: 1, content:'<i class="fa fa-arrows-alt"></i>', onClick : function($editor, clas){
				var $this = $(this);
				if(parseInt($this.data('command-value'))===1){
					$editor.addClass('fullPage');
					$editor.removeClass('toolbar-bottom');

					$this.data('command-value', '0');
					$editor.find('div.gos-editor').css('height', '100%');
				}else{
					$editor.removeClass('fullPage')
					if(clas.options.toolbarPosition && clas.options.toolbarPosition=='bottom'){
						$editor.addClass('toolbar-bottom');
					}
					$this.data('command-value', '1');
					
					if(clas.options.autoHeight){
						$editor.find('div.gos-editor').css('height', 'auto');
					}else{
						var $target = $editor.prev(),
							h = $target.height();
						h = h<0 ? parseInt($target.attr('rows')) * parseFloat($target.css('line-height')) : h;
						$editor.find('div.gos-editor').css('height', h.toString()+'px');
					}
				}
			}}
		];
	}

	return null;
})(jQuery);

