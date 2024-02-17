import Navbar from "@/components/navbar";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { ChevronRight, Mic2, Theater } from "lucide-react";
import Link from "next/link";

const lang = "en"
export default function HomePage() {

    return (
        <>
        <Navbar lang={lang}/>
        <main className="p-4 flex flex-wrap space-x-4">
            <Card className="max-w-full">
                <CardHeader>
                    <CardTitle>Events</CardTitle>
                    <CardDescription>Upcomming Events</CardDescription>
                </CardHeader>
                <CardContent>
                    <div className="divide-y divide-solid">

                        <div className="p-1">
                            <Link href={`/${lang}/events/xyz`} className="flex justify-between">
                                <span className="truncate flex-grow">Songbird Festival Davos 2024</span>
                                <ChevronRight className="w-4 flex-grow-0" />
                            </Link>
                        </div>

                        <div className="p-1">
                            <Link href={`/${lang}/events/xyz`} className="flex justify-between">
                                <span className="truncate flex-grow">Some Festival asdfasdfadfafd</span>
                                <ChevronRight className="w-4 flex-grow-0" />
                            </Link>
                        </div>

                        <div className="p-1">
                            <Link href={`/${lang}/events/xyz`} className="flex justify-between">
                                <span className="truncate flex-grow">Music Festival XYZ</span>
                                <ChevronRight className="w-4 flex-grow-0" />
                            </Link>
                        </div>
                    </div>
                </CardContent>
            </Card>

            <Card>
                <CardHeader>
                    <CardTitle>Profiles</CardTitle>
                    <CardDescription>Profiles you manage</CardDescription>
                </CardHeader>
                <CardContent>
                    <div className="divide-y divide-solid">
                        <div className="p-1">
                            <Link href={`/${lang}/venues/xyz`}>
                                <Theater className="inline mr-2" size={16}/>
                                Tom's Cafe
                                <ChevronRight className="float-right ml-2" />
                            </Link>
                        </div>

                        <div className="p-1">
                            <Link href={`/${lang}/venues/xyz`}>
                                <Theater className="inline mr-2" size={16}/>
                                Bob's Burgers 
                                <ChevronRight className="float-right ml-2" />
                            </Link>
                        </div>

                        <div className="p-1">
                            <Link href={`/${lang}/artists/xyz`}>
                                <Mic2 className="inline mr-2" size={16}/>
                                Awesome Band
                                <ChevronRight className="float-right ml-2" />
                            </Link>
                        </div>
                    </div>
                </CardContent>
            </Card>
        </main>
        </>
    )
}