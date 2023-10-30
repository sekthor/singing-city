import { Component, OnInit } from '@angular/core';
import { Venue } from 'src/app/models/venue';
import { VenueService } from 'src/app/services/venue.service';

@Component({
  selector: 'app-venue-list',
  templateUrl: './venue-list.component.html',
  styleUrls: ['./venue-list.component.scss']
})
export class VenueListComponent implements OnInit {

  venues: Venue[] = []
  selectedVenue?: Venue

  constructor(
    private venueService: VenueService
  ) {}

  ngOnInit(): void {
    this.venueService.getVenues().subscribe(venues => this.venues = venues);
  }

  selectVenue(venue: Venue) {
    this.selectedVenue = venue;
  }

  getAvailableSlotCount(venue: Venue) {
    return venue.slots.filter( slot => slot.artistID === 0 ).length
  }

}
