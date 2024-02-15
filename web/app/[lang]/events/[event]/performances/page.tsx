"use client"
import Navbar from "@/components/navbar";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Drawer, DrawerClose, DrawerContent, DrawerDescription, DrawerFooter, DrawerHeader, DrawerTitle, DrawerTrigger } from "@/components/ui/drawer";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Performance } from "@/model/performance";
import { Drum, Mail, Phone, Plug, Speaker } from 'lucide-react';
import React from "react";


function formatDate(date: Date): string {
    return `${date.getDate().toString().padStart(2,"0")}.${(date.getMonth()+1).toString().padStart(2,"0")}.${date.getFullYear()}`

}

function formatTime(date: Date): string {
    return `${date.getHours().toString().padStart(2,"0")}:${date.getMinutes().toString().padStart(2,"0")}`
}

function newDate(year: number, month: number, day: number, hour: number, minute: number): Date {
    let date = new Date()

    date.setFullYear(year)
    date.setMonth(month-1)
    date.setDate(day)
    date.setHours(hour)
    date.setMinutes(minute)

    return date
}


export default function EventPerformancesPage() {

    let performances: Performance[] = [
        { id: "abcdef", start: newDate(2024, 12, 1, 9, 0), end: new Date("2024-12-01 19:00"), fee: 50 }
    ]

    const [open, setOpen] = React.useState(false)
    const [selected, setSelected] = React.useState(performances[0])

    return (
        <>
        <Navbar lang="en"/>
        <main className="max-w-3xl mt-6 px-4 mx-auto">
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead>Date</TableHead>
                    <TableHead>Start</TableHead>
                    <TableHead>End</TableHead>
                    <TableHead>Venue</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                { performances.map(performance => 
                <TableRow key={performance.id} onClick={() => {setSelected(performance), setOpen(true)}}>
                    <TableCell>{formatDate(performance.start)}</TableCell>
                    <TableCell>{formatTime(performance.start)}</TableCell>
                    <TableCell>{formatTime(performance.end)}</TableCell>
                    <TableCell>Venue Name</TableCell>
                </TableRow>
                )}
            </TableBody>

        </Table>
        
        <Drawer open={open} onOpenChange={setOpen}>
            <DrawerContent>
                <DrawerHeader>
                <DrawerTitle>Timeslot Information</DrawerTitle>
                <DrawerDescription>This action cannot be undone.</DrawerDescription>
                </DrawerHeader>


                <div className="flex flex-wrap justify-center">

                    <Card className="m-2">
                        <CardHeader>
                            <CardTitle>Venue Name</CardTitle>
                            <CardDescription>Some Street 1, 5400 Baden</CardDescription>
                        </CardHeader>
                        <CardContent>
                            <ul>
                               <li><Phone size={16} className="inline"/><a href="tel:+41 79 420 69 69">+41 79 420 69 69</a></li>
                               <li><Mail size={16} className="inline"/><a href="mailto:someone@venue.ch">someone@venue.ch</a></li>
                            </ul>
                        </CardContent>
                    </Card>

                    <Card className="m-2">
                        <CardHeader>
                            <CardTitle>Equipment</CardTitle>
                            <CardDescription>This is what the venue provides</CardDescription>
                        </CardHeader>
                        <CardContent>
                            <ul>
                               <li><Plug size={16} className="inline"/>Electricity</li>
                               <li><Speaker size={16} className="inline"/>Amplifier</li>
                               <li><Drum size={16} className="inline"/>Drumset</li>
                               
                            </ul>
                        </CardContent>
                    </Card>

                    <Card className="m-2">
                        <CardHeader>
                            <CardTitle>Equipment</CardTitle>
                            <CardDescription>This is what the venue provides</CardDescription>
                        </CardHeader>
                        <CardContent>
                            <ul>
                               <li><Plug size={16} className="inline"/>Electricity</li>
                               <li><Speaker size={16} className="inline"/>Amplifier</li>
                               <li><Drum size={16} className="inline"/>Drumset</li>
                               
                            </ul>
                        </CardContent>
                    </Card>
                </div>
                    

                <DrawerFooter>
                <Button>Apply</Button>
                <DrawerClose>
                    <Button variant="outline">Close</Button>
                </DrawerClose>
                </DrawerFooter>
            </DrawerContent>
        </Drawer>

        </main>
        </>
    )
}