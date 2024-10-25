import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { VenueListComponent } from './components/venue-list/venue-list.component';
import { RegisterComponent } from './components/register/register.component';
import { VenueDetailComponent } from './components/venue-detail/venue-detail.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { ArtistDetailComponent } from './components/artist-detail/artist-detail.component';
import { authGuard } from './auth.guard';
import { ProfileComponent } from './components/profile/profile.component';
import { AdminComponent } from './components/admin/admin.component';
import { ArtistListComponent } from './components/artist-list/artist-list.component';
import { ResetPasswordComponent } from './components/reset-password/reset-password.component';
import { ForgotPasswordComponent } from './components/forgot-password/forgot-password.component';

const routes: Routes = [
  { path: "", redirectTo: "/dashboard", pathMatch: 'full' },
  { path: "dashboard", component: DashboardComponent, canActivate: [authGuard] },
  { path: "venues", component: VenueListComponent, canActivate: [authGuard] },
  { path: "venues/:id", component: VenueDetailComponent, canActivate: [authGuard] },
  { path: "artists", component: ArtistListComponent, canActivate: [authGuard] },
  { path: "artists/:id", component: ArtistDetailComponent, canActivate: [authGuard] },
  { path: "profile", component: ProfileComponent, canActivate: [authGuard] },
  { path: "admin", component: AdminComponent, canActivate: [authGuard] },
  { path: "login", component: LoginComponent },
  { path: "register", component: RegisterComponent },
  { path: "forgot-password", component: ForgotPasswordComponent },
  { path: "reset-password", component: ResetPasswordComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
