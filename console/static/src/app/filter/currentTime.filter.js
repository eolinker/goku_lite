(function () {
    'use strict';
    //author：广州银云信息科技有限公司
    angular.module('eolinker.filter')
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

        /*
         *计算当前时间过滤器
         */
        .filter('currentTimeFilter', [function () {
            return function (nowTime, options) {
                var fun = {};
                options = options || {};
                fun.getTime = function () {
                    var template = {
                        info: {
                            date: nowTime ? new Date(nowTime) : new Date(),
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
                    if (options.min == 'day') return (template.info.string);
                    if (template.info.time.hour < 10)
                        template.info.string += "0";

                    template.info.string += template.info.time.hour + ":";
                    if (template.info.time.minute < 10) template.info.string += '0';
                    template.info.string += template.info.time.minute + ":";
                    if (template.info.time.second < 10) template.info.string += '0';
                    template.info.string += template.info.time.second;
                    return (template.info.string);
                }
                return fun.getTime();
            }
        }])
        .filter('additionTimeFilter', [function (nowTimeQuery, time) {
            return function (nowTimeQuery, time) {
                var data = {
                    nowTime: null,
                    afterTime: null
                };
                if (nowTimeQuery) {
                    data.nowTime = new Date(nowTimeQuery[0] + "-" + nowTimeQuery[1] + "-" + nowTimeQuery[2]);
                    if (time < 12) {
                        data.nowTime.setDate(data.nowTime.getDate() + 30 * time);
                    } else {
                        data.nowTime.setDate(data.nowTime.getDate() + time / 12 * 365);
                    }
                    data.afterTime = data.nowTime.getFullYear() + "年" + (data.nowTime.getMonth() + 1) + "月" + data.nowTime.getDate() + "日";
                    return data.afterTime;
                }
            }
        }])
        .filter('durationTimeFilter', [function (startDate, endDate) {
            return function (startDate, endDate) {
                startDate = new Date(startDate);
                endDate = new Date(endDate);
                var diff = endDate.getTime() - startDate.getTime(); //时间差的毫秒数 
                //计算出相差天数 
                var days = Math.floor(diff / (24 * 3600 * 1000));
                //计算出小时数 
                var leave1 = diff % (24 * 3600 * 1000); //计算天数后剩余的毫秒数 
                var hours = Math.floor(leave1 / (3600 * 1000));
                //计算相差分钟数 
                var leave2 = leave1 % (3600 * 1000); //计算小时数后剩余的毫秒数 
                var minutes = Math.floor(leave2 / (60 * 1000));
                var returnStr = "";
                returnStr = minutes + "分" + returnStr;
                returnStr = hours + "小时" + returnStr;
                if (days > 0) {
                    returnStr = days + "天" + returnStr;
                }
                return returnStr;
            }
        }])

})();