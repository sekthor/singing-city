"use client"
import { MapContainer, Marker, TileLayer } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import Navbar from "@/components/navbar";
import { Venue } from "@/model/venue";
import { Icon } from "leaflet";

const customIcon = new Icon({
    iconUrl:"/pin.png",
    iconSize: [38,38]
})

export default function VenueMapPage() {
    let venues: Venue[] = [
        { id: "abc", name: "Some Bar", location:{longitude:46.79347, latitude:9.82048}}
    ]
    return (
        <>
            <Navbar lang="en"></Navbar>
            <main className="max-w-3xl mt-6 px-4 mx-auto">
            <h1 className="text-xl font-semibold mb-4">Songbird Festival 2024</h1>
            <MapContainer center={[46.80429, 9.83723]} zoom={13} className="h-96">
                <TileLayer
                    attribution="&copy; <a href='https://www.openstreetmap.org/copyright'>OpenStreetMap</a>"
                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                />

                {
                    venues.map(venue => 
                        <Marker 
                            position={[venue.location?.longitude||0, venue.location?.latitude||0]}
                            icon={customIcon}
                        ></Marker>
                    )
                }
            </MapContainer>
            </main>
        </>
    )
}