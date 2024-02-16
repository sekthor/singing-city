"use client"
import { Venue } from "@/model/venue"

import { MapContainer, Marker, Popup, TileLayer } from "react-leaflet"
import { Icon } from "leaflet"
import "leaflet/dist/leaflet.css";
import { useEffect, useState } from "react";


const customIcon = new Icon({
    iconUrl:"/pin.png",
    iconSize: [38,38]
})

type VenueMapProps = {
    venues: Venue[]
}

export default function VenueMap({venues}: VenueMapProps) {

    return (
        <>
        <MapContainer id="map" center={[46.80429, 9.83723]} zoom={13} attributionControl={false} className="h-96">
            <TileLayer
                attribution="&copy; <a href='https://www.openstreetmap.org/copyright'>OpenStreetMap</a>"
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />

            {
                venues.map(venue => 
                    <Marker 
                        position={[venue.location?.longitude||0, venue.location?.latitude||0]}
                        icon={customIcon}
                    >
                        <Popup>{venue.name}</Popup>
                    </Marker>
                )
            }
        </MapContainer>
        <label htmlFor="map" className="text-[0.5rem] float-right">Leaflet & Openstreetmap</label>
        </>
    )
}