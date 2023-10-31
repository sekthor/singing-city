import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Timeslot, Venue } from 'src/app/models/venue';
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

  newTimeslot: Timeslot = { ID:0, time: new Date(), artistID:0, venueID:0, pay:0, private:false }
  newDate: string = ""
  newTime: string = ""

  constructor(
    private venueService: VenueService,
    private route: ActivatedRoute,
    private userService: UserService
  ) {}

  ngOnInit(): void {
    let id = this.route.snapshot.paramMap.get("id")
    if (id !== null) {
      this.getVenue(parseInt(id))
      this.id = parseInt(id)
    }

    this.isRessourceOwner = (this.userService.getSubject() === id) 
    this.isArtist = this.userService.isArtist()
    this.setDate()
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
    this.newTimeslot.time = new Date(`${this.newDate}T${this.newTime}:00`)

    if (isNaN(+this.newTimeslot.time))
      return

    this.venueService.addTimeslot(this.id, this.newTimeslot).subscribe(
      response => {
        this.venue?.slots.push(this.newTimeslot)
        //this.sortTimeslots()
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

  applyForTimeslot(slot: Timeslot) {
    let artistId = this.userService.getSubject()
    this.venueService.applyForTimeslot(parseInt(artistId), slot.ID).subscribe(
      response => {
        console.log(response)
      },
      error => {
        console.log(error)
      }
    )
  }

  setDate() {
    let datepicker = <HTMLInputElement>document.getElementById("tsdate")

    console.log(datepicker)
    datepicker.value = "2023/12/08"
  }

  createGoogleMapsQuery(venue: Venue): string {
    let query = encodeURI(`${venue.address} ${venue.zip} ${venue.city}`)
    return `https://www.google.com/maps/search/?api=1&query=${query}`
  }

}
