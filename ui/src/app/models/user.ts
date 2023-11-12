import { Artist } from "./artist"
import { Venue } from "./venue"

export interface Profile {
  user: UserDTO
  artist: Artist
  venue: Venue
}
export interface UserDTO {
  ID: number
  email: string
  username: string
  type: number
}
export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  password: string
  username: string
  type: number
  name: string
  address: string
  zip: number
  city: string
}

export interface AuthToken {
  exp: number
  iat: number
  sub: string
  name: string
  type: number
}
