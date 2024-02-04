import Navbar from "@/components/navbar";
import { Badge } from "@/components/ui/badge";
import { Artist } from "@/model/artist";

export default function ArtistDetailPage() {
    let artist: Artist = {
        id: "uuid",
        name: "Cool Band Name",
        description: "very cool band",
        generes: ["death metal", "metal"]
    }
    return (
        <>
            <Navbar />
            <main className="max-w-3xl mx-auto p-2">
                <h1 className="text-3xl">{artist.name}</h1>
                <div className="flex space-x-2 py-2">
                    {artist.generes.map(genere => <Badge>{genere}</Badge>)}
                </div>
                <p>{artist.description}</p>
            </main>
        </>
    )
}