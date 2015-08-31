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

(function(e){e.jqplot.ciParser=function(t,n){function l(e,t){var n;if(t!=null){if(t.toString().indexOf("Date")>=0){n=/^\/Date\((-?[0-9]+)\)\/$/.exec(t);if(n)return parseInt(n[1],10)}return t}}var r=[],i,s,o,u,a,f;if(typeof t=="string")t=e.jqplot.JSON.parse(t,l);else{if(typeof t!="object")return null;for(a in t)for(o=0;o<t[a].length;o++)for(f in t[a][o])t[a][o][f]=l(f,t[a][o][f])}for(var c in t){i=[],s=t[c];switch(c){case"PriceTicks":for(o=0;o<s.length;o++)i.push([s[o].TickDate,s[o].Price]);break;case"PriceBars":for(o=0;o<s.length;o++)i.push([s[o].BarDate,s[o].Open,s[o].High,s[o].Low,s[o].Close])}r.push(i)}return r}})(jQuery);