import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import {AppMgmtServiceService} from '../app-mgmt-service.service'
import {ApplicationDetails} from 'autodevopsappservicemodel/clientjs/ApplicationMgmt_pb'

@Component({
  selector: 'app-app-list',
  templateUrl: './app-list.component.html',
  styleUrls: ['./app-list.component.scss']
})
export class AppListComponent implements OnInit {
  orgName: string;
  applications: ApplicationDetails[];

  constructor(private route: ActivatedRoute,
    private appMgmtService: AppMgmtServiceService) { }

  ngOnInit() {
    this.orgName = this.route.snapshot.paramMap.get('orgName');
    this.appMgmtService.getApplications(this.orgName).then( resp => {
      this.applications = resp.getApplicationsList();
    });
  }

}
