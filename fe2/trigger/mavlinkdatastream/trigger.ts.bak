import {Injectable, Injector} from "@angular/core";
import {Http} from "@angular/http";
import {Observable} from "rxjs/Observable";
import {
    WiContrib,
    WiServiceHandlerContribution,
    IValidationResult,
    ValidationResult,
    IContributionTypes,
    ActionResult,
    IActionResult,
    WiContribModelService,
    WiContributionUtils,
    IConnectorContribution
} from "wi-studio/app/contrib/wi-contrib";

@WiContrib({})
@Injectable()
export class MavLinkTriggerContribution extends WiServiceHandlerContribution {

    constructor(private injector: Injector, private http: Http) {
                    super(injector, http);
                }

    value = (fieldName: string, context: IContributionTypes): Observable<any> | any => {
        
        switch(fieldName) {
            case "connection":
                return Observable.create(observer => {
                    let connectionRefs = [];
                    WiContributionUtils.getConnections(this.http, "MavLink").subscribe((data: IConnectorContribution[]) => {
                        data.forEach(connection => {
                            if ((<any>connection).isValid) {
                                for(let i=0; i < connection.settings.length; i++) {
                                    if (connection.settings[i].name === "name") {
                                        connectionRefs.push({
                                            "unique_id": WiContributionUtils.getUniqueId(connection),
                                            "name": connection.settings[i].value
                                        });
                                        break;
                                    }
                                }
                            }
                        });
                        observer.next(connectionRefs);
                    });
                });
            default: 
                return null;
        }
            
    }

    validate = (fieldName: string, context: IContributionTypes): Observable<IValidationResult> | IValidationResult => {
        return Observable.create(observer => {
            let vresult: IValidationResult = ValidationResult.newValidationResult();
            observer.next(vresult);
        });
    }

    action = (actionId: string, context: IContributionTypes): Observable<IActionResult> | IActionResult => {
        return Observable.create(observer => {
            let aresult: IActionResult = ActionResult.newActionResult();
            observer.next(aresult);
        });
    }
}



