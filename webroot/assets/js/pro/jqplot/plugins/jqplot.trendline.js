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

(function(e){function t(e){var t=null;if(e.trendline&&e.trendline.show){var n=e.trendline.label.toString();n&&(t={label:n,color:e.trendline.color})}return t}function n(t,n,r,i,s){this._type&&(this._type==="line"||this._type=="bar")&&(this.trendline=new e.jqplot.Trendline,i=i||{},e.extend(!0,this.trendline,{color:this.color},r.trendline,i.trendline),this.trendline.renderer.init.call(this.trendline,null))}function r(t,n){n=e.extend(!0,{},this.trendline,n);if(this.trendline&&n.show){var r,i=n.data||this.data;r=u(i,this.trendline.type);var s=n.gridData||this.renderer.makeGridData.call(this,r.data);this.trendline.renderer.draw.call(this.trendline,t,s,{showLine:!0,shadow:this.trendline.shadow})}}function i(e,t,n){var r=n==null?"linear":n,i=e.length,s,o,u=0,a=0,f=0,l=0,c=0,h=[],p=[];if(r=="linear")p=e,h=t;else if(r=="exp"||r=="exponential")for(var d=0;d<t.length;d++)t[d]<=0?i--:(p.push(e[d]),h.push(Math.log(t[d])));for(var d=0;d<i;d++)u+=p[d],a+=h[d],l+=p[d]*h[d],f+=p[d]*p[d],c+=h[d]*h[d];return s=(i*l-u*a)/(i*f-u*u),o=(a-s*u)/i,[s,o]}function s(e,t){var n;return n=i(e,t,"linear"),[n[0],n[1]]}function o(e,t){var n,r=e,s=t;n=i(r,s,"exp");var o=Math.exp(n[0]),u=Math.exp(n[1]);return[o,u]}function u(e,t){var n=t==null?"linear":t,r,i,u=[],a=[],f=[];for(l=0;l<e.length;l++)e[l]!=null&&e[l][0]!=null&&e[l][1]!=null&&(u.push(e[l][0]),a.push(e[l][1]));if(n=="linear"){r=s(u,a);for(var l=0;l<u.length;l++)i=r[0]*u[l]+r[1],f.push([u[l],i])}else if(n=="exp"||n=="exponential"){r=o(u,a);for(var l=0;l<u.length;l++)i=r[1]*Math.pow(r[0],u[l]),f.push([u[l],i])}return{data:f,slope:r[0],intercept:r[1]}}e.jqplot.Trendline=function(){this.show=e.jqplot.config.enablePlugins,this.color="#666666",this.renderer=new e.jqplot.LineRenderer,this.rendererOptions={marker:{show:!1}},this.label="",this.type="linear",this.shadow=!0,this.markerRenderer={show:!1},this.lineWidth=1.5,this.shadowAngle=45,this.shadowOffset=1,this.shadowAlpha=.07,this.shadowDepth=3,this.isTrendline=!0},e.jqplot.postSeriesInitHooks.push(n),e.jqplot.postDrawSeriesHooks.push(r),e.jqplot.addLegendRowHooks.push(t)})(jQuery);