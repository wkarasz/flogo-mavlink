"use strict";
var __extends = this && this.__extends || function() {
        var t = Object.setPrototypeOf || {
            __proto__: []
        }
        instanceof Array && function(t, e) {
            t.__proto__ = e
        } || function(t, e) {
            for (var n in e) e.hasOwnProperty(n) && (t[n] = e[n])
        };
        return function(e, n) {
            function i() {
                this.constructor = e
            }
            t(e, n), e.prototype = null === n ? Object.create(n) : (i.prototype = n.prototype, new i)
        }
    }(),
    __decorate = this && this.__decorate || function(t, e, n, i) {
        var r, o = arguments.length,
            a = o < 3 ? e : null === i ? i = Object.getOwnPropertyDescriptor(e, n) : i;
        if ("object" == typeof Reflect && "function" == typeof Reflect.decorate) a = Reflect.decorate(t, e, n, i);
        else
            for (var l = t.length - 1; l >= 0; l--)(r = t[l]) && (a = (o < 3 ? r(a) : o > 3 ? r(e, n, a) : r(e, n)) || a);
        return o > 3 && a && Object.defineProperty(e, n, a), a
    },
    __metadata = this && this.__metadata || function(t, e) {
        if ("object" == typeof Reflect && "function" == typeof Reflect.metadata) return Reflect.metadata(t, e)
    };
Object.defineProperty(exports, "__esModule", {
    value: !0
});
var http_1 = require("@angular/http"),
    core_1 = require("@angular/core"),
    Observable_1 = require("rxjs/Observable"),
    lodash = require("lodash"),
    wi_contrib_1 = require("wi-studio/app/contrib/wi-contrib"),
    MavLinkContribModule = function(t) {
        function e(e, n, i) {
            var r = t.call(this, e, n, i) || this;
            return r.injector = e, r.http = n, r.contribModelService = i, r.value = function(t, e) {
                switch (t) {
                    case "mavlinkConnection":
                        return Observable_1.Observable.create(function(t) {
                            var e = [];
                            wi_contrib_1.WiContributionUtils.getConnections(r.http, r.category).subscribe(function(n) {
                                n.forEach(function(t) {
                                    for (var n = 0; n < t.settings.length; n++)
                                        if ("name" === t.settings[n].name) {
                                            e.push({
                                                unique_id: wi_contrib_1.WiContributionUtils.getUniqueId(t),
                                                name: t.settings[n].value
                                            });
                                            break
                                        }
                                }), t.next(e)
                            })
                        })
                }
                return null
            }, r.validate = function(t, e) {
                	return null
                }
            }, r.action = function(t, e) {
                var n = r.getModelService(),
                    i = wi_contrib_1.CreateFlowActionResult.newActionResult();
                if (e.handler && e.handler.settings && e.handler.settings.length > 0) {
                    var o = e.getField("mavlinkConnection");
                    if (o && o.value) {
                        var a = n.createTriggerElement("MavLink/mavlink-trigger");
                        if (a && a.handler && a.handler.settings && a.handler.settings.length > 0)
                            for (var l = 0; l < a.handler.settings.length; l++)
                                if ("mavlinkConnection" === a.settings[l].name) {
                                    a.settings[l].value = o.value;
                                    break
                                } var c = n.createFlow(e.getFlowName(), e.getFlowDescription());
                        i = i.addTriggerFlowMapping(lodash.cloneDeep(a), lodash.cloneDeep(c))
                    }
                }
                return wi_contrib_1.ActionResult.newActionResult().setSuccess(!0).setResult(i)
            }, r.category = "MavLink", r
        }
        return __extends(e, t), e.prototype.getContextVar = function(t, e) {
            return t.getField(e) ? void 0 === t.getField(e).value ? "" : t.getField(e).value : ""
        }, e
    }(wi_contrib_1.WiServiceHandlerContribution);
MavLinkContribModule = __decorate([core_1.Injectable(), wi_contrib_1.WiContrib({}), core_1.Injectable(), __metadata("design:paramtypes", [core_1.Injector, http_1.Http, wi_contrib_1.WiContribModelService])], MavLinkContribModule), exports.MavLinkContribModule = MavLinkContribModule;
//# sourceMappingURL=trigger.js.map
