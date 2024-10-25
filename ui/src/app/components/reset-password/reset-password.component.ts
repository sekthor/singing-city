import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ResetPasswordRequest } from 'src/app/models/user';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-reset-password',
  templateUrl: './reset-password.component.html',
  styleUrls: ['./reset-password.component.scss']
})
export class ResetPasswordComponent {

  resetRequest: ResetPasswordRequest = {password:"", code:""}
  error: string = ""

  constructor(
    private userService: UserService,
    private router: Router,
    private route: ActivatedRoute
  ) {}
  
  ngOnInit(): void {
    this.resetRequest.code = this.route.snapshot.queryParamMap.get("code") || ""
    console.log(this.resetRequest)
  }

  reset() {
    this.userService.resetPassword(this.resetRequest).subscribe(
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
