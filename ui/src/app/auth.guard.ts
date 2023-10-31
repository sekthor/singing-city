import { CanActivateFn, Router } from '@angular/router';
import { UserService } from './services/user.service';
import { inject } from '@angular/core';

export const authGuard: CanActivateFn = (route, state) => {
  let userservice = inject(UserService)
  let router = inject(Router)

  if (userservice.isLoggedIn())
    return true

  router.navigate(["/login"]) 
  return false
};
