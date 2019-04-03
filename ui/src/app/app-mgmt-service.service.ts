import { Injectable } from '@angular/core';
import {ApplicationMgmtPromiseClient} from 'autodevopsappservicemodel/clientjs/ApplicationMgmt_grpc_web_pb'
import {ComponentMgmtPromiseClient, ComponentMgmtClient} from 'autodevopsappservicemodel/clientjs/ComponentMgmt_grpc_web_pb'
import {GetApplicationsRequest, GetApplicationsResponse} from 'autodevopsappservicemodel/clientjs/ApplicationMgmt_pb'
import {GetComponentRequest, GetComponentResponse} from 'autodevopsappservicemodel/clientjs/ComponentMgmt_pb'
import { environment } from '../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AppMgmtServiceService {

  serviceUrl = environment.autoDevOpsHost;
  applicationMgmtClient = new ApplicationMgmtPromiseClient(this.serviceUrl, null, null);
  componentMgmtClient = new ComponentMgmtPromiseClient(this.serviceUrl, null, null);

  constructor() { }

  getApplications(orgName: string): Promise<GetApplicationsResponse> {
    const req = new GetApplicationsRequest();
    req.setOrgname(orgName)

    return this.applicationMgmtClient.getAllApplications(req, null);
  }

  getComponents(orgName: string, appName: string): Promise<GetComponentResponse> {
    const req = new GetComponentRequest();
    req.setOrgname(orgName)
    req.setAppname(appName);

    return this.componentMgmtClient.getComponents(req, null);
  }
}
