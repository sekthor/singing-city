"use client"
import { Select, SelectTrigger, SelectValue, SelectContent, SelectGroup, SelectLabel, SelectItem } from "@/components/ui/select";
import { Avatar, AvatarFallback } from "./ui/avatar";

const languages = ["en", "de"];

type NavbarProps = {
    lang: string
}

type Profile = {
    id: string;
    name: string;
}
type Profiles = {
    artists: Profile[];
    venues: Profile[];
}

function getProfileGroup(profiles: Profile[], label: string) {
        return <SelectGroup>
            <SelectLabel>{label}</SelectLabel>
            { profiles.map(profile => 
                <SelectItem key={profile.id} value={profile.id} className="flex">
                    <div key={profile.id} className="flex items-center">
                        <Avatar className="h-9 w-9">
                            <AvatarFallback>OM</AvatarFallback>
                        </Avatar>

                        <div className="ml-4 space-y-1">
                            <p className="text-sm font-medium leading-none">Olivia Martin</p>
                            <p className="text-sm text-muted-foreground">
                                olivia.martin@email.com
                            </p>
                        </div>
                    </div>
                </SelectItem>
            )}
        </SelectGroup>
}

export default function Navbar({ lang }:NavbarProps) {

    // the profiles the current user manages
    let managedProfiles: Profiles = {
        artists: [
            {id: "1", name: "Solo Project"}, 
            {id: "b", name: "Awesome Band"}],
        venues: [{id:"c", name: "Some Pub"}]
    }
    return (
        <nav className="p-2 border flex">
            <Select>
                <SelectTrigger className="w-[180px]">
                    <SelectValue placeholder="Select an event" />
                </SelectTrigger>
                <SelectContent>
                    {
                        managedProfiles.artists ? getProfileGroup(managedProfiles.artists, "Artists") : <></>
                    }
                    {
                        managedProfiles.venues ? getProfileGroup(managedProfiles.venues, "Venues") : <></>
                    }
                </SelectContent>
            </Select>

            <Select>
                <SelectTrigger className="w-[60px]">
                    <SelectValue placeholder={lang} />
                </SelectTrigger>
                <SelectContent>
                    <SelectGroup>
                        <SelectLabel>Language</SelectLabel>
                        {
                            languages.map(lang =>
                                <SelectItem value={lang} key={lang}>{lang}</SelectItem>
                            )
                        }
                    </SelectGroup>
                </SelectContent>
            </Select>
        </nav>
    )

}