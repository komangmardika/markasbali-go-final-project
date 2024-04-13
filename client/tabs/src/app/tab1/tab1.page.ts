import { Component } from '@angular/core';
import {ResetService} from "./services/reset.service";
import {Subject, takeUntil} from "rxjs";

@Component({
  selector: 'app-tab1',
  templateUrl: 'tab1.page.html',
  styleUrls: ['tab1.page.scss']
})
export class Tab1Page {
  public result = '';
  private destroy$: Subject<void> = new Subject<void>();

  constructor(
    private resetService: ResetService
  ) {}

  public resetDatabases(): void {
    this.result = 'Requesting please wait...';
    this.resetService.reset().pipe(takeUntil(this.destroy$)).subscribe(res => {
      this.result = res;
    }, err => {
      this.result = err;
    })
  }

  public seedDatabases(): void {
    this.result = 'Requesting please wait ...';
    this.resetService.seeding().pipe(takeUntil(this.destroy$)).subscribe(res => {
      this.result = res;
    }, err => {
      this.result = err;
    })
  }

}
