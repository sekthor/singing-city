import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { AdminInfo, AuthToken, ForgotPasswordRequest, Invite, LoginRequest, Profile, RegisterRequest, ResetPasswordRequest, UserDTO } from '../models/user';

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

  isAdmin(): boolean {
    if (this.token) {
      return this.token.type == 0
    }
    return false
  }

  register(register: RegisterRequest, invite: string): Observable<any> {
    let params = new HttpParams().set("invite", invite)
    return this.http.post(`/api/register`, register, { params: params })
  }

  addInvite(): Observable<Invite> {
    return this.http.post<Invite>("/api/invites", {})
  }

  getProfile(): Observable<Profile> {
    return this.http.get<Profile>(`/api/profile`)
  }

  updateUser(user: UserDTO): Observable<any> {
    return this.http.put(`/api/users/${user.ID}`, user)
  }

  getAdminInfo(): Observable<AdminInfo> {
    return this.http.get<AdminInfo>(`/api/admin`)
  }

  forgotPassword(req: ForgotPasswordRequest): Observable<any> {
    return this.http.post("/api/forgot-password", req)
  }

  resetPassword(req: ResetPasswordRequest): Observable<any> {
    return this.http.post("/api/reset-password", req)
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
