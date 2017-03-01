var page = require('webpage').create();

// // var url = 'https://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6&from=singlemessage&isappinstalled=1&key=d78209d75f0b7aa47345548e849491267f2f0122b63473851a7c4e2ef521f69c2349365376c66258988a4c10e7b398a11b7707ca56c4da9a0a6f9c3a5019d52da4e287d08b62767062a728687c0378ea&ascene=1&uin=MTMwMzUxMjg3Mw%3D%3D&devicetype=Windows+7&version=61000603&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%2BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%2BN8W';
var url = 'http://www.baidu.com';
var file = '1.png';

console.log('start...');

page.onLoadFinished = function() {
    console.log("page.onLoadFinished");

    // page.render(file);
    page.close();
    phantom.exit();
};

console.log('page.open');
page.open(url, function(status) {
    console.log('status: ' + status);
});   

