import { Artist } from "./artist"

export interface Venue {
  ID: number
  CreatedAt?: string
  name: string
  description: string
  slots: Timeslot[]
  contact: string
  address: string
  zip: number
  city: string
  phone?: string
}

export interface Timeslot {
  ID: number
  time: Date
  artist?: Artist
  artistID: number
  venueID: number
  pay: number
  private: boolean
  duration: number
}
