import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Artist } from 'src/app/models/artist';
import { ArtistService } from 'src/app/services/artist.service';

@Component({
  selector: 'app-artist-detail',
  templateUrl: './artist-detail.component.html',
  styleUrls: ['./artist-detail.component.scss']
})
export class ArtistDetailComponent implements OnInit {

  artist?: Artist
  artistID: string = ""

  constructor(
    private artistService: ArtistService,
    private route: ActivatedRoute) {}


  ngOnInit(): void {
    let id = this.route.snapshot.paramMap.get("id")
    if (id) {
      this.artistID = id
      this.getArtist(id)
    }
  }

  getArtist(id: string) {
    this.artistService.getArtistById(id).subscribe(
      artist => this.artist = artist,
      error => console.log(error)
    )
  }

}
