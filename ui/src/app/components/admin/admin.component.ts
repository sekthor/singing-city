import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AdminInfo } from 'src/app/models/user';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent implements OnInit {

  selectedTab: number = 0
  info : AdminInfo = {
    venues: [],
    artists: [],
    confirmed: [],
    pending: []
  }

  constructor(
    private userService: UserService,
    private router: Router){}

  ngOnInit(): void {
    if (!this.userService.isAdmin()) {
      this.router.navigate(["/dashboard"])
    }

    this.userService.getAdminInfo().subscribe(
      info => this.info = info,
      error => console.log(error)
    )
  }

  findVenueName(id: number): string {
    return this.info.venues.find(venue => id === venue.ID)?.name || ""
  }

  findAristName(id: number): string {
    return this.info.artists.find(artist => id === artist.ID)?.name || ""
  }

}
