package counter

var testData = []byte(`
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#375EAB">

  <title>File compress - The Go Programming Language</title>

<link type="text/css" rel="stylesheet" href="/lib/godoc/style.css">

<link rel="search" type="application/opensearchdescription+xml" title="godoc" href="/opensearch.xml" />

<link rel="stylesheet" href="/lib/godoc/jquery.treeview.css">
<script>window.initFuncs = [];</script>
<script src="/lib/godoc/jquery.js" defer></script>
<script src="/lib/godoc/jquery.treeview.js" defer></script>
<script src="/lib/godoc/jquery.treeview.edit.js" defer></script>


<script src="/lib/godoc/playground.js" defer></script>

<script>var goVersion = "go1.10.1";</script>
<script src="/lib/godoc/godocs.js" defer></script>
<script type="text/javascript">
var _gaq = _gaq || [];
_gaq.push(["_setAccount", "UA-11222381-2"]);
_gaq.push(["b._setAccount", "UA-49880327-6"]);
window.trackPageview = function() {
  _gaq.push(["_trackPageview", location.pathname+location.hash]);
  _gaq.push(["b._trackPageview", location.pathname+location.hash]);
};
window.trackPageview();
window.trackEvent = function(category, action, opt_label, opt_value, opt_noninteraction) {
  _gaq.push(["_trackEvent", category, action, opt_label, opt_value, opt_noninteraction]);
  _gaq.push(["b._trackEvent", category, action, opt_label, opt_value, opt_noninteraction]);
};
</script>
</head>
<body>

<div id='lowframe' style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
...
</div><!-- #lowframe -->

<div id="topbar" class="wide"><div class="container">
<div class="top-heading" id="heading-wide"><a href="/">The Go Programming Language</a></div>
<div class="top-heading" id="heading-narrow"><a href="/">Go</a></div>
<a href="#" id="menu-button"><span id="menu-button-arrow">&#9661;</span></a>
<form method="GET" action="/search">
<div id="menu">
<a href="/doc/">Documents</a>
<a href="/pkg/">Packages</a>
<a href="/project/">The Project</a>
<a href="/help/">Help</a>

<a href="/blog/">Blog</a>


<a id="playgroundButton" href="http://play.golang.org/" title="Show Go Playground">Play</a>

<span class="search-box"><input type="search" id="search" name="q" placeholder="Search" aria-label="Search" required><button type="submit"><span><!-- magnifying glass: --><svg width="24" height="24" viewBox="0 0 24 24"><title>submit search</title><path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/><path d="M0 0h24v24H0z" fill="none"/></svg></span></button></span>
</div>
</form>

</div></div>


<div id="playground" class="play">
	<div class="input"><textarea class="code" spellcheck="false">package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}</textarea></div>
	<div class="output"></div>
	<div class="buttons">
		<a class="run" title="Run this code [shift-enter]">Run</a>
		<a class="fmt" title="Format this code">Format</a>
		
		<a class="share" title="Share this code">Share</a>
		
	</div>
</div>


<div id="page" class="wide">
<div class="container">


  <h1>
    File compress
    <span class="text-muted"></span>
  </h1>



  <h2>compress</h2>





<div id="nav"></div>


<!--
	Copyright 2009 The Go Authors. All rights reserved.
	Use of this source code is governed by a BSD-style
	license that can be found in the LICENSE file.
-->

<p>
<span class="alert" style="font-size:120%"> ../../go/compress: file does not exist</span>
</p>


<div id="footer">
Build version go1.10.1.<br>
Except as <a href="https://developers.google.com/site-policies#restrictions">noted</a>,
the content of this page is licensed under the
Creative Commons Attribution 3.0 License,
and code is licensed under a <a href="/LICENSE">BSD license</a>.<br>
<a href="/doc/tos.html">Terms of Service</a> |
<a href="http://www.google.com/intl/en/policies/privacy/">Privacy Policy</a>
</div>

</div><!-- .container -->
</div><!-- #page -->
<script type="text/javascript">
(function() {
  var ga = document.createElement("script"); ga.type = "text/javascript"; ga.async = true;
  ga.src = ("https:" == document.location.protocol ? "https://ssl" : "http://www") + ".google-analytics.com/ga.js";
  var s = document.getElementsByTagName("script")[0]; s.parentNode.insertBefore(ga, s);
})();
</script>
</body>
</html>

`)

var testUrls = []string{
	"https://golang.org/",
	"https://golang.org/doc/",
	"https://golang.org/pkg/compress/",
	"https://golang.org/pkg/compress/gzip/",
	"https://golang.org/pkg/crypto/md5/",
	"https://golang.org/pkg/debug/pe/",
	"https://golang.org/pkg/log/syslog/",
	"https://golang.org/pkg/sort/",
	"https://golang.org/pkg/strconv/",
	"https://golang.org/pkg/strings/",
	"https://golang.org/pkg/sync/",
	"https://golang.org/pkg/strings/",
	"https://golang.org/pkg/time/",
	"https://golang.org/pkg/unicode/",
	"https://golang.org/pkg/unsafe/",
	"https://godoc.org/golang.org/x/benchmarks",
	"https://godoc.org/golang.org/x/net",
	"https://godoc.org/golang.org/x/mobile",
}
