(function() {
    'use strict';
    //author：广州银云信息科技有限公司
    angular.module('goku.filter')
    /*
    *计算当前时间过滤器
    */
    .filter('currentTimeFilter', [function() {
        return function() {
            var data={
                fun:{
                    getTime:null//获取当前时间功能函数
                }
            }
            data.fun.getTime = function() {
                var template={
                    info:{
                        date:new Date(),
                        time:{
                            year:null,
                            month:null,
                            day:null,
                            hour:null,
                            minute:null,
                            second:null
                        },
                        string:null//结果存储字符串
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

})();
