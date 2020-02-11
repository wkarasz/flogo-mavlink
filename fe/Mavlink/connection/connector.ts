import { Inject, Injectable, Injector } from "@angular/core";
import { Http, Response, URLSearchParams, Headers, RequestOptions } from "@angular/http";
import { WiContrib, WiServiceHandlerContribution, AUTHENTICATION_TYPE, WiContributionUtils, WiProxyCORSUtils } from "wi-studio/app/contrib/wi-contrib";
import { IConnectorContribution } from "wi-studio/common/models/contrib";
import { IValidationResult, ValidationResult, ValidationError } from "wi-studio/common/models/validation";
import { IActionResult, ActionResult } from "wi-studio/common/models/contrib";
import { Observable } from "rxjs/Observable";
import { IFieldDefinition } from "wi-studio/common/models/contrib";

import * as lodash from "lodash";

@Injectable()
@WiContrib({})
export class MavlinkConnectorService extends WiServiceHandlerContribution {
    constructor( @Inject(Injector) injector, private http: Http) {
        super(injector, http);
    }

    value = (fieldName: string, context: IConnectorContribution): Observable<any> | any => {
        return null;
    }

    validate = (fieldName: string, context: IConnectorContribution): Observable<IValidationResult> | IValidationResult => {

        return null;
    }

    action = (name: string, context: IConnectorContribution): Observable<IActionResult> | IActionResult => {
        if (name == "Save") {
            return Observable.create(observer => {
                let actionResult = {
                    context: context,
                    authType: AUTHENTICATION_TYPE.BASIC,
                    authData: {}
                }
                observer.next(ActionResult.newActionResult().setSuccess(true).setResult(actionResult));
            });
        }
        return null;
    }
}
