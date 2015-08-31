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

(function(e){e.jqplot.EnhancedLegendRenderer=function(){e.jqplot.TableLegendRenderer.call(this)},e.jqplot.EnhancedLegendRenderer.prototype=new e.jqplot.TableLegendRenderer,e.jqplot.EnhancedLegendRenderer.prototype.constructor=e.jqplot.EnhancedLegendRenderer,e.jqplot.EnhancedLegendRenderer.prototype.init=function(t){this.numberRows=null,this.numberColumns=null,this.seriesToggle="normal",this.seriesToggleReplot=!1,this.disableIEFading=!0,e.extend(!0,this,t),this.seriesToggle&&e.jqplot.postDrawHooks.push(n)},e.jqplot.EnhancedLegendRenderer.prototype.draw=function(n,r){var i=this;if(this.show){var s=this._series,o,u="position:absolute;";u+=this.background?"background:"+this.background+";":"",u+=this.border?"border:"+this.border+";":"",u+=this.fontSize?"font-size:"+this.fontSize+";":"",u+=this.fontFamily?"font-family:"+this.fontFamily+";":"",u+=this.textColor?"color:"+this.textColor+";":"",u+=this.marginTop!=null?"margin-top:"+this.marginTop+";":"",u+=this.marginBottom!=null?"margin-bottom:"+this.marginBottom+";":"",u+=this.marginLeft!=null?"margin-left:"+this.marginLeft+";":"",u+=this.marginRight!=null?"margin-right:"+this.marginRight+";":"",this._elem=e('<table class="jqplot-table-legend" style="'+u+'"></table>'),this.seriesToggle&&this._elem.css("z-index","3");var a=!1,f=!1,l,c;this.numberRows?(l=this.numberRows,this.numberColumns?c=this.numberColumns:c=Math.ceil(s.length/l)):this.numberColumns?(c=this.numberColumns,l=Math.ceil(s.length/this.numberColumns)):(l=s.length,c=1);var h,p,d,v,m,g,y,b,w,E,S=0;for(h=s.length-1;h>=0;h--)if(c==1&&s[h]._stack||s[h].renderer.constructor==e.jqplot.BezierCurveRenderer)f=!0;for(h=0;h<l;h++){d=e(document.createElement("tr")),d.addClass("jqplot-table-legend"),f?d.prependTo(this._elem):d.appendTo(this._elem);for(p=0;p<c;p++){if(S<s.length&&(s[S].show||s[S].showLabel)){o=s[S],g=this.labels[S]||o.label.toString();if(g){var x=o.color;f?h==l-1?a=!1:a=!0:h>0?a=!0:a=!1,y=a?this.rowSpacing:"0",v=e(document.createElement("td")),v.addClass("jqplot-table-legend jqplot-table-legend-swatch"),v.css({textAlign:"center",paddingTop:y}),w=e(document.createElement("div")),w.addClass("jqplot-table-legend-swatch-outline"),E=e(document.createElement("div")),E.addClass("jqplot-table-legend-swatch"),E.css({backgroundColor:x,borderColor:x}),v.append(w.append(E)),m=e(document.createElement("td")),m.addClass("jqplot-table-legend jqplot-table-legend-label"),m.css("paddingTop",y),this.escapeHtml?m.text(g):m.html(g),f?(this.showLabels&&m.prependTo(d),this.showSwatches&&v.prependTo(d)):(this.showSwatches&&v.appendTo(d),this.showLabels&&m.appendTo(d));if(this.seriesToggle){var T;if(typeof this.seriesToggle=="string"||typeof this.seriesToggle=="number")if(!e.jqplot.use_excanvas||!this.disableIEFading)T=this.seriesToggle;this.showSwatches&&(v.bind("click",{series:o,speed:T,plot:r,replot:this.seriesToggleReplot},t),v.addClass("jqplot-seriesToggle")),this.showLabels&&(m.bind("click",{series:o,speed:T,plot:r,replot:this.seriesToggleReplot},t),m.addClass("jqplot-seriesToggle")),!o.show&&o.showLabel&&(v.addClass("jqplot-series-hidden"),m.addClass("jqplot-series-hidden"))}a=!0}}S++}v=m=w=E=null}}return this._elem};var t=function(t){var n=t.data,r=n.series,i=n.replot,s=n.plot,o=n.speed,u=r.index,a=!1;if(r.canvas._elem.is(":hidden")||!r.show)a=!0;var f=function(){if(i){var t={};e.isPlainObject(i)&&e.extend(!0,t,i),s.replot(t);if(a&&o){var n=s.series[u];n.shadowCanvas._elem&&n.shadowCanvas._elem.hide().fadeIn(o),n.canvas._elem.hide().fadeIn(o),n.canvas._elem.nextAll(".jqplot-point-label.jqplot-series-"+n.index).hide().fadeIn(o)}}else{var n=s.series[u];n.canvas._elem.is(":hidden")||!n.show?((typeof s.options.legend.showSwatches=="undefined"||s.options.legend.showSwatches===!0)&&s.legend._elem.find("td").eq(u*2).addClass("jqplot-series-hidden"),(typeof s.options.legend.showLabels=="undefined"||s.options.legend.showLabels===!0)&&s.legend._elem.find("td").eq(u*2+1).addClass("jqplot-series-hidden")):((typeof s.options.legend.showSwatches=="undefined"||s.options.legend.showSwatches===!0)&&s.legend._elem.find("td").eq(u*2).removeClass("jqplot-series-hidden"),(typeof s.options.legend.showLabels=="undefined"||s.options.legend.showLabels===!0)&&s.legend._elem.find("td").eq(u*2+1).removeClass("jqplot-series-hidden"))}};r.toggleDisplay(t,f)},n=function(){if(this.legend.renderer.constructor==e.jqplot.EnhancedLegendRenderer&&this.legend.seriesToggle){var t=this.legend._elem.detach();this.eventCanvas._elem.after(t)}}})(jQuery);