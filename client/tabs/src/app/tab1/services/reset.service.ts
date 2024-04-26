import {Injectable} from "@angular/core";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {environment} from "../../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class ResetService {
  private apiUrl = environment.cliService;
  private token = environment.token
  constructor(private http: HttpClient) { }
  public reset(): Observable<any> {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Authorization': `X-TOKEN ${this.token}`
    });

    return this.http.get<any>(`${this.apiUrl}/reset`, { headers });
  }

  public seeding(): Observable<any> {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Authorization': `X-TOKEN ${this.token}`
    });

    return this.http.get<any>(`${this.apiUrl}/reset/seed`, { headers });
  }

  public async dbLst(): Promise<any> {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Authorization': `X-TOKEN ${this.token}`
    });

    return await this.http.get<any>(`${this.apiUrl}/common/db-list`, { headers }).toPromise();
  }
}
