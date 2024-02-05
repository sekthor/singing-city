import Navbar from "@/components/navbar";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Carousel, CarouselContent, CarouselItem, CarouselNext, CarouselPrevious } from "@/components/ui/carousel";
import { Artist } from "@/model/artist";

export default function ArtistDetailPage({ params }: { params: { lang: string }}) {
    let artist: Artist = {
        id: "uuid",
        name: "Cool Band Name",
        description: "very cool band",
        generes: ["death metal", "metal"],
        images: [
            "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/ACDC_-_Toronto_November_7%2C_2008.JPG/1920px-ACDC_-_Toronto_November_7%2C_2008.JPG",
            "https://upload.wikimedia.org/wikipedia/commons/8/80/ACDC_Tacoma.jpg"
        ]
    }
    return (
        <>
            <Navbar lang={params.lang} />
            <main className="max-w-3xl mx-auto mt-8 p-2">
                <h1 className="text-3xl">{artist.name}</h1>
                <div className="flex space-x-2 py-2">
                    {artist.generes.map(genere => <Badge>{genere}</Badge>)}
                </div>
                <Carousel className="my-6">
                    <CarouselContent>
                        {
                            artist.images.map(image => (
                                <CarouselItem key={image} className="flex justify-center">
                                    <img src={image} alt="image" className="rounded-lg h-72 object-cover"/>
                                </CarouselItem>
                            )
                        )}
                    </CarouselContent>
                    <CarouselPrevious />
                    <CarouselNext />
                </Carousel>
                <p>{artist.description}</p>
            </main>
        </>
    )
}