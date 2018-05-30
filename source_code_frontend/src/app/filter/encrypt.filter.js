(function () {
    'use strict';
    //author：广州银云信息科技有限公司
    angular.module('goku.filter')
        /*
         *计算当前时间过滤器
         */
        .filter('currentTimeFilter', [function () {
            return function () {
                var data = {
                    fun: {
                        getTime: null //获取当前时间功能函数
                    }
                }
                data.fun.getTime = function () {
                    var template = {
                        info: {
                            date: new Date(),
                            time: {
                                year: null,
                                month: null,
                                day: null,
                                hour: null,
                                minute: null,
                                second: null
                            },
                            string: null //结果存储字符串
                        }
                    }
                    template.info.time.year = template.info.date.getFullYear();
                    template.info.time.month = template.info.date.getMonth() + 1;
                    template.info.time.day = template.info.date.getDate();

                    template.info.time.hour = template.info.date.getHours();
                    template.info.time.minute = template.info.date.getMinutes(); //分
                    template.info.time.second = template.info.date.getSeconds();

                    template.info.string = template.info.time.year + "-";

                    if (template.info.time.month < 10)
                        template.info.string += "0";

                    template.info.string += template.info.time.month + "-";

                    if (template.info.time.day < 10)
                        template.info.string += "0";

                    template.info.string += template.info.time.day + " ";

                    if (template.info.time.hour < 10)
                        template.info.string += "0";

                    template.info.string += template.info.time.hour + ":";
                    if (template.info.time.minute < 10) template.info.string += '0';
                    template.info.string += template.info.time.minute + ":";
                    if (template.info.time.second < 10) template.info.string += '0';
                    template.info.string += template.info.time.second;
                    return (template.info.string);
                }
                return data.fun.getTime();
            }
        }])
        .filter('uuidFilter', [function() {
            var data = {
                fun: {
                    uuid: null //生成uuid功能函数
                }
            }
            data.fun.uuid = function() {
                var template = {
                    array: [],
                    hexSingal: "0123456789abcdef"
                }
                for (var i = 0; i < 36; i++) {
                    template.array[i] = template.hexSingal.substr(Math.floor(Math.random() * 0x10), 1);
                }
                template.array[14] = "4"; // bits 12-15 of the time_hi_and_version field to 0010
                template.array[19] = template.hexSingal.substr((template.array[19] & 0x3) | 0x8, 1); // bits 6-7 of the clock_seq_hi_and_reserved to 01
                template.array[8] = template.array[13] = template.array[18] = template.array[23] = "-";
                return template.array.join("");
            }
            return function() {
                return data.fun.uuid();
            }
        }])
        .filter('tokenFilter', ['$filter',function ($filter) {
            return function(name){
                return CryptoJS.enc.Hex.stringify(CryptoJS.SHA1((new Date()).getTime().toString()+$filter('uuidFilter')()+(name||'author:广州银云信息科技有限公司')+$filter('uuidFilter')()));
            }
        }])

})();