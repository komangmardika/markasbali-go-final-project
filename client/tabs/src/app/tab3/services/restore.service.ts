import {Injectable} from "@angular/core";
import {environment} from "../../../environments/environment";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class RestoreService {
  private apiUrl = environment.cliService;
  private token = environment.token
  constructor(private http: HttpClient) { }
  restore(): Observable<any> {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Authorization': `X-TOKEN ${this.token}`
    });

    return this.http.get<any>(`${this.apiUrl}/restore`, { headers });
  }
}
