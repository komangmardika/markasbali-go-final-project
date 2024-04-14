import {Component, OnInit} from '@angular/core';
import {CommonService} from "./services/common.service";
import {BackupListInterface} from "./models/backup-list.interface";

@Component({
  selector: 'app-tab4',
  templateUrl: 'tab4.page.html',
  styleUrls: ['tab4.page.scss']
})
export class Tab4Page implements OnInit{
  backupList: BackupListInterface[] = [];
  constructor(
    private commonService: CommonService
  ) {}

  ngOnInit() {
    this.commonService.latestBackedUpAllDatabases().subscribe(r => {
      this.backupList = r.data as BackupListInterface[];
    }, err => {
      console.log(err);
    })
  }

  refresh() {
    this.commonService.latestBackedUpAllDatabases().subscribe(r => {
      this.backupList = r.data as BackupListInterface[];
    }, err => {
      console.log(err);
    })
  }

}
