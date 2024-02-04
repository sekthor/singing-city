"use client"
import { Select, SelectTrigger, SelectValue, SelectContent, SelectGroup, SelectLabel, SelectItem } from "@/components/ui/select";

const languages = ["en", "de"];

type NavbarProps = {
    lang: string
}

export default function Navbar({ lang }:NavbarProps) {

    return (
        <nav className="p-2 border flex">
            <Select>
                <SelectTrigger className="w-[180px]">
                    <SelectValue placeholder="Select an event" />
                </SelectTrigger>
                <SelectContent>
                    <SelectGroup>
                        <SelectLabel>Fruits</SelectLabel>
                        <SelectItem value="apple">Apple</SelectItem>
                        <SelectItem value="banana">Banana</SelectItem>
                        <SelectItem value="blueberry">Blueberry</SelectItem>
                        <SelectItem value="grapes">Grapes</SelectItem>
                        <SelectItem value="pineapple">Pineapple</SelectItem>
                    </SelectGroup>
                </SelectContent>
            </Select>

            <Select>
                <SelectTrigger className="w-[80px]">
                    <SelectValue placeholder={lang} />
                </SelectTrigger>
                <SelectContent>
                    <SelectGroup>
                        <SelectLabel>Language</SelectLabel>
                        {
                            languages.map(lang =>
                                <SelectItem value={lang}>{lang}</SelectItem>
                            )
                        }
                    </SelectGroup>
                </SelectContent>
            </Select>
        </nav>
    )

}