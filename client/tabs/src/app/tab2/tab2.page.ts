import { Component } from '@angular/core';
import {Subject, take, takeUntil} from "rxjs";
import {BackupService} from "./services/backup.service";
import {CommonService} from "../tab4/services/common.service";
import {BackupListInterface} from "../tab4/models/backup-list.interface";

@Component({
  selector: 'app-tab2',
  templateUrl: 'tab2.page.html',
  styleUrls: ['tab2.page.scss']
})
export class Tab2Page {

  public result = '';
  private destroy$: Subject<void> = new Subject<void>();

  constructor(
    private backupService: BackupService,
    private commonService: CommonService,
  ) {}

  public backupDatabases(): void {
    this.result = 'Requesting please wait...';
    this.backupService.backup().pipe(takeUntil(this.destroy$)).subscribe(res => {
      this.result = res;
    }, err => {
      this.result = err;
    })
  }

  public getLatestBackup(): void {
    this.result = 'Requesting please wait...';
    this.commonService.latestBackedUpAllDatabases().pipe(takeUntil(this.destroy$)).subscribe(r => {
      this.result = r
    }, err => {
      this.result = err;
    })
  }

}
