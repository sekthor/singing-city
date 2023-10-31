import { HttpClient } from '@angular/common/http';
import { Injectable, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthToken, LoginRequest, RegisterRequest } from '../models/user';

import jwt_decode from "jwt-decode";
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class UserService implements OnInit {

  token: AuthToken | null = null;
  loggedIn: boolean = false;

  constructor(
    private cookie: CookieService,
    private http: HttpClient) { }

  ngOnInit(): void {
    //this.getTokenFromCookie()
  }

  getTokenFromCookie() {
    let token = this.cookie.get("Authorization")
    if (token) {
      this.token = jwt_decode(token) 
    }
  }

  login(login: LoginRequest): Observable<any> {
    return this.http.post(`/api/login`, login)
  }

  logout() {
    this.token = null
    this.loggedIn = false
    this.cookie.delete("Authorization")
  }

  isLoggedIn(): boolean {
    this.getTokenFromCookie()
    return this.token !== null
  }

  register(register: RegisterRequest): Observable<any> {
    return this.http.post(`/api/register`, register)
  }

  getSubject(): string {
    if (this.token) {
      return this.token.sub
    }
    return ""
  }

  getName(): string {
    if (this.token) {
      return this.token.name
    }
    return ""
  }

  isArtist(): boolean {
    return this.token?.type === 1
  }
}
