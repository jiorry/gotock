/**
 * jqPlot
 * Pure JavaScript plotting plugin using jQuery
 *
 * Version: 1.0.8
 * Revision: 1250
 *
 * Copyright (c) 2009-2013 Chris Leonello
 * jqPlot is currently available for use in all personal or commercial projects 
 * under both the MIT (http://www.opensource.org/licenses/mit-license.php) and GPL 
 * version 2.0 (http://www.gnu.org/licenses/gpl-2.0.html) licenses. This means that you can 
 * choose the license that best suits your project and use it accordingly. 
 *
 * Although not required, the author would appreciate an email letting him 
 * know of any substantial use of jqPlot.  You can reach the author at: 
 * chris at jqplot dot com or see http://www.jqplot.com/info.php .
 *
 * If you are feeling kind and generous, consider supporting the project by
 * making a donation at: http://www.jqplot.com/donate.php .
 *
 * sprintf functions contained in jqplot.sprintf.js by Ash Searle:
 *
 *     version 2007.04.27
 *     author Ash Searle
 *     http://hexmen.com/blog/2007/03/printf-sprintf/
 *     http://hexmen.com/js/sprintf.js
 *     The author (Ash Searle) has placed this code in the public domain:
 *     "This code is unrestricted: you are free to use it however you like."
 * 
 */

(function(e){function t(){e.jqplot.GenericCanvas.call(this),this.isDragging=!1,this.isOver=!1,this._neighbor,this._cursors=[]}function n(t,n){var r=t.series[n.seriesIndex],i=r.plugins.dragable,s=r.markerRenderer,o=i.markerRenderer;o.style=s.style,o.lineWidth=s.lineWidth+2.5,o.size=s.size+5;if(!i.color){var u=e.jqplot.getColorComponents(s.color),a=[u[0],u[1],u[2]],f=u[3]>=.6?u[3]*.6:u[3]*(2-u[3]);i.color="rgba("+a[0]+","+a[1]+","+a[2]+","+f+")"}o.color=i.color,o.init();var l=n.pointIndex>0?n.pointIndex-1:0,c=n.pointIndex+2;i._gridData=r.gridData.slice(l,c)}function r(e,t,n,r,i){if(i.plugins.dragable.dragCanvas.isDragging){var s=i.plugins.dragable.dragCanvas,o=s._neighbor,u=i.series[o.seriesIndex],a=u.plugins.dragable,f=u.gridData,l=a.constrainTo=="y"?o.gridData[0]:t.x,c=a.constrainTo=="x"?o.gridData[1]:t.y,h=u._xaxis.series_p2u(l),p=u._yaxis.series_p2u(c),d=s._ctx;d.clearRect(0,0,d.canvas.width,d.canvas.height),o.pointIndex>0?a._gridData[1]=[l,c]:a._gridData[0]=[l,c],i.series[o.seriesIndex].draw(s._ctx,{gridData:a._gridData,shadow:!1,preventJqPlotSeriesDrawTrigger:!0,color:a.color,markerOptions:{color:a.color,shadow:!1},trendline:{show:!1}}),i.target.trigger("jqplotSeriesPointChange",[o.seriesIndex,o.pointIndex,[h,p],[l,c]])}else if(r!=null){var v=i.series[r.seriesIndex];if(v.isDragable){var s=i.plugins.dragable.dragCanvas;s.isOver||(s._cursors.push(e.target.style.cursor),e.target.style.cursor="pointer"),s.isOver=!0}}else if(r==null){var s=i.plugins.dragable.dragCanvas;s.isOver&&(e.target.style.cursor=s._cursors.pop(),s.isOver=!1)}}function i(e,t,r,i,s){var o=s.plugins.dragable.dragCanvas;o._cursors.push(e.target.style.cursor);if(i!=null){var u=s.series[i.seriesIndex],a=u.plugins.dragable;u.isDragable&&!o.isDragging&&(o._neighbor=i,o.isDragging=!0,n(s,i),a.markerRenderer.draw(u.gridData[i.pointIndex][0],u.gridData[i.pointIndex][1],o._ctx),e.target.style.cursor="move",s.target.trigger("jqplotDragStart",[i.seriesIndex,i.pointIndex,t,r]))}else{var f=o._ctx;f.clearRect(0,0,f.canvas.width,f.canvas.height),o.isDragging=!1}}function s(e,t,n,r,i){if(i.plugins.dragable.dragCanvas.isDragging){var s=i.plugins.dragable.dragCanvas,o=s._ctx;o.clearRect(0,0,o.canvas.width,o.canvas.height),s.isDragging=!1;var u=s._neighbor,a=i.series[u.seriesIndex],f=a.plugins.dragable,l=f.constrainTo=="y"?u.data[0]:n[a.xaxis],c=f.constrainTo=="x"?u.data[1]:n[a.yaxis];a.data[u.pointIndex][0]=l,a.data[u.pointIndex][1]=c,i.drawSeries({preventJqPlotSeriesDrawTrigger:!0},u.seriesIndex),s._neighbor=null,e.target.style.cursor=s._cursors.pop(),i.target.trigger("jqplotDragStop",[t,n])}}e.jqplot.Dragable=function(t){this.markerRenderer=new e.jqplot.MarkerRenderer({shadow:!1}),this.shapeRenderer=new e.jqplot.ShapeRenderer,this.isDragging=!1,this.isOver=!1,this._ctx,this._elem,this._point,this._gridData,this.color,this.constrainTo="none",e.extend(!0,this,t)},t.prototype=new e.jqplot.GenericCanvas,t.prototype.constructor=t,e.jqplot.Dragable.parseOptions=function(t,n){var r=n||{};this.plugins.dragable=new e.jqplot.Dragable(r.dragable),this.isDragable=e.jqplot.config.enablePlugins},e.jqplot.Dragable.postPlotDraw=function(){this.plugins.dragable&&this.plugins.dragable.highlightCanvas&&(this.plugins.dragable.highlightCanvas.resetCanvas(),this.plugins.dragable.highlightCanvas=null),this.plugins.dragable={previousCursor:"auto",isOver:!1},this.plugins.dragable.dragCanvas=new t,this.eventCanvas._elem.before(this.plugins.dragable.dragCanvas.createElement(this._gridPadding,"jqplot-dragable-canvas",this._plotDimensions,this));var e=this.plugins.dragable.dragCanvas.setContext()},e.jqplot.preParseSeriesOptionsHooks.push(e.jqplot.Dragable.parseOptions),e.jqplot.postDrawHooks.push(e.jqplot.Dragable.postPlotDraw),e.jqplot.eventListenerHooks.push(["jqplotMouseMove",r]),e.jqplot.eventListenerHooks.push(["jqplotMouseDown",i]),e.jqplot.eventListenerHooks.push(["jqplotMouseUp",s])})(jQuery);