import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { VenueListComponent } from './components/venue-list/venue-list.component';
import { VenueDetailComponent } from './components/venue-detail/venue-detail.component';
import { LoginComponent } from './components/login/login.component';
import { FormsModule } from '@angular/forms';
import { RegisterComponent } from './components/register/register.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { ArtistDetailComponent } from './components/artist-detail/artist-detail.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    VenueListComponent,
    VenueDetailComponent,
    LoginComponent,
    RegisterComponent,
    DashboardComponent,
    ArtistDetailComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
