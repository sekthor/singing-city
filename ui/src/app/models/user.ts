import { Application } from "./application"
import { Artist } from "./artist"
import { Timeslot, Venue } from "./venue"

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
  address?: string
  zip?: number
  city?: string
  phone?: string
  genere?: string
  description?: string
}

export interface ForgotPasswordRequest {
  email: string
}

export interface ResetPasswordRequest {
  password: string
  code: string
}

export interface AuthToken {
  exp: number
  iat: number
  sub: string
  name: string
  type: number
}

export interface AdminInfo {
  artists: Artist[]
  venues: Venue[]
  confirmed: Timeslot[]
  pending: Application[]
  invites: Invite[]
}

export interface Invite {
  invite: string
}
