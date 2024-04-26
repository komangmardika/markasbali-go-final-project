import { Component } from '@angular/core';
import {Subject, takeUntil} from "rxjs";
import {CommonService} from "../tab4/services/common.service";
import {RestoreService} from "./services/restore.service";
import {ResetService} from "../tab1/services/reset.service";

@Component({
  selector: 'app-tab3',
  templateUrl: 'tab3.page.html',
  styleUrls: ['tab3.page.scss']
})
export class Tab3Page {
  public result = '';
  private destroy$: Subject<void> = new Subject<void>();
  alert = false;
  message = '';
  buttonList: string[] = [];

  constructor(
    private restoreService: RestoreService,
    private commonService: CommonService,
    private resetService: ResetService,
  ) {}

  public async restoreDatabases(): Promise<void> {
    this.result = 'Requesting please wait...';
    let f = await this.resetService.dbLst()
    this.restoreService.restore().pipe(takeUntil(this.destroy$)).subscribe(res => {
      this.result = res;
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
