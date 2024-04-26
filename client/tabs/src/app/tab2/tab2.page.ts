import { Component } from '@angular/core';
import {Subject, take, takeUntil} from "rxjs";
import {BackupService} from "./services/backup.service";
import {CommonService} from "../tab4/services/common.service";
import {ResetService} from "../tab1/services/reset.service";

@Component({
  selector: 'app-tab2',
  templateUrl: 'tab2.page.html',
  styleUrls: ['tab2.page.scss']
})
export class Tab2Page {

  public result = '';
  private destroy$: Subject<void> = new Subject<void>();

  alert = false;
  message = '';
  buttonList: string[] = [];

  constructor(
    private backupService: BackupService,
    private commonService: CommonService,
    private resetService: ResetService,
  ) {}

  public async backupDatabases(): Promise<void> {
    this.result = 'Requesting please wait...';
    let f = await this.resetService.dbLst()

    this.backupService.backup().pipe(takeUntil(this.destroy$)).subscribe(res => {
      this.result = res;
      this.alert = true;
      this.message = f.data.length + ' Databases has been backed up successfully'
      this.buttonList = ['Close'];
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
