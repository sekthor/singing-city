import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Timeslot, Venue } from '../models/venue';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Application } from '../models/application';

@Injectable({
  providedIn: 'root'
})
export class VenueService {

  constructor(
    private http: HttpClient
  ) { }

  getVenues(): Observable<Venue[]> {
    return this.http.get<Venue[]>(`/api/venues`)
  }

  getVenue(id: number): Observable<Venue> {
    return this.http.get<Venue>(`/api/venues/${id}`)
  }

  updateVenue(venue: Venue): Observable<any> {
    return this.http.put(`/api/venues/${venue.ID}`, venue)
  }

  addTimeslot(venueId: number, slot: Timeslot): Observable<any> {
    return this.http.post(`/api/timeslots/venues/${venueId}`, slot)
  }

  deleteTimeslot(venueId: number, tsid: number) {
    return this.http.delete(`/api/timeslots/${tsid}/venues/${venueId}`)
  }

  deleteTimeslotAsAdmin(tsid: number) {
    return this.http.delete(`/api/timeslots/${tsid}`)
  }

  applyForTimeslot(artistId: number, tsid: number) {
    return this.http.post(`api/timeslots/${tsid}/apply/${artistId}`, {})
  }

  getTimeslots(): Observable<Timeslot[]> {
    return this.http.get<Timeslot[]>(`api/timeslots`)
  }
}
