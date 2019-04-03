import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import {AppMgmtServiceService} from '../app-mgmt-service.service'
import {ComponentDetail, Package} from 'autodevopsappservicemodel/clientjs/ComponentMgmt_pb'
import { ComponentUI } from '../component-ui';

@Component({
  selector: 'app-app-detail',
  templateUrl: './app-detail.component.html',
  styleUrls: ['./app-detail.component.scss']
})
export class AppDetailComponent implements OnInit {
  orgName: string;
  appName: string;
  applicationComponents: ComponentUI[];

  constructor(private route: ActivatedRoute, 
    private appMgmtService: AppMgmtServiceService) { }

  ngOnInit() {
    this.orgName = this.route.snapshot.paramMap.get('orgName');
    this.appName = this.route.snapshot.paramMap.get('appName');

    this.appMgmtService.getComponents(this.orgName, this.appName).then( resp => {
      this.applicationComponents = resp.getComponentsList().
        map(c => new ComponentUI(c));  
    });
  }
}
