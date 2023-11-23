import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TranslateService } from '@ngx-translate/core';
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
  translationLink = ""

  constructor(
    private artistService: ArtistService,
    private route: ActivatedRoute,
    private translate: TranslateService) {}


  ngOnInit(): void {
    let id = this.route.snapshot.paramMap.get("id")
    if (id) {
      this.artistID = id
      this.getArtist(id)
    }

    this.translate.onLangChange.subscribe(
      () => this.translationLink = this.getTranslationLink(this.artist?.description || "")
    )
  }

  getArtist(id: string) {
    this.artistService.getArtistById(id).subscribe(
      artist => {
        this.artist = artist;
        this.translationLink = this.getTranslationLink(this.artist.description || "")
      },
      error => console.log(error)
    )
  }

  getTranslationLink(text: string): string {
    if (!text) 
      return text

    let src = "de"
    let dst = "en"

    if (this.translate.currentLang === "de") {
      dst = src
      src = "en"
    }
    
    return `https://translate.google.com/?sl=${src}&tl=${dst}&text=${text}&op=translate`
  }

}
