import {Component, OnDestroy, OnInit} from '@angular/core';
import {ResetService} from "./services/reset.service";
import {Subject, Subscription, takeUntil} from "rxjs";
import {WebSocketService} from "../tabs/services/ws.service";

@Component({
  selector: 'app-tab1',
  templateUrl: 'tab1.page.html',
  styleUrls: ['tab1.page.scss']
})
export class Tab1Page implements OnInit, OnDestroy{
  public result = '';
  private destroy$: Subject<void> = new Subject<void>();

  alert = false;
  message = '';
  buttonList: string[] = [];

  messageSubscription: Subscription | undefined;
  progressMessage = '';

  constructor(
    private resetService: ResetService,
    private webSocketService: WebSocketService,
  ) {}

  ngOnInit() {
    // Subscribe to incoming messages
    this.messageSubscription = this.webSocketService.getMessage().subscribe(
      (message) => {
        console.log('Received message:', message);
        if(message.type == 'seeding') {
          this.alert = true;
          this.message = 'Currently seeding database ' + message.message.toUpperCase() + ' please wait ...';
          this.buttonList = [];
        }
      },
      (error) => {
        console.error('WebSocket error:', error);
      },
      () => {
        console.log('WebSocket connection closed.');
      }
    );
  }

  ngOnDestroy() {
    // Unsubscribe from the WebSocket subscription when component is destroyed
    if (this.messageSubscription) {
      this.messageSubscription.unsubscribe();
    }
  }

  public async resetDatabases(): Promise<void> {
    let r = await this.resetService.dbLst();
    this.result = 'Requesting please wait...';
    this.resetService.reset().pipe(takeUntil(this.destroy$)).subscribe(res => {
      this.result = res;
      this.buttonList = ['Close'];
      this.alert = true;
      this.message = r.data.length + ' Databases has been reset successfully';
    }, err => {
      this.result = err;
    })
  }

  public async seedDatabases(): Promise<void> {

    let r = await this.resetService.dbLst();

    this.result = 'Requesting please wait ...';
    this.resetService.seeding().pipe(takeUntil(this.destroy$)).subscribe(res => {
      this.result = res;
      this.alert = true;
      this.progressMessage = '';
      this.buttonList = ['Close'];
      this.message = 'Databases has been seeded successfully';
    }, err => {
      this.result = err;
    })
  }


}
