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

const routes: Routes = [
  { path: "", redirectTo: "/dashboard", pathMatch: 'full' },
  { path: "dashboard", component: DashboardComponent, canActivate: [authGuard] },
  { path: "venues", component: VenueListComponent, canActivate: [authGuard] },
  { path: "venues/:id", component: VenueDetailComponent, canActivate: [authGuard] },
  { path: "artists/:id", component: ArtistDetailComponent, canActivate: [authGuard] },
  { path: "profile", component: ProfileComponent, canActivate: [authGuard] },
  { path: "login", component: LoginComponent },
  { path: "register", component: RegisterComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
