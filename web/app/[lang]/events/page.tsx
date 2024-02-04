"use client"

import Navbar from "@/components/navbar";
import { Table } from "@/components/ui/table";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

import { Event } from "@/model/event";

const events: Event[] = [
    { id: "", name: "Songbird Festival 2024", },
]

export default function({ params }: { params: { lang: string }}) {

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
                </TabsContent>

                <TabsContent value="past">
                </TabsContent>
            </Tabs>
        </main>
       </>
    )
}