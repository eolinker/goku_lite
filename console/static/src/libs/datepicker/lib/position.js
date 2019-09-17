angular.module("ui.bootstrap.position", []).factory("$uibPosition", ["$document", "$window", function (t, e) {
    var o, i, r = {
            normal: /(auto|scroll)/,
            hidden: /(auto|scroll|hidden)/
        },
        a = {
            auto: /\s?auto?\s?/i,
            primary: /^(top|bottom|left|right)$/,
            secondary: /^(top|bottom|left|right|center)$/,
            vertical: /^(top|bottom)$/
        },
        n = /(HTML|BODY)/;
    return {
        getRawNode: function (t) {
            return t.nodeName ? t : t[0] || t
        },
        parseStyle: function (t) {
            return t = parseFloat(t), isFinite(t) ? t : 0
        },
        offsetParent: function (o) {
            function i(t) {
                return "static" === (e.getComputedStyle(t).position || "static")
            }
            o = this.getRawNode(o);
            for (var r = o.offsetParent || t[0].documentElement; r && r !== t[0].documentElement && i(r);) r = r.offsetParent;
            return r || t[0].documentElement
        },
        scrollbarWidth: function (r) {
            if (r) {
                if (angular.isUndefined(i)) {
                    var a = t.find("body");
                    a.addClass("uib-position-body-scrollbar-measure"), i = e.innerWidth - a[0].clientWidth, i = isFinite(i) ? i : 0, a.removeClass("uib-position-body-scrollbar-measure")
                }
                return i
            }
            if (angular.isUndefined(o)) {
                var n = angular.element('<div class="uib-position-scrollbar-measure"></div>');
                t.find("body").append(n), o = n[0].offsetWidth - n[0].clientWidth, o = isFinite(o) ? o : 0, n.remove()
            }
            return o
        },
        scrollbarPadding: function (t) {
            t = this.getRawNode(t);
            var o = e.getComputedStyle(t),
                i = this.parseStyle(o.paddingRight),
                r = this.parseStyle(o.paddingBottom),
                a = this.scrollParent(t, !1, !0),
                h = this.scrollbarWidth(n.test(a.tagName));
            return {
                scrollbarWidth: h,
                widthOverflow: a.scrollWidth > a.clientWidth,
                right: i + h,
                originalRight: i,
                heightOverflow: a.scrollHeight > a.clientHeight,
                bottom: r + h,
                originalBottom: r
            }
        },
        isScrollable: function (t, o) {
            t = this.getRawNode(t);
            var i = o ? r.hidden : r.normal,
                a = e.getComputedStyle(t);
            return i.test(a.overflow + a.overflowY + a.overflowX)
        },
        scrollParent: function (o, i, a) {
            o = this.getRawNode(o);
            var n = i ? r.hidden : r.normal,
                h = t[0].documentElement,
                l = e.getComputedStyle(o);
            if (a && n.test(l.overflow + l.overflowY + l.overflowX)) return o;
            var s = "absolute" === l.position,
                f = o.parentElement || h;
            if (f === h || "fixed" === l.position) return h;
            for (; f.parentElement && f !== h;) {
                var d = e.getComputedStyle(f);
                if (s && "static" !== d.position && (s = !1), !s && n.test(d.overflow + d.overflowY + d.overflowX)) break;
                f = f.parentElement
            }
            return f
        },
        position: function (o, i) {
            o = this.getRawNode(o);
            var r = this.offset(o);
            if (i) {
                var a = e.getComputedStyle(o);
                r.top -= this.parseStyle(a.marginTop), r.left -= this.parseStyle(a.marginLeft)
            }
            var n = this.offsetParent(o),
                h = {
                    top: 0,
                    left: 0
                };
            return n !== t[0].documentElement && (h = this.offset(n), h.top += n.clientTop - n.scrollTop, h.left += n.clientLeft - n.scrollLeft), {
                width: Math.round(angular.isNumber(r.width) ? r.width : o.offsetWidth),
                height: Math.round(angular.isNumber(r.height) ? r.height : o.offsetHeight),
                top: Math.round(r.top - h.top),
                left: Math.round(r.left - h.left)
            }
        },
        offset: function (o) {
            o = this.getRawNode(o);
            var i = o.getBoundingClientRect();
            return {
                width: Math.round(angular.isNumber(i.width) ? i.width : o.offsetWidth),
                height: Math.round(angular.isNumber(i.height) ? i.height : o.offsetHeight),
                top: Math.round(i.top + (e.pageYOffset || t[0].documentElement.scrollTop)),
                left: Math.round(i.left + (e.pageXOffset || t[0].documentElement.scrollLeft))
            }
        },
        viewportOffset: function (o, i, r) {
            o = this.getRawNode(o), r = r !== !1;
            var a = o.getBoundingClientRect(),
                n = {
                    top: 0,
                    left: 0,
                    bottom: 0,
                    right: 0
                },
                h = i ? t[0].documentElement : this.scrollParent(o),
                l = h.getBoundingClientRect();
            if (n.top = l.top + h.clientTop, n.left = l.left + h.clientLeft, h === t[0].documentElement && (n.top += e.pageYOffset, n.left += e.pageXOffset), n.bottom = n.top + h.clientHeight, n.right = n.left + h.clientWidth, r) {
                var s = e.getComputedStyle(h);
                n.top += this.parseStyle(s.paddingTop), n.bottom -= this.parseStyle(s.paddingBottom), n.left += this.parseStyle(s.paddingLeft), n.right -= this.parseStyle(s.paddingRight)
            }
            return {
                top: Math.round(a.top - n.top),
                bottom: Math.round(n.bottom - a.bottom),
                left: Math.round(a.left - n.left),
                right: Math.round(n.right - a.right)
            }
        },
        parsePlacement: function (t) {
            var e = a.auto.test(t);
            return e && (t = t.replace(a.auto, "")), t = t.split("-"), t[0] = t[0] || "top", a.primary.test(t[0]) || (t[0] = "top"), t[1] = t[1] || "center", a.secondary.test(t[1]) || (t[1] = "center"), e ? t[2] = !0 : t[2] = !1, t
        },
        positionElements: function (t, o, i, r) {
            t = this.getRawNode(t), o = this.getRawNode(o);
            var n = angular.isDefined(o.offsetWidth) ? o.offsetWidth : o.prop("offsetWidth"),
                h = angular.isDefined(o.offsetHeight) ? o.offsetHeight : o.prop("offsetHeight");
            i = this.parsePlacement(i);
            var l = r ? this.offset(t) : this.position(t),
                s = {
                    top: 0,
                    left: 0,
                    placement: ""
                };
            if (i[2]) {
                var f = this.viewportOffset(t, r),
                    d = e.getComputedStyle(o),
                    p = {
                        width: n + Math.round(Math.abs(this.parseStyle(d.marginLeft) + this.parseStyle(d.marginRight))),
                        height: h + Math.round(Math.abs(this.parseStyle(d.marginTop) + this.parseStyle(d.marginBottom)))
                    };
                if (i[0] = "top" === i[0] && p.height > f.top && p.height <= f.bottom ? "bottom" : "bottom" === i[0] && p.height > f.bottom && p.height <= f.top ? "top" : "left" === i[0] && p.width > f.left && p.width <= f.right ? "right" : "right" === i[0] && p.width > f.right && p.width <= f.left ? "left" : i[0], i[1] = "top" === i[1] && p.height - l.height > f.bottom && p.height - l.height <= f.top ? "bottom" : "bottom" === i[1] && p.height - l.height > f.top && p.height - l.height <= f.bottom ? "top" : "left" === i[1] && p.width - l.width > f.right && p.width - l.width <= f.left ? "right" : "right" === i[1] && p.width - l.width > f.left && p.width - l.width <= f.right ? "left" : i[1], "center" === i[1])
                    if (a.vertical.test(i[0])) {
                        var g = l.width / 2 - n / 2;
                        f.left + g < 0 && p.width - l.width <= f.right ? i[1] = "left" : f.right + g < 0 && p.width - l.width <= f.left && (i[1] = "right")
                    } else {
                        var u = l.height / 2 - p.height / 2;
                        f.top + u < 0 && p.height - l.height <= f.bottom ? i[1] = "top" : f.bottom + u < 0 && p.height - l.height <= f.top && (i[1] = "bottom")
                    }
            }
            switch (i[0]) {
                case "top":
                    s.top = l.top - h;
                    break;
                case "bottom":
                    s.top = l.top + l.height;
                    break;
                case "left":
                    s.left = l.left - n;
                    break;
                case "right":
                    s.left = l.left + l.width
            }
            switch (i[1]) {
                case "top":
                    s.top = l.top;
                    break;
                case "bottom":
                    s.top = l.top + l.height - h;
                    break;
                case "left":
                    s.left = l.left;
                    break;
                case "right":
                    s.left = l.left + l.width - n;
                    break;
                case "center":
                    a.vertical.test(i[0]) ? s.left = l.left + l.width / 2 - n / 2 : s.top = l.top + l.height / 2 - h / 2
            }
            return s.top = Math.round(s.top), s.left = Math.round(s.left), s.placement = "center" === i[1] ? i[0] : i[0] + "-" + i[1], s
        },
        adjustTop: function (t, e, o, i) {
            if (t.indexOf("top") !== -1 && o !== i) return {
                top: e.top - i + "px"
            }
        },
        positionArrow: function (t, o) {
            t = this.getRawNode(t);
            var i = t.querySelector(".tooltip-inner, .popover-inner");
            if (i) {
                var r = angular.element(i).hasClass("tooltip-inner"),
                    n = r ? t.querySelector(".tooltip-arrow") : t.querySelector(".arrow");
                if (n) {
                    var h = {
                        top: "",
                        bottom: "",
                        left: "",
                        right: ""
                    };
                    if (o = this.parsePlacement(o), "center" === o[1]) return void angular.element(n).css(h);
                    var l = "border-" + o[0] + "-width",
                        s = e.getComputedStyle(n)[l],
                        f = "border-";
                    f += a.vertical.test(o[0]) ? o[0] + "-" + o[1] : o[1] + "-" + o[0], f += "-radius";
                    var d = e.getComputedStyle(r ? i : t)[f];
                    switch (o[0]) {
                        case "top":
                            h.bottom = r ? "0" : "-" + s;
                            break;
                        case "bottom":
                            h.top = r ? "0" : "-" + s;
                            break;
                        case "left":
                            h.right = r ? "0" : "-" + s;
                            break;
                        case "right":
                            h.left = r ? "0" : "-" + s
                    }
                    h[o[1]] = d, angular.element(n).css(h)
                }
            }
        }
    }
}]);