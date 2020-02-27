/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
import { Observable } from "rxjs/Observable";
import { Http } from "@angular/http";
import { Injectable, Injector, Inject } from "@angular/core";
import * as lodash from "lodash";
import {
    WiContrib,
    ActionResult,
    CreateFlowActionResult,
    ICreateFlowActionContext,
    WiServiceHandlerContribution,
    IValidationResult,
    ValidationResult,
    IFieldDefinition,
    ITriggerContribution,
    IConnectorContribution,
    IActionResult,
    WiContributionUtils
} from "wi-studio/app/contrib/wi-contrib";

@WiContrib({})
@Injectable()
export class MavLinkTriggerContribution extends WiServiceHandlerContribution {
    constructor( @Inject(Injector) injector, private http: Http) {
        super(injector, http);
    }

    value = (fieldName: string, context: ITriggerContribution): Observable<any> | any => {
        if (fieldName === "connection") {
            //Connector Type must match with the name defined in connector.json
            return Observable.create(observer => {
                let connectionRefs = [];
                WiContributionUtils.getConnections(this.http, "MavLink").subscribe((data: IConnectorContribution[]) => {
                    data.forEach(connection => {
                        for (let i = 0; i < connection.settings.length; i++) {
                            if (connection.settings[i].name === "name") {
                                connectionRefs.push({
                                    "unique_id": WiContributionUtils.getUniqueId(connection),
                                    "name": connection.settings[i].value
                                });
                                break;
                            }
                        }
                    });
                    observer.next(connectionRefs);
                });
            });
        } 
	return null;
    }

    validate = (fieldName: string, context: ITriggerContribution): Observable<IValidationResult> | IValidationResult => {
        if (fieldName === "connection") {
            let connection: IFieldDefinition = context.getField("connection");
            if (connection.value === null) {
                return ValidationResult.newValidationResult().setError("MavLink-1000", "MavLink Connection must be configured");
            }
        }
        return null;
    }
/*
    action = (fieldName: string, context: ICreateFlowActionContext): Observable<IActionResult> | IActionResult => {
        let modelService = this.getModelService();
        let result = CreateFlowActionResult.newActionResult();
        if (context.settings && context.settings.length > 0) {
            let connection = <IFieldDefinition>context.getField("sqsConnection");
            if (connection && connection.value) {
                let trigger = modelService.createTriggerElement("AWSSQS/sqsreceivemessage");
                if (trigger && trigger.settings  && trigger.settings.length > 0) {
                    for (let j = 0; j < trigger.settings.length; j++) {
                        if (trigger.settings[j].name === "sqsConnection") {
                            trigger.settings[j].value = connection.value;
                        }
                    }
                }
                let flowModel = modelService.createFlow(context.getFlowName(), context.getFlowDescription());
                result = result.addTriggerFlowMapping(lodash.cloneDeep(trigger), lodash.cloneDeep(flowModel));
            }
        }
        let actionResult = ActionResult.newActionResult().setSuccess(true).setResult(result);
        return actionResult;
    }
*/
}
