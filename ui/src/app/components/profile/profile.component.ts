import { Component, OnInit } from '@angular/core';
import { Profile } from 'src/app/models/user';
import { ArtistService } from 'src/app/services/artist.service';
import { UserService } from 'src/app/services/user.service';
import { TranslateService } from '@ngx-translate/core';
import { VenueService } from 'src/app/services/venue.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss']
})
export class ProfileComponent implements OnInit {

  profile: Profile
  userSuccess: string = ""
  userError:string = ""
  detailSuccess: string = ""
  detailError: string = ""

  constructor(
    private userService: UserService,
    private artistService: ArtistService,
    private venueService: VenueService,
    private translate: TranslateService
  ) {
    this.profile = {
      user: {ID:0, username: "", type: 0, email: ""},
      venue: {ID:0, name: "", description:"", slots:[], city:"", zip:0, address: "", contact:""},
      artist: {ID:0, name: "",contact:"", genere:""}
    }

    
  }

  ngOnInit(): void {
    this.userService.getProfile().subscribe(
      profile => {
        this.profile = profile
      },
    )
  }

  updateUser() {
    this.userSuccess = ""
    this.userError = ""
    this.userService.updateUser(this.profile.user).subscribe(
      response => {
        this.userSuccess = this.translate.instant("profile.success")
      },
      error => {
        console.log(error)
        this.userError = this.translate.instant("profile.failure")
      }
    )
  }

  updateArtist() {
    this.detailSuccess = ""
    this.detailError = ""
    this.artistService.updateArtist(this.profile.artist).subscribe(
      response => {
        this.detailSuccess = this.translate.instant("profile.success")
      },
      error => {
        console.log(error)
        this.detailError = this.translate.instant("profile.failure")
      }
    )
  }

  updateVenue() {
    this.detailSuccess = ""
    this.detailError = ""
    this.venueService.updateVenue(this.profile.venue).subscribe(
      response => {
        this.detailSuccess = this.translate.instant("profile.success")
      },
      error => {
        console.log(error)
        this.detailError = this.translate.instant("profile.failure")
      }
    )
  }
}
