import { Component, OnInit } from '@angular/core';
import { Application } from 'src/app/models/application';
import { Timeslot, Venue } from 'src/app/models/venue';
import { ApplicationService } from 'src/app/services/application.service';
import { UserService } from 'src/app/services/user.service';
import { VenueService } from 'src/app/services/venue.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  isArtist: boolean = false
  openApplications: Application[] = []
  venues: Venue[] = []
  confirmedTimeslots: Timeslot[] = []
  showConfirmed: boolean = false

  constructor(
    private userService: UserService,
    private venueService: VenueService,
    private applicationService: ApplicationService) {}

  ngOnInit(): void {
    this.isArtist = this.userService.isArtist()
    this.getOpenApplications()
    this.getConfirmedPerformances()
  }

  getOpenApplications() {

    let userType = this.isArtist ? "artist" : "venue"

    this.applicationService.getApplications(userType, this.userService.getSubject(), "open").subscribe(
      applications => {
        this.openApplications = applications
        this.sortTimeslots()
      },
      error => {
        console.log(error)
      }
    )
    if (this.isArtist) {
      this.venueService.getVenues().subscribe(
        venues => {
          return this.venues = venues;
        },
        error => console.log(error)
      )
    }
  }

  getConfirmedPerformances() {
    this.venueService.getTimeslots().subscribe(
      timeslots => {
        this.confirmedTimeslots = timeslots
        this.sortTimeslots()
      },
      error => console.log(error)
    )
  } 

  getVenueNameByVenueID(venueId: number): string {
    return this.venues.find(venue => venue.ID === venueId)?.name || ""
  }

  getVenueContactByVenueID(venueId: number): string {
    return this.venues.find(venue => venue.ID === venueId)?.contact || ""
  }

  getVenuePhoneByVenueID(venueId: number): string {
    return this.venues.find(venue => venue.ID === venueId)?.phone || ""
  }

  acceptApplication(application: Application) {
    this.applicationService.acceptApplication(application).subscribe(
      response => {
        this.getConfirmedPerformances()
        this.openApplications = this.openApplications.filter(app => app.ID != application.ID)
      },
      error => {
        console.log(error)
      }
    )
  }

  revokeApplication(application: Application) {
    this.applicationService.deleteApplication(application).subscribe(
      response => {
        this.getConfirmedPerformances()
        this.openApplications = this.openApplications.filter(app => app.ID != application.ID)
      },
      error => {
        console.log(error)
      }
    )
  }


  sortTimeslots() {
    this.openApplications.sort((a,b) => {
      let c = new Date(a.timeslot.time).getTime() 
      let d = new Date(b.timeslot.time).getTime() 
      return c - d
    })

    this.confirmedTimeslots.sort((a,b)=>{
      let c = new Date(a.time).getTime() 
      let d = new Date(b.time).getTime() 
      return c - d
    })
  }

}
