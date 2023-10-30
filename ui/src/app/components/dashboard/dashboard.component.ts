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
      timeslots => this.confirmedTimeslots = timeslots,
      error => console.log(error)
    )
  } 

  getVenueNameByVenueID(venueId: number): string {
    return this.venues.find(venue => venue.ID === venueId)?.name || ""
  }

  getVenueContactByVenueID(venueId: number): string {
    return this.venues.find(venue => venue.ID === venueId)?.contact || ""
  }

  acceptApplication(application: Application) {
    this.applicationService.acceptApplication(application).subscribe(
      response => {
        console.log("success")
      },
      error => {
        console.log(error)
      }
    )

  }

}
