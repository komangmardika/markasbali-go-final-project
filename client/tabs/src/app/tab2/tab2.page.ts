import { Component } from '@angular/core';
import {Subject, takeUntil} from "rxjs";
import {BackupService} from "./services/backup.service";

@Component({
  selector: 'app-tab2',
  templateUrl: 'tab2.page.html',
  styleUrls: ['tab2.page.scss']
})
export class Tab2Page {

  public result = '';
  private destroy$: Subject<void> = new Subject<void>();

  constructor(
    private backupService: BackupService
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

  }

}
