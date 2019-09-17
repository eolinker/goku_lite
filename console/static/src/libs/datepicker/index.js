! function () {
    "use strict";
    angular.module("ui.bootstrap.datepickerPopup", ["ui.bootstrap.datepicker", "ui.bootstrap.position"]).value("$datepickerPopupLiteralWarning", !0).constant("uibDatepickerPopupConfig", {
        altInputFormats: [],
        appendToBody: !1,
        clearText: "Clear",
        closeOnDateSelection: !0,
        closeText: "Done",
        currentText: "Today",
        datepickerPopup: "yyyy-MM-dd",
        datepickerPopupTemplateUrl: "./libs/datepicker/html/datepickerPopup/popup.html",
        datepickerTemplateUrl: "./libs/datepicker/html/datepicker/datepicker.html",
        html5Types: {
            date: "yyyy-MM-dd",
            "datetime-local": "yyyy-MM-dd",
            month: "yyyy-MM"
        },
        onOpenFocus: !0,
        showButtonBar: !1,
        placement: "auto bottom-left"
    }).controller("UibDatepickerPopupController", ["$scope", "$element", "$attrs", "$compile", "$log", "$parse", "$window", "$document", "$rootScope", "$uibPosition", "dateFilter", "uibDateParser", "uibDatepickerPopupConfig", "$timeout", "uibDatepickerConfig", "$datepickerPopupLiteralWarning", function (e, t, a, n, i, o, r, p, u, l, s, c, d, f, m, g) {
        function k(t) {
            var a = c.parse(t, v, e.date);
            if (isNaN(a))
                for (var n = 0; n < C.length; n++)
                    if (a = c.parse(t, C[n], e.date), !isNaN(a)) return a;
            return a
        }

        function D(e) {
            if (angular.isNumber(e) && (e = new Date(e)), !e) return null;
            if (angular.isDate(e) && !isNaN(e)) return e;
            if (angular.isString(e)) {
                var t = k(e);
                if (!isNaN(t)) return c.toTimezone(t, U.getOption("timezone"))
            }
            return U.getOption("allowInvalid") ? e : void 0
        }

        function h(e, t) {
            var n = e || t;
            return !a.ngRequired && !n || (angular.isNumber(n) && (n = new Date(n)), !n || !(!angular.isDate(n) || isNaN(n)) || !!angular.isString(n) && !isNaN(k(n)))
        }

        function $(a) {
            if (e.isOpen || !e.disabled) {
                var n = S[0],
                    i = t[0].contains(a.target),
                    o = void 0 !== n.contains && n.contains(a.target);
                !e.isOpen || i || o || e.$apply(function () {
                    e.isOpen = !1
                })
            }
        }

        function b(a) {
            27 === a.which && e.isOpen ? (a.preventDefault(), a.stopPropagation(), e.$apply(function () {
                e.isOpen = !1
            }), t[0].focus()) : 40 !== a.which || e.isOpen || (a.preventDefault(), a.stopPropagation(), e.$apply(function () {
                e.isOpen = !0
            }))
        }

        function O() {
            if (e.isOpen) {
                var n = angular.element(S[0].querySelector(".uib-datepicker-popup")),
                    i = a.popupPlacement ? a.popupPlacement : d.placement,
                    o = l.positionElements(t, n, i, P);
                n.css({
                    top: o.top + "px",
                    left: o.left + "px"
                }), n.hasClass("uib-position-measure") && n.removeClass("uib-position-measure")
            }
        }

        function y(e) {
            var t;
            return angular.version.minor < 6 ? (t = angular.isObject(e.$options) ? e.$options : {
                timezone: null
            }, t.getOption = function (e) {
                return t[e]
            }) : t = e.$options, t
        }
        var v, w, P, T, x, M, N, z, B, F, U, S, C, E = !1,
            I = [];
        this.init = function (i) {
            if (F = i, U = y(F), w = angular.isDefined(a.closeOnDateSelection) ? e.$parent.$eval(a.closeOnDateSelection) : d.closeOnDateSelection, P = angular.isDefined(a.datepickerAppendToBody) ? e.$parent.$eval(a.datepickerAppendToBody) : d.appendToBody, T = angular.isDefined(a.onOpenFocus) ? e.$parent.$eval(a.onOpenFocus) : d.onOpenFocus, x = angular.isDefined(a.datepickerPopupTemplateUrl) ? a.datepickerPopupTemplateUrl : d.datepickerPopupTemplateUrl, M = angular.isDefined(a.datepickerTemplateUrl) ? a.datepickerTemplateUrl : d.datepickerTemplateUrl, C = angular.isDefined(a.altInputFormats) ? e.$parent.$eval(a.altInputFormats) : d.altInputFormats, e.showButtonBar = angular.isDefined(a.showButtonBar) ? e.$parent.$eval(a.showButtonBar) : d.showButtonBar, d.html5Types[a.type] ? (v = d.html5Types[a.type], E = !0) : (v = a.uibDatepickerPopup || d.datepickerPopup, a.$observe("uibDatepickerPopup", function (e, t) {
                    var a = e || d.datepickerPopup;
                    if (a !== v && (v = a, F.$modelValue = null, !v)) throw new Error("uibDatepickerPopup must have a date format specified.")
                })), !v) throw new Error("uibDatepickerPopup must have a date format specified.");
            if (E && a.uibDatepickerPopup) throw new Error("HTML5 date input types do not support custom formats.");
            N = angular.element('<div style="position:absolute;top:30px;left:-190px;z-index:-1" uib-datepicker-popup-wrap><div uib-datepicker></div></div>'), N.attr({
                "ng-model": "date",
                "ng-change": "dateSelection(date)",
                "template-url": x
            }), z = angular.element(N.children()[0]), z.attr("template-url", M), e.datepickerOptions || (e.datepickerOptions = {}), E && "month" === a.type && (e.datepickerOptions.datepickerMode = "month", e.datepickerOptions.minMode = "month"), z.attr("datepicker-options", "datepickerOptions"), E ? F.$formatters.push(function (t) {
                return e.date = c.fromTimezone(t, U.getOption("timezone")), t
            }) : (F.$$parserName = "date", F.$validators.date = h, F.$parsers.unshift(D), F.$formatters.push(function (t) {
                return F.$isEmpty(t) ? (e.date = t, t) : (angular.isNumber(t) && (t = new Date(t)), e.date = c.fromTimezone(t, U.getOption("timezone")), c.filter(e.date, v))
            })), F.$viewChangeListeners.push(function () {
                e.date = k(F.$viewValue)
            }), t.on("keydown", b), S = n(N)(e), N.remove(), P ? p.find("body").append(S) : t.after(S), e.$on("$destroy", function () {
                for (e.isOpen === !0 && (u.$$phase || e.$apply(function () {
                        e.isOpen = !1
                    })), S.remove(), t.off("keydown", b), p.off("click", $), B && B.off("scroll", O), angular.element(r).off("resize", O); I.length;) I.shift()()
            })
        }, e.getText = function (t) {
            return e[t + "Text"] || d[t + "Text"]
        }, e.isDisabled = function (t) {
            "today" === t && (t = c.fromTimezone(new Date, U.getOption("timezone")));
            var a = {};
            return angular.forEach(["minDate", "maxDate"], function (t) {
                e.datepickerOptions[t] ? angular.isDate(e.datepickerOptions[t]) ? a[t] = new Date(e.datepickerOptions[t]) : (g && i.warn("Literal date support has been deprecated, please switch to date object usage"), a[t] = new Date(s(e.datepickerOptions[t], "medium"))) : a[t] = null
            }), e.datepickerOptions && a.minDate && e.compare(t, a.minDate) < 0 || a.maxDate && e.compare(t, a.maxDate) > 0
        }, e.compare = function (e, t) {
            return new Date(e.getFullYear(), e.getMonth(), e.getDate()) - new Date(t.getFullYear(), t.getMonth(), t.getDate())
        }, e.dateSelection = function (a) {
            e.datepickerOptions._default ? e.date = new Date(e.datepickerOptions._default.year || a.getFullYear(), e.datepickerOptions._default.month || a.getMonth(), e.datepickerOptions._default.day || a.getDate()) : e.date = a;
            var n = e.date ? c.filter(e.date, v) : null;
            t.val(n), F.$setViewValue(n), w && (e.isOpen = !1, t[0].focus())
        }, e.keydown = function (a) {
            27 === a.which && (a.stopPropagation(), e.isOpen = !1, t[0].focus())
        }, e.select = function (t, a) {
            if (a.stopPropagation(), "today" === t) {
                var n = new Date;
                angular.isDate(e.date) ? (t = new Date(e.date), t.setFullYear(n.getFullYear(), n.getMonth(), n.getDate())) : (t = c.fromTimezone(n, U.getOption("timezone")), t.setHours(0, 0, 0, 0))
            }
            e.dateSelection(t)
        }, e.close = function (a) {
            a.stopPropagation(), e.isOpen = !1, t[0].focus()
        }, e.disabled = angular.isDefined(a.disabled) || !1, a.ngDisabled && I.push(e.$parent.$watch(o(a.ngDisabled), function (t) {
            e.disabled = t
        })), e.$watch("isOpen", function (n) {
            n ? e.disabled ? e.isOpen = !1 : f(function () {
                O(), T && e.$broadcast("uib:datepicker.focus"), p.on("click", $);
                var n = a.popupPlacement ? a.popupPlacement : d.placement;
                P || l.parsePlacement(n)[2] ? (B = B || angular.element(l.scrollParent(t)), B && B.on("scroll", O)) : B = null, angular.element(r).on("resize", O)
            }, 0, !1) : (p.off("click", $), B && B.off("scroll", O), angular.element(r).off("resize", O))
        }), e.$on("uib:datepicker.mode", function () {
            f(O, 0, !1)
        })
    }]).directive("uibDatepickerPopup", function () {
        return {
            require: ["ngModel", "uibDatepickerPopup"],
            controller: "UibDatepickerPopupController",
            scope: {
                datepickerOptions: "=?",
                isOpen: "=?",
                currentText: "@",
                clearText: "@",
                closeText: "@"
            },
            link: function (e, t, a, n) {
                var i = n[0],
                    o = n[1];
                o.init(i)
            }
        }
    }).directive("uibDatepickerPopupWrap", function () {
        return {
            restrict: "A",
            transclude: !0,
            templateUrl: function (e, t) {
                return t.templateUrl || "./libs/datepicker/html/datepickerPopup/popup.html"
            }
        }
    })
}();