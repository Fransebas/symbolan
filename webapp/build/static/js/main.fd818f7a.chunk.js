(this.webpackJsonpwebapp=this.webpackJsonpwebapp||[]).push([[0],{11:function(e,t,a){},12:function(e,t,a){},19:function(e,t,a){"use strict";a.r(t);var n=a(0),l=a.n(n),r=a(5),c=a.n(r),o=(a(11),a(2)),i=(a(12),a(1)),m=a.n(i);function s(e){var t=e.formulaList,a=e.name,n=t.map((function(e){return l.a.createElement("label",null,l.a.createElement(m.a.Node,{inline:!0},e," "))}));return l.a.createElement("div",null,a," : ",n)}var u=function(){var e=Object(n.useState)(""),t=Object(o.a)(e,2),a=t[0],r=t[1],c=Object(n.useState)(""),i=Object(o.a)(c,2),u=i[0],E=i[1],p=Object(n.useState)(""),d=Object(o.a)(p,2),h=d[0],b=d[1],v=Object(n.useState)(null),x=Object(o.a)(v,2);return x[0],x[1],l.a.createElement("div",{className:"App"},l.a.createElement("div",{className:"column"},l.a.createElement("div",null,l.a.createElement("h1",{style:{margin:"0px",bottomMargin:"0px"}}," Symbolan "),l.a.createElement("h6",{style:{margin:"0px"}}," (Very original name) ")),l.a.createElement("div",null,l.a.createElement("h3",null," Source code here (WIP) "),l.a.createElement("a",{href:"https://github.com/Fransebas/symbolan"}," https://github.com/Fransebas/symbolan ")),l.a.createElement("h3",null," Derivate "),l.a.createElement("div",{className:"flexItem inputDiv"},l.a.createElement("input",{value:a,onChange:function(e){r(e.target.value)}}),l.a.createElement("button",{id:"submitButton",onClick:function(){var e=new XMLHttpRequest,t={};e.addEventListener("load",(function(){t=JSON.parse(e.responseText),E(t.Original),b(t.Result)})),e.open("POST","http://localhost:7777/derivative"),e.send(a)}}," Derivate ")),l.a.createElement(m.a.Context,{input:"tex"},l.a.createElement("div",{id:"resultBox"},l.a.createElement("h5",{className:"resultTitle"}," Input: "),l.a.createElement(m.a.Node,null,u," "),l.a.createElement("h5",{className:"resultTitle"}," Differential: "),l.a.createElement(m.a.Node,null,h," "),l.a.createElement("div",{className:"bottomBox"}))),l.a.createElement("div",{className:"explanation"},l.a.createElement("h5",{className:"resultTitle"}," Valid functions "),l.a.createElement(m.a.Context,{input:"tex"},l.a.createElement("div",null,l.a.createElement(s,{formulaList:["+","-","*","/","^"],name:"Elementary Functions"}),l.a.createElement(s,{formulaList:["sin(x)","cos(x)","tan(x)"],name:"Trigonometric"}),l.a.createElement(s,{formulaList:["sinh(x)","cosh(x)","tanh(x)"],name:"Hiperbolics"}),l.a.createElement(s,{formulaList:["log(x)","e^x"],name:"Logs and exponentials"}),l.a.createElement(s,{formulaList:["x","y","z","x_{56}"],name:"Variables"}),l.a.createElement(s,{formulaList:["a","b","c","a_{88}"],name:"Constants"}),l.a.createElement("p",null," Where ",l.a.createElement(m.a.Node,{inline:!0}," x ")," can be any function "),l.a.createElement("h5",null," Example:"),l.a.createElement("p",null," sin(7x^2) + log(sinh(y)) + a_9 "))),l.a.createElement("div",{className:"bottomBox"}))))};Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));c.a.render(l.a.createElement(l.a.StrictMode,null,l.a.createElement(u,null)),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(e){e.unregister()})).catch((function(e){console.error(e.message)}))},6:function(e,t,a){e.exports=a(19)}},[[6,1,2]]]);
//# sourceMappingURL=main.fd818f7a.chunk.js.map