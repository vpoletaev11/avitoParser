package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vpoletaev11/avitoParser/scrapper"
)

// NewDepAndServer returns: dependencies, Sqlmock interface to add mocks and httptest server, that emulating avito ad page
func NewDepAndServer() (scrapper.Dep, sqlmock.Sqlmock, *httptest.Server) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, avitoAdHTML)
	}))

	dep := scrapper.Dep{
		DB:     db,
		Client: ts.Client(),
	}

	return dep, sqlMock, ts
}

const avitoAdHTML = `<!DOCTYPE html>
            
<html> <head> <script>
 try {
 window.firstHiddenTime = document.visibilityState === 'hidden' ? 0 : Infinity;
 document.addEventListener('visibilitychange', function (event) {
 window.firstHiddenTime = Math.min(window.firstHiddenTime, event.timeStamp);
 }, { once: true });
 if ('PerformanceLongTaskTiming' in window) {
 var globalStats = window.__statsLongTasks = { tasks: [] };
 globalStats.observer = new PerformanceObserver(function(list) {
 globalStats.tasks = globalStats.tasks.concat(list.getEntries());
 });
 globalStats.observer.observe({ entryTypes: ['longtask'] });
 }
 if (PerformanceObserver && (PerformanceObserver.supportedEntryTypes || []).some(function(e) {
 return e === 'element'
 })) {
 if (!window.oet) {
 window.oet = [];
 }
 new PerformanceObserver(function(l) {
 window.oet.push.apply(window.oet, l.getEntries());
 }).observe({ entryTypes: ['element'] });
 }
 } catch (e) {
 console.error(e);
 }
 </script>
    <script>
 window.dataLayer = [{"dynx_user":"a","dynx_region":"moskva","dynx_prodid":2009709110,"dynx_price":1000000,"dynx_category":"kollektsionirovanie","dynx_vertical":4,"dynx_pagetype":"item"},{"pageType":"Item","itemID":2009709110,"vertical":"GENERAL","categoryId":36,"categorySlug":"kollektsionirovanie","microCategoryId":320,"locationId":637640,"isShop":0,"isClientType1":0,"itemPrice":1000000,"withDelivery":1,"sostoyanie":"Б\/у","vid_tovara":"Монеты","type_of_trade":"Продаю своё"}];
 (function(w, d, s, l, i) {
 w[l] = w[l] || [];
 w[l].push({
 'gtm.start': new Date().getTime(),
 event: 'gtm.js'
 });
 var f = d.getElementsByTagName(s)[0],
 j = d.createElement(s),
 dl = l != 'dataLayer' ? '&l=' + l : '';
 j.async = true;
 j.src = '//www.googletagmanager.com/gtm.js?id=' + i + dl;
 f.parentNode.insertBefore(j, f);
 })(window, document, 'script', 'dataLayer', 'GTM-KP9Q9H');
 </script>
  
 <meta charset="utf-8">
   <meta name="format-detection" content="telephone=no">  <meta name="google-site-verification" content="7iEzRRMJ2_0p66pVS7wTYYvhZZSFBdzL5FVml4IKUS0" />
            <link rel="alternate" media="only screen and (max-width: 640px)" href="https://m.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110">
      <title>Монета из макдоналдса купить в Москве | Хобби и отдых | Авито</title>
   <link rel="dns-prefetch" href="//an.yandex.ru/">
  <link rel="preload" href="//an.yandex.ru/system/context.js" as="script">
  <script src="//an.yandex.ru/system/context.js" async></script>
    <link rel="dns-prefetch" href="//yastatic.net/" >
  <link rel="preload" href="//yastatic.net/pcode/adfox/loader.js" as="script">
  <script src="//yastatic.net/pcode/adfox/loader.js" async></script>
   <link rel="dns-prefetch" href="//securepubads.g.doubleclick.net/" >
   <link rel="preload" href="//www.googletagservices.com/tag/js/gpt.js" as="script">
  <script src="//www.googletagservices.com/tag/js/gpt.js" async></script>
      <meta name="yandex-verification" content="499bdc75d3636c55" /><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/edcf951a50e7eae10aeb.css"><script src="https://static.avito.ru/s/cc/bundles/9a7c46aa215ff01519cc.js" ></script><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/03e71f68fc31b7c38ef3.css"><script src="https://static.avito.ru/s/cc/bundles/3d308baee20bc5fa8cb3.js" ></script><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/fd12012944cd8b497cfc.css"><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/c2db2eef73d0e6f185a5.css"><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/c9136147f2896e653e00.css"><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/12d40f81b7e415ffe00e.css"><script src="https://static.avito.ru/s/cc/bundles/54287266aa808b5a71e5.js" ></script><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/7a1ee2a151af71ca5ee9.css"><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/c3298eb836d19670aaed.css"><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/7740a44321f872ffa8e6.css"><script src="https://static.avito.ru/s/cc/bundles/da8ad765686aa4a78d1b.js" ></script><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/eed641c3b0c0d2537cef.css"><script src="https://static.avito.ru/s/cc/bundles/8ca2773f4359b2483f95.js" ></script><link rel="stylesheet" href="https://static.avito.ru/@avito/sd-item-view/2.4.0/prod/web/styles/da3e05a67ce211513054.css"><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/909a107b39faef2b89f1.css"><script src="https://static.avito.ru/s/cc/bundles/f6486134a2ff82ae03d0.js" ></script><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/a2c96a91b5c90c14964f.css"><script src="https://static.avito.ru/s/cc/bundles/de1f5bb87a3d9e6cee72.js" ></script><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/73f729913972c9bea736.css"><script src="https://static.avito.ru/s/cc/bundles/b080a0729918eac78279.js" ></script><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/9f79063bd9b93af23d78.css"><script src="https://static.avito.ru/s/cc/bundles/c69d96ebf1286d9032b6.js" ></script><link rel="stylesheet" href="https://static.avito.ru/@avito/au-discount/1.0.0/prod/web/styles/d0b173e67e036b2c31d1.css"><link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/d0c2b7cb570e4aa104c5.css"><script src="https://static.avito.ru/s/cc/bundles/8489c81ca1a5b6b32f44.js" ></script><link rel="apple-touch-icon-precomposed" sizes="180x180" href="https://www.avito.st/s/common/touch-icons/common/apple-touch-icon-180x180-precomposed.png?57be3fb" /><link rel="apple-touch-icon-precomposed" sizes="152x152" href="https://www.avito.st/s/common/touch-icons/common/apple-touch-icon-152x152-precomposed.png?cac4f2a" /><link rel="apple-touch-icon-precomposed" sizes="144x144" href="https://www.avito.st/s/common/touch-icons/common/apple-touch-icon-144x144-precomposed.png?9615e61" /><link rel="apple-touch-icon-precomposed" sizes="120x120" href="https://www.avito.st/s/common/touch-icons/common/apple-touch-icon-120x120-precomposed.png?2a32f09" /><link rel="apple-touch-icon-precomposed" sizes="114x114" href="https://www.avito.st/s/common/touch-icons/common/apple-touch-icon-114x114-precomposed.png?174e153" /><link rel="apple-touch-icon-precomposed" sizes="76x76" href="https://www.avito.st/s/common/touch-icons/common/apple-touch-icon-76x76-precomposed.png?28e6cfb" /><link rel="apple-touch-icon-precomposed" sizes="72x72" href="https://www.avito.st/s/common/touch-icons/common/apple-touch-icon-72x72-precomposed.png?aeb90b3" /><link rel="apple-touch-icon-precomposed" sizes="57x57" href="https://www.avito.st/s/common/touch-icons/common/apple-touch-icon-57x57-precomposed.png?fd7ac94" /><meta name="msapplication-TileColor" content="#000000"><meta name="msapplication-TileImage" content="/s/common/touch-icons/common/mstile-144x144.png"><meta name="msapplication-config" content="browserconfig.xml" /><link href="https://www.avito.st/favicon.ico?9de48a5" rel="shortcut icon" type="image/x-icon" /><link href="https://www.avito.st/ya-tableau-manifest-ru.json?5ac8b8a" rel="yandex-tableau-widget" /><link href="https://www.avito.st/open_search_ru.xml?4b0fd3d" rel="search" type="application/opensearchdescription+xml" title="Авито" /><meta property="og:description" content="Монета из макдоналдса: объявление о продаже в Москве на Авито. Отсуствует бумажная обертка" /><meta name="description" content="Монета из макдоналдса: объявление о продаже в Москве на Авито. Отсуствует бумажная обертка" /><meta name="mrc__share_title" content="Монета из макдоналдса купить в Москве | Хобби и отдых | Авито" /><meta name="mrc__share_description" content="Монета из макдоналдса: объявление о продаже в Москве на Авито. Отсуствует бумажная обертка" /><link rel="image_src" href="https://www.avito.ru/img/share/auto/9582784370" /><meta property="og:title" content="Монета из макдоналдса купить в Москве | Хобби и отдых | Авито" /><meta property="og:type" content="website" /><meta property="og:url" content="https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110" /><meta property="og:site_name" content="Авито" /><meta property="og:locale" content="ru_RU" /><meta property="fb:app_id" content="472292516308756" /><meta property="product:price:amount" content="1000000" /><meta property="product:price:currency" content="RUB"/><meta property="og:image" content="https://www.avito.ru/img/share/auto/9582784370" /><meta property="og:image:alt" content="Монета из макдоналдса" /><meta property="og:image" content="https://www.avito.ru/img/share/auto/9582785024" /><meta property="og:image:alt" content="Монета из макдоналдса" /><meta property="og:image" content="https://www.avito.ru/img/share/auto/9582784370" /><meta property="og:image:alt" content="Монета из макдоналдса" /><link rel="canonical" href="https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110" /><link rel="alternate" href="android-app&#x3A;&#x2F;&#x2F;com.avito.android&#x2F;ru.avito&#x2F;1&#x2F;items&#x2F;2009709110" /><link rel="alternate" href="ios-app&#x3A;&#x2F;&#x2F;417281773&#x2F;ru.avito&#x2F;1&#x2F;items&#x2F;2009709110" />  <script src="https://static.avito.ru/s/cc/chunks/23865871b706fee44dda.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/11616c6625b532e7659e.js" ></script>
 <script src="https://static.avito.ru/s/cc/chunks/1336ad25f42de143db58.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/5d86817686121c006504.js" ></script>
    <script>
 window.capturedErrors = [];
 window.addEventListener('error', function(error) {
 window.capturedErrors.push(error);
 });
 </script> <script>
 window.avito = window.avito || {};
 window.avito.platform = 'desktop';
 window.avito.siteName = 'Авито';
 window.avito.staticPrefix = 'https://www.avito.st';
 window.avito.supportPrefix = 'https://support.avito.ru';
 window.avito.pageId = 'item';
 window.avito.categoryId = 36;
 window.avito.locationId = null;
 window.avito.fromPage = '';
 window.avito.sentry = {
 dsn: 'https://1ede3b886c8b4efd9230509682fe2f12@sntr.avito.ru/41',
 release: "rc-202010121331-172125"
 };
 window.avito.clickstream = {
 buildId: "rc-202010121331-172125",
 buildUid: null
 };
 window.avito.isAuthenticated = '' === '1';
 window.avito.isVerifiedCompany = '' === '1';
 window.avito.metrics = {};
 window.avito.metrics.categoryId = 36;
 window.avito.metrics.browser = '.0';
 window.avito.filtersGroup = 'desktop_catalog_filters';
 window.avito.experiments = window.avito.experiments || {};
 window.avito.abFeatures = window.avito.abFeatures || {};
 // Messenger config section
 window.avito.socketserver = 'wss://socket.avito.ru/socket';
 window.avito.httpfallback = '';
 window.avito.socketImageUploadUrl = 'https://socket.avito.ru/images';
 window.avito.socketMetricsUrl = 'https://socket.avito.ru/metrics';
 window.avito.isTestAccount = '' === '1';
 window.avito.userId = '0';
  </script>
  <script src="/s/a/j/dfp/px.js?ch=1"></script> <script src="/s/a/j/dfp/px.js?ch=2"></script>
     
 <script>
 window.avito.ads = {
 userGroup: 49
 };
 </script>
  <script src="https://static.avito.ru/s/cc/chunks/f08e81b0b81c9ed89b2e.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/d944b92d78da83c396a7.js" ></script>
  <script src="https://static.avito.ru/s/cc/bundles/29c93c440a7039dd499d.js" ></script>
  <script src="https://static.avito.ru/s/cc/bundles/789cc9cbb491c1b8a08e.js" ></script>
   <script src="https://static.avito.ru/s/cc/chunks/e0da6af759ed6c02053f.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/da88e189e4172c9a136c.js" ></script>
   <script src="https://static.avito.ru/s/cc/bundles/2d9826e2a23d7bc2f5c7.js" ></script>
  <script src="https://static.avito.ru/s/cc/bundles/8e8838711c32bb27965c.js" ></script>
 <script src="https://static.avito.ru/s/cc/bundles/744e0f5353f043f08632.js" ></script>
  
 <script>
 var avito = avito || {};
 avito.item = avito.item || {};
 avito.abFeatures = avito.abFeatures || {};
 avito.item.url = '/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110';
 avito.item.id = '2009709110';
 avito.item.rootCategoryId = 7;
 avito.item.price = '1000000';
 avito.item.countryHost = 'www.avito.ru';
 avito.item.siteName = 'Авито';
 avito.item.isRealty = 0;
 avito.item.isMyItem = 0;
 avito.item.hasCvPackage = 0;
 avito.item.tokenName = 'token[2496499932144]';
 avito.item.tokenValue = '20548fab04c2aace';
 avito.item.searchHash = '';
 avito.item.userHashId = '119541467';
 avito.item.image = 'https://70.img.avito.st/208x156/9582784370.jpg';
 avito.item.location = '\u041C\u043E\u0441\u043A\u0432\u0430';
 avito.item.title = '\u041C\u043E\u043D\u0435\u0442\u0430\u0020\u0438\u0437\u0020\u043C\u0430\u043A\u0434\u043E\u043D\u0430\u043B\u0434\u0441\u0430';
 avito.item.priceFormatted = '1 000 000&nbsp;₽';
 avito.item.vin = '';
 avito.item.locationId = 637640;
 avito.item.categoryId = 36;
                 avito.abFeatures.isCriteoTestTransactionsDefaultGroup = true;       </script>
    <link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/627a87e060f07ac42424.css">
<script src="https://static.avito.ru/s/cc/bundles/4527b754df3986f12bd9.js" ></script>
  <script>
 window.avito = window.avito || {};
 window.avito.banners = {"ldr_top":[{"sellingSystem":"AdFox","parameters":{"block_id":"8_bztzl-fqry","stat_id":"100003649","params":{"puid2":1,"puid18":"item","puid1":49,"puid7":2207,"puid5":7,"puid33":0,"puid14":"79.165.72.11","puid41":"\/moskva\/kollektsionirovanie\/moneta_iz_makdonaldsa_2009709110","puid15":"77","puid6":36,"puid4":637640,"puid3":637640,"puid26":0,"puid30":"0","puid31":0,"puid23":"1000000-1499999","puid25":1,"puid34":"Монета из макдоналдса","puid35":1000000,"puid36":"\/\/70.img.avito.st\/640x480\/9582784370.jpg","puid63":"46"},"alid":"62a6297cd89b68466e283fb7177402af"}},{"sellingSystem":"DFP","parameters":{"yandex_rtb_low_block_id":"R-A-419506-33","yandex_rtb_low_stat_id":"100003649","dfp_slot":"\/7870\/AR\/hobbi_i_otdyh\/kollektsionirovanie\/Item\/ldr_top","google_adexchange_pf":"10","sizes":[[1000,120],[1000,90],[970,90],[728,90],[980,120],[960,90],[950,90],[1,1],[900,90],[1000,200],[644,105]],"count":1,"targetings":{"bp":"ldr_top"},"isrmp":1,"page_type":"item","AudRandom":49,"keyword":"","g_metro":2207,"master_category":7,"pmin":null,"pmax":null,"cat_interests":[],"params_interests":[],"socd":[],"abp":0,"ip":"79.165.72.11","referer":"\/moskva\/kollektsionirovanie\/moneta_iz_makdonaldsa_2009709110","item_property":["77"],"slave_category":36,"g_country":621540,"g_reg":637640,"g_city":637640,"g_domofond_reg":"moskva-c3584","Usid":0,"stlogin":"0","phcodedesc":null,"Uname_Get4C":"","has_shop":0,"company_id":null,"price":"1000000-1499999","stcompany":1,"par_title":"Монета из макдоналдса","par_price":1000000,"par_picture":"\/\/70.img.avito.st\/640x480\/9582784370.jpg","alid":"62a6297cd89b68466e283fb7177402af"}}],"tgb1":[{"sellingSystem":"AdFox","parameters":{"block_id":"8_bzubs-fvlq","stat_id":"100003649","params":{"puid2":1,"puid18":"item","puid1":49,"puid7":2207,"puid5":7,"puid33":0,"puid14":"79.165.72.11","puid41":"\/moskva\/kollektsionirovanie\/moneta_iz_makdonaldsa_2009709110","puid15":"77","puid6":36,"puid4":637640,"puid3":637640,"puid26":0,"puid30":"0","puid31":0,"puid23":"1000000-1499999","puid25":1,"puid34":"Монета из макдоналдса","puid35":1000000,"puid36":"\/\/70.img.avito.st\/640x480\/9582784370.jpg","puid63":"30"},"alid":"7ea4295905f3b735723f09f510abf3e4"}},{"sellingSystem":"DFP","parameters":{"yandex_rtb_low_block_id":"R-A-419506-8","yandex_rtb_low_stat_id":"100003649","dfp_slot":"\/7870\/AR\/hobbi_i_otdyh\/kollektsionirovanie\/Item\/tgb1","google_adexchange_pf":"10","sizes":[[240,160],[240,133],[240,145],[240,200],[300,195],[300,250],[240,400],[300,300]],"count":1,"targetings":{"bp":"tgb1"},"isrmp":1,"page_type":"item","AudRandom":49,"keyword":"","g_metro":2207,"master_category":7,"pmin":null,"pmax":null,"cat_interests":[],"params_interests":[],"socd":[],"abp":0,"ip":"79.165.72.11","referer":"\/moskva\/kollektsionirovanie\/moneta_iz_makdonaldsa_2009709110","item_property":["77"],"slave_category":36,"g_country":621540,"g_reg":637640,"g_city":637640,"g_domofond_reg":"moskva-c3584","Usid":0,"stlogin":"0","phcodedesc":null,"Uname_Get4C":"","has_shop":0,"company_id":null,"price":"1000000-1499999","stcompany":1,"par_title":"Монета из макдоналдса","par_price":1000000,"par_picture":"\/\/70.img.avito.st\/640x480\/9582784370.jpg","alid":"7ea4295905f3b735723f09f510abf3e4"}}],"tgb2":[{"sellingSystem":"DFP","parameters":{"dfp_slot":"\/7870\/AR\/hobbi_i_otdyh\/kollektsionirovanie\/Item\/tgb2","google_adexchange_pf":"5","yandex_rtb_low_block_id":"R-A-419506-55","yandex_rtb_low_stat_id":"100003649","sizes":[[240,160],[240,133],[240,145],[240,200],[300,195],[300,250],[240,400],[300,300]],"count":1,"targetings":{"bp":"tgb2"},"isrmp":1,"page_type":"item","AudRandom":49,"keyword":"","g_metro":2207,"master_category":7,"pmin":null,"pmax":null,"cat_interests":[],"params_interests":[],"socd":[],"abp":0,"ip":"79.165.72.11","referer":"\/moskva\/kollektsionirovanie\/moneta_iz_makdonaldsa_2009709110","item_property":["77"],"slave_category":36,"g_country":621540,"g_reg":637640,"g_city":637640,"g_domofond_reg":"moskva-c3584","Usid":0,"stlogin":"0","phcodedesc":null,"Uname_Get4C":"","has_shop":0,"company_id":null,"price":"1000000-1499999","stcompany":1,"par_title":"Монета из макдоналдса","par_price":1000000,"par_picture":"\/\/70.img.avito.st\/640x480\/9582784370.jpg","alid":"8efe4f81a598e059f4d55444e51e7fff"}}],"wl":[{"sellingSystem":"AdFox","parameters":{"block_id":"8_bztzo-fvle","stat_id":"100003649","params":{"puid2":1,"puid18":"item","puid1":49,"puid7":2207,"puid5":7,"puid33":0,"puid14":"79.165.72.11","puid41":"\/moskva\/kollektsionirovanie\/moneta_iz_makdonaldsa_2009709110","puid15":"77","puid6":36,"puid4":637640,"puid3":637640,"puid26":0,"puid30":"0","puid31":0,"puid23":"1000000-1499999","puid25":1,"puid34":"Монета из макдоналдса","puid35":1000000,"puid36":"\/\/70.img.avito.st\/640x480\/9582784370.jpg","puid63":"30"},"alid":"36c802e1ac1c8ce22f726dadd3401a54"}},{"sellingSystem":"DFP","parameters":{"dfp_slot":"\/7870\/AR\/hobbi_i_otdyh\/kollektsionirovanie\/Item\/wl","google_adexchange_pf":"10","yandex_rtb_low_block_id":"R-A-419506-15","yandex_rtb_low_stat_id":"100003649","sizes":[[240,240],[240,400],[300,250],[308,400],[300,600]],"count":1,"targetings":{"bp":"wl"},"isrmp":1,"page_type":"item","AudRandom":49,"keyword":"","g_metro":2207,"master_category":7,"pmin":null,"pmax":null,"cat_interests":[],"params_interests":[],"socd":[],"abp":0,"ip":"79.165.72.11","referer":"\/moskva\/kollektsionirovanie\/moneta_iz_makdonaldsa_2009709110","item_property":["77"],"slave_category":36,"g_country":621540,"g_reg":637640,"g_city":637640,"g_domofond_reg":"moskva-c3584","Usid":0,"stlogin":"0","phcodedesc":null,"Uname_Get4C":"","has_shop":0,"company_id":null,"price":"1000000-1499999","stcompany":1,"par_title":"Монета из макдоналдса","par_price":1000000,"par_picture":"\/\/70.img.avito.st\/640x480\/9582784370.jpg","alid":"36c802e1ac1c8ce22f726dadd3401a54"}}],"ldr_low":[{"sellingSystem":"Yandex RTB","parameters":{"block_id":"R-A-419506-97","stat_id":"100003649","alid":"fd43ba60987dae7b8337f51728005824"}},{"sellingSystem":"DFP","parameters":{"yandex_rtb_low_stat_id":"100003649","dfp_slot":"\/7870\/AR\/hobbi_i_otdyh\/kollektsionirovanie\/Item\/ldr_low","google_adexchange_pf":"5","yandex_rtb_low_block_id":"R-A-419506-29","sizes":[[728,90],[1000,90],[1000,120],[970,90],[980,120],[900,90]],"count":1,"targetings":{"bp":"ldr_low"},"isrmp":1,"page_type":"item","AudRandom":49,"keyword":"","g_metro":2207,"master_category":7,"pmin":null,"pmax":null,"cat_interests":[],"params_interests":[],"socd":[],"abp":0,"ip":"79.165.72.11","referer":"\/moskva\/kollektsionirovanie\/moneta_iz_makdonaldsa_2009709110","item_property":["77"],"slave_category":36,"g_country":621540,"g_reg":637640,"g_city":637640,"g_domofond_reg":"moskva-c3584","Usid":0,"stlogin":"0","phcodedesc":null,"Uname_Get4C":"","has_shop":0,"company_id":null,"price":"1000000-1499999","stcompany":1,"par_title":"Монета из макдоналдса","par_price":1000000,"par_picture":"\/\/70.img.avito.st\/640x480\/9582784370.jpg","alid":"fd43ba60987dae7b8337f51728005824"}}],"btni":[{"sellingSystem":"DFP","parameters":{"dfp_slot":"\/7870\/AR\/hobbi_i_otdyh\/kollektsionirovanie\/Item\/btni","sizes":[[263,23],[263,25],[288,30],[287,30],[342,30],[264,30],[342,40],[309,25]],"count":1,"targetings":{"bp":"btni"},"isrmp":1,"page_type":"item","AudRandom":49,"keyword":"","g_metro":2207,"master_category":7,"pmin":null,"pmax":null,"cat_interests":[],"params_interests":[],"socd":[],"abp":0,"ip":"79.165.72.11","referer":"\/moskva\/kollektsionirovanie\/moneta_iz_makdonaldsa_2009709110","item_property":["77"],"slave_category":36,"g_country":621540,"g_reg":637640,"g_city":637640,"g_domofond_reg":"moskva-c3584","Usid":0,"stlogin":"0","phcodedesc":null,"Uname_Get4C":"","has_shop":0,"company_id":null,"price":"1000000-1499999","stcompany":1,"par_title":"Монета из макдоналдса","par_price":1000000,"par_picture":"\/\/70.img.avito.st\/640x480\/9582784370.jpg","alid":"1b309f12faf97d64254baae2743f3f6b"}}]};
 window.avito.rmp = window.avito.rmp || {};
 window.avito.rmp.dfpTargetings = {"isrmp":1,"page_type":"item","AudRandom":49,"keyword":"","g_metro":2207,"master_category":7,"pmin":null,"pmax":null,"cat_interests":[],"params_interests":[],"socd":[],"abp":0,"ip":"79.165.72.11","referer":"\/moskva\/kollektsionirovanie\/moneta_iz_makdonaldsa_2009709110","item_property":["77"],"slave_category":36,"g_country":621540,"g_reg":637640,"g_city":637640,"g_domofond_reg":"moskva-c3584","Usid":0,"stlogin":"0","phcodedesc":null,"Uname_Get4C":"","has_shop":0,"company_id":null,"price":"1000000-1499999","stcompany":1,"par_title":"Монета из макдоналдса","par_price":1000000,"par_picture":"\/\/70.img.avito.st\/640x480\/9582784370.jpg"};
 window.avito.rmp.enabledBanners = {"ldr_top":{"code":"ldr_top"},"tgb1":{"code":"tgb1"},"tgb2":{"code":"tgb2"},"wl":{"code":"wl"},"ldr_low":{"code":"ldr_low"},"btni":{"code":"btni"}};
 window.avito.rmp.enableEventSampling = false;
 window.avito.rmp.newYandexSearchBanner = null
 window.avito.abFeatures = window.avito.abFeatures || {};
    window.avito.abFeatures.lazyLoadDFPTest = true;
    window.avito_desktop = true;
 </script>
 <script src="https://static.avito.ru/s/cc/bundles/6ebac8d0c0473e482aeb.js" ></script>
   <script src="https://static.avito.ru/s/cc/chunks/d7953237868db32d9ee0.js" async></script>
<script src="https://static.avito.ru/s/cc/chunks/963553aefed75e903a64.js" async></script>
<script src="https://static.avito.ru/s/cc/chunks/6a786cec75ecda5c32f3.js" async></script>
<script src="https://static.avito.ru/s/cc/bundles/e6f95a26c8450adb5278.js" async></script>
    <script src="https://static.avito.ru/s/cc/bundles/e2f1f20445490466a64a.js" async></script>
  <script src="https://static.avito.ru/s/cc/bundles/eacafaba71bc713ea5cc.js" ></script>
 </head> <body class=" "  >
    <noscript> <iframe src=//www.googletagmanager.com/ns.html?id=GTM-KP9Q9H height="0" width="0" style="display:none;visibility:hidden"></iframe> </noscript>
                              
<div class="js-header-container header-container   header-container_no-bottom-margin    header-responsive">
    <div class='js-header' data-state='{"body_class":"","country":{"host":"www.avito.ru","country_slug":"rossiya","site_name":"Авито","currency_sign":"₽"},"currentPage":"item","headerInfo":[],"addButtonText":"Подать объявление","hideAddButton":false,"isNCEnabled":false,"isShowAvitoPro":false,"isOrdersPageEnabled":false,"luri":"moskva","menu":{"catalog":{"title":"Объявления","link":"catalog","active":true},"shops":{"title":"Магазины","link":"shops"},"business":{"title":"Бизнес","link":"business"},"support":{"title":"Помощь","absoluteLink":"support.avito.ru"}},"messenger":[],"servicesClassName":"header-services","user":{"isAuthenticated":false,"id":0,"name":"","hasShopSubscription":false,"isLegalPerson":false,"avatar":null},"userAccount":{"balance":{"bonus":"","real":"0","total":"0"},"isSeparationBalance":null},"user_location_id":637640,"responsive":true,"hierarchy":[],"inTNS8114test":true,"now":1602557569,"_dashboard":{}}'><div
 class="header-root-1FCTt header-services header-responsive-yeqX8"
 data-marker="header/navbar"><div class="header-inner-3iFNe header-clearfix-kI6fL"><ul class="header-list-IUZFq header-nav-wQVeb header-clearfix-kI6fL"><li class="header-nav-item-1OJG-"><a
 class="header-link-TLsAU header-nav-link-126h3"
 href="/moskva"
   >Объявления</a></li><li class="header-nav-item-1OJG-"><a
 class="header-link-TLsAU header-nav-link-126h3"
 href="/shops/moskva"
    >Магазины</a></li><li class="header-nav-item-1OJG-"><a
 class="header-link-TLsAU header-nav-link-126h3"
 href="/business"
   >Бизнес</a></li><li class="header-nav-item-1OJG-"><a
 class="header-link-TLsAU header-nav-link-126h3"
 href="//support.avito.ru"
  target="_blank" rel="noopener noreferrer" >Помощь</a></li></ul><div class="header-services-menu-2tz5y"><div class="header-services-menu-item-3H7kQ" data-marker="header/favorites"><a class="header-services-menu-link-fsJlE"
 href="/favorites"
 title="Избранное"
 ><span class="header-services-menu-icon-wrap-STcWG"><span class="header-services-menu-icon-PXhUE"><svg width="21" height="24" xmlns="http://www.w3.org/2000/svg"><path d="M10.918 5.085a5.256 5.256 0 0 1 7.524 0c2.077 2.114 2.077 5.541 0 7.655l-7.405 7.534a.75.75 0 0 1-1.074 0L2.558 12.74c-2.077-2.114-2.077-5.54 0-7.655a5.256 5.256 0 0 1 7.524 0c.15.152.289.312.418.479.13-.167.269-.327.418-.479z" fill="#CCC" fill-rule="nonzero"/></svg></span><i class="header-icon-count-2EGgu header-icon-count_red-3f61L header-icon-count_hidden-3av6Y"></i></span></a></div><div class="header-services-menu-item_username-32omV "><a class="header-services-menu-link-fsJlE header-services-menu-link-not-authenticated-3Uyu_"
 href="#login?authsrc=h"
 data-marker="header/login-button">Вход и регистрация</a></div></div><div class="header-button-wrapper-2UC-r"><a class="button-button-Dtqx2 button-button-origin-12oVr button-button-origin-blue-358Vt"
 href="#login?authsrc=ca&amp;next=%2Fadditem">Подать объявление</a></div></div></div></div>    <div class='js-header-navigation' data-state='{"alternativeCategoryMenu":null,"categoryMenu":[{"title":"Авто","categoryId":"Obj_Category_ROOT_TRANSPORT"},{"title":"Недвижимость","categoryId":"Obj_Category_ROOT_REAL_ESTATE"},{"title":"Работа","categoryId":"Obj_Category_ROOT_JOB"},{"title":"Услуги","categoryId":"Obj_Category_ROOT_SERVICES"}],"categoryTree":[{"id":25983,"mcId":1,"name":"Все категории","subs":[{"id":25984,"mcId":2,"name":"Транспорт","subs":[{"id":25985,"mcId":14,"name":"Автомобили","subs":[],"url":"\/moskva\/avtomobili?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":9,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":25986,"mcId":15,"name":"Мотоциклы и мототехника","subs":[],"url":"\/moskva\/mototsikly_i_mototehnika?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":14,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26025,"mcId":16,"name":"Грузовики и спецтехника","subs":[],"url":"\/moskva\/gruzoviki_i_spetstehnika?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":81,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26040,"mcId":12,"name":"Водный транспорт","subs":[],"url":"\/moskva\/vodnyy_transport?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":11,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":25999,"mcId":17,"name":"Запчасти и аксессуары","subs":[],"url":"\/moskva\/zapchasti_i_aksessuary?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":10,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"url":"\/moskva\/transport?cd=1","current":false,"currentParent":false,"opened":false,"level":1,"categoryId":1,"params":[],"count":6,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26113,"mcId":5,"name":"Недвижимость","subs":[{"id":26125,"mcId":30,"name":"Квартиры","subs":[],"url":"\/moskva\/kvartiry?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":24,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26115,"mcId":31,"name":"Комнаты","subs":[],"url":"\/moskva\/komnaty?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":23,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26116,"mcId":32,"name":"Дома, дачи, коттеджи","subs":[],"url":"\/moskva\/doma_dachi_kottedzhi?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":25,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26119,"mcId":34,"name":"Гаражи и машиноместа","subs":[],"url":"\/moskva\/garazhi_i_mashinomesta?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":85,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26118,"mcId":33,"name":"Земельные участки","subs":[],"url":"\/moskva\/zemelnye_uchastki?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":26,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26114,"mcId":35,"name":"Коммерческая недвижимость","subs":[],"url":"\/moskva\/kommercheskaya_nedvizhimost?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":42,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26120,"mcId":36,"name":"Недвижимость за рубежом","subs":[],"url":"\/moskva\/nedvizhimost_za_rubezhom?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":86,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"url":"\/moskva\/nedvizhimost?cd=1","current":false,"currentParent":false,"opened":false,"level":1,"categoryId":4,"params":[],"count":8,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26400,"mcId":10,"name":"Работа","subs":[{"id":26427,"mcId":61,"name":"Вакансии","subs":[],"url":"\/moskva\/vakansii?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":111,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26401,"mcId":62,"name":"Резюме","subs":[],"url":"\/moskva\/rezume?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":112,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"url":"\/moskva\/rabota?cd=1","current":false,"currentParent":false,"opened":false,"level":1,"categoryId":110,"params":[],"count":3,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26486,"mcId":63,"name":"Услуги","subs":[],"url":"\/moskva\/predlozheniya_uslug?cd=1","current":false,"currentParent":false,"opened":false,"level":1,"categoryId":114,"params":[],"count":24,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26127,"mcId":6,"name":"Личные вещи","subs":[{"id":26128,"mcId":37,"name":"Одежда, обувь, аксессуары","subs":[],"url":"\/moskva\/odezhda_obuv_aksessuary?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":27,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26153,"mcId":38,"name":"Детская одежда и обувь","subs":[],"url":"\/moskva\/detskaya_odezhda_i_obuv?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":29,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26173,"mcId":39,"name":"Товары для детей и игрушки","subs":[],"url":"\/moskva\/tovary_dlya_detey_i_igrushki?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":30,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26183,"mcId":40,"name":"Часы и украшения","subs":[],"url":"\/moskva\/chasy_i_ukrasheniya?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":28,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26187,"mcId":41,"name":"Красота и здоровье","subs":[],"url":"\/moskva\/krasota_i_zdorove?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":88,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"url":"\/moskva\/lichnye_veschi?cd=1","current":false,"currentParent":false,"opened":false,"level":1,"categoryId":5,"params":[],"count":6,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26047,"mcId":3,"name":"Для дома и дачи","subs":[{"id":26048,"mcId":20,"name":"Бытовая техника","subs":[],"url":"\/moskva\/bytovaya_tehnika?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":21,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26073,"mcId":21,"name":"Мебель и интерьер","subs":[],"url":"\/moskva\/mebel_i_interer?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":20,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26084,"mcId":22,"name":"Посуда и товары для кухни","subs":[],"url":"\/moskva\/posuda_i_tovary_dlya_kuhni?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":87,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26087,"mcId":18,"name":"Продукты питания","subs":[],"url":"\/moskva\/produkty_pitaniya?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":82,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26088,"mcId":23,"name":"Ремонт и строительство","subs":[],"url":"\/moskva\/remont_i_stroitelstvo?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":19,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26097,"mcId":19,"name":"Растения","subs":[],"url":"\/moskva\/rasteniya?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":106,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"url":"\/moskva\/dlya_doma_i_dachi?cd=1","current":false,"currentParent":false,"opened":false,"level":1,"categoryId":2,"params":[],"count":7,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26195,"mcId":7,"name":"Бытовая электроника","subs":[{"id":26196,"mcId":43,"name":"Аудио и видео","subs":[],"url":"\/moskva\/audio_i_video?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":32,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26216,"mcId":44,"name":"Игры, приставки и программы","subs":[],"url":"\/moskva\/igry_pristavki_i_programmy?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":97,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26221,"mcId":45,"name":"Настольные компьютеры","subs":[],"url":"\/moskva\/nastolnye_kompyutery?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":31,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26222,"mcId":46,"name":"Ноутбуки","subs":[],"url":"\/moskva\/noutbuki?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":98,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26223,"mcId":47,"name":"Оргтехника и расходники","subs":[],"url":"\/moskva\/orgtehnika_i_rashodniki?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":99,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26236,"mcId":48,"name":"Планшеты и электронные книги","subs":[],"url":"\/moskva\/planshety_i_elektronnye_knigi?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":96,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26249,"mcId":49,"name":"Телефоны","subs":[],"url":"\/moskva\/telefony?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":84,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26292,"mcId":42,"name":"Товары для компьютера","subs":[],"url":"\/moskva\/tovary_dlya_kompyutera?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":101,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26209,"mcId":50,"name":"Фототехника","subs":[],"url":"\/moskva\/fototehnika?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":105,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"url":"\/moskva\/bytovaya_elektronika?cd=1","current":false,"currentParent":false,"opened":false,"level":1,"categoryId":6,"params":[],"count":10,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26315,"mcId":8,"name":"Хобби и отдых","subs":[{"id":26316,"mcId":51,"name":"Билеты и путешествия","subs":[],"url":"\/moskva\/bilety_i_puteshestviya?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":33,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26339,"mcId":53,"name":"Велосипеды","subs":[],"url":"\/moskva\/velosipedy?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":34,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26345,"mcId":54,"name":"Книги и журналы","subs":[],"url":"\/moskva\/knigi_i_zhurnaly?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":83,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26349,"mcId":55,"name":"Коллекционирование","subs":[],"url":"\/moskva\/kollektsionirovanie?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":36,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26373,"mcId":52,"name":"Музыкальные инструменты","subs":[],"url":"\/moskva\/muzykalnye_instrumenty?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":38,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26324,"mcId":56,"name":"Охота и рыбалка","subs":[],"url":"\/moskva\/ohota_i_rybalka?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":102,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26325,"mcId":57,"name":"Спорт и отдых","subs":[],"url":"\/moskva\/sport_i_otdyh?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":39,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"url":"\/moskva\/hobbi_i_otdyh?cd=1","current":false,"currentParent":false,"opened":false,"level":1,"categoryId":7,"params":[],"count":8,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26098,"mcId":4,"name":"Животные","subs":[{"id":26099,"mcId":24,"name":"Собаки","subs":[],"url":"\/moskva\/sobaki?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":89,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26100,"mcId":25,"name":"Кошки","subs":[],"url":"\/moskva\/koshki?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":90,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26101,"mcId":26,"name":"Птицы","subs":[],"url":"\/moskva\/ptitsy?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":91,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26102,"mcId":27,"name":"Аквариум","subs":[],"url":"\/moskva\/akvarium?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":92,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26103,"mcId":28,"name":"Другие животные","subs":[],"url":"\/moskva\/drugie_zhivotnye?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":93,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26112,"mcId":29,"name":"Товары для животных","subs":[],"url":"\/moskva\/tovary_dlya_zhivotnyh?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":94,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"url":"\/moskva\/zhivotnye?cd=1","current":false,"currentParent":false,"opened":false,"level":1,"categoryId":35,"params":[],"count":7,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26382,"mcId":9,"name":"Для бизнеса","subs":[{"id":26383,"mcId":59,"name":"Готовый бизнес","subs":[],"url":"\/moskva\/gotoviy_biznes?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":116,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null},{"id":26393,"mcId":60,"name":"Оборудование для бизнеса","subs":[],"url":"\/moskva\/oborudovanie_dlya_biznesa?cd=1","current":false,"currentParent":false,"opened":false,"level":2,"categoryId":40,"params":[],"count":1,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"url":"\/moskva\/dlya_biznesa?cd=1","current":false,"currentParent":false,"opened":false,"level":1,"categoryId":8,"params":[],"count":3,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"url":"\/moskva?cd=1","current":true,"currentParent":false,"opened":false,"level":0,"categoryId":null,"params":[],"count":82,"shortListMaxLength":null,"shortListCollapsedLength":null,"longListMaxLength":null,"longListCollapsedLength":null,"iconUrl":null,"customUrl":null}],"commonCategories":{"Obj_Category_VERTICAL_AUTO":{"slug":null,"id":0},"Obj_Category_VERTICAL_REALTY":{"slug":"transport","id":1},"Obj_Category_VERTICAL_JOB":{"slug":"dlya_doma_i_dachi","id":2},"Obj_Category_VERTICAL_SERVICES":{"slug":null,"id":3},"Obj_Category_ROOT_TRANSPORT":{"slug":"transport","id":1},"Obj_Category_ROOT_REAL_ESTATE":{"slug":"nedvizhimost","id":4},"Obj_Category_ROOT_JOB":{"slug":"rabota","id":110},"Obj_Category_JOB_VACANCIES":{"slug":"vakansii","id":111},"Obj_Category_JOB_RESUME":{"slug":"rezume","id":112},"Obj_Category_ROOT_SERVICES":{"slug":"predlozheniya_uslug","id":114},"Obj_Category_TRANSPORT_CARS":{"slug":"avtomobili","id":9}},"constant":{"Obj_Category_VERTICAL_AUTO":0,"Obj_Category_VERTICAL_REALTY":1,"Obj_Category_VERTICAL_JOB":2,"Obj_Category_VERTICAL_SERVICES":3,"Obj_Category_ROOT_TRANSPORT":1,"Obj_Category_ROOT_REAL_ESTATE":4,"Obj_Category_ROOT_JOB":110,"Obj_Category_JOB_VACANCIES":111,"Obj_Category_JOB_RESUME":112,"Obj_Category_ROOT_SERVICES":114,"Obj_Category_TRANSPORT_CARS":9},"country":{"country_slug":"rossiya","site_name":"Авито","currency_sign":"₽"},"luri":"moskva","verticalId":4,"responsive":true,"orderAllCategories":[{"id":0,"values":[1,2,8]},{"id":1,"values":[4,6]},{"id":2,"values":[110,114,7]},{"id":3,"values":[5,35]}],"now":1602557569,"_dashboard":{}}'><div 
 class="header-navigation-basic-i28MZ header-container-basic header-navigation-responsive-2_5tJ"
 data-marker="header/navigation"><div class="header-navigation-basic-inner-226Ce  header-container-basic-inner"><div class="header-navigation-logo-2aaur"><span class="logo-root-fxfjv"><a class="logo-logo-2YITg" href="/" title=&quot;Авито &amp;mdash; сайт объявлений&quot;></a></span></div><div class="header-navigation-categories-87Lbp"><div><ul class="simple-with-more-rubricator-category-list-1B8Ve"><li class="simple-with-more-rubricator-category-item-1oRcq "><a class="simple-with-more-rubricator-link-27kbj simple-with-more-rubricator-category-link-3ngHO"
 href="/moskva/transport"
 data-marker="navigation/link"
 data-category-id="1"
 >Авто</a></li><li class="simple-with-more-rubricator-category-item-1oRcq "><a class="simple-with-more-rubricator-link-27kbj simple-with-more-rubricator-category-link-3ngHO"
 href="/moskva/nedvizhimost"
 data-marker="navigation/link"
 data-category-id="4"
 >Недвижимость</a></li><li class="simple-with-more-rubricator-category-item-1oRcq "><a class="simple-with-more-rubricator-link-27kbj simple-with-more-rubricator-category-link-3ngHO"
 href="/moskva/rabota"
 data-marker="navigation/link"
 data-category-id="110"
 >Работа</a></li><li class="simple-with-more-rubricator-category-item-1oRcq "><a class="simple-with-more-rubricator-link-27kbj simple-with-more-rubricator-category-link-3ngHO"
 href="/moskva/predlozheniya_uslug"
 data-marker="navigation/link"
 data-category-id="114"
 >Услуги</a></li><li class="simple-with-more-rubricator-category-item-1oRcq"><button class="simple-with-more-rubricator-link-27kbj simple-with-more-rubricator-category-link-3ngHO simple-with-more-rubricator-category-link_more-3cOco"
 data-marker="navigation/more-button"
 type="button" data-location-id="">ещё</button></li></ul><div
 class="simple-with-more-rubricator-more-popup-2fDTp"
 data-marker="navigation/more-popup"><div
 class="simple-with-more-rubricator-more-popup-arrow-13hlF"
 ></div><div><div class="simple-with-more-rubricator-header-categories-all-2Yo_9 js-header-more-content"><div class="simple-with-more-rubricator-header-categories-all__all-1ElCY"><a href="/moskva">Все категории</a></div><div
 class="simple-with-more-rubricator-header-categories-all__column-wrapper-Ognfc"
 ><div class="simple-with-more-rubricator-header-categories-all__column-3KQAH"><ul class="simple-with-more-rubricator-header-categories-all__list-3UY03"><li class=" simple-with-more-rubricator-header-categories-all__item_parent-yGrsI"><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/nedvizhimost?cd=1"
 data-category-id="26113"
 >Недвижимость</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/kvartiry?cd=1"
 data-category-id="26125"
 >Квартиры</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/komnaty?cd=1"
 data-category-id="26115"
 >Комнаты</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/doma_dachi_kottedzhi?cd=1"
 data-category-id="26116"
 >Дома, дачи, коттеджи</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/garazhi_i_mashinomesta?cd=1"
 data-category-id="26119"
 >Гаражи и машиноместа</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/zemelnye_uchastki?cd=1"
 data-category-id="26118"
 >Земельные участки</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/kommercheskaya_nedvizhimost?cd=1"
 data-category-id="26114"
 >Коммерческая недвижимость</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/nedvizhimost_za_rubezhom?cd=1"
 data-category-id="26120"
 >Недвижимость за рубежом</a></li></ul><ul class="simple-with-more-rubricator-header-categories-all__list-3UY03"><li class=" simple-with-more-rubricator-header-categories-all__item_parent-yGrsI"><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/rabota?cd=1"
 data-category-id="26400"
 >Работа</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/vakansii?cd=1"
 data-category-id="26427"
 >Вакансии</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/rezume?cd=1"
 data-category-id="26401"
 >Резюме</a></li></ul><ul class="simple-with-more-rubricator-header-categories-all__list-3UY03"><li class=" simple-with-more-rubricator-header-categories-all__item_parent-yGrsI"><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/zhivotnye?cd=1"
 data-category-id="26098"
 >Животные</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/sobaki?cd=1"
 data-category-id="26099"
 >Собаки</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/koshki?cd=1"
 data-category-id="26100"
 >Кошки</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/ptitsy?cd=1"
 data-category-id="26101"
 >Птицы</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/akvarium?cd=1"
 data-category-id="26102"
 >Аквариум</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/drugie_zhivotnye?cd=1"
 data-category-id="26103"
 >Другие животные</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/tovary_dlya_zhivotnyh?cd=1"
 data-category-id="26112"
 >Товары для животных</a></li></ul></div><div class="simple-with-more-rubricator-header-categories-all__column-3KQAH"><ul class="simple-with-more-rubricator-header-categories-all__list-3UY03"><li class=" simple-with-more-rubricator-header-categories-all__item_parent-yGrsI"><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/lichnye_veschi?cd=1"
 data-category-id="26127"
 >Личные вещи</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/odezhda_obuv_aksessuary?cd=1"
 data-category-id="26128"
 >Одежда, обувь, аксессуары</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/detskaya_odezhda_i_obuv?cd=1"
 data-category-id="26153"
 >Детская одежда и обувь</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/tovary_dlya_detey_i_igrushki?cd=1"
 data-category-id="26173"
 >Товары для детей и игрушки</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/chasy_i_ukrasheniya?cd=1"
 data-category-id="26183"
 >Часы и украшения</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/krasota_i_zdorove?cd=1"
 data-category-id="26187"
 >Красота и здоровье</a></li></ul><ul class="simple-with-more-rubricator-header-categories-all__list-3UY03"><li class=" simple-with-more-rubricator-header-categories-all__item_parent-yGrsI"><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/bytovaya_elektronika?cd=1"
 data-category-id="26195"
 >Бытовая электроника</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/audio_i_video?cd=1"
 data-category-id="26196"
 >Аудио и видео</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/igry_pristavki_i_programmy?cd=1"
 data-category-id="26216"
 >Игры, приставки и программы</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/nastolnye_kompyutery?cd=1"
 data-category-id="26221"
 >Настольные компьютеры</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/noutbuki?cd=1"
 data-category-id="26222"
 >Ноутбуки</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/orgtehnika_i_rashodniki?cd=1"
 data-category-id="26223"
 >Оргтехника и расходники</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/planshety_i_elektronnye_knigi?cd=1"
 data-category-id="26236"
 >Планшеты и электронные книги</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/telefony?cd=1"
 data-category-id="26249"
 >Телефоны</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/tovary_dlya_kompyutera?cd=1"
 data-category-id="26292"
 >Товары для компьютера</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/fototehnika?cd=1"
 data-category-id="26209"
 >Фототехника</a></li></ul></div><div class="simple-with-more-rubricator-header-categories-all__column-3KQAH"><ul class="simple-with-more-rubricator-header-categories-all__list-3UY03"><li class=" simple-with-more-rubricator-header-categories-all__item_parent-yGrsI"><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href=""
 data-category-id=""
 ></a></li></ul><ul class="simple-with-more-rubricator-header-categories-all__list-3UY03"><li class=" simple-with-more-rubricator-header-categories-all__item_parent-yGrsI"><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href=""
 data-category-id=""
 ></a></li></ul><ul class="simple-with-more-rubricator-header-categories-all__list-3UY03"><li class=" simple-with-more-rubricator-header-categories-all__item_parent-yGrsI"><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/hobbi_i_otdyh?cd=1"
 data-category-id="26315"
 >Хобби и отдых</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/bilety_i_puteshestviya?cd=1"
 data-category-id="26316"
 >Билеты и путешествия</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/velosipedy?cd=1"
 data-category-id="26339"
 >Велосипеды</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/knigi_i_zhurnaly?cd=1"
 data-category-id="26345"
 >Книги и журналы</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/kollektsionirovanie?cd=1"
 data-category-id="26349"
 >Коллекционирование</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/muzykalnye_instrumenty?cd=1"
 data-category-id="26373"
 >Музыкальные инструменты</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/ohota_i_rybalka?cd=1"
 data-category-id="26324"
 >Охота и рыбалка</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/sport_i_otdyh?cd=1"
 data-category-id="26325"
 >Спорт и отдых</a></li></ul></div><div class="simple-with-more-rubricator-header-categories-all__column-3KQAH"><ul class="simple-with-more-rubricator-header-categories-all__list-3UY03"><li class=" simple-with-more-rubricator-header-categories-all__item_parent-yGrsI"><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/dlya_doma_i_dachi?cd=1"
 data-category-id="26047"
 >Для дома и дачи</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/bytovaya_tehnika?cd=1"
 data-category-id="26048"
 >Бытовая техника</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/mebel_i_interer?cd=1"
 data-category-id="26073"
 >Мебель и интерьер</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/posuda_i_tovary_dlya_kuhni?cd=1"
 data-category-id="26084"
 >Посуда и товары для кухни</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/produkty_pitaniya?cd=1"
 data-category-id="26087"
 >Продукты питания</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/remont_i_stroitelstvo?cd=1"
 data-category-id="26088"
 >Ремонт и строительство</a></li><li class=""><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href="/moskva/rasteniya?cd=1"
 data-category-id="26097"
 >Растения</a></li></ul><ul class="simple-with-more-rubricator-header-categories-all__list-3UY03"><li class=" simple-with-more-rubricator-header-categories-all__item_parent-yGrsI"><a class="simple-with-more-rubricator-header-categories-all__link-k_Jr3 js-header-categories-all__link"
 href=""
 data-category-id=""
 ></a></li></ul></div></div></div></div></div></div></div></div></div></div>  </div>
   <div class="layout-internal layout-responsive">
     
<div class="b-search-form  b-search-form_item"> <script>
  </script> <form
 id="search_form"
 class="search-form__form js-search-form-catalog js-search-form"
 autocomplete="off"
 action="/search"
 method="post"
 data-initial-request-counter='1'
 data-hide-counter=""
 data-total-count=''
 data-marker="search-form"> <div class="search-form-main-controls js-search-form-main-controls">
  <input type="hidden" class="js-search-map" name="map" value="">
     <input type="hidden" class="js-token" name="token[2496499932144]" value="20548fab04c2aace">
  
 <div class="search-form__row search-form__row_1 clearfix">
  <div class="search-form__category"> <div class="form-select-v2">
 
 <select id="category" name="category_id"
 class="js-search-form-category "
 data-marker="search-form/category"> <option value="">Любая категория</option>
   <option value="1" class="opt-group"  >Транспорт</option>
    <option value="9" >Автомобили</option>
    <option value="14" >Мотоциклы и мототехника</option>
    <option value="81" >Грузовики и спецтехника</option>
    <option value="11" >Водный транспорт</option>
    <option value="10" >Запчасти и аксессуары</option>
    <option value="4" class="opt-group"  >Недвижимость</option>
    <option value="24" >Квартиры</option>
    <option value="23" >Комнаты</option>
    <option value="25" >Дома, дачи, коттеджи</option>
    <option value="26" >Земельные участки</option>
    <option value="85" >Гаражи и машиноместа</option>
    <option value="42" >Коммерческая недвижимость</option>
    <option value="86" >Недвижимость за рубежом</option>
    <option value="110" class="opt-group"  >Работа</option>
    <option value="111" >Вакансии</option>
    <option value="112" >Резюме</option>
    <option value="113" class="opt-group"  >Услуги</option>
    <option value="114" >Предложение услуг</option>
    <option value="5" class="opt-group"  >Личные вещи</option>
    <option value="27" >Одежда, обувь, аксессуары</option>
    <option value="29" >Детская одежда и обувь</option>
    <option value="30" >Товары для детей и игрушки</option>
    <option value="28" >Часы и украшения</option>
    <option value="88" >Красота и здоровье</option>
    <option value="2" class="opt-group"  >Для дома и дачи</option>
    <option value="21" >Бытовая техника</option>
    <option value="20" >Мебель и интерьер</option>
    <option value="87" >Посуда и товары для кухни</option>
    <option value="82" >Продукты питания</option>
    <option value="19" >Ремонт и строительство</option>
    <option value="106" >Растения</option>
    <option value="6" class="opt-group"  >Бытовая электроника</option>
    <option value="32" >Аудио и видео</option>
    <option value="97" >Игры, приставки и программы</option>
    <option value="31" >Настольные компьютеры</option>
    <option value="98" >Ноутбуки</option>
    <option value="99" >Оргтехника и расходники</option>
    <option value="96" >Планшеты и электронные книги</option>
    <option value="84" >Телефоны</option>
    <option value="101" >Товары для компьютера</option>
    <option value="105" >Фототехника</option>
    <option value="7" class="opt-group"  >Хобби и отдых</option>
    <option value="33" >Билеты и путешествия</option>
    <option value="34" >Велосипеды</option>
    <option value="83" >Книги и журналы</option>
    <option value="36"  selected>Коллекционирование</option>
    <option value="38" >Музыкальные инструменты</option>
    <option value="102" >Охота и рыбалка</option>
    <option value="39" >Спорт и отдых</option>
    <option value="35" class="opt-group"  >Животные</option>
    <option value="89" >Собаки</option>
    <option value="90" >Кошки</option>
    <option value="91" >Птицы</option>
    <option value="92" >Аквариум</option>
    <option value="93" >Другие животные</option>
    <option value="94" >Товары для животных</option>
    <option value="8" class="opt-group"  >Для бизнеса</option>
    <option value="116" >Готовый бизнес</option>
    <option value="40" >Оборудование для бизнеса</option>
   </select> </div> </div>
  <div class="search-form__submit"> <input
 type="submit"
 value="Найти"
 class="search button button-origin js-search-button"
 data-marker="search-form/submit-button"> </div>
                                                                           
 <div class="hidden js-show-elements" data-show-elements="[&quot;smallRadius&quot;,&quot;metro&quot;,&quot;districts&quot;]"></div> <div class="search-form__direction"> <div id="directions" class="form-select-v2 param  hidden" data-marker="search-form/directions"> <select
  disabled
  name="smallRadius[]" id="directions-select"
 class="directions"  data-filter="1"> <option data-prev-alias="smallRadius" value="">Радиус / Метро / Район</option> </select>
  <select multiple class="hidden-input-for-tab" id="directions-multiple"></select> </div> <div
 class="search-form__change-filters disabled js-change-filters"
  hidden  data-current-tab="smallRadius"
 data-selected-elements='[]'
 ></div> </div>
    
 <div class="search-form__radius"> <div
 class="
 form-select-v2
 js-search-form-radius
 "> <select
  id="radius"
 name="radius"
 class="js-search-form-radius-select"
 title="Радиус поиска"
 data-is-small-radius="1"
 data-marker="search-form/radius"> <option value="">Радиус / Метро / Район</option> </select> </div> <div
 class="search-form__change-radius disabled  js-change-small-radius "
  ></div> </div>
   <input
 hidden
 name="geoCoords"
 class="hidden js-geo-coords"
 data-default-coords="{&quot;latitude&quot;:55.755814,&quot;longitude&quot;:37.617635,&quot;zoom&quot;:16}"
 data-geo-coords="{&quot;latitude&quot;:55.755814,&quot;longitude&quot;:37.617635,&quot;zoom&quot;:16}" />
     <div class="search-form__location">
   <div class="form-select-v2"> <select
 id="region"
 name="location_id"
 class="js-search-form-region"
 data-marker="search-form/region">
 <option
 value="621540"
 data-parent-id=""
   >По всей России</option><option
 value="637640"
 data-parent-id="621540"
  data-metro-map="1"  selected >Москва</option><option
 value="637680"
 data-parent-id="621540"
   >Московская область</option> <option value="0">Выбрать другой...</option> </select> </div> <div
 class="search-form__change-location disabled js-change-location"
 data-location-id="637640"
 data-location-name="Москва"
 data-category-id="36"
 data-local-priority=""
 ></div> </div>
  <div class="search-form__key-words"> <div id="search_holder" class="search-form__key-words__search-holder"> <input id="search"
 type="text" name="name" value=""
 placeholder="Поиск по объявлениям"
  spellcheck="false" data-suggest="true" maxlength="100"
 data-suggest-delay=""
 data-marker="search-form/suggest"> </div> </div> </div> <div
    class="search-form__row search-form__row_2 js-pre-filters hidden"
 id="pre-filters">
  <label class="form-checkbox" data-marker="search-form/by-title"> <input type="checkbox" class="js-by-title" name="bt" > <span class="form-checkbox__label">только в названиях</span> </label> <label class="form-checkbox" data-marker="search-form/with-images"> <input type="checkbox" class="js-with-images" name="i" > <span class="form-checkbox__label">только с фото</span> </label>
    <label
 class="save-link_wrapper save-link_add "
 data-is-saved=""
 >
  <a
 href="/autosearch/add"
 class="save-link js-search-form-save-link"
 data-action="add"
 >Сохранить поиск</a>
  </label>
   </div> </div>
  </form>
 </div>
 </div>
  
 <div class="item-view-page-layout item-view-page-layout_content item-view-page-layout_responsive"> <div class="l-content clearfix">
      <div class="avito-ads-container avito-ads-container_ldr_top ad_1000x120 js-ads-loading avito-ads-placeholder avito-ads-placeholder_ldr_top"> <div id="template_ldr_top" class="avito-ads-template"> <div class="js-banner-1000x120 item-view-ads-ldr_top avito-ads-content"> </div> </div> </div>
  <div class="item-navigation ">
  
<div itemscope itemtype="http://schema.org/BreadcrumbList" class="breadcrumbs js-breadcrumbs breadcrumbs_gray">
  <span itemscope itemprop="itemListElement" itemtype="http://schema.org/ListItem"> <a
 itemprop="item"
 itemtype="http://schema.org/ListItem"
 class="js-breadcrumbs-link js-breadcrumbs-link-interaction"
 href="/moskva"
 title="Все объявления в Москве"> <span itemprop="name">Москва</span> </a> <meta itemprop="position" content="1"> </span>
  <span class="breadcrumbs-separator breadcrumbs-separator_gray">·</span>
   <span itemscope itemprop="itemListElement" itemtype="http://schema.org/ListItem"> <a
 itemprop="item"
 itemtype="http://schema.org/ListItem"
 class="js-breadcrumbs-link js-breadcrumbs-link-interaction"
 href="/moskva/hobbi_i_otdyh"
 title="Хобби и отдых в Москве"> <span itemprop="name">Хобби и отдых</span> </a> <meta itemprop="position" content="2"> </span>
  <span class="breadcrumbs-separator breadcrumbs-separator_gray">·</span>
   <span itemscope itemprop="itemListElement" itemtype="http://schema.org/ListItem"> <a
 itemprop="item"
 itemtype="http://schema.org/ListItem"
 class="js-breadcrumbs-link js-breadcrumbs-link-interaction"
 href="/moskva/kollektsionirovanie"
 title="Коллекционирование в Москве"> <span itemprop="name">Коллекционирование</span> </a> <meta itemprop="position" content="3"> </span>
  <span class="breadcrumbs-separator breadcrumbs-separator_gray">·</span>
   <span itemscope itemprop="itemListElement" itemtype="http://schema.org/ListItem"> <a
 itemprop="item"
 itemtype="http://schema.org/ListItem"
 class="js-breadcrumbs-link js-breadcrumbs-link-interaction"
 href="/moskva/kollektsionirovanie/monety-ASgBAgICAUQcmgE"
 title="Монеты"> <span itemprop="name">Монеты</span> </a> <meta itemprop="position" content="4"> </span>
  </div>
  <div class="item-navigation-next item-navigation-next_gray"> <a class="js-item-view-back-button" href="/moskva/kollektsionirovanie" rel="nofollow" data-from-block=1>В каталог</a>&#160;&#160;&#160;<a class="js-item-view-next-button" href="/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110/next?back=/moskva/kollektsionirovanie" rel="nofollow">Следующее &rarr;</a> </div> </div>

 <div
 class="item-view js-item-view item-view_position"
 itemscope itemtype="http://schema.org/Product">
   
<div class="sticky-header js-sticky-header sticky-header_responsive" data-relative-node="toggle-sticker-header"> <div class="item-view-page-layout"> <div class="sticky-header-content"> <div class="sticky-header-left">
  <div class="sticky-header-favorites">
  <a href="/favorites/add/2009709110"
 data-action
 data-options="&#x7B;&quot;isFavorite&quot;&#x3A;false,&quot;categorySlug&quot;&#x3A;&quot;kollektsionirovanie&quot;,&quot;compare&quot;&#x3A;false,&quot;searchHash&quot;&#x3A;null&#x7D;"
 class="sticky-header-favorites-link add-favorite add-favorite_small js-add-favorite-header"></a> </div>
  <div class="sticky-header-notes"> <span class="sticky-header-notes-link item-notes-icon_add item-notes-icon_add_header js-add-note-header"></span> </div>
   <div class="sticky-header-prop sticky-header-title">
 Монета из макдоналдса
 </div> <div class="sticky-header-prop sticky-header-price">
   
<div itemprop="offers" itemscope itemtype="http://schema.org/Offer" class="price-value price-value_side-header" id="price-value"> <span class="price-value-string js-price-value-string">
  <meta itemprop="availability" content="https://schema.org/LimitedAvailability" /> <meta itemprop="priceCurrency" content="RUB" />
  <span class="js-item-price" itemprop="price" content="1000000">1 000 000</span>&nbsp;<span class="price-value-prices-list-item-currency_sign"><span class="font_arial-rub">₽</span></span> 
 </span>
  </div>
 </div> </div> <div class="sticky-header-right">
   <div class="sticky-header-prop sticky-header-seller"> <span class="sticky-header-seller-text" title="Виктор">Виктор</span> </div>
   <div class="sticky-header-prop sticky-header-contacts">
  <span
 class="js-safedeal-node-header"
 data-marker="safedeal-item-header"></span>
    <div class="sticky-header-contact-bar js-sticky-header-contact-bar"></div>
  </div>
  </div> </div> </div> </div>
              <div class="item-view-header">
  </div> <div class="banners-header-after-container">
  </div>
 
 <div class="item-view-content"> <div class="item-view-content-left">
    
<div class="item-view-title-info js-item-view-title-info"> <div class="title-info title-info_mode-with-favorite"> <div class="title-info-main"> <h1 class="title-info-title">
  <span class="title-info-title-text" itemprop="name">Монета из макдоналдса</span>
 </h1> </div> <div id="toggle-sticker-header" class="js-toggle-sticker-header"></div>
  <div class="title-info-metadata">
    </div>
  <div class="title-info-actions"> <div class="title-info-actions-item">
  
<a
 data-side=""
 data-action
 data-favorite-mode="button"
 data-options="&#x7B;&quot;isFavorite&quot;&#x3A;false,&quot;categorySlug&quot;&#x3A;&quot;kollektsionirovanie&quot;,&quot;compare&quot;&#x3A;false,&quot;searchHash&quot;&#x3A;null&#x7D;"
 href="/favorites/add/2009709110"
 class="button button-origin button-origin_small add-favorite-button js-add-favorite"> <i class="add-favorite-button-icon"></i> <span class="add-favorite-button-text">Добавить в избранное</span> </a>
  
<a
 data-side=""
 href="#login?next=%2Fmoskva%2Fkollektsionirovanie%2Fmoneta_iz_makdonaldsa_2009709110%3Fopen-add-note&authsrc=n"
 class="button button-origin button-origin_small item-notes-button js-item-add-note-button"> <i class="item-notes-icon_add"></i>Добавить заметку
</a>
   <div class="title-info-metadata-item-redesign">
   10 октября в 16:03  </div>
  </div> </div>
    </div> </div>
    
 <div class="item-view-main js-item-view-main">
   <div class="item-view-gallery" data-hero="true">
                 <div class="gallery gallery_state-clicked js-gallery" > <div class="gallery-imgs-wrapper js-gallery-imgs-wrapper ">
   <div class="gallery-imgs-container js-gallery-imgs-container">
   <div class="gallery-img-wrapper js-gallery-img-wrapper"> <div class="gallery-img-frame js-gallery-img-frame"
 data-url="https://70.img.avito.st/640x480/9582784370.jpg"
 data-title="Монета из макдоналдса">
  </div> </div>
   <div class="gallery-img-wrapper js-gallery-img-wrapper"> <div class="gallery-img-frame js-gallery-img-frame"
 data-url="https://24.img.avito.st/640x480/9582785024.jpg"
 data-title="Монета из макдоналдса">
  </div> </div>
   </div>
    <div class="gallery-navigation gallery-navigation_prev js-gallery-navigation" data-dir="prev"> <span class="gallery-navigation-icon"></span> </div> <div class="gallery-navigation gallery-navigation_next js-gallery-navigation" data-dir="next"> <span class="gallery-navigation-icon"></span> </div>
  </div>
   <div class="gallery-list-wrapper "> <ul class="gallery-list js-gallery-list">
   <li class="gallery-list-item gallery-list-item_selected js-gallery-list-item" data-index="0" data-type="image"> <div class="gallery-list-item-link">
  <img src="https://70.img.avito.st/image/1/r1QfhrawA70_JG9bFMiz-KknCbc_JG-_qQ" alt="Монета из макдоналдса"/> </div> </li>
    <li class="gallery-list-item  js-gallery-list-item" data-index="1" data-type="image"> <div class="gallery-list-item-link">
  <img src="https://24.img.avito.st/image/1/3ZVIB7awcXxopR3-NknBOf6me3ZopR1-_g" alt="Монета из макдоналдса"/> </div> </li>
    </ul> </div>
   </div>
                       
<div class="gallery-extended gallery-extended_state-clicked gallery-extended_state-hide js-gallery-extended" data-shop-id="" > <div class="gallery-extended-content-wrapper js-gallery-extended-content-wrapper"> <div class="gallery-extended-content js-gallery-extended-content"> <div class="gallery-extended-container js-gallery-extended-container">
  <div class="gallery-extended-close js-gallery-extended-close"></div> <div class="gallery-extended-imgs-control">
   <div class="gallery-extended-img-nav gallery-extended-img-nav_type-prev js-gallery-extended-img-nav" data-dir="prev"> <span class="gallery-extended-img-nav-icon"></span> </div> <div class="gallery-extended-img-nav gallery-extended-img-nav_type-next js-gallery-extended-img-nav" data-dir="next"> <span class="gallery-extended-img-nav-icon"></span> </div>
   <div class="gallery-extended-imgs-wrapper js-gallery-extended-imgs-wrapper">
     <div class="gallery-extended-img-frame gallery-extended-img-frame_state-selected js-gallery-extended-img-frame"
 data-url="https://70.img.avito.st/image/1/r1QfhraxA70pMYGwTZjsDe4lA7mjJwm_"
 data-image-id="9582784370"
 data-alt-urls="[&quot;https:\/\/70.img.avito.st\/640x480\/9582784370.jpg&quot;]"
 data-title="Монета из макдоналдса"> </div>
     <div class="gallery-extended-img-frame js-gallery-extended-img-frame"
 data-url="https://24.img.avito.st/image/1/3ZVIB7axcXx-sPNxfmyezLmkcXj0pnt-"
 data-image-id="9582785024"
 data-alt-urls="[&quot;https:\/\/24.img.avito.st\/640x480\/9582785024.jpg&quot;]"
 data-title="Монета из макдоналдса"> </div>
   </div> </div> </div> </div> </div>
   <div class="gallery-extended-list-wrapper js-gallery-extended-list-wrapper"> <ul class="gallery-extended-list js-gallery-extended-list">
    <li class="gallery-extended-list-item  gallery-extended-list-item_selected js-gallery-extended-list-item" data-index="0" data-type="image">
  <span class="gallery-extended-list-item-link" title="Монета из макдоналдса &#8212; фотография №1" style="background-image: url(https://70.img.avito.st/image/1/r1QfhrawA70_JG9bFMiz-KknCbc_JG-_qQ);"></span> </li>
     <li class="gallery-extended-list-item  js-gallery-extended-list-item" data-index="1" data-type="image">
  <span class="gallery-extended-list-item-link" title="Монета из макдоналдса &#8212; фотография №2" style="background-image: url(https://24.img.avito.st/image/1/3ZVIB7awcXxopR3-NknBOf6me3ZopR1-_g);"></span> </li>
    </ul> </div>
  <div class="gallery-extended-list-navigation gallery-extended-list-navigation_state-hide js-gallery-extended-list-navigation" data-dir="up"> <span class="gallery-extended-list-navigation-icon"></span> </div> <div class="gallery-extended-list-navigation gallery-extended-list-navigation_state-hide gallery-extended-list-navigation_bottom js-gallery-extended-list-navigation_bottom" data-dir="bottom"> <span class="gallery-extended-list-navigation-icon"></span> </div>
   </div>
  </div>
                                   <div class="item-view-block item-view-map js-item-view-map" itemscope itemtype="http://schema.org/Place"> <div class="item-map js-item-map"> <div class="item-map-location">
  
<div class="item-address">
         <div itemprop="address" itemscope itemtype="http://schema.org/PostalAddress"><span class="item-address__string">
 Москва
 </span><div class="item-address-georeferences"><span class="item-address-georeferences"><span class="item-address-georeferences-item"><span class="item-address-georeferences-item-icons"><i class="item-address-georeferences-item-icons__icon"
 style="background-color: #4FB04F"></i></span><span class="item-address-georeferences-item__content">Беломорская</span></span><span class="item-address-georeferences-item"><span class="item-address-georeferences-item-icons"><i class="item-address-georeferences-item-icons__icon"
 style="background-color: #4FB04F"></i></span><span class="item-address-georeferences-item__content">Речной вокзал</span><span>,</span><span class="item-address-georeferences-item__after"> 1,2 км</span></span><span class="item-address-georeferences-item"><span class="item-address-georeferences-item-icons"><i class="item-address-georeferences-item-icons__icon"
 style="background-color: #4FB04F"></i></span><span class="item-address-georeferences-item__content">Ховрино</span><span>,</span><span class="item-address-georeferences-item__after"> 1,5 км</span></span></span></div></div>
</div>

 <div class="item-map-location__control">
  <div class="item-map-control"> <a data-text-open="Показать карту"
 data-text-close="Скрыть карту"
 class="item-map-slider-toggle js-item-map-slider-toggle ">
 Показать карту
 </a> </div>
  </div> </div>
  <div class="b-search-map item-map-wrapper js-item-map-wrapper"
 data-map-zoom="16"
 data-map-lat="55.865369"
 data-map-lon="37.475692"
 data-map-type="dynamic"
 data-item-id="2009709110"
 data-location-id="637640"
 data-category-id="36"
 data-shop-id=""> <div class="search-map" id="search-map"></div> </div>
  </div> </div>
   <div class="item-view-block">
 
<div class="item-description"> <div class="item-description-text" itemprop="description">
  <p>Отсуствует бумажная обертка</p>  </div> </div>
 </div>
          <div
 class="js-icebreakers__wrapper icebreakers__wrapper_borderBottom"
 data-icebreakers=&#x7B;&quot;id&quot;&#x3A;2009709110,&quot;contact&quot;&#x3A;&quot;&#x0421;&#x043F;&#x0440;&#x043E;&#x0441;&#x0438;&#x0442;&#x0435;&#x20;&#x0443;&#x20;&#x043F;&#x0440;&#x043E;&#x0434;&#x0430;&#x0432;&#x0446;&#x0430;&quot;,&quot;texts&quot;&#x3A;&#x5B;&#x7B;&quot;id&quot;&#x3A;10701421,&quot;messageText&quot;&#x3A;&quot;&#x0417;&#x0434;&#x0440;&#x0430;&#x0432;&#x0441;&#x0442;&#x0432;&#x0443;&#x0439;&#x0442;&#x0435;&#x21;&#x20;&#x0415;&#x0449;&#x0451;&#x20;&#x043F;&#x0440;&#x043E;&#x0434;&#x0430;&#x0451;&#x0442;&#x0435;&#x3F;&quot;,&quot;previewText&quot;&#x3A;&quot;&#x0415;&#x0449;&#x0451;&#x20;&#x043F;&#x0440;&#x043E;&#x0434;&#x0430;&#x0451;&#x0442;&#x0435;&#x3F;&quot;&#x7D;,&#x7B;&quot;id&quot;&#x3A;10701422,&quot;messageText&quot;&#x3A;&quot;&#x0417;&#x0434;&#x0440;&#x0430;&#x0432;&#x0441;&#x0442;&#x0432;&#x0443;&#x0439;&#x0442;&#x0435;&#x21;&#x20;&#x0421;&#x043A;&#x0430;&#x0436;&#x0438;&#x0442;&#x0435;,&#x20;&#x0442;&#x043E;&#x0440;&#x0433;&#x20;&#x0443;&#x043C;&#x0435;&#x0441;&#x0442;&#x0435;&#x043D;&#x3F;&quot;,&quot;previewText&quot;&#x3A;&quot;&#x0422;&#x043E;&#x0440;&#x0433;&#x20;&#x0443;&#x043C;&#x0435;&#x0441;&#x0442;&#x0435;&#x043D;&#x3F;&quot;&#x7D;,&#x7B;&quot;id&quot;&#x3A;10701423,&quot;messageText&quot;&#x3A;&quot;&#x0417;&#x0434;&#x0440;&#x0430;&#x0432;&#x0441;&#x0442;&#x0432;&#x0443;&#x0439;&#x0442;&#x0435;&#x21;&#x20;&#x0421;&#x043C;&#x043E;&#x0436;&#x0435;&#x0442;&#x0435;&#x20;&#x043E;&#x0442;&#x043F;&#x0440;&#x0430;&#x0432;&#x0438;&#x0442;&#x044C;&#x20;&#x0410;&#x0432;&#x0438;&#x0442;&#x043E;&#x20;&#x0414;&#x043E;&#x0441;&#x0442;&#x0430;&#x0432;&#x043A;&#x043E;&#x0439;&#x3F;&#x20;&#x0422;&#x0443;&#x0442;&#x20;&#x043D;&#x0430;&#x043F;&#x0438;&#x0441;&#x0430;&#x043D;&#x043E;,&#x20;&#x043A;&#x0430;&#x043A;&#x20;&#x043E;&#x043D;&#x0430;&#x20;&#x0440;&#x0430;&#x0431;&#x043E;&#x0442;&#x0430;&#x0435;&#x0442;&#x3A;&#x20;https&#x3A;&#x5C;&#x2F;&#x5C;&#x2F;support.avito.ru&#x5C;&#x2F;categories&#x5C;&#x2F;115000474347.&quot;,&quot;previewText&quot;&#x3A;&quot;&#x041E;&#x0442;&#x043F;&#x0440;&#x0430;&#x0432;&#x0438;&#x0442;&#x0435;&#x20;&#x0410;&#x0432;&#x0438;&#x0442;&#x043E;&#x20;&#x0414;&#x043E;&#x0441;&#x0442;&#x0430;&#x0432;&#x043A;&#x043E;&#x0439;&#x3F;&quot;&#x7D;,&#x7B;&quot;id&quot;&#x3A;10701424,&quot;messageText&quot;&#x3A;&quot;&#x0417;&#x0434;&#x0440;&#x0430;&#x0432;&#x0441;&#x0442;&#x0432;&#x0443;&#x0439;&#x0442;&#x0435;&#x21;&#x20;&#x041A;&#x043E;&#x0433;&#x0434;&#x0430;&#x20;&#x043C;&#x043E;&#x0436;&#x043D;&#x043E;&#x20;&#x043F;&#x043E;&#x0441;&#x043C;&#x043E;&#x0442;&#x0440;&#x0435;&#x0442;&#x044C;&#x3F;&quot;,&quot;previewText&quot;&#x3A;&quot;&#x041A;&#x043E;&#x0433;&#x0434;&#x0430;&#x20;&#x043C;&#x043E;&#x0436;&#x043D;&#x043E;&#x20;&#x043F;&#x043E;&#x0441;&#x043C;&#x043E;&#x0442;&#x0440;&#x0435;&#x0442;&#x044C;&#x3F;&quot;&#x7D;&#x5D;&#x7D;> <div class="icebreakers__wrapper">
 Спросите у продавца
 <div class="icebreakerBubble__wrapper">
  <button class="icebreaker__bubble">
 Ещё продаёте?
 </button>
  <button class="icebreaker__bubble">
 Торг уместен?
 </button>
  <button class="icebreaker__bubble">
 Отправите Авито Доставкой?
 </button>
  <button class="icebreaker__bubble">
 Когда можно посмотреть?
 </button>
  </div> </div> </div>
  </div>
     
 <div class="item-view-socials">
 <div class="item-socials"> <div class="item-socials-actions clearfix">
  <div class="item-socials-share">
  
<div class="js-social-share social-share"
 data-services="vkontakte,odnoklassniki,facebook,twitter,moimir,lj"
 data-title="Объявление на Авито - Монета из макдоналдса"
 data-description="Отсуствует бумажная обертка"
 data-url="https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110"
 data-image="https://www.avito.ru/img/share/auto/9582784370"> </div>
 </div>
    <div class="item-socials-abuse"> <button class="js-abuse-button button button-origin">Пожаловаться</button> <input class="js-token" type="hidden" name="token[2496499932144]" value="20548fab04c2aace"> <div id="abuse" data-abuse='{"itemId":2009709110,"isAuth":false}'  data-recaptcha-enabled="1"></div> </div>
   </div> </div>
 </div>
     <div class="avito-ads-item-mid">
    </div>
   <div class="item-view-similars">
 <div class="similars js-similars similars_column-4"
 data-show-more-btn="1"> <div class="similars-inner similars-inner_hidden js-similars-inner"> <div
 class="similars-list js-similars-list"
 data-serp-link=""
 data-serp-link-text=""
 data-from-page=""> </div> </div> </div>
 </div>
   </div> <div class="item-view-content-right"> <div class="item-view-info js-item-view-info js-sticky-fallback  item-view_position"> <div class="item-view-contacts js-item-view-contacts">
    
 <div class="item-view-price js-item-view-price"> <div class="item-view-price-content js-item-view-price-content">
 
<div class="item-price"> <div class="item-price-wrapper">
    
<div itemprop="offers" itemscope itemtype="http://schema.org/Offer" class="price-value price-value_side-card" id="price-value"> <span class="price-value-string js-price-value-string">
  <meta itemprop="availability" content="https://schema.org/LimitedAvailability" /> <meta itemprop="priceCurrency" content="RUB" />
  <span class="js-item-price" itemprop="price" content="1000000">1 000 000</span>&nbsp;<span class="price-value-prices-list-item-currency_sign"><span class="font_arial-rub">₽</span></span> 
 </span>
  </div>
   </div>
   </div>
  <div class="item-price-banner js-item-price-banner"> <div class="item-price-banner-holder">
    <div class="avito-ads-container avito-ads-container_btni"> <div id="template_btni" class="avito-ads-template"> <div class="item-price-banner-content js-item-price-banner-content"> </div> </div> </div>
 </div> </div>
  </div> </div>
 
 <script>
 var itemTitle = document.querySelector('.js-item-view-title-info');
 var itemPrice = document.querySelector('.js-item-view-price');
 var itemPriceContent = document.querySelector('.js-item-view-price-content');
 var itemPriceCorrectSize = itemTitle.offsetHeight;
 var itemPriceSize = itemPrice.offsetHeight;
 itemPrice.style.height = itemPriceCorrectSize  + 'px';
  var itemInfo = document.querySelector('.js-item-view-info');
 var gap = itemPriceCorrectSize - itemPriceSize;
 var stickyGap = 25;
 var itemTitleTop;
 var getScrollTop = function () {
 return (document.scrollingElement || document.documentElement).scrollTop;
 };
 var getItemTitleTop = function () {
 if (itemTitleTop) {
 return itemTitleTop;
 }
 return itemTitle.getBoundingClientRect().top + getScrollTop();
 };
 var isWide = function () {
 return document.documentElement.clientWidth > 1334;
 };
 window.addEventListener('DOMContentLoaded', function () {
 itemTitleTop = itemTitle.getBoundingClientRect().top + getScrollTop();
 });
 var recalc = function () {
 if (!isWide()) {
 return;
 }
 var transform = getScrollTop() - getItemTitleTop() + stickyGap;
 if (transform < 0) {
 transform = 0;
 } else if (transform > gap) {
 transform = gap;
 }
 itemPriceContent.style.transform = 'translateY(' + transform + 'px)';
 itemInfo.style.top = stickyGap - transform + 'px';
 };
 var resize = function () {
 if (isWide()) {
 recalc();
 } else {
 itemPriceContent.style.transform = '';
 itemInfo.style.top = '';
 }
 };
 window.addEventListener('resize', resize);
 window.addEventListener('scroll', recalc);
 resize();
  </script>
     <div class="item-view-actions "> <div class="item-actions js-item-actions">
   
 <div
 class="item-actions-line"
 data-marker="safedeal-item-card">
 <div class="js-ssr-5f851681d625f" data-props='{&quot;itemId&quot;:2009709110,&quot;isLogin&quot;:false,&quot;isUserSeller&quot;:false,&quot;isBuyerLocationAppropriate&quot;:false,&quot;isNewTrustFactorsEnabled&quot;:false}'>
            <script>
                window.avito = window.avito || {};
                window.avito.__SAFEDEAL_ITEM_STATE__ = {"onboarding":{"service":null},"delivery":{"loading":false,"available":false,"enabled":false,"reserved":false,"orders":[],"deliveryConditions":null},"deliveryCourier":{"loading":false,"available":false,"enabled":false,"reserved":false,"orders":[],"deliveryConditions":null}};
            </script>
            
        </div> </div>
      <div class="js-contact-bar contact-bar_wrapper">
 <div class="js-ssr-5f851681d71bb" data-props='&quot;{\&quot;itemId\&quot;:2009709110,\&quot;locationId\&quot;:637640,\&quot;categoryId\&quot;:36,\&quot;price\&quot;:1000000,\&quot;isVacancies\&quot;:false,\&quot;isResume\&quot;:false,\&quot;isLogin\&quot;:false,\&quot;side\&quot;:\&quot;card\&quot;,\&quot;useDeliveryPopup\&quot;:true,\&quot;hasCvPackage\&quot;:false,\&quot;seller\&quot;:{\&quot;name\&quot;:\&quot;\u0412\u0438\u043a\u0442\u043e\u0440\&quot;,\&quot;avatar\&quot;:\&quot;https:\\\/\\\/19.img.avito.st\\\/avatar\\\/social\\\/256x256\\\/3685744219.jpg\&quot;,\&quot;profileUrl\&quot;:\&quot;\\\/user\\\/052b87f0dc4df3f7a2f77d3e669553d6\\\/profile?id=2009709110&amp;src=item\&quot;},\&quot;shop\&quot;:null,\&quot;contacts\&quot;:{\&quot;list\&quot;:[{\&quot;type\&quot;:\&quot;messenger\&quot;,\&quot;value\&quot;:{\&quot;isAuthorized\&quot;:false,\&quot;isFullWidth\&quot;:true,\&quot;hasHidePhoneOnboarding\&quot;:true,\&quot;sellerIdHash\&quot;:\&quot;119541467\&quot;,\&quot;itemId\&quot;:2009709110,\&quot;itemUrl\&quot;:\&quot;\\\/moskva\\\/kollektsionirovanie\\\/moneta_iz_makdonaldsa_2009709110\&quot;,\&quot;itemCVViewed\&quot;:false,\&quot;categoryId\&quot;:36,\&quot;buttonText\&quot;:\&quot;\u041d\u0430\u043f\u0438\u0441\u0430\u0442\u044c \u0441\u043e\u043e\u0431\u0449\u0435\u043d\u0438\u0435\&quot;,\&quot;replyTime\&quot;:\&quot;\&quot;,\&quot;isMiniMessenger\&quot;:true,\&quot;logParams\&quot;:{\&quot;userId\&quot;:119541467,\&quot;wsrc\&quot;:\&quot;item\&quot;,\&quot;s\&quot;:\&quot;mi\&quot;},\&quot;experiments\&quot;:[]}}]},\&quot;hidePhone\&quot;:true,\&quot;searchHash\&quot;:\&quot;\&quot;,\&quot;publicProfileInfo\&quot;:{\&quot;tenureMedal\&quot;:\&quot;gold\&quot;,\&quot;isCompany\&quot;:false,\&quot;isChargeableCV\&quot;:false,\&quot;hideSellerName\&quot;:false,\&quot;itemCategoryIsServices\&quot;:false,\&quot;isShortRentItem\&quot;:false,\&quot;showVerifiedSellerMark\&quot;:false,\&quot;publicProfileLink\&quot;:\&quot;\\\/user\\\/052b87f0dc4df3f7a2f77d3e669553d6\\\/profile?id=2009709110&amp;src=item\&quot;,\&quot;isAbuseEnabled\&quot;:true,\&quot;phoneNote\&quot;:\&quot;\u0421\u043a\u0430\u0436\u0438\u0442\u0435 \u043f\u0440\u043e\u0434\u0430\u0432\u0446\u0443, \u0447\u0442\u043e \u043d\u0430\u0448\u043b\u0438 \u044d\u0442\u043e \u043e\u0431\u044a\u044f\u0432\u043b\u0435\u043d\u0438\u0435 \u043d\u0430 \u0410\u0432\u0438\u0442\u043e.\&quot;,\&quot;isShowScamAlert\&quot;:true,\&quot;sellerScam\&quot;:\&quot;\u043f\u0440\u043e\u0434\u0430\u0432\u0446\u0430\&quot;,\&quot;howOldInfo\&quot;:\&quot;\u041d\u0430 \u0410\u0432\u0438\u0442\u043e c \u0430\u0432\u0433\u0443\u0441\u0442\u0430 2017\&quot;,\&quot;showMedal\&quot;:false,\&quot;vin\&quot;:null,\&quot;shopLink\&quot;:\&quot;\&quot;,\&quot;isShop\&quot;:false,\&quot;showShopLink\&quot;:false,\&quot;shopInfo\&quot;:{\&quot;isShop\&quot;:false,\&quot;shopName\&quot;:\&quot;\&quot;,\&quot;shopLink\&quot;:\&quot;\&quot;},\&quot;isVerifiedSeller\&quot;:false,\&quot;contactManager\&quot;:\&quot;\&quot;,\&quot;itemSellerName\&quot;:\&quot;\u0412\u0438\u043a\u0442\u043e\u0440\&quot;,\&quot;sellerName\&quot;:\&quot;\u0427\u0430\u0441\u0442\u043d\u043e\u0435 \u043b\u0438\u0446\u043e\&quot;,\&quot;publicProfile\&quot;:{\&quot;link\&quot;:\&quot;\\\/user\\\/052b87f0dc4df3f7a2f77d3e669553d6\\\/profile?id=2009709110&amp;src=item\&quot;,\&quot;avatar\&quot;:\&quot;https:\\\/\\\/19.img.avito.st\\\/avatar\\\/social\\\/256x256\\\/3685744219.jpg\&quot;},\&quot;showClosedItemsCount\&quot;:false},\&quot;antifraudOldInfo\&quot;:false}&quot;'><div data-reactroot=""><div class="contact-bar-wrapper-2FTgk"><span class="tooltip-tooltip-box-2rApK"><span class="tooltip-target-wrapper-XcPdv"><div><div class="messenger-button-root-3gcie messenger-button-root_fullwidth-1Qoze"><div class="messenger-button-onboardingButton-3r8ky"><div class="text-text-1PdBw text-size-l-2gTpu"><span class="messenger-button-onboardingButtonText-2_Xgw">Без звонков</span></div><div class="text-text-1PdBw text-size-s-1PUdo"><span class="messenger-button-onboardingHintText-bu8lf">Пользователь предпочитает сообщения</span></div></div><a width="12" href="#login?next=%3Fwritein%3D2009709110&amp;authsrc=mi" target="_self" data-marker="messenger-button/link" class="button-button-2Fo5k button-size-l-3LVJf button-primary-1RhOG width-width-12-2VZLz"><div class="messenger-button-cardButtonText-2eVYZ">Написать сообщение<div class="text-text-1PdBw text-size-s-1PUdo"><div class="item-owner-status-root-3jRSs item-owner-status-root_noOfflineText-3jZsd"></div></div></div></a></div></div></span></span></div></div></div> </div>
   
 </div> </div>
 
 <div class="item-view-seller-info js-item-view-seller-info ">
   
<div class="seller-info  js-seller-info ">
  <div class="seller-info-prop js-seller-info-prop_seller-name seller-info-prop_layout-two-col">
 <div class="seller-info-col"> <div class="seller-info-value"> <div class="seller-info-name js-seller-info-name">
   <a href="https://www.avito.ru/user/052b87f0dc4df3f7a2f77d3e669553d6/profile?id=2009709110&src=item&page_from=from_item_card&iid=2009709110" title="Нажмите, чтобы перейти в профиль">
 Виктор
 </a>
  </div>
   </div>
  <div>Частное лицо</div>
 
 <div class="seller-info-value">
   <div>
 На Авито c августа 2017 </div>
    </div>
 </div>    <div class="seller-info-avatar ">
  <a
 class="seller-info-avatar-image  js-public-profile-link"
 href="https://www.avito.ru/user/052b87f0dc4df3f7a2f77d3e669553d6/profile?id=2009709110&src=item&page_from=from_item_card_icon&iid=2009709110"
 title="Нажмите, чтобы перейти в профиль"style="background-image: url('https://19.img.avito.st/avatar/social/256x256/3685744219.jpg')">Профиль</a>
  </div>
   </div>
   <div class="js-favorite-seller-buttons seller-info-favorite-seller-buttons" data-item-id='2009709110' data-props='{"isShop":false,"isLogin":false,"publicProfileLink":"\/user\/052b87f0dc4df3f7a2f77d3e669553d6\/profile?id=2009709110&src=item","summary":"","userKey":"052b87f0dc4df3f7a2f77d3e669553d6","isSubscribed":false,"subscriptionCounter":0}'></div>
   
 <span
 class="seller-info-timing"
 elementtiming="tns.seller-info"> <span class="seller-info-timing_content">timing</span> </span> </div>
 </div>
  <div class="item-view-search-info-redesign"> <span data-marker="item-view/item-id">№ 2009709110</span>
  ,   
<div class="title-info-metadata-item title-info-metadata-views"> <i class="title-info-icon-views"></i>804 (+14)</div>
  </div>
 
 <span
 class="item-view-timing"
 elementtiming="bx.contacts"> <span class="item-view-timing_content">timing</span> </span> </div>
  </div>
  <div class="item-view-ads">
    <div class="avito-ads-container avito-ads-container_wl"> <div id="template_wl" class="avito-ads-template"> <div class=" avito-ads-content"> </div> </div> </div>
    <div class="avito-ads-container avito-ads-container_tgb1"> <div id="template_tgb1" class="avito-ads-template"> <div class="ad_300x133 avito-ads-content"> </div> </div> </div>

 <div class="avito-ads-tgb2-sticky-container">
    <div class="avito-ads-container avito-ads-container_tgb2"> <div id="template_tgb2" class="avito-ads-template"> <div class="ad_300x133 avito-ads-content"> </div> </div> </div>
 </div> </div>
   </div> </div>
  <div class="item-view-low-ads">
    <div class="avito-ads-container avito-ads-container_ldr_low"> <div id="template_ldr_low" class="avito-ads-template"> <div class="item-view-ads-ldr_tow avito-ads-content"> </div> </div> </div>
 </div>
  </div>
  <div class="slide-alert js-slide-alert">
  </div>
  <div class="item-tooltip js-item-tooltip tooltip tooltip_bottom"> <i class="item-tooltip-arrow js-item-tooltip-arrow tooltip-arrow"></i> <div class="item-tooltip-content js-item-tooltip-content tooltip__content"></div> </div>
   <script type="text/template" id="js-cookie-support"> <div class="cookie-support-icon"></div> <div class="cookie-support-title">Произошла ошибка</div> <div class="cookie-support-body">Для продолжения работы включите поддержку cookies<br>в&nbsp;настройках вашего браузера.</div> <button type="button" class="button button-origin js-reload-page">Я включил поддержку cookies</button> </script>
 </div>
  <div
 class="js-footer-app layout-internal col-12"
 data-source-data='&#x7B;&quot;luri&quot;&#x3A;&quot;moskva&quot;,&quot;countrySlug&quot;&#x3A;&quot;rossiya&quot;,&quot;supportPrefix&quot;&#x3A;&quot;https&#x3A;&#x5C;&#x2F;&#x5C;&#x2F;support.avito.ru&quot;,&quot;siteName&quot;&#x3A;&quot;&#x0410;&#x0432;&#x0438;&#x0442;&#x043E;&quot;,&quot;city&quot;&#x3A;null,&quot;mobileVersionUrl&quot;&#x3A;&quot;m.avito.ru&#x5C;&#x2F;moskva&#x5C;&#x2F;kollektsionirovanie&#x5C;&#x2F;moneta_iz_makdonaldsa_2009709110&#x3F;nomobile&#x3D;0&quot;,&quot;isShopBackground&quot;&#x3A;null,&quot;isShopPlank&quot;&#x3A;null,&quot;isCompanyPage&quot;&#x3A;false,&quot;isTechPage&quot;&#x3A;false,&quot;isBrowserMobile&quot;&#x3A;false&#x7D;'> </div>
   </div>
  
 <div id="counters-invisible" class="counters-invisible">
     <noscript> <img src="/stat/u?1602557569" alt=""/> </noscript>
  <script>
 var ci_id = 2009709110, ci_location = 637640, ci_category = 36, ci_root_category = 7;
 </script>
  <script>
  if (window.devicePixelRatio > 1) {
 avito.tracking.trackGTMEvent('tracking', 'retina', 'item');
 }
  </script>

<script>
         if (avito.tracking && avito.tracking.initCriteo) {
 var isRealty = false;
 avito.tracking.initCriteo(isRealty ? [10019, 39534] : 39534, 0, "");
 }
      
 if (avito.abFeatures.isCriteoTestTransactionsDefaultGroup) {
 avito.tracking.trackCriteo(
 { event: "viewItem", avito: "1", item: "2009709110", user_segment: 5 }
 );
 } else if (avito.abFeatures.isCriteoTestTransactionsPushRecGroup) {
 avito.tracking.trackCriteo(
 { event: "viewItem", avito: "1", item: "2009709110", user_segment: 6 }
 );
 } else if (avito.abFeatures.isCriteoTestTransactionsPushMoreAutoGroup) {
 avito.tracking.trackCriteo(
 { event: "viewItem", avito: "1", item: "2009709110", user_segment: 7 }
 );
 } else if (avito.abFeatures.isCriteoTestTransactionsPushMLBlendGroup) {
 avito.tracking.trackCriteo(
 { event: "viewItem", avito: "1", item: "2009709110", user_segment: 8 }
 );
 } else if (avito.abFeatures.isCriteoTestTransactionsPushMoreAutoWithMLBlendGroup) {
 avito.tracking.trackCriteo(
 { event: "viewItem", avito: "1", item: "2009709110", user_segment: 9 }
 );
 } else {
 avito.tracking.trackCriteo(
 { event: "viewItem", avito: "1", item: "2009709110" }
 );
 }
  avito.tracking.trackCriteoTransaction = function(idPrefix) {
 if (!idPrefix) {
 return;
 }
 var utmFromCookie = document.cookie.match(/_utmz=[^;]*utmcsr=([^|]*)/);
 var utmFromQueryString = location.href.indexOf('utm_source=criteo') !== -1 ? 1 : 0;
 if (utmFromCookie && utmFromCookie.length) {
 utmFromCookie = utmFromCookie.pop();
 }
 const criteoEvents = [{
 event: "manualDising"
 }];
 if (
 avito.abFeatures.isCriteoTestTransactionsPushRecGroup
 || avito.abFeatures.isCriteoTestTransactionsPushMoreAutoGroup
 || avito.abFeatures.isCriteoTestTransactionsPushMLBlendGroup
 || avito.abFeatures.isCriteoTestTransactionsPushMoreAutoWithMLBlendGroup
 ) {
 criteoEvents.push({
 event: "setAccount",
 account: [39534, 28472]
 });
 }
 criteoEvents.push({
 event: "trackTransaction",
 id: [idPrefix, Math.floor(Math.random()*99999999999)].join('_'),
 deduplication: utmFromQueryString || Number(utmFromCookie === 'criteo'),  item: [{
 id: "2009709110",
  price: 1000000.00,
 quantity: 1
 }]
 });
 avito.tracking.trackCriteo(...criteoEvents);
 };
     
 var _comscore = _comscore || [];
 _comscore.push({ c1: "2", c2: "9829393", c4: document.location.href, c5: "hobbi_i_otdyh" });
 (function() {
 var s = document.createElement("script"), el = document.getElementsByTagName("script")[0]; s.async = true;
 s.src = (document.location.protocol == "https:" ? "https://sb" : "http://b") + ".scorecardresearch.com/beacon.js";
 el.parentNode.insertBefore(s, el);
 })();
  
</script>
 <noscript> <img src="https://sb.scorecardresearch.com/p?c1=2&c2=9829393&cv=2.0&c4=avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110&c5=hobbi_i_otdyh&cj=1" /> </noscript>
    
 <script>
 var img = new Image();
 img.src = '//www.tns-counter.ru/V13a***R>' + document.referrer.replace(/\*/g,'%2a') + '*avito_ru/ru/CP1251/tmsec=avito_7-36/' + Math.round(Math.random() * 1000000000);
 </script> <noscript><img src="//www.tns-counter.ru/V13a****avito_ru/ru/CP1251/tmsec=avito_7-36/" width="1" height="1" alt="" /></noscript>
 
<script type="text/javascript">
 window.avito = window.avito || {};
</script>
  <!-- Yandex.Metrika counter --> <script>
 (function(m,e,t,r,i,k,a){m[i]=m[i]||function(){(m[i].a=m[i].a||[]).push(arguments)};
 m[i].l=1*new Date();k=e.createElement(t),a=e.getElementsByTagName(t)[0],k.async=1,k.src=r,a.parentNode.insertBefore(k,a)})
 (window, document, "script", "https://mc.yandex.ru/metrika/tag.js", "ym");
 ym(34241905, "init", {
 id:34241905,
  clickmap:true,
 trackLinks:true,
 accurateTrackBounce:true,
 ecommerce:"dataLayer"
 });
 ym(34241905, 'params', { "ab-features": null });
 </script> <noscript><div><img src="https://mc.yandex.ru/watch/34241905" style="position:absolute; left:-9999px;" alt="" /></div></noscript> <!-- /Yandex.Metrika counter -->
  <script>var google_conversion_id=987009030,google_conversion_label="f8JaCLLjvAQQhqDS1gM",google_custom_params=window.google_tag_params,google_remarketing_only=!0;</script><script src="//www.googleadservices.com/pagead/conversion.js"></script> <noscript> <div style="display:inline;"> <img height="1" width="1" style="border-style:none;" alt="" src="//googleads.g.doubleclick.net/pagead/viewthroughconversion/987009030/?value=0&amp;label=f8JaCLLjvAQQhqDS1gM&amp;guid=ON&amp;script=0"/> </div> </noscript>
  </div>
  <script src="https://static.avito.ru/s/cc/bundles/4c297f13b49d60bb4be1.js" async></script>
      <img style="display: none" height="1" width="1" alt=""
 src="//wf.frontend.weborama.fr/streampixel/?wamid=2337&Wvar=%7B%22g_country%22%3A621540%2C%22g_reg%22%3A637640%2C%22master_category%22%3A7%2C%22item_property%22%3A%5B%2277%22%5D%2C%22slave_category%22%3A36%7D&d.r=1602557569"
 />
   <div id="js-auth"
 data-captcha-enabled="1"
 data-experiments-loginPopup="null"
 data-experiments-loginSuggest="null"
 ></div>
  <input type="hidden" class="js-token" name="token[2496499932144]" value="20548fab04c2aace">
  <script src="https://static.avito.ru/deps/object-assign/4.1.1/prod/web/main.js" ></script>
<script src="https://static.avito.ru/deps/react/16.13.1/prod/web/main.js" ></script>
<script src="https://static.avito.ru/deps/scheduler/0.19.1/prod/web/main.js" ></script>
<script src="https://static.avito.ru/deps/react-dom/16.13.1/prod/web/main.js" ></script>
<script src="https://static.avito.ru/deps/prop-types/15.7.2/prod/web/main.js" ></script>
<script src="https://static.avito.ru/deps/react-is/16.13.1/prod/web/main.js" ></script>
<script src="https://static.avito.ru/deps/redux/4.0.5/prod/web/main.js" ></script>
<script src="https://static.avito.ru/deps/react-redux/7.2.1/prod/web/main.js" ></script>
<script src="https://static.avito.ru/deps/react-popper/1.3.7/prod/web/main.js" ></script>
<script src="https://static.avito.ru/deps/redux-thunk/2.3.0/prod/web/main.js" ></script>
  <script src="https://static.avito.ru/s/cc/chunks/815c8579b1c4a47707ea.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/3e374094164914d0fda1.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/baacb4793d867dd94482.js" ></script>
    <script src="https://static.avito.ru/s/cc/chunks/aa40d1104fb6ed7e9e4a.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/a429261266bc8beb5b5a.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/06f52c58ae54a635b3bb.js" ></script>
     <script src="https://static.avito.ru/s/cc/chunks/15cd13f44912e3bf72ec.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/7e80d92183c9c7afa895.js" ></script>
    <script src="https://static.avito.ru/s/cc/chunks/d15f167c7e17d3b411ce.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/febf11a9e293f024b350.js" ></script>
  <script src="https://static.avito.ru/s/cc/bundles/d1fb75b06937dc42a14b.js" ></script>
  <script src="https://static.avito.ru/s/cc/chunks/a3228bc11af80e240d6a.js" defer></script>
<script src="https://static.avito.ru/s/cc/chunks/d89179589d18cfe037b6.js" defer></script>
<script src="https://static.avito.ru/s/cc/chunks/689126823527e3b5ffb3.js" defer></script>
<script src="https://static.avito.ru/s/cc/chunks/a914f63bee0ffbee4554.js" defer></script>
<script src="https://static.avito.ru/s/cc/chunks/3b5e55b81d10a63acdec.js" defer></script>
<script src="https://static.avito.ru/s/cc/bundles/17ec5bb017e9e2c432cb.js" defer></script>
   <script src="https://static.avito.ru/s/cc/chunks/bba4ee8329383fb6476a.js" defer></script>
<script src="https://static.avito.ru/s/cc/bundles/3dadd1c18c485dc8daab.js" defer></script>
  <script src="https://static.avito.ru/s/cc/bundles/041345a1ca8ae3252dd7.js" defer></script>
 <script src="https://static.avito.ru/s/cc/bundles/429cd8b1adbf24f43746.js" defer></script>
  <script src="https://static.avito.ru/s/cc/chunks/07a47296e7d5ac437116.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/966d1b7c53fda474ce6e.js" ></script>
  
 <script>
 var itemMainNode = document.querySelector('.item-view-main');
 var itemContactsNode = document.querySelector('.item-view-contacts');
  if (itemMainNode && itemContactsNode) {
 var contactsHeight = itemContactsNode.clientHeight;
 itemMainNode.style.minHeight = contactsHeight + 'px';
 }
  </script>
     <script src="https://www.avito.st/s/cc/raw/888cabd27b474cf6d57d00d7a63eaee4.js" ></script>
  <script src="https://static.avito.ru/s/cc/chunks/5b3919fef9c9a8a12d70.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/e1efa7f509bb9da490d0.js" ></script>
  <link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/d585cea3a81c421c492b.css">
<script src="https://static.avito.ru/s/cc/chunks/f6d6c97d1e9a906bfca9.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/aeba4bfb2103b511e54d.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/6ad6d5364ab222db2643.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/df157359147311c9c074.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/992ea7449fed3c88efb2.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/814cf8bfc03d7486eb53.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/af06d5b744df76bee5c5.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/558953a4250f68cb8451.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/54fada59136da7cc7059.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/f145cbbd9db127132d0a.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/de75fad963437feeb434.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/82ec0d70932d63b1034e.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/af5a10279967e247e146.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/f1e7d8cd342e4d78baf9.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/2de713d1d80849ffe6d4.js" ></script>
  <link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/f90b01c6e72afad092be.css">
<script src="https://static.avito.ru/s/cc/chunks/10fca16ec66ccf3ecaa4.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/46101108cb2989ad44c5.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/5d41f36e6ae48da46442.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/459cbd0d914841f1267b.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/9fcae49d4cd9bec39028.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/0c957753492aca4c070d.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/2c0a7943de918a873109.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/0604c1a3a3dcd0d55eea.js" ></script>
 <script src="https://static.avito.ru/s/cc/bundles/25f18f974db65ecd9ec3.js" ></script>
 <script src="https://static.avito.ru/s/cc/bundles/bf89dcf6b43d4a36f345.js" ></script>
 <link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/c0fd049c8d230d16c36a.css">
<script src="https://static.avito.ru/s/cc/chunks/1b163eda98a05e3c8f93.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/f8f7f83e7b71f4e82b7e.js" ></script>
     <script src="https://static.avito.ru/@avito/sd-item-view/2.4.0/prod/web/bundles/da3e05a67ce211513054.js" ></script>
   <link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/3227406328aaf25ed828.css">
<script src="https://static.avito.ru/s/cc/chunks/4ab8a58a28c37f8b5d53.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/3d8be3d4d34b45a03c26.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/fbd3e8de86895e0489c6.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/f3cacc12602c48165a8c.js" ></script>
  <script src="https://static.avito.ru/s/cc/bundles/b03a9af78f5264a8c45d.js" ></script>
 <link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/ca299e2b29103524d999.css">
<script src="https://static.avito.ru/s/cc/chunks/4f39de0272800aeeb8d4.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/d9416211e42c03a997aa.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/abf0112abd33092518c5.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/70722726a3e8702b1306.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/6338e821b2e7d7a82274.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/c7f7e647722735f4e23a.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/52e7500cb13b7e8f4efb.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/03c2f30f4137cc082966.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/93fd8395d430e878aaa7.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/eb8b6c9e039cc0cd5261.js" ></script>
      <link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/2397da9194a16bc42a93.css">
<script src="https://static.avito.ru/s/cc/chunks/26d02d6335a0f97988ce.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/40b928fda3d5f7967eca.js" ></script>
        <script src="https://static.avito.ru/@avito/au-discount/1.0.0/prod/web/bundles/d0b173e67e036b2c31d1.js" ></script>
   <link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/67b2e4c5bf01dc54f6c7.css">
<script src="https://static.avito.ru/s/cc/chunks/2d55ec34f2e889cc2c93.js" ></script>
<script src="https://static.avito.ru/s/cc/chunks/49d177325ecec982ffe6.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/d7e68bc427a7eb901cc1.js" ></script>
   <script src="https://static.avito.ru/@avito/bx-item-view/0.53.0/prod/web/bundles/00134f1c53234380fc91.js" ></script>
   <link rel="stylesheet" href="https://static.avito.ru/s/cc/styles/4b74785b902e5552e3ca.css">
<script src="https://static.avito.ru/s/cc/chunks/1d3712b33a4a50f3ecf9.js" ></script>
<script src="https://static.avito.ru/s/cc/bundles/d4989c0bcfa90e837fb3.js" ></script>
          <script>
    (function(apps) {
        for (var key in apps) {
            if (!apps.hasOwnProperty(key)) {
                continue;
            }

            var app = apps[key];

            if (
                avito &&
                avito.bundles &&
                avito.bundles[app.name] &&
                avito.bundles[app.name][app.version] &&
                typeof avito.bundles[app.name][app.version].render === 'function'
            ) {
                render = avito.bundles[app.name][app.version].render;

                app.instances.forEach(function(instance) {
                    var mountNode = document.querySelector('.' + instance.selector);

                    var props = {};

                    try {
                        props = JSON.parse(instance.props);
                    } catch(error) {
                        console.error('Failed to parse instance.props', error);
                    }

                    render(mountNode, props);
                });
            }
        };
    })({"bx-item-view":{"name":"@avito\/bx-item-view","version":"0.53.0","instances":[{"selector":"js-ssr-5f851681d71bb","props":["\"{\\\"itemId\\\":2009709110,\\\"locationId\\\":637640,\\\"categoryId\\\":36,\\\"price\\\":1000000,\\\"isVacancies\\\":false,\\\"isResume\\\":false,\\\"isLogin\\\":false,\\\"side\\\":\\\"card\\\",\\\"useDeliveryPopup\\\":true,\\\"hasCvPackage\\\":false,\\\"seller\\\":{\\\"name\\\":\\\"\\u0412\\u0438\\u043a\\u0442\\u043e\\u0440\\\",\\\"avatar\\\":\\\"https:\\\\\\\/\\\\\\\/19.img.avito.st\\\\\\\/avatar\\\\\\\/social\\\\\\\/256x256\\\\\\\/3685744219.jpg\\\",\\\"profileUrl\\\":\\\"\\\\\\\/user\\\\\\\/052b87f0dc4df3f7a2f77d3e669553d6\\\\\\\/profile?id=2009709110&src=item\\\"},\\\"shop\\\":null,\\\"contacts\\\":{\\\"list\\\":[{\\\"type\\\":\\\"messenger\\\",\\\"value\\\":{\\\"isAuthorized\\\":false,\\\"isFullWidth\\\":true,\\\"hasHidePhoneOnboarding\\\":true,\\\"sellerIdHash\\\":\\\"119541467\\\",\\\"itemId\\\":2009709110,\\\"itemUrl\\\":\\\"\\\\\\\/moskva\\\\\\\/kollektsionirovanie\\\\\\\/moneta_iz_makdonaldsa_2009709110\\\",\\\"itemCVViewed\\\":false,\\\"categoryId\\\":36,\\\"buttonText\\\":\\\"\\u041d\\u0430\\u043f\\u0438\\u0441\\u0430\\u0442\\u044c \\u0441\\u043e\\u043e\\u0431\\u0449\\u0435\\u043d\\u0438\\u0435\\\",\\\"replyTime\\\":\\\"\\\",\\\"isMiniMessenger\\\":true,\\\"logParams\\\":{\\\"userId\\\":119541467,\\\"wsrc\\\":\\\"item\\\",\\\"s\\\":\\\"mi\\\"},\\\"experiments\\\":[]}}]},\\\"hidePhone\\\":true,\\\"searchHash\\\":\\\"\\\",\\\"publicProfileInfo\\\":{\\\"tenureMedal\\\":\\\"gold\\\",\\\"isCompany\\\":false,\\\"isChargeableCV\\\":false,\\\"hideSellerName\\\":false,\\\"itemCategoryIsServices\\\":false,\\\"isShortRentItem\\\":false,\\\"showVerifiedSellerMark\\\":false,\\\"publicProfileLink\\\":\\\"\\\\\\\/user\\\\\\\/052b87f0dc4df3f7a2f77d3e669553d6\\\\\\\/profile?id=2009709110&src=item\\\",\\\"isAbuseEnabled\\\":true,\\\"phoneNote\\\":\\\"\\u0421\\u043a\\u0430\\u0436\\u0438\\u0442\\u0435 \\u043f\\u0440\\u043e\\u0434\\u0430\\u0432\\u0446\\u0443, \\u0447\\u0442\\u043e \\u043d\\u0430\\u0448\\u043b\\u0438 \\u044d\\u0442\\u043e \\u043e\\u0431\\u044a\\u044f\\u0432\\u043b\\u0435\\u043d\\u0438\\u0435 \\u043d\\u0430 \\u0410\\u0432\\u0438\\u0442\\u043e.\\\",\\\"isShowScamAlert\\\":true,\\\"sellerScam\\\":\\\"\\u043f\\u0440\\u043e\\u0434\\u0430\\u0432\\u0446\\u0430\\\",\\\"howOldInfo\\\":\\\"\\u041d\\u0430 \\u0410\\u0432\\u0438\\u0442\\u043e c \\u0430\\u0432\\u0433\\u0443\\u0441\\u0442\\u0430 2017\\\",\\\"showMedal\\\":false,\\\"vin\\\":null,\\\"shopLink\\\":\\\"\\\",\\\"isShop\\\":false,\\\"showShopLink\\\":false,\\\"shopInfo\\\":{\\\"isShop\\\":false,\\\"shopName\\\":\\\"\\\",\\\"shopLink\\\":\\\"\\\"},\\\"isVerifiedSeller\\\":false,\\\"contactManager\\\":\\\"\\\",\\\"itemSellerName\\\":\\\"\\u0412\\u0438\\u043a\\u0442\\u043e\\u0440\\\",\\\"sellerName\\\":\\\"\\u0427\\u0430\\u0441\\u0442\\u043d\\u043e\\u0435 \\u043b\\u0438\\u0446\\u043e\\\",\\\"publicProfile\\\":{\\\"link\\\":\\\"\\\\\\\/user\\\\\\\/052b87f0dc4df3f7a2f77d3e669553d6\\\\\\\/profile?id=2009709110&src=item\\\",\\\"avatar\\\":\\\"https:\\\\\\\/\\\\\\\/19.img.avito.st\\\\\\\/avatar\\\\\\\/social\\\\\\\/256x256\\\\\\\/3685744219.jpg\\\"},\\\"showClosedItemsCount\\\":false},\\\"antifraudOldInfo\\\":false}\""]}]},"safedeal-item-view":{"name":"@avito\/sd-item-view","version":"2.4.0","instances":[{"selector":"js-ssr-5f851681d625f","props":["{\"itemId\":2009709110,\"isLogin\":false,\"isUserSeller\":false,\"isBuyerLocationAppropriate\":false,\"isNewTrustFactorsEnabled\":false}"]}]},"item-stat-and-vas":{"name":"@avito\/item-stat-and-vas","version":"2.0.1","instances":[]},"au-discount":{"name":"@avito\/au-discount","version":"1.0.0","instances":[]},"profile-sidebar-navigation":{"name":"@avito\/profile-sidebar-navigation","version":"3.3.1","instances":[]}});
</script><img src="https://redirect.frontend.weborama.fr/rd?url=https%3A%2F%2Fwww.avito.ru%2Fadvertisement%2Fweborama.gif%3Fwebouuid%3D{WEBO_CID}" alt="" width="1" height="1"></body> </html>`
