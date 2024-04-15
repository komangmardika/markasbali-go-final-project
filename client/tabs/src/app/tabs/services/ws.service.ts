import {Injectable} from "@angular/core";
import {WebSocketSubject} from "rxjs/internal/observable/dom/WebSocketSubject";
import {webSocket} from "rxjs/webSocket";

@Injectable({
  providedIn: 'root'
})
export class WebSocketService {
  private socket$: WebSocketSubject<any>;

  constructor() {
    // Replace 'ws://localhost:8080/ws' with your WebSocket server URL
    this.socket$ = webSocket('ws://localhost:8080/ws');
  }

  // Method to send a message over WebSocket
  sendMessage(message: any): void {
    this.socket$.next(message);
  }

  // Method to listen for incoming messages
  getMessage(): WebSocketSubject<any> {
    return this.socket$;
  }
}
