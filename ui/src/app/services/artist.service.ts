import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Artist } from '../models/artist';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ArtistService {

  constructor(private http: HttpClient) { }

  getArtistById(id: string): Observable<Artist> {
    return this.http.get<Artist>(`api/artists/${id}`)
  }
  
  updateArtist(artist: Artist): Observable<any> {
    return this.http.put(`/api/artists/${artist.ID}`, artist)
  }

  getArtists(): Observable<Artist[]> {
    return this.http.get<Artist[]>(`/api/artists`)
  }
}
