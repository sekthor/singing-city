import { Artist } from "./artist"
import { Timeslot } from "./venue"

export interface Application {
    ID: number
    artist: Artist
    timeslot: Timeslot
}