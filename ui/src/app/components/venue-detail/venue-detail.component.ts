import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TranslateService } from '@ngx-translate/core';
import { Application } from 'src/app/models/application';
import { Timeslot, Venue } from 'src/app/models/venue';
import { ApplicationService } from 'src/app/services/application.service';
import { UserService } from 'src/app/services/user.service';
import { VenueService } from 'src/app/services/venue.service';

@Component({
  selector: 'app-venue-detail',
  templateUrl: './venue-detail.component.html',
  styleUrls: ['./venue-detail.component.scss']
})
export class VenueDetailComponent implements OnInit {

  venue?: Venue
  isRessourceOwner: boolean = false
  isArtist: boolean = true
  id: number = 0;

  openApplications: Application[] = []

  newTimeslot: Timeslot = { ID:0, time: new Date(), artistID:0, venueID:0, pay:0, private:false, duration:0 }
  newDate: string = ""
  newTime: string = ""

  constructor(
    private venueService: VenueService,
    private route: ActivatedRoute,
    private userService: UserService,
    private applicationService: ApplicationService,
    private translate: TranslateService
  ) {}

  ngOnInit(): void {
    this.translate.use(localStorage.getItem("lang") || "en")
    let id = this.route.snapshot.paramMap.get("id")
    if (id !== null) {
      this.getVenue(parseInt(id))
      this.id = parseInt(id)
    }

    this.isRessourceOwner = (this.userService.getSubject() === id) 
    this.isArtist = this.userService.isArtist()

    if (this.isArtist)
      this.applicationService.getApplications(
        "artist", this.userService.getSubject(), "open").subscribe(
          applications => this.openApplications = applications
        )
  }

  getVenue(id: number) {
    this.venueService.getVenue(id).subscribe(
      venue => {
        this.venue = venue
        //this.sortTimeslots()
      },
      error => {
        console.log(error)
      }
    )
  }

  addTimeSlot() {
    let datestring = `${this.newDate}T${this.newTime}:00`
    this.newTimeslot.time = new Date(`${this.newDate}T${this.newTime}:00`)

    if (isNaN(this.newTimeslot.time.getDate())) {
      console.log("invalid date")
      return
    }

    this.venueService.addTimeslot(this.id, this.newTimeslot).subscribe(
      response => {
        let tsCopy = Object.assign({}, this.newTimeslot)
        this.venue?.slots.push(tsCopy)
      },
      error => {
        console.log(error)
      }
    )  
  }

  deleteTimeslot(slot: Timeslot) {
    this.venueService.deleteTimeslot(this.id, slot.ID).subscribe(
      response => {
        if (this.venue) {
          this.venue.slots = this.venue.slots.filter(s => s.ID !== slot.ID)
        }
      },
      error => {

      }
    )
  }

  sortTimeslots() {
    if (this.venue)
      this.venue?.slots.sort((a: Timeslot, b: Timeslot) => { return a.time.getTime() - b.time.getTime()})
  }

  applyForTimeslot(event: any, slot: Timeslot) {
    let artistId = this.userService.getSubject()
    this.venueService.applyForTimeslot(parseInt(artistId), slot.ID).subscribe(
      response => {
        console.log(response)
      },
      error => {
        console.log(error)
      }
    )
    event.target.disabled = true
  }

  createGoogleMapsQuery(venue: Venue): string {
    let query = encodeURI(`${venue.address} ${venue.zip} ${venue.city}`)
    return `https://www.google.com/maps/search/?api=1&query=${query}`
  }

  hasAlreadyApplied(ts: Timeslot): boolean{
    return this.openApplications
      .find(application => ts.ID === application.timeslot.ID) !== undefined
  }

  getValidTimes(): string[] {
    let times: string[] = []
    for (let hour = 0; hour < 24; hour++) {
      let h = hour.toString().padStart(2, '0')
      times.push(`${h}:00`,`${h}:15`,`${h}:30`,`${h}:45`)
    }
    return times
  }

}
