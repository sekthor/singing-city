"use client"

import { useSearchParams } from 'next/navigation'
import Navbar from "@/components/navbar"
import { Button } from "@/components/ui/button"
import { Event } from "@/model/event"
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { TabsContent } from '@radix-ui/react-tabs'
import { Venue } from '@/model/venue'

import dynamic from 'next/dynamic'

export default function EventDetailPage({ params }: { params: { lang: string }}) {
    
    // we import the venueMap like this, because leaflet wants to access "window"
    // this is not possible with Server Side Rendering (SSR)
    const MapWithNoSSR = dynamic(() => import("@/components/venuemap/venuemap"), {
        ssr: false
    });

    let event: Event = {
        id: "uuid",
        name: "Songbird Festival 2024",
        description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras non malesuada lacus. Nam et eros scelerisque, tincidunt ex quis, posuere nulla. Morbi imperdiet tortor nec justo gravida, at fermentum purus euismod. Integer mollis vitae enim eget auctor. Curabitur lacinia vitae odio at commodo. Proin interdum ligula suscipit ultrices elementum. Ut sed iaculis lacus. Ut tincidunt sodales mauris sed tempor. Aliquam nulla velit, suscipit a egestas eget, suscipit et arcu. In venenatis eu odio quis consectetur. Praesent placerat orci sit amet ligula laoreet, id imperdiet sem dictum. In felis nunc, consectetur sit amet nunc non, venenatis fermentum eros. Integer congue congue turpis, eget dictum turpis tincidunt id. Suspendisse auctor enim vitae rutrum volutpat. Aenean ex lorem, dictum a magna non, dictum finibus mauris. Nam tellus nibh, luctus in bibendum a, sagittis lacinia orci. Donec pellentesque dolor non nibh luctus, a mattis nunc pulvinar. Maecenas a nisi vel neque sodales rhoncus. Ut vel nisi urna. Suspendisse non nunc vehicula, molestie urna a, aliquet enim. Vestibulum lobortis, sem nec vulputate elementum, ipsum magna rhoncus nulla, eu ullamcorper ante lacus vitae dolor. Quisque a felis enim. In augue tortor, placerat quis mattis suscipit, porttitor id orci. Etiam ut varius lacus, vel bibendum lorem. Suspendisse id libero pharetra, lobortis sapien ut, dictum odio. Nulla sed elementum elit, vitae ultrices turpis. Aenean purus nunc, feugiat vel nibh quis, tincidunt aliquet ex. Sed commodo vestibulum risus, at sagittis purus pharetra laoreet. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Phasellus eget felis tortor. Proin a leo eros. Quisque eu volutpat dolor. Nullam vestibulum erat ut dui consequat, non euismod ipsum feugiat. Etiam pretium nisi at molestie accumsan. Ut sed gravida enim, sed vulputate nisl. Duis nec libero condimentum, fermentum quam nec, sodales nunc."
    }

    let venues: Venue[] = [
        { id: "abc", name: "Some Bar", location:{longitude:46.79347, latitude:9.82048}}
    ]

    const searchParams = useSearchParams()
    const invitation = searchParams.get('invitation')

    return (
        <>
        <Navbar lang={params.lang}/>
        <main className="max-w-3xl mt-6 px-4 mx-auto">
            <h1 className="text-xl font-semibold">{event.name}</h1>

            <Tabs defaultValue="info">
            <TabsList className='my-4'>
                <TabsTrigger value="info">Info</TabsTrigger>
                <TabsTrigger value="venues">Venues</TabsTrigger>
                <TabsTrigger value="program">Program</TabsTrigger>
                <TabsTrigger value="timeslots">Timeslots</TabsTrigger>
            </TabsList>

            <TabsContent value="info">

            <p>{event.description}</p>
            {
                invitation ? 
                <Button className="float-right">Accept Invitation</Button> :
                <Button className="float-right"variant="outline">Apply to Join</Button>
            }
            </TabsContent>

            <TabsContent value='venues'>
                <MapWithNoSSR venues={venues}/>
            </TabsContent>

            </Tabs>

                
        </main>
        </>
    )
}