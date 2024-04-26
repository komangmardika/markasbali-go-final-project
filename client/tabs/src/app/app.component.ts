import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {Subscription} from "rxjs";
import {WebSocketService} from "./tabs/services/ws.service";
import {OverlayEventDetail} from "@ionic/core/components";
import {IonModal} from "@ionic/angular";

@Component({
  selector: 'app-root',
  templateUrl: 'app.component.html',
  styleUrls: ['app.component.scss'],
})
export class AppComponent implements OnInit, OnDestroy{
  messages:string[] = [];
  messageSubscription: Subscription | undefined;
  @ViewChild(IonModal, { static: true }) modal: IonModal | undefined;
  modalOpen = false;

  constructor(private webSocketService: WebSocketService) {}

  ngOnInit() {
    // Subscribe to incoming messages
    this.messageSubscription = this.webSocketService.getMessage().subscribe(
      (message) => {
        console.log('Received message:', message);
        if(message.type == 'error') {
          this.messages.push(message.message);
          this.modalOpen = true
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

  cancel() {
    if(this.modal) {
      this.modal.dismiss(null, 'cancel');
      this.messages = []
      this.modalOpen = true;
    }
  }

  onWillDismiss(event: Event) {
    const ev = event as CustomEvent<OverlayEventDetail<string>>;
    this.modalOpen = false;
  }
}
