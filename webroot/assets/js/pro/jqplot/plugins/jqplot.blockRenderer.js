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

(function(e){e.jqplot.BlockRenderer=function(){e.jqplot.LineRenderer.call(this)},e.jqplot.BlockRenderer.prototype=new e.jqplot.LineRenderer,e.jqplot.BlockRenderer.prototype.constructor=e.jqplot.BlockRenderer,e.jqplot.BlockRenderer.prototype.init=function(t){this.css={padding:"2px",border:"1px solid #999",textAlign:"center"},this.escapeHtml=!1,this.insertBreaks=!0,this.varyBlockColors=!1,e.extend(!0,this,t),this.css.backgroundColor?this.color=this.css.backgroundColor:this.css.background?this.color=this.css.background:this.varyBlockColors||(this.css.background=this.color),this.canvas=new e.jqplot.BlockCanvas,this.shadowCanvas=new e.jqplot.BlockCanvas,this.canvas._plotDimensions=this._plotDimensions,this.shadowCanvas._plotDimensions=this._plotDimensions,this._type="block",this.moveBlock=function(e,t,n,r){var i=this.canvas._elem.children(":eq("+e+")");this.data[e][0]=t,this.data[e][1]=n,this._plotData[e][0]=t,this._plotData[e][1]=n,this._stackData[e][0]=t,this._stackData[e][1]=n,this.gridData[e][0]=this._xaxis.series_u2p(t),this.gridData[e][1]=this._yaxis.series_u2p(n);var s=i.outerWidth(),o=i.outerHeight(),u=this.gridData[e][0]-s/2+"px",a=this.gridData[e][1]-o/2+"px";r?(parseInt(r,10)&&(r=parseInt(r,10)),i.animate({left:u,top:a},r)):i.css({left:u,top:a}),i=null}},e.jqplot.BlockRenderer.prototype.draw=function(t,n,r){this.plugins.pointLabels&&(this.plugins.pointLabels.show=!1);var i,s,o,n,u,a,f,l,c,h,p=r!=undefined?r:{},d=new e.jqplot.ColorGenerator(this.seriesColors);this.canvas._elem.empty();for(i=0;i<this.gridData.length;i++)o=this.data[i],n=this.gridData[i],u="",a={},typeof o[2]=="string"?u=o[2]:typeof o[2]=="object"&&(a=o[2]),typeof o[3]=="object"&&(a=o[3]),this.insertBreaks&&(u=u.replace(/ /g,"<br />")),a=e.extend(!0,{},this.css,a),s=e('<div style="position:absolute;margin-left:auto;margin-right:auto;"></div>'),this.canvas._elem.append(s),this.escapeHtml?s.text(u):s.html(u),delete a.position,delete a.marginRight,delete a.marginLeft,!a.background&&!a.backgroundColor&&!a.backgroundImage&&(a.background=d.next()),s.css(a),f=s.outerWidth(),l=s.outerHeight(),c=n[0]-f/2+"px",h=n[1]-l/2+"px",s.css({left:c,top:h}),s=null},e.jqplot.BlockCanvas=function(){e.jqplot.ElemContainer.call(this),this._ctx},e.jqplot.BlockCanvas.prototype=new e.jqplot.ElemContainer,e.jqplot.BlockCanvas.prototype.constructor=e.jqplot.BlockCanvas,e.jqplot.BlockCanvas.prototype.createElement=function(t,n,r){this._offsets=t;var i="jqplot-blockCanvas";n!=undefined&&(i=n);var s;this._elem?s=this._elem.get(0):s=document.createElement("div"),r!=undefined&&(this._plotDimensions=r);var o=this._plotDimensions.width-this._offsets.left-this._offsets.right+"px",u=this._plotDimensions.height-this._offsets.top-this._offsets.bottom+"px";return this._elem=e(s),this._elem.css({position:"absolute",width:o,height:u,left:this._offsets.left,top:this._offsets.top}),this._elem.addClass(i),this._elem},e.jqplot.BlockCanvas.prototype.setContext=function(){return this._ctx={canvas:{width:0,height:0},clearRect:function(){return null}},this._ctx}})(jQuery);