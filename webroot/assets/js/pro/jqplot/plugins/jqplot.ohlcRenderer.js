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

(function(e){e.jqplot.OHLCRenderer=function(){e.jqplot.LineRenderer.call(this),this.candleStick=!1,this.tickLength="auto",this.bodyWidth="auto",this.openColor=null,this.closeColor=null,this.wickColor=null,this.fillUpBody=!1,this.fillDownBody=!0,this.upBodyColor=null,this.downBodyColor=null,this.hlc=!1,this.lineWidth=1.5,this._tickLength,this._bodyWidth},e.jqplot.OHLCRenderer.prototype=new e.jqplot.LineRenderer,e.jqplot.OHLCRenderer.prototype.constructor=e.jqplot.OHLCRenderer,e.jqplot.OHLCRenderer.prototype.init=function(t){t=t||{},this.lineWidth=t.lineWidth||1.5,e.jqplot.LineRenderer.prototype.init.call(this,t),this._type="ohlc";var n=this._yaxis._dataBounds,r=this._plotData;if(r[0].length<5){this.renderer.hlc=!0;for(var i=0;i<r.length;i++){if(r[i][2]<n.min||n.min==null)n.min=r[i][2];if(r[i][1]>n.max||n.max==null)n.max=r[i][1]}}else for(var i=0;i<r.length;i++){if(r[i][3]<n.min||n.min==null)n.min=r[i][3];if(r[i][2]>n.max||n.max==null)n.max=r[i][2]}},e.jqplot.OHLCRenderer.prototype.draw=function(t,n,r){var i=this.data,s=this._xaxis.min,o=this._xaxis.max,u=0,a=i.length,f=this._xaxis.series_u2p,l=this._yaxis.series_u2p,c,h,p,d,v,m,g,y,b,w=this.renderer,E=r!=undefined?r:{},S=E.shadow!=undefined?E.shadow:this.shadow,x=E.fill!=undefined?E.fill:this.fill,T=E.fillAndStroke!=undefined?E.fillAndStroke:this.fillAndStroke;w.bodyWidth=E.bodyWidth!=undefined?E.bodyWidth:w.bodyWidth,w.tickLength=E.tickLength!=undefined?E.tickLength:w.tickLength,t.save();if(this.show){var N,C,k,L,A;for(var c=0;c<i.length;c++)i[c][0]<s?u=c:i[c][0]<o&&(a=c+1);var O=this.gridData[a-1][0]-this.gridData[u][0],M=a-u;try{var _=Math.abs(this._xaxis.series_u2p(parseInt(this._xaxis._intervalStats[0].sortedIntervals[0].interval,10))-this._xaxis.series_u2p(0))}catch(D){var _=O/M}w.candleStick?typeof w.bodyWidth=="number"?w._bodyWidth=w.bodyWidth:w._bodyWidth=Math.min(20,_/1.65):typeof w.tickLength=="number"?w._tickLength=w.tickLength:w._tickLength=Math.min(10,_/3.5);for(var c=u;c<a;c++)N=f(i[c][0]),w.hlc?(C=null,k=l(i[c][1]),L=l(i[c][2]),A=l(i[c][3])):(C=l(i[c][1]),k=l(i[c][2]),L=l(i[c][3]),A=l(i[c][4])),b={},w.candleStick&&!w.hlc?(m=w._bodyWidth,g=N-m/2,A<C?(w.wickColor?b.color=w.wickColor:w.downBodyColor&&(b.color=w.upBodyColor),p=e.extend(!0,{},E,b),w.shapeRenderer.draw(t,[[N,k],[N,A]],p),w.shapeRenderer.draw(t,[[N,C],[N,L]],p),b={},d=A,v=C-A,w.fillUpBody?b.fillRect=!0:(b.strokeRect=!0,m-=this.lineWidth,g=N-m/2),w.upBodyColor&&(b.color=w.upBodyColor,b.fillStyle=w.upBodyColor),y=[g,d,m,v]):A>C?(w.wickColor?b.color=w.wickColor:w.downBodyColor&&(b.color=w.downBodyColor),p=e.extend(!0,{},E,b),w.shapeRenderer.draw(t,[[N,k],[N,C]],p),w.shapeRenderer.draw(t,[[N,A],[N,L]],p),b={},d=C,v=A-C,w.fillDownBody?b.fillRect=!0:(b.strokeRect=!0,m-=this.lineWidth,g=N-m/2),w.downBodyColor&&(b.color=w.downBodyColor,b.fillStyle=w.downBodyColor),y=[g,d,m,v]):(w.wickColor&&(b.color=w.wickColor),p=e.extend(!0,{},E,b),w.shapeRenderer.draw(t,[[N,k],[N,L]],p),b={},b.fillRect=!1,b.strokeRect=!1,g=[N-m/2,C],d=[N+m/2,A],m=null,v=null,y=[g,d]),p=e.extend(!0,{},E,b),w.shapeRenderer.draw(t,y,p)):(h=E.color,w.openColor&&(E.color=w.openColor),w.hlc||w.shapeRenderer.draw(t,[[N-w._tickLength,C],[N,C]],E),E.color=h,w.wickColor&&(E.color=w.wickColor),w.shapeRenderer.draw(t,[[N,k],[N,L]],E),E.color=h,w.closeColor&&(E.color=w.closeColor),w.shapeRenderer.draw(t,[[N,A],[N+w._tickLength,A]],E),E.color=h)}t.restore()},e.jqplot.OHLCRenderer.prototype.drawShadow=function(e,t,n){},e.jqplot.OHLCRenderer.checkOptions=function(e,t,n){n.highlighter||(n.highlighter={showMarker:!1,tooltipAxes:"y",yvalues:4,formatString:'<table class="jqplot-highlighter"><tr><td>date:</td><td>%s</td></tr><tr><td>open:</td><td>%s</td></tr><tr><td>hi:</td><td>%s</td></tr><tr><td>low:</td><td>%s</td></tr><tr><td>close:</td><td>%s</td></tr></table>'})}})(jQuery);