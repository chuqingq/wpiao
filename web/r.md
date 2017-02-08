* 短链接转长连接：即拿到biz、sn、mid、idx。curl "http://mp.weixin.qq.com/s/Nw_Jiahy6tTswuOtPv0-Zg"直接就有。
短链接：http://mp.weixin.qq.com/s/Nw_Jiahy6tTswuOtPv0-Zg
从短链接中获取到的：http://mp.weixin.qq.com/s?__biz=MzI2ODQ4NzA4MQ==&mid=2247483985&idx=1&sn=9e9ab04cca51
合法的长连接：      http://mp.weixin.qq.com/s?__biz=MzI2ODQ4NzA4MQ==&mid=2247483985&idx=1&sn=9e9ab04cca51af0c6c97b21710422fff
    var msg_link = "http://mp.weixin.qq.com/s?__biz=MzI2ODQ4NzA4MQ==\x26amp;mid=2247483985\x26amp;idx=1\x26amp;sn=9e9ab04cca51af0c6c97b21710422fff\x26amp;chksm=eaef9297dd981b8184d11f8230ba88c7c611fe6bbef3d315e568e14faa3065a4f3ffa14c8c70#rd";
                    http://mp.weixin.qq.com/s?__biz=MzI2ODQ4NzA4MQ==&mid=2247483985&idx=1&sn=9e9ab04cca51af0c6c97b21710422fff&chksm=eaef9297dd981b8184d11f8230ba88c7c611fe6bbef3d315e568e14faa3065a4f3ffa14c8c70#rd

* http://mp.weixin.qq.com/mp/newappmsgvote?action=show&__biz=MzI2ODQ4NzA4MQ==&supervoteid=481832252&uin=&key=&pass_ticket=&wxtoken=&mid=2247483985&idx=1
    var voteInfo={"title":"德孝模范评选微信投票","vote_permission":1,"expire_time":1481342400,"total_person":15981,"vote_subject":[{"type":2,"title":"家庭和睦模范投票（最多选3人）","options":[{"name":"家庭和睦模范候选人---赵鲜霞","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3IC4zcTlBeWWI4P69JBf2OtadPZOMBRRoWXfUh19QyFHj9FnHWicVzNA\/0?wx_fmt=jpeg","cnt":3884,"selected":false},{"name":"家庭和睦模范候选人---杨红霞","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3Kp2un8gwxnLN0A6PzMTk1npw4DZGp89fqVAXEOUibZmXl2OeicIPhI7A\/0?wx_fmt=jpeg","cnt":2628,"selected":false},{"name":"家庭和睦模范候选人---张家龙","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia32DgPqfK5ArTpY0iboYkG4JtajOdVR4C8157rJeJ4c5BEvWiatFG7bnnQ\/0?wx_fmt=jpeg","cnt":3049,"selected":false},{"name":"家庭和睦模范候选人---成小宽","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3s3fkmk8OYSHWB84hicD9NJ7j8UmU6AUjMicHfDBXtM4brlrUVicLooib1A\/0?wx_fmt=jpeg","cnt":3489,"selected":false},{"name":"家庭和睦模范候选人---成红兵","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia34WsyRobSFwgF2ibuzDeTGnqK3WawZoyzgNSDhqlYfbjthfKg7WCxChg\/0?wx_fmt=jpeg","cnt":5127,"selected":false}],"total_cnt":18177,"vote_id":481832253},{"type":2,"title":"孝老爱亲模范投票（最多选3人）","options":[{"name":"孝老爱亲模范候选人---马新会","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3KTOqXyfTocnLH6hxAJbWiaJRoCibcDw2ibsSRLCUleMR770kFUcnGL4fw\/0?wx_fmt=jpeg","cnt":1524,"selected":false},{"name":"孝老爱亲模范候选人---成锁红","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia39oWF9pBdUTqlIGIrzkWbcXnl8VjTicVwPnkMS3lUUqsfT8DTfTmB3ng\/0?wx_fmt=jpeg","cnt":4017,"selected":false},{"name":"孝老爱亲模范候选人---成海霞","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3qv73dSJCx6bz8aib0kvqsYTXNXfIaKQGc6m2j1y7r3aY1ZNmH7LyAwQ\/0?wx_fmt=jpeg","cnt":1533,"selected":false},{"name":"孝老爱亲模范候选人---成冲林","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3AqdzMvDkrGoiaJArUNddP2JY3qNZoicKXoI27pibgsXdLcODZWUNDXEtg\/0?wx_fmt=jpeg","cnt":5591,"selected":false},{"name":"孝老爱亲模范候选人---白兴龙","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3VqLbKZQyNiawFHpfLv4aTFicW1B0lSCsK91LRjcVVqTktMJyNKFLHZfA\/0?wx_fmt=jpeg","cnt":769,"selected":false},{"name":"孝老爱亲模范候选人---成红霞","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3uwgibYlQog0suCibJ48cV9jKloibk99rfG6lJGEyXyQDBGbdS2I3VciccQ\/0?wx_fmt=jpeg","cnt":955,"selected":false},{"name":"孝老爱亲模范候选人---马宽锋","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia382XGN3s90zgMQxxCzNlVaqYBrhAF54Tu96ArTia7MEzLI3jZvItDVuA\/0?wx_fmt=jpeg","cnt":4426,"selected":false}],"total_cnt":18815,"vote_id":481832254},{"type":2,"title":"助人为乐模范投票（最多选3人）","options":[{"name":"助人为乐模范候选人---赵芳研","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3eLuDklwXiaX4yqDU6vNf2yrJA8rI41vSOghYtlYSQJkeTeDRCumJuLg\/0?wx_fmt=jpeg","cnt":4719,"selected":false},{"name":"助人为乐模范候选人---张连会","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3uxbXfbddKF27RvBbibRp2wC2umxLohB4tOCJdtq930zSiaMINyVMpd1w\/0?wx_fmt=jpeg","cnt":3147,"selected":false},{"name":"助人为乐模范候选人---马进龙","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia36YjhCSm22ZXG5CR9et6lCa4Ko7iaIW7a16TOEhkcQm2TGKDFENBaReg\/0?wx_fmt=jpeg","cnt":4119,"selected":false},{"name":"助人为乐模范候选人---李千万","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/Tibp1QKkyT4K7Lz2VJSOjN4lmnYnrozia3jNnZcNyCWUTALlPCfweHZXzcIgFfTU00SDaSBRJotIuMkYZIO1EPeA\/0?wx_fmt=jpeg","cnt":4050,"selected":false},{"name":"助人为乐模范候选人---张永战","cnt":2173,"selected":false}],"total_cnt":18208,"vote_id":481832255}],"super_vote_id":481832252,"del_flag":0};


示例2：
http://mp.weixin.qq.com/s/WEBkpBjBdOAIXxu9fknV9w
    var msg_link = "http://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==\x26amp;mid=2650886522\x26amp;idx=1\x26amp;sn=317f363e12cd7c45e6bbc0de9916a6c6\x26amp;chksm=8b588222bc2f0b34afedf82c17944f80bcbb8ed01ded7c8ac35cb9af873d5cde32ecccc18d40#rd";
    http://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6&chksm=8b588222bc2f0b34afedf82c17944f80bcbb8ed01ded7c8ac35cb9af873d5cde32ecccc18d40#rd
http://mp.weixin.qq.com/mp/newappmsgvote?action=show&__biz=MzA5NjYwOTg0Nw==&supervoteid=684888406&uin=&key=&pass_ticket=&wxtoken=&mid=2650886522&idx=1
    var voteInfo={"title":"站前小学“书香家庭”评选活动","vote_permission":1,"expire_time":1486900800,"total_person":34451,"vote_subject":[{"type":1,"title":"参加站前小学“书香家庭”评选活动的家庭","options":[{"name":"一（1）班魏佳宝家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYRA9uNvaUhJHX5BwLJutC2vbTtOIprXAMLHyyq5yEc5GtPBtc207hww\/0?wx_fmt=jpeg","cnt":1407,"selected":false},{"name":"一（2）班吴佳骏家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txY3alszWGDDYFElkBTqKp2dWq71Vu55P3fFdJ74KXtAgiaqfNu6RVIbQw\/0?wx_fmt=jpeg","cnt":434,"selected":false},{"name":"一（3）班张自焜家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYpnUqSbgQJRb0u0iaZ4yicNH0jN0TJtnB29ZY3kCMONeS83h2vw5X2MAA\/0?wx_fmt=jpeg","cnt":1459,"selected":false},{"name":"一（4）班谢博宇家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYkCTicPXSQJvzxIkbXxlrRQsSFcT11ia6ZCuPV3Bj3I7Rdf7sOY15H12Q\/0?wx_fmt=jpeg","cnt":2701,"selected":false},{"name":"二（1）班张思华家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYyDibhqdgMba72YHHJx0OAVHqZ2ZrrzGkjzQF5XNwQIAblDzE5EacoKg\/0?wx_fmt=jpeg","cnt":426,"selected":false},{"name":"二（2）班董书彤家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txY3M61xhhlC0pAaEibnSSBKyRC3icibvx4qjEbUiaWC2PXFeNjZFLkSdAMBQ\/0?wx_fmt=jpeg","cnt":558,"selected":false},{"name":"二（3）班高天然家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txY2COBxW67jQSu3BWOxVjczUMt8iaCsjOPdP2tlNicFibSwgWNoOajqeTMg\/0?wx_fmt=jpeg","cnt":730,"selected":false},{"name":"二（4）班张靖雯家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYHwNUrnnHUlxDA8XcUkaiaGpricIEJQInmwvgDZrr6k5wp3PwmDx0fm9A\/0?wx_fmt=jpeg","cnt":3450,"selected":false},{"name":"三（1）班王紫郡家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYj9JibZZSKC1U97kmUJkat8olpynbficKRkwhvLbIWxv52ZFia5ndibmOxQ\/0?wx_fmt=jpeg","cnt":355,"selected":false},{"name":"三（2）班孙震家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYDhFerq0cNndu5Bficl7yzicdcluN8ElmT9dlnvUtukUfF8EibLwGSqntQ\/0?wx_fmt=jpeg","cnt":899,"selected":false},{"name":"三（3）班杨昊喆家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txY3Nkjl0XkQ2ibK6NuJtXIL5CaV6ciasaCgOdfte6jwQbxn0T5lft3NOsw\/0?wx_fmt=jpeg","cnt":2751,"selected":false},{"name":"三（4）班王珺锴家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYoTIlVpKsXgvoVHrXiaQkgMhwUicXujxuzswTI5scZ6H6HeHjKpEf8iapA\/0?wx_fmt=jpeg","cnt":525,"selected":false},{"name":"四（1）班李恺辛家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txY4iaGccx1NfUpiavpXDzgqajy8Nbslyk6EKZGt0y4WUia6jVkfgyXSaUwg\/0?wx_fmt=jpeg","cnt":1918,"selected":false},{"name":"四（2）班任佳怡家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYET98jSnICkj7Yaic2rib2OkZWSzpHiaiaNu5BtnSqhRXZaiaicZOgplib6Ynw\/0?wx_fmt=jpeg","cnt":1567,"selected":false},{"name":"四（3）班张添义家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYQXwzGLmHd39kklkn6DOwbcs1cPefiaFianyTb2dgzblzk4lUJ1xAvVnA\/0?wx_fmt=jpeg","cnt":4339,"selected":false},{"name":"四（4）班代舒宇家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txY3MM9MBRs2p099jZSd6FEZcP5OEDTlshZGUkY2iafc68k4RdUkEQSlBg\/0?wx_fmt=jpeg","cnt":353,"selected":false},{"name":"五（1）班鄂晓萌家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYs8FKfkuIlFbTL22X7u9ldiclHkHyQdficbs60wKictHC8lSLZmsAic00pw\/0?wx_fmt=jpeg","cnt":3747,"selected":false},{"name":"五（2）班吴昊家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYS4wwCJVcib6HUu440HWw7lbfzdzqOWUz6eXxgh7BFiaMh9cPUn9NuHibA\/0?wx_fmt=jpeg","cnt":916,"selected":false},{"name":"五（3）班金爽家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYibfITQzNobaVrHHeVceaReWdrpruqNIedQM70CfQvLyqz2zejmNf3Qw\/0?wx_fmt=jpeg","cnt":356,"selected":false},{"name":"五（4）班盛曦瑶家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYwsiabP6JG0eOWosIGdiapYfiaRznDYK42WsNia548XHCY8RI2MlI9kibGDQ\/0?wx_fmt=jpeg","cnt":2238,"selected":false},{"name":"六（1）班宋歌、宋朝家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYDMWYfBGcQiazHP0OQ4nOLu4dWeVYQMbVJM0n9tx5dNibrJZazhpfMuEQ\/0?wx_fmt=jpeg","cnt":1057,"selected":false},{"name":"六（2）班王嘉欣家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYDkQxIkv5pKKPXVicnYgQia3mb5wdDJCqedjBZNJiblIlS8gmDxZlnxOLg\/0?wx_fmt=jpeg","cnt":725,"selected":false},{"name":"六（3）班高逸可家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYN3eC4y6xIWNTLicaKs0UGfagicW8HjqKR9IdYKj46lokelELcpmEDoPA\/0?wx_fmt=jpeg","cnt":650,"selected":false},{"name":"六（4）班付航家庭","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/356Z1iboRaSyU9t6fAr2eKInSib4DA6txYIyEeeGKHwDqVBAhzp4dAM9xxepVj1X5wTqHJ1k16NdicgG8N7NrNWUQ\/0?wx_fmt=jpeg","cnt":890,"selected":false}],"total_cnt":34451,"vote_id":684888407}],"super_vote_id":684888406,"del_flag":0};

https://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6&from=singlemessage&isappinstalled=1&key=9d6f9969a2b984dc250e21860f4a6307f3fa4b2602489c90d12415a9e5654dba6301cb25d89fa7e7a97bf7a0affaaa54fb7f19522ed87aa6b96178fa43e80ba08b69e2f9872097c68e8aba4c9a9728dd&ascene=1&uin=MTMwMzUxMjg3Mw%3D%3D&devicetype=Windows+7&version=61000603&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%2BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%2BN8W
一个合法的newappmsgvote响应：
            var _idx = '1',
                _mid = '2650886522',
                _wxtoken = '4161491645',
                _reprint_ticket = '',
                _source_mid = '',
                _source_idx = '',
                _data = {
                    action:'vote',
                    __biz:'MzA5NjYwOTg0Nw==',
                    uin:'MTMwMzUxMjg3Mw==',
                    key:'9d6f9969a2b984dc250e21860f4a6307f3fa4b2602489c90d12415a9e5654dba6301cb25d89fa7e7a97bf7a0affaaa54fb7f19522ed87aa6b96178fa43e80ba08b69e2f9872097c68e8aba4c9a9728dd',
                    pass_ticket:'WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%2BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%2BN8W',
                    f:'json',
                    json:JSON.stringify({
                        "super_vote_item":values,
                        "super_vote_id":voteInfo.super_vote_id
                    })
                };
            _idx && (_data["idx"] = _idx);
            _mid && (_data["mid"] = _mid);
            _wxtoken && (_data["wxtoken"] = _wxtoken);
            _reprint_ticket && (_data["reprint_ticket"] = _reprint_ticket);
            _source_mid && (_data["source_mid"] = _source_mid);
            _source_idx && (_data["source_idx"] = _source_idx);

            Z.ajax({
                url:'/mp/newappmsgvote',
                data:_data,
                dataType:'json',
                type:'post',
                success:function(res){
                    if (res.base_resp && res.base_resp.ret == 0){
                        //人数加1，selected修改为true                           
                        for (var i = 0; i < subject.length; i++) {
                            
                            for (var j = 0; j < subject[i].options.length; j++) {                                    
                                if($('input[name=vote_'+i+']').eq(j).is(':checked')){
                                    subject[i].options[j].selected = true;
                                    subject[i].options[j].cnt++;
                                    subject[i].total_cnt ++;
                                }                                    
                            };
                        };
                        voteInfo.has_vote = 1;
                        has_click_vote = false;
                        afterVote();
                    }else if(res.base_resp && res.base_resp.ret == -7){
                        has_click_vote = false;
                        alert('关注公众号后才可以投票');
                    }else if(res.base_resp && res.base_resp.ret == -6){
                        has_click_vote = false;
                        alert('投票过于频繁，请稍后重试！');
                    }else{
                        has_click_vote = false;
                        alert('投票失败，请稍后重试!');                        
                    }
                },
                error:function(){
                    has_click_vote = false;
                    alert('投票失败，请稍后重试!');
                }
            });

POST /mp/newappmsgvote HTTP/1.1
Host: mp.weixin.qq.com
Connection: keep-alive
Content-Length: 529
Pragma: no-cache
Cache-Control: no-cache
Accept: application/json
Origin: https://mp.weixin.qq.com
X-Requested-With: XMLHttpRequest
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36
Content-Type: application/x-www-form-urlencoded
Referer: https://mp.weixin.qq.com/mp/newappmsgvote?action=show&__biz=MzA5NjYwOTg0Nw==&supervoteid=684888406&uin=MTMwMzUxMjg3Mw%3D%3D&key=9d6f9969a2b984dc250e21860f4a6307f3fa4b2602489c90d12415a9e5654dba6301cb25d89fa7e7a97bf7a0affaaa54fb7f19522ed87aa6b96178fa43e80ba08b69e2f9872097c68e8aba4c9a9728dd&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%252BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%252BN8W&wxtoken=4161491645&mid=2650886522&idx=1
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6
Cookie: pgv_pvid=6036349385; ptui_loginuin=chuqingq@qq.com; ptcz=5b73a14956cd4baa8c163612e38b7624fb4964a3180d3deb48a7b37ed82c04d9; pt2gguin=o0289123732; pgv_pvi=724569088; wxtokenkey=2b79df2e602db7f424503674dc4110099d22d5643bba51d85420e6815345083b; wxticket=245928144; wxticketkey=cb1bb0a2c1bacb81ab807c93fefe13af9d22d5643bba51d85420e6815345083b; wap_sid=CKmOyO0EEkBDVjFhRVlrekNnUDVqVElBQU9LeDVhT3E4OE9hNUgtcEZjbG4yTnFxTWhqSk1nZFNNeDZONzBQRDlVeXhVelRQGAQg/BEot4jKxAswxIHqxAU=; wap_sid2=CKmOyO0EElxaMi12NElCSmZ5Yk83NGc3RzJlZVB4RHBJbWNqUVNTdXk4YkRDUXFyZldzMkJBbUtQZXN6VlpXbDJPaGs2Qkg0QWttckJUYzJaNHNTM2gwZGczeUM4M29EQUFBfg==

form data：
action:vote
__biz:MzA5NjYwOTg0Nw==
uin:MTMwMzUxMjg3Mw==
key:9d6f9969a2b984dc250e21860f4a6307f3fa4b2602489c90d12415a9e5654dba6301cb25d89fa7e7a97bf7a0affaaa54fb7f19522ed87aa6b96178fa43e80ba08b69e2f9872097c68e8aba4c9a9728dd
pass_ticket:WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%2BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%2BN8W
f:json
json:{"super_vote_item":[{"vote_id":684888407,"item_idx_list":{"item_idx":["23"]}}],"super_vote_id":684888406}
idx:1
mid:2650886522
wxtoken:4161491645


decodeURIComponent("%7B%22super_vote_item%22%3A%5B%7B%22vote_id%22%3A684888407%2C%22item_idx_list%22%3A%7B%22item_idx%22%3A%5B%220%22%5D%7D%7D%5D%2C%22super_vote_id%22%3A684888406%7D")
"{"super_vote_item":[{"vote_id":684888407,"item_idx_list":{"item_idx":["0"]}}],"super_vote_id":684888406}"

curl（bash）请求：
curl 'https://mp.weixin.qq.com/mp/newappmsgvote' -H 'Pragma: no-cache' -H 'Origin: https://mp.weixin.qq.com' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: zh-CN,zh;q=0.8,en;q=0.6' -H 'User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: application/json' -H 'Cache-Control: no-cache' -H 'X-Requested-With: XMLHttpRequest' -H 'Cookie: pgv_pvid=6036349385; ptui_loginuin=chuqingq@qq.com; ptcz=5b73a14956cd4baa8c163612e38b7624fb4964a3180d3deb48a7b37ed82c04d9; pt2gguin=o0289123732; pgv_pvi=724569088; wxtokenkey=2b79df2e602db7f424503674dc4110099d22d5643bba51d85420e6815345083b; wxticket=245928144; wxticketkey=cb1bb0a2c1bacb81ab807c93fefe13af9d22d5643bba51d85420e6815345083b; wap_sid=CKmOyO0EEkBDVjFhRVlrekNnUDVqVElBQU9LeDVhT3E4OE9hNUgtcEZjbG4yTnFxTWhqSk1nZFNNeDZONzBQRDlVeXhVelRQGAQg/BEot4jKxAswxIHqxAU=; wap_sid2=CKmOyO0EElxaMi12NElCSmZ5Yk83NGc3RzJlZVB4RHBJbWNqUVNTdXk4YkRDUXFyZldzMkJBbUtQZXN6VlpXbDJPaGs2Qkg0QWttckJUYzJaNHNTM2gwZGczeUM4M29EQUFBfg==' -H 'Connection: keep-alive' -H 'Referer: https://mp.weixin.qq.com/mp/newappmsgvote?action=show&__biz=MzA5NjYwOTg0Nw==&supervoteid=684888406&uin=MTMwMzUxMjg3Mw%3D%3D&key=9d6f9969a2b984dc250e21860f4a6307f3fa4b2602489c90d12415a9e5654dba6301cb25d89fa7e7a97bf7a0affaaa54fb7f19522ed87aa6b96178fa43e80ba08b69e2f9872097c68e8aba4c9a9728dd&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%252BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%252BN8W&wxtoken=4161491645&mid=2650886522&idx=1' --data 'action=vote&__biz=MzA5NjYwOTg0Nw%3D%3D&uin=MTMwMzUxMjg3Mw%3D%3D&key=9d6f9969a2b984dc250e21860f4a6307f3fa4b2602489c90d12415a9e5654dba6301cb25d89fa7e7a97bf7a0affaaa54fb7f19522ed87aa6b96178fa43e80ba08b69e2f9872097c68e8aba4c9a9728dd&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%252BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%252BN8W&f=json&json=%7B%22super_vote_item%22%3A%5B%7B%22vote_id%22%3A684888407%2C%22item_idx_list%22%3A%7B%22item_idx%22%3A%5B%2223%22%5D%7D%7D%5D%2C%22super_vote_id%22%3A684888406%7D&idx=1&mid=2650886522&wxtoken=4161491645' --compressed
curl 'https://mp.weixin.qq.com/mp/newappmsgvote' -H 'Pragma: no-cache' -H 'Origin: https://mp.weixin.qq.com' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: zh-CN,zh;q=0.8,en;q=0.6' -H 'User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: application/json' -H 'Cache-Control: no-cache' -H 'X-Requested-With: XMLHttpRequest' -H 'Cookie: pgv_pvid=6036349385; ptui_loginuin=chuqingq@qq.com; ptcz=5b73a14956cd4baa8c163612e38b7624fb4964a3180d3deb48a7b37ed82c04d9; pt2gguin=o0289123732; pgv_pvi=724569088; wxtokenkey=2b79df2e602db7f424503674dc4110099d22d5643bba51d85420e6815345083b; wxticket=245928144; wxticketkey=cb1bb0a2c1bacb81ab807c93fefe13af9d22d5643bba51d85420e6815345083b; wap_sid=CKmOyO0EEkBDVjFhRVlrekNnUDVqVElBQU9LeDVhT3E4OE9hNUgtcEZjbG4yTnFxTWhqSk1nZFNNeDZONzBQRDlVeXhVelRQGAQg/BEot4jKxAswxIHqxAU=; wap_sid2=CKmOyO0EElxaMi12NElCSmZ5Yk83NGc3RzJlZVB4RHBJbWNqUVNTdXk4YkRDUXFyZldzMkJBbUtQZXN6VlpXbDJPaGs2Qkg0QWttckJUYzJaNHNTM2gwZGczeUM4M29EQUFBfg==' -H 'Connection: keep-alive' -H 'Referer: https://mp.weixin.qq.com/mp/newappmsgvote?action=show&__biz=MzA5NjYwOTg0Nw==&supervoteid=684888406&uin=MTMwMzUxMjg3Mw%3D%3D&key=9d6f9969a2b984dc250e21860f4a6307f3fa4b2602489c90d12415a9e5654dba6301cb25d89fa7e7a97bf7a0affaaa54fb7f19522ed87aa6b96178fa43e80ba08b69e2f9872097c68e8aba4c9a9728dd&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%252BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%252BN8W&wxtoken=4161491645&mid=2650886522&idx=1' --data 'action=vote&__biz=MzA5NjYwOTg0Nw%3D%3D&uin=MTMwMzUxMjg3Mw%3D%3D&key=9d6f9969a2b984dc250e21860f4a6307f3fa4b2602489c90d12415a9e5654dba6301cb25d89fa7e7a97bf7a0affaaa54fb7f19522ed87aa6b96178fa43e80ba08b69e2f9872097c68e8aba4c9a9728dd&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%252BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%252BN8W&f=json&json=%7B%22super_vote_item%22%3A%5B%7B%22vote_id%22%3A684888407%2C%22item_idx_list%22%3A%7B%22item_idx%22%3A%5B%2223%22%5D%7D%7D%5D%2C%22super_vote_id%22%3A684888406%7D&idx=1&mid=2650886522' --compressed
把wxtoken去掉，仍然是相同的响应，因此猜测不需要wxtoken


响应：
{"base_resp":{"ret":-6,"errmsg":"freq"}}
