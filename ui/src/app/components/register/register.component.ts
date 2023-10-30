import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RegisterRequest } from 'src/app/models/user';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent {

  registerRequest: RegisterRequest = { 
    email: "",
    username: "",
    password: "",
    type: 1,
    name: "",
    address: "",
    zip: 0,
    city: ""
  }

  error: string = ""

  constructor(
    private userService: UserService,
    private router: Router) {}

  register() {
    document.getElementById("errormsg")?.classList.add("hide")

    this.userService.register(this.registerRequest).subscribe(
      response => {
        this.router.navigate(["/login"], { queryParams: { origin: 'register' } })
      },
      error => {
          this.error = error.error.error
          document.getElementById("errormsg")?.classList.remove("hide")
      }
    )
  }
}
