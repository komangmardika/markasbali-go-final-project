import { Component } from '@angular/core';
import {Subject, takeUntil} from "rxjs";
import {CommonService} from "../tab4/services/common.service";
import {RestoreService} from "./services/restore.service";

@Component({
  selector: 'app-tab3',
  templateUrl: 'tab3.page.html',
  styleUrls: ['tab3.page.scss']
})
export class Tab3Page {
  public result = '';
  private destroy$: Subject<void> = new Subject<void>();
  constructor(
    private restoreService: RestoreService,
    private commonService: CommonService,
  ) {}

  public restoreDatabases(): void {
    this.result = 'Requesting please wait...';
    this.restoreService.restore().pipe(takeUntil(this.destroy$)).subscribe(res => {
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
