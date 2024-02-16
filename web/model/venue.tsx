import { Location } from "./location";
export interface Venue {
    id: string;
    name: string;
    description?: string;
    location?: Location;
}