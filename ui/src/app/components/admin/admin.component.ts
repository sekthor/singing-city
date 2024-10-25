import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Application } from 'src/app/models/application';
import { AdminInfo } from 'src/app/models/user';
import { Timeslot } from 'src/app/models/venue';
import { ApplicationService } from 'src/app/services/application.service';
import { UserService } from 'src/app/services/user.service';
import { VenueService } from 'src/app/services/venue.service';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent implements OnInit {

  selectedTab: number = 0
  info: AdminInfo = {
    venues: [],
    artists: [],
    confirmed: [],
    pending: [],
    invites: []
  }

  constructor(
    private userService: UserService,
    private router: Router,
    private venueService: VenueService,
    private applicationService: ApplicationService
    ) { }

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

  deleteTimeslot(ts: Timeslot) {
    this.venueService.deleteTimeslotAsAdmin(ts.ID).subscribe(
      response => {
        this.info.confirmed = this.info.confirmed.filter(slot => ts.ID !== slot.ID)
      },
      error => console.log(error)
    )
  }

  deleteApplication(applicaton: Application) {
    this.applicationService.deleteApplication(applicaton).subscribe(
      response => this.info.pending = this.info.pending.filter(app => applicaton.ID !== app.ID),
      error => console.log(error)
    )
  }

  addInvite() {
    this.userService.addInvite().subscribe(
      invite => this.info.invites.push(invite),
      error => console.log(error)
    )
  }

  sortTimeslots() {
    this.info.confirmed.sort((a,b) =>{
      let c = new Date(a.time).getTime() 
      let d = new Date(b.time).getTime() 
      return c - d
    })
    this.info.pending.sort((a,b) =>{
      let c = new Date(a.timeslot.time).getTime() 
      let d = new Date(b.timeslot.time).getTime() 
      return c - d
    })
  }

}
