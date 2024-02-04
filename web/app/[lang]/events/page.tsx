"use client"

import Navbar from "@/components/navbar";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

import { Event } from "@/model/event";
import Link from "next/link";

const events: Event[] = [
    { id: "", name: "Songbird Festival 2024", },
]

export default function EventsPage({ params }: { params: { lang: string }}) {

    return (
        <>
        <Navbar lang={params.lang} /> 
        <main className="max-w-3xl mx-auto p-2">
            <h1 className="text-xl my-4">Events</h1>


            <Tabs defaultValue="upcoming" className="w-[400px]">
                <TabsList>
                    <TabsTrigger value="upcoming">Upcoming</TabsTrigger>
                    <TabsTrigger value="past">Past</TabsTrigger>
                </TabsList>

                <TabsContent value="upcoming">
                    <Table>
                        <TableHeader>
                            <TableRow>
                                <TableHead>Event</TableHead>
                                <TableHead>Start</TableHead>
                                <TableHead>End</TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                        {
                            events.map(e => 
                            
                            <TableRow key={e.id}>
                                <TableCell>
                                    <Link href={`/${params.lang}/events/${e.name}`}>{e.name}</Link>
                                </TableCell>
                                <TableCell>{e.start?.toString()}</TableCell>
                                <TableCell>{e.end?.toString()}</TableCell>
                            </TableRow>
                        )}
                        </TableBody>
                    </Table>
                </TabsContent>

                <TabsContent value="past">
                </TabsContent>
            </Tabs>
        </main>
       </>
    )
}