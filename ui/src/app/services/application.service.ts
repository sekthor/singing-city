import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Application } from '../models/application';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApplicationService {

  constructor(
    private http: HttpClient,
  ) { }

  getApplications(userType: string, userID: string, status: string):Observable<Application[]> {
    let params = new HttpParams().set("status", status)
    return this.http.get<Application[]>(`api/applications/${userType}/${userID}`, { params: params })
  }

  acceptApplication(application: Application) {
    return this.http.post(`api/applications/${application.ID}/accept`, {})
  }
}
