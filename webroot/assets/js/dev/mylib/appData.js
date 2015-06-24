define(
	'appData',
	['ajax', 'util', 'websock'],

	function(ajax, util, websock) {
		var appData = {
			userVO : null,
			etag : '',
			branches : null,
			appInitDataState : ''
		};

		// appData.buildJsonEventResult = function(result){
		// 	var events = []
		// 	for(var i=0; i< result.length;i++){
		// 		events.push(JSON.parse(result[i]))
		// 	}
		// 	return events
		// }

		appData.branchName = function(bid, user){
			var info = JSON.parse(user.json_data);
			if(info['branch_names'] && info['branch_names'][String(bid)])
				return info['branch_names'][String(bid)];

			return '';
		};

		appData.init = function(){
			if(this.appInitDataState == 'running'){
				return;
			}

			if(appData.branches==null){
				this.appInitDataState = 'running';
				websock.sendCall("AppInitData", null);
			}
		}

		$(document).on('socket.AppInitDataResult', function(e, args){
			for(var i=0;i<args.branches.length;i++){
				args.branches[i].users = ajax.datasetDecode(args.branches[i].users);
				args.branches[i].groups = ajax.datasetDecode(args.branches[i].groups);
			}
			
			appData.userVO = args.user
			appData.branches = args.branches;
			appData.etag = args.etag
			appData.eventMessage.loadEvents()

		}).on('socket.ErrorResult', function(e, args){
			console.log('socket error:', args.message);

		});
		
		appData.eventMessage = {
			groupData : null,
			loadEvents : function(){
				console.log('---------------------------loadEvents---------------------------------')
				ajax.NewClient("/api/open")
					.sendAlone('module.home.EventGroups', null).done(function(result){
						appData.eventMessage.groupData = result.data;
						appData.appInitDataState = 'ok';
					});
			}
		};

		$(document).on('socket.EventTagResult', function(e, args){
			console.log('EventTagResult', args);
			appData.etag = args;
			appData.eventMessage.loadEvents()

		})

		appData.user = {
			allusers  : function(){
				if(this.users){
					return this.users;
				}
				var i, users = [];
				for(i=0;i<appData.branches.length;i++){
					users = $.merge(users, appData.branches[i].users );
				}
				return users;
			},

			allgroups  : function(){
				if(this.groups){
					return this.groups;
				}
				var i, groups = [];
				for(i=0;i<appData.branches.length;i++){
					groups = $.merge(groups, appData.branches[i].groups );
				}
				return groups;
			},

			isInGroup : function(groupId, user){
				if($.type(user) == 'number'){
					user = util.objectFind('id', user, this.allusers());
				}

				return user.groups.indexOf(parseInt(groupId))>-1
			},

			simpleData2ContactData : function(dataItems){
				if($.type(dataItems) == 'string'){
					dataItems = util.parseArray(dataItems);
				}

				var result={groups:[], users:[]},
					i, item, u, 
					allusers = this.allusers(),
					allgroups = this.allgroups();

				for(i=0;i<dataItems.length;i++){
					item = dataItems[i];
					switch(parseInt(item[0])){
						case 0:
							u = util.objectFind('id', parseInt(item[1]), allusers)
							if(u)
								result.users.push(u);
							break;
						case 1:
							u = util.objectFind('id', parseInt(item[1]), allgroups)
							if(u)
								result.groups.push(u);
							break;
						case 9:
							break;
					}
				}
				return result;
			}
		};

		// appData.init();

		return appData;
	}		
);
