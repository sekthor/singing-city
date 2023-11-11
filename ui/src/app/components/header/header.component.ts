import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TranslateService } from '@ngx-translate/core';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  isLoggedIn: boolean = false
  language: string = ""

  constructor(
    private userService: UserService, 
    private router: Router, 
    private translate: TranslateService) { }

  ngOnInit(): void {
    this.isLoggedIn = this.userService.isLoggedIn()
    this.language = this.translate.currentLang
  }

  logout() {
    this.userService.logout() 
    this.isLoggedIn = false
    this.router.navigate(["/"]).then(() => window.location.reload())
  }

  setLanguage(language: string) {
    this.translate.use(language)
  }
}
