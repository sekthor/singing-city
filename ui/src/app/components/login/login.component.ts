import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { LoginRequest } from 'src/app/models/user';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  loginRequest: LoginRequest = {email: "", password: ""}
  error: string = ""
  newUser: boolean = false

  constructor(
    private userService: UserService,
    private router: Router,
    private route: ActivatedRoute) {
  }

  ngOnInit(): void {
    this.newUser = (this.route.snapshot.queryParamMap.get("origin") === "register") 
  }

  login() {
    document.getElementById("errormsg")?.classList.add("hide")
    this.userService.login(this.loginRequest).subscribe(
        response => {
          this.router.navigate(["/"]).then(() => window.location.reload())
        },
        error => {
          this.error = error.error.error
          document.getElementById("errormsg")?.classList.remove("hide")
        }
      )
  }

}
