import { Artist } from "./artist"

export interface Venue {
  ID: number
  name: string
  description: string
  slots: Timeslot[]
  contact: string
  address: string
  zip: number
  city: string
}

export interface Timeslot {
  ID: number
  time: Date
  artist?: Artist
  artistID: number
  venueID: number
  pay: number
}
