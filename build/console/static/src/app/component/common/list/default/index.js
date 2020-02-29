(function () {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 默认列表组件
     * @extend {object} authorityObject 权限类{operate}
     * @extend {object} mainObject 主类{setting:{colspan,warning},item:{default,fun}}
     * @extend {object} showObject 列表ITEM显示
     * @extend {object} otherObject 批量操作等额外对象
     * @extend {Object} pageObject 页码
     */
    angular.module('eolinker')
        .component('listDefaultCommonComponent', {
            templateUrl: 'app/component/common/list/default/index.html',
            controller: indexController,
            bindings: {
                authorityObject: '<',
                mainObject: '<',
                showObject: '<',
                otherObject: '=',
                list: '<',
                pageObject: '<'
            }
        })

        indexController.$inject = [];

        function indexController() {
            var vm = this;
            vm.data = {
                selectAllMore: false,
                listOrderBy: {}
            }
            vm.fun = {};
            var fun = {},
                privateFun = {};
            fun.generateHtml = function (type, array) {
                var template = {
                    html: '',
                }
                switch (type) {
                    case 'select': {
                        template.html = `<td class="${!vm.mainObject.setting.page ? 'select_checkbox w_30' : 'select_checkbox w_50'}"`;
                        if (vm.mainObject.setting.batch) {
                            template.html += ' ng-show="$ctrl.otherObject.batch.isOperating==true"';
                        }
                        template.html += '>' + '<button  type="button" class="eo-checkbox iconfont" ng-class="{\'icon-duihao\':$ctrl.otherObject.batch.indexAddress[item.' + vm.mainObject.item.primaryKey + ']}"  ' + (vm.mainObject.item.batchItemExpression || '') + ' >{{$ctrl.otherObject.batch.indexAddress[item.' + vm.mainObject.item.primaryKey + ']?\'\':\'&nbsp;\'}}</button></td>'
                        break;
                    }
                }
                return template.html;
            }
            vm.fun.showMore = function (arg) {
                var template = {};
                try {
                    template.point = arg.$event.target.classList[0];
                } catch (e) {
                    template.point = 'default';
                }
                switch (template.point) {
                    case 'more-btn':
                    case 'more-btn-icon': {
                        arg.$event.stopPropagation();
                        arg.item.clickMore = !arg.item.clickMore;
                        break;
                    }
                }
            }
            vm.fun.autoSortFun = (inputArg) => {
                if (inputArg.listItem.sort !== true) return;
                if (vm.mainObject.setting.batch && ((vm.otherObject.batch.isOperating && vm.mainObject.setting.batchInitStatus !== 'open') || vm.otherObject.batchMenu && vm.otherObject.batchMenu.isOperating)) return;
                let listOrderBy = {
                    orderBy: inputArg.item.sortOrderByVal
                }
                if (vm.data.listOrderBy.orderBy === inputArg.item.sortOrderByVal) {
                    listOrderBy.asc = inputArg.item.asc == 0 ? 1 : 0;
                } else {
                    listOrderBy.asc = inputArg.item.asc;
                }
                let isStop = vm.mainObject.baseFun.autoSortFun({
                    listOrderBy: listOrderBy
                });
                if (!isStop && vm.mainObject.setting.sortStorageKey) {
                    if (vm.data.listOrderBy.orderBy === inputArg.item.sortOrderByVal) {
                        inputArg.item.asc = listOrderBy.asc;
                    }
                    vm.data.listOrderBy = listOrderBy;
                    window.localStorage.setItem(vm.mainObject.setting.sortStorageKey, angular.toJson(listOrderBy));
                }
            }
            /**
             * @desc 拖动鼠标放开操作
             */
            privateFun.dragMouseup = (inputMark, inputWidth) => {
                vm.mainObject.setting.dragCacheObj[inputMark] = inputWidth;
                window.localStorage.setItem(vm.mainObject.setting.dragCacheVar, JSON.stringify(vm.mainObject.setting.dragCacheObj));
            }
            vm.fun.range = function (inputLength, inputObject) {
                inputLength = inputLength || 1;
                if (!vm.list[inputObject.$index + 1] || ((vm.list[inputObject.$index + 1].listDepth || 0) <= (inputObject.item.listDepth || 0))) inputLength--;
                return new Array(inputLength);
            };
            fun.getTargetEvent = function ($event, inputPointAttr) {
                var itemIndex = $event.getAttribute(inputPointAttr || 'eo-attr-index');
                if (itemIndex) {
                    return $event;
                } else {
                    return fun.getTargetEvent($event.parentNode, inputPointAttr);
                }
            }
            fun.getTargetIndex = function ($event, inputPointAttr) {
                var itemIndex = $event.getAttribute(inputPointAttr || 'eo-attr-index');
                if (itemIndex) {
                    return itemIndex;
                } else {
                    return fun.getTargetIndex($event.parentNode, inputPointAttr);
                }
            }
            fun.operateLevel = function (inputDepth, $event, inputIndex) {
                var tmp = {
                    operateName: angular.element($event).hasClass('ng-hide') ? 'removeClass' : 'addClass'
                },
                    tmpParentIsShrinkIndex = inputIndex,
                    itemIndex = inputIndex;
                while ($event && inputDepth < $event.getAttribute('eo-attr-depth')) {
                    switch (tmp.operateName) {
                        case 'addClass': {
                            vm.list[itemIndex].isHide = true;
                            break;
                        }
                        case 'removeClass': {
                            var tmpParentShrinkObject = vm.list[tmpParentIsShrinkIndex];
                            if (vm.list[itemIndex].isShrink && vm.list[itemIndex].listDepth <= tmpParentShrinkObject.listDepth) {
                                vm.list[itemIndex].isHide = false;
                                tmpParentIsShrinkIndex = itemIndex;
                            } else if (vm.list[itemIndex].listDepth <= tmpParentShrinkObject.listDepth) {
                                vm.list[itemIndex].isHide = false;
                                tmpParentIsShrinkIndex = itemIndex;
                            } else if (!tmpParentShrinkObject.isShrink) {
                                vm.list[itemIndex].isHide = false;
                            }
                            break;
                        }
                    }
                    itemIndex++;
                    $event = $event.nextElementSibling;
                }
            }
            vm.fun.shrinkList = function ($event) {
                $event.stopPropagation();
                var tmp = {};
                tmp.targetDom = fun.getTargetEvent($event.target);
                tmp.itemIndex = fun.getTargetIndex($event.target);
                vm.list[tmp.itemIndex].isShrink = !vm.list[tmp.itemIndex].isShrink;
                fun.operateLevel(tmp.targetDom.getAttribute('eo-attr-depth'), tmp.targetDom.nextElementSibling, parseInt(tmp.itemIndex) + 1);
            }
            /**
             * 初始化单项表格
             */
            vm.$onInit = function () {
                var template = {
                    itemHtml: '',
                    operateHtml: '',
                    thItemHtml: '',
                    moreFunArrHtml: ''
                },
                    tmpItemDragHtml = '',
                    tmpDragObj;
                if (vm.mainObject.setting.draggable) {
                    try {
                        let tmpOriginDragCacheObj = vm.mainObject.setting.dragCacheObj;
                        vm.mainObject.setting.dragCacheObj = Object.assign({}, tmpOriginDragCacheObj, JSON.parse(window.localStorage.getItem(vm.mainObject.setting.dragCacheVar)) || {});
                        for (let key in vm.mainObject.setting.dragCacheObj) {
                            if (vm.mainObject.setting.dragCacheObj[key] === '0px') {
                                vm.mainObject.setting.dragCacheObj[key] = tmpOriginDragCacheObj[key];
                            }
                        }
                    } catch (JSON_PARSE_ERROR) {
                        console.error(JSON_PARSE_ERROR);
                    }
                    tmpDragObj = {
                        setting: {
                            object: 'width',
                            affectCount: 2,
                            minWidth: 30
                        },
                        baseFun: {
                            mouseup: privateFun.dragMouseup
                        }
                    }
                }
                if (vm.mainObject.setting.batch) {
                    template.itemHtml = fun.generateHtml('select');
                }
                if (vm.mainObject.item.default) {
                    if (vm.mainObject.setting.autoSort) {
                        let storageOrderBy = window.localStorage[vm.mainObject.setting.sortStorageKey];
                        if (storageOrderBy) {
                            vm.data.listOrderBy = JSON.parse(storageOrderBy);
                        } else {
                            vm.data.listOrderBy = vm.mainObject.setting.sortDefaultVal.storageVal;
                        }
                    }
                    angular.forEach(vm.mainObject.item.default, function (listItem, thKey) {
                        let thItemContent = '';
                        switch (listItem.thType) {
                            case 'html': {
                                thItemContent = listItem.thHtml;
                                break;
                            }
                            default: {
                                thItemContent = `<span>${listItem.key}</span>`;
                            }
                        }
                        if (vm.mainObject.setting.autoSort && listItem.sort) {
                            if (vm.data.listOrderBy.orderBy === listItem.sortOrderByVal) {
                                listItem.asc = vm.data.listOrderBy.asc;
                            } else {
                                listItem.asc = vm.mainObject.setting.sortDefaultVal.sortOrder === 'asc' ? 1 : 0;
                            }
                        }
                        if (vm.mainObject.setting.draggable) {
                            listItem.draggableMainObject = Object.assign({}, tmpDragObj, {
                                mark: listItem.draggableCacheMark
                            });
                        }
                        template.thItemHtml += `<th class="${vm.mainObject.setting.draggaRootClass ? (vm.mainObject.setting.draggaRootClass + thKey) : ('th_drag_' + thKey + '_ldcc')} po_re ${listItem.thClass || ''} ${listItem.class || ''}" 
                        ${vm.mainObject.setting.autoSort && listItem.sort ? `ng-class="{\'hover_th_ldcc\':$ctrl.mainObject.item.default[${thKey}].sort&&(!$ctrl.otherObject.batchMenu.isOperating&&(!$ctrl.otherObject.batch.isOperating||$ctrl.mainObject.setting.batchInitStatus===\'open\'))}" ng-mousedown="$ctrl.fun.autoSortFun({item:$ctrl.mainObject.item.default[${thKey}],listItem:$ctrl.mainObject.item.default[${thKey}]})"` : ''}
                        ${vm.mainObject.setting.draggable ? `style="width:${vm.mainObject.setting.dragCacheObj[listItem.draggableCacheMark]}"` : ''}  
                        ${listItem.showVariable ? `ng-show="$ctrl.showObject['${listItem.showPoint}']['${listItem.showVariable}']==${listItem.show}"` : listItem.showPoint ? `ng-show="$ctrl.showObject['${listItem.showPoint}']==${listItem.show}" ` : ''} 
                        >${thItemContent}  
                        ${vm.mainObject.setting.autoSort && listItem.sort ? `<span  ng-if="$ctrl.mainObject.item.default[${thKey}].sort" class="iconfont  focus_orderby" ng-class="{'un_focus_orderBy':$ctrl.data.listOrderBy.orderBy!==${typeof(listItem.sortOrderByVal)==="string"?(`'${listItem.sortOrderByVal}'`):listItem.sortOrderByVal},'icon-xuanzeqizhankai_o':!$ctrl.mainObject.item.default[${thKey}].asc,'icon-xuanzeqishouqi_o':$ctrl.mainObject.item.default[${thKey}].asc}"></span>` : ''}
                        <div class="divide_line_ldcc ${vm.mainObject.setting.draggable ? 'ccr drag_divide_line_ldcc' : ''}" ${vm.mainObject.setting.draggable ? `drag-Change-Spacing-Common-Directive main-object='$ctrl.mainObject.item.default[${thKey}].draggableMainObject' container-affect-class="conatiner_ldcc" affect-Class="${vm.mainObject.setting.draggaRootClass ? (vm.mainObject.setting.draggaRootClass + thKey) : ('th_drag_' + thKey + '_ldcc')}"` : ''}>&nbsp;</div></th>`;
                        tmpItemDragHtml += `<th class="${vm.mainObject.setting.draggaRootClass ? (vm.mainObject.setting.draggaRootClass + thKey) : ('th_drag_' + thKey + '_ldcc')} po_re ${listItem.thClass || ''} ${listItem.class || ''}" ${vm.mainObject.setting.draggable ? `style="width:${vm.mainObject.setting.dragCacheObj[listItem.draggableCacheMark]}"` : ''} ${listItem.showVariable ? `ng-show="$ctrl.showObject['${listItem.showPoint}']['${listItem.showVariable}']==${listItem.show}"` : listItem.showPoint ? `ng-show="$ctrl.showObject['${listItem.showPoint}']==${listItem.show}" ` : ''}  >&nbsp;</th>`;
                        switch (listItem.type) {
                            case 'depthHtml': {
                                vm.data.isDepth = true;
                                template.itemHtml += `<td class="td-tbd text-td-tbd plr5 ${listItem.class || ''}">` +
                                    `<div class="depth-td-tbd" ng-init="item.listDepthArray=$ctrl.fun.range(item.listDepth+1,{item:item,$index:$outerIndex})">` +
                                    '<button type="button" class="btn-shrink iconfont" ng-click="$ctrl.fun.shrinkList($event)" ng-class="{\'icon-pinleizengjia\':item.isShrink,\'icon-pinleijianshao\':!item.isShrink}" ng-if="$ctrl.list[$index+1].listDepth>item.listDepth"></button>' +
                                    '<span class="divide-td-tbd" ng-class="{\'first-divide-td-tbd\':item.listDepth==$index}" ng-repeat="key in item.listDepthArray track by $index" ng-style="{\'left\':(15*$index+30)+\'px\'}" ng-hide="item.isShrink&&item.listDepth==$index"></span>' +
                                    listItem.html +
                                    '</div>' +
                                    '</td>';
                                break;
                            }
                            default: {
                                template.itemHtml += `<td class="${listItem.isUnneccessary ? 'eo_theme_ldt_tdt ' : ''} ${listItem.class || ''} ${listItem.contentClass || ''}" ` +
                                    `${listItem.showVariable ? `ng-show="$ctrl.showObject['${listItem.showPoint}']['${listItem.showVariable}']==${listItem.show}"` : listItem.showPoint ? `ng-show="$ctrl.showObject['${listItem.showPoint}']==${listItem.show}" ` : ''} ` +
                                    `${listItem.switch ? `ng-switch="item.${listItem.switch}" ` : ''}` +
                                    `${listItem.title ? `title="${listItem.title}" ` : ''}` +
                                    `>${listItem.html}</td>\n`;
                            }
                        }
                    })
                }
                if (vm.mainObject.item.operate) {
                    angular.forEach(vm.mainObject.item.operate.funArr, function (button, key) {
                        switch (button.type) {
                            case 'more': {
                                angular.forEach(button.funArr, function (moreButton, moreButtonKey) {
                                    if (moreButton.type == 'html') {
                                        template.moreFunArrHtml += moreButton.html;
                                    } else {
                                        template.moreFunArrHtml += '<li class="' + (moreButton.class || ' ') + ' eo_theme_gd_li_fli" ' + (moreButton.itemExpression || '') + (moreButton.showPoint ? ` ng-show="item['${moreButton.showPoint}']==${moreButton.show}" ` : '') + (moreButton.fun ? ' ng-click="$ctrl.fun.common($ctrl.mainObject.item.operate.funArr[' + key + '].funArr[' + moreButtonKey + '],{item:item,$index:$outerIndex,$event:$event})"' : '') + '>' + moreButton.key + '</li>';
                                    }
                                })
                                template.operateHtml += '<div ng-mouseleave="item.clickMore=false;" class="' + (button.class ? button.class : '') + 'more-btn eo-operate-btn po_re "' + (button.showPoint ? ' ng-show="item[$ctrl.mainObject.item.operate.funArr[' + key + '].showPoint]==$ctrl.mainObject.item.operate.funArr[' + key + '].show"' : '') + (button.itemExpression || '') + 'ng-click="$ctrl.fun.showMore({item:item,$event:$event})"' + (button.disabled ? 'ng-disabled="||($ctrl.mainObject.item.operate.funArr[' + key + '].disabled&&item[$ctrl.mainObject.item.operate.funArr[' + key + '].disabledVar]==$ctrl.mainObject.item.operate.funArr[' + key + '].disabled)||(' + vm.mainObject.setting.disabled + '&&$ctrl.mainObject.item.operate.funArr[' + key + '].status!==\'allowed\')"' : '') + '>{{$ctrl.mainObject.item.operate.funArr[' + key + '].key}}<span class="more-btn-icon iconfont icon-xuanzeqizhankai_o"></span>' +
                                    '<ul class="more-function n_dv_c n_list1_style po_ab " ng-show="item.clickMore==true">' + template.moreFunArrHtml + '</ul>' +
                                    '</div>';
                                break;
                            }
    
                            case 'html': {
                                template.operateHtml += button.html;
                                break;
                            }
                            default: {
                                template.operateHtml += '<button class="';
                                if (button.class) {
                                    template.operateHtml += button.class;
                                }
                                template.operateHtml += ' eo-operate-btn"';
                                if (button.itemExpression) {
                                    template.operateHtml += button.itemExpression;
                                }
                                if (button.authority) {
                                    template.operateHtml += ' ng-if="$ctrl.authorityObject[\'' + button.authority + '\']"';
                                }
                                if (button.authorityPoint) {
                                    template.operateHtml += ' ng-if="item[\'' + button.authorityPoint + '\']==' + button.authorityVal + '"';
                                }
                                if (button.showPoint) {
                                    template.operateHtml += ' ng-show="item[\'' + button.showPoint + '\']==' + button.show + '"';
                                }
                                if (button.showHtml) {
                                    template.operateHtml += ' ng-show="' + button.showHtml + '"';
                                }
                                if (button.disabledHtml) {
                                    template.operateHtml += ' ng-disabled="' + button.disabledHtml + '" ';
                                }
                                if (button.disabledVar) {
                                    template.operateHtml += ' ng-disabled="item[\'' + button.disabledVar + '\']==' + button.disabled + '" ';
                                }
                                if (button.fun) {
                                    template.operateHtml += ' ng-click="$ctrl.fun.common($ctrl.mainObject.item.operate.funArr[' + key + '],{item:item,$index:$outerIndex,$event:$event})"';
                                }
                                template.operateHtml += ">";
                                if (button.icon) {
                                    template.operateHtml += '<span class="iconfont icon-' + button.icon + '"></span>'
                                }
                                template.operateHtml += button.key + '</button>';
                            }
                        }
                    })
                    template.operateHtml = '<td class="operate-td  ' + (vm.mainObject.item.operate.class || '') + '" ' + (vm.mainObject.item.operate.listExpression || '') + (vm.mainObject.setting.batch ? 'ng-hide="$ctrl.mainObject.setting.batchInitStatus!==\'open\'&&$ctrl.otherObject.batch.isOperating"' : '') + ' ng-if="$ctrl.authorityObject.operate">' +
                        '<div ' + (vm.mainObject.item.operate.hideKey ? ('ng-if="!(' + vm.mainObject.item.operate.hideKey + ')"') : '') + '>' +
                        template.operateHtml +
                        '</div>' +
                        '</td>';
                }
                let tmpThHtml = (vm.mainObject.setting.batch ? `<th class="${!vm.mainObject.setting.page ? 'select_checkbox w_30' : 'select_checkbox w_50'}"  ng-show="$ctrl.otherObject.batch.isOperating">` +
                    (!vm.mainObject.setting.page ? `<button  type="button" ng-click="$ctrl.fun.selectAll()" class="eo-checkbox iconfont" ng-class="{\'icon-duihao\':$ctrl.otherObject.batch.selectAll==true}">{{$ctrl.otherObject.batch.selectAll?"":"&nbsp;"}}</button> ` :
                        `<div class="select_all_box f_row f_ac cp" ng-click="$ctrl.fun.selectAll({$event:$event})">
                        <button  type="button"  class="eo-checkbox iconfont" ng-class="{\'icon-duihao\':$ctrl.otherObject.batch.selectAll==true}">{{$ctrl.otherObject.batch.selectAll?"":"&nbsp;"}}</button>
                        <div ng-mouseleave="$ctrl.data.selectAllMore=false;" class="btn_all_show_more select_all_show_more">
                        <span class="btn_all_show_more iconfont icon-xuanzeqizhankai fs12 fwb"></span>
            </div>`) + '</th>' : '') + template.thItemHtml + '<th class="{{$ctrl.mainObject.item.operate.class}}" ng-style="$ctrl.mainObject.item.operate.style"' + (vm.mainObject.item.operate&&vm.mainObject.item.operate.listExpression? vm.mainObject.item.operate.listExpression: '') + '  ng-if="$ctrl.authorityObject.operate&&$ctrl.mainObject.item.operate" ng-hide="$ctrl.mainObject.setting.batchInitStatus!==\'open\'&&$ctrl.otherObject.batch.isOperating">' + (vm.mainObject.setting.operateThKey || '操作') + `</th>${vm.mainObject.setting.draggable ? '<th></th>' : ''}`;
                let tmpDragHtml = (vm.mainObject.setting.batch ? `<th class="${!vm.mainObject.setting.page ? 'select_checkbox w_30' : 'select_checkbox w_50'}"  ng-show="$ctrl.otherObject.batch.isOperating">&nbsp;</th>` : '') + tmpItemDragHtml + `<th class="{{$ctrl.mainObject.item.operate.class}}" ng-style="$ctrl.mainObject.item.operate.style" ${vm.mainObject.item.operate&&vm.mainObject.item.operate.listExpression? vm.mainObject.item.operate.listExpression: ''} ng-if="$ctrl.authorityObject.operate&&$ctrl.mainObject.item.operate" ng-hide="$ctrl.mainObject.setting.batchInitStatus!==\'open\'&&$ctrl.otherObject.batch.isOperating">&nbsp;</th>${vm.mainObject.setting.draggable ? '<th></th>' : ''}`;
    
                template.html = `<tr ng-hide="item.isHide&&$ctrl.data.isDepth"  eo-attr-index="{{$index}}" eo-attr-depth="{{item.listDepth}}"    class="${(vm.mainObject.setting.unhover ? 'unhover-tr' : 'hover-tr')}" ng-style="$ctrl.mainObject.setting.style" {{trExpression}}  ${vm.mainObject.item.primaryKey ? `ng-class="{'eo_theme_lct_tra':$ctrl.otherObject.batch.indexAddress[item.${vm.mainObject.item.primaryKey}],{{trNgClass}}}"` : 'ng-class="{{{trNgClass}}}"'}  ng-repeat='($outerIndex,item) in $ctrl.list' ng-click="$ctrl.fun.click({item:item,$index:$index,$event:$event})" ng-init="item.$index=$index">${template.itemHtml + template.operateHtml}${vm.mainObject.setting.draggable ? '<td></td>' : ''}</tr>`;
                try {
                    template.html = template.html.replace('{{trExpression}}', vm.mainObject.setting.trExpression || '');
                    template.html = template.html.replace('{{trNgClass}}', vm.mainObject.setting.trNgClass || '');
                } catch (REPLACE_ERR) {
                    console.error(REPLACE_ERR)
                }
                vm.data.tableHtml = '<article class="eo_theme_ldt first_level_article ' + (vm.mainObject.setting.isFixedHeight ? ' eo_theme_lrd fixed-height-list ' : '') + (vm.mainObject.setting.isGreyShading ? 'eo_theme_lrt' : '') + ' ">' +
                    '<div class="conatiner_ldcc '+ (vm.mainObject.setting.draggaRootClass || "") + (vm.mainObject.setting.draggable ? ' conatiner_ldcc_draggable ' : '') + ((vm.mainObject.setting.page||vm.mainObject.setting.fixFoot) ? 'conatiner_ldcc_has_footer' : '') + ' "><div class="thead_container_ldcc ' + (vm.mainObject.setting.trClass || '') + ' " ><table ' + (vm.mainObject.setting && vm.mainObject.setting.batch ? 'ng-class="{\'batchOperating\':$ctrl.otherObject.batch.isOperating}"' : '') + '>' +
                    '<thead>' +
                    `<tr>${tmpThHtml}</tr>` +
                    `</thead></table></div>` +
                    (vm.mainObject.setting.page ? `<div ng-click="$ctrl.fun.selectAll({$event:$event})"  ng-mouseover="$ctrl.data.selectAllMore=true" ng-mouseleave="$ctrl.data.selectAllMore=false" class="select_all_placeholder po_ab" ng-show="$ctrl.data.selectAllMore===true" >
                    <ul class="select_all_ul n_dv_c n_list1_style"><li class="btn_select_all eo_theme_gd_li_fli  select_all_item">选择所有数据 （共{{$ctrl.pageObject.pageInfo.msgCount}}条）</li><li class="btn_select_view eo_theme_gd_li_fli select_all_item">选择可见数据 （共{{(($ctrl.pageObject.pageInfo.page*$ctrl.pageObject.pageInfo.pageSize+($ctrl.pageObject.pageInfo.extraOprNum||0))>$ctrl.pageObject.pageInfo.msgCount)?$ctrl.pageObject.pageInfo.msgCount:($ctrl.pageObject.pageInfo.pageSize*$ctrl.pageObject.pageInfo.page+($ctrl.pageObject.pageInfo.extraOprNum||0))}}条）</li></ul> 
                 </div>` : '') +
                    `<div class="tbody_container_ldcc ${vm.mainObject.setting.trClass || ''} " >
                    <table` +
                    (vm.mainObject.setting.scroll ? ' infinite-scroll="$ctrl.mainObject.baseFun.scrollLoading()" infinite-scroll-parent infinite-scroll-distance="$ctrl.mainObject.setting.scrollRemainRatio||0">' : '>') +
                    ('<thead class="vis_hid unplaceholder_thead_ldcc"><tr>' + tmpDragHtml + '</tr></thead>') +
                    '<tbody class="list-default-tbody">' + template.html +
                    '</tbody></table><div class="none_div" ng-if="$ctrl.list.length===0"> ' + (vm.mainObject.setting.warning || "尚无任何内容") + ' </div>'+(vm.mainObject.setting.defaultFoot?`<div class="bottom-count-div "><span>共{{$ctrl.list.length}}条记录</span></div>`:"")+'</div></div>' +
                    (vm.mainObject.setting.page ?
                        `<div class="footer">
                    <span ng-if="$ctrl.mainObject.setting.batchInitStatus!=='open'&&$ctrl.otherObject.batch.isOperating">已选择{{$ctrl.otherObject.batch.query.length}}条记录</span>
                    <span ng-if="$ctrl.mainObject.setting.batchInitStatus==='open'||!$ctrl.otherObject.batch.isOperating">已加载{{(($ctrl.pageObject.pageInfo.page*$ctrl.pageObject.pageInfo.pageSize+($ctrl.pageObject.pageInfo.extraOprNum||0))>$ctrl.pageObject.pageInfo.msgCount)?$ctrl.pageObject.pageInfo.msgCount:($ctrl.pageObject.pageInfo.pageSize*$ctrl.pageObject.pageInfo.page+($ctrl.pageObject.pageInfo.extraOprNum||0))}}条记录，共{{$ctrl.pageObject.pageInfo.msgCount}}条记录</span>
                    </div>` : vm.mainObject.setting.fixFoot ?`<div class="footer"><span>共{{$ctrl.list.length}}条记录</span></div>`:"") +
                    '</article>';
                //     <footer class="pageFooter" ng-show="$ctrl.list.length>0">
                //     <uib-pagination total-items="$ctrl.pageObject.pageInfo.msgCount"
                //         items-per-page="$ctrl.pageObject.pageInfo.pageSize" ng-model="$ctrl.pageObject.pageInfo.page"
                //         max-size="$ctrl.pageObject.pageInfo.maxSize" boundary-link-Number="true" rotate="false"
                //         next-text="&#xeb5b;" previous-text="&#xeb5a;" ng-change="$ctrl.fun.changePage()"></uib-pagination>
                // </footer>
                // <span class="mlr15 divide-span" ng-if="$ctrl.list.length>0"></span>
    
            }
            fun.countItemSelectIsAll = function (inputBool) {
                if (inputBool) {
                    if (vm.mainObject.setting.page) {
                        let returnFlag = false;
                        for (var i = 0; i < vm.list.length; i++) {
                            if (vm.otherObject.batch.query.indexOf(vm.list[i][vm.mainObject.item.primaryKey]) === -1) {
                                returnFlag = true;
                                break;
                            }
                        }
                        if (vm.list.length && !returnFlag) vm.otherObject.batch.selectAll = true;
                    } else {
                        if ((vm.otherObject.batch.query || []).length == (vm.list || []).length) {
                            vm.otherObject.batch.selectAll = true;
                        }
                    }
                } else {
                    vm.otherObject.batch.selectAll = false;
                }
            }
            vm.fun.click = function (arg) {
                var template = {
                    $index: 0,
                    batchFun: (arg) => {
                        template.$index = vm.otherObject.batch.query.indexOf(arg.item[vm.mainObject.item.primaryKey])
                        if (vm.otherObject.batch.indexAddress[arg.item[vm.mainObject.item.primaryKey]]) {
                            vm.otherObject.batch.query.splice(template.$index, 1);
                            delete vm.otherObject.batch.indexAddress[arg.item[vm.mainObject.item.primaryKey]];
                            fun.countItemSelectIsAll(false);
                        } else {
                            vm.otherObject.batch.query.push(arg.item[vm.mainObject.item.primaryKey]);
                            vm.otherObject.batch.indexAddress[arg.item[vm.mainObject.item.primaryKey]] = arg.$index + 1;
                            fun.countItemSelectIsAll(true);
                        }
                        if (vm.mainObject.baseFun && vm.mainObject.baseFun.batchFilter) {
                            vm.mainObject.baseFun.batchFilter(template.$index === -1 ? 'select' : 'cancel', arg)
                        }
                    }
                }
                if (vm.mainObject.setting.batch && vm.otherObject.batch.isOperating && !vm.mainObject.setting.clickAffectBatch) {
                    template.batchFun(arg);
                } else if (vm.mainObject.baseFun && vm.mainObject.baseFun.click) {
                    vm.mainObject.baseFun.click(arg, template.batchFun);
                }
            }
            /**
             * 初始化单项表格
             */
            vm.fun.selectAll = function (arg) {
                if (vm.mainObject.baseFun && vm.mainObject.baseFun.selectAll) {
                    if (vm.mainObject.setting.page) {
                        let point = 'default';
                        try {
                            point = arg.$event.target.classList[0];
                        } catch (e) { }
                        switch (point) {
                            case "btn_all_show_more": {
                                vm.data.selectAllMore = true;
                                break;
                            }
                            case "eo-checkbox":
                            case "btn_select_all": {
                                if (point === "btn_select_all" || !vm.otherObject.batch.selectAll) {
                                    //选择所有数据
                                    vm.mainObject.baseFun.selectAll('selectAll');
                                    vm.data.selectAllMore = false;
                                } else {
                                    vm.mainObject.baseFun.selectAll('cancelAll');
                                }
                                break;
                            }
                            case "btn_select_view": {
                                //选择可见数据
                                vm.mainObject.baseFun.selectAll('selectView');
                                vm.data.selectAllMore = false;
                                break;
                            }
                        }
                    } else {
                        vm.mainObject.baseFun.selectAll(arg);
                    }
                } else {
                    let selectViewDataFun = () => {
                        if (vm.otherObject.batch.selectAll) {
                            vm.otherObject.batch.indexAddress = {};
                            vm.otherObject.batch.query = [];
                        } else {
                            vm.otherObject.batch.query = [];
                            let tmpList = vm.otherObject.allList ? vm.otherObject.allList : vm.list;
                            tmpList.map(function (val, key) {
                                vm.otherObject.batch.query.push(val[vm.mainObject.item.primaryKey]);
                                vm.otherObject.batch.indexAddress[val[vm.mainObject.item.primaryKey]] = key + 1;
                            })
                        }
                        vm.otherObject.batch.selectAll = !vm.otherObject.batch.selectAll;
                    }
                    if (vm.mainObject.setting.page) {
                        let point = 'default';
                        try {
                            point = arg.$event.target.classList[0];
                        } catch (e) { }
                        switch (point) {
                            case "btn_all_show_more": {
                                vm.data.selectAllMore = true;
                                break;
                            }
                            case "eo-checkbox":
                            case "btn_select_all": {
                                if (point === "btn_select_all" || !vm.otherObject.batch.selectAll) {
                                    //选择所有数据
                                    if (vm.pageObject.pageInfo.page === Math.ceil(vm.pageObject.pageInfo.msgCount / vm.pageObject.pageInfo.pageSize)) {
                                        vm.otherObject.batch.query = [];
                                        vm.otherObject.batch.indexAddress = {};
                                        vm.list.map(function (val, key) {
                                            vm.otherObject.batch.query.push(val[vm.mainObject.item.primaryKey]);
                                            vm.otherObject.batch.indexAddress[val[vm.mainObject.item.primaryKey]] = key + 1;
                                        })
                                    } else {
                                        vm.otherObject.batch.query = angular.copy(vm.otherObject.allQueryID);
                                        vm.otherObject.allQueryID.map(function (val, key) {
                                            vm.otherObject.batch.indexAddress[val] = key + 1;
                                        })
                                    }
                                    vm.data.selectAllMore = false;
                                    vm.otherObject.batch.selectAll = true;
                                } else {
                                    vm.otherObject.batch.indexAddress = {};
                                    vm.otherObject.batch.query = [];
                                    vm.otherObject.batch.selectAll = false;
                                }
                                break;
                            }
                            case "btn_select_view": {
                                //选择可见数据
                                vm.data.selectAllMore = false;
                                vm.otherObject.batch.query = [];
                                vm.otherObject.batch.indexAddress = {};
                                vm.list.map(function (val, key) {
                                    vm.otherObject.batch.query.push(val[vm.mainObject.item.primaryKey]);
                                    vm.otherObject.batch.indexAddress[val[vm.mainObject.item.primaryKey]] = key + 1;
                                })
                                vm.otherObject.batch.selectAll = true;
                                break;
                            }
                        }
                    } else {
                        selectViewDataFun();
                    }
    
                }
            }
            /**
             * @description 统筹绑定调用页面列表功能单击函数
             * @param {extend} obejct 方式值
             * @param {object} arg 共用体变量，后根据传值函数回调方法
             */
            vm.fun.common = function (extend, arg) {
                if (arg.$event) {
                    arg.$event.stopPropagation();
                }
                var template = {
                    params: angular.copy(arg)
                }
                switch (typeof (extend.params)) {
                    case 'string': {
                        return eval('extend.fun(' + extend.params + ')');
                    }
                    default: {
                        for (var key in extend.params) {
                            if (extend.params[key] == null) {
                                template.params[key] = arg[key];
                            } else {
                                template.params[key] = extend.params[key];
                            }
                        }
                        return extend.fun(template.params);
                    }
                }
            }
            vm.fun.changePage = function () {
                if (vm.pageObject.changeFun) {
                    vm.pageObject.changeFun();
                } else {
                    $state.go($state.current.name, {
                        page: vm.pageObject.pageInfo.page
                    });
                }
    
            }
        }
})();