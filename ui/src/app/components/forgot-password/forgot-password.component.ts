import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ForgotPasswordRequest } from 'src/app/models/user';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-forgot-password',
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.scss']
})
export class ForgotPasswordComponent {
  resetRequest: ForgotPasswordRequest = {email:""}
  error: string = ""

  constructor(
    private userService: UserService,
    private router: Router,
    private route: ActivatedRoute
  ) {}

  requestReset() {
    
    this.userService.forgotPassword(this.resetRequest).subscribe(
        response => {
          document.getElementById("resetsuccess")?.classList.remove("hide")
          document.getElementById("errormsg")?.classList.add("hide")
        },
        error => {
          this.error = error.error.error
          document.getElementById("errormsg")?.classList.remove("hide")
          document.getElementById("resetsuccess")?.classList.add("hide")
        }
    )
  }
}
