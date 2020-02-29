(function () {
    'use strict';
    //author：广州银云信息科技有限公司
    angular.module('eolinker.filter')


        .filter('HtmlFilter', function () {
            return function (input) {
                var HtmlUtil = {
                    /*2.用浏览器内部转换器实现html解码*/
                    htmlDecode: function (text) {
                        //1.首先动态创建一个容器标签元素，如DIV

                        var temp = document.createElement("div");
                        //2.然后将要转换的字符串设置为这个元素的innerHTML(ie，火狐，google都支持)
                        temp.innerHTML = text;
                        //3.最后返回这个元素的innerText(ie支持)或者textContent(火狐，google支持)，即得到经过HTML解码的字符串了。

                        var output = temp.innerText || temp.textContent;
                        temp = null;

                        return output;

                    },
                    /*4.用正则表达式实现html解码*/
                    htmlDecodeByRegExp: function (str) {

                        var s = "";

                        if (str.length == 0) return "";


                        s = str.replace(/&lt;/g, "<");
                        s = s.replace(/&gt;/g, ">");
                        s = s.replace(/&amp;/g, "&");
                        s = s.replace(/&nbsp;/g, " ");
                        //s = s.replace(/&#39;/g, "\\\'");
                        s = s.replace(/&quot;/g, "\\\"");
                        s = s.replace(/&#65279;/g, "");
                        s = s.replace(/(\\\\ufeff)/g, "");
                        return s;

                    }
                };
                return HtmlUtil.htmlDecodeByRegExp(input);
            }
        })

})();